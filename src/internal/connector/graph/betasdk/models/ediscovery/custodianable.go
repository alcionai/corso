package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Custodianable 
type Custodianable interface {
    DataSourceContainerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetApplyHoldToSources()(*bool)
    GetEmail()(*string)
    GetSiteSources()([]SiteSourceable)
    GetUnifiedGroupSources()([]UnifiedGroupSourceable)
    GetUserSources()([]UserSourceable)
    SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetApplyHoldToSources(value *bool)()
    SetEmail(value *string)()
    SetSiteSources(value []SiteSourceable)()
    SetUnifiedGroupSources(value []UnifiedGroupSourceable)()
    SetUserSources(value []UserSourceable)()
}
