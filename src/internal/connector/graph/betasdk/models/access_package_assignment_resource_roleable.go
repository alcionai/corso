package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentResourceRoleable 
type AccessPackageAssignmentResourceRoleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageAssignments()([]AccessPackageAssignmentable)
    GetAccessPackageResourceRole()(AccessPackageResourceRoleable)
    GetAccessPackageResourceScope()(AccessPackageResourceScopeable)
    GetAccessPackageSubject()(AccessPackageSubjectable)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetStatus()(*string)
    SetAccessPackageAssignments(value []AccessPackageAssignmentable)()
    SetAccessPackageResourceRole(value AccessPackageResourceRoleable)()
    SetAccessPackageResourceScope(value AccessPackageResourceScopeable)()
    SetAccessPackageSubject(value AccessPackageSubjectable)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetStatus(value *string)()
}
