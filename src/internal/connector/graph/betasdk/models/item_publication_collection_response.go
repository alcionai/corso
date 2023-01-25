package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPublicationCollectionResponse 
type ItemPublicationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ItemPublicationable
}
// NewItemPublicationCollectionResponse instantiates a new ItemPublicationCollectionResponse and sets the default values.
func NewItemPublicationCollectionResponse()(*ItemPublicationCollectionResponse) {
    m := &ItemPublicationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemPublicationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPublicationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPublicationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPublicationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemPublicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPublicationable, len(val))
            for i, v := range val {
                res[i] = v.(ItemPublicationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemPublicationCollectionResponse) GetValue()([]ItemPublicationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemPublicationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemPublicationCollectionResponse) SetValue(value []ItemPublicationable)() {
    m.value = value
}
