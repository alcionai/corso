package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VulnerableManagedDevice this entity represents a device associated with a task.
type VulnerableManagedDevice struct {
    Entity
    // The device name.
    displayName *string
    // The last sync date.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Intune managed device ID.
    managedDeviceId *string
}
// NewVulnerableManagedDevice instantiates a new vulnerableManagedDevice and sets the default values.
func NewVulnerableManagedDevice()(*VulnerableManagedDevice) {
    m := &VulnerableManagedDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVulnerableManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVulnerableManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVulnerableManagedDevice(), nil
}
// GetDisplayName gets the displayName property value. The device name.
func (m *VulnerableManagedDevice) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VulnerableManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
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
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last sync date.
func (m *VulnerableManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. The Intune managed device ID.
func (m *VulnerableManagedDevice) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// Serialize serializes information the current object
func (m *VulnerableManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
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
// SetDisplayName sets the displayName property value. The device name.
func (m *VulnerableManagedDevice) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last sync date.
func (m *VulnerableManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The Intune managed device ID.
func (m *VulnerableManagedDevice) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
