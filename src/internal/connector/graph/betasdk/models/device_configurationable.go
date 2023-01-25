package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationable 
type DeviceConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceConfigurationAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceManagementApplicabilityRuleDeviceMode()(DeviceManagementApplicabilityRuleDeviceModeable)
    GetDeviceManagementApplicabilityRuleOsEdition()(DeviceManagementApplicabilityRuleOsEditionable)
    GetDeviceManagementApplicabilityRuleOsVersion()(DeviceManagementApplicabilityRuleOsVersionable)
    GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable)
    GetDeviceStatuses()([]DeviceConfigurationDeviceStatusable)
    GetDeviceStatusOverview()(DeviceConfigurationDeviceOverviewable)
    GetDisplayName()(*string)
    GetGroupAssignments()([]DeviceConfigurationGroupAssignmentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetSupportsScopeTags()(*bool)
    GetUserStatuses()([]DeviceConfigurationUserStatusable)
    GetUserStatusOverview()(DeviceConfigurationUserOverviewable)
    GetVersion()(*int32)
    SetAssignments(value []DeviceConfigurationAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceManagementApplicabilityRuleDeviceMode(value DeviceManagementApplicabilityRuleDeviceModeable)()
    SetDeviceManagementApplicabilityRuleOsEdition(value DeviceManagementApplicabilityRuleOsEditionable)()
    SetDeviceManagementApplicabilityRuleOsVersion(value DeviceManagementApplicabilityRuleOsVersionable)()
    SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)()
    SetDeviceStatuses(value []DeviceConfigurationDeviceStatusable)()
    SetDeviceStatusOverview(value DeviceConfigurationDeviceOverviewable)()
    SetDisplayName(value *string)()
    SetGroupAssignments(value []DeviceConfigurationGroupAssignmentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetSupportsScopeTags(value *bool)()
    SetUserStatuses(value []DeviceConfigurationUserStatusable)()
    SetUserStatusOverview(value DeviceConfigurationUserOverviewable)()
    SetVersion(value *int32)()
}
