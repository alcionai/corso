package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteAssistanceSettings 
type RemoteAssistanceSettings struct {
    Entity
    // Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
    allowSessionsToUnenrolledDevices *bool
    // Indicates if sessions to block chat function. This setting is configurable by the admin. Default value is false.
    blockChat *bool
    // State of remote assistance for the account
    remoteAssistanceState *RemoteAssistanceState
}
// NewRemoteAssistanceSettings instantiates a new remoteAssistanceSettings and sets the default values.
func NewRemoteAssistanceSettings()(*RemoteAssistanceSettings) {
    m := &RemoteAssistanceSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRemoteAssistanceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoteAssistanceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteAssistanceSettings(), nil
}
// GetAllowSessionsToUnenrolledDevices gets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) GetAllowSessionsToUnenrolledDevices()(*bool) {
    return m.allowSessionsToUnenrolledDevices
}
// GetBlockChat gets the blockChat property value. Indicates if sessions to block chat function. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) GetBlockChat()(*bool) {
    return m.blockChat
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoteAssistanceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowSessionsToUnenrolledDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSessionsToUnenrolledDevices(val)
        }
        return nil
    }
    res["blockChat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockChat(val)
        }
        return nil
    }
    res["remoteAssistanceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAssistanceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceState(val.(*RemoteAssistanceState))
        }
        return nil
    }
    return res
}
// GetRemoteAssistanceState gets the remoteAssistanceState property value. State of remote assistance for the account
func (m *RemoteAssistanceSettings) GetRemoteAssistanceState()(*RemoteAssistanceState) {
    return m.remoteAssistanceState
}
// Serialize serializes information the current object
func (m *RemoteAssistanceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowSessionsToUnenrolledDevices", m.GetAllowSessionsToUnenrolledDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockChat", m.GetBlockChat())
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAssistanceState() != nil {
        cast := (*m.GetRemoteAssistanceState()).String()
        err = writer.WriteStringValue("remoteAssistanceState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowSessionsToUnenrolledDevices sets the allowSessionsToUnenrolledDevices property value. Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) SetAllowSessionsToUnenrolledDevices(value *bool)() {
    m.allowSessionsToUnenrolledDevices = value
}
// SetBlockChat sets the blockChat property value. Indicates if sessions to block chat function. This setting is configurable by the admin. Default value is false.
func (m *RemoteAssistanceSettings) SetBlockChat(value *bool)() {
    m.blockChat = value
}
// SetRemoteAssistanceState sets the remoteAssistanceState property value. State of remote assistance for the account
func (m *RemoteAssistanceSettings) SetRemoteAssistanceState(value *RemoteAssistanceState)() {
    m.remoteAssistanceState = value
}
