package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterEvaluationSummaryable 
type AssignmentFilterEvaluationSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterDisplayName()(*string)
    GetAssignmentFilterId()(*string)
    GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAssignmentFilterPlatform()(*DevicePlatformType)
    GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType)
    GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResultable)
    GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEvaluationResult()(*AssignmentFilterEvaluationResult)
    GetOdataType()(*string)
    SetAssignmentFilterDisplayName(value *string)()
    SetAssignmentFilterId(value *string)()
    SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAssignmentFilterPlatform(value *DevicePlatformType)()
    SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)()
    SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResultable)()
    SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEvaluationResult(value *AssignmentFilterEvaluationResult)()
    SetOdataType(value *string)()
}
