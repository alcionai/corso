package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSiteHistoryable 
type BrowserSiteHistoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowRedirect()(*bool)
    GetComment()(*string)
    GetCompatibilityMode()(*BrowserSiteCompatibilityMode)
    GetLastModifiedBy()(IdentitySetable)
    GetMergeType()(*BrowserSiteMergeType)
    GetOdataType()(*string)
    GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTargetEnvironment()(*BrowserSiteTargetEnvironment)
    SetAllowRedirect(value *bool)()
    SetComment(value *string)()
    SetCompatibilityMode(value *BrowserSiteCompatibilityMode)()
    SetLastModifiedBy(value IdentitySetable)()
    SetMergeType(value *BrowserSiteMergeType)()
    SetOdataType(value *string)()
    SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTargetEnvironment(value *BrowserSiteTargetEnvironment)()
}
