package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingParameterSchemaCollectionResponse 
type AttributeMappingParameterSchemaCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []AttributeMappingParameterSchemaable
}
// NewAttributeMappingParameterSchemaCollectionResponse instantiates a new AttributeMappingParameterSchemaCollectionResponse and sets the default values.
func NewAttributeMappingParameterSchemaCollectionResponse()(*AttributeMappingParameterSchemaCollectionResponse) {
    m := &AttributeMappingParameterSchemaCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAttributeMappingParameterSchemaCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingParameterSchemaCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingParameterSchemaCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingParameterSchemaCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeMappingParameterSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeMappingParameterSchemaable, len(val))
            for i, v := range val {
                res[i] = v.(AttributeMappingParameterSchemaable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AttributeMappingParameterSchemaCollectionResponse) GetValue()([]AttributeMappingParameterSchemaable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AttributeMappingParameterSchemaCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AttributeMappingParameterSchemaCollectionResponse) SetValue(value []AttributeMappingParameterSchemaable)() {
    m.value = value
}
