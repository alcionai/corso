package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StrongAuthenticationPhoneAppDetail 
type StrongAuthenticationPhoneAppDetail struct {
    Entity
    // The authenticationType property
    authenticationType *string
    // The authenticatorFlavor property
    authenticatorFlavor *string
    // The deviceId property
    deviceId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The deviceName property
    deviceName *string
    // The deviceTag property
    deviceTag *string
    // The deviceToken property
    deviceToken *string
    // The hashFunction property
    hashFunction *string
    // The lastAuthenticatedDateTime property
    lastAuthenticatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The notificationType property
    notificationType *string
    // The oathSecretKey property
    oathSecretKey *string
    // The oathTokenMetadata property
    oathTokenMetadata OathTokenMetadataable
    // The oathTokenTimeDriftInSeconds property
    oathTokenTimeDriftInSeconds *int32
    // The phoneAppVersion property
    phoneAppVersion *string
    // The tenantDeviceId property
    tenantDeviceId *string
    // The tokenGenerationIntervalInSeconds property
    tokenGenerationIntervalInSeconds *int32
}
// NewStrongAuthenticationPhoneAppDetail instantiates a new StrongAuthenticationPhoneAppDetail and sets the default values.
func NewStrongAuthenticationPhoneAppDetail()(*StrongAuthenticationPhoneAppDetail) {
    m := &StrongAuthenticationPhoneAppDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateStrongAuthenticationPhoneAppDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStrongAuthenticationPhoneAppDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStrongAuthenticationPhoneAppDetail(), nil
}
// GetAuthenticationType gets the authenticationType property value. The authenticationType property
func (m *StrongAuthenticationPhoneAppDetail) GetAuthenticationType()(*string) {
    return m.authenticationType
}
// GetAuthenticatorFlavor gets the authenticatorFlavor property value. The authenticatorFlavor property
func (m *StrongAuthenticationPhoneAppDetail) GetAuthenticatorFlavor()(*string) {
    return m.authenticatorFlavor
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceTag gets the deviceTag property value. The deviceTag property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceTag()(*string) {
    return m.deviceTag
}
// GetDeviceToken gets the deviceToken property value. The deviceToken property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceToken()(*string) {
    return m.deviceToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StrongAuthenticationPhoneAppDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val)
        }
        return nil
    }
    res["authenticatorFlavor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticatorFlavor(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    res["deviceToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceToken(val)
        }
        return nil
    }
    res["hashFunction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashFunction(val)
        }
        return nil
    }
    res["lastAuthenticatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAuthenticatedDateTime(val)
        }
        return nil
    }
    res["notificationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationType(val)
        }
        return nil
    }
    res["oathSecretKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOathSecretKey(val)
        }
        return nil
    }
    res["oathTokenMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOathTokenMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOathTokenMetadata(val.(OathTokenMetadataable))
        }
        return nil
    }
    res["oathTokenTimeDriftInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOathTokenTimeDriftInSeconds(val)
        }
        return nil
    }
    res["phoneAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneAppVersion(val)
        }
        return nil
    }
    res["tenantDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDeviceId(val)
        }
        return nil
    }
    res["tokenGenerationIntervalInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenGenerationIntervalInSeconds(val)
        }
        return nil
    }
    return res
}
// GetHashFunction gets the hashFunction property value. The hashFunction property
func (m *StrongAuthenticationPhoneAppDetail) GetHashFunction()(*string) {
    return m.hashFunction
}
// GetLastAuthenticatedDateTime gets the lastAuthenticatedDateTime property value. The lastAuthenticatedDateTime property
func (m *StrongAuthenticationPhoneAppDetail) GetLastAuthenticatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastAuthenticatedDateTime
}
// GetNotificationType gets the notificationType property value. The notificationType property
func (m *StrongAuthenticationPhoneAppDetail) GetNotificationType()(*string) {
    return m.notificationType
}
// GetOathSecretKey gets the oathSecretKey property value. The oathSecretKey property
func (m *StrongAuthenticationPhoneAppDetail) GetOathSecretKey()(*string) {
    return m.oathSecretKey
}
// GetOathTokenMetadata gets the oathTokenMetadata property value. The oathTokenMetadata property
func (m *StrongAuthenticationPhoneAppDetail) GetOathTokenMetadata()(OathTokenMetadataable) {
    return m.oathTokenMetadata
}
// GetOathTokenTimeDriftInSeconds gets the oathTokenTimeDriftInSeconds property value. The oathTokenTimeDriftInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) GetOathTokenTimeDriftInSeconds()(*int32) {
    return m.oathTokenTimeDriftInSeconds
}
// GetPhoneAppVersion gets the phoneAppVersion property value. The phoneAppVersion property
func (m *StrongAuthenticationPhoneAppDetail) GetPhoneAppVersion()(*string) {
    return m.phoneAppVersion
}
// GetTenantDeviceId gets the tenantDeviceId property value. The tenantDeviceId property
func (m *StrongAuthenticationPhoneAppDetail) GetTenantDeviceId()(*string) {
    return m.tenantDeviceId
}
// GetTokenGenerationIntervalInSeconds gets the tokenGenerationIntervalInSeconds property value. The tokenGenerationIntervalInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) GetTokenGenerationIntervalInSeconds()(*int32) {
    return m.tokenGenerationIntervalInSeconds
}
// Serialize serializes information the current object
func (m *StrongAuthenticationPhoneAppDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationType", m.GetAuthenticationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authenticatorFlavor", m.GetAuthenticatorFlavor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceTag", m.GetDeviceTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceToken", m.GetDeviceToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hashFunction", m.GetHashFunction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastAuthenticatedDateTime", m.GetLastAuthenticatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationType", m.GetNotificationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("oathSecretKey", m.GetOathSecretKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("oathTokenMetadata", m.GetOathTokenMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("oathTokenTimeDriftInSeconds", m.GetOathTokenTimeDriftInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneAppVersion", m.GetPhoneAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDeviceId", m.GetTenantDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("tokenGenerationIntervalInSeconds", m.GetTokenGenerationIntervalInSeconds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationType sets the authenticationType property value. The authenticationType property
func (m *StrongAuthenticationPhoneAppDetail) SetAuthenticationType(value *string)() {
    m.authenticationType = value
}
// SetAuthenticatorFlavor sets the authenticatorFlavor property value. The authenticatorFlavor property
func (m *StrongAuthenticationPhoneAppDetail) SetAuthenticatorFlavor(value *string)() {
    m.authenticatorFlavor = value
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceTag sets the deviceTag property value. The deviceTag property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
// SetDeviceToken sets the deviceToken property value. The deviceToken property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceToken(value *string)() {
    m.deviceToken = value
}
// SetHashFunction sets the hashFunction property value. The hashFunction property
func (m *StrongAuthenticationPhoneAppDetail) SetHashFunction(value *string)() {
    m.hashFunction = value
}
// SetLastAuthenticatedDateTime sets the lastAuthenticatedDateTime property value. The lastAuthenticatedDateTime property
func (m *StrongAuthenticationPhoneAppDetail) SetLastAuthenticatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastAuthenticatedDateTime = value
}
// SetNotificationType sets the notificationType property value. The notificationType property
func (m *StrongAuthenticationPhoneAppDetail) SetNotificationType(value *string)() {
    m.notificationType = value
}
// SetOathSecretKey sets the oathSecretKey property value. The oathSecretKey property
func (m *StrongAuthenticationPhoneAppDetail) SetOathSecretKey(value *string)() {
    m.oathSecretKey = value
}
// SetOathTokenMetadata sets the oathTokenMetadata property value. The oathTokenMetadata property
func (m *StrongAuthenticationPhoneAppDetail) SetOathTokenMetadata(value OathTokenMetadataable)() {
    m.oathTokenMetadata = value
}
// SetOathTokenTimeDriftInSeconds sets the oathTokenTimeDriftInSeconds property value. The oathTokenTimeDriftInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) SetOathTokenTimeDriftInSeconds(value *int32)() {
    m.oathTokenTimeDriftInSeconds = value
}
// SetPhoneAppVersion sets the phoneAppVersion property value. The phoneAppVersion property
func (m *StrongAuthenticationPhoneAppDetail) SetPhoneAppVersion(value *string)() {
    m.phoneAppVersion = value
}
// SetTenantDeviceId sets the tenantDeviceId property value. The tenantDeviceId property
func (m *StrongAuthenticationPhoneAppDetail) SetTenantDeviceId(value *string)() {
    m.tenantDeviceId = value
}
// SetTokenGenerationIntervalInSeconds sets the tokenGenerationIntervalInSeconds property value. The tokenGenerationIntervalInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) SetTokenGenerationIntervalInSeconds(value *int32)() {
    m.tokenGenerationIntervalInSeconds = value
}
