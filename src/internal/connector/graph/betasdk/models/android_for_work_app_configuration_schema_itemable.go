package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkAppConfigurationSchemaItemable 
type AndroidForWorkAppConfigurationSchemaItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataType()(*AndroidForWorkAppConfigurationSchemaItemDataType)
    GetDefaultBoolValue()(*bool)
    GetDefaultIntValue()(*int32)
    GetDefaultStringArrayValue()([]string)
    GetDefaultStringValue()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetSchemaItemKey()(*string)
    GetSelections()([]KeyValuePairable)
    SetDataType(value *AndroidForWorkAppConfigurationSchemaItemDataType)()
    SetDefaultBoolValue(value *bool)()
    SetDefaultIntValue(value *int32)()
    SetDefaultStringArrayValue(value []string)()
    SetDefaultStringValue(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetSchemaItemKey(value *string)()
    SetSelections(value []KeyValuePairable)()
}
