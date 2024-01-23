package metadata

import (
	"context"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/logger"
)

type SharingMode int

const (
	SharingModeCustom    SharingMode = 0
	SharingModeInherited SharingMode = 1
)

type GV2Type string

const (
	GV2App       GV2Type = "application"
	GV2Device    GV2Type = "device"
	GV2Group     GV2Type = "group"
	GV2SiteUser  GV2Type = "site_user"
	GV2SiteGroup GV2Type = "site_group"
	GV2User      GV2Type = "user"
)

// FilePermission is used to store permissions of a specific resource owner
// to a drive item.
type Permission struct {
	ID         string     `json:"id,omitempty"`
	Roles      []string   `json:"role,omitempty"`
	Email      string     `json:"email,omitempty"`    // DEPRECATED: Replaced with EntityID in newer backups
	EntityID   string     `json:"entityId,omitempty"` // this is the resource owner's ID
	EntityType GV2Type    `json:"entityType,omitempty"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

// Equal checks equality of two UserPermission objects
func (p Permission) Equals(other Permission) bool {
	if p.EntityType != other.EntityType {
		return false
	}

	// NOTE: v1 of permissions only contain emails, v2 only contains IDs.
	// The current one will contains both ID and email.
	if len(p.EntityID) > 0 && len(other.EntityID) > 0 &&
		p.EntityID != other.EntityID {
		return false
	}

	// In cases where we have shared an item with an external user,
	// the user will not have an id
	if len(p.EntityID) == 0 && len(other.EntityID) == 0 {
		if len(p.Email) > 0 && len(other.Email) > 0 &&
			p.Email != other.Email {
			return false
		}
	}

	// Possible that one is empty and the other is not
	if p.EntityID != other.EntityID {
		return false
	}

	// We cannot just compare id/email because of #3117
	p1r := p.Roles
	p2r := other.Roles

	slices.Sort(p1r)
	slices.Sort(p2r)

	return slices.Equal(p1r, p2r)
}

// DiffLinkShares is just a wrapper on top of DiffPermissions but we
// filter out link shares which do not have any associated users. This
// is useful for two reason:
//   - When a user creates a link share on parent after creating a child
//     link with `retainInheritedPermissisons`, all the previous link shares
//     are inherited onto the child but without any users associated with
//     the share. We have to drop the empty ones to make sure we reset.
//   - We are restoring link shares so that we can restore permissions for
//     the user, but restoring links without users is not useful.
func DiffLinkShares(current, expected []LinkShare) ([]LinkShare, []LinkShare) {
	filteredCurrent := []LinkShare{}
	filteredExpected := []LinkShare{}

	for _, ls := range current {
		// It is useless to restore link shares without associated
		// entities. When we do a link share restore, it always creates
		// a new link(we cannot reuse the old link). Since we
		// have no way of knowing the users who previously had access to
		// an item via link, we don't have a reason to restore the link.
		if len(ls.Entities) == 0 {
			continue
		}

		filteredCurrent = append(filteredCurrent, ls)
	}

	for _, ls := range expected {
		if len(ls.Entities) == 0 {
			continue
		}

		filteredExpected = append(filteredExpected, ls)
	}

	return DiffPermissions(filteredCurrent, filteredExpected)
}

// DiffPermissions compares the before and after set, returning
// the permissions that were added and removed (in that order)
// in the after set.
func DiffPermissions[T interface{ Equals(T) bool }](current, expected []T) ([]T, []T) {
	var (
		added   = []T{}
		removed = []T{}
	)

	for _, cp := range expected {
		found := false

		for _, pp := range current {
			if cp.Equals(pp) {
				found = true
				break
			}
		}

		if !found {
			added = append(added, cp)
		}
	}

	for _, pp := range current {
		found := false

		for _, cp := range expected {
			if cp.Equals(pp) {
				found = true
				break
			}
		}

		if !found {
			removed = append(removed, pp)
		}
	}

	return added, removed
}

func FilterPermissions(ctx context.Context, perms []models.Permissionable) []Permission {
	up := []Permission{}

	for _, p := range perms {
		if p.GetGrantedToV2() == nil {
			// For link shares, we get permissions without a user
			// specified
			continue
		}

		// Below are the mapping from roles to "Advanced" permissions
		// screen entries:
		//
		// owner - Full Control
		// write - Design | Edit | Contribute (no difference in /permissions api)
		// read  - Read
		// empty - Restricted View
		//
		// helpful docs:
		// https://devblogs.microsoft.com/microsoft365dev/controlling-app-access-on-specific-sharepoint-site-collections/
		roles := p.GetRoles()

		ent, ok := getIdentityDetails(ctx, p.GetGrantedToV2())
		if !ok {
			// We log the inability to handle certain type of
			// permissions within the getIdentityDetails function and so
			// we just skip here
			continue
		}

		up = append(up, Permission{
			ID:         ptr.Val(p.GetId()),
			Roles:      roles,
			EntityID:   ent.ID,
			Email:      ent.Email, // not necessary if we have email, but useful for debugging
			EntityType: ent.EntityType,
			Expiration: p.GetExpirationDateTime(),
		})
	}

	return up
}

func FilterLinkShares(ctx context.Context, perms []models.Permissionable) []LinkShare {
	up := []LinkShare{}

	for _, p := range perms {
		link := p.GetLink()
		if link == nil {
			// Non link share based permissions are handled separately
			continue
		}

		var (
			roles = p.GetRoles()
			gv2   = p.GetGrantedToIdentitiesV2()
		)

		idens := []Entity{}

		for _, g := range gv2 {
			ent, ok := getIdentityDetails(ctx, g)
			if !ok {
				continue
			}

			idens = append(idens, ent)
		}

		up = append(up, LinkShare{
			ID: ptr.Val(p.GetId()),
			Link: LinkShareLink{
				Scope:            ptr.Val(link.GetScope()),
				Type:             ptr.Val(link.GetTypeEscaped()),
				WebURL:           ptr.Val(link.GetWebUrl()),
				PreventsDownload: ptr.Val(link.GetPreventsDownload()),
			},
			Roles:       roles,
			Entities:    idens,
			HasPassword: ptr.Val(p.GetHasPassword()),
			Expiration:  p.GetExpirationDateTime(),
		})
	}

	return up
}

func getIdentityDetails(ctx context.Context, gv2 models.SharePointIdentitySetable) (Entity, bool) {
	switch true {
	case gv2.GetUser() != nil:
		add := gv2.GetUser().GetAdditionalData()
		email, _ := str.AnyToString(add["email"]) // empty will be dropped automatically when writing

		return Entity{
			ID:         ptr.Val(gv2.GetUser().GetId()),
			Email:      email,
			EntityType: GV2User,
		}, true
	case gv2.GetSiteUser() != nil:
		return Entity{
			ID:         ptr.Val(gv2.GetSiteUser().GetId()),
			EntityType: GV2SiteUser,
		}, true
	case gv2.GetGroup() != nil:
		return Entity{
			ID:         ptr.Val(gv2.GetGroup().GetId()),
			EntityType: GV2Group,
		}, true
	case gv2.GetSiteGroup() != nil:
		return Entity{
			ID:         ptr.Val(gv2.GetSiteGroup().GetId()),
			EntityType: GV2SiteGroup,
		}, true
	case gv2.GetApplication() != nil:
		return Entity{
			ID:         ptr.Val(gv2.GetApplication().GetId()),
			EntityType: GV2App,
		}, true
	case gv2.GetDevice() != nil:
		return Entity{
			ID:         ptr.Val(gv2.GetDevice().GetId()),
			EntityType: GV2Device,
		}, true
	default:
		logger.Ctx(ctx).Info("untracked permission")
		return Entity{}, false
	}
}
