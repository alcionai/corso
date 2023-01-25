package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppConfigurationSchemaItemable 
type AndroidManagedStoreAppConfigurationSchemaItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataType()(*AndroidManagedStoreAppConfigurationSchemaItemDataType)
    GetDefaultBoolValue()(*bool)
    GetDefaultIntValue()(*int32)
    GetDefaultStringArrayValue()([]string)
    GetDefaultStringValue()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIndex()(*int32)
    GetOdataType()(*string)
    GetParentIndex()(*int32)
    GetSchemaItemKey()(*string)
    GetSelections()([]KeyValuePairable)
    SetDataType(value *AndroidManagedStoreAppConfigurationSchemaItemDataType)()
    SetDefaultBoolValue(value *bool)()
    SetDefaultIntValue(value *int32)()
    SetDefaultStringArrayValue(value []string)()
    SetDefaultStringValue(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIndex(value *int32)()
    SetOdataType(value *string)()
    SetParentIndex(value *int32)()
    SetSchemaItemKey(value *string)()
    SetSelections(value []KeyValuePairable)()
}
