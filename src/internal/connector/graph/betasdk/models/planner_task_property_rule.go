package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskPropertyRule 
type PlannerTaskPropertyRule struct {
    PlannerPropertyRule
    // Rules and restrictions for applied categories. This value does not currently support overrides. Accepted values for the default rule and individual overrides are allow, block.
    appliedCategories PlannerFieldRulesable
    // Rules and restrictions for assignments. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, addSelf, addOther, remove, removeSelf, removeOther, block.
    assignments PlannerFieldRulesable
    // Rules and restrictions for checklist. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, update, check, reorder, block.
    checkLists PlannerFieldRulesable
    // Rules and restrictions for deleting the task. Accepted values are allow and block.
    delete []string
    // Rules and restrictions for changing the due date of the task. Accepted values are allow and block.
    dueDate []string
    // Rules and restrictions for moving the task between buckets or plans. Accepted values are allow, moveBetweenPlans, moveBetweenBuckets, and block.
    move []string
    // Rules and restrictions for changing the notes of the task. Accepted values are allow and block.
    notes []string
    // Rules and restrictions for changing the order of the task. Accepted values are allow and block.
    order []string
    // Rules and restrictions for changing the completion percentage of the task. Accepted values are allow, setToComplete, setToNotStarted, setToInProgress, and block.
    percentComplete []string
    // Rules and restrictions for changing the preview type of the task. Accepted values are allow and block.
    previewType []string
    // Rules and restrictions for changing the priority of the task. Accepted values are allow and block.
    priority []string
    // Rules and restrictions for references. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, block.
    references PlannerFieldRulesable
    // Rules and restrictions for changing the start date of the task. Accepted values are allow and block.
    startDate []string
    // Rules and restrictions for changing the title of the task. Accepted values are allow and block.
    title []string
}
// NewPlannerTaskPropertyRule instantiates a new PlannerTaskPropertyRule and sets the default values.
func NewPlannerTaskPropertyRule()(*PlannerTaskPropertyRule) {
    m := &PlannerTaskPropertyRule{
        PlannerPropertyRule: *NewPlannerPropertyRule(),
    }
    odataTypeValue := "#microsoft.graph.plannerTaskPropertyRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerTaskPropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskPropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskPropertyRule(), nil
}
// GetAppliedCategories gets the appliedCategories property value. Rules and restrictions for applied categories. This value does not currently support overrides. Accepted values for the default rule and individual overrides are allow, block.
func (m *PlannerTaskPropertyRule) GetAppliedCategories()(PlannerFieldRulesable) {
    return m.appliedCategories
}
// GetAssignments gets the assignments property value. Rules and restrictions for assignments. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, addSelf, addOther, remove, removeSelf, removeOther, block.
func (m *PlannerTaskPropertyRule) GetAssignments()(PlannerFieldRulesable) {
    return m.assignments
}
// GetCheckLists gets the checkLists property value. Rules and restrictions for checklist. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, update, check, reorder, block.
func (m *PlannerTaskPropertyRule) GetCheckLists()(PlannerFieldRulesable) {
    return m.checkLists
}
// GetDelete gets the delete property value. Rules and restrictions for deleting the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetDelete()([]string) {
    return m.delete
}
// GetDueDate gets the dueDate property value. Rules and restrictions for changing the due date of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetDueDate()([]string) {
    return m.dueDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskPropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPropertyRule.GetFieldDeserializers()
    res["appliedCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedCategories(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignments(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["checkLists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckLists(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["delete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDelete(res)
        }
        return nil
    }
    res["dueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDueDate(res)
        }
        return nil
    }
    res["move"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMove(res)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["order"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrder(res)
        }
        return nil
    }
    res["percentComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPercentComplete(res)
        }
        return nil
    }
    res["previewType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPreviewType(res)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPriority(res)
        }
        return nil
    }
    res["references"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferences(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStartDate(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTitle(res)
        }
        return nil
    }
    return res
}
// GetMove gets the move property value. Rules and restrictions for moving the task between buckets or plans. Accepted values are allow, moveBetweenPlans, moveBetweenBuckets, and block.
func (m *PlannerTaskPropertyRule) GetMove()([]string) {
    return m.move
}
// GetNotes gets the notes property value. Rules and restrictions for changing the notes of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetNotes()([]string) {
    return m.notes
}
// GetOrder gets the order property value. Rules and restrictions for changing the order of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetOrder()([]string) {
    return m.order
}
// GetPercentComplete gets the percentComplete property value. Rules and restrictions for changing the completion percentage of the task. Accepted values are allow, setToComplete, setToNotStarted, setToInProgress, and block.
func (m *PlannerTaskPropertyRule) GetPercentComplete()([]string) {
    return m.percentComplete
}
// GetPreviewType gets the previewType property value. Rules and restrictions for changing the preview type of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetPreviewType()([]string) {
    return m.previewType
}
// GetPriority gets the priority property value. Rules and restrictions for changing the priority of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetPriority()([]string) {
    return m.priority
}
// GetReferences gets the references property value. Rules and restrictions for references. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, block.
func (m *PlannerTaskPropertyRule) GetReferences()(PlannerFieldRulesable) {
    return m.references
}
// GetStartDate gets the startDate property value. Rules and restrictions for changing the start date of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetStartDate()([]string) {
    return m.startDate
}
// GetTitle gets the title property value. Rules and restrictions for changing the title of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) GetTitle()([]string) {
    return m.title
}
// Serialize serializes information the current object
func (m *PlannerTaskPropertyRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerPropertyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appliedCategories", m.GetAppliedCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignments", m.GetAssignments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("checkLists", m.GetCheckLists())
        if err != nil {
            return err
        }
    }
    if m.GetDelete() != nil {
        err = writer.WriteCollectionOfStringValues("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    if m.GetDueDate() != nil {
        err = writer.WriteCollectionOfStringValues("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    if m.GetMove() != nil {
        err = writer.WriteCollectionOfStringValues("move", m.GetMove())
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        err = writer.WriteCollectionOfStringValues("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetOrder() != nil {
        err = writer.WriteCollectionOfStringValues("order", m.GetOrder())
        if err != nil {
            return err
        }
    }
    if m.GetPercentComplete() != nil {
        err = writer.WriteCollectionOfStringValues("percentComplete", m.GetPercentComplete())
        if err != nil {
            return err
        }
    }
    if m.GetPreviewType() != nil {
        err = writer.WriteCollectionOfStringValues("previewType", m.GetPreviewType())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        err = writer.WriteCollectionOfStringValues("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("references", m.GetReferences())
        if err != nil {
            return err
        }
    }
    if m.GetStartDate() != nil {
        err = writer.WriteCollectionOfStringValues("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    if m.GetTitle() != nil {
        err = writer.WriteCollectionOfStringValues("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliedCategories sets the appliedCategories property value. Rules and restrictions for applied categories. This value does not currently support overrides. Accepted values for the default rule and individual overrides are allow, block.
func (m *PlannerTaskPropertyRule) SetAppliedCategories(value PlannerFieldRulesable)() {
    m.appliedCategories = value
}
// SetAssignments sets the assignments property value. Rules and restrictions for assignments. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, addSelf, addOther, remove, removeSelf, removeOther, block.
func (m *PlannerTaskPropertyRule) SetAssignments(value PlannerFieldRulesable)() {
    m.assignments = value
}
// SetCheckLists sets the checkLists property value. Rules and restrictions for checklist. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, update, check, reorder, block.
func (m *PlannerTaskPropertyRule) SetCheckLists(value PlannerFieldRulesable)() {
    m.checkLists = value
}
// SetDelete sets the delete property value. Rules and restrictions for deleting the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetDelete(value []string)() {
    m.delete = value
}
// SetDueDate sets the dueDate property value. Rules and restrictions for changing the due date of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetDueDate(value []string)() {
    m.dueDate = value
}
// SetMove sets the move property value. Rules and restrictions for moving the task between buckets or plans. Accepted values are allow, moveBetweenPlans, moveBetweenBuckets, and block.
func (m *PlannerTaskPropertyRule) SetMove(value []string)() {
    m.move = value
}
// SetNotes sets the notes property value. Rules and restrictions for changing the notes of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetNotes(value []string)() {
    m.notes = value
}
// SetOrder sets the order property value. Rules and restrictions for changing the order of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetOrder(value []string)() {
    m.order = value
}
// SetPercentComplete sets the percentComplete property value. Rules and restrictions for changing the completion percentage of the task. Accepted values are allow, setToComplete, setToNotStarted, setToInProgress, and block.
func (m *PlannerTaskPropertyRule) SetPercentComplete(value []string)() {
    m.percentComplete = value
}
// SetPreviewType sets the previewType property value. Rules and restrictions for changing the preview type of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetPreviewType(value []string)() {
    m.previewType = value
}
// SetPriority sets the priority property value. Rules and restrictions for changing the priority of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetPriority(value []string)() {
    m.priority = value
}
// SetReferences sets the references property value. Rules and restrictions for references. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, block.
func (m *PlannerTaskPropertyRule) SetReferences(value PlannerFieldRulesable)() {
    m.references = value
}
// SetStartDate sets the startDate property value. Rules and restrictions for changing the start date of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetStartDate(value []string)() {
    m.startDate = value
}
// SetTitle sets the title property value. Rules and restrictions for changing the title of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetTitle(value []string)() {
    m.title = value
}
