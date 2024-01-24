package driveish

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/tform"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func ComparatorEqualPerms(expect, result []common.PermissionInfo) func() bool {
	return func() bool {
		return len(expect) == len(result)
	}
}

// was getting used by sharepoint, but sharepoint was also skipping permissions
// tests. Keeping here for reference.
// func ComparatorExpectFewerPerms(expect, result []common.PermissionInfo) func() bool {
// 	return func() bool {
// 		return len(expect) <= len(result)
// 	}
// }

func CheckRestoration(
	ctx context.Context,
	ac api.Client,
	drive models.Driveable,
	envs common.Envs,
	permissionsComparator func(expect, result []common.PermissionInfo) func() bool,
) {
	var (
		driveID   = ptr.Val(drive.GetId())
		driveName = ptr.Val(drive.GetName())
	)

	ctx = clues.Add(
		ctx,
		"drive_id", driveID,
		"drive_name", driveName)

	root := populateSanitree(ctx, ac, driveID, envs.RestoreContainer)

	sourceTree, ok := root.Children[envs.SourceContainer]
	common.Assert(
		ctx,
		func() bool { return ok },
		"should find root-level source data folder",
		envs.SourceContainer,
		"not found")

	restoreTree, ok := root.Children[envs.RestoreContainer]
	common.Assert(
		ctx,
		func() bool { return ok },
		"should find root-level restore folder",
		envs.RestoreContainer,
		"not found")

	var permissionCheck common.ContainerComparatorFn[
		models.DriveItemable, models.DriveItemable,
		models.DriveItemable, models.DriveItemable]

	if permissionsComparator != nil {
		permissionCheck = checkRestoredDriveItemPermissions(permissionsComparator)
	}

	common.AssertEqualTrees[models.DriveItemable](
		ctx,
		sourceTree,
		restoreTree.Children[envs.SourceContainer],
		permissionCheck,
		nil)

	common.Infof(ctx, "Success")
}

func permissionIn(
	ctx context.Context,
	ac api.Client,
	driveID, itemID string,
	cannotAllowErrors bool,
) []common.PermissionInfo {
	pi := []common.PermissionInfo{}

	pcr, err := ac.Drives().GetItemPermission(ctx, driveID, itemID)
	if err != nil {
		if cannotAllowErrors {
			common.Fatal(ctx, "getting permission", err)
		}

		common.Infof(
			ctx,
			"ignoring error getting permissions for %q\nerror: %s,%+v",
			itemID,
			err.Error(),
			clues.ToCore(err))

		return []common.PermissionInfo{}
	}

	for _, perm := range pcr.GetValue() {
		if perm.GetGrantedToV2() == nil {
			continue
		}

		var (
			gv2      = perm.GetGrantedToV2()
			permInfo = common.PermissionInfo{}
			entityID string
		)

		// TODO: replace with filterUserPermissions in onedrive item.go
		if gv2.GetUser() != nil {
			entityID = ptr.Val(gv2.GetUser().GetId())
		} else if gv2.GetGroup() != nil {
			entityID = ptr.Val(gv2.GetGroup().GetId())
		}

		roles := common.FilterSlice(perm.GetRoles(), owner)
		for _, role := range roles {
			permInfo.EntityID = entityID
			permInfo.Roles = append(permInfo.Roles, role)
		}

		if len(roles) > 0 {
			slices.Sort[[]string, string](permInfo.Roles)

			pi = append(pi, permInfo)
		}
	}

	return pi
}

/*
TODO: replace this check with testElementsMatch
from internal/connecter/graph_connector_helper_test.go
*/
func checkRestoredDriveItemPermissions(
	comparator func(expect, result []common.PermissionInfo) func() bool,
) common.ContainerComparatorFn[
	models.DriveItemable, models.DriveItemable,
	models.DriveItemable, models.DriveItemable,
] {
	return func(
		ctx context.Context,
		expect, result *common.Sanitree[models.DriveItemable, models.DriveItemable],
	) {
		expectPerms, err := tform.AnyValueToT[[]common.PermissionInfo](
			expandPermissions,
			expect.Expand)
		common.Assert(
			ctx,
			func() bool { return err == nil },
			"should find permissions in 'expect' node Expand data",
			expect.Name,
			err)

		resultPerms, err := tform.AnyValueToT[[]common.PermissionInfo](
			expandPermissions,
			result.Expand)
		common.Assert(
			ctx,
			func() bool { return err == nil },
			"should find permissions in 'result' node Expand data",
			result.Name,
			err)

		if len(expectPerms) == 0 {
			common.Infof(ctx, "no permissions found in folder: %s", expect.Name)
			return
		}

		common.Assert(
			ctx,
			comparator(expectPerms, resultPerms),
			"wrong number of restored permissions",
			expectPerms,
			resultPerms)

		for _, perm := range expectPerms {
			eqID := func(pi common.PermissionInfo) bool {
				return strings.EqualFold(pi.EntityID, perm.EntityID)
			}

			i := slices.IndexFunc(resultPerms, eqID)

			common.Assert(
				ctx,
				func() bool { return i >= 0 },
				"restore is missing permission",
				perm.EntityID,
				resultPerms)

			// permissions should be sorted, so a by-index comparison works
			restored := resultPerms[i]

			common.Assert(
				ctx,
				func() bool { return slices.Equal(perm.Roles, restored.Roles) },
				"different roles restored",
				perm.Roles,
				restored.Roles)
		}
	}
}
