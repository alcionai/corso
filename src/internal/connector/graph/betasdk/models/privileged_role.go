package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRole 
type PrivilegedRole struct {
    Entity
    // The assignments for this role. Read-only. Nullable.
    assignments []PrivilegedRoleAssignmentable
    // Role name.
    name *string
    // The settings for this role. Read-only. Nullable.
    settings PrivilegedRoleSettingsable
    // The summary information for this role. Read-only. Nullable.
    summary PrivilegedRoleSummaryable
}
// NewPrivilegedRole instantiates a new privilegedRole and sets the default values.
func NewPrivilegedRole()(*PrivilegedRole) {
    m := &PrivilegedRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRole(), nil
}
// GetAssignments gets the assignments property value. The assignments for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetAssignments()([]PrivilegedRoleAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedRoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedRoleAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
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
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(PrivilegedRoleSettingsable))
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val.(PrivilegedRoleSummaryable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Role name.
func (m *PrivilegedRole) GetName()(*string) {
    return m.name
}
// GetSettings gets the settings property value. The settings for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetSettings()(PrivilegedRoleSettingsable) {
    return m.settings
}
// GetSummary gets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) GetSummary()(PrivilegedRoleSummaryable) {
    return m.summary
}
// Serialize serializes information the current object
func (m *PrivilegedRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The assignments for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetAssignments(value []PrivilegedRoleAssignmentable)() {
    m.assignments = value
}
// SetName sets the name property value. Role name.
func (m *PrivilegedRole) SetName(value *string)() {
    m.name = value
}
// SetSettings sets the settings property value. The settings for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetSettings(value PrivilegedRoleSettingsable)() {
    m.settings = value
}
// SetSummary sets the summary property value. The summary information for this role. Read-only. Nullable.
func (m *PrivilegedRole) SetSummary(value PrivilegedRoleSummaryable)() {
    m.summary = value
}
