package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleAssignment 
type RoleAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type of the admin relationship(s) associated with the role assignment. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges.
    assignmentType *DelegatedPrivilegeStatus
    // The OdataType property
    odataType *string
    // The collection of roles assigned.
    roles []RoleDefinitionable
}
// NewRoleAssignment instantiates a new roleAssignment and sets the default values.
func NewRoleAssignment()(*RoleAssignment) {
    m := &RoleAssignment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleAssignment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleAssignment) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentType gets the assignmentType property value. The type of the admin relationship(s) associated with the role assignment. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges.
func (m *RoleAssignment) GetAssignmentType()(*DelegatedPrivilegeStatus) {
    return m.assignmentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedPrivilegeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentType(val.(*DelegatedPrivilegeStatus))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(RoleDefinitionable)
            }
            m.SetRoles(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RoleAssignment) GetOdataType()(*string) {
    return m.odataType
}
// GetRoles gets the roles property value. The collection of roles assigned.
func (m *RoleAssignment) GetRoles()([]RoleDefinitionable) {
    return m.roles
}
// Serialize serializes information the current object
func (m *RoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentType() != nil {
        cast := (*m.GetAssignmentType()).String()
        err := writer.WriteStringValue("assignmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoles()))
        for i, v := range m.GetRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("roles", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleAssignment) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentType sets the assignmentType property value. The type of the admin relationship(s) associated with the role assignment. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges.
func (m *RoleAssignment) SetAssignmentType(value *DelegatedPrivilegeStatus)() {
    m.assignmentType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RoleAssignment) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoles sets the roles property value. The collection of roles assigned.
func (m *RoleAssignment) SetRoles(value []RoleDefinitionable)() {
    m.roles = value
}
