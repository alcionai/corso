package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateCatalogItemCollectionResponse 
type WindowsUpdateCatalogItemCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []WindowsUpdateCatalogItemable
}
// NewWindowsUpdateCatalogItemCollectionResponse instantiates a new WindowsUpdateCatalogItemCollectionResponse and sets the default values.
func NewWindowsUpdateCatalogItemCollectionResponse()(*WindowsUpdateCatalogItemCollectionResponse) {
    m := &WindowsUpdateCatalogItemCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindowsUpdateCatalogItemCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateCatalogItemCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdateCatalogItemCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateCatalogItemCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsUpdateCatalogItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsUpdateCatalogItemable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsUpdateCatalogItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *WindowsUpdateCatalogItemCollectionResponse) GetValue()([]WindowsUpdateCatalogItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *WindowsUpdateCatalogItemCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WindowsUpdateCatalogItemCollectionResponse) SetValue(value []WindowsUpdateCatalogItemable)() {
    m.value = value
}
