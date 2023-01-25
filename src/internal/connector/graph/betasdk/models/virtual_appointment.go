package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualAppointment 
type VirtualAppointment struct {
    Entity
    // The join web URL of the virtual appointment for clients with waiting room and browser join. Optional.
    appointmentClientJoinWebUrl *string
    // The client information for the virtual appointment, including name, email, and SMS phone number. Optional.
    appointmentClients []VirtualAppointmentUserable
    // The identifier of the appointment from the scheduling system, associated with the current virtual appointment. Optional.
    externalAppointmentId *string
    // The URL of the appointment resource from the scheduling system, associated with the current virtual appointment. Optional.
    externalAppointmentUrl *string
    // The settings associated with the virtual appointment resource. Optional.
    settings VirtualAppointmentSettingsable
}
// NewVirtualAppointment instantiates a new virtualAppointment and sets the default values.
func NewVirtualAppointment()(*VirtualAppointment) {
    m := &VirtualAppointment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualAppointmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualAppointmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualAppointment(), nil
}
// GetAppointmentClientJoinWebUrl gets the appointmentClientJoinWebUrl property value. The join web URL of the virtual appointment for clients with waiting room and browser join. Optional.
func (m *VirtualAppointment) GetAppointmentClientJoinWebUrl()(*string) {
    return m.appointmentClientJoinWebUrl
}
// GetAppointmentClients gets the appointmentClients property value. The client information for the virtual appointment, including name, email, and SMS phone number. Optional.
func (m *VirtualAppointment) GetAppointmentClients()([]VirtualAppointmentUserable) {
    return m.appointmentClients
}
// GetExternalAppointmentId gets the externalAppointmentId property value. The identifier of the appointment from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) GetExternalAppointmentId()(*string) {
    return m.externalAppointmentId
}
// GetExternalAppointmentUrl gets the externalAppointmentUrl property value. The URL of the appointment resource from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) GetExternalAppointmentUrl()(*string) {
    return m.externalAppointmentUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualAppointment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appointmentClientJoinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppointmentClientJoinWebUrl(val)
        }
        return nil
    }
    res["appointmentClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualAppointmentUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualAppointmentUserable, len(val))
            for i, v := range val {
                res[i] = v.(VirtualAppointmentUserable)
            }
            m.SetAppointmentClients(res)
        }
        return nil
    }
    res["externalAppointmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAppointmentId(val)
        }
        return nil
    }
    res["externalAppointmentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAppointmentUrl(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualAppointmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(VirtualAppointmentSettingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The settings associated with the virtual appointment resource. Optional.
func (m *VirtualAppointment) GetSettings()(VirtualAppointmentSettingsable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *VirtualAppointment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appointmentClientJoinWebUrl", m.GetAppointmentClientJoinWebUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAppointmentClients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppointmentClients()))
        for i, v := range m.GetAppointmentClients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appointmentClients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalAppointmentId", m.GetExternalAppointmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalAppointmentUrl", m.GetExternalAppointmentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppointmentClientJoinWebUrl sets the appointmentClientJoinWebUrl property value. The join web URL of the virtual appointment for clients with waiting room and browser join. Optional.
func (m *VirtualAppointment) SetAppointmentClientJoinWebUrl(value *string)() {
    m.appointmentClientJoinWebUrl = value
}
// SetAppointmentClients sets the appointmentClients property value. The client information for the virtual appointment, including name, email, and SMS phone number. Optional.
func (m *VirtualAppointment) SetAppointmentClients(value []VirtualAppointmentUserable)() {
    m.appointmentClients = value
}
// SetExternalAppointmentId sets the externalAppointmentId property value. The identifier of the appointment from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) SetExternalAppointmentId(value *string)() {
    m.externalAppointmentId = value
}
// SetExternalAppointmentUrl sets the externalAppointmentUrl property value. The URL of the appointment resource from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) SetExternalAppointmentUrl(value *string)() {
    m.externalAppointmentUrl = value
}
// SetSettings sets the settings property value. The settings associated with the virtual appointment resource. Optional.
func (m *VirtualAppointment) SetSettings(value VirtualAppointmentSettingsable)() {
    m.settings = value
}
