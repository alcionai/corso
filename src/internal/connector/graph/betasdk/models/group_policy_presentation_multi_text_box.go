package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationMultiTextBox 
type GroupPolicyPresentationMultiTextBox struct {
    GroupPolicyUploadedPresentation
    // An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
    maxLength *int64
    // An unsigned integer that specifies the maximum number of strings. Default value is 0.
    maxStrings *int64
    // Requirement to enter a value in the text box. Default value is false.
    required *bool
}
// NewGroupPolicyPresentationMultiTextBox instantiates a new GroupPolicyPresentationMultiTextBox and sets the default values.
func NewGroupPolicyPresentationMultiTextBox()(*GroupPolicyPresentationMultiTextBox) {
    m := &GroupPolicyPresentationMultiTextBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationMultiTextBox";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationMultiTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationMultiTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationMultiTextBox(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationMultiTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["maxLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLength(val)
        }
        return nil
    }
    res["maxStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxStrings(val)
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
    return res
}
// GetMaxLength gets the maxLength property value. An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
func (m *GroupPolicyPresentationMultiTextBox) GetMaxLength()(*int64) {
    return m.maxLength
}
// GetMaxStrings gets the maxStrings property value. An unsigned integer that specifies the maximum number of strings. Default value is 0.
func (m *GroupPolicyPresentationMultiTextBox) GetMaxStrings()(*int64) {
    return m.maxStrings
}
// GetRequired gets the required property value. Requirement to enter a value in the text box. Default value is false.
func (m *GroupPolicyPresentationMultiTextBox) GetRequired()(*bool) {
    return m.required
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationMultiTextBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("maxLength", m.GetMaxLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maxStrings", m.GetMaxStrings())
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
    return nil
}
// SetMaxLength sets the maxLength property value. An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
func (m *GroupPolicyPresentationMultiTextBox) SetMaxLength(value *int64)() {
    m.maxLength = value
}
// SetMaxStrings sets the maxStrings property value. An unsigned integer that specifies the maximum number of strings. Default value is 0.
func (m *GroupPolicyPresentationMultiTextBox) SetMaxStrings(value *int64)() {
    m.maxStrings = value
}
// SetRequired sets the required property value. Requirement to enter a value in the text box. Default value is false.
func (m *GroupPolicyPresentationMultiTextBox) SetRequired(value *bool)() {
    m.required = value
}
