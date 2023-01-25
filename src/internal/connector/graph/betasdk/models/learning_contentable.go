package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LearningContentable 
type LearningContentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalTags()([]string)
    GetContentWebUrl()(*string)
    GetContributors()([]string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetExternalId()(*string)
    GetFormat()(*string)
    GetIsActive()(*bool)
    GetIsPremium()(*bool)
    GetIsSearchable()(*bool)
    GetLanguageTag()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumberOfPages()(*int32)
    GetSkillTags()([]string)
    GetSourceName()(*string)
    GetThumbnailWebUrl()(*string)
    GetTitle()(*string)
    SetAdditionalTags(value []string)()
    SetContentWebUrl(value *string)()
    SetContributors(value []string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetExternalId(value *string)()
    SetFormat(value *string)()
    SetIsActive(value *bool)()
    SetIsPremium(value *bool)()
    SetIsSearchable(value *bool)()
    SetLanguageTag(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumberOfPages(value *int32)()
    SetSkillTags(value []string)()
    SetSourceName(value *string)()
    SetThumbnailWebUrl(value *string)()
    SetTitle(value *string)()
}
