package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicyTemplateReferenceable 
type DeviceManagementConfigurationPolicyTemplateReferenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetTemplateDisplayName()(*string)
    GetTemplateDisplayVersion()(*string)
    GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily)
    GetTemplateId()(*string)
    SetOdataType(value *string)()
    SetTemplateDisplayName(value *string)()
    SetTemplateDisplayVersion(value *string)()
    SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)()
    SetTemplateId(value *string)()
}
