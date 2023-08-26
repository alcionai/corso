package utils

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
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

func GetAccountAndConnect(
	ctx context.Context,
	pst path.ServiceType,
	overrides map[string]string,
) (repository.Repository, *storage.Storage, *account.Account, *control.Options, error) {
	cfg, err := config.GetConfigRepoDetails(ctx, true, true, overrides)
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
	pst path.ServiceType,
	overrides map[string]string,
) (repository.Repository, *account.Account, error) {
	r, stg, acc, opts, err := GetAccountAndConnect(ctx, pst, overrides)
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting and connecting account")
		return nil, nil, err
	}

	storageConfig, err := config.NewStorageConfigFrom(*stg)
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting storage configuration")
		return nil, nil, err
	}

	m365Config, err := acc.M365Config()
	if err != nil {
		logger.CtxErr(ctx, err).Info("getting m365 configuration")
		return nil, nil, err
	}

	// repo config gets set during repo connect and init.
	// This call confirms we have the correct values.
	err = config.WriteRepoConfig(ctx, storageConfig, m365Config, opts.Repo, r.GetID())
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
