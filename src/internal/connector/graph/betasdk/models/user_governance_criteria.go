package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserGovernanceCriteria 
type UserGovernanceCriteria struct {
    GovernanceCriteria
    // The userId property
    userId *string
}
// NewUserGovernanceCriteria instantiates a new UserGovernanceCriteria and sets the default values.
func NewUserGovernanceCriteria()(*UserGovernanceCriteria) {
    m := &UserGovernanceCriteria{
        GovernanceCriteria: *NewGovernanceCriteria(),
    }
    odataTypeValue := "#microsoft.graph.userGovernanceCriteria";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserGovernanceCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserGovernanceCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserGovernanceCriteria(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserGovernanceCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GovernanceCriteria.GetFieldDeserializers()
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the userId property value. The userId property
func (m *UserGovernanceCriteria) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *UserGovernanceCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GovernanceCriteria.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserId sets the userId property value. The userId property
func (m *UserGovernanceCriteria) SetUserId(value *string)() {
    m.userId = value
}
