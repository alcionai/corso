package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewSetable 
type AccessReviewSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDecisions()([]AccessReviewInstanceDecisionItemable)
    GetDefinitions()([]AccessReviewScheduleDefinitionable)
    GetHistoryDefinitions()([]AccessReviewHistoryDefinitionable)
    GetPolicy()(AccessReviewPolicyable)
    SetDecisions(value []AccessReviewInstanceDecisionItemable)()
    SetDefinitions(value []AccessReviewScheduleDefinitionable)()
    SetHistoryDefinitions(value []AccessReviewHistoryDefinitionable)()
    SetPolicy(value AccessReviewPolicyable)()
}
