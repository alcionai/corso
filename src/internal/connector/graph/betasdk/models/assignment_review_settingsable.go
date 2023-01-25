package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentReviewSettingsable 
type AssignmentReviewSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessReviewTimeoutBehavior()(*AccessReviewTimeoutBehavior)
    GetDurationInDays()(*int32)
    GetIsAccessRecommendationEnabled()(*bool)
    GetIsApprovalJustificationRequired()(*bool)
    GetIsEnabled()(*bool)
    GetOdataType()(*string)
    GetRecurrenceType()(*string)
    GetReviewers()([]UserSetable)
    GetReviewerType()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessReviewTimeoutBehavior(value *AccessReviewTimeoutBehavior)()
    SetDurationInDays(value *int32)()
    SetIsAccessRecommendationEnabled(value *bool)()
    SetIsApprovalJustificationRequired(value *bool)()
    SetIsEnabled(value *bool)()
    SetOdataType(value *string)()
    SetRecurrenceType(value *string)()
    SetReviewers(value []UserSetable)()
    SetReviewerType(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
