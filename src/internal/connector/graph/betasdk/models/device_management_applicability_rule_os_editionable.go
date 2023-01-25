package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementApplicabilityRuleOsEditionable 
type DeviceManagementApplicabilityRuleOsEditionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetOdataType()(*string)
    GetOsEditionTypes()([]Windows10EditionType)
    GetRuleType()(*DeviceManagementApplicabilityRuleType)
    SetName(value *string)()
    SetOdataType(value *string)()
    SetOsEditionTypes(value []Windows10EditionType)()
    SetRuleType(value *DeviceManagementApplicabilityRuleType)()
}
