package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagement 
type EntitlementManagement struct {
    Entity
    // The accessPackageAssignmentApprovals property
    accessPackageAssignmentApprovals []Approvalable
    // Represents the policy that governs which subjects can request or be assigned an access package via an access package assignment.
    accessPackageAssignmentPolicies []AccessPackageAssignmentPolicyable
    // Represents access package assignment requests created by or on behalf of a user.
    accessPackageAssignmentRequests []AccessPackageAssignmentRequestable
    // Represents the resource-specific role which a subject has been assigned through an access package assignment.
    accessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRoleable
    // The assignment of an access package to a subject for a period of time.
    accessPackageAssignments []AccessPackageAssignmentable
    // A container of access packages.
    accessPackageCatalogs []AccessPackageCatalogable
    // A reference to the geolocation environment in which a resource is located.
    accessPackageResourceEnvironments []AccessPackageResourceEnvironmentable
    // Represents a request to add or remove a resource to or from a catalog respectively.
    accessPackageResourceRequests []AccessPackageResourceRequestable
    // A reference to both a scope within a resource, and a role in that resource for that scope.
    accessPackageResourceRoleScopes []AccessPackageResourceRoleScopeable
    // A reference to a resource associated with an access package catalog.
    accessPackageResources []AccessPackageResourceable
    // Represents access package objects.
    accessPackages []AccessPackageable
    // Represents references to a directory or domain of another organization whose users can request access.
    connectedOrganizations []ConnectedOrganizationable
    // Represents the settings that control the behavior of Azure AD entitlement management.
    settings EntitlementManagementSettingsable
    // The subjects property
    subjects []AccessPackageSubjectable
}
// NewEntitlementManagement instantiates a new EntitlementManagement and sets the default values.
func NewEntitlementManagement()(*EntitlementManagement) {
    m := &EntitlementManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEntitlementManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagement(), nil
}
// GetAccessPackageAssignmentApprovals gets the accessPackageAssignmentApprovals property value. The accessPackageAssignmentApprovals property
func (m *EntitlementManagement) GetAccessPackageAssignmentApprovals()([]Approvalable) {
    return m.accessPackageAssignmentApprovals
}
// GetAccessPackageAssignmentPolicies gets the accessPackageAssignmentPolicies property value. Represents the policy that governs which subjects can request or be assigned an access package via an access package assignment.
func (m *EntitlementManagement) GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicyable) {
    return m.accessPackageAssignmentPolicies
}
// GetAccessPackageAssignmentRequests gets the accessPackageAssignmentRequests property value. Represents access package assignment requests created by or on behalf of a user.
func (m *EntitlementManagement) GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable) {
    return m.accessPackageAssignmentRequests
}
// GetAccessPackageAssignmentResourceRoles gets the accessPackageAssignmentResourceRoles property value. Represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *EntitlementManagement) GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable) {
    return m.accessPackageAssignmentResourceRoles
}
// GetAccessPackageAssignments gets the accessPackageAssignments property value. The assignment of an access package to a subject for a period of time.
func (m *EntitlementManagement) GetAccessPackageAssignments()([]AccessPackageAssignmentable) {
    return m.accessPackageAssignments
}
// GetAccessPackageCatalogs gets the accessPackageCatalogs property value. A container of access packages.
func (m *EntitlementManagement) GetAccessPackageCatalogs()([]AccessPackageCatalogable) {
    return m.accessPackageCatalogs
}
// GetAccessPackageResourceEnvironments gets the accessPackageResourceEnvironments property value. A reference to the geolocation environment in which a resource is located.
func (m *EntitlementManagement) GetAccessPackageResourceEnvironments()([]AccessPackageResourceEnvironmentable) {
    return m.accessPackageResourceEnvironments
}
// GetAccessPackageResourceRequests gets the accessPackageResourceRequests property value. Represents a request to add or remove a resource to or from a catalog respectively.
func (m *EntitlementManagement) GetAccessPackageResourceRequests()([]AccessPackageResourceRequestable) {
    return m.accessPackageResourceRequests
}
// GetAccessPackageResourceRoleScopes gets the accessPackageResourceRoleScopes property value. A reference to both a scope within a resource, and a role in that resource for that scope.
func (m *EntitlementManagement) GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScopeable) {
    return m.accessPackageResourceRoleScopes
}
// GetAccessPackageResources gets the accessPackageResources property value. A reference to a resource associated with an access package catalog.
func (m *EntitlementManagement) GetAccessPackageResources()([]AccessPackageResourceable) {
    return m.accessPackageResources
}
// GetAccessPackages gets the accessPackages property value. Represents access package objects.
func (m *EntitlementManagement) GetAccessPackages()([]AccessPackageable) {
    return m.accessPackages
}
// GetConnectedOrganizations gets the connectedOrganizations property value. Represents references to a directory or domain of another organization whose users can request access.
func (m *EntitlementManagement) GetConnectedOrganizations()([]ConnectedOrganizationable) {
    return m.connectedOrganizations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentApprovals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approvalable, len(val))
            for i, v := range val {
                res[i] = v.(Approvalable)
            }
            m.SetAccessPackageAssignmentApprovals(res)
        }
        return nil
    }
    res["accessPackageAssignmentPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentPolicyable)
            }
            m.SetAccessPackageAssignmentPolicies(res)
        }
        return nil
    }
    res["accessPackageAssignmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentRequestable)
            }
            m.SetAccessPackageAssignmentRequests(res)
        }
        return nil
    }
    res["accessPackageAssignmentResourceRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentResourceRoleable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentResourceRoleable)
            }
            m.SetAccessPackageAssignmentResourceRoles(res)
        }
        return nil
    }
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
    res["accessPackageCatalogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageCatalogable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageCatalogable)
            }
            m.SetAccessPackageCatalogs(res)
        }
        return nil
    }
    res["accessPackageResourceEnvironments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceEnvironmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceEnvironmentable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageResourceEnvironmentable)
            }
            m.SetAccessPackageResourceEnvironments(res)
        }
        return nil
    }
    res["accessPackageResourceRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRequestable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageResourceRequestable)
            }
            m.SetAccessPackageResourceRequests(res)
        }
        return nil
    }
    res["accessPackageResourceRoleScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRoleScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageResourceRoleScopeable)
            }
            m.SetAccessPackageResourceRoleScopes(res)
        }
        return nil
    }
    res["accessPackageResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageResourceable)
            }
            m.SetAccessPackageResources(res)
        }
        return nil
    }
    res["accessPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageable)
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["connectedOrganizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectedOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectedOrganizationable, len(val))
            for i, v := range val {
                res[i] = v.(ConnectedOrganizationable)
            }
            m.SetConnectedOrganizations(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(EntitlementManagementSettingsable))
        }
        return nil
    }
    res["subjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageSubjectable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageSubjectable)
            }
            m.SetSubjects(res)
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. Represents the settings that control the behavior of Azure AD entitlement management.
func (m *EntitlementManagement) GetSettings()(EntitlementManagementSettingsable) {
    return m.settings
}
// GetSubjects gets the subjects property value. The subjects property
func (m *EntitlementManagement) GetSubjects()([]AccessPackageSubjectable) {
    return m.subjects
}
// Serialize serializes information the current object
func (m *EntitlementManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignmentApprovals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentApprovals()))
        for i, v := range m.GetAccessPackageAssignmentApprovals() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentApprovals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageAssignmentPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentPolicies()))
        for i, v := range m.GetAccessPackageAssignmentPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageAssignmentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentRequests()))
        for i, v := range m.GetAccessPackageAssignmentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageAssignmentResourceRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignmentResourceRoles()))
        for i, v := range m.GetAccessPackageAssignmentResourceRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentResourceRoles", cast)
        if err != nil {
            return err
        }
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
    if m.GetAccessPackageCatalogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageCatalogs()))
        for i, v := range m.GetAccessPackageCatalogs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageCatalogs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceEnvironments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResourceEnvironments()))
        for i, v := range m.GetAccessPackageResourceEnvironments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceEnvironments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResourceRequests()))
        for i, v := range m.GetAccessPackageResourceRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResourceRoleScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResourceRoleScopes()))
        for i, v := range m.GetAccessPackageResourceRoleScopes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResourceRoleScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackageResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageResources()))
        for i, v := range m.GetAccessPackageResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectedOrganizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectedOrganizations()))
        for i, v := range m.GetConnectedOrganizations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("connectedOrganizations", cast)
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
    if m.GetSubjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjects()))
        for i, v := range m.GetSubjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("subjects", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageAssignmentApprovals sets the accessPackageAssignmentApprovals property value. The accessPackageAssignmentApprovals property
func (m *EntitlementManagement) SetAccessPackageAssignmentApprovals(value []Approvalable)() {
    m.accessPackageAssignmentApprovals = value
}
// SetAccessPackageAssignmentPolicies sets the accessPackageAssignmentPolicies property value. Represents the policy that governs which subjects can request or be assigned an access package via an access package assignment.
func (m *EntitlementManagement) SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicyable)() {
    m.accessPackageAssignmentPolicies = value
}
// SetAccessPackageAssignmentRequests sets the accessPackageAssignmentRequests property value. Represents access package assignment requests created by or on behalf of a user.
func (m *EntitlementManagement) SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)() {
    m.accessPackageAssignmentRequests = value
}
// SetAccessPackageAssignmentResourceRoles sets the accessPackageAssignmentResourceRoles property value. Represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *EntitlementManagement) SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)() {
    m.accessPackageAssignmentResourceRoles = value
}
// SetAccessPackageAssignments sets the accessPackageAssignments property value. The assignment of an access package to a subject for a period of time.
func (m *EntitlementManagement) SetAccessPackageAssignments(value []AccessPackageAssignmentable)() {
    m.accessPackageAssignments = value
}
// SetAccessPackageCatalogs sets the accessPackageCatalogs property value. A container of access packages.
func (m *EntitlementManagement) SetAccessPackageCatalogs(value []AccessPackageCatalogable)() {
    m.accessPackageCatalogs = value
}
// SetAccessPackageResourceEnvironments sets the accessPackageResourceEnvironments property value. A reference to the geolocation environment in which a resource is located.
func (m *EntitlementManagement) SetAccessPackageResourceEnvironments(value []AccessPackageResourceEnvironmentable)() {
    m.accessPackageResourceEnvironments = value
}
// SetAccessPackageResourceRequests sets the accessPackageResourceRequests property value. Represents a request to add or remove a resource to or from a catalog respectively.
func (m *EntitlementManagement) SetAccessPackageResourceRequests(value []AccessPackageResourceRequestable)() {
    m.accessPackageResourceRequests = value
}
// SetAccessPackageResourceRoleScopes sets the accessPackageResourceRoleScopes property value. A reference to both a scope within a resource, and a role in that resource for that scope.
func (m *EntitlementManagement) SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScopeable)() {
    m.accessPackageResourceRoleScopes = value
}
// SetAccessPackageResources sets the accessPackageResources property value. A reference to a resource associated with an access package catalog.
func (m *EntitlementManagement) SetAccessPackageResources(value []AccessPackageResourceable)() {
    m.accessPackageResources = value
}
// SetAccessPackages sets the accessPackages property value. Represents access package objects.
func (m *EntitlementManagement) SetAccessPackages(value []AccessPackageable)() {
    m.accessPackages = value
}
// SetConnectedOrganizations sets the connectedOrganizations property value. Represents references to a directory or domain of another organization whose users can request access.
func (m *EntitlementManagement) SetConnectedOrganizations(value []ConnectedOrganizationable)() {
    m.connectedOrganizations = value
}
// SetSettings sets the settings property value. Represents the settings that control the behavior of Azure AD entitlement management.
func (m *EntitlementManagement) SetSettings(value EntitlementManagementSettingsable)() {
    m.settings = value
}
// SetSubjects sets the subjects property value. The subjects property
func (m *EntitlementManagement) SetSubjects(value []AccessPackageSubjectable)() {
    m.subjects = value
}
