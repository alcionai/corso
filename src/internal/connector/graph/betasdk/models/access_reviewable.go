package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewable 
type AccessReviewable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessFlowTemplateId()(*string)
    GetCreatedBy()(UserIdentityable)
    GetDecisions()([]AccessReviewDecisionable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstances()([]AccessReviewable)
    GetMyDecisions()([]AccessReviewDecisionable)
    GetReviewedEntity()(Identityable)
    GetReviewers()([]AccessReviewReviewerable)
    GetReviewerType()(*string)
    GetSettings()(AccessReviewSettingsable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    SetBusinessFlowTemplateId(value *string)()
    SetCreatedBy(value UserIdentityable)()
    SetDecisions(value []AccessReviewDecisionable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstances(value []AccessReviewable)()
    SetMyDecisions(value []AccessReviewDecisionable)()
    SetReviewedEntity(value Identityable)()
    SetReviewers(value []AccessReviewReviewerable)()
    SetReviewerType(value *string)()
    SetSettings(value AccessReviewSettingsable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
}
