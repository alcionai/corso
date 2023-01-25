package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskPropertyRuleable 
type PlannerTaskPropertyRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerPropertyRuleable
    GetAppliedCategories()(PlannerFieldRulesable)
    GetAssignments()(PlannerFieldRulesable)
    GetCheckLists()(PlannerFieldRulesable)
    GetDelete()([]string)
    GetDueDate()([]string)
    GetMove()([]string)
    GetNotes()([]string)
    GetOrder()([]string)
    GetPercentComplete()([]string)
    GetPreviewType()([]string)
    GetPriority()([]string)
    GetReferences()(PlannerFieldRulesable)
    GetStartDate()([]string)
    GetTitle()([]string)
    SetAppliedCategories(value PlannerFieldRulesable)()
    SetAssignments(value PlannerFieldRulesable)()
    SetCheckLists(value PlannerFieldRulesable)()
    SetDelete(value []string)()
    SetDueDate(value []string)()
    SetMove(value []string)()
    SetNotes(value []string)()
    SetOrder(value []string)()
    SetPercentComplete(value []string)()
    SetPreviewType(value []string)()
    SetPriority(value []string)()
    SetReferences(value PlannerFieldRulesable)()
    SetStartDate(value []string)()
    SetTitle(value []string)()
}
