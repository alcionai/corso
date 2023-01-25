package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceRegistrationPolicy 
type DeviceRegistrationPolicy struct {
    Entity
    // Specifies the authorization policy for controlling registration of new devices using Azure AD Join within your organization. Required. For more information, see What is a device identity?.
    azureADJoin AzureAdJoinPolicyable
    // Specifies the authorization policy for controlling registration of new devices using Azure AD registered within your organization. Required. For more information, see What is a device identity?.
    azureADRegistration AzureADRegistrationPolicyable
    // The description of the device registration policy. It is always set to Tenant-wide policy that manages intial provisioning controls using quota restrictions, additional authentication and authorization checks. Read-only.
    description *string
    // The name of the device registration policy. It is always set to Device Registration Policy. Read-only.
    displayName *string
    // The multiFactorAuthConfiguration property
    multiFactorAuthConfiguration *MultiFactorAuthConfiguration
    // Specifies the maximum number of devices that a user can have within your organization before blocking new device registrations. The default value is set to 50. If this property is not specified during the policy update operation, it is automatically reset to 0 to indicate that users are not allowed to join any devices.
    userDeviceQuota *int32
}
// NewDeviceRegistrationPolicy instantiates a new DeviceRegistrationPolicy and sets the default values.
func NewDeviceRegistrationPolicy()(*DeviceRegistrationPolicy) {
    m := &DeviceRegistrationPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceRegistrationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceRegistrationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceRegistrationPolicy(), nil
}
// GetAzureADJoin gets the azureADJoin property value. Specifies the authorization policy for controlling registration of new devices using Azure AD Join within your organization. Required. For more information, see What is a device identity?.
func (m *DeviceRegistrationPolicy) GetAzureADJoin()(AzureAdJoinPolicyable) {
    return m.azureADJoin
}
// GetAzureADRegistration gets the azureADRegistration property value. Specifies the authorization policy for controlling registration of new devices using Azure AD registered within your organization. Required. For more information, see What is a device identity?.
func (m *DeviceRegistrationPolicy) GetAzureADRegistration()(AzureADRegistrationPolicyable) {
    return m.azureADRegistration
}
// GetDescription gets the description property value. The description of the device registration policy. It is always set to Tenant-wide policy that manages intial provisioning controls using quota restrictions, additional authentication and authorization checks. Read-only.
func (m *DeviceRegistrationPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the device registration policy. It is always set to Device Registration Policy. Read-only.
func (m *DeviceRegistrationPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceRegistrationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureADJoin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureAdJoinPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADJoin(val.(AzureAdJoinPolicyable))
        }
        return nil
    }
    res["azureADRegistration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureADRegistrationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADRegistration(val.(AzureADRegistrationPolicyable))
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
    res["multiFactorAuthConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiFactorAuthConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultiFactorAuthConfiguration(val.(*MultiFactorAuthConfiguration))
        }
        return nil
    }
    res["userDeviceQuota"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDeviceQuota(val)
        }
        return nil
    }
    return res
}
// GetMultiFactorAuthConfiguration gets the multiFactorAuthConfiguration property value. The multiFactorAuthConfiguration property
func (m *DeviceRegistrationPolicy) GetMultiFactorAuthConfiguration()(*MultiFactorAuthConfiguration) {
    return m.multiFactorAuthConfiguration
}
// GetUserDeviceQuota gets the userDeviceQuota property value. Specifies the maximum number of devices that a user can have within your organization before blocking new device registrations. The default value is set to 50. If this property is not specified during the policy update operation, it is automatically reset to 0 to indicate that users are not allowed to join any devices.
func (m *DeviceRegistrationPolicy) GetUserDeviceQuota()(*int32) {
    return m.userDeviceQuota
}
// Serialize serializes information the current object
func (m *DeviceRegistrationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("azureADJoin", m.GetAzureADJoin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("azureADRegistration", m.GetAzureADRegistration())
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
    if m.GetMultiFactorAuthConfiguration() != nil {
        cast := (*m.GetMultiFactorAuthConfiguration()).String()
        err = writer.WriteStringValue("multiFactorAuthConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("userDeviceQuota", m.GetUserDeviceQuota())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureADJoin sets the azureADJoin property value. Specifies the authorization policy for controlling registration of new devices using Azure AD Join within your organization. Required. For more information, see What is a device identity?.
func (m *DeviceRegistrationPolicy) SetAzureADJoin(value AzureAdJoinPolicyable)() {
    m.azureADJoin = value
}
// SetAzureADRegistration sets the azureADRegistration property value. Specifies the authorization policy for controlling registration of new devices using Azure AD registered within your organization. Required. For more information, see What is a device identity?.
func (m *DeviceRegistrationPolicy) SetAzureADRegistration(value AzureADRegistrationPolicyable)() {
    m.azureADRegistration = value
}
// SetDescription sets the description property value. The description of the device registration policy. It is always set to Tenant-wide policy that manages intial provisioning controls using quota restrictions, additional authentication and authorization checks. Read-only.
func (m *DeviceRegistrationPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the device registration policy. It is always set to Device Registration Policy. Read-only.
func (m *DeviceRegistrationPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMultiFactorAuthConfiguration sets the multiFactorAuthConfiguration property value. The multiFactorAuthConfiguration property
func (m *DeviceRegistrationPolicy) SetMultiFactorAuthConfiguration(value *MultiFactorAuthConfiguration)() {
    m.multiFactorAuthConfiguration = value
}
// SetUserDeviceQuota sets the userDeviceQuota property value. Specifies the maximum number of devices that a user can have within your organization before blocking new device registrations. The default value is set to 50. If this property is not specified during the policy update operation, it is automatically reset to 0 to indicate that users are not allowed to join any devices.
func (m *DeviceRegistrationPolicy) SetUserDeviceQuota(value *int32)() {
    m.userDeviceQuota = value
}
