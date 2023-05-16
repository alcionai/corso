package metadata

import (
	"time"

	"golang.org/x/exp/slices"
)

type SharingMode int

const (
	SharingModeCustom = SharingMode(iota)
	SharingModeInherited
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

// isSamePermission checks equality of two UserPermission objects
func (p Permission) Equals(other Permission) bool {
	// EntityID can be empty for older backups and Email can be empty
	// for newer ones. It is not possible for both to be empty.  Also,
	// if EntityID/Email for one is not empty then the other will also
	// have EntityID/Email as we backup permissions for all the
	// parents and children when we have a change in permissions.
	if p.EntityID != "" && p.EntityID != other.EntityID {
		return false
	}

	if p.Email != "" && p.Email != other.Email {
		return false
	}

	p1r := p.Roles
	p2r := other.Roles

	slices.Sort(p1r)
	slices.Sort(p2r)

	return slices.Equal(p1r, p2r)
}

// DiffPermissions compares the before and after set, returning
// the permissions that were added and removed (in that order)
// in the after set.
func DiffPermissions(before, after []Permission) ([]Permission, []Permission) {
	var (
		added   = []Permission{}
		removed = []Permission{}
	)

	for _, cp := range after {
		found := false

		for _, pp := range before {
			if cp.Equals(pp) {
				found = true
				break
			}
		}

		if !found {
			added = append(added, cp)
		}
	}

	for _, pp := range before {
		found := false

		for _, cp := range after {
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
