package kopia

import (
	"context"
	"path/filepath"
	"sync"
	"time"

	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/compression"
	"github.com/kopia/kopia/repo/content"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	defaultKopiaConfigDir  = "/tmp/"
	defaultKopiaConfigFile = "repository.config"
	defaultCompressor      = "s2-default"
	// Interval of 0 disables scheduling.
	defaultSchedulingInterval = time.Second * 0
)

const defaultConfigErrTmpl = "setting default repo config values"

var (
	errInit    = errors.New("initializing repo")
	errConnect = errors.New("connecting repo")
)

// Having all fields set to 0 causes it to keep max-int versions of snapshots.
var (
	zeroOpt          = policy.OptionalInt(0)
	defaultRetention = policy.RetentionPolicy{
		KeepLatest:  &zeroOpt,
		KeepHourly:  &zeroOpt,
		KeepWeekly:  &zeroOpt,
		KeepDaily:   &zeroOpt,
		KeepMonthly: &zeroOpt,
		KeepAnnual:  &zeroOpt,
	}
)

type ErrorRepoAlreadyExists struct {
	common.Err
}

func RepoAlreadyExistsError(e error) error {
	return ErrorRepoAlreadyExists{*common.EncapsulateError(e)}
}

func IsRepoAlreadyExistsError(e error) bool {
	var erae ErrorRepoAlreadyExists
	return errors.As(e, &erae)
}

var _ snapshotManager = &conn{}

type conn struct {
	storage storage.Storage
	repo.Repository
	mu       sync.Mutex
	refCount int
}

func NewConn(s storage.Storage) *conn {
	return &conn{
		storage: s,
	}
}

func (w *conn) Initialize(ctx context.Context) error {
	bst, err := blobStoreByProvider(ctx, w.storage)
	if err != nil {
		return errors.Wrap(err, errInit.Error())
	}
	defer bst.Close(ctx)

	cfg, err := w.storage.CommonConfig()
	if err != nil {
		return err
	}

	// todo - issue #75: nil here should be a storage.NewRepoOptions()
	if err = repo.Initialize(ctx, bst, nil, cfg.CorsoPassphrase); err != nil {
		if errors.Is(err, repo.ErrAlreadyInitialized) {
			return RepoAlreadyExistsError(err)
		}

		return errors.Wrap(err, errInit.Error())
	}

	return w.commonConnect(
		ctx,
		cfg.KopiaCfgDir,
		bst,
		cfg.CorsoPassphrase,
		defaultCompressor,
	)
}

func (w *conn) Connect(ctx context.Context) error {
	bst, err := blobStoreByProvider(ctx, w.storage)
	if err != nil {
		return errors.Wrap(err, errInit.Error())
	}
	defer bst.Close(ctx)

	cfg, err := w.storage.CommonConfig()
	if err != nil {
		return err
	}

	return w.commonConnect(
		ctx,
		cfg.KopiaCfgDir,
		bst,
		cfg.CorsoPassphrase,
		defaultCompressor,
	)
}

func (w *conn) commonConnect(
	ctx context.Context,
	configDir string,
	bst blob.Storage,
	password, compressor string,
) error {
	var opts *repo.ConnectOptions
	if len(configDir) > 0 {
		opts = &repo.ConnectOptions{
			CachingOptions: content.CachingOptions{
				CacheDirectory: configDir,
			},
		}
	} else {
		configDir = defaultKopiaConfigDir
	}

	cfgFile := filepath.Join(configDir, defaultKopiaConfigFile)

	// todo - issue #75: nil here should be storage.ConnectOptions()
	if err := repo.Connect(
		ctx,
		cfgFile,
		bst,
		password,
		opts,
	); err != nil {
		return errors.Wrap(err, errConnect.Error())
	}

	if err := w.open(ctx, cfgFile, password); err != nil {
		return err
	}

	return w.setDefaultConfigValues(ctx)
}

func blobStoreByProvider(ctx context.Context, s storage.Storage) (blob.Storage, error) {
	switch s.Provider {
	case storage.ProviderS3:
		return s3BlobStorage(ctx, s)
	default:
		return nil, errors.New("storage provider details are required")
	}
}

func (w *conn) Close(ctx context.Context) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.refCount == 0 {
		return nil
	}

	w.refCount--

	if w.refCount > 0 {
		return nil
	}

	return w.close(ctx)
}

// close closes the kopia handle. Safe to run without the mutex because other
// functions check only the refCount variable.
func (w *conn) close(ctx context.Context) error {
	err := w.Repository.Close(ctx)
	w.Repository = nil

	return errors.Wrap(err, "closing repository connection")
}

