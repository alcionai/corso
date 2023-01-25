package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineStateable 
type SecurityBaselineStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetSecurityBaselineTemplateId()(*string)
    GetSettingStates()([]SecurityBaselineSettingStateable)
    GetState()(*SecurityBaselineComplianceState)
    GetUserPrincipalName()(*string)
    SetDisplayName(value *string)()
    SetSecurityBaselineTemplateId(value *string)()
    SetSettingStates(value []SecurityBaselineSettingStateable)()
    SetState(value *SecurityBaselineComplianceState)()
    SetUserPrincipalName(value *string)()
}
