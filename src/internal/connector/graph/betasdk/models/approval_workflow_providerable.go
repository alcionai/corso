package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalWorkflowProviderable 
type ApprovalWorkflowProviderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessFlows()([]BusinessFlowable)
    GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlowable)
    GetDisplayName()(*string)
    GetPolicyTemplates()([]GovernancePolicyTemplateable)
    SetBusinessFlows(value []BusinessFlowable)()
    SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlowable)()
    SetDisplayName(value *string)()
    SetPolicyTemplates(value []GovernancePolicyTemplateable)()
}
