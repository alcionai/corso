package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkNineWorkEasConfiguration 
type AndroidForWorkNineWorkEasConfiguration struct {
    AndroidForWorkEasEmailProfileBase
    // Toggles syncing the calendar. If set to false the calendar is turned off on the device.
    syncCalendar *bool
    // Toggles syncing contacts. If set to false contacts are turned off on the device.
    syncContacts *bool
    // Toggles syncing tasks. If set to false tasks are turned off on the device.
    syncTasks *bool
}
// NewAndroidForWorkNineWorkEasConfiguration instantiates a new AndroidForWorkNineWorkEasConfiguration and sets the default values.
func NewAndroidForWorkNineWorkEasConfiguration()(*AndroidForWorkNineWorkEasConfiguration) {
    m := &AndroidForWorkNineWorkEasConfiguration{
        AndroidForWorkEasEmailProfileBase: *NewAndroidForWorkEasEmailProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkNineWorkEasConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidForWorkNineWorkEasConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkNineWorkEasConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkNineWorkEasConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkNineWorkEasConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidForWorkEasEmailProfileBase.GetFieldDeserializers()
    res["syncCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncCalendar(val)
        }
        return nil
    }
    res["syncContacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncContacts(val)
        }
        return nil
    }
    res["syncTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncTasks(val)
        }
        return nil
    }
    return res
}
// GetSyncCalendar gets the syncCalendar property value. Toggles syncing the calendar. If set to false the calendar is turned off on the device.
func (m *AndroidForWorkNineWorkEasConfiguration) GetSyncCalendar()(*bool) {
    return m.syncCalendar
}
// GetSyncContacts gets the syncContacts property value. Toggles syncing contacts. If set to false contacts are turned off on the device.
func (m *AndroidForWorkNineWorkEasConfiguration) GetSyncContacts()(*bool) {
    return m.syncContacts
}
// GetSyncTasks gets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
func (m *AndroidForWorkNineWorkEasConfiguration) GetSyncTasks()(*bool) {
    return m.syncTasks
}
// Serialize serializes information the current object
func (m *AndroidForWorkNineWorkEasConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidForWorkEasEmailProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("syncCalendar", m.GetSyncCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncContacts", m.GetSyncContacts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncTasks", m.GetSyncTasks())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSyncCalendar sets the syncCalendar property value. Toggles syncing the calendar. If set to false the calendar is turned off on the device.
func (m *AndroidForWorkNineWorkEasConfiguration) SetSyncCalendar(value *bool)() {
    m.syncCalendar = value
}
// SetSyncContacts sets the syncContacts property value. Toggles syncing contacts. If set to false contacts are turned off on the device.
func (m *AndroidForWorkNineWorkEasConfiguration) SetSyncContacts(value *bool)() {
    m.syncContacts = value
}
// SetSyncTasks sets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
func (m *AndroidForWorkNineWorkEasConfiguration) SetSyncTasks(value *bool)() {
    m.syncTasks = value
}
