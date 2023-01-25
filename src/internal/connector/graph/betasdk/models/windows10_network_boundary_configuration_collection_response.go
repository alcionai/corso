package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10NetworkBoundaryConfigurationCollectionResponse 
type Windows10NetworkBoundaryConfigurationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Windows10NetworkBoundaryConfigurationable
}
// NewWindows10NetworkBoundaryConfigurationCollectionResponse instantiates a new Windows10NetworkBoundaryConfigurationCollectionResponse and sets the default values.
func NewWindows10NetworkBoundaryConfigurationCollectionResponse()(*Windows10NetworkBoundaryConfigurationCollectionResponse) {
    m := &Windows10NetworkBoundaryConfigurationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindows10NetworkBoundaryConfigurationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10NetworkBoundaryConfigurationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10NetworkBoundaryConfigurationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10NetworkBoundaryConfigurationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10NetworkBoundaryConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10NetworkBoundaryConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10NetworkBoundaryConfigurationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Windows10NetworkBoundaryConfigurationCollectionResponse) GetValue()([]Windows10NetworkBoundaryConfigurationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Windows10NetworkBoundaryConfigurationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10NetworkBoundaryConfigurationCollectionResponse) SetValue(value []Windows10NetworkBoundaryConfigurationable)() {
    m.value = value
}
