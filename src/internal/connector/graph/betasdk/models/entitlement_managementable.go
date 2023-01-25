package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementable 
type EntitlementManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageAssignmentApprovals()([]Approvalable)
    GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicyable)
    GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable)
    GetAccessPackageAssignments()([]AccessPackageAssignmentable)
    GetAccessPackageCatalogs()([]AccessPackageCatalogable)
    GetAccessPackageResourceEnvironments()([]AccessPackageResourceEnvironmentable)
    GetAccessPackageResourceRequests()([]AccessPackageResourceRequestable)
    GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScopeable)
    GetAccessPackageResources()([]AccessPackageResourceable)
    GetAccessPackages()([]AccessPackageable)
    GetConnectedOrganizations()([]ConnectedOrganizationable)
    GetSettings()(EntitlementManagementSettingsable)
    GetSubjects()([]AccessPackageSubjectable)
    SetAccessPackageAssignmentApprovals(value []Approvalable)()
    SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicyable)()
    SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)()
    SetAccessPackageAssignments(value []AccessPackageAssignmentable)()
    SetAccessPackageCatalogs(value []AccessPackageCatalogable)()
    SetAccessPackageResourceEnvironments(value []AccessPackageResourceEnvironmentable)()
    SetAccessPackageResourceRequests(value []AccessPackageResourceRequestable)()
    SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScopeable)()
    SetAccessPackageResources(value []AccessPackageResourceable)()
    SetAccessPackages(value []AccessPackageable)()
    SetConnectedOrganizations(value []ConnectedOrganizationable)()
    SetSettings(value EntitlementManagementSettingsable)()
    SetSubjects(value []AccessPackageSubjectable)()
}
