package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RbacApplicationMultiple 
type RbacApplicationMultiple struct {
    Entity
    // The resourceNamespaces property
    resourceNamespaces []UnifiedRbacResourceNamespaceable
    // The roleAssignments property
    roleAssignments []UnifiedRoleAssignmentMultipleable
    // The roleDefinitions property
    roleDefinitions []UnifiedRoleDefinitionable
}
// NewRbacApplicationMultiple instantiates a new RbacApplicationMultiple and sets the default values.
func NewRbacApplicationMultiple()(*RbacApplicationMultiple) {
    m := &RbacApplicationMultiple{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRbacApplicationMultipleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRbacApplicationMultipleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRbacApplicationMultiple(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RbacApplicationMultiple) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["resourceNamespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRbacResourceNamespaceable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRbacResourceNamespaceable)
            }
            m.SetResourceNamespaces(res)
        }
        return nil
    }
    res["roleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentMultipleable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleAssignmentMultipleable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleDefinitionable)
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    return res
}
// GetResourceNamespaces gets the resourceNamespaces property value. The resourceNamespaces property
func (m *RbacApplicationMultiple) GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable) {
    return m.resourceNamespaces
}
// GetRoleAssignments gets the roleAssignments property value. The roleAssignments property
func (m *RbacApplicationMultiple) GetRoleAssignments()([]UnifiedRoleAssignmentMultipleable) {
    return m.roleAssignments
}
// GetRoleDefinitions gets the roleDefinitions property value. The roleDefinitions property
func (m *RbacApplicationMultiple) GetRoleDefinitions()([]UnifiedRoleDefinitionable) {
    return m.roleDefinitions
}
// Serialize serializes information the current object
func (m *RbacApplicationMultiple) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetResourceNamespaces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceNamespaces()))
        for i, v := range m.GetResourceNamespaces() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceNamespaces", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceNamespaces sets the resourceNamespaces property value. The resourceNamespaces property
func (m *RbacApplicationMultiple) SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)() {
    m.resourceNamespaces = value
}
// SetRoleAssignments sets the roleAssignments property value. The roleAssignments property
func (m *RbacApplicationMultiple) SetRoleAssignments(value []UnifiedRoleAssignmentMultipleable)() {
    m.roleAssignments = value
}
// SetRoleDefinitions sets the roleDefinitions property value. The roleDefinitions property
func (m *RbacApplicationMultiple) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    m.roleDefinitions = value
}
