package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceProtectionOverviewable 
type DeviceProtectionOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCleanDeviceCount()(*int32)
    GetCriticalFailuresDeviceCount()(*int32)
    GetInactiveThreatAgentDeviceCount()(*int32)
    GetOdataType()(*string)
    GetPendingFullScanDeviceCount()(*int32)
    GetPendingManualStepsDeviceCount()(*int32)
    GetPendingOfflineScanDeviceCount()(*int32)
    GetPendingQuickScanDeviceCount()(*int32)
    GetPendingRestartDeviceCount()(*int32)
    GetPendingSignatureUpdateDeviceCount()(*int32)
    GetTotalReportedDeviceCount()(*int32)
    GetUnknownStateThreatAgentDeviceCount()(*int32)
    SetCleanDeviceCount(value *int32)()
    SetCriticalFailuresDeviceCount(value *int32)()
    SetInactiveThreatAgentDeviceCount(value *int32)()
    SetOdataType(value *string)()
    SetPendingFullScanDeviceCount(value *int32)()
    SetPendingManualStepsDeviceCount(value *int32)()
    SetPendingOfflineScanDeviceCount(value *int32)()
    SetPendingQuickScanDeviceCount(value *int32)()
    SetPendingRestartDeviceCount(value *int32)()
    SetPendingSignatureUpdateDeviceCount(value *int32)()
    SetTotalReportedDeviceCount(value *int32)()
    SetUnknownStateThreatAgentDeviceCount(value *int32)()
}
