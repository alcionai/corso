package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

const (
	Wildcard = "*"
)

// RequireProps validates the existence of the properties
// in the map.  Expects the format map[propName]propVal.
func RequireProps(props map[string]string) error {
	for name, val := range props {
		if len(val) == 0 {
			return errors.New(name + " is required to perform this command")
		}
	}

	return nil
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
	hidden    bool
	preRelese bool
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
		cc.preRelese = true
	}
}

// AddCommand adds a clone of the subCommand to the parent,
// and returns both the clone and its pflags.
func AddCommand(parent, c *cobra.Command, opts ...cmdOpt) (*cobra.Command, *pflag.FlagSet) {
	cc := &cmdCfg{}
	cc.populate(opts...)

	parent.AddCommand(c)
	c.Hidden = cc.hidden

	if cc.preRelese {
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
		logger.Ctx(ctx).Infow("analytics event failure", "err", err)
	}

	bus.SetRepoID(repoID)
	bus.Event(ctx, events.CorsoStart, data)
}
