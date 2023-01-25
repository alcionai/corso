package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable 
type DeviceManagementConfigurationIntegerSettingValueDefinitionTemplateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxValue()(*int32)
    GetMinValue()(*int32)
    GetOdataType()(*string)
    SetMaxValue(value *int32)()
    SetMinValue(value *int32)()
    SetOdataType(value *string)()
}
