package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalyDeviceable 
type UserExperienceAnalyticsAnomalyDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnomalyId()(*string)
    GetAnomalyOnDeviceFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAnomalyOnDeviceLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceId()(*string)
    GetDeviceManufacturer()(*string)
    GetDeviceModel()(*string)
    GetDeviceName()(*string)
    GetOsName()(*string)
    GetOsVersion()(*string)
    SetAnomalyId(value *string)()
    SetAnomalyOnDeviceFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAnomalyOnDeviceLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceId(value *string)()
    SetDeviceManufacturer(value *string)()
    SetDeviceModel(value *string)()
    SetDeviceName(value *string)()
    SetOsName(value *string)()
    SetOsVersion(value *string)()
}
