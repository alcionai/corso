package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Connector provides operations to manage the collection of site entities.
type Connector struct {
    Entity
    // The external IP address as detected by the the connector server. Read-only.
    externalIp *string
    // The machine name the connector is installed and running on.
    machineName *string
    // The connectorGroup that the connector is a member of. Read-only.
    memberOf []ConnectorGroupable
    // The status property
    status *ConnectorStatus
}
// NewConnector instantiates a new connector and sets the default values.
func NewConnector()(*Connector) {
    m := &Connector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnector(), nil
}
// GetExternalIp gets the externalIp property value. The external IP address as detected by the the connector server. Read-only.
func (m *Connector) GetExternalIp()(*string) {
    return m.externalIp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Connector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["externalIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalIp(val)
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
    res["memberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectorGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectorGroupable, len(val))
            for i, v := range val {
                res[i] = v.(ConnectorGroupable)
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ConnectorStatus))
        }
        return nil
    }
    return res
}
// GetMachineName gets the machineName property value. The machine name the connector is installed and running on.
func (m *Connector) GetMachineName()(*string) {
    return m.machineName
}
// GetMemberOf gets the memberOf property value. The connectorGroup that the connector is a member of. Read-only.
func (m *Connector) GetMemberOf()([]ConnectorGroupable) {
    return m.memberOf
}
// GetStatus gets the status property value. The status property
func (m *Connector) GetStatus()(*ConnectorStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *Connector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("externalIp", m.GetExternalIp())
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
    if m.GetMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExternalIp sets the externalIp property value. The external IP address as detected by the the connector server. Read-only.
func (m *Connector) SetExternalIp(value *string)() {
    m.externalIp = value
}
// SetMachineName sets the machineName property value. The machine name the connector is installed and running on.
func (m *Connector) SetMachineName(value *string)() {
    m.machineName = value
}
// SetMemberOf sets the memberOf property value. The connectorGroup that the connector is a member of. Read-only.
func (m *Connector) SetMemberOf(value []ConnectorGroupable)() {
    m.memberOf = value
}
// SetStatus sets the status property value. The status property
func (m *Connector) SetStatus(value *ConnectorStatus)() {
    m.status = value
}
