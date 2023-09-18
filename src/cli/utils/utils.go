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

var ErrNotYetImplemented = clues.New("not yet implemented")

// GetAccountAndConnectWithOverrides is a wrapper for GetAccountAndConnect
// that also gets the storage provider and any storage provider specific
// flag overrides from the command line.
func GetAccountAndConnectWithOverrides(
	ctx context.Context,
	cmd *cobra.Command,
	pst path.ServiceType,
) (repository.Repository, *storage.Storage, *account.Account, *control.Options, error) {
	provider, overrides, err := GetStorageProviderAndOverrides(ctx, cmd)
	if err != nil {
		return nil, nil, nil, nil, clues.Stack(err)
	}

	return GetAccountAndConnect(ctx, pst, provider, overrides)
}

func GetAccountAndConnect(
	ctx context.Context,
	pst path.ServiceType,
	provider storage.ProviderType,
	overrides map[string]string,
) (repository.Repository, *storage.Storage, *account.Account, *control.Options, error) {
	cfg, err := config.GetConfigRepoDetails(
		ctx,
		provider,
		true,
		true,
		overrides)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	repoID := cfg.RepoID
	if len(repoID) == 0 {
		repoID = events.RepoIDNotFound
	}

	opts := ControlWithConfig(cfg)

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, repoID, opts)
	if err != nil {
		return nil, nil, nil, nil, clues.Wrap(err, "connecting to the "+cfg.Storage.Provider.String()+" repository")
	}

	// this initializes our graph api client configurations,
	// including control options such as concurency limitations.
	if _, err := r.ConnectToM365(ctx, pst); err != nil {
		return nil, nil, nil, nil, clues.Wrap(err, "connecting to m365")
	}

	return r, &cfg.Storage, &cfg.Account, &opts, nil
}

func AccountConnectAndWriteRepoConfig(
	ctx context.Context,
	cmd *cobra.Command,
	pst path.ServiceType,
) (repository.Repository, *account.Account, error) {
	r, stg, acc, opts, err := GetAccountAndConnectWithOverrides(
		ctx,
		cmd,
		pst)
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting and connecting account")
		return nil, nil, err
	}

	sc, err := stg.StorageConfig()
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting storage configuration")
		return nil, nil, err
	}

	s3Config := sc.(*storage.S3Config)

	m365Config, err := acc.M365Config()
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting m365 configuration")
		return nil, nil, err
	}

	// repo config gets set during repo connect and init.
	// This call confirms we have the correct values.
	err = config.WriteRepoConfig(ctx, s3Config, m365Config, opts.Repo, r.GetID())
	if err != nil {
		logger.CtxErr(ctx, err).Info("writing to repository configuration")
		return nil, nil, err
	}

	return r, acc, nil
}

// CloseRepo handles closing a repo.
func CloseRepo(ctx context.Context, r repository.Repository) {
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

	overrides := map[string]string{}

	switch provider {
	case storage.ProviderS3:
		overrides = flags.S3FlagOverrides(cmd)
	}

	return provider, overrides, nil
}

func MakeAbsoluteFilePath(p string) (string, error) {
	if len(p) == 0 {
		return "", clues.New("empty path")
	}

	// Special case handling for "~". filepath.Abs will not expand it.
	// If the path starts with "~", expand it to the user's home directory.
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
