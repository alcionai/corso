package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotSettings 
type WindowsAutopilotSettings struct {
    Entity
    // Last data sync date time with DDS service.
    lastManualSyncTriggerDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last data sync date time with DDS service.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The syncStatus property
    syncStatus *WindowsAutopilotSyncStatus
}
// NewWindowsAutopilotSettings instantiates a new windowsAutopilotSettings and sets the default values.
func NewWindowsAutopilotSettings()(*WindowsAutopilotSettings) {
    m := &WindowsAutopilotSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsAutopilotSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastManualSyncTriggerDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastManualSyncTriggerDateTime(val)
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
    res["syncStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncStatus(val.(*WindowsAutopilotSyncStatus))
        }
        return nil
    }
    return res
}
// GetLastManualSyncTriggerDateTime gets the lastManualSyncTriggerDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) GetLastManualSyncTriggerDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastManualSyncTriggerDateTime
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetSyncStatus gets the syncStatus property value. The syncStatus property
func (m *WindowsAutopilotSettings) GetSyncStatus()(*WindowsAutopilotSyncStatus) {
    return m.syncStatus
}
// Serialize serializes information the current object
func (m *WindowsAutopilotSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastManualSyncTriggerDateTime", m.GetLastManualSyncTriggerDateTime())
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
    if m.GetSyncStatus() != nil {
        cast := (*m.GetSyncStatus()).String()
        err = writer.WriteStringValue("syncStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastManualSyncTriggerDateTime sets the lastManualSyncTriggerDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) SetLastManualSyncTriggerDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastManualSyncTriggerDateTime = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last data sync date time with DDS service.
func (m *WindowsAutopilotSettings) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetSyncStatus sets the syncStatus property value. The syncStatus property
func (m *WindowsAutopilotSettings) SetSyncStatus(value *WindowsAutopilotSyncStatus)() {
    m.syncStatus = value
}
