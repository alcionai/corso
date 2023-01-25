package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleScopeTag role Scope Tag
type RoleScopeTag struct {
    Entity
    // The list of assignments for this Role Scope Tag.
    assignments []RoleScopeTagAutoAssignmentable
    // Description of the Role Scope Tag.
    description *string
    // The display or friendly name of the Role Scope Tag.
    displayName *string
    // Description of the Role Scope Tag. This property is read-only.
    isBuiltIn *bool
}
// NewRoleScopeTag instantiates a new roleScopeTag and sets the default values.
func NewRoleScopeTag()(*RoleScopeTag) {
    m := &RoleScopeTag{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRoleScopeTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleScopeTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleScopeTag(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for this Role Scope Tag.
func (m *RoleScopeTag) GetAssignments()([]RoleScopeTagAutoAssignmentable) {
    return m.assignments
}
// GetDescription gets the description property value. Description of the Role Scope Tag.
func (m *RoleScopeTag) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display or friendly name of the Role Scope Tag.
func (m *RoleScopeTag) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleScopeTag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleScopeTagAutoAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleScopeTagAutoAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(RoleScopeTagAutoAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["isBuiltIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
func (m *RoleScopeTag) GetIsBuiltIn()(*bool) {
    return m.isBuiltIn
}
// Serialize serializes information the current object
func (m *RoleScopeTag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for this Role Scope Tag.
func (m *RoleScopeTag) SetAssignments(value []RoleScopeTagAutoAssignmentable)() {
    m.assignments = value
}
// SetDescription sets the description property value. Description of the Role Scope Tag.
func (m *RoleScopeTag) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display or friendly name of the Role Scope Tag.
func (m *RoleScopeTag) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsBuiltIn sets the isBuiltIn property value. Description of the Role Scope Tag. This property is read-only.
func (m *RoleScopeTag) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