func (w *conn) open(ctx context.Context, configPath, password string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.refCount++

	// TODO(ashmrtnz): issue #75: nil here should be storage.ConnectionOptions().
	rep, err := repo.Open(ctx, configPath, password, nil)
	if err != nil {
		return errors.Wrap(err, "opening repository connection")
	}

	w.Repository = rep

	return nil
}

func (w *conn) wrap() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.refCount == 0 {
		return errors.New("conn already closed")
	}

	w.refCount++

	return nil
}

func (w *conn) setDefaultConfigValues(ctx context.Context) error {
	p, err := w.getGlobalPolicyOrEmpty(ctx)
	if err != nil {
		return errors.Wrap(err, defaultConfigErrTmpl)
	}

	changed, err := updateCompressionOnPolicy(defaultCompressor, p)
	if err != nil {
		return errors.Wrap(err, defaultConfigErrTmpl)
	}

	if updateRetentionOnPolicy(defaultRetention, p) {
		changed = true
	}

	if updateSchedulingOnPolicy(defaultSchedulingInterval, p) {
		changed = true
	}

	if !changed {
		return nil
	}

	return errors.Wrap(
		w.writeGlobalPolicy(ctx, "UpdateGlobalPolicyWithDefaults", p),
		"updating global policy with defaults",
	)
}

// Compression attempts to set the global compression policy for the kopia repo
// to the given compressor.
func (w *conn) Compression(ctx context.Context, compressor string) error {
	// Redo this check so we can exit without looking up a policy if a bad
	// compressor was given.
	comp := compression.Name(compressor)
	if err := checkCompressor(comp); err != nil {
		return err
	}

	p, err := w.getGlobalPolicyOrEmpty(ctx)
	if err != nil {
		return err
	}

	changed, err := updateCompressionOnPolicy(compressor, p)
	if err != nil {
		return err
	}

	if !changed {
		return nil
	}

	return errors.Wrap(
		w.writeGlobalPolicy(ctx, "UpdateGlobalCompressionPolicy", p),
		"updating global compression policy",
	)
}

func updateCompressionOnPolicy(compressor string, p *policy.Policy) (bool, error) {
	comp := compression.Name(compressor)

	if err := checkCompressor(comp); err != nil {
		return false, err
	}

	if comp == p.CompressionPolicy.CompressorName {
		return false, nil
	}

	p.CompressionPolicy = policy.CompressionPolicy{
		CompressorName: comp,
	}

	return true, nil
}

func updateRetentionOnPolicy(retention policy.RetentionPolicy, p *policy.Policy) bool {
	if retention == p.RetentionPolicy {
		return false
	}

	p.RetentionPolicy = retention

	return true
}

func updateSchedulingOnPolicy(
	interval time.Duration,
	p *policy.Policy,
) bool {
	if p.SchedulingPolicy.Interval() == interval {
		return false
	}

	p.SchedulingPolicy.SetInterval(interval)

	return true
}

func (w *conn) getGlobalPolicyOrEmpty(ctx context.Context) (*policy.Policy, error) {
	si := policy.GlobalPolicySourceInfo
	return w.getPolicyOrEmpty(ctx, si)
}

func (w *conn) getPolicyOrEmpty(ctx context.Context, si snapshot.SourceInfo) (*policy.Policy, error) {
	p, err := policy.GetDefinedPolicy(ctx, w.Repository, si)
	if err != nil {
		if errors.Is(err, policy.ErrPolicyNotFound) {
			return &policy.Policy{}, nil
		}

		return nil, errors.Wrapf(err, "getting backup policy for %+v", si)
	}

	return p, nil
}

func (w *conn) writeGlobalPolicy(
	ctx context.Context,
	purpose string,
	p *policy.Policy,
) error {
	si := policy.GlobalPolicySourceInfo
	return w.writePolicy(ctx, purpose, si, p)
}

func (w *conn) writePolicy(
	ctx context.Context,
	purpose string,
	si snapshot.SourceInfo,
	p *policy.Policy,
) error {
	err := repo.WriteSession(
		ctx,
		w.Repository,
		repo.WriteSessionOptions{Purpose: purpose},
		func(innerCtx context.Context, rw repo.RepositoryWriter) error {
			return policy.SetPolicy(ctx, rw, si, p)
		},
	)

	return errors.Wrapf(err, "updating policy for %+v", si)
}

func checkCompressor(compressor compression.Name) error {
	for c := range compression.ByName {
		if c == compressor {
			return nil
		}
	}

	return errors.Errorf("unknown compressor type %s", compressor)
}

func (w *conn) LoadSnapshots(
	ctx context.Context,
	ids []manifest.ID,
) ([]*snapshot.Manifest, error) {
	return snapshot.LoadSnapshots(ctx, w.Repository, ids)
}
