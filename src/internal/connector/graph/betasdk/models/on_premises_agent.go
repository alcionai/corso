package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesAgent 
type OnPremisesAgent struct {
    Entity
    // List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroupable
    // The external IP address as detected by the service for the agent machine. Read-only
    externalIp *string
    // The name of the machine that the aggent is running on. Read-only
    machineName *string
    // The status property
    status *AgentStatus
    // The supportedPublishingTypes property
    supportedPublishingTypes []OnPremisesPublishingType
}
// NewOnPremisesAgent instantiates a new OnPremisesAgent and sets the default values.
func NewOnPremisesAgent()(*OnPremisesAgent) {
    m := &OnPremisesAgent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnPremisesAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesAgent(), nil
}
// GetAgentGroups gets the agentGroups property value. List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
func (m *OnPremisesAgent) GetAgentGroups()([]OnPremisesAgentGroupable) {
    return m.agentGroups
}
// GetExternalIp gets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
func (m *OnPremisesAgent) GetExternalIp()(*string) {
    return m.externalIp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentGroupable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesAgentGroupable)
            }
            m.SetAgentGroups(res)
        }
        return nil
    }
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAgentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AgentStatus))
        }
        return nil
    }
    res["supportedPublishingTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOnPremisesPublishingType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesPublishingType, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesPublishingType))
            }
            m.SetSupportedPublishingTypes(res)
        }
        return nil
    }
    return res
}
// GetMachineName gets the machineName property value. The name of the machine that the aggent is running on. Read-only
func (m *OnPremisesAgent) GetMachineName()(*string) {
    return m.machineName
}
// GetStatus gets the status property value. The status property
func (m *OnPremisesAgent) GetStatus()(*AgentStatus) {
    return m.status
}
// GetSupportedPublishingTypes gets the supportedPublishingTypes property value. The supportedPublishingTypes property
func (m *OnPremisesAgent) GetSupportedPublishingTypes()([]OnPremisesPublishingType) {
    return m.supportedPublishingTypes
}
// Serialize serializes information the current object
func (m *OnPremisesAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgentGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgentGroups()))
        for i, v := range m.GetAgentGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agentGroups", cast)
        if err != nil {
            return err
        }
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedPublishingTypes() != nil {
        err = writer.WriteCollectionOfStringValues("supportedPublishingTypes", SerializeOnPremisesPublishingType(m.GetSupportedPublishingTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgentGroups sets the agentGroups property value. List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
func (m *OnPremisesAgent) SetAgentGroups(value []OnPremisesAgentGroupable)() {
    m.agentGroups = value
}
// SetExternalIp sets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
func (m *OnPremisesAgent) SetExternalIp(value *string)() {
    m.externalIp = value
}
// SetMachineName sets the machineName property value. The name of the machine that the aggent is running on. Read-only
func (m *OnPremisesAgent) SetMachineName(value *string)() {
    m.machineName = value
}
// SetStatus sets the status property value. The status property
func (m *OnPremisesAgent) SetStatus(value *AgentStatus)() {
    m.status = value
}
// SetSupportedPublishingTypes sets the supportedPublishingTypes property value. The supportedPublishingTypes property
func (m *OnPremisesAgent) SetSupportedPublishingTypes(value []OnPremisesPublishingType)() {
    m.supportedPublishingTypes = value
}
