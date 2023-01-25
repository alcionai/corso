package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderAdvancedThreatProtectionConfiguration 
type WindowsDefenderAdvancedThreatProtectionConfiguration struct {
    DeviceConfiguration
    // Auto populate onboarding blob programmatically from Advanced Threat protection service
    advancedThreatProtectionAutoPopulateOnboardingBlob *bool
    // Windows Defender AdvancedThreatProtection Offboarding Blob.
    advancedThreatProtectionOffboardingBlob *string
    // Name of the file from which AdvancedThreatProtectionOffboardingBlob was obtained.
    advancedThreatProtectionOffboardingFilename *string
    // Windows Defender AdvancedThreatProtection Onboarding Blob.
    advancedThreatProtectionOnboardingBlob *string
    // Name of the file from which AdvancedThreatProtectionOnboardingBlob was obtained.
    advancedThreatProtectionOnboardingFilename *string
    // Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
    allowSampleSharing *bool
    // Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
    enableExpeditedTelemetryReporting *bool
}
// NewWindowsDefenderAdvancedThreatProtectionConfiguration instantiates a new WindowsDefenderAdvancedThreatProtectionConfiguration and sets the default values.
func NewWindowsDefenderAdvancedThreatProtectionConfiguration()(*WindowsDefenderAdvancedThreatProtectionConfiguration) {
    m := &WindowsDefenderAdvancedThreatProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsDefenderAdvancedThreatProtectionConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsDefenderAdvancedThreatProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDefenderAdvancedThreatProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDefenderAdvancedThreatProtectionConfiguration(), nil
}
// GetAdvancedThreatProtectionAutoPopulateOnboardingBlob gets the advancedThreatProtectionAutoPopulateOnboardingBlob property value. Auto populate onboarding blob programmatically from Advanced Threat protection service
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionAutoPopulateOnboardingBlob()(*bool) {
    return m.advancedThreatProtectionAutoPopulateOnboardingBlob
}
// GetAdvancedThreatProtectionOffboardingBlob gets the advancedThreatProtectionOffboardingBlob property value. Windows Defender AdvancedThreatProtection Offboarding Blob.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOffboardingBlob()(*string) {
    return m.advancedThreatProtectionOffboardingBlob
}
// GetAdvancedThreatProtectionOffboardingFilename gets the advancedThreatProtectionOffboardingFilename property value. Name of the file from which AdvancedThreatProtectionOffboardingBlob was obtained.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOffboardingFilename()(*string) {
    return m.advancedThreatProtectionOffboardingFilename
}
// GetAdvancedThreatProtectionOnboardingBlob gets the advancedThreatProtectionOnboardingBlob property value. Windows Defender AdvancedThreatProtection Onboarding Blob.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOnboardingBlob()(*string) {
    return m.advancedThreatProtectionOnboardingBlob
}
// GetAdvancedThreatProtectionOnboardingFilename gets the advancedThreatProtectionOnboardingFilename property value. Name of the file from which AdvancedThreatProtectionOnboardingBlob was obtained.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOnboardingFilename()(*string) {
    return m.advancedThreatProtectionOnboardingFilename
}
// GetAllowSampleSharing gets the allowSampleSharing property value. Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAllowSampleSharing()(*bool) {
    return m.allowSampleSharing
}
// GetEnableExpeditedTelemetryReporting gets the enableExpeditedTelemetryReporting property value. Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetEnableExpeditedTelemetryReporting()(*bool) {
    return m.enableExpeditedTelemetryReporting
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["advancedThreatProtectionAutoPopulateOnboardingBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionAutoPopulateOnboardingBlob(val)
        }
        return nil
    }
    res["advancedThreatProtectionOffboardingBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOffboardingBlob(val)
        }
        return nil
    }
    res["advancedThreatProtectionOffboardingFilename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOffboardingFilename(val)
        }
        return nil
    }
    res["advancedThreatProtectionOnboardingBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOnboardingBlob(val)
        }
        return nil
    }
    res["advancedThreatProtectionOnboardingFilename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOnboardingFilename(val)
        }
        return nil
    }
    res["allowSampleSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSampleSharing(val)
        }
        return nil
    }
    res["enableExpeditedTelemetryReporting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableExpeditedTelemetryReporting(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("advancedThreatProtectionAutoPopulateOnboardingBlob", m.GetAdvancedThreatProtectionAutoPopulateOnboardingBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOffboardingBlob", m.GetAdvancedThreatProtectionOffboardingBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOffboardingFilename", m.GetAdvancedThreatProtectionOffboardingFilename())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOnboardingBlob", m.GetAdvancedThreatProtectionOnboardingBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOnboardingFilename", m.GetAdvancedThreatProtectionOnboardingFilename())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowSampleSharing", m.GetAllowSampleSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableExpeditedTelemetryReporting", m.GetEnableExpeditedTelemetryReporting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedThreatProtectionAutoPopulateOnboardingBlob sets the advancedThreatProtectionAutoPopulateOnboardingBlob property value. Auto populate onboarding blob programmatically from Advanced Threat protection service
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionAutoPopulateOnboardingBlob(value *bool)() {
    m.advancedThreatProtectionAutoPopulateOnboardingBlob = value
}
// SetAdvancedThreatProtectionOffboardingBlob sets the advancedThreatProtectionOffboardingBlob property value. Windows Defender AdvancedThreatProtection Offboarding Blob.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOffboardingBlob(value *string)() {
    m.advancedThreatProtectionOffboardingBlob = value
}
// SetAdvancedThreatProtectionOffboardingFilename sets the advancedThreatProtectionOffboardingFilename property value. Name of the file from which AdvancedThreatProtectionOffboardingBlob was obtained.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOffboardingFilename(value *string)() {
    m.advancedThreatProtectionOffboardingFilename = value
}
// SetAdvancedThreatProtectionOnboardingBlob sets the advancedThreatProtectionOnboardingBlob property value. Windows Defender AdvancedThreatProtection Onboarding Blob.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOnboardingBlob(value *string)() {
    m.advancedThreatProtectionOnboardingBlob = value
}
// SetAdvancedThreatProtectionOnboardingFilename sets the advancedThreatProtectionOnboardingFilename property value. Name of the file from which AdvancedThreatProtectionOnboardingBlob was obtained.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOnboardingFilename(value *string)() {
    m.advancedThreatProtectionOnboardingFilename = value
}
// SetAllowSampleSharing sets the allowSampleSharing property value. Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAllowSampleSharing(value *bool)() {
    m.allowSampleSharing = value
}
// SetEnableExpeditedTelemetryReporting sets the enableExpeditedTelemetryReporting property value. Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetEnableExpeditedTelemetryReporting(value *bool)() {
    m.enableExpeditedTelemetryReporting = value
}
