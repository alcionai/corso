package common

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// Resource Lookup Handling
// ---------------------------------------------------------------------------

func GetResourceClient(rc resource.Category, ac api.Client) (*resourceClient, error) {
	switch rc {
	case resource.Users:
		return &resourceClient{enum: rc, getter: ac.Users()}, nil
	case resource.Sites:
		return &resourceClient{enum: rc, getter: ac.Sites()}, nil
	case resource.Groups:
		return &resourceClient{enum: rc, getter: ac.Groups()}, nil
	default:
		return nil, clues.New("unrecognized owner resource type").With("resource_enum", rc)
	}
}

type resourceClient struct {
	enum   resource.Category
	getter getIDAndNamer
}

type getIDAndNamer interface {
	GetIDAndName(
		ctx context.Context,
		owner string,
		cc api.CallConfig,
	) (
		ownerID string,
		ownerName string,
		err error,
	)
}

var _ idname.GetResourceIDAndNamer = &resourceClient{}

// GetResourceIDAndNameFrom looks up the resource's canonical id and display name.
// If the resource is present in the idNameSwapper, then that interface's id and
// name values are returned.  As a fallback, the resource calls the discovery
// api to fetch the user or site using the resource value. This fallback assumes
// that the resource is a well formed ID or display name of appropriate design
// (PrincipalName for users, WebURL for sites).
func (r resourceClient) GetResourceIDAndNameFrom(
	ctx context.Context,
	owner string,
	ins idname.Cacher,
) (idname.Provider, error) {
	if ins != nil {
		if n, ok := ins.NameOf(owner); ok {
			return idname.NewProvider(owner, n), nil
		} else if i, ok := ins.IDOf(owner); ok {
			return idname.NewProvider(i, owner), nil
		}
	}

	ctx = clues.Add(ctx, "owner_identifier", owner)

	var (
		id, name string
		err      error
	)

	id, name, err = r.getter.GetIDAndName(ctx, owner, api.CallConfig{})
	if err != nil {
		if graph.IsErrUserNotFound(err) {
			return nil, clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		if graph.IsErrResourceLocked(err) {
			return nil, clues.Stack(graph.ErrResourceLocked, err)
		}

		return nil, err
	}

	if len(id) == 0 || len(name) == 0 {
		return nil, clues.Stack(graph.ErrResourceOwnerNotFound)
	}

	return idname.NewProvider(id, name), nil
}
