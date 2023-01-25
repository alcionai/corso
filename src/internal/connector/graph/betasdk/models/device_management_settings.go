package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettings 
type DeviceManagementSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The property to determine if Android device administrator enrollment is enabled for this account.
    androidDeviceAdministratorEnrollmentEnabled *bool
    // Provider type for Derived Credentials.
    derivedCredentialProvider *DerivedCredentialProviderType
    // The Derived Credential Provider self-service URI.
    derivedCredentialUrl *string
    // The number of days a device is allowed to go without checking in to remain compliant.
    deviceComplianceCheckinThresholdDays *int32
    // When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
    deviceInactivityBeforeRetirementInDay *int32
    // Determines whether the autopilot diagnostic feature is enabled or not.
    enableAutopilotDiagnostics *bool
    // Determines whether the device group membership report feature is enabled or not.
    enableDeviceGroupMembershipReport *bool
    // Determines whether the enhanced troubleshooting UX is enabled or not.
    enableEnhancedTroubleshootingExperience *bool
    // Determines whether the log collection feature should be available for use.
    enableLogCollection *bool
    // Is feature enabled or not for enhanced jailbreak detection.
    enhancedJailBreak *bool
    // The property to determine whether to ignore unsupported compliance settings on certian models of devices.
    ignoreDevicesForUnsupportedSettingsEnabled *bool
    // Is feature enabled or not for scheduled action for rule.
    isScheduledActionEnabled *bool
    // The OdataType property
    odataType *string
    // Device should be noncompliant when there is no compliance policy targeted when this is true
    secureByDefault *bool
}
// NewDeviceManagementSettings instantiates a new deviceManagementSettings and sets the default values.
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAndroidDeviceAdministratorEnrollmentEnabled gets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
func (m *DeviceManagementSettings) GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool) {
    return m.androidDeviceAdministratorEnrollmentEnabled
}
// GetDerivedCredentialProvider gets the derivedCredentialProvider property value. Provider type for Derived Credentials.
func (m *DeviceManagementSettings) GetDerivedCredentialProvider()(*DerivedCredentialProviderType) {
    return m.derivedCredentialProvider
}
// GetDerivedCredentialUrl gets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
func (m *DeviceManagementSettings) GetDerivedCredentialUrl()(*string) {
    return m.derivedCredentialUrl
}
// GetDeviceComplianceCheckinThresholdDays gets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    return m.deviceComplianceCheckinThresholdDays
}
// GetDeviceInactivityBeforeRetirementInDay gets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
func (m *DeviceManagementSettings) GetDeviceInactivityBeforeRetirementInDay()(*int32) {
    return m.deviceInactivityBeforeRetirementInDay
}
// GetEnableAutopilotDiagnostics gets the enableAutopilotDiagnostics property value. Determines whether the autopilot diagnostic feature is enabled or not.
func (m *DeviceManagementSettings) GetEnableAutopilotDiagnostics()(*bool) {
    return m.enableAutopilotDiagnostics
}
// GetEnableDeviceGroupMembershipReport gets the enableDeviceGroupMembershipReport property value. Determines whether the device group membership report feature is enabled or not.
func (m *DeviceManagementSettings) GetEnableDeviceGroupMembershipReport()(*bool) {
    return m.enableDeviceGroupMembershipReport
}
// GetEnableEnhancedTroubleshootingExperience gets the enableEnhancedTroubleshootingExperience property value. Determines whether the enhanced troubleshooting UX is enabled or not.
func (m *DeviceManagementSettings) GetEnableEnhancedTroubleshootingExperience()(*bool) {
    return m.enableEnhancedTroubleshootingExperience
}
// GetEnableLogCollection gets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
func (m *DeviceManagementSettings) GetEnableLogCollection()(*bool) {
    return m.enableLogCollection
}
// GetEnhancedJailBreak gets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
func (m *DeviceManagementSettings) GetEnhancedJailBreak()(*bool) {
    return m.enhancedJailBreak
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["androidDeviceAdministratorEnrollmentEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceAdministratorEnrollmentEnabled(val)
        }
        return nil
    }
    res["derivedCredentialProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDerivedCredentialProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialProvider(val.(*DerivedCredentialProviderType))
        }
        return nil
    }
    res["derivedCredentialUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialUrl(val)
        }
        return nil
    }
    res["deviceComplianceCheckinThresholdDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceCheckinThresholdDays(val)
        }
        return nil
    }
    res["deviceInactivityBeforeRetirementInDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceInactivityBeforeRetirementInDay(val)
        }
        return nil
    }
    res["enableAutopilotDiagnostics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAutopilotDiagnostics(val)
        }
        return nil
    }
    res["enableDeviceGroupMembershipReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDeviceGroupMembershipReport(val)
        }
        return nil
    }
    res["enableEnhancedTroubleshootingExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableEnhancedTroubleshootingExperience(val)
        }
        return nil
    }
    res["enableLogCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableLogCollection(val)
        }
        return nil
    }
    res["enhancedJailBreak"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnhancedJailBreak(val)
        }
        return nil
    }
    res["ignoreDevicesForUnsupportedSettingsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreDevicesForUnsupportedSettingsEnabled(val)
        }
        return nil
    }
    res["isScheduledActionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScheduledActionEnabled(val)
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
    res["secureByDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureByDefault(val)
        }
        return nil
    }
    return res
}
// GetIgnoreDevicesForUnsupportedSettingsEnabled gets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
func (m *DeviceManagementSettings) GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool) {
    return m.ignoreDevicesForUnsupportedSettingsEnabled
}
// GetIsScheduledActionEnabled gets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    return m.isScheduledActionEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetSecureByDefault gets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    return m.secureByDefault
}
// Serialize serializes information the current object
func (m *DeviceManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("androidDeviceAdministratorEnrollmentEnabled", m.GetAndroidDeviceAdministratorEnrollmentEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDerivedCredentialProvider() != nil {
        cast := (*m.GetDerivedCredentialProvider()).String()
        err := writer.WriteStringValue("derivedCredentialProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("derivedCredentialUrl", m.GetDerivedCredentialUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceComplianceCheckinThresholdDays", m.GetDeviceComplianceCheckinThresholdDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceInactivityBeforeRetirementInDay", m.GetDeviceInactivityBeforeRetirementInDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableAutopilotDiagnostics", m.GetEnableAutopilotDiagnostics())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableDeviceGroupMembershipReport", m.GetEnableDeviceGroupMembershipReport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableEnhancedTroubleshootingExperience", m.GetEnableEnhancedTroubleshootingExperience())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableLogCollection", m.GetEnableLogCollection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enhancedJailBreak", m.GetEnhancedJailBreak())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreDevicesForUnsupportedSettingsEnabled", m.GetIgnoreDevicesForUnsupportedSettingsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScheduledActionEnabled", m.GetIsScheduledActionEnabled())
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
        err := writer.WriteBoolValue("secureByDefault", m.GetSecureByDefault())
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
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAndroidDeviceAdministratorEnrollmentEnabled sets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
func (m *DeviceManagementSettings) SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)() {
    m.androidDeviceAdministratorEnrollmentEnabled = value
}
// SetDerivedCredentialProvider sets the derivedCredentialProvider property value. Provider type for Derived Credentials.
func (m *DeviceManagementSettings) SetDerivedCredentialProvider(value *DerivedCredentialProviderType)() {
    m.derivedCredentialProvider = value
}
// SetDerivedCredentialUrl sets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
func (m *DeviceManagementSettings) SetDerivedCredentialUrl(value *string)() {
    m.derivedCredentialUrl = value
}
// SetDeviceComplianceCheckinThresholdDays sets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    m.deviceComplianceCheckinThresholdDays = value
}
// SetDeviceInactivityBeforeRetirementInDay sets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
func (m *DeviceManagementSettings) SetDeviceInactivityBeforeRetirementInDay(value *int32)() {
    m.deviceInactivityBeforeRetirementInDay = value
}
// SetEnableAutopilotDiagnostics sets the enableAutopilotDiagnostics property value. Determines whether the autopilot diagnostic feature is enabled or not.
func (m *DeviceManagementSettings) SetEnableAutopilotDiagnostics(value *bool)() {
    m.enableAutopilotDiagnostics = value
}
// SetEnableDeviceGroupMembershipReport sets the enableDeviceGroupMembershipReport property value. Determines whether the device group membership report feature is enabled or not.
func (m *DeviceManagementSettings) SetEnableDeviceGroupMembershipReport(value *bool)() {
    m.enableDeviceGroupMembershipReport = value
}
// SetEnableEnhancedTroubleshootingExperience sets the enableEnhancedTroubleshootingExperience property value. Determines whether the enhanced troubleshooting UX is enabled or not.
func (m *DeviceManagementSettings) SetEnableEnhancedTroubleshootingExperience(value *bool)() {
    m.enableEnhancedTroubleshootingExperience = value
}
// SetEnableLogCollection sets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
func (m *DeviceManagementSettings) SetEnableLogCollection(value *bool)() {
    m.enableLogCollection = value
}
// SetEnhancedJailBreak sets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
func (m *DeviceManagementSettings) SetEnhancedJailBreak(value *bool)() {
    m.enhancedJailBreak = value
}
// SetIgnoreDevicesForUnsupportedSettingsEnabled sets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
func (m *DeviceManagementSettings) SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)() {
    m.ignoreDevicesForUnsupportedSettingsEnabled = value
}
// SetIsScheduledActionEnabled sets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    m.isScheduledActionEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSecureByDefault sets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    m.secureByDefault = value
}
