package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComanagementAuthorityConfiguration 
type DeviceComanagementAuthorityConfiguration struct {
    DeviceEnrollmentConfiguration
    // CoManagement Authority configuration ConfigurationManagerAgentCommandLineArgument
    configurationManagerAgentCommandLineArgument *string
    // CoManagement Authority configuration InstallConfigurationManagerAgent
    installConfigurationManagerAgent *bool
    // CoManagement Authority configuration ManagedDeviceAuthority
    managedDeviceAuthority *int32
}
// NewDeviceComanagementAuthorityConfiguration instantiates a new DeviceComanagementAuthorityConfiguration and sets the default values.
func NewDeviceComanagementAuthorityConfiguration()(*DeviceComanagementAuthorityConfiguration) {
    m := &DeviceComanagementAuthorityConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.deviceComanagementAuthorityConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceComanagementAuthorityConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComanagementAuthorityConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComanagementAuthorityConfiguration(), nil
}
// GetConfigurationManagerAgentCommandLineArgument gets the configurationManagerAgentCommandLineArgument property value. CoManagement Authority configuration ConfigurationManagerAgentCommandLineArgument
func (m *DeviceComanagementAuthorityConfiguration) GetConfigurationManagerAgentCommandLineArgument()(*string) {
    return m.configurationManagerAgentCommandLineArgument
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComanagementAuthorityConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["configurationManagerAgentCommandLineArgument"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerAgentCommandLineArgument(val)
        }
        return nil
    }
    res["installConfigurationManagerAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallConfigurationManagerAgent(val)
        }
        return nil
    }
    res["managedDeviceAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceAuthority(val)
        }
        return nil
    }
    return res
}
// GetInstallConfigurationManagerAgent gets the installConfigurationManagerAgent property value. CoManagement Authority configuration InstallConfigurationManagerAgent
func (m *DeviceComanagementAuthorityConfiguration) GetInstallConfigurationManagerAgent()(*bool) {
    return m.installConfigurationManagerAgent
}
// GetManagedDeviceAuthority gets the managedDeviceAuthority property value. CoManagement Authority configuration ManagedDeviceAuthority
func (m *DeviceComanagementAuthorityConfiguration) GetManagedDeviceAuthority()(*int32) {
    return m.managedDeviceAuthority
}
// Serialize serializes information the current object
func (m *DeviceComanagementAuthorityConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("configurationManagerAgentCommandLineArgument", m.GetConfigurationManagerAgentCommandLineArgument())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("installConfigurationManagerAgent", m.GetInstallConfigurationManagerAgent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("managedDeviceAuthority", m.GetManagedDeviceAuthority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationManagerAgentCommandLineArgument sets the configurationManagerAgentCommandLineArgument property value. CoManagement Authority configuration ConfigurationManagerAgentCommandLineArgument
func (m *DeviceComanagementAuthorityConfiguration) SetConfigurationManagerAgentCommandLineArgument(value *string)() {
    m.configurationManagerAgentCommandLineArgument = value
}
// SetInstallConfigurationManagerAgent sets the installConfigurationManagerAgent property value. CoManagement Authority configuration InstallConfigurationManagerAgent
func (m *DeviceComanagementAuthorityConfiguration) SetInstallConfigurationManagerAgent(value *bool)() {
    m.installConfigurationManagerAgent = value
}
// SetManagedDeviceAuthority sets the managedDeviceAuthority property value. CoManagement Authority configuration ManagedDeviceAuthority
func (m *DeviceComanagementAuthorityConfiguration) SetManagedDeviceAuthority(value *int32)() {
    m.managedDeviceAuthority = value
}
