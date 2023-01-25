package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementActionInfoable 
type ManagementActionInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManagementActionId()(*string)
    GetManagementTemplateId()(*string)
    GetManagementTemplateVersion()(*int32)
    GetOdataType()(*string)
    SetManagementActionId(value *string)()
    SetManagementTemplateId(value *string)()
    SetManagementTemplateVersion(value *int32)()
    SetOdataType(value *string)()
}
