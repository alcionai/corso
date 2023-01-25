package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagementEligibleDeviceCollectionResponse 
type ComanagementEligibleDeviceCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ComanagementEligibleDeviceable
}
// NewComanagementEligibleDeviceCollectionResponse instantiates a new ComanagementEligibleDeviceCollectionResponse and sets the default values.
func NewComanagementEligibleDeviceCollectionResponse()(*ComanagementEligibleDeviceCollectionResponse) {
    m := &ComanagementEligibleDeviceCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateComanagementEligibleDeviceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagementEligibleDeviceCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagementEligibleDeviceCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagementEligibleDeviceCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComanagementEligibleDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComanagementEligibleDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(ComanagementEligibleDeviceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ComanagementEligibleDeviceCollectionResponse) GetValue()([]ComanagementEligibleDeviceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ComanagementEligibleDeviceCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ComanagementEligibleDeviceCollectionResponse) SetValue(value []ComanagementEligibleDeviceable)() {
    m.value = value
}
