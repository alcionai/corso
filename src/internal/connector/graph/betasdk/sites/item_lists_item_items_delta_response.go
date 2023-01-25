package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ItemListsItemItemsDeltaResponse provides operations to call the delta method.
type ItemListsItemItemsDeltaResponse struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.BaseDeltaFunctionResponse
    // The value property
    value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ListItemable
}
// NewItemListsItemItemsDeltaResponse instantiates a new ItemListsItemItemsDeltaResponse and sets the default values.
func NewItemListsItemItemsDeltaResponse()(*ItemListsItemItemsDeltaResponse) {
    m := &ItemListsItemItemsDeltaResponse{
        BaseDeltaFunctionResponse: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateItemListsItemItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemItemsDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListsItemItemsDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ListItemable, len(val))
            for i, v := range val {
                res[i] = v.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ListItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemListsItemItemsDeltaResponse) GetValue()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ListItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemListsItemItemsDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
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
func (m *ItemListsItemItemsDeltaResponse) SetValue(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ListItemable)() {
    m.value = value
}
