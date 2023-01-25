package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageable 
type AccessPackageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageAssignmentPolicies()([]AccessPackageAssignmentPolicyable)
    GetAccessPackageCatalog()(AccessPackageCatalogable)
    GetAccessPackageResourceRoleScopes()([]AccessPackageResourceRoleScopeable)
    GetAccessPackagesIncompatibleWith()([]AccessPackageable)
    GetCatalogId()(*string)
    GetCreatedBy()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIncompatibleAccessPackages()([]AccessPackageable)
    GetIncompatibleGroups()([]Groupable)
    GetIsHidden()(*bool)
    GetIsRoleScopesVisible()(*bool)
    GetModifiedBy()(*string)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessPackageAssignmentPolicies(value []AccessPackageAssignmentPolicyable)()
    SetAccessPackageCatalog(value AccessPackageCatalogable)()
    SetAccessPackageResourceRoleScopes(value []AccessPackageResourceRoleScopeable)()
    SetAccessPackagesIncompatibleWith(value []AccessPackageable)()
    SetCatalogId(value *string)()
    SetCreatedBy(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIncompatibleAccessPackages(value []AccessPackageable)()
    SetIncompatibleGroups(value []Groupable)()
    SetIsHidden(value *bool)()
    SetIsRoleScopesVisible(value *bool)()
    SetModifiedBy(value *string)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
