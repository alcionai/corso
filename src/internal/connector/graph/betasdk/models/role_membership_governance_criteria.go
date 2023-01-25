package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleMembershipGovernanceCriteria 
type RoleMembershipGovernanceCriteria struct {
    GovernanceCriteria
    // The roleId property
    roleId *string
    // The roleTemplateId property
    roleTemplateId *string
}
// NewRoleMembershipGovernanceCriteria instantiates a new RoleMembershipGovernanceCriteria and sets the default values.
func NewRoleMembershipGovernanceCriteria()(*RoleMembershipGovernanceCriteria) {
    m := &RoleMembershipGovernanceCriteria{
        GovernanceCriteria: *NewGovernanceCriteria(),
    }
    odataTypeValue := "#microsoft.graph.roleMembershipGovernanceCriteria";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRoleMembershipGovernanceCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleMembershipGovernanceCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleMembershipGovernanceCriteria(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleMembershipGovernanceCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GovernanceCriteria.GetFieldDeserializers()
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleTemplateId(val)
        }
        return nil
    }
    return res
}
// GetRoleId gets the roleId property value. The roleId property
func (m *RoleMembershipGovernanceCriteria) GetRoleId()(*string) {
    return m.roleId
}
// GetRoleTemplateId gets the roleTemplateId property value. The roleTemplateId property
func (m *RoleMembershipGovernanceCriteria) GetRoleTemplateId()(*string) {
    return m.roleTemplateId
}
// Serialize serializes information the current object
func (m *RoleMembershipGovernanceCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GovernanceCriteria.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleTemplateId", m.GetRoleTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRoleId sets the roleId property value. The roleId property
func (m *RoleMembershipGovernanceCriteria) SetRoleId(value *string)() {
    m.roleId = value
}
// SetRoleTemplateId sets the roleTemplateId property value. The roleTemplateId property
func (m *RoleMembershipGovernanceCriteria) SetRoleTemplateId(value *string)() {
    m.roleTemplateId = value
}
