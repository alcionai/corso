package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementTemplateDetailedInfoable 
type ManagementTemplateDetailedInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*ManagementCategory)
    GetDisplayName()(*string)
    GetManagementTemplateId()(*string)
    GetOdataType()(*string)
    GetVersion()(*int32)
    SetCategory(value *ManagementCategory)()
    SetDisplayName(value *string)()
    SetManagementTemplateId(value *string)()
    SetOdataType(value *string)()
    SetVersion(value *int32)()
}
