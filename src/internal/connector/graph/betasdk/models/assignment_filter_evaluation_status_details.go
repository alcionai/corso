package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterEvaluationStatusDetails a class containing information about the payloads on which filter has been applied.
type AssignmentFilterEvaluationStatusDetails struct {
    Entity
    // PayloadId on which filter has been applied.
    payloadId *string
}
// NewAssignmentFilterEvaluationStatusDetails instantiates a new assignmentFilterEvaluationStatusDetails and sets the default values.
func NewAssignmentFilterEvaluationStatusDetails()(*AssignmentFilterEvaluationStatusDetails) {
    m := &AssignmentFilterEvaluationStatusDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterEvaluationStatusDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterEvaluationStatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    return res
}
// GetPayloadId gets the payloadId property value. PayloadId on which filter has been applied.
func (m *AssignmentFilterEvaluationStatusDetails) GetPayloadId()(*string) {
    return m.payloadId
}
// Serialize serializes information the current object
func (m *AssignmentFilterEvaluationStatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPayloadId sets the payloadId property value. PayloadId on which filter has been applied.
func (m *AssignmentFilterEvaluationStatusDetails) SetPayloadId(value *string)() {
    m.payloadId = value
}
