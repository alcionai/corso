package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSiteable 
type BrowserSiteable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowRedirect()(*bool)
    GetComment()(*string)
    GetCompatibilityMode()(*BrowserSiteCompatibilityMode)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHistory()([]BrowserSiteHistoryable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMergeType()(*BrowserSiteMergeType)
    GetStatus()(*BrowserSiteStatus)
    GetTargetEnvironment()(*BrowserSiteTargetEnvironment)
    GetWebUrl()(*string)
    SetAllowRedirect(value *bool)()
    SetComment(value *string)()
    SetCompatibilityMode(value *BrowserSiteCompatibilityMode)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHistory(value []BrowserSiteHistoryable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMergeType(value *BrowserSiteMergeType)()
    SetStatus(value *BrowserSiteStatus)()
    SetTargetEnvironment(value *BrowserSiteTargetEnvironment)()
    SetWebUrl(value *string)()
}
