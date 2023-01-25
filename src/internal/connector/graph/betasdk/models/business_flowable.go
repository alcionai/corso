package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessFlowable 
type BusinessFlowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomData()(*string)
    GetDeDuplicationId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetPolicy()(GovernancePolicyable)
    GetPolicyTemplateId()(*string)
    GetRecordVersion()(*string)
    GetSchemaId()(*string)
    GetSettings()(BusinessFlowSettingsable)
    SetCustomData(value *string)()
    SetDeDuplicationId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetPolicy(value GovernancePolicyable)()
    SetPolicyTemplateId(value *string)()
    SetRecordVersion(value *string)()
    SetSchemaId(value *string)()
    SetSettings(value BusinessFlowSettingsable)()
}
