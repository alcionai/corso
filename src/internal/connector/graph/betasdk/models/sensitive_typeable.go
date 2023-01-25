package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitiveTypeable 
type SensitiveTypeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassificationMethod()(*ClassificationMethod)
    GetDescription()(*string)
    GetName()(*string)
    GetPublisherName()(*string)
    GetRulePackageId()(*string)
    GetRulePackageType()(*string)
    GetScope()(*SensitiveTypeScope)
    GetSensitiveTypeSource()(*SensitiveTypeSource)
    GetState()(*string)
    SetClassificationMethod(value *ClassificationMethod)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetPublisherName(value *string)()
    SetRulePackageId(value *string)()
    SetRulePackageType(value *string)()
    SetScope(value *SensitiveTypeScope)()
    SetSensitiveTypeSource(value *SensitiveTypeSource)()
    SetState(value *string)()
}
