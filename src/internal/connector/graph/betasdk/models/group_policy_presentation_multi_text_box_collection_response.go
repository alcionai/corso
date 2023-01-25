package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationMultiTextBoxCollectionResponse 
type GroupPolicyPresentationMultiTextBoxCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []GroupPolicyPresentationMultiTextBoxable
}
// NewGroupPolicyPresentationMultiTextBoxCollectionResponse instantiates a new GroupPolicyPresentationMultiTextBoxCollectionResponse and sets the default values.
func NewGroupPolicyPresentationMultiTextBoxCollectionResponse()(*GroupPolicyPresentationMultiTextBoxCollectionResponse) {
    m := &GroupPolicyPresentationMultiTextBoxCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateGroupPolicyPresentationMultiTextBoxCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationMultiTextBoxCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationMultiTextBoxCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationMultiTextBoxCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyPresentationMultiTextBoxFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentationMultiTextBoxable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyPresentationMultiTextBoxable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *GroupPolicyPresentationMultiTextBoxCollectionResponse) GetValue()([]GroupPolicyPresentationMultiTextBoxable) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationMultiTextBoxCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *GroupPolicyPresentationMultiTextBoxCollectionResponse) SetValue(value []GroupPolicyPresentationMultiTextBoxable)() {
    m.value = value
}
