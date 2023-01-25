package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRbacResourceScope 
type UnifiedRbacResourceScope struct {
    Entity
    // The displayName property
    displayName *string
    // The scope property
    scope *string
    // The type property
    type_escaped *string
}
// NewUnifiedRbacResourceScope instantiates a new unifiedRbacResourceScope and sets the default values.
func NewUnifiedRbacResourceScope()(*UnifiedRbacResourceScope) {
    m := &UnifiedRbacResourceScope{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRbacResourceScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRbacResourceScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRbacResourceScope(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *UnifiedRbacResourceScope) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRbacResourceScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetScope gets the scope property value. The scope property
func (m *UnifiedRbacResourceScope) GetScope()(*string) {
    return m.scope
}
// GetType gets the type property value. The type property
func (m *UnifiedRbacResourceScope) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *UnifiedRbacResourceScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *UnifiedRbacResourceScope) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetScope sets the scope property value. The scope property
func (m *UnifiedRbacResourceScope) SetScope(value *string)() {
    m.scope = value
}
// SetType sets the type property value. The type property
func (m *UnifiedRbacResourceScope) SetType(value *string)() {
    m.type_escaped = value
}
