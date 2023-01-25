package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationComboBoxCollectionResponse 
type GroupPolicyPresentationComboBoxCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []GroupPolicyPresentationComboBoxable
}
// NewGroupPolicyPresentationComboBoxCollectionResponse instantiates a new GroupPolicyPresentationComboBoxCollectionResponse and sets the default values.
func NewGroupPolicyPresentationComboBoxCollectionResponse()(*GroupPolicyPresentationComboBoxCollectionResponse) {
    m := &GroupPolicyPresentationComboBoxCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateGroupPolicyPresentationComboBoxCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationComboBoxCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationComboBoxCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationComboBoxCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyPresentationComboBoxFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentationComboBoxable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyPresentationComboBoxable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *GroupPolicyPresentationComboBoxCollectionResponse) GetValue()([]GroupPolicyPresentationComboBoxable) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationComboBoxCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GroupPolicyPresentationComboBoxCollectionResponse) SetValue(value []GroupPolicyPresentationComboBoxable)() {
    m.value = value
}
