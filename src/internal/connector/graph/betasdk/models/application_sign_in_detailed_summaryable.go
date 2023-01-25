package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationSignInDetailedSummaryable 
type ApplicationSignInDetailedSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregatedEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetSignInCount()(*int64)
    GetStatus()(SignInStatusable)
    SetAggregatedEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetSignInCount(value *int64)()
    SetStatus(value SignInStatusable)()
}
