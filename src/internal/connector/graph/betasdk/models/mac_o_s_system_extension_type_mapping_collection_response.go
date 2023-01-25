package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSystemExtensionTypeMappingCollectionResponse 
type MacOSSystemExtensionTypeMappingCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []MacOSSystemExtensionTypeMappingable
}
// NewMacOSSystemExtensionTypeMappingCollectionResponse instantiates a new MacOSSystemExtensionTypeMappingCollectionResponse and sets the default values.
func NewMacOSSystemExtensionTypeMappingCollectionResponse()(*MacOSSystemExtensionTypeMappingCollectionResponse) {
    m := &MacOSSystemExtensionTypeMappingCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMacOSSystemExtensionTypeMappingCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSystemExtensionTypeMappingCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSystemExtensionTypeMappingCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSystemExtensionTypeMappingCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSSystemExtensionTypeMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSystemExtensionTypeMappingable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSSystemExtensionTypeMappingable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MacOSSystemExtensionTypeMappingCollectionResponse) GetValue()([]MacOSSystemExtensionTypeMappingable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MacOSSystemExtensionTypeMappingCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MacOSSystemExtensionTypeMappingCollectionResponse) SetValue(value []MacOSSystemExtensionTypeMappingable)() {
    m.value = value
}
