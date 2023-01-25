package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomSecurityAttributeDefinitionable 
type CustomSecurityAttributeDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedValues()([]AllowedValueable)
    GetAttributeSet()(*string)
    GetDescription()(*string)
    GetIsCollection()(*bool)
    GetIsSearchable()(*bool)
    GetName()(*string)
    GetStatus()(*string)
    GetType()(*string)
    GetUsePreDefinedValuesOnly()(*bool)
    SetAllowedValues(value []AllowedValueable)()
    SetAttributeSet(value *string)()
    SetDescription(value *string)()
    SetIsCollection(value *bool)()
    SetIsSearchable(value *bool)()
    SetName(value *string)()
    SetStatus(value *string)()
    SetType(value *string)()
    SetUsePreDefinedValuesOnly(value *bool)()
}
