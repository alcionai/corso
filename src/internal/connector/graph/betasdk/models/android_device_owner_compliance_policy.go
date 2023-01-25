package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerCompliancePolicy 
type AndroidDeviceOwnerCompliancePolicy struct {
    DeviceCompliancePolicy
    // MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
    advancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Require that devices have enabled device threat protection.
    deviceThreatProtectionEnabled *bool
    // Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
    deviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Minimum Android security patch level.
    minAndroidSecurityPatchLevel *string
    // Maximum Android version.
    osMaximumVersion *string
    // Minimum Android version.
    osMinimumVersion *string
    // Number of days before the password expires. Valid values 1 to 365
    passwordExpirationDays *int32
    // Minimum password length. Valid values 4 to 16
    passwordMinimumLength *int32
    // Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
    passwordMinimumLetterCharacters *int32
    // Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
    passwordMinimumLowerCaseCharacters *int32
    // Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
    passwordMinimumNonLetterCharacters *int32
    // Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
    passwordMinimumNumericCharacters *int32
    // Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
    passwordMinimumSymbolCharacters *int32
    // Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
    passwordMinimumUpperCaseCharacters *int32
    // Minutes of inactivity before a password is required.
    passwordMinutesOfInactivityBeforeLock *int32
    // Number of previous passwords to block. Valid values 1 to 24
    passwordPreviousPasswordCountToBlock *int32
    // Require a password to unlock device.
    passwordRequired *bool
    // Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
    passwordRequiredType *AndroidDeviceOwnerRequiredPasswordType
    // If setting is set to true, checks that the Intune app installed on fully managed, dedicated, or corporate-owned work profile Android Enterprise enrolled devices, is the one provided by Microsoft from the Managed Google Playstore. If the check fails, the device will be reported as non-compliant.
    securityRequireIntuneAppIntegrity *bool
    // Require the device to pass the SafetyNet basic integrity check.
    securityRequireSafetyNetAttestationBasicIntegrity *bool
    // Require the device to pass the SafetyNet certified device check.
    securityRequireSafetyNetAttestationCertifiedDevice *bool
    // Require encryption on Android devices.
    storageRequireEncryption *bool
}
// NewAndroidDeviceOwnerCompliancePolicy instantiates a new AndroidDeviceOwnerCompliancePolicy and sets the default values.
func NewAndroidDeviceOwnerCompliancePolicy()(*AndroidDeviceOwnerCompliancePolicy) {
    m := &AndroidDeviceOwnerCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerCompliancePolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerCompliancePolicy(), nil
}
// GetAdvancedThreatProtectionRequiredSecurityLevel gets the advancedThreatProtectionRequiredSecurityLevel property value. MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) GetAdvancedThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    return m.advancedThreatProtectionRequiredSecurityLevel
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidDeviceOwnerCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    return m.deviceThreatProtectionEnabled
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    return m.deviceThreatProtectionRequiredSecurityLevel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["advancedThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["deviceThreatProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionEnabled(val)
        }
        return nil
    }
    res["deviceThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["minAndroidSecurityPatchLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["osMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumVersion(val)
        }
        return nil
    }
    res["osMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumVersion(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinimumLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLowerCaseCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNonLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNumericCharacters(val)
        }
        return nil
    }
    res["passwordMinimumSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumSymbolCharacters(val)
        }
        return nil
    }
    res["passwordMinimumUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumUpperCaseCharacters(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passwordPreviousPasswordCountToBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordCountToBlock(val)
        }
        return nil
    }
    res["passwordRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequired(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidDeviceOwnerRequiredPasswordType))
        }
        return nil
    }
    res["securityRequireIntuneAppIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireIntuneAppIntegrity(val)
        }
        return nil
    }
    res["securityRequireSafetyNetAttestationBasicIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireSafetyNetAttestationBasicIntegrity(val)
        }
        return nil
    }
    res["securityRequireSafetyNetAttestationCertifiedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireSafetyNetAttestationCertifiedDevice(val)
        }
        return nil
    }
    res["storageRequireEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRequireEncryption(val)
        }
        return nil
    }
    return res
}
// GetMinAndroidSecurityPatchLevel gets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidDeviceOwnerCompliancePolicy) GetMinAndroidSecurityPatchLevel()(*string) {
    return m.minAndroidSecurityPatchLevel
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidDeviceOwnerCompliancePolicy) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidDeviceOwnerCompliancePolicy) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinimumLetterCharacters gets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumLetterCharacters()(*int32) {
    return m.passwordMinimumLetterCharacters
}
// GetPasswordMinimumLowerCaseCharacters gets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumLowerCaseCharacters()(*int32) {
    return m.passwordMinimumLowerCaseCharacters
}
// GetPasswordMinimumNonLetterCharacters gets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumNonLetterCharacters()(*int32) {
    return m.passwordMinimumNonLetterCharacters
}
// GetPasswordMinimumNumericCharacters gets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumNumericCharacters()(*int32) {
    return m.passwordMinimumNumericCharacters
}
// GetPasswordMinimumSymbolCharacters gets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumSymbolCharacters()(*int32) {
    return m.passwordMinimumSymbolCharacters
}
// GetPasswordMinimumUpperCaseCharacters gets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinimumUpperCaseCharacters()(*int32) {
    return m.passwordMinimumUpperCaseCharacters
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    return m.passwordMinutesOfInactivityBeforeLock
}
// GetPasswordPreviousPasswordCountToBlock gets the passwordPreviousPasswordCountToBlock property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordPreviousPasswordCountToBlock()(*int32) {
    return m.passwordPreviousPasswordCountToBlock
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerCompliancePolicy) GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetSecurityRequireIntuneAppIntegrity gets the securityRequireIntuneAppIntegrity property value. If setting is set to true, checks that the Intune app installed on fully managed, dedicated, or corporate-owned work profile Android Enterprise enrolled devices, is the one provided by Microsoft from the Managed Google Playstore. If the check fails, the device will be reported as non-compliant.
func (m *AndroidDeviceOwnerCompliancePolicy) GetSecurityRequireIntuneAppIntegrity()(*bool) {
    return m.securityRequireIntuneAppIntegrity
}
// GetSecurityRequireSafetyNetAttestationBasicIntegrity gets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidDeviceOwnerCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool) {
    return m.securityRequireSafetyNetAttestationBasicIntegrity
}
// GetSecurityRequireSafetyNetAttestationCertifiedDevice gets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidDeviceOwnerCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool) {
    return m.securityRequireSafetyNetAttestationCertifiedDevice
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidDeviceOwnerCompliancePolicy) GetStorageRequireEncryption()(*bool) {
    return m.storageRequireEncryption
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetAdvancedThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("advancedThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceThreatProtectionEnabled", m.GetDeviceThreatProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetDeviceThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("deviceThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minAndroidSecurityPatchLevel", m.GetMinAndroidSecurityPatchLevel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMaximumVersion", m.GetOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumVersion", m.GetOsMinimumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLetterCharacters", m.GetPasswordMinimumLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLowerCaseCharacters", m.GetPasswordMinimumLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNonLetterCharacters", m.GetPasswordMinimumNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNumericCharacters", m.GetPasswordMinimumNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumSymbolCharacters", m.GetPasswordMinimumSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumUpperCaseCharacters", m.GetPasswordMinimumUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeLock", m.GetPasswordMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordCountToBlock", m.GetPasswordPreviousPasswordCountToBlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordRequired", m.GetPasswordRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireIntuneAppIntegrity", m.GetSecurityRequireIntuneAppIntegrity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireSafetyNetAttestationBasicIntegrity", m.GetSecurityRequireSafetyNetAttestationBasicIntegrity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireSafetyNetAttestationCertifiedDevice", m.GetSecurityRequireSafetyNetAttestationCertifiedDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRequireEncryption", m.GetStorageRequireEncryption())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedThreatProtectionRequiredSecurityLevel sets the advancedThreatProtectionRequiredSecurityLevel property value. MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) SetAdvancedThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    m.advancedThreatProtectionRequiredSecurityLevel = value
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidDeviceOwnerCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    m.deviceThreatProtectionEnabled = value
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *AndroidDeviceOwnerCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    m.deviceThreatProtectionRequiredSecurityLevel = value
}
// SetMinAndroidSecurityPatchLevel sets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidDeviceOwnerCompliancePolicy) SetMinAndroidSecurityPatchLevel(value *string)() {
    m.minAndroidSecurityPatchLevel = value
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidDeviceOwnerCompliancePolicy) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidDeviceOwnerCompliancePolicy) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinimumLetterCharacters sets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumLetterCharacters(value *int32)() {
    m.passwordMinimumLetterCharacters = value
}
// SetPasswordMinimumLowerCaseCharacters sets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumLowerCaseCharacters(value *int32)() {
    m.passwordMinimumLowerCaseCharacters = value
}
// SetPasswordMinimumNonLetterCharacters sets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumNonLetterCharacters(value *int32)() {
    m.passwordMinimumNonLetterCharacters = value
}
// SetPasswordMinimumNumericCharacters sets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumNumericCharacters(value *int32)() {
    m.passwordMinimumNumericCharacters = value
}
// SetPasswordMinimumSymbolCharacters sets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumSymbolCharacters(value *int32)() {
    m.passwordMinimumSymbolCharacters = value
}
// SetPasswordMinimumUpperCaseCharacters sets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinimumUpperCaseCharacters(value *int32)() {
    m.passwordMinimumUpperCaseCharacters = value
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    m.passwordMinutesOfInactivityBeforeLock = value
}
// SetPasswordPreviousPasswordCountToBlock sets the passwordPreviousPasswordCountToBlock property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordPreviousPasswordCountToBlock(value *int32)() {
    m.passwordPreviousPasswordCountToBlock = value
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerCompliancePolicy) SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetSecurityRequireIntuneAppIntegrity sets the securityRequireIntuneAppIntegrity property value. If setting is set to true, checks that the Intune app installed on fully managed, dedicated, or corporate-owned work profile Android Enterprise enrolled devices, is the one provided by Microsoft from the Managed Google Playstore. If the check fails, the device will be reported as non-compliant.
func (m *AndroidDeviceOwnerCompliancePolicy) SetSecurityRequireIntuneAppIntegrity(value *bool)() {
    m.securityRequireIntuneAppIntegrity = value
}
// SetSecurityRequireSafetyNetAttestationBasicIntegrity sets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidDeviceOwnerCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)() {
    m.securityRequireSafetyNetAttestationBasicIntegrity = value
}
// SetSecurityRequireSafetyNetAttestationCertifiedDevice sets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidDeviceOwnerCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)() {
    m.securityRequireSafetyNetAttestationCertifiedDevice = value
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidDeviceOwnerCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    m.storageRequireEncryption = value
}
