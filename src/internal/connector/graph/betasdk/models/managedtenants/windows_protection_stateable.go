package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// WindowsProtectionStateable 
type WindowsProtectionStateable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAntiMalwareVersion()(*string)
    GetAttentionRequired()(*bool)
    GetDeviceDeleted()(*bool)
    GetDevicePropertyRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEngineVersion()(*string)
    GetFullScanOverdue()(*bool)
    GetFullScanRequired()(*bool)
    GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastFullScanSignatureVersion()(*string)
    GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastQuickScanSignatureVersion()(*string)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareProtectionEnabled()(*bool)
    GetManagedDeviceHealthState()(*string)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetNetworkInspectionSystemEnabled()(*bool)
    GetQuickScanOverdue()(*bool)
    GetRealTimeProtectionEnabled()(*bool)
    GetRebootRequired()(*bool)
    GetSignatureUpdateOverdue()(*bool)
    GetSignatureVersion()(*string)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    SetAntiMalwareVersion(value *string)()
    SetAttentionRequired(value *bool)()
    SetDeviceDeleted(value *bool)()
    SetDevicePropertyRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEngineVersion(value *string)()
    SetFullScanOverdue(value *bool)()
    SetFullScanRequired(value *bool)()
    SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastFullScanSignatureVersion(value *string)()
    SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastQuickScanSignatureVersion(value *string)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareProtectionEnabled(value *bool)()
    SetManagedDeviceHealthState(value *string)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetNetworkInspectionSystemEnabled(value *bool)()
    SetQuickScanOverdue(value *bool)()
    SetRealTimeProtectionEnabled(value *bool)()
    SetRebootRequired(value *bool)()
    SetSignatureUpdateOverdue(value *bool)()
    SetSignatureVersion(value *string)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
}
