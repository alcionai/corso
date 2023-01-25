package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccess 
type PrivilegedAccess struct {
    Entity
    // The display name of the provider managed by PIM.
    displayName *string
    // A collection of resources for the provider.
    resources []GovernanceResourceable
    // A collection of role assignment requests for the provider.
    roleAssignmentRequests []GovernanceRoleAssignmentRequestable
    // A collection of role assignments for the provider.
    roleAssignments []GovernanceRoleAssignmentable
    // A collection of role defintions for the provider.
    roleDefinitions []GovernanceRoleDefinitionable
    // A collection of role settings for the provider.
    roleSettings []GovernanceRoleSettingable
}
// NewPrivilegedAccess instantiates a new PrivilegedAccess and sets the default values.
func NewPrivilegedAccess()(*PrivilegedAccess) {
    m := &PrivilegedAccess{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccess(), nil
}
// GetDisplayName gets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceResourceable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    res["roleAssignmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleAssignmentRequestable)
            }
            m.SetRoleAssignmentRequests(res)
        }
        return nil
    }
    res["roleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleAssignmentable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleDefinitionable)
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRoleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRoleSettingable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRoleSettingable)
            }
            m.SetRoleSettings(res)
        }
        return nil
    }
    return res
}
// GetResources gets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) GetResources()([]GovernanceResourceable) {
    return m.resources
}
// GetRoleAssignmentRequests gets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequestable) {
    return m.roleAssignmentRequests
}
// GetRoleAssignments gets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) GetRoleAssignments()([]GovernanceRoleAssignmentable) {
    return m.roleAssignments
}
// GetRoleDefinitions gets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) GetRoleDefinitions()([]GovernanceRoleDefinitionable) {
    return m.roleDefinitions
}
// GetRoleSettings gets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) GetRoleSettings()([]GovernanceRoleSettingable) {
    return m.roleSettings
}
// Serialize serializes information the current object
func (m *PrivilegedAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignmentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignmentRequests()))
        for i, v := range m.GetRoleAssignmentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignmentRequests", cast)
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
    if m.GetRoleSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleSettings()))
        for i, v := range m.GetRoleSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the provider managed by PIM.
func (m *PrivilegedAccess) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetResources sets the resources property value. A collection of resources for the provider.
func (m *PrivilegedAccess) SetResources(value []GovernanceResourceable)() {
    m.resources = value
}
// SetRoleAssignmentRequests sets the roleAssignmentRequests property value. A collection of role assignment requests for the provider.
func (m *PrivilegedAccess) SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequestable)() {
    m.roleAssignmentRequests = value
}
// SetRoleAssignments sets the roleAssignments property value. A collection of role assignments for the provider.
func (m *PrivilegedAccess) SetRoleAssignments(value []GovernanceRoleAssignmentable)() {
    m.roleAssignments = value
}
// SetRoleDefinitions sets the roleDefinitions property value. A collection of role defintions for the provider.
func (m *PrivilegedAccess) SetRoleDefinitions(value []GovernanceRoleDefinitionable)() {
    m.roleDefinitions = value
}
// SetRoleSettings sets the roleSettings property value. A collection of role settings for the provider.
func (m *PrivilegedAccess) SetRoleSettings(value []GovernanceRoleSettingable)() {
    m.roleSettings = value
}
