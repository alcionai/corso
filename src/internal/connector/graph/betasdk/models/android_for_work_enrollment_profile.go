package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkEnrollmentProfile enrollment Profile used to enroll COSU devices using Google's Cloud Management.
type AndroidForWorkEnrollmentProfile struct {
    Entity
    // Tenant GUID the enrollment profile belongs to.
    accountId *string
    // Date time the enrollment profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description for the enrollment profile.
    description *string
    // Display name for the enrollment profile.
    displayName *string
    // Total number of Android devices that have enrolled using this enrollment profile.
    enrolledDeviceCount *int32
    // Date time the enrollment profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // String used to generate a QR code for the token.
    qrCodeContent *string
    // String used to generate a QR code for the token.
    qrCodeImage MimeContentable
    // Date time the most recently created token will expire.
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Value of the most recently created token for this enrollment profile.
    tokenValue *string
}
// NewAndroidForWorkEnrollmentProfile instantiates a new androidForWorkEnrollmentProfile and sets the default values.
func NewAndroidForWorkEnrollmentProfile()(*AndroidForWorkEnrollmentProfile) {
    m := &AndroidForWorkEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidForWorkEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkEnrollmentProfile(), nil
}
// GetAccountId gets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidForWorkEnrollmentProfile) GetAccountId()(*string) {
    return m.accountId
}
// GetCreatedDateTime gets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidForWorkEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrolledDeviceCount gets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetEnrolledDeviceCount()(*int32) {
    return m.enrolledDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enrolledDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDeviceCount(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["qrCodeContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeContent(val)
        }
        return nil
    }
    res["qrCodeImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeImage(val.(MimeContentable))
        }
        return nil
    }
    res["tokenExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpirationDateTime(val)
        }
        return nil
    }
    res["tokenValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenValue(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidForWorkEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetQrCodeContent gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) GetQrCodeContent()(*string) {
    return m.qrCodeContent
}
// GetQrCodeImage gets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) GetQrCodeImage()(MimeContentable) {
    return m.qrCodeImage
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidForWorkEnrollmentProfile) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.tokenExpirationDateTime
}
// GetTokenValue gets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetTokenValue()(*string) {
    return m.tokenValue
}
// Serialize serializes information the current object
func (m *AndroidForWorkEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enrolledDeviceCount", m.GetEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("qrCodeContent", m.GetQrCodeContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("qrCodeImage", m.GetQrCodeImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("tokenExpirationDateTime", m.GetTokenExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenValue", m.GetTokenValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountId sets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidForWorkEnrollmentProfile) SetAccountId(value *string)() {
    m.accountId = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidForWorkEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrolledDeviceCount sets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetEnrolledDeviceCount(value *int32)() {
    m.enrolledDeviceCount = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidForWorkEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetQrCodeContent sets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) SetQrCodeContent(value *string)() {
    m.qrCodeContent = value
}
// SetQrCodeImage sets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) SetQrCodeImage(value MimeContentable)() {
    m.qrCodeImage = value
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidForWorkEnrollmentProfile) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenExpirationDateTime = value
}
// SetTokenValue sets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetTokenValue(value *string)() {
    m.tokenValue = value
}
