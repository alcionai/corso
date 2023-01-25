package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalyable 
type UserExperienceAnalyticsAnomalyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnomalyFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAnomalyId()(*string)
    GetAnomalyLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAnomalyName()(*string)
    GetAnomalyType()(*UserExperienceAnalyticsAnomalyType)
    GetAssetName()(*string)
    GetAssetPublisher()(*string)
    GetAssetVersion()(*string)
    GetDetectionModelId()(*string)
    GetDeviceImpactedCount()(*int32)
    GetIssueId()(*string)
    GetSeverity()(*UserExperienceAnalyticsAnomalySeverity)
    GetState()(*UserExperienceAnalyticsAnomalyState)
    SetAnomalyFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAnomalyId(value *string)()
    SetAnomalyLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAnomalyName(value *string)()
    SetAnomalyType(value *UserExperienceAnalyticsAnomalyType)()
    SetAssetName(value *string)()
    SetAssetPublisher(value *string)()
    SetAssetVersion(value *string)()
    SetDetectionModelId(value *string)()
    SetDeviceImpactedCount(value *int32)()
    SetIssueId(value *string)()
    SetSeverity(value *UserExperienceAnalyticsAnomalySeverity)()
    SetState(value *UserExperienceAnalyticsAnomalyState)()
}
