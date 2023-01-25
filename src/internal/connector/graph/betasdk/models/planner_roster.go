package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerRoster provides operations to call the add method.
type PlannerRoster struct {
    Entity
    // Retrieves the members of the plannerRoster.
    members []PlannerRosterMemberable
    // Retrieves the plans contained by the plannerRoster.
    plans []PlannerPlanable
}
// NewPlannerRoster instantiates a new plannerRoster and sets the default values.
func NewPlannerRoster()(*PlannerRoster) {
    m := &PlannerRoster{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerRosterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerRosterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerRoster(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerRoster) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerRosterMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerRosterMemberable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerRosterMemberable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["plans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanable)
            }
            m.SetPlans(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Retrieves the members of the plannerRoster.
func (m *PlannerRoster) GetMembers()([]PlannerRosterMemberable) {
    return m.members
}
// GetPlans gets the plans property value. Retrieves the plans contained by the plannerRoster.
func (m *PlannerRoster) GetPlans()([]PlannerPlanable) {
    return m.plans
}
// Serialize serializes information the current object
func (m *PlannerRoster) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("plans", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMembers sets the members property value. Retrieves the members of the plannerRoster.
func (m *PlannerRoster) SetMembers(value []PlannerRosterMemberable)() {
    m.members = value
}
// SetPlans sets the plans property value. Retrieves the plans contained by the plannerRoster.
func (m *PlannerRoster) SetPlans(value []PlannerPlanable)() {
    m.plans = value
}
