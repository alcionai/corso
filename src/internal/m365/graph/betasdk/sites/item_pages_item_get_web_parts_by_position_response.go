package sites

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"

	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/m365/graph/betasdk/models"
)

// ItemPagesItemGetWebPartsByPositionResponse provides operations to call the getWebPartsByPosition method.
type ItemPagesItemGetWebPartsByPositionResponse struct {
	msmodel.BaseCollectionPaginationCountResponse
	// The value property
	value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable
}

// NewItemPagesItemGetWebPartsByPositionResponse instantiates a new ItemPagesItemGetWebPartsByPositionResponse and sets the default values.
//
//nolint:wsl,lll
func NewItemPagesItemGetWebPartsByPositionResponse() *ItemPagesItemGetWebPartsByPositionResponse {
	m := &ItemPagesItemGetWebPartsByPositionResponse{
		BaseCollectionPaginationCountResponse: *msmodel.NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateItemPagesItemGetWebPartsByPositionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
//
//nolint:lll
func CreateItemPagesItemGetWebPartsByPositionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewItemPagesItemGetWebPartsByPositionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
//
//nolint:lll,wsl
func (m *ItemPagesItemGetWebPartsByPositionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateWebPartFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable, len(val))
			for i, v := range val {
				res[i] = v.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
//
//nolint:lll
func (m *ItemPagesItemGetWebPartsByPositionResponse) GetValue() []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable {
	return m.value
}

// Serialize serializes information the current object
//
//nolint:lll,wsl
func (m *ItemPagesItemGetWebPartsByPositionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
//
//nolint:lll
func (m *ItemPagesItemGetWebPartsByPositionResponse) SetValue(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable) {
	m.value = value
}
