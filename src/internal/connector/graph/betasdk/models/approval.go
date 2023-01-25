package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Approval 
type Approval struct {
    Entity
    // The steps property
    steps []ApprovalStepable
}
// NewApproval instantiates a new approval and sets the default values.
func NewApproval()(*Approval) {
    m := &Approval{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApproval(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Approval) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["steps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalStepable, len(val))
            for i, v := range val {
                res[i] = v.(ApprovalStepable)
            }
            m.SetSteps(res)
        }
        return nil
    }
    return res
}
// GetSteps gets the steps property value. The steps property
func (m *Approval) GetSteps()([]ApprovalStepable) {
    return m.steps
}
// Serialize serializes information the current object
func (m *Approval) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSteps()))
        for i, v := range m.GetSteps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("steps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSteps sets the steps property value. The steps property
func (m *Approval) SetSteps(value []ApprovalStepable)() {
    m.steps = value
}
