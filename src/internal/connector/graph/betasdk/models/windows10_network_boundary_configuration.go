package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10NetworkBoundaryConfiguration 
type Windows10NetworkBoundaryConfiguration struct {
    DeviceConfiguration
    // Windows Network Isolation Policy
    windowsNetworkIsolationPolicy WindowsNetworkIsolationPolicyable
}
// NewWindows10NetworkBoundaryConfiguration instantiates a new Windows10NetworkBoundaryConfiguration and sets the default values.
func NewWindows10NetworkBoundaryConfiguration()(*Windows10NetworkBoundaryConfiguration) {
    m := &Windows10NetworkBoundaryConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10NetworkBoundaryConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10NetworkBoundaryConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10NetworkBoundaryConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10NetworkBoundaryConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10NetworkBoundaryConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["windowsNetworkIsolationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsNetworkIsolationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsNetworkIsolationPolicy(val.(WindowsNetworkIsolationPolicyable))
        }
        return nil
    }
    return res
}
// GetWindowsNetworkIsolationPolicy gets the windowsNetworkIsolationPolicy property value. Windows Network Isolation Policy
func (m *Windows10NetworkBoundaryConfiguration) GetWindowsNetworkIsolationPolicy()(WindowsNetworkIsolationPolicyable) {
    return m.windowsNetworkIsolationPolicy
}
// Serialize serializes information the current object
func (m *Windows10NetworkBoundaryConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("windowsNetworkIsolationPolicy", m.GetWindowsNetworkIsolationPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWindowsNetworkIsolationPolicy sets the windowsNetworkIsolationPolicy property value. Windows Network Isolation Policy
func (m *Windows10NetworkBoundaryConfiguration) SetWindowsNetworkIsolationPolicy(value WindowsNetworkIsolationPolicyable)() {
    m.windowsNetworkIsolationPolicy = value
}
