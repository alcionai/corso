package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// WebPartCollectionResponse provides operations to manage the webParts property of the microsoft.graph.sitePage entity.
type WebPartCollectionResponse struct {
	msmodel.BaseCollectionPaginationCountResponse
	// The value property
	value []WebPartable
}

// NewWebPartCollectionResponse instantiates a new WebPartCollectionResponse and sets the default values.
func NewWebPartCollectionResponse() *WebPartCollectionResponse {
	m := &WebPartCollectionResponse{
		BaseCollectionPaginationCountResponse: *msmodel.NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateWebPartCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebPartCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewWebPartCollectionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *WebPartCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateWebPartFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]WebPartable, len(val))
			for i, v := range val {
				res[i] = v.(WebPartable)
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *WebPartCollectionResponse) GetValue() []WebPartable {
	return m.value
}

// Serialize serializes information the current object
func (m *WebPartCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *WebPartCollectionResponse) SetValue(value []WebPartable) {
	m.value = value
}
