package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceCertificateState provides operations to call the add method.
type ManagedDeviceCertificateState struct {
    Entity
    // Extended key usage
    certificateEnhancedKeyUsage *string
    // Error code
    certificateErrorCode *int32
    // Certificate expiry date
    certificateExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Issuance date
    certificateIssuanceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Certificate Issuance State Options.
    certificateIssuanceState *CertificateIssuanceStates
    // Issuer
    certificateIssuer *string
    // Key length
    certificateKeyLength *int32
    // Key Storage Provider (KSP) Import Options.
    certificateKeyStorageProvider *KeyStorageProviderOption
    // Key Usage Options.
    certificateKeyUsage *KeyUsages
    // Last certificate issuance state change
    certificateLastIssuanceStateChangedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Certificate profile display name
    certificateProfileDisplayName *string
    // Certificate Revocation Status.
    certificateRevokeStatus *CertificateRevocationStatus
    // Serial number
    certificateSerialNumber *string
    // Subject Alternative Name Options.
    certificateSubjectAlternativeNameFormat *SubjectAlternativeNameType
    // Subject alternative name format string for custom formats
    certificateSubjectAlternativeNameFormatString *string
    // Subject Name Format Options.
    certificateSubjectNameFormat *SubjectNameFormat
    // Subject name format string for custom subject name formats
    certificateSubjectNameFormatString *string
    // Thumbprint
    certificateThumbprint *string
    // Validity period
    certificateValidityPeriod *int32
    // Certificate Validity Period Options.
    certificateValidityPeriodUnits *CertificateValidityPeriodScale
    // Device display name
    deviceDisplayName *string
    // Supported platform types.
    devicePlatform *DevicePlatformType
    // Last certificate issuance state change
    lastCertificateStateChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // User display name
    userDisplayName *string
}
// NewManagedDeviceCertificateState instantiates a new managedDeviceCertificateState and sets the default values.
func NewManagedDeviceCertificateState()(*ManagedDeviceCertificateState) {
    m := &ManagedDeviceCertificateState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceCertificateStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceCertificateStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceCertificateState(), nil
}
// GetCertificateEnhancedKeyUsage gets the certificateEnhancedKeyUsage property value. Extended key usage
func (m *ManagedDeviceCertificateState) GetCertificateEnhancedKeyUsage()(*string) {
    return m.certificateEnhancedKeyUsage
}
// GetCertificateErrorCode gets the certificateErrorCode property value. Error code
func (m *ManagedDeviceCertificateState) GetCertificateErrorCode()(*int32) {
    return m.certificateErrorCode
}
// GetCertificateExpirationDateTime gets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedDeviceCertificateState) GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateExpirationDateTime
}
// GetCertificateIssuanceDateTime gets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedDeviceCertificateState) GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateIssuanceDateTime
}
// GetCertificateIssuanceState gets the certificateIssuanceState property value. Certificate Issuance State Options.
func (m *ManagedDeviceCertificateState) GetCertificateIssuanceState()(*CertificateIssuanceStates) {
    return m.certificateIssuanceState
}
// GetCertificateIssuer gets the certificateIssuer property value. Issuer
func (m *ManagedDeviceCertificateState) GetCertificateIssuer()(*string) {
    return m.certificateIssuer
}
// GetCertificateKeyLength gets the certificateKeyLength property value. Key length
func (m *ManagedDeviceCertificateState) GetCertificateKeyLength()(*int32) {
    return m.certificateKeyLength
}
// GetCertificateKeyStorageProvider gets the certificateKeyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *ManagedDeviceCertificateState) GetCertificateKeyStorageProvider()(*KeyStorageProviderOption) {
    return m.certificateKeyStorageProvider
}
// GetCertificateKeyUsage gets the certificateKeyUsage property value. Key Usage Options.
func (m *ManagedDeviceCertificateState) GetCertificateKeyUsage()(*KeyUsages) {
    return m.certificateKeyUsage
}
// GetCertificateLastIssuanceStateChangedDateTime gets the certificateLastIssuanceStateChangedDateTime property value. Last certificate issuance state change
func (m *ManagedDeviceCertificateState) GetCertificateLastIssuanceStateChangedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateLastIssuanceStateChangedDateTime
}
// GetCertificateProfileDisplayName gets the certificateProfileDisplayName property value. Certificate profile display name
func (m *ManagedDeviceCertificateState) GetCertificateProfileDisplayName()(*string) {
    return m.certificateProfileDisplayName
}
// GetCertificateRevokeStatus gets the certificateRevokeStatus property value. Certificate Revocation Status.
func (m *ManagedDeviceCertificateState) GetCertificateRevokeStatus()(*CertificateRevocationStatus) {
    return m.certificateRevokeStatus
}
// GetCertificateSerialNumber gets the certificateSerialNumber property value. Serial number
func (m *ManagedDeviceCertificateState) GetCertificateSerialNumber()(*string) {
    return m.certificateSerialNumber
}
// GetCertificateSubjectAlternativeNameFormat gets the certificateSubjectAlternativeNameFormat property value. Subject Alternative Name Options.
func (m *ManagedDeviceCertificateState) GetCertificateSubjectAlternativeNameFormat()(*SubjectAlternativeNameType) {
    return m.certificateSubjectAlternativeNameFormat
}
// GetCertificateSubjectAlternativeNameFormatString gets the certificateSubjectAlternativeNameFormatString property value. Subject alternative name format string for custom formats
func (m *ManagedDeviceCertificateState) GetCertificateSubjectAlternativeNameFormatString()(*string) {
    return m.certificateSubjectAlternativeNameFormatString
}
// GetCertificateSubjectNameFormat gets the certificateSubjectNameFormat property value. Subject Name Format Options.
func (m *ManagedDeviceCertificateState) GetCertificateSubjectNameFormat()(*SubjectNameFormat) {
    return m.certificateSubjectNameFormat
}
// GetCertificateSubjectNameFormatString gets the certificateSubjectNameFormatString property value. Subject name format string for custom subject name formats
func (m *ManagedDeviceCertificateState) GetCertificateSubjectNameFormatString()(*string) {
    return m.certificateSubjectNameFormatString
}
// GetCertificateThumbprint gets the certificateThumbprint property value. Thumbprint
func (m *ManagedDeviceCertificateState) GetCertificateThumbprint()(*string) {
    return m.certificateThumbprint
}
// GetCertificateValidityPeriod gets the certificateValidityPeriod property value. Validity period
func (m *ManagedDeviceCertificateState) GetCertificateValidityPeriod()(*int32) {
    return m.certificateValidityPeriod
}
// GetCertificateValidityPeriodUnits gets the certificateValidityPeriodUnits property value. Certificate Validity Period Options.
func (m *ManagedDeviceCertificateState) GetCertificateValidityPeriodUnits()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodUnits
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Device display name
func (m *ManagedDeviceCertificateState) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetDevicePlatform gets the devicePlatform property value. Supported platform types.
func (m *ManagedDeviceCertificateState) GetDevicePlatform()(*DevicePlatformType) {
    return m.devicePlatform
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceCertificateState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateEnhancedKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateEnhancedKeyUsage(val)
        }
        return nil
    }
    res["certificateErrorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateErrorCode(val)
        }
        return nil
    }
    res["certificateExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateExpirationDateTime(val)
        }
        return nil
    }
    res["certificateIssuanceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuanceDateTime(val)
        }
        return nil
    }
    res["certificateIssuanceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateIssuanceStates)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuanceState(val.(*CertificateIssuanceStates))
        }
        return nil
    }
    res["certificateIssuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuer(val)
        }
        return nil
    }
    res["certificateKeyLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyLength(val)
        }
        return nil
    }
    res["certificateKeyStorageProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyStorageProviderOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyStorageProvider(val.(*KeyStorageProviderOption))
        }
        return nil
    }
    res["certificateKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyUsages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyUsage(val.(*KeyUsages))
        }
        return nil
    }
    res["certificateLastIssuanceStateChangedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateLastIssuanceStateChangedDateTime(val)
        }
        return nil
    }
    res["certificateProfileDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateProfileDisplayName(val)
        }
        return nil
    }
    res["certificateRevokeStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateRevocationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevokeStatus(val.(*CertificateRevocationStatus))
        }
        return nil
    }
    res["certificateSerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSerialNumber(val)
        }
        return nil
    }
    res["certificateSubjectAlternativeNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectAlternativeNameType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectAlternativeNameFormat(val.(*SubjectAlternativeNameType))
        }
        return nil
    }
    res["certificateSubjectAlternativeNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectAlternativeNameFormatString(val)
        }
        return nil
    }
    res["certificateSubjectNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectNameFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectNameFormat(val.(*SubjectNameFormat))
        }
        return nil
    }
    res["certificateSubjectNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectNameFormatString(val)
        }
        return nil
    }
    res["certificateThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateThumbprint(val)
        }
        return nil
    }
    res["certificateValidityPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriod(val)
        }
        return nil
    }
    res["certificateValidityPeriodUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateValidityPeriodScale)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodUnits(val.(*CertificateValidityPeriodScale))
        }
        return nil
    }
    res["deviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
    res["devicePlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["lastCertificateStateChangeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCertificateStateChangeDateTime(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    return res
}
// GetLastCertificateStateChangeDateTime gets the lastCertificateStateChangeDateTime property value. Last certificate issuance state change
func (m *ManagedDeviceCertificateState) GetLastCertificateStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCertificateStateChangeDateTime
}
// GetUserDisplayName gets the userDisplayName property value. User display name
func (m *ManagedDeviceCertificateState) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// Serialize serializes information the current object
func (m *ManagedDeviceCertificateState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificateEnhancedKeyUsage", m.GetCertificateEnhancedKeyUsage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateErrorCode", m.GetCertificateErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateExpirationDateTime", m.GetCertificateExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateIssuanceDateTime", m.GetCertificateIssuanceDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateIssuanceState() != nil {
        cast := (*m.GetCertificateIssuanceState()).String()
        err = writer.WriteStringValue("certificateIssuanceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateIssuer", m.GetCertificateIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateKeyLength", m.GetCertificateKeyLength())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateKeyStorageProvider() != nil {
        cast := (*m.GetCertificateKeyStorageProvider()).String()
        err = writer.WriteStringValue("certificateKeyStorageProvider", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertificateKeyUsage() != nil {
        cast := (*m.GetCertificateKeyUsage()).String()
        err = writer.WriteStringValue("certificateKeyUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateLastIssuanceStateChangedDateTime", m.GetCertificateLastIssuanceStateChangedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateProfileDisplayName", m.GetCertificateProfileDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateRevokeStatus() != nil {
        cast := (*m.GetCertificateRevokeStatus()).String()
        err = writer.WriteStringValue("certificateRevokeStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSerialNumber", m.GetCertificateSerialNumber())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateSubjectAlternativeNameFormat() != nil {
        cast := (*m.GetCertificateSubjectAlternativeNameFormat()).String()
        err = writer.WriteStringValue("certificateSubjectAlternativeNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSubjectAlternativeNameFormatString", m.GetCertificateSubjectAlternativeNameFormatString())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateSubjectNameFormat() != nil {
        cast := (*m.GetCertificateSubjectNameFormat()).String()
        err = writer.WriteStringValue("certificateSubjectNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSubjectNameFormatString", m.GetCertificateSubjectNameFormatString())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateThumbprint", m.GetCertificateThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateValidityPeriod", m.GetCertificateValidityPeriod())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateValidityPeriodUnits() != nil {
        cast := (*m.GetCertificateValidityPeriodUnits()).String()
        err = writer.WriteStringValue("certificateValidityPeriodUnits", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDevicePlatform() != nil {
        cast := (*m.GetDevicePlatform()).String()
        err = writer.WriteStringValue("devicePlatform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCertificateStateChangeDateTime", m.GetLastCertificateStateChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateEnhancedKeyUsage sets the certificateEnhancedKeyUsage property value. Extended key usage
func (m *ManagedDeviceCertificateState) SetCertificateEnhancedKeyUsage(value *string)() {
    m.certificateEnhancedKeyUsage = value
}
// SetCertificateErrorCode sets the certificateErrorCode property value. Error code
func (m *ManagedDeviceCertificateState) SetCertificateErrorCode(value *int32)() {
    m.certificateErrorCode = value
}
// SetCertificateExpirationDateTime sets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedDeviceCertificateState) SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateExpirationDateTime = value
}
// SetCertificateIssuanceDateTime sets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedDeviceCertificateState) SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateIssuanceDateTime = value
}
// SetCertificateIssuanceState sets the certificateIssuanceState property value. Certificate Issuance State Options.
func (m *ManagedDeviceCertificateState) SetCertificateIssuanceState(value *CertificateIssuanceStates)() {
    m.certificateIssuanceState = value
}
// SetCertificateIssuer sets the certificateIssuer property value. Issuer
func (m *ManagedDeviceCertificateState) SetCertificateIssuer(value *string)() {
    m.certificateIssuer = value
}
// SetCertificateKeyLength sets the certificateKeyLength property value. Key length
func (m *ManagedDeviceCertificateState) SetCertificateKeyLength(value *int32)() {
    m.certificateKeyLength = value
}
// SetCertificateKeyStorageProvider sets the certificateKeyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *ManagedDeviceCertificateState) SetCertificateKeyStorageProvider(value *KeyStorageProviderOption)() {
    m.certificateKeyStorageProvider = value
}
// SetCertificateKeyUsage sets the certificateKeyUsage property value. Key Usage Options.
func (m *ManagedDeviceCertificateState) SetCertificateKeyUsage(value *KeyUsages)() {
    m.certificateKeyUsage = value
}
// SetCertificateLastIssuanceStateChangedDateTime sets the certificateLastIssuanceStateChangedDateTime property value. Last certificate issuance state change
func (m *ManagedDeviceCertificateState) SetCertificateLastIssuanceStateChangedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateLastIssuanceStateChangedDateTime = value
}
// SetCertificateProfileDisplayName sets the certificateProfileDisplayName property value. Certificate profile display name
func (m *ManagedDeviceCertificateState) SetCertificateProfileDisplayName(value *string)() {
    m.certificateProfileDisplayName = value
}
// SetCertificateRevokeStatus sets the certificateRevokeStatus property value. Certificate Revocation Status.
func (m *ManagedDeviceCertificateState) SetCertificateRevokeStatus(value *CertificateRevocationStatus)() {
    m.certificateRevokeStatus = value
}
// SetCertificateSerialNumber sets the certificateSerialNumber property value. Serial number
func (m *ManagedDeviceCertificateState) SetCertificateSerialNumber(value *string)() {
    m.certificateSerialNumber = value
}
// SetCertificateSubjectAlternativeNameFormat sets the certificateSubjectAlternativeNameFormat property value. Subject Alternative Name Options.
func (m *ManagedDeviceCertificateState) SetCertificateSubjectAlternativeNameFormat(value *SubjectAlternativeNameType)() {
    m.certificateSubjectAlternativeNameFormat = value
}
// SetCertificateSubjectAlternativeNameFormatString sets the certificateSubjectAlternativeNameFormatString property value. Subject alternative name format string for custom formats
func (m *ManagedDeviceCertificateState) SetCertificateSubjectAlternativeNameFormatString(value *string)() {
    m.certificateSubjectAlternativeNameFormatString = value
}
// SetCertificateSubjectNameFormat sets the certificateSubjectNameFormat property value. Subject Name Format Options.
func (m *ManagedDeviceCertificateState) SetCertificateSubjectNameFormat(value *SubjectNameFormat)() {
    m.certificateSubjectNameFormat = value
}
// SetCertificateSubjectNameFormatString sets the certificateSubjectNameFormatString property value. Subject name format string for custom subject name formats
func (m *ManagedDeviceCertificateState) SetCertificateSubjectNameFormatString(value *string)() {
    m.certificateSubjectNameFormatString = value
}
// SetCertificateThumbprint sets the certificateThumbprint property value. Thumbprint
func (m *ManagedDeviceCertificateState) SetCertificateThumbprint(value *string)() {
    m.certificateThumbprint = value
}
// SetCertificateValidityPeriod sets the certificateValidityPeriod property value. Validity period
func (m *ManagedDeviceCertificateState) SetCertificateValidityPeriod(value *int32)() {
    m.certificateValidityPeriod = value
}
// SetCertificateValidityPeriodUnits sets the certificateValidityPeriodUnits property value. Certificate Validity Period Options.
func (m *ManagedDeviceCertificateState) SetCertificateValidityPeriodUnits(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodUnits = value
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Device display name
func (m *ManagedDeviceCertificateState) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetDevicePlatform sets the devicePlatform property value. Supported platform types.
func (m *ManagedDeviceCertificateState) SetDevicePlatform(value *DevicePlatformType)() {
    m.devicePlatform = value
}
// SetLastCertificateStateChangeDateTime sets the lastCertificateStateChangeDateTime property value. Last certificate issuance state change
func (m *ManagedDeviceCertificateState) SetLastCertificateStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCertificateStateChangeDateTime = value
}
// SetUserDisplayName sets the userDisplayName property value. User display name
func (m *ManagedDeviceCertificateState) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
