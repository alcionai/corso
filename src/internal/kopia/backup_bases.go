package kopia

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/logger"
)

// TODO(ashmrtn): Move this into some inject package. Here to avoid import
// cycles.
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

	// TODO(ashmrtn): Remove other functions and just have these once other code
	// is updated. Here for now so changes in this file can be made.
	NewMergeBases() []BackupBase
	NewUniqueAssistBases() []BackupBase
}

type backupBases struct {
	mergeBases  []BackupBase
	assistBases []BackupBase

	// disableAssistBases denote whether any assist bases should be returned to
	// kopia during snapshot operation.
	disableAssistBases bool
}

func (bb *backupBases) SnapshotAssistBases() []ManifestEntry {
	if bb.disableAssistBases {
		return nil
	}

	res := []ManifestEntry{}

	for _, ab := range bb.assistBases {
		res = append(res, ManifestEntry{
			Manifest: ab.ItemDataSnapshot,
			Reasons:  ab.Reasons,
		})
	}

	for _, mb := range bb.mergeBases {
		res = append(res, ManifestEntry{
			Manifest: mb.ItemDataSnapshot,
			Reasons:  mb.Reasons,
		})
	}

	// Need to use the actual variables here because the functions will return nil
	// depending on what's been marked as disabled.
	return res
}

func (bb *backupBases) ConvertToAssistBase(manifestID manifest.ID) {
	idx := slices.IndexFunc(
		bb.mergeBases,
		func(base BackupBase) bool {
			return base.ItemDataSnapshot.ID == manifestID
		})
	if idx >= 0 {
		bb.assistBases = append(bb.assistBases, bb.mergeBases[idx])
		bb.mergeBases = slices.Delete(bb.mergeBases, idx, idx+1)
	}
}

func (bb backupBases) Backups() []BackupEntry {
	res := []BackupEntry{}

	for _, mb := range bb.mergeBases {
		res = append(res, BackupEntry{
			Backup:  mb.Backup,
			Reasons: mb.Reasons,
		})
	}

	return res
}

func (bb backupBases) UniqueAssistBackups() []BackupEntry {
	if bb.disableAssistBases {
		return nil
	}

	res := []BackupEntry{}

	for _, ab := range bb.assistBases {
		res = append(res, BackupEntry{
			Backup:  ab.Backup,
			Reasons: ab.Reasons,
		})
	}

	return res
}

func (bb *backupBases) MinBackupVersion() int {
	min := version.NoBackup

	if bb == nil {
		return min
	}

	for _, base := range bb.mergeBases {
		if min == version.NoBackup || base.Backup.Version < min {
			min = base.Backup.Version
		}
	}

	return min
}

func (bb backupBases) MergeBases() []ManifestEntry {
	res := []ManifestEntry{}

	for _, mb := range bb.mergeBases {
		res = append(res, ManifestEntry{
			Manifest: mb.ItemDataSnapshot,
			Reasons:  mb.Reasons,
		})
	}

	return res
}

func (bb backupBases) NewMergeBases() []BackupBase {
	return slices.Clone(bb.mergeBases)
}

func (bb *backupBases) DisableMergeBases() {
	// Turn all merge bases into assist bases. We don't want to remove them
	// completely because we still want to allow kopia assisted incrementals
	// unless that's also explicitly disabled. However, we can't just leave them
	// in the merge set since then we won't return the bases when merging backup
	// details.
	bb.assistBases = append(bb.assistBases, bb.mergeBases...)

	bb.mergeBases = nil
}

func (bb backupBases) UniqueAssistBases() []ManifestEntry {
	if bb.disableAssistBases {
		return nil
	}

	res := []ManifestEntry{}

	for _, ab := range bb.assistBases {
		res = append(res, ManifestEntry{
			Manifest: ab.ItemDataSnapshot,
			Reasons:  ab.Reasons,
		})
	}

	return res
}

func (bb backupBases) NewUniqueAssistBases() []BackupBase {
	if bb.disableAssistBases {
		return nil
	}

	return slices.Clone(bb.assistBases)
}

func (bb *backupBases) DisableAssistBases() {
	bb.disableAssistBases = true
}

func getMissingBases(
	reasonToKey func(identity.Reasoner) string,
	seen map[string]struct{},
	toCheck []BackupBase,
) []BackupBase {
	var res []BackupBase

	for _, base := range toCheck {
		useReasons := []identity.Reasoner{}

		for _, r := range base.Reasons {
			k := reasonToKey(r)
			if _, ok := seen[k]; ok {
				// This Reason is already "covered" by a previously seen base. Skip
				// adding the Reason to the base being examined.
				continue
			}

			useReasons = append(useReasons, r)
		}

		if len(useReasons) > 0 {
			base.Reasons = useReasons
			res = append(res, base)
		}
	}

	return res
}

