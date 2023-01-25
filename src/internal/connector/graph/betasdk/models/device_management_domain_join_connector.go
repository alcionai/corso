package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDomainJoinConnector a Domain Join Connector is a connector that is responsible to allocate (and delete) machine account blobs
type DeviceManagementDomainJoinConnector struct {
    Entity
    // The connector display name.
    displayName *string
    // Last time connector contacted Intune.
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ODJ request states.
    state *DeviceManagementDomainJoinConnectorState
    // The version of the connector.
    version *string
}
// NewDeviceManagementDomainJoinConnector instantiates a new deviceManagementDomainJoinConnector and sets the default values.
func NewDeviceManagementDomainJoinConnector()(*DeviceManagementDomainJoinConnector) {
    m := &DeviceManagementDomainJoinConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDomainJoinConnector(), nil
}
// GetDisplayName gets the displayName property value. The connector display name.
func (m *DeviceManagementDomainJoinConnector) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDomainJoinConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["lastConnectionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConnectionDateTime(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementDomainJoinConnectorState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*DeviceManagementDomainJoinConnectorState))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Last time connector contacted Intune.
func (m *DeviceManagementDomainJoinConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastConnectionDateTime
}
// GetState gets the state property value. The ODJ request states.
func (m *DeviceManagementDomainJoinConnector) GetState()(*DeviceManagementDomainJoinConnectorState) {
    return m.state
}
// GetVersion gets the version property value. The version of the connector.
func (m *DeviceManagementDomainJoinConnector) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceManagementDomainJoinConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("lastConnectionDateTime", m.GetLastConnectionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The connector display name.
func (m *DeviceManagementDomainJoinConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Last time connector contacted Intune.
func (m *DeviceManagementDomainJoinConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// SetState sets the state property value. The ODJ request states.
func (m *DeviceManagementDomainJoinConnector) SetState(value *DeviceManagementDomainJoinConnectorState)() {
    m.state = value
}
// SetVersion sets the version property value. The version of the connector.
func (m *DeviceManagementDomainJoinConnector) SetVersion(value *string)() {
    m.version = value
}
