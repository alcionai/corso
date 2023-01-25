package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppAssignedDeviceLicense 
type IosVppAppAssignedDeviceLicense struct {
    IosVppAppAssignedLicense
    // The device name.
    deviceName *string
    // The managed device ID.
    managedDeviceId *string
}
// NewIosVppAppAssignedDeviceLicense instantiates a new IosVppAppAssignedDeviceLicense and sets the default values.
func NewIosVppAppAssignedDeviceLicense()(*IosVppAppAssignedDeviceLicense) {
    m := &IosVppAppAssignedDeviceLicense{
        IosVppAppAssignedLicense: *NewIosVppAppAssignedLicense(),
    }
    return m
}
// CreateIosVppAppAssignedDeviceLicenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppAssignedDeviceLicenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppAppAssignedDeviceLicense(), nil
}
// GetDeviceName gets the deviceName property value. The device name.
func (m *IosVppAppAssignedDeviceLicense) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppAppAssignedDeviceLicense) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosVppAppAssignedLicense.GetFieldDeserializers()
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
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. The managed device ID.
func (m *IosVppAppAssignedDeviceLicense) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// Serialize serializes information the current object
func (m *IosVppAppAssignedDeviceLicense) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosVppAppAssignedLicense.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceName sets the deviceName property value. The device name.
func (m *IosVppAppAssignedDeviceLicense) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The managed device ID.
func (m *IosVppAppAssignedDeviceLicense) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
