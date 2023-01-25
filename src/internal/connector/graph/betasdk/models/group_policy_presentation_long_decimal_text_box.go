package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationLongDecimalTextBox 
type GroupPolicyPresentationLongDecimalTextBox struct {
    GroupPolicyUploadedPresentation
    // An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
    defaultValue *int64
    // An unsigned long that specifies the maximum allowed value. The default value is 9999.
    maxValue *int64
    // An unsigned long that specifies the minimum allowed value. The default value is 0.
    minValue *int64
    // Requirement to enter a value in the parameter box. The default value is false.
    required *bool
    // If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
    spin *bool
    // An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
    spinStep *int64
}
// NewGroupPolicyPresentationLongDecimalTextBox instantiates a new GroupPolicyPresentationLongDecimalTextBox and sets the default values.
func NewGroupPolicyPresentationLongDecimalTextBox()(*GroupPolicyPresentationLongDecimalTextBox) {
    m := &GroupPolicyPresentationLongDecimalTextBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationLongDecimalTextBox";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationLongDecimalTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationLongDecimalTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationLongDecimalTextBox(), nil
}
// GetDefaultValue gets the defaultValue property value. An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
func (m *GroupPolicyPresentationLongDecimalTextBox) GetDefaultValue()(*int64) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationLongDecimalTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["maxValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxValue(val)
        }
        return nil
    }
    res["minValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinValue(val)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["spin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpin(val)
        }
        return nil
    }
    res["spinStep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpinStep(val)
        }
        return nil
    }
    return res
}
// GetMaxValue gets the maxValue property value. An unsigned long that specifies the maximum allowed value. The default value is 9999.
func (m *GroupPolicyPresentationLongDecimalTextBox) GetMaxValue()(*int64) {
    return m.maxValue
}
// GetMinValue gets the minValue property value. An unsigned long that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationLongDecimalTextBox) GetMinValue()(*int64) {
    return m.minValue
}
// GetRequired gets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationLongDecimalTextBox) GetRequired()(*bool) {
    return m.required
}
// GetSpin gets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationLongDecimalTextBox) GetSpin()(*bool) {
    return m.spin
}
// GetSpinStep gets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationLongDecimalTextBox) GetSpinStep()(*int64) {
    return m.spinStep
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationLongDecimalTextBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maxValue", m.GetMaxValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("minValue", m.GetMinValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("spin", m.GetSpin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("spinStep", m.GetSpinStep())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
func (m *GroupPolicyPresentationLongDecimalTextBox) SetDefaultValue(value *int64)() {
    m.defaultValue = value
}
// SetMaxValue sets the maxValue property value. An unsigned long that specifies the maximum allowed value. The default value is 9999.
func (m *GroupPolicyPresentationLongDecimalTextBox) SetMaxValue(value *int64)() {
    m.maxValue = value
}
// SetMinValue sets the minValue property value. An unsigned long that specifies the minimum allowed value. The default value is 0.
func (m *GroupPolicyPresentationLongDecimalTextBox) SetMinValue(value *int64)() {
    m.minValue = value
}
// SetRequired sets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationLongDecimalTextBox) SetRequired(value *bool)() {
    m.required = value
}
// SetSpin sets the spin property value. If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
func (m *GroupPolicyPresentationLongDecimalTextBox) SetSpin(value *bool)() {
    m.spin = value
}
// SetSpinStep sets the spinStep property value. An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
func (m *GroupPolicyPresentationLongDecimalTextBox) SetSpinStep(value *int64)() {
    m.spinStep = value
}
