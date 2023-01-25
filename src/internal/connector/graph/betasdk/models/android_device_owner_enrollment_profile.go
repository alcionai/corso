package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerEnrollmentProfile enrollment Profile used to enroll Android Enterprise devices using Google's Cloud Management.
type AndroidDeviceOwnerEnrollmentProfile struct {
    Entity
    // Tenant GUID the enrollment profile belongs to.
    accountId *string
    // Boolean that indicates that the Wi-Fi network should be configured during device provisioning. When set to TRUE, device provisioning will use Wi-Fi related properties to automatically connect to Wi-Fi networks. When set to FALSE or undefined, other Wi-Fi related properties will be ignored. Default value is TRUE. Returned by default.
    configureWifi *bool
    // Date time the enrollment profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description for the enrollment profile.
    description *string
    // Display name for the enrollment profile.
    displayName *string
    // Total number of Android devices that have enrolled using this enrollment profile.
    enrolledDeviceCount *int32
    // The enrollment mode for an enrollment profile.
    enrollmentMode *AndroidDeviceOwnerEnrollmentMode
    // The enrollment token type for an enrollment profile.
    enrollmentTokenType *AndroidDeviceOwnerEnrollmentTokenType
    // Total number of AOSP devices that have enrolled using the current token.
    enrollmentTokenUsageCount *int32
    // Boolean indicating if this profile is an Android AOSP for Teams device profile.
    isTeamsDeviceProfile *bool
    // Date time the enrollment profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // String used to generate a QR code for the token.
    qrCodeContent *string
    // String used to generate a QR code for the token.
    qrCodeImage MimeContentable
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // Date time the most recently created token was created.
    tokenCreationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date time the most recently created token will expire.
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Value of the most recently created token for this enrollment profile.
    tokenValue *string
    // Boolean that indicates if hidden wifi networks are enabled
    wifiHidden *bool
    // String that contains the wi-fi login password
    wifiPassword *string
    // This enum represents Wi-Fi Security Types for Android Device Owner AOSP Scenarios.
    wifiSecurityType *AospWifiSecurityType
    // String that contains the wi-fi login ssid
    wifiSsid *string
}
// NewAndroidDeviceOwnerEnrollmentProfile instantiates a new androidDeviceOwnerEnrollmentProfile and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfile()(*AndroidDeviceOwnerEnrollmentProfile) {
    m := &AndroidDeviceOwnerEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerEnrollmentProfile(), nil
}
// GetAccountId gets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetAccountId()(*string) {
    return m.accountId
}
// GetConfigureWifi gets the configureWifi property value. Boolean that indicates that the Wi-Fi network should be configured during device provisioning. When set to TRUE, device provisioning will use Wi-Fi related properties to automatically connect to Wi-Fi networks. When set to FALSE or undefined, other Wi-Fi related properties will be ignored. Default value is TRUE. Returned by default.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetConfigureWifi()(*bool) {
    return m.configureWifi
}
// GetCreatedDateTime gets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrolledDeviceCount gets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrolledDeviceCount()(*int32) {
    return m.enrolledDeviceCount
}
// GetEnrollmentMode gets the enrollmentMode property value. The enrollment mode for an enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentMode()(*AndroidDeviceOwnerEnrollmentMode) {
    return m.enrollmentMode
}
// GetEnrollmentTokenType gets the enrollmentTokenType property value. The enrollment token type for an enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenType()(*AndroidDeviceOwnerEnrollmentTokenType) {
    return m.enrollmentTokenType
}
// GetEnrollmentTokenUsageCount gets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenUsageCount()(*int32) {
    return m.enrollmentTokenUsageCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["configureWifi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigureWifi(val)
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
    res["enrollmentMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentMode(val.(*AndroidDeviceOwnerEnrollmentMode))
        }
        return nil
    }
    res["enrollmentTokenType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTokenType(val.(*AndroidDeviceOwnerEnrollmentTokenType))
        }
        return nil
    }
    res["enrollmentTokenUsageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTokenUsageCount(val)
        }
        return nil
    }
    res["isTeamsDeviceProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTeamsDeviceProfile(val)
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
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["tokenCreationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenCreationDateTime(val)
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
    res["wifiHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiHidden(val)
        }
        return nil
    }
    res["wifiPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiPassword(val)
        }
        return nil
    }
    res["wifiSecurityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAospWifiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiSecurityType(val.(*AospWifiSecurityType))
        }
        return nil
    }
    res["wifiSsid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiSsid(val)
        }
        return nil
    }
    return res
}
// GetIsTeamsDeviceProfile gets the isTeamsDeviceProfile property value. Boolean indicating if this profile is an Android AOSP for Teams device profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetIsTeamsDeviceProfile()(*bool) {
    return m.isTeamsDeviceProfile
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetQrCodeContent gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeContent()(*string) {
    return m.qrCodeContent
}
// GetQrCodeImage gets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeImage()(MimeContentable) {
    return m.qrCodeImage
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetTokenCreationDateTime gets the tokenCreationDateTime property value. Date time the most recently created token was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.tokenCreationDateTime
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.tokenExpirationDateTime
}
// GetTokenValue gets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenValue()(*string) {
    return m.tokenValue
}
// GetWifiHidden gets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiHidden()(*bool) {
    return m.wifiHidden
}
// GetWifiPassword gets the wifiPassword property value. String that contains the wi-fi login password
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiPassword()(*string) {
    return m.wifiPassword
}
// GetWifiSecurityType gets the wifiSecurityType property value. This enum represents Wi-Fi Security Types for Android Device Owner AOSP Scenarios.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSecurityType()(*AospWifiSecurityType) {
    return m.wifiSecurityType
}
// GetWifiSsid gets the wifiSsid property value. String that contains the wi-fi login ssid
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSsid()(*string) {
    return m.wifiSsid
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("configureWifi", m.GetConfigureWifi())
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
    if m.GetEnrollmentMode() != nil {
        cast := (*m.GetEnrollmentMode()).String()
        err = writer.WriteStringValue("enrollmentMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentTokenType() != nil {
        cast := (*m.GetEnrollmentTokenType()).String()
        err = writer.WriteStringValue("enrollmentTokenType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enrollmentTokenUsageCount", m.GetEnrollmentTokenUsageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTeamsDeviceProfile", m.GetIsTeamsDeviceProfile())
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
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("tokenCreationDateTime", m.GetTokenCreationDateTime())
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
    {
        err = writer.WriteBoolValue("wifiHidden", m.GetWifiHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("wifiPassword", m.GetWifiPassword())
        if err != nil {
            return err
        }
    }
    if m.GetWifiSecurityType() != nil {
        cast := (*m.GetWifiSecurityType()).String()
        err = writer.WriteStringValue("wifiSecurityType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("wifiSsid", m.GetWifiSsid())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountId sets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetAccountId(value *string)() {
    m.accountId = value
}
// SetConfigureWifi sets the configureWifi property value. Boolean that indicates that the Wi-Fi network should be configured during device provisioning. When set to TRUE, device provisioning will use Wi-Fi related properties to automatically connect to Wi-Fi networks. When set to FALSE or undefined, other Wi-Fi related properties will be ignored. Default value is TRUE. Returned by default.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetConfigureWifi(value *bool)() {
    m.configureWifi = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrolledDeviceCount sets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrolledDeviceCount(value *int32)() {
    m.enrolledDeviceCount = value
}
// SetEnrollmentMode sets the enrollmentMode property value. The enrollment mode for an enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentMode(value *AndroidDeviceOwnerEnrollmentMode)() {
    m.enrollmentMode = value
}
// SetEnrollmentTokenType sets the enrollmentTokenType property value. The enrollment token type for an enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenType(value *AndroidDeviceOwnerEnrollmentTokenType)() {
    m.enrollmentTokenType = value
}
// SetEnrollmentTokenUsageCount sets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenUsageCount(value *int32)() {
    m.enrollmentTokenUsageCount = value
}
// SetIsTeamsDeviceProfile sets the isTeamsDeviceProfile property value. Boolean indicating if this profile is an Android AOSP for Teams device profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetIsTeamsDeviceProfile(value *bool)() {
    m.isTeamsDeviceProfile = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetQrCodeContent sets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeContent(value *string)() {
    m.qrCodeContent = value
}
// SetQrCodeImage sets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeImage(value MimeContentable)() {
    m.qrCodeImage = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetTokenCreationDateTime sets the tokenCreationDateTime property value. Date time the most recently created token was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenCreationDateTime = value
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenExpirationDateTime = value
}
// SetTokenValue sets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenValue(value *string)() {
    m.tokenValue = value
}
// SetWifiHidden sets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiHidden(value *bool)() {
    m.wifiHidden = value
}
// SetWifiPassword sets the wifiPassword property value. String that contains the wi-fi login password
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiPassword(value *string)() {
    m.wifiPassword = value
}
// SetWifiSecurityType sets the wifiSecurityType property value. This enum represents Wi-Fi Security Types for Android Device Owner AOSP Scenarios.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSecurityType(value *AospWifiSecurityType)() {
    m.wifiSecurityType = value
}
// SetWifiSsid sets the wifiSsid property value. String that contains the wi-fi login ssid
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSsid(value *string)() {
    m.wifiSsid = value
}
