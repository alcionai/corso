package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentResourceRole provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AccessPackageAssignmentResourceRole struct {
    Entity
    // The access package assignments resulting in this role assignment. Read-only. Nullable.
    accessPackageAssignments []AccessPackageAssignmentable
    // The accessPackageResourceRole property
    accessPackageResourceRole AccessPackageResourceRoleable
    // The accessPackageResourceScope property
    accessPackageResourceScope AccessPackageResourceScopeable
    // Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
    accessPackageSubject AccessPackageSubjectable
    // A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
    originId *string
    // The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
    originSystem *string
    // The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
    status *string
}
// NewAccessPackageAssignmentResourceRole instantiates a new accessPackageAssignmentResourceRole and sets the default values.
func NewAccessPackageAssignmentResourceRole()(*AccessPackageAssignmentResourceRole) {
    m := &AccessPackageAssignmentResourceRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentResourceRole(), nil
}
// GetAccessPackageAssignments gets the accessPackageAssignments property value. The access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageAssignments()([]AccessPackageAssignmentable) {
    return m.accessPackageAssignments
}
// GetAccessPackageResourceRole gets the accessPackageResourceRole property value. The accessPackageResourceRole property
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceRole()(AccessPackageResourceRoleable) {
    return m.accessPackageResourceRole
}
// GetAccessPackageResourceScope gets the accessPackageResourceScope property value. The accessPackageResourceScope property
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceScope()(AccessPackageResourceScopeable) {
    return m.accessPackageResourceScope
}
// GetAccessPackageSubject gets the accessPackageSubject property value. Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageSubject()(AccessPackageSubjectable) {
    return m.accessPackageSubject
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentResourceRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentable)
            }
            m.SetAccessPackageAssignments(res)
        }
        return nil
    }
    res["accessPackageResourceRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResourceRole(val.(AccessPackageResourceRoleable))
        }
        return nil
    }
    res["accessPackageResourceScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResourceScope(val.(AccessPackageResourceScopeable))
        }
        return nil
    }
    res["accessPackageSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageSubject(val.(AccessPackageSubjectable))
        }
        return nil
    }
    res["originId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginId(val)
        }
        return nil
    }
    res["originSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginSystem(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetOriginId gets the originId property value. A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) GetOriginId()(*string) {
    return m.originId
}
// GetOriginSystem gets the originSystem property value. The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) GetOriginSystem()(*string) {
    return m.originSystem
}
// GetStatus gets the status property value. The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
func (m *AccessPackageAssignmentResourceRole) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentResourceRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignments()))
        for i, v := range m.GetAccessPackageAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageResourceRole", m.GetAccessPackageResourceRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageResourceScope", m.GetAccessPackageResourceScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageSubject", m.GetAccessPackageSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originId", m.GetOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originSystem", m.GetOriginSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageAssignments sets the accessPackageAssignments property value. The access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageAssignments(value []AccessPackageAssignmentable)() {
    m.accessPackageAssignments = value
}
// SetAccessPackageResourceRole sets the accessPackageResourceRole property value. The accessPackageResourceRole property
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceRole(value AccessPackageResourceRoleable)() {
    m.accessPackageResourceRole = value
}
// SetAccessPackageResourceScope sets the accessPackageResourceScope property value. The accessPackageResourceScope property
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceScope(value AccessPackageResourceScopeable)() {
    m.accessPackageResourceScope = value
}
// SetAccessPackageSubject sets the accessPackageSubject property value. Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageSubject(value AccessPackageSubjectable)() {
    m.accessPackageSubject = value
}
// SetOriginId sets the originId property value. A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) SetOriginId(value *string)() {
    m.originId = value
}
// SetOriginSystem sets the originSystem property value. The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) SetOriginSystem(value *string)() {
    m.originSystem = value
}
// SetStatus sets the status property value. The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
func (m *AccessPackageAssignmentResourceRole) SetStatus(value *string)() {
    m.status = value
}
