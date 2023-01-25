package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelServer entity that represents a single Microsoft Tunnel server
type MicrosoftTunnelServer struct {
    Entity
    // The digest of the current agent image running on this server
    agentImageDigest *string
    // The display name for the server. This property is required when a server is created and cannot be cleared during updates.
    displayName *string
    // Indicates when the server last checked in
    lastCheckinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The digest of the current server image running on this server
    serverImageDigest *string
    // Enum of possible MicrosoftTunnelServer health status types
    tunnelServerHealthStatus *MicrosoftTunnelServerHealthStatus
}
// NewMicrosoftTunnelServer instantiates a new microsoftTunnelServer and sets the default values.
func NewMicrosoftTunnelServer()(*MicrosoftTunnelServer) {
    m := &MicrosoftTunnelServer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftTunnelServerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelServerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelServer(), nil
}
// GetAgentImageDigest gets the agentImageDigest property value. The digest of the current agent image running on this server
func (m *MicrosoftTunnelServer) GetAgentImageDigest()(*string) {
    return m.agentImageDigest
}
// GetDisplayName gets the displayName property value. The display name for the server. This property is required when a server is created and cannot be cleared during updates.
func (m *MicrosoftTunnelServer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelServer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentImageDigest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgentImageDigest(val)
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
    res["lastCheckinDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckinDateTime(val)
        }
        return nil
    }
    res["serverImageDigest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerImageDigest(val)
        }
        return nil
    }
    res["tunnelServerHealthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftTunnelServerHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTunnelServerHealthStatus(val.(*MicrosoftTunnelServerHealthStatus))
        }
        return nil
    }
    return res
}
// GetLastCheckinDateTime gets the lastCheckinDateTime property value. Indicates when the server last checked in
func (m *MicrosoftTunnelServer) GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCheckinDateTime
}
// GetServerImageDigest gets the serverImageDigest property value. The digest of the current server image running on this server
func (m *MicrosoftTunnelServer) GetServerImageDigest()(*string) {
    return m.serverImageDigest
}
// GetTunnelServerHealthStatus gets the tunnelServerHealthStatus property value. Enum of possible MicrosoftTunnelServer health status types
func (m *MicrosoftTunnelServer) GetTunnelServerHealthStatus()(*MicrosoftTunnelServerHealthStatus) {
    return m.tunnelServerHealthStatus
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelServer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("agentImageDigest", m.GetAgentImageDigest())
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
        err = writer.WriteTimeValue("lastCheckinDateTime", m.GetLastCheckinDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverImageDigest", m.GetServerImageDigest())
        if err != nil {
            return err
        }
    }
    if m.GetTunnelServerHealthStatus() != nil {
        cast := (*m.GetTunnelServerHealthStatus()).String()
        err = writer.WriteStringValue("tunnelServerHealthStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgentImageDigest sets the agentImageDigest property value. The digest of the current agent image running on this server
func (m *MicrosoftTunnelServer) SetAgentImageDigest(value *string)() {
    m.agentImageDigest = value
}
// SetDisplayName sets the displayName property value. The display name for the server. This property is required when a server is created and cannot be cleared during updates.
func (m *MicrosoftTunnelServer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastCheckinDateTime sets the lastCheckinDateTime property value. Indicates when the server last checked in
func (m *MicrosoftTunnelServer) SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckinDateTime = value
}
// SetServerImageDigest sets the serverImageDigest property value. The digest of the current server image running on this server
func (m *MicrosoftTunnelServer) SetServerImageDigest(value *string)() {
    m.serverImageDigest = value
}
// SetTunnelServerHealthStatus sets the tunnelServerHealthStatus property value. Enum of possible MicrosoftTunnelServer health status types
func (m *MicrosoftTunnelServer) SetTunnelServerHealthStatus(value *MicrosoftTunnelServerHealthStatus)() {
    m.tunnelServerHealthStatus = value
}
