package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeClientCheckinStatusable 
type OfficeClientCheckinStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppliedPolicies()([]string)
    GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceName()(*string)
    GetDevicePlatform()(*string)
    GetDevicePlatformVersion()(*string)
    GetErrorMessage()(*string)
    GetOdataType()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetWasSuccessful()(*bool)
    SetAppliedPolicies(value []string)()
    SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceName(value *string)()
    SetDevicePlatform(value *string)()
    SetDevicePlatformVersion(value *string)()
    SetErrorMessage(value *string)()
    SetOdataType(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetWasSuccessful(value *bool)()
}
