package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationTextBox 
type GroupPolicyPresentationTextBox struct {
    GroupPolicyUploadedPresentation
    // Localized default string displayed in the text box. The default value is empty.
    defaultValue *string
    // An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
    maxLength *int64
    // Requirement to enter a value in the text box. Default value is false.
    required *bool
}
// NewGroupPolicyPresentationTextBox instantiates a new GroupPolicyPresentationTextBox and sets the default values.
func NewGroupPolicyPresentationTextBox()(*GroupPolicyPresentationTextBox) {
    m := &GroupPolicyPresentationTextBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationTextBox";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationTextBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationTextBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationTextBox(), nil
}
// GetDefaultValue gets the defaultValue property value. Localized default string displayed in the text box. The default value is empty.
func (m *GroupPolicyPresentationTextBox) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationTextBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
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
func (m *GroupPolicyPresentationTextBox) GetMaxLength()(*int64) {
    return m.maxLength
}
// GetRequired gets the required property value. Requirement to enter a value in the text box. Default value is false.
func (m *GroupPolicyPresentationTextBox) GetRequired()(*bool) {
    return m.required
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationTextBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maxLength", m.GetMaxLength())
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
// SetDefaultValue sets the defaultValue property value. Localized default string displayed in the text box. The default value is empty.
func (m *GroupPolicyPresentationTextBox) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetMaxLength sets the maxLength property value. An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
func (m *GroupPolicyPresentationTextBox) SetMaxLength(value *int64)() {
    m.maxLength = value
}
// SetRequired sets the required property value. Requirement to enter a value in the text box. Default value is false.
func (m *GroupPolicyPresentationTextBox) SetRequired(value *bool)() {
    m.required = value
}
