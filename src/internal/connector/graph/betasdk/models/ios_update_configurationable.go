package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosUpdateConfigurationable 
type IosUpdateConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveHoursEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetActiveHoursStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable)
    GetDesiredOsVersion()(*string)
    GetEnforcedSoftwareUpdateDelayInDays()(*int32)
    GetIsEnabled()(*bool)
    GetScheduledInstallDays()([]DayOfWeek)
    GetUpdateScheduleType()(*IosSoftwareUpdateScheduleType)
    GetUtcTimeOffsetInMinutes()(*int32)
    SetActiveHoursEnd(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetActiveHoursStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)()
    SetDesiredOsVersion(value *string)()
    SetEnforcedSoftwareUpdateDelayInDays(value *int32)()
    SetIsEnabled(value *bool)()
    SetScheduledInstallDays(value []DayOfWeek)()
    SetUpdateScheduleType(value *IosSoftwareUpdateScheduleType)()
    SetUtcTimeOffsetInMinutes(value *int32)()
}
