package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationDetailable 
type AuthenticationDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*string)
    GetAuthenticationMethodDetail()(*string)
    GetAuthenticationStepDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAuthenticationStepRequirement()(*string)
    GetAuthenticationStepResultDetail()(*string)
    GetOdataType()(*string)
    GetSucceeded()(*bool)
    SetAuthenticationMethod(value *string)()
    SetAuthenticationMethodDetail(value *string)()
    SetAuthenticationStepDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAuthenticationStepRequirement(value *string)()
    SetAuthenticationStepResultDetail(value *string)()
    SetOdataType(value *string)()
    SetSucceeded(value *bool)()
}
