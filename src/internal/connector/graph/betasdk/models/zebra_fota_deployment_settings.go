package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaDeploymentSettings the Zebra FOTA deployment complex type that describes the settings required to create a FOTA deployment.
type ZebraFotaDeploymentSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Minimum battery level (%) required for both download and installation. Default: -1 (System defaults). Maximum is 100.
    batteryRuleMinimumBatteryLevelPercentage *int32
    // Flag indicating if charger is required. When set to false, the client can install updates whether the device is in or out of the charger. Applied only for installation. Defaults to false.
    batteryRuleRequireCharger *bool
    // Deploy update for devices with this model only.
    deviceModel *string
    // Represents various network types for Zebra FOTA deployment.
    downloadRuleNetworkType *ZebraFotaNetworkType
    // Date and time in the device time zone when the download will start (e.g., 2018-07-25T10:20:32). The default value is UTC now and the maximum is 10 days from deployment creation.
    downloadRuleStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A description provided by Zebra for the the firmware artifact to update the device to (e.g.: LifeGuard Update 120 (released 29-June-2022).
    firmwareTargetArtifactDescription *string
    // Deployment's Board Support Package (BSP. E.g.: '01.18.02.00'). Required only for custom update type.
    firmwareTargetBoardSupportPackageVersion *string
    // Target OS Version (e.g.: '8.1.0'). Required only for custom update type.
    firmwareTargetOsVersion *string
    // Target patch name (e.g.: 'U06'). Required only for custom update type.
    firmwareTargetPatch *string
    // Date and time in device time zone when the install will start. Default - download startDate if configured, otherwise defaults to NOW. Ignored when deployment update type was set to auto.
    installRuleStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Time of day after which the install cannot start. Possible range is 00:30:00 to 23:59:59. Should be greater than 'installRuleWindowStartTime' by 30 mins. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 23:59:59. Respected for all values of update type, including AUTO.
    installRuleWindowEndTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Time of day (00:00:00 - 23:30:00) when installation should begin. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 00:00:00. Respected for all values of update type, including AUTO.
    installRuleWindowStartTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The OdataType property
    odataType *string
    // Maximum 28 days. Default is 28 days. Sequence of dates are: 1) Download start date. 2) Install start date. 3) Schedule end date. If any of the values are not provided, the date provided in the preceding step of the sequence is used. If no values are provided, the string value of the current UTC is used.
    scheduleDurationInDays *int32
    // Represents various schedule modes for Zebra FOTA deployment.
    scheduleMode *ZebraFotaScheduleMode
    // This attribute indicates the deployment time offset (e.g.180 represents an offset of +03:00, and -270 represents an offset of -04:30). The time offset is the time timezone where the devices are located. The deployment start and end data uses this timezone
    timeZoneOffsetInMinutes *int32
    // Represents various update types for Zebra FOTA deployment.
    updateType *ZebraFotaUpdateType
}
// NewZebraFotaDeploymentSettings instantiates a new zebraFotaDeploymentSettings and sets the default values.
func NewZebraFotaDeploymentSettings()(*ZebraFotaDeploymentSettings) {
    m := &ZebraFotaDeploymentSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateZebraFotaDeploymentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaDeploymentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaDeploymentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ZebraFotaDeploymentSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBatteryRuleMinimumBatteryLevelPercentage gets the batteryRuleMinimumBatteryLevelPercentage property value. Minimum battery level (%) required for both download and installation. Default: -1 (System defaults). Maximum is 100.
func (m *ZebraFotaDeploymentSettings) GetBatteryRuleMinimumBatteryLevelPercentage()(*int32) {
    return m.batteryRuleMinimumBatteryLevelPercentage
}
// GetBatteryRuleRequireCharger gets the batteryRuleRequireCharger property value. Flag indicating if charger is required. When set to false, the client can install updates whether the device is in or out of the charger. Applied only for installation. Defaults to false.
func (m *ZebraFotaDeploymentSettings) GetBatteryRuleRequireCharger()(*bool) {
    return m.batteryRuleRequireCharger
}
// GetDeviceModel gets the deviceModel property value. Deploy update for devices with this model only.
func (m *ZebraFotaDeploymentSettings) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetDownloadRuleNetworkType gets the downloadRuleNetworkType property value. Represents various network types for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) GetDownloadRuleNetworkType()(*ZebraFotaNetworkType) {
    return m.downloadRuleNetworkType
}
// GetDownloadRuleStartDateTime gets the downloadRuleStartDateTime property value. Date and time in the device time zone when the download will start (e.g., 2018-07-25T10:20:32). The default value is UTC now and the maximum is 10 days from deployment creation.
func (m *ZebraFotaDeploymentSettings) GetDownloadRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.downloadRuleStartDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ZebraFotaDeploymentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["batteryRuleMinimumBatteryLevelPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuleMinimumBatteryLevelPercentage(val)
        }
        return nil
    }
    res["batteryRuleRequireCharger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuleRequireCharger(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["downloadRuleNetworkType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaNetworkType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadRuleNetworkType(val.(*ZebraFotaNetworkType))
        }
        return nil
    }
    res["downloadRuleStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadRuleStartDateTime(val)
        }
        return nil
    }
    res["firmwareTargetArtifactDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetArtifactDescription(val)
        }
        return nil
    }
    res["firmwareTargetBoardSupportPackageVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetBoardSupportPackageVersion(val)
        }
        return nil
    }
    res["firmwareTargetOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetOsVersion(val)
        }
        return nil
    }
    res["firmwareTargetPatch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetPatch(val)
        }
        return nil
    }
    res["installRuleStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallRuleStartDateTime(val)
        }
        return nil
    }
    res["installRuleWindowEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallRuleWindowEndTime(val)
        }
        return nil
    }
    res["installRuleWindowStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallRuleWindowStartTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["scheduleDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleDurationInDays(val)
        }
        return nil
    }
    res["scheduleMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaScheduleMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleMode(val.(*ZebraFotaScheduleMode))
        }
        return nil
    }
    res["timeZoneOffsetInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZoneOffsetInMinutes(val)
        }
        return nil
    }
    res["updateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaUpdateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateType(val.(*ZebraFotaUpdateType))
        }
        return nil
    }
    return res
}
// GetFirmwareTargetArtifactDescription gets the firmwareTargetArtifactDescription property value. A description provided by Zebra for the the firmware artifact to update the device to (e.g.: LifeGuard Update 120 (released 29-June-2022).
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetArtifactDescription()(*string) {
    return m.firmwareTargetArtifactDescription
}
// GetFirmwareTargetBoardSupportPackageVersion gets the firmwareTargetBoardSupportPackageVersion property value. Deployment's Board Support Package (BSP. E.g.: '01.18.02.00'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetBoardSupportPackageVersion()(*string) {
    return m.firmwareTargetBoardSupportPackageVersion
}
// GetFirmwareTargetOsVersion gets the firmwareTargetOsVersion property value. Target OS Version (e.g.: '8.1.0'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetOsVersion()(*string) {
    return m.firmwareTargetOsVersion
}
// GetFirmwareTargetPatch gets the firmwareTargetPatch property value. Target patch name (e.g.: 'U06'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetPatch()(*string) {
    return m.firmwareTargetPatch
}
// GetInstallRuleStartDateTime gets the installRuleStartDateTime property value. Date and time in device time zone when the install will start. Default - download startDate if configured, otherwise defaults to NOW. Ignored when deployment update type was set to auto.
func (m *ZebraFotaDeploymentSettings) GetInstallRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.installRuleStartDateTime
}
// GetInstallRuleWindowEndTime gets the installRuleWindowEndTime property value. Time of day after which the install cannot start. Possible range is 00:30:00 to 23:59:59. Should be greater than 'installRuleWindowStartTime' by 30 mins. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 23:59:59. Respected for all values of update type, including AUTO.
func (m *ZebraFotaDeploymentSettings) GetInstallRuleWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.installRuleWindowEndTime
}
// GetInstallRuleWindowStartTime gets the installRuleWindowStartTime property value. Time of day (00:00:00 - 23:30:00) when installation should begin. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 00:00:00. Respected for all values of update type, including AUTO.
func (m *ZebraFotaDeploymentSettings) GetInstallRuleWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.installRuleWindowStartTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ZebraFotaDeploymentSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetScheduleDurationInDays gets the scheduleDurationInDays property value. Maximum 28 days. Default is 28 days. Sequence of dates are: 1) Download start date. 2) Install start date. 3) Schedule end date. If any of the values are not provided, the date provided in the preceding step of the sequence is used. If no values are provided, the string value of the current UTC is used.
func (m *ZebraFotaDeploymentSettings) GetScheduleDurationInDays()(*int32) {
    return m.scheduleDurationInDays
}
// GetScheduleMode gets the scheduleMode property value. Represents various schedule modes for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) GetScheduleMode()(*ZebraFotaScheduleMode) {
    return m.scheduleMode
}
// GetTimeZoneOffsetInMinutes gets the timeZoneOffsetInMinutes property value. This attribute indicates the deployment time offset (e.g.180 represents an offset of +03:00, and -270 represents an offset of -04:30). The time offset is the time timezone where the devices are located. The deployment start and end data uses this timezone
func (m *ZebraFotaDeploymentSettings) GetTimeZoneOffsetInMinutes()(*int32) {
    return m.timeZoneOffsetInMinutes
}
// GetUpdateType gets the updateType property value. Represents various update types for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) GetUpdateType()(*ZebraFotaUpdateType) {
    return m.updateType
}
// Serialize serializes information the current object
func (m *ZebraFotaDeploymentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("batteryRuleMinimumBatteryLevelPercentage", m.GetBatteryRuleMinimumBatteryLevelPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("batteryRuleRequireCharger", m.GetBatteryRuleRequireCharger())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    if m.GetDownloadRuleNetworkType() != nil {
        cast := (*m.GetDownloadRuleNetworkType()).String()
        err := writer.WriteStringValue("downloadRuleNetworkType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("downloadRuleStartDateTime", m.GetDownloadRuleStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetArtifactDescription", m.GetFirmwareTargetArtifactDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetBoardSupportPackageVersion", m.GetFirmwareTargetBoardSupportPackageVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetOsVersion", m.GetFirmwareTargetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetPatch", m.GetFirmwareTargetPatch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("installRuleStartDateTime", m.GetInstallRuleStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("installRuleWindowEndTime", m.GetInstallRuleWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("installRuleWindowStartTime", m.GetInstallRuleWindowStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("scheduleDurationInDays", m.GetScheduleDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetScheduleMode() != nil {
        cast := (*m.GetScheduleMode()).String()
        err := writer.WriteStringValue("scheduleMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeZoneOffsetInMinutes", m.GetTimeZoneOffsetInMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateType() != nil {
        cast := (*m.GetUpdateType()).String()
        err := writer.WriteStringValue("updateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ZebraFotaDeploymentSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBatteryRuleMinimumBatteryLevelPercentage sets the batteryRuleMinimumBatteryLevelPercentage property value. Minimum battery level (%) required for both download and installation. Default: -1 (System defaults). Maximum is 100.
func (m *ZebraFotaDeploymentSettings) SetBatteryRuleMinimumBatteryLevelPercentage(value *int32)() {
    m.batteryRuleMinimumBatteryLevelPercentage = value
}
// SetBatteryRuleRequireCharger sets the batteryRuleRequireCharger property value. Flag indicating if charger is required. When set to false, the client can install updates whether the device is in or out of the charger. Applied only for installation. Defaults to false.
func (m *ZebraFotaDeploymentSettings) SetBatteryRuleRequireCharger(value *bool)() {
    m.batteryRuleRequireCharger = value
}
// SetDeviceModel sets the deviceModel property value. Deploy update for devices with this model only.
func (m *ZebraFotaDeploymentSettings) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetDownloadRuleNetworkType sets the downloadRuleNetworkType property value. Represents various network types for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) SetDownloadRuleNetworkType(value *ZebraFotaNetworkType)() {
    m.downloadRuleNetworkType = value
}
// SetDownloadRuleStartDateTime sets the downloadRuleStartDateTime property value. Date and time in the device time zone when the download will start (e.g., 2018-07-25T10:20:32). The default value is UTC now and the maximum is 10 days from deployment creation.
func (m *ZebraFotaDeploymentSettings) SetDownloadRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.downloadRuleStartDateTime = value
}
// SetFirmwareTargetArtifactDescription sets the firmwareTargetArtifactDescription property value. A description provided by Zebra for the the firmware artifact to update the device to (e.g.: LifeGuard Update 120 (released 29-June-2022).
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetArtifactDescription(value *string)() {
    m.firmwareTargetArtifactDescription = value
}
// SetFirmwareTargetBoardSupportPackageVersion sets the firmwareTargetBoardSupportPackageVersion property value. Deployment's Board Support Package (BSP. E.g.: '01.18.02.00'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetBoardSupportPackageVersion(value *string)() {
    m.firmwareTargetBoardSupportPackageVersion = value
}
// SetFirmwareTargetOsVersion sets the firmwareTargetOsVersion property value. Target OS Version (e.g.: '8.1.0'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetOsVersion(value *string)() {
    m.firmwareTargetOsVersion = value
}
// SetFirmwareTargetPatch sets the firmwareTargetPatch property value. Target patch name (e.g.: 'U06'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetPatch(value *string)() {
    m.firmwareTargetPatch = value
}
// SetInstallRuleStartDateTime sets the installRuleStartDateTime property value. Date and time in device time zone when the install will start. Default - download startDate if configured, otherwise defaults to NOW. Ignored when deployment update type was set to auto.
func (m *ZebraFotaDeploymentSettings) SetInstallRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.installRuleStartDateTime = value
}
// SetInstallRuleWindowEndTime sets the installRuleWindowEndTime property value. Time of day after which the install cannot start. Possible range is 00:30:00 to 23:59:59. Should be greater than 'installRuleWindowStartTime' by 30 mins. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 23:59:59. Respected for all values of update type, including AUTO.
func (m *ZebraFotaDeploymentSettings) SetInstallRuleWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.installRuleWindowEndTime = value
}
// SetInstallRuleWindowStartTime sets the installRuleWindowStartTime property value. Time of day (00:00:00 - 23:30:00) when installation should begin. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 00:00:00. Respected for all values of update type, including AUTO.
func (m *ZebraFotaDeploymentSettings) SetInstallRuleWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.installRuleWindowStartTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ZebraFotaDeploymentSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetScheduleDurationInDays sets the scheduleDurationInDays property value. Maximum 28 days. Default is 28 days. Sequence of dates are: 1) Download start date. 2) Install start date. 3) Schedule end date. If any of the values are not provided, the date provided in the preceding step of the sequence is used. If no values are provided, the string value of the current UTC is used.
func (m *ZebraFotaDeploymentSettings) SetScheduleDurationInDays(value *int32)() {
    m.scheduleDurationInDays = value
}
// SetScheduleMode sets the scheduleMode property value. Represents various schedule modes for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) SetScheduleMode(value *ZebraFotaScheduleMode)() {
    m.scheduleMode = value
}
// SetTimeZoneOffsetInMinutes sets the timeZoneOffsetInMinutes property value. This attribute indicates the deployment time offset (e.g.180 represents an offset of +03:00, and -270 represents an offset of -04:30). The time offset is the time timezone where the devices are located. The deployment start and end data uses this timezone
func (m *ZebraFotaDeploymentSettings) SetTimeZoneOffsetInMinutes(value *int32)() {
    m.timeZoneOffsetInMinutes = value
}
// SetUpdateType sets the updateType property value. Represents various update types for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) SetUpdateType(value *ZebraFotaUpdateType)() {
    m.updateType = value
}
