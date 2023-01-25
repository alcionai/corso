package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NdesConnector entity which represents an OnPrem Ndes connector.
type NdesConnector struct {
    Entity
    // The build version of the Ndes Connector.
    connectorVersion *string
    // The friendly name of the Ndes Connector.
    displayName *string
    // Timestamp when on-prem certificate connector was enrolled in Intune.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last connection time for the Ndes Connector
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the machine running on-prem certificate connector service.
    machineName *string
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // The current status of the Ndes Connector.
    state *NdesConnectorState
}
// NewNdesConnector instantiates a new ndesConnector and sets the default values.
func NewNdesConnector()(*NdesConnector) {
    m := &NdesConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateNdesConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNdesConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNdesConnector(), nil
}
// GetConnectorVersion gets the connectorVersion property value. The build version of the Ndes Connector.
func (m *NdesConnector) GetConnectorVersion()(*string) {
    return m.connectorVersion
}
// GetDisplayName gets the displayName property value. The friendly name of the Ndes Connector.
func (m *NdesConnector) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Timestamp when on-prem certificate connector was enrolled in Intune.
func (m *NdesConnector) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrolledDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NdesConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorVersion(val)
        }
        return nil
    }
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
    res["enrolledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDateTime(val)
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
    res["machineName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineName(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNdesConnectorState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*NdesConnectorState))
        }
        return nil
    }
    return res
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
func (m *NdesConnector) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastConnectionDateTime
}
// GetMachineName gets the machineName property value. Name of the machine running on-prem certificate connector service.
func (m *NdesConnector) GetMachineName()(*string) {
    return m.machineName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *NdesConnector) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetState gets the state property value. The current status of the Ndes Connector.
func (m *NdesConnector) GetState()(*NdesConnectorState) {
    return m.state
}
// Serialize serializes information the current object
func (m *NdesConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("connectorVersion", m.GetConnectorVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrolledDateTime", m.GetEnrolledDateTime())
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
    {
        err = writer.WriteStringValue("machineName", m.GetMachineName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
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
    return nil
}
// SetConnectorVersion sets the connectorVersion property value. The build version of the Ndes Connector.
func (m *NdesConnector) SetConnectorVersion(value *string)() {
    m.connectorVersion = value
}
// SetDisplayName sets the displayName property value. The friendly name of the Ndes Connector.
func (m *NdesConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Timestamp when on-prem certificate connector was enrolled in Intune.
func (m *NdesConnector) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Last connection time for the Ndes Connector
func (m *NdesConnector) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// SetMachineName sets the machineName property value. Name of the machine running on-prem certificate connector service.
func (m *NdesConnector) SetMachineName(value *string)() {
    m.machineName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *NdesConnector) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetState sets the state property value. The current status of the Ndes Connector.
func (m *NdesConnector) SetState(value *NdesConnectorState)() {
    m.state = value
}
