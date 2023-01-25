package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineSettingStateable 
type SecurityBaselineSettingStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContributingPolicies()([]SecurityBaselineContributingPolicyable)
    GetErrorCode()(*string)
    GetSettingCategoryId()(*string)
    GetSettingCategoryName()(*string)
    GetSettingId()(*string)
    GetSettingName()(*string)
    GetSourcePolicies()([]SettingSourceable)
    GetState()(*SecurityBaselineComplianceState)
    SetContributingPolicies(value []SecurityBaselineContributingPolicyable)()
    SetErrorCode(value *string)()
    SetSettingCategoryId(value *string)()
    SetSettingCategoryName(value *string)()
    SetSettingId(value *string)()
    SetSettingName(value *string)()
    SetSourcePolicies(value []SettingSourceable)()
    SetState(value *SecurityBaselineComplianceState)()
}
