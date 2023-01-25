package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppliedConditionalAccessPolicyable 
type AppliedConditionalAccessPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationStrength()(AuthenticationStrengthable)
    GetConditionsNotSatisfied()(*ConditionalAccessConditions)
    GetConditionsSatisfied()(*ConditionalAccessConditions)
    GetDisplayName()(*string)
    GetEnforcedGrantControls()([]string)
    GetEnforcedSessionControls()([]string)
    GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable)
    GetId()(*string)
    GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable)
    GetOdataType()(*string)
    GetResult()(*AppliedConditionalAccessPolicyResult)
    GetSessionControlsNotSatisfied()([]string)
    SetAuthenticationStrength(value AuthenticationStrengthable)()
    SetConditionsNotSatisfied(value *ConditionalAccessConditions)()
    SetConditionsSatisfied(value *ConditionalAccessConditions)()
    SetDisplayName(value *string)()
    SetEnforcedGrantControls(value []string)()
    SetEnforcedSessionControls(value []string)()
    SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)()
    SetId(value *string)()
    SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)()
    SetOdataType(value *string)()
    SetResult(value *AppliedConditionalAccessPolicyResult)()
    SetSessionControlsNotSatisfied(value []string)()
}
