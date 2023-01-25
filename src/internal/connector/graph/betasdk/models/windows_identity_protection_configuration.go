package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsIdentityProtectionConfiguration 
type WindowsIdentityProtectionConfiguration struct {
    DeviceConfiguration
    // Boolean value used to enable enhanced anti-spoofing for facial feature recognition on Windows Hello face authentication.
    enhancedAntiSpoofingForFacialFeaturesEnabled *bool
    // Integer value specifies the period (in days) that a PIN can be used before the system requires the user to change it. Valid values are 0 to 730 inclusive. Valid values 0 to 730
    pinExpirationInDays *int32
    // Possible values of the ConfigurationUsage list.
    pinLowercaseCharactersUsage *ConfigurationUsage
    // Integer value that sets the maximum number of characters allowed for the work PIN. Valid values are 4 to 127 inclusive and greater than or equal to the value set for the minimum PIN. Valid values 4 to 127
    pinMaximumLength *int32
    // Integer value that sets the minimum number of characters required for the Windows Hello for Business PIN. Valid values are 4 to 127 inclusive and less than or equal to the value set for the maximum PIN. Valid values 4 to 127
    pinMinimumLength *int32
    // Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset. Valid values 0 to 50
    pinPreviousBlockCount *int32
    // Boolean value that enables a user to change their PIN by using the Windows Hello for Business PIN recovery service.
    pinRecoveryEnabled *bool
    // Possible values of the ConfigurationUsage list.
    pinSpecialCharactersUsage *ConfigurationUsage
    // Possible values of the ConfigurationUsage list.
    pinUppercaseCharactersUsage *ConfigurationUsage
    // Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
    securityDeviceRequired *bool
    // Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
    unlockWithBiometricsEnabled *bool
    // Boolean value that enables Windows Hello for Business to use certificates to authenticate on-premise resources.
    useCertificatesForOnPremisesAuthEnabled *bool
    // Boolean value used to enable the Windows Hello security key as a logon credential.
    useSecurityKeyForSignin *bool
    // Boolean value that blocks Windows Hello for Business as a method for signing into Windows.
    windowsHelloForBusinessBlocked *bool
}
// NewWindowsIdentityProtectionConfiguration instantiates a new WindowsIdentityProtectionConfiguration and sets the default values.
func NewWindowsIdentityProtectionConfiguration()(*WindowsIdentityProtectionConfiguration) {
    m := &WindowsIdentityProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsIdentityProtectionConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsIdentityProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsIdentityProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsIdentityProtectionConfiguration(), nil
}
// GetEnhancedAntiSpoofingForFacialFeaturesEnabled gets the enhancedAntiSpoofingForFacialFeaturesEnabled property value. Boolean value used to enable enhanced anti-spoofing for facial feature recognition on Windows Hello face authentication.
func (m *WindowsIdentityProtectionConfiguration) GetEnhancedAntiSpoofingForFacialFeaturesEnabled()(*bool) {
    return m.enhancedAntiSpoofingForFacialFeaturesEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsIdentityProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["enhancedAntiSpoofingForFacialFeaturesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnhancedAntiSpoofingForFacialFeaturesEnabled(val)
        }
        return nil
    }
    res["pinExpirationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinExpirationInDays(val)
        }
        return nil
    }
    res["pinLowercaseCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinLowercaseCharactersUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["pinMaximumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinMaximumLength(val)
        }
        return nil
    }
    res["pinMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinMinimumLength(val)
        }
        return nil
    }
    res["pinPreviousBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinPreviousBlockCount(val)
        }
        return nil
    }
    res["pinRecoveryEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRecoveryEnabled(val)
        }
        return nil
    }
    res["pinSpecialCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinSpecialCharactersUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["pinUppercaseCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinUppercaseCharactersUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["securityDeviceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDeviceRequired(val)
        }
        return nil
    }
    res["unlockWithBiometricsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlockWithBiometricsEnabled(val)
        }
        return nil
    }
    res["useCertificatesForOnPremisesAuthEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseCertificatesForOnPremisesAuthEnabled(val)
        }
        return nil
    }
    res["useSecurityKeyForSignin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSecurityKeyForSignin(val)
        }
        return nil
    }
    res["windowsHelloForBusinessBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsHelloForBusinessBlocked(val)
        }
        return nil
    }
    return res
}
// GetPinExpirationInDays gets the pinExpirationInDays property value. Integer value specifies the period (in days) that a PIN can be used before the system requires the user to change it. Valid values are 0 to 730 inclusive. Valid values 0 to 730
func (m *WindowsIdentityProtectionConfiguration) GetPinExpirationInDays()(*int32) {
    return m.pinExpirationInDays
}
// GetPinLowercaseCharactersUsage gets the pinLowercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) GetPinLowercaseCharactersUsage()(*ConfigurationUsage) {
    return m.pinLowercaseCharactersUsage
}
// GetPinMaximumLength gets the pinMaximumLength property value. Integer value that sets the maximum number of characters allowed for the work PIN. Valid values are 4 to 127 inclusive and greater than or equal to the value set for the minimum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) GetPinMaximumLength()(*int32) {
    return m.pinMaximumLength
}
// GetPinMinimumLength gets the pinMinimumLength property value. Integer value that sets the minimum number of characters required for the Windows Hello for Business PIN. Valid values are 4 to 127 inclusive and less than or equal to the value set for the maximum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) GetPinMinimumLength()(*int32) {
    return m.pinMinimumLength
}
// GetPinPreviousBlockCount gets the pinPreviousBlockCount property value. Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset. Valid values 0 to 50
func (m *WindowsIdentityProtectionConfiguration) GetPinPreviousBlockCount()(*int32) {
    return m.pinPreviousBlockCount
}
// GetPinRecoveryEnabled gets the pinRecoveryEnabled property value. Boolean value that enables a user to change their PIN by using the Windows Hello for Business PIN recovery service.
func (m *WindowsIdentityProtectionConfiguration) GetPinRecoveryEnabled()(*bool) {
    return m.pinRecoveryEnabled
}
// GetPinSpecialCharactersUsage gets the pinSpecialCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) GetPinSpecialCharactersUsage()(*ConfigurationUsage) {
    return m.pinSpecialCharactersUsage
}
// GetPinUppercaseCharactersUsage gets the pinUppercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) GetPinUppercaseCharactersUsage()(*ConfigurationUsage) {
    return m.pinUppercaseCharactersUsage
}
// GetSecurityDeviceRequired gets the securityDeviceRequired property value. Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
func (m *WindowsIdentityProtectionConfiguration) GetSecurityDeviceRequired()(*bool) {
    return m.securityDeviceRequired
}
// GetUnlockWithBiometricsEnabled gets the unlockWithBiometricsEnabled property value. Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
func (m *WindowsIdentityProtectionConfiguration) GetUnlockWithBiometricsEnabled()(*bool) {
    return m.unlockWithBiometricsEnabled
}
// GetUseCertificatesForOnPremisesAuthEnabled gets the useCertificatesForOnPremisesAuthEnabled property value. Boolean value that enables Windows Hello for Business to use certificates to authenticate on-premise resources.
func (m *WindowsIdentityProtectionConfiguration) GetUseCertificatesForOnPremisesAuthEnabled()(*bool) {
    return m.useCertificatesForOnPremisesAuthEnabled
}
// GetUseSecurityKeyForSignin gets the useSecurityKeyForSignin property value. Boolean value used to enable the Windows Hello security key as a logon credential.
func (m *WindowsIdentityProtectionConfiguration) GetUseSecurityKeyForSignin()(*bool) {
    return m.useSecurityKeyForSignin
}
// GetWindowsHelloForBusinessBlocked gets the windowsHelloForBusinessBlocked property value. Boolean value that blocks Windows Hello for Business as a method for signing into Windows.
func (m *WindowsIdentityProtectionConfiguration) GetWindowsHelloForBusinessBlocked()(*bool) {
    return m.windowsHelloForBusinessBlocked
}
// Serialize serializes information the current object
func (m *WindowsIdentityProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enhancedAntiSpoofingForFacialFeaturesEnabled", m.GetEnhancedAntiSpoofingForFacialFeaturesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinExpirationInDays", m.GetPinExpirationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetPinLowercaseCharactersUsage() != nil {
        cast := (*m.GetPinLowercaseCharactersUsage()).String()
        err = writer.WriteStringValue("pinLowercaseCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinMaximumLength", m.GetPinMaximumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinMinimumLength", m.GetPinMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinPreviousBlockCount", m.GetPinPreviousBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pinRecoveryEnabled", m.GetPinRecoveryEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPinSpecialCharactersUsage() != nil {
        cast := (*m.GetPinSpecialCharactersUsage()).String()
        err = writer.WriteStringValue("pinSpecialCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPinUppercaseCharactersUsage() != nil {
        cast := (*m.GetPinUppercaseCharactersUsage()).String()
        err = writer.WriteStringValue("pinUppercaseCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDeviceRequired", m.GetSecurityDeviceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("unlockWithBiometricsEnabled", m.GetUnlockWithBiometricsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useCertificatesForOnPremisesAuthEnabled", m.GetUseCertificatesForOnPremisesAuthEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useSecurityKeyForSignin", m.GetUseSecurityKeyForSignin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsHelloForBusinessBlocked", m.GetWindowsHelloForBusinessBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnhancedAntiSpoofingForFacialFeaturesEnabled sets the enhancedAntiSpoofingForFacialFeaturesEnabled property value. Boolean value used to enable enhanced anti-spoofing for facial feature recognition on Windows Hello face authentication.
func (m *WindowsIdentityProtectionConfiguration) SetEnhancedAntiSpoofingForFacialFeaturesEnabled(value *bool)() {
    m.enhancedAntiSpoofingForFacialFeaturesEnabled = value
}
// SetPinExpirationInDays sets the pinExpirationInDays property value. Integer value specifies the period (in days) that a PIN can be used before the system requires the user to change it. Valid values are 0 to 730 inclusive. Valid values 0 to 730
func (m *WindowsIdentityProtectionConfiguration) SetPinExpirationInDays(value *int32)() {
    m.pinExpirationInDays = value
}
// SetPinLowercaseCharactersUsage sets the pinLowercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) SetPinLowercaseCharactersUsage(value *ConfigurationUsage)() {
    m.pinLowercaseCharactersUsage = value
}
// SetPinMaximumLength sets the pinMaximumLength property value. Integer value that sets the maximum number of characters allowed for the work PIN. Valid values are 4 to 127 inclusive and greater than or equal to the value set for the minimum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) SetPinMaximumLength(value *int32)() {
    m.pinMaximumLength = value
}
// SetPinMinimumLength sets the pinMinimumLength property value. Integer value that sets the minimum number of characters required for the Windows Hello for Business PIN. Valid values are 4 to 127 inclusive and less than or equal to the value set for the maximum PIN. Valid values 4 to 127
func (m *WindowsIdentityProtectionConfiguration) SetPinMinimumLength(value *int32)() {
    m.pinMinimumLength = value
}
// SetPinPreviousBlockCount sets the pinPreviousBlockCount property value. Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset. Valid values 0 to 50
func (m *WindowsIdentityProtectionConfiguration) SetPinPreviousBlockCount(value *int32)() {
    m.pinPreviousBlockCount = value
}
// SetPinRecoveryEnabled sets the pinRecoveryEnabled property value. Boolean value that enables a user to change their PIN by using the Windows Hello for Business PIN recovery service.
func (m *WindowsIdentityProtectionConfiguration) SetPinRecoveryEnabled(value *bool)() {
    m.pinRecoveryEnabled = value
}
// SetPinSpecialCharactersUsage sets the pinSpecialCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) SetPinSpecialCharactersUsage(value *ConfigurationUsage)() {
    m.pinSpecialCharactersUsage = value
}
// SetPinUppercaseCharactersUsage sets the pinUppercaseCharactersUsage property value. Possible values of the ConfigurationUsage list.
func (m *WindowsIdentityProtectionConfiguration) SetPinUppercaseCharactersUsage(value *ConfigurationUsage)() {
    m.pinUppercaseCharactersUsage = value
}
// SetSecurityDeviceRequired sets the securityDeviceRequired property value. Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
func (m *WindowsIdentityProtectionConfiguration) SetSecurityDeviceRequired(value *bool)() {
    m.securityDeviceRequired = value
}
// SetUnlockWithBiometricsEnabled sets the unlockWithBiometricsEnabled property value. Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
func (m *WindowsIdentityProtectionConfiguration) SetUnlockWithBiometricsEnabled(value *bool)() {
    m.unlockWithBiometricsEnabled = value
}
// SetUseCertificatesForOnPremisesAuthEnabled sets the useCertificatesForOnPremisesAuthEnabled property value. Boolean value that enables Windows Hello for Business to use certificates to authenticate on-premise resources.
func (m *WindowsIdentityProtectionConfiguration) SetUseCertificatesForOnPremisesAuthEnabled(value *bool)() {
    m.useCertificatesForOnPremisesAuthEnabled = value
}
// SetUseSecurityKeyForSignin sets the useSecurityKeyForSignin property value. Boolean value used to enable the Windows Hello security key as a logon credential.
func (m *WindowsIdentityProtectionConfiguration) SetUseSecurityKeyForSignin(value *bool)() {
    m.useSecurityKeyForSignin = value
}
// SetWindowsHelloForBusinessBlocked sets the windowsHelloForBusinessBlocked property value. Boolean value that blocks Windows Hello for Business as a method for signing into Windows.
func (m *WindowsIdentityProtectionConfiguration) SetWindowsHelloForBusinessBlocked(value *bool)() {
    m.windowsHelloForBusinessBlocked = value
}
