package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerCompliancePolicy 
type AospDeviceOwnerCompliancePolicy struct {
    DeviceCompliancePolicy
    // Minimum Android security patch level.
    minAndroidSecurityPatchLevel *string
    // Maximum Android version.
    osMaximumVersion *string
    // Minimum Android version.
    osMinimumVersion *string
    // Minimum password length. Valid values 4 to 16
    passwordMinimumLength *int32
    // Minutes of inactivity before a password is required. Valid values 1 to 8640
    passwordMinutesOfInactivityBeforeLock *int32
    // Require a password to unlock device.
    passwordRequired *bool
    // Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
    passwordRequiredType *AndroidDeviceOwnerRequiredPasswordType
    // Devices must not be jailbroken or rooted.
    securityBlockJailbrokenDevices *bool
    // Require encryption on Android devices.
    storageRequireEncryption *bool
}
// NewAospDeviceOwnerCompliancePolicy instantiates a new AospDeviceOwnerCompliancePolicy and sets the default values.
func NewAospDeviceOwnerCompliancePolicy()(*AospDeviceOwnerCompliancePolicy) {
    m := &AospDeviceOwnerCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.aospDeviceOwnerCompliancePolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAospDeviceOwnerCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAospDeviceOwnerCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAospDeviceOwnerCompliancePolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AospDeviceOwnerCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
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
    res["securityBlockJailbrokenDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityBlockJailbrokenDevices(val)
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
func (m *AospDeviceOwnerCompliancePolicy) GetMinAndroidSecurityPatchLevel()(*string) {
    return m.minAndroidSecurityPatchLevel
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Android version.
func (m *AospDeviceOwnerCompliancePolicy) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Android version.
func (m *AospDeviceOwnerCompliancePolicy) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AospDeviceOwnerCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required. Valid values 1 to 8640
func (m *AospDeviceOwnerCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    return m.passwordMinutesOfInactivityBeforeLock
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock device.
func (m *AospDeviceOwnerCompliancePolicy) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AospDeviceOwnerCompliancePolicy) GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetSecurityBlockJailbrokenDevices gets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AospDeviceOwnerCompliancePolicy) GetSecurityBlockJailbrokenDevices()(*bool) {
    return m.securityBlockJailbrokenDevices
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AospDeviceOwnerCompliancePolicy) GetStorageRequireEncryption()(*bool) {
    return m.storageRequireEncryption
}
// Serialize serializes information the current object
func (m *AospDeviceOwnerCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
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
        err = writer.WriteBoolValue("securityBlockJailbrokenDevices", m.GetSecurityBlockJailbrokenDevices())
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
// SetMinAndroidSecurityPatchLevel sets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AospDeviceOwnerCompliancePolicy) SetMinAndroidSecurityPatchLevel(value *string)() {
    m.minAndroidSecurityPatchLevel = value
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Android version.
func (m *AospDeviceOwnerCompliancePolicy) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Android version.
func (m *AospDeviceOwnerCompliancePolicy) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AospDeviceOwnerCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required. Valid values 1 to 8640
func (m *AospDeviceOwnerCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    m.passwordMinutesOfInactivityBeforeLock = value
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock device.
func (m *AospDeviceOwnerCompliancePolicy) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AospDeviceOwnerCompliancePolicy) SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetSecurityBlockJailbrokenDevices sets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AospDeviceOwnerCompliancePolicy) SetSecurityBlockJailbrokenDevices(value *bool)() {
    m.securityBlockJailbrokenDevices = value
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AospDeviceOwnerCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    m.storageRequireEncryption = value
}
