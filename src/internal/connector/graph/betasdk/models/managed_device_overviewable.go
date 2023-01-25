package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceOverviewable 
type ManagedDeviceOverviewable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceExchangeAccessStateSummary()(DeviceExchangeAccessStateSummaryable)
    GetDeviceOperatingSystemSummary()(DeviceOperatingSystemSummaryable)
    GetDualEnrolledDeviceCount()(*int32)
    GetEnrolledDeviceCount()(*int32)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceModelsAndManufacturers()(ManagedDeviceModelsAndManufacturersable)
    GetMdmEnrolledCount()(*int32)
    SetDeviceExchangeAccessStateSummary(value DeviceExchangeAccessStateSummaryable)()
    SetDeviceOperatingSystemSummary(value DeviceOperatingSystemSummaryable)()
    SetDualEnrolledDeviceCount(value *int32)()
    SetEnrolledDeviceCount(value *int32)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceModelsAndManufacturers(value ManagedDeviceModelsAndManufacturersable)()
    SetMdmEnrolledCount(value *int32)()
}
