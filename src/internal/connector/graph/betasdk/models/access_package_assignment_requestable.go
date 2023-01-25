package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestable 
type AccessPackageAssignmentRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAccessPackageAssignment()(AccessPackageAssignmentable)
    GetAnswers()([]AccessPackageAnswerable)
    GetCompletedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomExtensionHandlerInstances()([]CustomExtensionHandlerInstanceable)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsValidationOnly()(*bool)
    GetJustification()(*string)
    GetRequestor()(AccessPackageSubjectable)
    GetRequestState()(*string)
    GetRequestStatus()(*string)
    GetRequestType()(*string)
    GetSchedule()(RequestScheduleable)
    SetAccessPackage(value AccessPackageable)()
    SetAccessPackageAssignment(value AccessPackageAssignmentable)()
    SetAnswers(value []AccessPackageAnswerable)()
    SetCompletedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomExtensionHandlerInstances(value []CustomExtensionHandlerInstanceable)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsValidationOnly(value *bool)()
    SetJustification(value *string)()
    SetRequestor(value AccessPackageSubjectable)()
    SetRequestState(value *string)()
    SetRequestStatus(value *string)()
    SetRequestType(value *string)()
    SetSchedule(value RequestScheduleable)()
}
