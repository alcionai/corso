package testdata

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
)

func AssertMetadataEqual(t *testing.T, expect, got metadata.Metadata) {
	assert.Equal(t, expect.FileName, got.FileName, "fileName")
	assert.Equal(t, expect.SharingMode, got.SharingMode, "sharingMode")
	assert.Equal(t, len(expect.Permissions), len(got.Permissions), "permissions count")

	for i, ep := range expect.Permissions {
		gp := got.Permissions[i]

		assert.Equal(t, ep.EntityType, gp.EntityType, "permission %d entityType", i)
		assert.Equal(t, ep.EntityID, gp.EntityID, "permission %d entityID", i)
		assert.Equal(t, ep.ID, gp.ID, "permission %d ID", i)
		assert.ElementsMatch(t, ep.Roles, gp.Roles, "permission %d roles", i)
	}
}

func NewStubPermissionResponse(
	gv2 metadata.GV2Type,
	permID, entityID string,
	roles []string,
) models.PermissionCollectionResponseable {
	var (
		p    = models.NewPermission()
		pcr  = models.NewPermissionCollectionResponse()
		spis = models.NewSharePointIdentitySet()
	)

	switch gv2 {
	case metadata.GV2User:
		i := models.NewIdentity()
		i.SetId(&entityID)
		i.SetDisplayName(&entityID)

		spis.SetUser(i)
	}

	p.SetGrantedToV2(spis)
	p.SetId(&permID)
	p.SetRoles(roles)

	pcr.SetValue([]models.Permissionable{p})

	return pcr
}
