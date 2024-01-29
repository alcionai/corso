package its

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type M365IntgSuite struct {
	tester.Suite
}

func TestM365IntgSuite(t *testing.T) {
	suite.Run(t, &M365IntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *M365IntgSuite) TestGetM365() {
	t := suite.T()
	result := GetM365(t)

	assert.NotEmpty(t, result.Acct)
	assert.NotEmpty(t, result.Creds)
	assert.NotEmpty(t, result.AC)
	assert.NotEmpty(t, result.GockAC)
	assert.NotEmpty(t, result.TenantID)

	assertIDs(
		t,
		result.Site,
		[]string{id, weburl, provider, driveid, driverootfolderid},
		[]string{})
	assertIDs(
		t,
		result.SecondarySite,
		[]string{id, weburl, provider, driveid, driverootfolderid},
		[]string{})
	assertIDs(
		t,
		result.Group,
		[]string{id, email, provider, testcontainerid},
		[]string{id, weburl, provider, displayname, driveid, driverootfolderid})
	assertIDs(
		t,
		result.SecondaryGroup,
		[]string{id, email, provider, testcontainerid},
		[]string{id, weburl, provider, displayname, driveid, driverootfolderid})
	assertIDs(
		t,
		result.User,
		[]string{id, email, provider, driveid, driverootfolderid},
		[]string{})
	assertIDs(
		t,
		result.SecondaryUser,
		[]string{id, email, provider, driveid, driverootfolderid},
		[]string{})
	assertIDs(
		t,
		result.TertiaryUser,
		[]string{id, email, provider, driveid, driverootfolderid},
		[]string{})
}

const (
	provider          = "provider"
	id                = "id"
	email             = "email"
	displayname       = "displayname"
	driveid           = "driveid"
	driverootfolderid = "driverootfolderid"
	testcontainerid   = "testcontainerid"
	weburl            = "weburl"
)

func assertIDs(
	t *testing.T,
	ids IDs,
	expect []string,
	expectRootSite []string,
) {
	assert.NotEmpty(t, ids)

	if slices.Contains(expect, provider) {
		assert.NotNil(t, ids.Provider)
		assert.NotEmpty(t, ids.Provider.ID())
		assert.NotEmpty(t, ids.Provider.Name())
	} else {
		assert.Nil(t, ids.Provider)
	}

	if slices.Contains(expect, id) {
		assert.NotEmpty(t, ids.ID)
	} else {
		assert.Empty(t, ids.ID)
	}

	if slices.Contains(expect, email) {
		assert.NotEmpty(t, ids.Email)
	} else {
		assert.Empty(t, ids.Email)
	}

	if slices.Contains(expect, driveid) {
		assert.NotEmpty(t, ids.DriveID)
	} else {
		assert.Empty(t, ids.DriveID)
	}

	if slices.Contains(expect, driverootfolderid) {
		assert.NotEmpty(t, ids.DriveRootFolderID)
	} else {
		assert.Empty(t, ids.DriveRootFolderID)
	}

	if slices.Contains(expect, testcontainerid) {
		assert.NotEmpty(t, ids.TestContainerID)
	} else {
		assert.Empty(t, ids.TestContainerID)
	}

	if slices.Contains(expect, weburl) {
		assert.NotEmpty(t, ids.WebURL)
	} else {
		assert.Empty(t, ids.WebURL)
	}

	if slices.Contains(expectRootSite, provider) {
		assert.NotNil(t, ids.RootSite.Provider)
		assert.NotEmpty(t, ids.RootSite.Provider.ID())
		assert.NotEmpty(t, ids.RootSite.Provider.Name())
	} else {
		assert.Nil(t, ids.RootSite.Provider)
	}

	if slices.Contains(expectRootSite, id) {
		assert.NotEmpty(t, ids.RootSite.ID)
	} else {
		assert.Empty(t, ids.RootSite.ID)
	}

	if slices.Contains(expectRootSite, driveid) {
		assert.NotEmpty(t, ids.RootSite.DriveID)
	} else {
		assert.Empty(t, ids.RootSite.DriveID)
	}

	if slices.Contains(expectRootSite, displayname) {
		assert.NotEmpty(t, ids.RootSite.DisplayName)
	} else {
		assert.Empty(t, ids.RootSite.DisplayName)
	}

	if slices.Contains(expectRootSite, driverootfolderid) {
		assert.NotEmpty(t, ids.RootSite.DriveRootFolderID)
	} else {
		assert.Empty(t, ids.RootSite.DriveRootFolderID)
	}

	if slices.Contains(expectRootSite, weburl) {
		assert.NotEmpty(t, ids.RootSite.WebURL)
	} else {
		assert.Empty(t, ids.RootSite.WebURL)
	}
}
