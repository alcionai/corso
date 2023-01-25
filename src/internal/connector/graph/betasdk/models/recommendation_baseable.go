package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendationBaseable 
type RecommendationBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSteps()([]ActionStepable)
    GetBenefits()(*string)
    GetCategory()(*RecommendationCategory)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentScore()(*float64)
    GetDisplayName()(*string)
    GetFeatureAreas()([]RecommendationFeatureAreas)
    GetImpactedResources()([]ImpactedResourceable)
    GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetImpactType()(*string)
    GetInsights()(*string)
    GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaxScore()(*float64)
    GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPriority()(*RecommendationPriority)
    GetRecommendationType()(*RecommendationType)
    GetRemediationImpact()(*string)
    GetStatus()(*RecommendationStatus)
    SetActionSteps(value []ActionStepable)()
    SetBenefits(value *string)()
    SetCategory(value *RecommendationCategory)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentScore(value *float64)()
    SetDisplayName(value *string)()
    SetFeatureAreas(value []RecommendationFeatureAreas)()
    SetImpactedResources(value []ImpactedResourceable)()
    SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetImpactType(value *string)()
    SetInsights(value *string)()
    SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaxScore(value *float64)()
    SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPriority(value *RecommendationPriority)()
    SetRecommendationType(value *RecommendationType)()
    SetRemediationImpact(value *string)()
    SetStatus(value *RecommendationStatus)()
}
