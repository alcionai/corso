package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CorsConfiguration_v2CollectionResponse 
type CorsConfiguration_v2CollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []CorsConfiguration_v2able
}
// NewCorsConfiguration_v2CollectionResponse instantiates a new CorsConfiguration_v2CollectionResponse and sets the default values.
func NewCorsConfiguration_v2CollectionResponse()(*CorsConfiguration_v2CollectionResponse) {
    m := &CorsConfiguration_v2CollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCorsConfiguration_v2CollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCorsConfiguration_v2CollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCorsConfiguration_v2CollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CorsConfiguration_v2CollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCorsConfiguration_v2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CorsConfiguration_v2able, len(val))
            for i, v := range val {
                res[i] = v.(CorsConfiguration_v2able)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CorsConfiguration_v2CollectionResponse) GetValue()([]CorsConfiguration_v2able) {
    return m.value
}
// Serialize serializes information the current object
func (m *CorsConfiguration_v2CollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CorsConfiguration_v2CollectionResponse) SetValue(value []CorsConfiguration_v2able)() {
    m.value = value
}
