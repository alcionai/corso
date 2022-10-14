//go:build testing

package onedrive

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
)

//revive:disable:context-as-argument
func MustGetDefaultDriveID(
	t *testing.T,
	ctx context.Context,
	service graph.Service,
	userID string,
) string {
	//revive:enable:context-as-argument
	d, err := service.Client().UsersById(userID).Drive().Get(ctx, nil)
	if err != nil {
		err = errors.Wrapf(
			err,
			"failed to retrieve user drives. user: %s, details: %s",
			userID,
			support.ConnectorStackErrorTrace(err),
		)

		require.NoError(t, err)
	}

	require.NotNil(t, d.GetId())

	return *d.GetId()
}
