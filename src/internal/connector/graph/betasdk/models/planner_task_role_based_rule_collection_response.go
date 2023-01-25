package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskRoleBasedRuleCollectionResponse 
type PlannerTaskRoleBasedRuleCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []PlannerTaskRoleBasedRuleable
}
// NewPlannerTaskRoleBasedRuleCollectionResponse instantiates a new PlannerTaskRoleBasedRuleCollectionResponse and sets the default values.
func NewPlannerTaskRoleBasedRuleCollectionResponse()(*PlannerTaskRoleBasedRuleCollectionResponse) {
    m := &PlannerTaskRoleBasedRuleCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePlannerTaskRoleBasedRuleCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskRoleBasedRuleCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskRoleBasedRuleCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskRoleBasedRuleCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerTaskRoleBasedRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerTaskRoleBasedRuleable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerTaskRoleBasedRuleable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PlannerTaskRoleBasedRuleCollectionResponse) GetValue()([]PlannerTaskRoleBasedRuleable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PlannerTaskRoleBasedRuleCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *PlannerTaskRoleBasedRuleCollectionResponse) SetValue(value []PlannerTaskRoleBasedRuleable)() {
    m.value = value
}
