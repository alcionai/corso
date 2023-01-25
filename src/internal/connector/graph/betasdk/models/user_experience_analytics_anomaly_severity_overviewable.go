package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalySeverityOverviewable 
type UserExperienceAnalyticsAnomalySeverityOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHighSeverityAnomalyCount()(*int32)
    GetInformationalSeverityAnomalyCount()(*int32)
    GetLowSeverityAnomalyCount()(*int32)
    GetMediumSeverityAnomalyCount()(*int32)
    GetOdataType()(*string)
    SetHighSeverityAnomalyCount(value *int32)()
    SetInformationalSeverityAnomalyCount(value *int32)()
    SetLowSeverityAnomalyCount(value *int32)()
    SetMediumSeverityAnomalyCount(value *int32)()
    SetOdataType(value *string)()
}
