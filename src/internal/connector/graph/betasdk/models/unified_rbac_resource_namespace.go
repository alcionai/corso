package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRbacResourceNamespace provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UnifiedRbacResourceNamespace struct {
    Entity
    // Name of the resource namespace. Typically, the same name as the id property, such as microsoft.aad.b2c. Required. Supports $filter (eq, startsWith).
    name *string
    // Operations that an authorized principal are allowed to perform.
    resourceActions []UnifiedRbacResourceActionable
}
// NewUnifiedRbacResourceNamespace instantiates a new unifiedRbacResourceNamespace and sets the default values.
func NewUnifiedRbacResourceNamespace()(*UnifiedRbacResourceNamespace) {
    m := &UnifiedRbacResourceNamespace{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRbacResourceNamespace(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacResourceNamespace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["resourceActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRbacResourceActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRbacResourceActionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRbacResourceActionable)
            }
            m.SetResourceActions(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the resource namespace. Typically, the same name as the id property, such as microsoft.aad.b2c. Required. Supports $filter (eq, startsWith).
func (m *UnifiedRbacResourceNamespace) GetName()(*string) {
    return m.name
}
// GetResourceActions gets the resourceActions property value. Operations that an authorized principal are allowed to perform.
func (m *UnifiedRbacResourceNamespace) GetResourceActions()([]UnifiedRbacResourceActionable) {
    return m.resourceActions
}
// Serialize serializes information the current object
func (m *UnifiedRbacResourceNamespace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetResourceActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceActions()))
        for i, v := range m.GetResourceActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name of the resource namespace. Typically, the same name as the id property, such as microsoft.aad.b2c. Required. Supports $filter (eq, startsWith).
func (m *UnifiedRbacResourceNamespace) SetName(value *string)() {
    m.name = value
}
// SetResourceActions sets the resourceActions property value. Operations that an authorized principal are allowed to perform.
func (m *UnifiedRbacResourceNamespace) SetResourceActions(value []UnifiedRbacResourceActionable)() {
    m.resourceActions = value
}
