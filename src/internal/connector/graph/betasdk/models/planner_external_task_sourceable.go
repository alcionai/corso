package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerExternalTaskSourceable 
type PlannerExternalTaskSourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerTaskCreationable
    GetContextScenarioId()(*string)
    GetDisplayLinkType()(*PlannerExternalTaskSourceDisplayType)
    GetDisplayNameSegments()([]string)
    GetExternalContextId()(*string)
    GetExternalObjectId()(*string)
    GetExternalObjectVersion()(*string)
    GetWebUrl()(*string)
    SetContextScenarioId(value *string)()
    SetDisplayLinkType(value *PlannerExternalTaskSourceDisplayType)()
    SetDisplayNameSegments(value []string)()
    SetExternalContextId(value *string)()
    SetExternalObjectId(value *string)()
    SetExternalObjectVersion(value *string)()
    SetWebUrl(value *string)()
}