// MergeBackupBases reduces the two BackupBases into a single BackupBase.
// Assumes the passed in BackupBases represents a prior backup version (across
// some migration that disrupts lookup), and that the BackupBases used to call
// this function contains the current version.
//
// This call should be made prior to Disable*Bases being called on either the
// called BackupBases or the passed in BackupBases.
//
// reasonToKey should be a function that, given a Reasoner, will produce some
// string that represents Reasoner in the context of the merge operation. For
// example, to merge BackupBases across a ProtectedResource migration, the
// Reasoner's service and category can be used as the key.
//
// Selection priority, for each reason key generated by reasonsToKey, follows
// these rules:
//  1. If the called BackupBases has an entry for a given reason, ignore the
//     other BackupBases matching that reason.
//  2. If the called BackupBases has only AssistBases, look for a matching
//     MergeBase manifest in the other BackupBases.
//  3. If the called BackupBases has no entry for a reason, look for a matching
//     MergeBase in the other BackupBases.
func (bb *backupBases) MergeBackupBases(
	ctx context.Context,
	other BackupBases,
	reasonToKey func(reason identity.Reasoner) string,
) BackupBases {
	if other == nil || (len(other.NewMergeBases()) == 0 && len(other.NewUniqueAssistBases()) == 0) {
		return bb
	}

	if bb == nil || (len(bb.NewMergeBases()) == 0 && len(bb.NewUniqueAssistBases()) == 0) {
		return other
	}

	toMerge := map[string]struct{}{}
	assist := map[string]struct{}{}

	// Track the bases in bb. We need to know the Reason(s) covered by merge bases
	// and the Reason(s) covered by assist bases separately because the former
	// dictates whether we need to select a merge base and an assist base from
	// other while the latter dictates whether we need to select an assist base
	// from other.
	for _, m := range bb.MergeBases() {
		for _, r := range m.Reasons {
			k := reasonToKey(r)

			toMerge[k] = struct{}{}
			assist[k] = struct{}{}
		}
	}

	for _, m := range bb.UniqueAssistBases() {
		for _, r := range m.Reasons {
			k := reasonToKey(r)
			assist[k] = struct{}{}
		}
	}

	addMerge := getMissingBases(reasonToKey, toMerge, other.NewMergeBases())
	addAssist := getMissingBases(reasonToKey, assist, other.NewUniqueAssistBases())

	res := &backupBases{
		mergeBases:  append(addMerge, bb.NewMergeBases()...),
		assistBases: append(addAssist, bb.NewUniqueAssistBases()...),
	}

	return res
}

func fixupMinRequirements(
	ctx context.Context,
	baseSet []BackupBase,
) []BackupBase {
	res := make([]BackupBase, 0, len(baseSet))

	for _, base := range baseSet {
		var (
			backupID       model.StableID
			snapID         manifest.ID
			snapIncomplete bool
			deetsID        string
		)

		if base.Backup != nil {
			backupID = base.Backup.ID

			deetsID = base.Backup.StreamStoreID
			if len(deetsID) == 0 {
				deetsID = base.Backup.DetailsID
			}
		}

		if base.ItemDataSnapshot != nil {
			snapID = base.ItemDataSnapshot.ID

			snapIncomplete = len(base.ItemDataSnapshot.IncompleteReason) > 0
		}

		ictx := clues.Add(
			ctx,
			"base_backup_id", backupID,
			"base_item_data_snapshot_id", snapID,
			"base_details_id", deetsID)

		switch {
		case len(backupID) == 0:
			logger.Ctx(ictx).Info("dropping base missing backup model")
			continue

		case len(snapID) == 0:
			logger.Ctx(ictx).Info("dropping base missing item data snapshot")
			continue

		case snapIncomplete:
			logger.Ctx(ictx).Info("dropping base with incomplete item data snapshot")
			continue

		case len(deetsID) == 0:
			logger.Ctx(ictx).Info("dropping base missing backup details")
			continue

		case len(base.Reasons) == 0:
			// Not sure how we'd end up here, but just to make sure we're really
			// getting what we expect.
			logger.Ctx(ictx).Info("dropping base with no marked Reasons")
			continue
		}

		res = append(res, base)
	}

	return res
}

