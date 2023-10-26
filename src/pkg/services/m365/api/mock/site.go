package mock

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

func UserIdentity(userID string, userEmail string) models.IdentitySetable {
	user := models.NewIdentitySet()
	userIdentity := models.NewUserIdentity()
	userIdentity.SetId(ptr.To(userID))

	if len(userEmail) > 0 {
		userIdentity.SetAdditionalData(map[string]any{
			"email": userEmail,
		})
	}

	user.SetUser(userIdentity)

	return user
}

func GroupIdentitySet(groupID string, groupEmail string) models.IdentitySetable {
	groupData := map[string]any{}
	if len(groupEmail) > 0 {
		groupData["email"] = groupEmail
	}

	if len(groupID) > 0 {
		groupData["id"] = groupID
	}

	group := models.NewIdentitySet()
	group.SetAdditionalData(map[string]any{"group": groupData})

	return group
}

func DummySite(owner models.IdentitySetable) models.Siteable {
	site := models.NewSite()
	site.SetId(ptr.To("id"))
	site.SetWebUrl(ptr.To("weburl"))
	site.SetDisplayName(ptr.To("displayname"))

	drive := models.NewDrive()
	drive.SetOwner(owner)
	site.SetDrive(drive)

	return site
}
