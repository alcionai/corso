package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationText 
type GroupPolicyPresentationText struct {
    GroupPolicyUploadedPresentation
}
// NewGroupPolicyPresentationText instantiates a new GroupPolicyPresentationText and sets the default values.
func NewGroupPolicyPresentationText()(*GroupPolicyPresentationText) {
    m := &GroupPolicyPresentationText{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationText";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationTextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationTextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationText(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationText) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationText) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