func fixupReasons(
	ctx context.Context,
	baseSet []BackupBase,
) []BackupBase {
	// Associate a Reason with a set of bases since the basesByReason map needs a
	// string key.
	type baseEntry struct {
		bases  []BackupBase
		reason identity.Reasoner
	}

	var (
		basesByReason = map[string]baseEntry{}
		// res holds a mapping from backup ID -> base. We need this additional level
		// of indirection when determining what to return because a base may be
		// selected for multiple reasons. This map allows us to consolidate that
		// into a single base result for all reasons easily.
		res = map[model.StableID]BackupBase{}
	)

	// Organize all the base(s) by the Reason(s) they were chosen. A base can
	// exist in multiple slices in the map if it was selected for multiple
	// Reasons.
	for _, base := range baseSet {
		for _, reason := range base.Reasons {
			foundBases := basesByReason[reasonKey(reason)]
			foundBases.reason = reason
			foundBases.bases = append(foundBases.bases, base)

			basesByReason[reasonKey(reason)] = foundBases
		}
	}

	// Go through the map and check that the length of each slice is 1. If it's
	// longer than that then we somehow got multiple bases for the same Reason and
	// should drop the extras.
	for _, bases := range basesByReason {
		ictx := clues.Add(
			ctx,
			"verify_service", bases.reason.Service().String(),
			"verify_category", bases.reason.Category().String())

		// Not sure how we'd actually get here but handle it anyway.
		if len(bases.bases) == 0 {
			logger.Ctx(ictx).Info("no bases found for reason")
			continue
		}

		// We've got at least one base for this Reason. The below finds which base
		// to keep based on the creation time of the bases. If there's multiple
		// bases in the input slice then we'll log information about the ones that
		// we didn't add to the result set.

		// Sort in reverse chronological order so that it's easy to find the
		// youngest base.
		slices.SortFunc(bases.bases, func(a, b BackupBase) int {
			return -a.Backup.CreationTime.Compare(b.Backup.CreationTime)
		})

		keepBase := bases.bases[0]

		// Add the youngest base to the result set. We add each Reason for selecting
		// the base individually so that bases dropped for a particular Reason (or
		// dropped completely because they overlap for all Reasons) happens without
		// additional logic. The dropped (Reason, base) pair will just never be
		// added to the result set to begin with.
		b, ok := res[keepBase.Backup.ID]
		if ok {
			// We've already seen this base, just add this Reason to it as well.
			b.Reasons = append(b.Reasons, bases.reason)
			res[keepBase.Backup.ID] = b

			continue
		}

		// We haven't seen this base before. We want to clear all the Reasons for it
		// except the one we're currently examining. That allows us to just not add
		// bases that are duplicates for a Reason to res and still end up with the
		// correct output.
		keepBase.Reasons = []identity.Reasoner{bases.reason}
		res[keepBase.Backup.ID] = keepBase

		// Don't log about dropped bases if there was only one base.
		if len(bases.bases) == 1 {
			continue
		}

		// This is purely for debugging, but log the base(s) that we dropped for
		// this Reason.
		var dropped []model.StableID

		for _, b := range bases.bases[1:] {
			dropped = append(dropped, b.Backup.ID)
		}

		logger.Ctx(ictx).Infow(
			"dropping bases for reason",
			"dropped_backup_ids", dropped)
	}

	return maps.Values(res)
}

// fixupAndVerify goes through the set of backups and snapshots used for merging
// and ensures:
//   - the reasons for selecting merge snapshots are distinct
//   - all bases have a backup model with item and details snapshot IDs
//   - all bases have both a backup and item data snapshot present
//   - all bases have item data snapshots with no incomplete reason
//
// Backups that have overlapping reasons or that are not complete are removed
// from the set. Dropping these is safe because it only affects how much data we
// pull. On the other hand, *not* dropping them is unsafe as it will muck up
// merging when we add stuff to kopia (possibly multiple entries for the same
// item etc).
func (bb *backupBases) fixupAndVerify(ctx context.Context) {
	// Start off by removing bases that don't meet the minimum requirements of
	// having a backup model and item data snapshot or having a backup details ID.
	// These requirements apply to both merge and assist bases.
	bb.mergeBases = fixupMinRequirements(ctx, bb.mergeBases)
	bb.assistBases = fixupMinRequirements(ctx, bb.assistBases)

	// Remove merge bases that have overlapping Reasons. It's alright to call this
	// on assist bases too because we only expect at most one assist base per
	// Reason.
	bb.mergeBases = fixupReasons(ctx, bb.mergeBases)
	bb.assistBases = fixupReasons(ctx, bb.assistBases)
}
