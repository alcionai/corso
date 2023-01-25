package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerBucketPropertyRule 
type PlannerBucketPropertyRule struct {
    PlannerPropertyRule
    // The order property
    order []string
    // The title property
    title []string
}
// NewPlannerBucketPropertyRule instantiates a new PlannerBucketPropertyRule and sets the default values.
func NewPlannerBucketPropertyRule()(*PlannerBucketPropertyRule) {
    m := &PlannerBucketPropertyRule{
        PlannerPropertyRule: *NewPlannerPropertyRule(),
    }
    odataTypeValue := "#microsoft.graph.plannerBucketPropertyRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerBucketPropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerBucketPropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerBucketPropertyRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerBucketPropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPropertyRule.GetFieldDeserializers()
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
// GetOrder gets the order property value. The order property
func (m *PlannerBucketPropertyRule) GetOrder()([]string) {
    return m.order
}
// GetTitle gets the title property value. The title property
func (m *PlannerBucketPropertyRule) GetTitle()([]string) {
    return m.title
}
// Serialize serializes information the current object
func (m *PlannerBucketPropertyRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerPropertyRule.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOrder() != nil {
        err = writer.WriteCollectionOfStringValues("order", m.GetOrder())
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
// SetOrder sets the order property value. The order property
func (m *PlannerBucketPropertyRule) SetOrder(value []string)() {
    m.order = value
}
// SetTitle sets the title property value. The title property
func (m *PlannerBucketPropertyRule) SetTitle(value []string)() {
    m.title = value
}
