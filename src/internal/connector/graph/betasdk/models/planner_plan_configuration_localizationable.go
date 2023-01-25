package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanConfigurationLocalizationable 
type PlannerPlanConfigurationLocalizationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuckets()([]PlannerPlanConfigurationBucketLocalizationable)
    GetLanguageTag()(*string)
    GetPlanTitle()(*string)
    SetBuckets(value []PlannerPlanConfigurationBucketLocalizationable)()
    SetLanguageTag(value *string)()
    SetPlanTitle(value *string)()
}
