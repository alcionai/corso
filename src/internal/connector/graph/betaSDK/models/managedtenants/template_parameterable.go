package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TemplateParameterable 
type TemplateParameterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetJsonAllowedValues()(*string)
    GetJsonDefaultValue()(*string)
    GetOdataType()(*string)
    GetValueType()(*ManagementParameterValueType)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetJsonAllowedValues(value *string)()
    SetJsonDefaultValue(value *string)()
    SetOdataType(value *string)()
    SetValueType(value *ManagementParameterValueType)()
}
