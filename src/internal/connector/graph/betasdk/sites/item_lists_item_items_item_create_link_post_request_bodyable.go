package sites

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ItemListsItemItemsItemCreateLinkPostRequestBodyable 
type ItemListsItemItemsItemCreateLinkPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPassword()(*string)
    GetRecipients()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DriveRecipientable)
    GetRetainInheritedPermissions()(*bool)
    GetScope()(*string)
    GetType()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPassword(value *string)()
    SetRecipients(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DriveRecipientable)()
    SetRetainInheritedPermissions(value *bool)()
    SetScope(value *string)()
    SetType(value *string)()
}
