package groups

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type backupHandler interface {
	// gets all channels for the group
	getChannels(
		ctx context.Context,
	) ([]models.Channelable, error)

	// gets all messages by delta in the channel
	getChannelMessagesDelta(
		ctx context.Context,
		channelID, prevDelta string,
	) ([]models.ChatMessageable, api.DeltaUpdate, error)

	// includeContainer evaluates whether the channel is included
	// in he provided scope.
	includeContainer(
		ctx context.Context,
		qp graph.QueryParams,
		ch models.Channelable,
		scope selectors.GroupsScope,
	) bool

	// canonicalPath constructs the service and category specific path for
	// the given builder.
	canonicalPath(
		folders *path.Builder,
		tenantID string,
	) (path.Path, error)
}
