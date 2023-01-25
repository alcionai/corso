package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceMobileAppConfigurationable 
type ManagedDeviceMobileAppConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]ManagedDeviceMobileAppConfigurationAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatusable)
    GetDeviceStatusSummary()(ManagedDeviceMobileAppConfigurationDeviceSummaryable)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetTargetedMobileApps()([]string)
    GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatusable)
    GetUserStatusSummary()(ManagedDeviceMobileAppConfigurationUserSummaryable)
    GetVersion()(*int32)
    SetAssignments(value []ManagedDeviceMobileAppConfigurationAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatusable)()
    SetDeviceStatusSummary(value ManagedDeviceMobileAppConfigurationDeviceSummaryable)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetTargetedMobileApps(value []string)()
    SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatusable)()
    SetUserStatusSummary(value ManagedDeviceMobileAppConfigurationUserSummaryable)()
    SetVersion(value *int32)()
}
