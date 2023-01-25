package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkloadActionable 
type WorkloadActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionId()(*string)
    GetCategory()(*WorkloadActionCategory)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLicenses()([]string)
    GetOdataType()(*string)
    GetService()(*string)
    GetSettings()([]Settingable)
    SetActionId(value *string)()
    SetCategory(value *WorkloadActionCategory)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLicenses(value []string)()
    SetOdataType(value *string)()
    SetService(value *string)()
    SetSettings(value []Settingable)()
}
