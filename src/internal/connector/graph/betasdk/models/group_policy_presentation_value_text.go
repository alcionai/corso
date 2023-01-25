package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValueText 
type GroupPolicyPresentationValueText struct {
    GroupPolicyPresentationValue
    // A string value for the associated presentation.
    value *string
}
// NewGroupPolicyPresentationValueText instantiates a new GroupPolicyPresentationValueText and sets the default values.
func NewGroupPolicyPresentationValueText()(*GroupPolicyPresentationValueText) {
    m := &GroupPolicyPresentationValueText{
        GroupPolicyPresentationValue: *NewGroupPolicyPresentationValue(),
    }
    return m
}
// CreateGroupPolicyPresentationValueTextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationValueTextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationValueText(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationValueText) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentationValue.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. A string value for the associated presentation.
func (m *GroupPolicyPresentationValueText) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValueText) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentationValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. A string value for the associated presentation.
func (m *GroupPolicyPresentationValueText) SetValue(value *string)() {
    m.value = value
}
