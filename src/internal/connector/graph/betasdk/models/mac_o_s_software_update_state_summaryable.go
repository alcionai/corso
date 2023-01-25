package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateStateSummaryable 
type MacOSSoftwareUpdateStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProductKey()(*string)
    GetState()(*MacOSSoftwareUpdateState)
    GetUpdateCategory()(*MacOSSoftwareUpdateCategory)
    GetUpdateVersion()(*string)
    SetDisplayName(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProductKey(value *string)()
    SetState(value *MacOSSoftwareUpdateState)()
    SetUpdateCategory(value *MacOSSoftwareUpdateCategory)()
    SetUpdateVersion(value *string)()
}
