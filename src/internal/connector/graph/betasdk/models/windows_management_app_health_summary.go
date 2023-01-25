package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagementAppHealthSummary 
type WindowsManagementAppHealthSummary struct {
    Entity
    // Healthy device count.
    healthyDeviceCount *int32
    // Unhealthy device count.
    unhealthyDeviceCount *int32
    // Unknown device count.
    unknownDeviceCount *int32
}
// NewWindowsManagementAppHealthSummary instantiates a new WindowsManagementAppHealthSummary and sets the default values.
func NewWindowsManagementAppHealthSummary()(*WindowsManagementAppHealthSummary) {
    m := &WindowsManagementAppHealthSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsManagementAppHealthSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsManagementAppHealthSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagementAppHealthSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagementAppHealthSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthyDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthyDeviceCount(val)
        }
        return nil
    }
    res["unhealthyDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnhealthyDeviceCount(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetHealthyDeviceCount gets the healthyDeviceCount property value. Healthy device count.
func (m *WindowsManagementAppHealthSummary) GetHealthyDeviceCount()(*int32) {
    return m.healthyDeviceCount
}
// GetUnhealthyDeviceCount gets the unhealthyDeviceCount property value. Unhealthy device count.
func (m *WindowsManagementAppHealthSummary) GetUnhealthyDeviceCount()(*int32) {
    return m.unhealthyDeviceCount
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Unknown device count.
func (m *WindowsManagementAppHealthSummary) GetUnknownDeviceCount()(*int32) {
    return m.unknownDeviceCount
}
// Serialize serializes information the current object
func (m *WindowsManagementAppHealthSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("healthyDeviceCount", m.GetHealthyDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unhealthyDeviceCount", m.GetUnhealthyDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHealthyDeviceCount sets the healthyDeviceCount property value. Healthy device count.
func (m *WindowsManagementAppHealthSummary) SetHealthyDeviceCount(value *int32)() {
    m.healthyDeviceCount = value
}
// SetUnhealthyDeviceCount sets the unhealthyDeviceCount property value. Unhealthy device count.
func (m *WindowsManagementAppHealthSummary) SetUnhealthyDeviceCount(value *int32)() {
    m.unhealthyDeviceCount = value
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Unknown device count.
func (m *WindowsManagementAppHealthSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
