package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyUploadedPresentation 
type GroupPolicyUploadedPresentation struct {
    GroupPolicyPresentation
}
// NewGroupPolicyUploadedPresentation instantiates a new GroupPolicyUploadedPresentation and sets the default values.
func NewGroupPolicyUploadedPresentation()(*GroupPolicyUploadedPresentation) {
    m := &GroupPolicyUploadedPresentation{
        GroupPolicyPresentation: *NewGroupPolicyPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyUploadedPresentation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyUploadedPresentationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyUploadedPresentationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.groupPolicyPresentationCheckBox":
                        return NewGroupPolicyPresentationCheckBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationComboBox":
                        return NewGroupPolicyPresentationComboBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationDecimalTextBox":
                        return NewGroupPolicyPresentationDecimalTextBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationDropdownList":
                        return NewGroupPolicyPresentationDropdownList(), nil
                    case "#microsoft.graph.groupPolicyPresentationListBox":
                        return NewGroupPolicyPresentationListBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationLongDecimalTextBox":
                        return NewGroupPolicyPresentationLongDecimalTextBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationMultiTextBox":
                        return NewGroupPolicyPresentationMultiTextBox(), nil
                    case "#microsoft.graph.groupPolicyPresentationText":
                        return NewGroupPolicyPresentationText(), nil
                    case "#microsoft.graph.groupPolicyPresentationTextBox":
                        return NewGroupPolicyPresentationTextBox(), nil
                }
            }
        }
    }
    return NewGroupPolicyUploadedPresentation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyUploadedPresentation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *GroupPolicyUploadedPresentation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
