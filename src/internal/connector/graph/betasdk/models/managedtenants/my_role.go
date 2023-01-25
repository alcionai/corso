package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MyRole 
type MyRole struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of role assignments for the managed tenant.
    assignments []RoleAssignmentable
    // The OdataType property
    odataType *string
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
}
// NewMyRole instantiates a new myRole and sets the default values.
func NewMyRole()(*MyRole) {
    m := &MyRole{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMyRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMyRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMyRole(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MyRole) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. A collection of role assignments for the managed tenant.
func (m *MyRole) GetAssignments()([]RoleAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MyRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentable)
            }
            m.SetAssignments(res)
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
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MyRole) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *MyRole) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *MyRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignments", cast)
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
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *MyRole) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. A collection of role assignments for the managed tenant.
func (m *MyRole) SetAssignments(value []RoleAssignmentable)() {
    m.assignments = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MyRole) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *MyRole) SetTenantId(value *string)() {
    m.tenantId = value
}
