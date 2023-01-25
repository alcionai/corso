package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateConfigurationable 
type MacOSSoftwareUpdateConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllOtherUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetConfigDataUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetCriticalUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetCustomUpdateTimeWindows()([]CustomUpdateTimeWindowable)
    GetFirmwareUpdateBehavior()(*MacOSSoftwareUpdateBehavior)
    GetUpdateScheduleType()(*MacOSSoftwareUpdateScheduleType)
    GetUpdateTimeWindowUtcOffsetInMinutes()(*int32)
    SetAllOtherUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetConfigDataUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetCriticalUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetCustomUpdateTimeWindows(value []CustomUpdateTimeWindowable)()
    SetFirmwareUpdateBehavior(value *MacOSSoftwareUpdateBehavior)()
    SetUpdateScheduleType(value *MacOSSoftwareUpdateScheduleType)()
    SetUpdateTimeWindowUtcOffsetInMinutes(value *int32)()
}
