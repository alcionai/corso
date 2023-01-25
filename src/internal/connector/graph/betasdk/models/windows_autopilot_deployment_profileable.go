package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfileable 
type WindowsAutopilotDeploymentProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedDevices()([]WindowsAutopilotDeviceIdentityable)
    GetAssignments()([]WindowsAutopilotDeploymentProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceNameTemplate()(*string)
    GetDeviceType()(*WindowsAutopilotDeviceType)
    GetDisplayName()(*string)
    GetEnableWhiteGlove()(*bool)
    GetEnrollmentStatusScreenSettings()(WindowsEnrollmentStatusScreenSettingsable)
    GetExtractHardwareHash()(*bool)
    GetLanguage()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementServiceAppId()(*string)
    GetOutOfBoxExperienceSettings()(OutOfBoxExperienceSettingsable)
    GetRoleScopeTagIds()([]string)
    SetAssignedDevices(value []WindowsAutopilotDeviceIdentityable)()
    SetAssignments(value []WindowsAutopilotDeploymentProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceNameTemplate(value *string)()
    SetDeviceType(value *WindowsAutopilotDeviceType)()
    SetDisplayName(value *string)()
    SetEnableWhiteGlove(value *bool)()
    SetEnrollmentStatusScreenSettings(value WindowsEnrollmentStatusScreenSettingsable)()
    SetExtractHardwareHash(value *bool)()
    SetLanguage(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementServiceAppId(value *string)()
    SetOutOfBoxExperienceSettings(value OutOfBoxExperienceSettingsable)()
    SetRoleScopeTagIds(value []string)()
}
