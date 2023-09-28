package utils

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

type RepoDetailsAndOpts struct {
	Repo config.RepoDetails
	Opts control.Options
}

var ErrNotYetImplemented = clues.New("not yet implemented")

// GetAccountAndConnect is a wrapper for GetAccountAndConnectWithOverrides
// that automatically gets the storage provider and any storage provider specific
// flag overrides from the command line.
func GetAccountAndConnect(
	ctx context.Context,
	cmd *cobra.Command,
	pst path.ServiceType,
) (repository.Repositoryer, RepoDetailsAndOpts, error) {
	provider, overrides, err := GetStorageProviderAndOverrides(ctx, cmd)
	if err != nil {
		return nil, RepoDetailsAndOpts{}, clues.Stack(err)
	}

	return GetAccountAndConnectWithOverrides(ctx, pst, provider, overrides)
}

func GetAccountAndConnectWithOverrides(
	ctx context.Context,
	pst path.ServiceType,
	provider storage.ProviderType,
	overrides map[string]string,
) (repository.Repositoryer, RepoDetailsAndOpts, error) {
	cfg, err := config.GetConfigRepoDetails(
		ctx,
		provider,
		true,
		true,
		overrides)
	if err != nil {
		return nil, RepoDetailsAndOpts{}, err
	}

	repoID := cfg.RepoID
	if len(repoID) == 0 {
		repoID = events.RepoIDNotFound
	}

	opts := ControlWithConfig(cfg)

	r, err := repository.New(
		ctx,
		cfg.Account,
		cfg.Storage,
		opts,
		repoID)
	if err != nil {
		return nil, RepoDetailsAndOpts{}, clues.Wrap(err, "creating a repository controller")
	}

	if err := r.Connect(ctx); err != nil {
		return nil, RepoDetailsAndOpts{}, clues.Wrap(err, "connecting to the "+cfg.Storage.Provider.String()+" repository")
	}

	// this initializes our graph api client configurations,
	// including control options such as concurency limitations.
	if _, err := r.ConnectToM365(ctx, pst); err != nil {
		return nil, RepoDetailsAndOpts{}, clues.Wrap(err, "connecting to m365")
	}

	rdao := RepoDetailsAndOpts{
		Repo: cfg,
		Opts: opts,
	}

	return r, rdao, nil
}

func AccountConnectAndWriteRepoConfig(
	ctx context.Context,
	cmd *cobra.Command,
	pst path.ServiceType,
) (repository.Repositoryer, *account.Account, error) {
	r, rdao, err := GetAccountAndConnect(ctx, cmd, pst)
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting and connecting account")
		return nil, nil, err
	}

	sc, err := rdao.Repo.Storage.StorageConfig()
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting storage configuration")
		return nil, nil, err
	}

	m365Config, err := rdao.Repo.Account.M365Config()
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting m365 configuration")
		return nil, nil, err
	}

	// repo config gets set during repo connect and init.
	// This call confirms we have the correct values.
	err = config.WriteRepoConfig(ctx, sc, m365Config, rdao.Opts.Repo, r.GetID())
	if err != nil {
		logger.CtxErr(ctx, err).Info("writing to repository configuration")
		return nil, nil, err
	}

	return r, &rdao.Repo.Account, nil
}

// CloseRepo handles closing a repo.
func CloseRepo(ctx context.Context, r repository.Repositoryer) {
	if err := r.Close(ctx); err != nil {
		fmt.Print("Error closing repository:", err)
	}
}

// HasNoFlagsAndShownHelp shows the Help output if no flags
// were provided to the command.  Returns true if the help
// was shown.
// Use for when the non-flagged usage of a command
// (ex: corso backup restore exchange) is expected to no-op.
func HasNoFlagsAndShownHelp(cmd *cobra.Command) bool {
	if cmd.Flags().NFlag() == 0 {
		cobra.CheckErr(cmd.Help())
		return true
	}

	return false
}

type cmdCfg struct {
	hidden     bool
	preRelease bool
	preview    bool
}

type cmdOpt func(*cmdCfg)

func (cc *cmdCfg) populate(opts ...cmdOpt) {
	for _, opt := range opts {
		opt(cc)
	}
}

func HideCommand() cmdOpt {
	return func(cc *cmdCfg) {
		cc.hidden = true
	}
}

