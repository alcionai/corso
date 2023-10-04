package backup

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"

	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// Kopia does not do comparisons properly for empty tags right now so add some
	// placeholder value to them.
	legacyDefaultTagValue = "0"

	// Kopia CLI prefixes all user tags with "tag:"[1]. Maintaining this will
	// ensure we don't accidentally take reserved tags and that tags can be
	// displayed with kopia CLI.
	// (permalinks)
	// [1] https://github.com/kopia/kopia/blob/05e729a7858a6e86cb48ba29fb53cb6045efce2b/cli/command_snapshot_create.go#L169
	LegacyUserTagPrefix = "tag:"

	tenantIDKey      = "tenant"
	resourceIDKey    = "resource"
	serviceCatPrefix = "sc-"

	// Sentinel value for tags. Could technically be empty but we'll store
	// something for now.
	//nolint
	defaultTagValue = "1"
)

func serviceCatString(
	service path.ServiceType,
	category path.CategoryType,
) string {
	return serviceCatPrefix + service.String() + category.String()
}

// reasonTags returns the set of key-value pairs that can be used as tags to
// represent this Reason.
// nolint
func reasonTags(r identity.Reasoner) map[string]string {
	return map[string]string{
		tenantIDKey:   r.Tenant(),
		resourceIDKey: r.ProtectedResource(),
		serviceCatString(r.Service(), r.Category()): defaultTagValue,
	}
}

// nolint
type BackupEntry struct {
	*Backup
	Reasons []identity.Reasoner
}

type ManifestEntry struct {
	*snapshot.Manifest
	// Reasons contains the ResourceOwners and Service/Categories that caused this
	// snapshot to be selected as a base. We can't reuse OwnersCats here because
	// it's possible some ResourceOwners will have a subset of the Categories as
	// the reason for selecting a snapshot. For example:
	// 1. backup user1 email,contacts -> B1
	// 2. backup user1 contacts -> B2 (uses B1 as base)
	// 3. backup user1 email,contacts,events (uses B1 for email, B2 for contacts)
	Reasons []identity.Reasoner
}

// MakeTagKV normalizes the provided key to protect it from clobbering
// similarly named tags from non-user input (user inputs are still open
// to collisions amongst eachother).
// Returns the normalized Key plus a default value.  If you're embedding a
// key-only tag, the returned default value msut be used instead of an
// empty string.
func MakeTagKV(k string) (string, string) {
	return LegacyUserTagPrefix + k, legacyDefaultTagValue
}

func (me ManifestEntry) GetTag(key string) (string, bool) {
	k, _ := MakeTagKV(key)
	v, ok := me.Tags[k]

	return v, ok
}

// nolint
type BackupBases interface {
	// ConvertToAssistBase converts the base with the given item data snapshot ID
	// from a merge base to an assist base.
	ConvertToAssistBase(manifestID manifest.ID)
	Backups() []BackupEntry
	UniqueAssistBackups() []BackupEntry
	MinBackupVersion() int
	MergeBases() []ManifestEntry
	DisableMergeBases()
	UniqueAssistBases() []ManifestEntry
	DisableAssistBases()
	MergeBackupBases(
		ctx context.Context,
		other BackupBases,
		reasonToKey func(identity.Reasoner) string,
	) BackupBases
	// SnapshotAssistBases returns the set of bases to use for kopia assisted
	// incremental snapshot operations. It consists of the union of merge bases
	// and assist bases. If DisableAssistBases has been called then it returns
	// nil.
	SnapshotAssistBases() []ManifestEntry
}
