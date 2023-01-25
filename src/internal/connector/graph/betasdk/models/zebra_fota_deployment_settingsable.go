package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaDeploymentSettingsable 
type ZebraFotaDeploymentSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBatteryRuleMinimumBatteryLevelPercentage()(*int32)
    GetBatteryRuleRequireCharger()(*bool)
    GetDeviceModel()(*string)
    GetDownloadRuleNetworkType()(*ZebraFotaNetworkType)
    GetDownloadRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFirmwareTargetArtifactDescription()(*string)
    GetFirmwareTargetBoardSupportPackageVersion()(*string)
    GetFirmwareTargetOsVersion()(*string)
    GetFirmwareTargetPatch()(*string)
    GetInstallRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstallRuleWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetInstallRuleWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetOdataType()(*string)
    GetScheduleDurationInDays()(*int32)
    GetScheduleMode()(*ZebraFotaScheduleMode)
    GetTimeZoneOffsetInMinutes()(*int32)
    GetUpdateType()(*ZebraFotaUpdateType)
    SetBatteryRuleMinimumBatteryLevelPercentage(value *int32)()
    SetBatteryRuleRequireCharger(value *bool)()
    SetDeviceModel(value *string)()
    SetDownloadRuleNetworkType(value *ZebraFotaNetworkType)()
    SetDownloadRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFirmwareTargetArtifactDescription(value *string)()
    SetFirmwareTargetBoardSupportPackageVersion(value *string)()
    SetFirmwareTargetOsVersion(value *string)()
    SetFirmwareTargetPatch(value *string)()
    SetInstallRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstallRuleWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetInstallRuleWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetOdataType(value *string)()
    SetScheduleDurationInDays(value *int32)()
    SetScheduleMode(value *ZebraFotaScheduleMode)()
    SetTimeZoneOffsetInMinutes(value *int32)()
    SetUpdateType(value *ZebraFotaUpdateType)()
}
