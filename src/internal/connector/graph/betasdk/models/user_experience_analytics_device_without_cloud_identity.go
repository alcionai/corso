package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceWithoutCloudIdentity the user experience analytics Device without Cloud Identity.
type UserExperienceAnalyticsDeviceWithoutCloudIdentity struct {
    Entity
    // Azure Active Directory Device Id
    azureAdDeviceId *string
    // The tenant attach device's name.
    deviceName *string
}
// NewUserExperienceAnalyticsDeviceWithoutCloudIdentity instantiates a new userExperienceAnalyticsDeviceWithoutCloudIdentity and sets the default values.
func NewUserExperienceAnalyticsDeviceWithoutCloudIdentity()(*UserExperienceAnalyticsDeviceWithoutCloudIdentity) {
    m := &UserExperienceAnalyticsDeviceWithoutCloudIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceWithoutCloudIdentity(), nil
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. Azure Active Directory Device Id
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetAzureAdDeviceId()(*string) {
    return m.azureAdDeviceId
}
// GetDeviceName gets the deviceName property value. The tenant attach device's name.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
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
    return res
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
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
    return nil
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. Azure Active Directory Device Id
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
// SetDeviceName sets the deviceName property value. The tenant attach device's name.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) SetDeviceName(value *string)() {
    m.deviceName = value
}
