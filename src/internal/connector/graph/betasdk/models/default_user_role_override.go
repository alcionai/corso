package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultUserRoleOverride provides operations to manage the collection of site entities.
type DefaultUserRoleOverride struct {
    Entity
    // The isDefault property
    isDefault *bool
    // The rolePermissions property
    rolePermissions []UnifiedRolePermissionable
}
// NewDefaultUserRoleOverride instantiates a new defaultUserRoleOverride and sets the default values.
func NewDefaultUserRoleOverride()(*DefaultUserRoleOverride) {
    m := &DefaultUserRoleOverride{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDefaultUserRoleOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultUserRoleOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultUserRoleOverride(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultUserRoleOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["rolePermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRolePermissionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRolePermissionable)
            }
            m.SetRolePermissions(res)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. The isDefault property
func (m *DefaultUserRoleOverride) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetRolePermissions gets the rolePermissions property value. The rolePermissions property
func (m *DefaultUserRoleOverride) GetRolePermissions()([]UnifiedRolePermissionable) {
    return m.rolePermissions
}
// Serialize serializes information the current object
func (m *DefaultUserRoleOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    if m.GetRolePermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *DefaultUserRoleOverride) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetRolePermissions sets the rolePermissions property value. The rolePermissions property
func (m *DefaultUserRoleOverride) SetRolePermissions(value []UnifiedRolePermissionable)() {
    m.rolePermissions = value
}
