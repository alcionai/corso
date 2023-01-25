package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentable 
type AccessPackageAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAccessPackageAssignmentPolicy()(AccessPackageAssignmentPolicyable)
    GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable)
    GetAccessPackageId()(*string)
    GetAssignmentPolicyId()(*string)
    GetAssignmentState()(*string)
    GetAssignmentStatus()(*string)
    GetCatalogId()(*string)
    GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsExtended()(*bool)
    GetSchedule()(RequestScheduleable)
    GetTarget()(AccessPackageSubjectable)
    GetTargetId()(*string)
    SetAccessPackage(value AccessPackageable)()
    SetAccessPackageAssignmentPolicy(value AccessPackageAssignmentPolicyable)()
    SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)()
    SetAccessPackageId(value *string)()
    SetAssignmentPolicyId(value *string)()
    SetAssignmentState(value *string)()
    SetAssignmentStatus(value *string)()
    SetCatalogId(value *string)()
    SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsExtended(value *bool)()
    SetSchedule(value RequestScheduleable)()
    SetTarget(value AccessPackageSubjectable)()
    SetTargetId(value *string)()
}
