package kopia

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/compression"
	"github.com/kopia/kopia/repo/content"
	"github.com/kopia/kopia/repo/format"
	"github.com/kopia/kopia/repo/maintenance"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/kopia/retention"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	defaultKopiaConfigDir   = "/tmp/"
	kopiaConfigFileTemplate = "repository-%s.config"
	defaultCompressor       = "zstd-better-compression"
	// Interval of 0 disables scheduling.
	defaultSchedulingInterval = time.Second * 0
)

var (
	ErrSettingDefaultConfig = clues.New("setting default repo config values")
	ErrorRepoAlreadyExists  = clues.New("repo already exists")
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

type (
	manifestFinder interface {
		FindManifests(
			ctx context.Context,
			tags map[string]string,
		) ([]*manifest.EntryMetadata, error)
	}

	snapshotManager interface {
		manifestFinder
		LoadSnapshot(
			ctx context.Context,
			id manifest.ID,
		) (*snapshot.Manifest, error)
	}

	snapshotLoader interface {
		SnapshotRoot(man *snapshot.Manifest) (fs.Entry, error)
	}
)

var (
	_ snapshotManager = &conn{}
	_ snapshotLoader  = &conn{}
)

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

func (w *conn) Initialize(
	ctx context.Context,
	opts repository.Options,
	retentionOpts repository.Retention,
	repoNameHash string,
) error {
	bst, err := blobStoreByProvider(ctx, opts, w.storage)
	if err != nil {
		return clues.Wrap(err, "initializing storage")
	}
	defer bst.Close(ctx)

	cfg, err := w.storage.CommonConfig()
	if err != nil {
		return clues.StackWC(ctx, err)
	}

	rOpts := retention.NewOpts()
	if err := rOpts.Set(retentionOpts); err != nil {
		return clues.WrapWC(ctx, err, "setting retention configuration")
	}

	blobCfg, _, err := rOpts.AsConfigs(ctx)
	if err != nil {
		return clues.Stack(err)
	}

	// Minimal config for retention if caller requested it.
	kopiaOpts := repo.NewRepositoryOptions{
		RetentionMode:   blobCfg.RetentionMode,
		RetentionPeriod: blobCfg.RetentionPeriod,
	}

	if err = repo.Initialize(ctx, bst, &kopiaOpts, cfg.CorsoPassphrase); err != nil {
		if errors.Is(err, repo.ErrAlreadyInitialized) {
			return clues.StackWC(ctx, ErrorRepoAlreadyExists, err)
		}

		return clues.WrapWC(ctx, err, "initializing repo")
	}

	err = w.commonConnect(
		ctx,
		opts,
		cfg.KopiaCfgDir,
		repoNameHash,
		bst,
		cfg.CorsoPassphrase,
		defaultCompressor)
	if err != nil {
		return err
	}

	if err := w.setDefaultConfigValues(ctx); err != nil {
		return clues.StackWC(ctx, err)
	}

	// Calling with all parameters here will set extend object locks for
	// maintenance. Parameters for actual retention should have been set during
	// initialization and won't be updated again.
	return clues.Stack(w.setRetentionParameters(ctx, retentionOpts)).OrNil()
}

func (w *conn) Connect(ctx context.Context, opts repository.Options, repoNameHash string) error {
	bst, err := blobStoreByProvider(ctx, opts, w.storage)
	if err != nil {
		return clues.Wrap(err, "initializing storage")
	}
	defer bst.Close(ctx)

	cfg, err := w.storage.CommonConfig()
	if err != nil {
		return clues.StackWC(ctx, err)
	}

	return w.commonConnect(
		ctx,
		opts,
		cfg.KopiaCfgDir,
		repoNameHash,
		bst,
		cfg.CorsoPassphrase,
		defaultCompressor)
}

func (w *conn) commonConnect(
	ctx context.Context,
	opts repository.Options,
	configDir string,
	repoNameHash string,
	bst blob.Storage,
	password, compressor string,
) error {
	kopiaOpts := &repo.ConnectOptions{
		ClientOptions: repo.ClientOptions{
			Username: opts.User,
			Hostname: opts.Host,
			ReadOnly: opts.ReadOnly,
		},
	}

	if len(configDir) > 0 {
		kopiaOpts.CachingOptions = content.CachingOptions{
			CacheDirectory: configDir,
		}
	} else {
		configDir = defaultKopiaConfigDir
	}

	cfgFile := filepath.Join(configDir, fmt.Sprintf(kopiaConfigFileTemplate, repoNameHash))

	// todo - issue #75: nil here should be storage.ConnectOptions()
	if err := repo.Connect(
		ctx,
		cfgFile,
		bst,
		password,
		kopiaOpts); err != nil {
		return clues.WrapWC(ctx, err, "connecting to kopia repo")
	}

	if err := w.open(ctx, cfgFile, password); err != nil {
		return clues.StackWC(ctx, err)
	}

	return nil
}

func blobStoreByProvider(
	ctx context.Context,
	opts repository.Options,
	s storage.Storage,
) (blob.Storage, error) {
	switch s.Provider {
	case storage.ProviderS3:
		return s3BlobStorage(ctx, opts, s)
	case storage.ProviderFilesystem:
		return filesystemStorage(ctx, opts, s)
	default:
		return nil, clues.NewWC(ctx, "storage provider details are required")
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

	if err != nil {
		return clues.WrapWC(ctx, err, "closing repository connection")
	}

	return nil
}

func (w *conn) open(ctx context.Context, configPath, password string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.refCount++

	// TODO(ashmrtnz): issue #75: nil here should be storage.ConnectionOptions().
	rep, err := repo.Open(ctx, configPath, password, nil)
	if err != nil {
		return clues.WrapWC(ctx, err, "opening repository connection")
	}

	w.Repository = rep

	return nil
}

func (w *conn) wrap() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.refCount == 0 {
		return clues.New("conn not established or already closed")
	}

	w.refCount++

	return nil
}

func (w *conn) setDefaultConfigValues(ctx context.Context) error {
	p, err := w.getGlobalPolicyOrEmpty(ctx)
	if err != nil {
		return clues.Stack(ErrSettingDefaultConfig, err)
	}

	changed, err := updateCompressionOnPolicy(defaultCompressor, p)
	if err != nil {
		return clues.Stack(ErrSettingDefaultConfig, err)
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

	if err := w.writeGlobalPolicy(ctx, "UpdateGlobalPolicyWithDefaults", p); err != nil {
		return clues.Wrap(err, "updating global policy with defaults")
	}

	return nil
}

// Compression attempts to set the global compression policy for the kopia repo
// to the given compressor.
func (w *conn) Compression(ctx context.Context, compressor string) error {
	// Redo this check so we can exit without looking up a policy if a bad
	// compressor was given.
	comp := compression.Name(compressor)
	if err := checkCompressor(comp); err != nil {
		return clues.StackWC(ctx, err)
	}

	p, err := w.getGlobalPolicyOrEmpty(ctx)
	if err != nil {
		return err
	}

	changed, err := updateCompressionOnPolicy(compressor, p)
	if err != nil {
		return clues.StackWC(ctx, err)
	}

	if !changed {
		return nil
	}

	if err := w.writeGlobalPolicy(ctx, "UpdateGlobalCompressionPolicy", p); err != nil {
		return clues.Wrap(err, "updating global compression policy")
	}

	return nil
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

func updateRetentionOnPolicy(retPolicy policy.RetentionPolicy, p *policy.Policy) bool {
	if retPolicy == p.RetentionPolicy {
		return false
	}

	p.RetentionPolicy = retPolicy

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

		return nil, clues.WrapWC(ctx, err, "getting backup policy").With("source_info", si)
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
	ctx = clues.Add(ctx, "source_info", si)

	writeOpts := repo.WriteSessionOptions{Purpose: purpose}
	ctr := func(ictx context.Context, rw repo.RepositoryWriter) error {
		if err := policy.SetPolicy(ctx, rw, si, p); err != nil {
			return clues.StackWC(ictx, err)
		}

		return nil
	}

	if err := repo.WriteSession(ctx, w.Repository, writeOpts, ctr); err != nil {
		return clues.WrapWC(ctx, err, "updating policy")
	}

	return nil
}

func checkCompressor(compressor compression.Name) error {
	for c := range compression.ByName {
		if c == compressor {
			return nil
		}
	}

	return clues.Stack(clues.New("unknown compressor type"), clues.New(string(compressor)))
}

func (w *conn) setRetentionParameters(
	ctx context.Context,
	rrOpts repository.Retention,
) error {
	if rrOpts.Mode == nil && rrOpts.Duration == nil && rrOpts.Extend == nil {
		return nil
	}

	// Somewhat confusing case, when we have no retention but a non-zero duration
	// it acts like we passed in only the duration and returns an error about
	// having to set both. Return a clearer error here instead.
	if ptr.Val(rrOpts.Mode) == repository.NoRetention && ptr.Val(rrOpts.Duration) != 0 {
		return clues.NewWC(ctx, "duration must be 0 if rrOpts is disabled")
	}

	dr, ok := w.Repository.(repo.DirectRepository)
	if !ok {
		return clues.NewWC(ctx, "getting handle to repo")
	}

	blobCfg, params, err := getRetentionConfigs(ctx, dr)
	if err != nil {
		return clues.Stack(err)
	}

	opts := retention.OptsFromConfigs(*blobCfg, *params)
	if err := opts.Set(rrOpts); err != nil {
		return clues.StackWC(ctx, err)
	}

	return clues.Stack(persistRetentionConfigs(ctx, dr, opts)).OrNil()
}

func getRetentionConfigs(
	ctx context.Context,
	dr repo.DirectRepository,
) (*format.BlobStorageConfiguration, *maintenance.Params, error) {
	blobCfg, err := dr.FormatManager().BlobCfgBlob()
	if err != nil {
		return nil, nil, clues.WrapWC(ctx, err, "getting storage config")
	}

	params, err := maintenance.GetParams(ctx, dr)
	if err != nil {
		return nil, nil, clues.WrapWC(ctx, err, "getting maintenance config")
	}

	return &blobCfg, params, nil
}

func persistRetentionConfigs(
	ctx context.Context,
	dr repo.DirectRepository,
	opts *retention.Opts,
) error {
	// Persist changes.
	if !opts.BlobChanged() && !opts.ParamsChanged() {
		return nil
	}

	blobCfg, params, err := opts.AsConfigs(ctx)
	if err != nil {
		return clues.Stack(err)
	}

	mp, err := dr.FormatManager().GetMutableParameters()
	if err != nil {
		return clues.WrapWC(ctx, err, "getting mutable parameters")
	}

	requiredFeatures, err := dr.FormatManager().RequiredFeatures()
	if err != nil {
		return clues.WrapWC(ctx, err, "getting required features")
	}

	// Must be the case that only blob changed.
	if !opts.ParamsChanged() {
		return clues.WrapWC(
			ctx,
			dr.FormatManager().SetParameters(ctx, mp, blobCfg, requiredFeatures),
			"persisting storage config").
			OrNil()
	}

	// Both blob and maintenance changed. A DirectWriteSession is required to
	// update the maintenance config but not the blob config.
	err = repo.DirectWriteSession(
		ctx,
		dr,
		repo.WriteSessionOptions{
			Purpose: "Corso immutable backups config",
		},
		func(ctx context.Context, dw repo.DirectRepositoryWriter) error {
			// Set the maintenance config first as we can bail out of the write
			// session later.
			if err := maintenance.SetParams(ctx, dw, &params); err != nil {
				return clues.WrapWC(ctx, err, "maintenance config")
			}

			if !opts.BlobChanged() {
				return nil
			}

			return clues.WrapWC(
				ctx,
				dr.FormatManager().SetParameters(ctx, mp, blobCfg, requiredFeatures),
				"storage config").
				OrNil()
		})

	return clues.WrapWC(ctx, err, "persisting config changes").OrNil()
}

func (w *conn) LoadSnapshot(
	ctx context.Context,
	id manifest.ID,
) (*snapshot.Manifest, error) {
	man, err := snapshot.LoadSnapshot(ctx, w.Repository, id)
	if err != nil {
		return nil, clues.StackWC(ctx, err)
	}

	return man, nil
}

func (w *conn) SnapshotRoot(man *snapshot.Manifest) (fs.Entry, error) {
	return snapshotfs.SnapshotRoot(w.Repository, man)
}

func (w *conn) UpdatePassword(
	ctx context.Context,
	password string,
	opts repository.Options,
	repoNameHash string,
) error {
	if len(password) <= 0 {
		return clues.New("empty password provided")
	}

	kopiaRef := NewConn(w.storage)
	if err := kopiaRef.Connect(ctx, opts, repoNameHash); err != nil {
		return clues.Wrap(err, "connecting kopia client")
	}

	defer kopiaRef.Close(ctx)

	kopiaRepo := kopiaRef.Repository.(repo.DirectRepository)
	if err := kopiaRepo.FormatManager().ChangePassword(ctx, password); err != nil {
		return clues.Wrap(err, "unable to update password")
	}

	return nil
}
