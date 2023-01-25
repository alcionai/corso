package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// AzureADDeviceCollectionResponse 
type AzureADDeviceCollectionResponse struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.BaseCollectionPaginationCountResponse
    // The value property
    value []AzureADDeviceable
}
// NewAzureADDeviceCollectionResponse instantiates a new AzureADDeviceCollectionResponse and sets the default values.
func NewAzureADDeviceCollectionResponse()(*AzureADDeviceCollectionResponse) {
    m := &AzureADDeviceCollectionResponse{
        BaseCollectionPaginationCountResponse: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAzureADDeviceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADDeviceCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADDeviceCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADDeviceCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAzureADDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AzureADDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(AzureADDeviceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AzureADDeviceCollectionResponse) GetValue()([]AzureADDeviceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AzureADDeviceCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AzureADDeviceCollectionResponse) SetValue(value []AzureADDeviceable)() {
    m.value = value
}
