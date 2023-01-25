package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublishedResource provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PublishedResource struct {
    Entity
    // List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroupable
    // Display Name of the publishedResource.
    displayName *string
    // The publishingType property
    publishingType *OnPremisesPublishingType
    // Name of the publishedResource.
    resourceName *string
}
// NewPublishedResource instantiates a new publishedResource and sets the default values.
func NewPublishedResource()(*PublishedResource) {
    m := &PublishedResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePublishedResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublishedResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublishedResource(), nil
}
// GetAgentGroups gets the agentGroups property value. List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *PublishedResource) GetAgentGroups()([]OnPremisesAgentGroupable) {
    return m.agentGroups
}
// GetDisplayName gets the displayName property value. Display Name of the publishedResource.
func (m *PublishedResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublishedResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    return res
}
// GetPublishingType gets the publishingType property value. The publishingType property
func (m *PublishedResource) GetPublishingType()(*OnPremisesPublishingType) {
    return m.publishingType
}
// GetResourceName gets the resourceName property value. Name of the publishedResource.
func (m *PublishedResource) GetResourceName()(*string) {
    return m.resourceName
}
// Serialize serializes information the current object
func (m *PublishedResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgentGroups sets the agentGroups property value. List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *PublishedResource) SetAgentGroups(value []OnPremisesAgentGroupable)() {
    m.agentGroups = value
}
// SetDisplayName sets the displayName property value. Display Name of the publishedResource.
func (m *PublishedResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPublishingType sets the publishingType property value. The publishingType property
func (m *PublishedResource) SetPublishingType(value *OnPremisesPublishingType)() {
    m.publishingType = value
}
// SetResourceName sets the resourceName property value. Name of the publishedResource.
func (m *PublishedResource) SetResourceName(value *string)() {
    m.resourceName = value
}
