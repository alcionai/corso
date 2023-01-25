package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItemable 
type AccessReviewInstanceDecisionItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessReviewId()(*string)
    GetAppliedBy()(UserIdentityable)
    GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetApplyResult()(*string)
    GetDecision()(*string)
    GetInsights()([]GovernanceInsightable)
    GetInstance()(AccessReviewInstanceable)
    GetJustification()(*string)
    GetPrincipal()(Identityable)
    GetPrincipalLink()(*string)
    GetPrincipalResourceMembership()(DecisionItemPrincipalResourceMembershipable)
    GetRecommendation()(*string)
    GetResource()(AccessReviewInstanceDecisionItemResourceable)
    GetResourceLink()(*string)
    GetReviewedBy()(UserIdentityable)
    GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTarget()(AccessReviewInstanceDecisionItemTargetable)
    SetAccessReviewId(value *string)()
    SetAppliedBy(value UserIdentityable)()
    SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetApplyResult(value *string)()
    SetDecision(value *string)()
    SetInsights(value []GovernanceInsightable)()
    SetInstance(value AccessReviewInstanceable)()
    SetJustification(value *string)()
    SetPrincipal(value Identityable)()
    SetPrincipalLink(value *string)()
    SetPrincipalResourceMembership(value DecisionItemPrincipalResourceMembershipable)()
    SetRecommendation(value *string)()
    SetResource(value AccessReviewInstanceDecisionItemResourceable)()
    SetResourceLink(value *string)()
    SetReviewedBy(value UserIdentityable)()
    SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTarget(value AccessReviewInstanceDecisionItemTargetable)()
}