func MarkPreReleaseCommand() cmdOpt {
	return func(cc *cmdCfg) {
		cc.hidden = true
		cc.preRelease = true
	}
}

func MarkPreviewCommand() cmdOpt {
	return func(cc *cmdCfg) {
		cc.preview = true
	}
}

// AddCommand adds a clone of the subCommand to the parent,
// and returns both the clone and its pflags.
func AddCommand(parent, c *cobra.Command, opts ...cmdOpt) (*cobra.Command, *pflag.FlagSet) {
	cc := &cmdCfg{}
	cc.populate(opts...)

	parent.AddCommand(c)
	c.Hidden = cc.hidden

	if cc.preRelease {
		// There is a default deprecated message that always shows so we do some terminal magic to overwrite it
		c.Deprecated = "\n\033[1F\033[K" +
			"==================================================================================================\n" +
			"\tWARNING!!! THIS IS A PRE-RELEASE COMMAND THAT MAY NOT FUNCTION PROPERLY, OR AT ALL\n" +
			"==================================================================================================\n"
	}

	if cc.preview {
		// There is a default deprecated message that always shows so we do some terminal magic to overwrite it
		c.Deprecated = "\n\033[1F\033[K" +
			"=============================================================================================================\n" +
			"\tWARNING!!! THIS IS A FEATURE PREVIEW THAT MAY NOT FUNCTION PROPERLY AND MAY BREAK ACROSS RELEASES\n" +
			"=============================================================================================================\n"
	}

	c.Flags().SortFlags = false

	return c, c.Flags()
}

// separates the provided folders into two sets: folders that use a pathContains
// comparison (the default), and folders that use a pathPrefix comparison.
// Any element beginning with a path.PathSeparator (ie: '/') is moved to the prefix
// comparison set.  If folders is nil, returns only containsFolders with the any matcher.
func splitFoldersIntoContainsAndPrefix(folders []string) ([]string, []string) {
	var (
		containsFolders = []string{}
		prefixFolders   = []string{}
	)

	if len(folders) == 0 {
		return selectors.Any(), nil
	}

	// separate folder selection inputs by behavior.
	// any input beginning with a '/' character acts as a prefix match.
	for _, f := range folders {
		if len(f) == 0 {
			continue
		}

		if f[0] == path.PathSeparator {
			prefixFolders = append(prefixFolders, f)
		} else {
			containsFolders = append(containsFolders, f)
		}
	}

	return containsFolders, prefixFolders
}

// SendStartCorsoEvent utility sends corso start event at start of each action
func SendStartCorsoEvent(
	ctx context.Context,
	s storage.Storage,
	tenID string,
	data map[string]any,
	repoID string,
	opts control.Options,
) {
	bus, err := events.NewBus(ctx, s, tenID, opts)
	if err != nil {
		logger.CtxErr(ctx, err).Info("sending start event")
	}

	bus.SetRepoID(repoID)
	bus.Event(ctx, events.CorsoStart, data)
}

// GetStorageProviderAndOverrides returns the storage provider type and
// any flags specified on the command line which are storage provider specific.
func GetStorageProviderAndOverrides(
	ctx context.Context,
	cmd *cobra.Command,
) (storage.ProviderType, map[string]string, error) {
	provider, err := config.GetStorageProviderFromConfigFile(ctx)
	if err != nil {
		return provider, nil, clues.Stack(err)
	}

	switch provider {
	case storage.ProviderS3:
		return provider, flags.S3FlagOverrides(cmd), nil
	case storage.ProviderFilesystem:
		return provider, flags.FilesystemFlagOverrides(cmd), nil
	}

	return provider, nil, clues.New("unknown storage provider: " + provider.String())
}

// MakeAbsoluteFilePath does directory path expansions & conversions, namely:
// 1. Expands "~" prefix to the user's home directory, and converts to absolute path.
// 2. Relative paths are converted to absolute paths.
// 3. Absolute paths are returned as-is.
// 4. Empty paths are not allowed, an error is returned.
func MakeAbsoluteFilePath(p string) (string, error) {
	if len(p) == 0 {
		return "", clues.New("empty path")
	}

	// Special case handling for "~". filepath.Abs will not expand it.
	if p[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", clues.Wrap(err, "getting user home directory")
		}

		p = filepath.Join(homeDir, p[1:])
	}

	abs, err := filepath.Abs(p)
	if err != nil {
		return "", clues.Stack(err)
	}

	return abs, nil
}
