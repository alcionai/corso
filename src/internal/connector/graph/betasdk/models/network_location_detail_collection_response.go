package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkLocationDetailCollectionResponse 
type NetworkLocationDetailCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []NetworkLocationDetailable
}
// NewNetworkLocationDetailCollectionResponse instantiates a new NetworkLocationDetailCollectionResponse and sets the default values.
func NewNetworkLocationDetailCollectionResponse()(*NetworkLocationDetailCollectionResponse) {
    m := &NetworkLocationDetailCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateNetworkLocationDetailCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNetworkLocationDetailCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkLocationDetailCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkLocationDetailCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNetworkLocationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkLocationDetailable, len(val))
            for i, v := range val {
                res[i] = v.(NetworkLocationDetailable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *NetworkLocationDetailCollectionResponse) GetValue()([]NetworkLocationDetailable) {
    return m.value
}
// Serialize serializes information the current object
func (m *NetworkLocationDetailCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *NetworkLocationDetailCollectionResponse) SetValue(value []NetworkLocationDetailable)() {
    m.value = value
}
