package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationValueBoolean 
type GroupPolicyPresentationValueBoolean struct {
    GroupPolicyPresentationValue
    // An boolean value for the associated presentation.
    value *bool
}
// NewGroupPolicyPresentationValueBoolean instantiates a new GroupPolicyPresentationValueBoolean and sets the default values.
func NewGroupPolicyPresentationValueBoolean()(*GroupPolicyPresentationValueBoolean) {
    m := &GroupPolicyPresentationValueBoolean{
        GroupPolicyPresentationValue: *NewGroupPolicyPresentationValue(),
    }
    return m
}
// CreateGroupPolicyPresentationValueBooleanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationValueBooleanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationValueBoolean(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationValueBoolean) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentationValue.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
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
// GetValue gets the value property value. An boolean value for the associated presentation.
func (m *GroupPolicyPresentationValueBoolean) GetValue()(*bool) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValueBoolean) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentationValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. An boolean value for the associated presentation.
func (m *GroupPolicyPresentationValueBoolean) SetValue(value *bool)() {
    m.value = value
}
