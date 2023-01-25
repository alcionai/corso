package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImpactedResourceable 
type ImpactedResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAdditionalDetails()([]KeyValueable)
    GetApiUrl()(*string)
    GetDisplayName()(*string)
    GetLastModifiedBy()(*string)
    GetLastModifiedDateTime()(*string)
    GetOwner()(*string)
    GetPortalUrl()(*string)
    GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRank()(*int32)
    GetRecommendationId()(*string)
    GetResourceType()(*string)
    GetStatus()(*RecommendationStatus)
    GetSubjectId()(*string)
    SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAdditionalDetails(value []KeyValueable)()
    SetApiUrl(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedDateTime(value *string)()
    SetOwner(value *string)()
    SetPortalUrl(value *string)()
    SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRank(value *int32)()
    SetRecommendationId(value *string)()
    SetResourceType(value *string)()
    SetStatus(value *RecommendationStatus)()
    SetSubjectId(value *string)()
}
