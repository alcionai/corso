package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReputationCategoryCollectionResponse 
type ReputationCategoryCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ReputationCategoryable
}
// NewReputationCategoryCollectionResponse instantiates a new ReputationCategoryCollectionResponse and sets the default values.
func NewReputationCategoryCollectionResponse()(*ReputationCategoryCollectionResponse) {
    m := &ReputationCategoryCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateReputationCategoryCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReputationCategoryCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReputationCategoryCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReputationCategoryCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateReputationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReputationCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(ReputationCategoryable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ReputationCategoryCollectionResponse) GetValue()([]ReputationCategoryable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ReputationCategoryCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ReputationCategoryCollectionResponse) SetValue(value []ReputationCategoryable)() {
    m.value = value
}
