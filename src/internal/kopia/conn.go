package kopia

import (
	"context"
	"sync"

	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/compression"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/storage"
)

const (
	defaultKopiaConfigFilePath = "/tmp/repository.config"
)

var (
	errInit    = errors.New("initializing repo")
	errConnect = errors.New("connecting repo")
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
	if err = repo.Initialize(ctx, bst, nil, cfg.CorsoPassword); err != nil {
		if errors.Is(err, repo.ErrAlreadyInitialized) {
			return RepoAlreadyExistsError(err)
		}
		return errors.Wrap(err, errInit.Error())
	}

	// todo - issue #75: nil here should be a storage.ConnectOptions()
	if err := repo.Connect(
		ctx,
		defaultKopiaConfigFilePath,
		bst,
		cfg.CorsoPassword,
		nil,
	); err != nil {
		return errors.Wrap(err, errConnect.Error())
	}

	return w.open(ctx, cfg.CorsoPassword)
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

	// todo - issue #75: nil here should be storage.ConnectOptions()
	if err := repo.Connect(
		ctx,
		defaultKopiaConfigFilePath,
		bst,
		cfg.CorsoPassword,
		nil,
	); err != nil {
		return errors.Wrap(err, errConnect.Error())
	}

	return w.open(ctx, cfg.CorsoPassword)
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

func (w *conn) open(ctx context.Context, password string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.refCount++

	// TODO(ashmrtnz): issue #75: nil here should be storage.ConnectionOptions().
	rep, err := repo.Open(ctx, defaultKopiaConfigFilePath, password, nil)
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

// Compression attempts to set the global compression policy for the kopia repo
// to the given compressor.
func (w *conn) Compression(ctx context.Context, compressor string) error {
	comp := compression.Name(compressor)

	if err := checkCompressor(comp); err != nil {
		return err
	}

	si := policy.GlobalPolicySourceInfo

	p, err := w.getPolicyOrEmpty(ctx, si)
	if err != nil {
		return err
	}

	if compressor == string(p.CompressionPolicy.CompressorName) {
		return nil
	}

	p.CompressionPolicy = policy.CompressionPolicy{
		CompressorName: compression.Name(comp),
	}

	err = repo.WriteSession(
		ctx,
		w.Repository,
		repo.WriteSessionOptions{Purpose: "UpdateGlobalCompressionPolicy"},
		func(innerCtx context.Context, rw repo.RepositoryWriter) error {
			return policy.SetPolicy(ctx, rw, si, p)
		},
	)

	return errors.Wrap(err, "updating global compression policy")
}

func (w *conn) getPolicyOrEmpty(ctx context.Context, si snapshot.SourceInfo) (*policy.Policy, error) {
	p, err := policy.GetDefinedPolicy(ctx, w.Repository, si)
	if err != nil {
		if errors.Is(err, policy.ErrPolicyNotFound) {
			return &policy.Policy{}, nil
		}

		return nil, errors.Wrap(err, "getting global backup policy")
	}

	return p, nil
}

func checkCompressor(compressor compression.Name) error {
	for c := range compression.ByName {
		if c == compressor {
			return nil
		}
	}

	return errors.Errorf("unknown compressor type %s", compressor)
}
