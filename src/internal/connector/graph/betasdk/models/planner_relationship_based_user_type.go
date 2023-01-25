package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerRelationshipBasedUserType 
type PlannerRelationshipBasedUserType struct {
    PlannerTaskConfigurationRoleBase
    // The role property
    role *PlannerRelationshipUserRoles
}
// NewPlannerRelationshipBasedUserType instantiates a new PlannerRelationshipBasedUserType and sets the default values.
func NewPlannerRelationshipBasedUserType()(*PlannerRelationshipBasedUserType) {
    m := &PlannerRelationshipBasedUserType{
        PlannerTaskConfigurationRoleBase: *NewPlannerTaskConfigurationRoleBase(),
    }
    odataTypeValue := "#microsoft.graph.plannerRelationshipBasedUserType";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerRelationshipBasedUserTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerRelationshipBasedUserTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerRelationshipBasedUserType(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerRelationshipBasedUserType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerTaskConfigurationRoleBase.GetFieldDeserializers()
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerRelationshipUserRoles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*PlannerRelationshipUserRoles))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. The role property
func (m *PlannerRelationshipBasedUserType) GetRole()(*PlannerRelationshipUserRoles) {
    return m.role
}
// Serialize serializes information the current object
func (m *PlannerRelationshipBasedUserType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerTaskConfigurationRoleBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRole sets the role property value. The role property
func (m *PlannerRelationshipBasedUserType) SetRole(value *PlannerRelationshipUserRoles)() {
    m.role = value
}
