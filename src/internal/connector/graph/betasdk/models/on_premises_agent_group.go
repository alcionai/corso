package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesAgentGroup provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type OnPremisesAgentGroup struct {
    Entity
    // List of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
    agents []OnPremisesAgentable
    // Display name of the onPremisesAgentGroup.
    displayName *string
    // Indicates if the onPremisesAgentGroup is the default agent group. Only a single agent group can be the default onPremisesAgentGroup and is set by the system.
    isDefault *bool
    // List of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
    publishedResources []PublishedResourceable
    // The publishingType property
    publishingType *OnPremisesPublishingType
}
// NewOnPremisesAgentGroup instantiates a new onPremisesAgentGroup and sets the default values.
func NewOnPremisesAgentGroup()(*OnPremisesAgentGroup) {
    m := &OnPremisesAgentGroup{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnPremisesAgentGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesAgentGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesAgentGroup(), nil
}
// GetAgents gets the agents property value. List of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) GetAgents()([]OnPremisesAgentable) {
    return m.agents
}
// GetDisplayName gets the displayName property value. Display name of the onPremisesAgentGroup.
func (m *OnPremisesAgentGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesAgentGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesAgentable)
            }
            m.SetAgents(res)
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
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["publishedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublishedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublishedResourceable, len(val))
            for i, v := range val {
                res[i] = v.(PublishedResourceable)
            }
            m.SetPublishedResources(res)
        }
        return nil
    }
    res["publishingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnPremisesPublishingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingType(val.(*OnPremisesPublishingType))
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates if the onPremisesAgentGroup is the default agent group. Only a single agent group can be the default onPremisesAgentGroup and is set by the system.
func (m *OnPremisesAgentGroup) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetPublishedResources gets the publishedResources property value. List of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) GetPublishedResources()([]PublishedResourceable) {
    return m.publishedResources
}
// GetPublishingType gets the publishingType property value. The publishingType property
func (m *OnPremisesAgentGroup) GetPublishingType()(*OnPremisesPublishingType) {
    return m.publishingType
}
// Serialize serializes information the current object
func (m *OnPremisesAgentGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgents()))
        for i, v := range m.GetAgents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agents", cast)
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
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    if m.GetPublishedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPublishedResources()))
        for i, v := range m.GetPublishedResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("publishedResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublishingType() != nil {
        cast := (*m.GetPublishingType()).String()
        err = writer.WriteStringValue("publishingType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgents sets the agents property value. List of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) SetAgents(value []OnPremisesAgentable)() {
    m.agents = value
}
// SetDisplayName sets the displayName property value. Display name of the onPremisesAgentGroup.
func (m *OnPremisesAgentGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsDefault sets the isDefault property value. Indicates if the onPremisesAgentGroup is the default agent group. Only a single agent group can be the default onPremisesAgentGroup and is set by the system.
func (m *OnPremisesAgentGroup) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetPublishedResources sets the publishedResources property value. List of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) SetPublishedResources(value []PublishedResourceable)() {
    m.publishedResources = value
}
// SetPublishingType sets the publishingType property value. The publishingType property
func (m *OnPremisesAgentGroup) SetPublishingType(value *OnPremisesPublishingType)() {
    m.publishingType = value
}
