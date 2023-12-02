// funcs as the original structs in the msgraph-sdk-go package, which do not
// follow some of the golint rules.
//
//nolint:revive
package custom

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gotest.tools/v3/assert"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/tester"
)

type driveUnitSuite struct {
	tester.Suite
}

func TestDriveUnitSuite(t *testing.T) {
	suite.Run(t, &driveUnitSuite{
		Suite: tester.NewUnitSuite(t),
	})
}

func (suite *driveUnitSuite) TestToLiteDriveItemable() {
	id := "itemID"

	table := []struct {
		name         string
		itemFunc     func() models.DriveItemable
		validateFunc func(
			t *testing.T,
			expected models.DriveItemable,
			got LiteDriveItemable)
	}{
		{
			name: "nil item",
			itemFunc: func() models.DriveItemable {
				return nil
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.Nil(t, got)
			},
		},
		{
			name: "uninitialized values",
			itemFunc: func() models.DriveItemable {
				di := models.NewDriveItem()

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				assert.Equal(t, ptr.Val(got.GetId()), "")
				assert.Equal(t, ptr.Val(got.GetName()), "")
				assert.Equal(t, ptr.Val(got.GetSize()), int64(0))
				assert.Equal(t, ptr.Val(got.GetCreatedDateTime()), time.Time{})
				assert.Equal(t, ptr.Val(got.GetLastModifiedDateTime()), time.Time{})
				require.Nil(t, got.GetFolder())
				require.Nil(t, got.GetFile())
				require.Nil(t, got.GetPackageEscaped())
				require.Nil(t, got.GetShared())
				require.Nil(t, got.GetMalware())
				require.Nil(t, got.GetDeleted())
				require.Nil(t, got.GetRoot())
				require.Nil(t, got.GetCreatedBy())
				require.Nil(t, got.GetParentReference())
				assert.Equal(t, len(got.GetAdditionalData()), 0)
			},
		},
		{
			name: "ID, name, size, created, modified",
			itemFunc: func() models.DriveItemable {
				name := "itemName"
				size := int64(123)
				created := time.Now().Add(-time.Second).Truncate(time.Nanosecond)
				modified := time.Now().Truncate(time.Nanosecond)

				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetName(&name)
				di.SetSize(&size)
				di.SetCreatedDateTime(&created)
				di.SetLastModifiedDateTime(&modified)

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
				assert.Equal(t, ptr.Val(got.GetName()), ptr.Val(expected.GetName()))
				assert.Equal(t, ptr.Val(got.GetSize()), ptr.Val(expected.GetSize()))
				require.True(
					t,
					got.GetCreatedDateTime().Equal(ptr.Val(expected.GetCreatedDateTime())))
				require.True(
					t,
					got.GetLastModifiedDateTime().Equal(ptr.Val(expected.GetLastModifiedDateTime())))
			},
		},
		{
			name: "Folder item",
			itemFunc: func() models.DriveItemable {
				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetFolder(models.NewFolder())

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetFolder())
				require.Nil(t, got.GetFile())
				require.Nil(t, got.GetPackageEscaped())
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
			},
		},
		{
			name: "Package item",
			itemFunc: func() models.DriveItemable {
				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetPackageEscaped(models.NewPackageEscaped())

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetPackageEscaped())
				require.Nil(t, got.GetFile())
				require.Nil(t, got.GetFolder())
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
			},
		},
		{
			name: "File item",
			itemFunc: func() models.DriveItemable {
				mime := "mimeType"
				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetFile(models.NewFile())
				di.GetFile().SetMimeType(&mime)

				// additional data
				di.SetAdditionalData(map[string]interface{}{
					"@microsoft.graph.downloadUrl": "downloadURL",
				})

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetFile())
				require.Nil(t, got.GetFolder())
				require.Nil(t, got.GetPackageEscaped())
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
				assert.Equal(
					t,
					ptr.Val(got.GetFile().GetMimeType()),
					ptr.Val(expected.GetFile().GetMimeType()))

				// additional data
				urlExpected, err := str.AnyValueToString(
					"@microsoft.graph.downloadUrl",
					expected.GetAdditionalData())
				require.NoError(t, err)

				urlGot, err := str.AnyValueToString(
					"@microsoft.graph.downloadUrl",
					got.GetAdditionalData())
				require.NoError(t, err)

				assert.Equal(
					t,
					urlExpected,
					urlGot)
			},
		},

		{
			name: "Shared item",
			itemFunc: func() models.DriveItemable {
				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetShared(models.NewShared())
				di.SetFile(models.NewFile())

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetShared())
				require.NotNil(t, got.GetFile())
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
			},
		},
		{
			name: "Malware item",
			itemFunc: func() models.DriveItemable {
				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetMalware(models.NewMalware())
				di.SetFile(models.NewFile())

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetMalware())
				require.NotNil(t, got.GetFile())
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
			},
		},
		{
			name: "Deleted item",
			itemFunc: func() models.DriveItemable {
				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetDeleted(models.NewDeleted())
				di.SetFile(models.NewFile())

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetDeleted())
				require.NotNil(t, got.GetFile())
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
			},
		},
		{
			name: "Root item",
			itemFunc: func() models.DriveItemable {
				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetRoot(models.NewRoot())
				di.SetFolder(models.NewFolder())

				return di
			},

			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetRoot())
				require.NotNil(t, got.GetFolder())
				assert.Equal(t, ptr.Val(got.GetId()), ptr.Val(expected.GetId()))
			},
		},
		{
			name: "Parent reference",
			itemFunc: func() models.DriveItemable {
				parentID := "parentID"
				parentPath := "/parentPath"
				parentName := "parentName"
				parentDriveID := "parentDriveID"

				parentRef := models.NewItemReference()
				parentRef.SetId(&parentID)
				parentRef.SetPath(&parentPath)
				parentRef.SetName(&parentName)
				parentRef.SetDriveId(&parentDriveID)

				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetParentReference(parentRef)

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetParentReference())
				assert.Equal(
					t,
					ptr.Val(got.GetParentReference().GetId()),
					ptr.Val(expected.GetParentReference().GetId()))
				assert.Equal(
					t,
					ptr.Val(got.GetParentReference().GetPath()),
					ptr.Val(expected.GetParentReference().GetPath()))
				assert.Equal(
					t,
					ptr.Val(got.GetParentReference().GetName()),
					ptr.Val(expected.GetParentReference().GetName()))
				assert.Equal(
					t,
					ptr.Val(got.GetParentReference().GetDriveId()),
					ptr.Val(expected.GetParentReference().GetDriveId()))
			},
		},
		{
			name: "Created by",
			itemFunc: func() models.DriveItemable {
				createdBy := models.NewIdentitySet()

				createdBy.SetUser(models.NewUser())
				createdBy.GetUser().SetAdditionalData(map[string]interface{}{
					"email": "email@me",
				})

				di := models.NewDriveItem()

				di.SetId(&id)
				di.SetCreatedBy(createdBy)

				return di
			},
			validateFunc: func(
				t *testing.T,
				expected models.DriveItemable,
				got LiteDriveItemable,
			) {
				require.NotNil(t, got.GetCreatedBy())
				require.NotNil(t, got.GetCreatedBy().GetUser())
				emailExpected, err := str.AnyValueToString(
					"email",
					expected.GetCreatedBy().GetUser().GetAdditionalData())
				require.NoError(t, err)

				emailGot, err := str.AnyValueToString(
					"email",
					got.GetCreatedBy().GetUser().GetAdditionalData())
				require.NoError(t, err)

				assert.Equal(t, emailExpected, emailGot)
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			expected := test.itemFunc()
			got := ToLiteDriveItemable(expected)
			test.validateFunc(suite.T(), expected, got)
		})
	}
}
