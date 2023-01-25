package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCompliancePolicyable 
type DeviceManagementCompliancePolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceManagementConfigurationPolicyAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreationSource()(*string)
    GetDescription()(*string)
    GetIsAssigned()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetRoleScopeTagIds()([]string)
    GetScheduledActionsForRule()([]DeviceManagementComplianceScheduledActionForRuleable)
    GetSettingCount()(*int32)
    GetSettings()([]DeviceManagementConfigurationSettingable)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    SetAssignments(value []DeviceManagementConfigurationPolicyAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreationSource(value *string)()
    SetDescription(value *string)()
    SetIsAssigned(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetRoleScopeTagIds(value []string)()
    SetScheduledActionsForRule(value []DeviceManagementComplianceScheduledActionForRuleable)()
    SetSettingCount(value *int32)()
    SetSettings(value []DeviceManagementConfigurationSettingable)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
}
