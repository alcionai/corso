package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MatchingDlpRuleable 
type MatchingDlpRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]DlpActionInfoable)
    GetIsMostRestrictive()(*bool)
    GetOdataType()(*string)
    GetPolicyId()(*string)
    GetPolicyName()(*string)
    GetPriority()(*int32)
    GetRuleId()(*string)
    GetRuleMode()(*RuleMode)
    GetRuleName()(*string)
    SetActions(value []DlpActionInfoable)()
    SetIsMostRestrictive(value *bool)()
    SetOdataType(value *string)()
    SetPolicyId(value *string)()
    SetPolicyName(value *string)()
    SetPriority(value *int32)()
    SetRuleId(value *string)()
    SetRuleMode(value *RuleMode)()
    SetRuleName(value *string)()
}
