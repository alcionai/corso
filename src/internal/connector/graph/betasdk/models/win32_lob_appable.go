package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppable 
type Win32LobAppable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAvailableUninstall()(*bool)
    GetApplicableArchitectures()(*WindowsArchitecture)
    GetDetectionRules()([]Win32LobAppDetectionable)
    GetDisplayVersion()(*string)
    GetInstallCommandLine()(*string)
    GetInstallExperience()(Win32LobAppInstallExperienceable)
    GetMinimumCpuSpeedInMHz()(*int32)
    GetMinimumFreeDiskSpaceInMB()(*int32)
    GetMinimumMemoryInMB()(*int32)
    GetMinimumNumberOfProcessors()(*int32)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    GetMinimumSupportedWindowsRelease()(*string)
    GetMsiInformation()(Win32LobAppMsiInformationable)
    GetRequirementRules()([]Win32LobAppRequirementable)
    GetReturnCodes()([]Win32LobAppReturnCodeable)
    GetRules()([]Win32LobAppRuleable)
    GetSetupFilePath()(*string)
    GetUninstallCommandLine()(*string)
    SetAllowAvailableUninstall(value *bool)()
    SetApplicableArchitectures(value *WindowsArchitecture)()
    SetDetectionRules(value []Win32LobAppDetectionable)()
    SetDisplayVersion(value *string)()
    SetInstallCommandLine(value *string)()
    SetInstallExperience(value Win32LobAppInstallExperienceable)()
    SetMinimumCpuSpeedInMHz(value *int32)()
    SetMinimumFreeDiskSpaceInMB(value *int32)()
    SetMinimumMemoryInMB(value *int32)()
    SetMinimumNumberOfProcessors(value *int32)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
    SetMinimumSupportedWindowsRelease(value *string)()
    SetMsiInformation(value Win32LobAppMsiInformationable)()
    SetRequirementRules(value []Win32LobAppRequirementable)()
    SetReturnCodes(value []Win32LobAppReturnCodeable)()
    SetRules(value []Win32LobAppRuleable)()
    SetSetupFilePath(value *string)()
    SetUninstallCommandLine(value *string)()
}
