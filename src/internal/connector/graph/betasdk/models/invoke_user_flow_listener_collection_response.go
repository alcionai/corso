package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InvokeUserFlowListenerCollectionResponse 
type InvokeUserFlowListenerCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []InvokeUserFlowListenerable
}
// NewInvokeUserFlowListenerCollectionResponse instantiates a new InvokeUserFlowListenerCollectionResponse and sets the default values.
func NewInvokeUserFlowListenerCollectionResponse()(*InvokeUserFlowListenerCollectionResponse) {
    m := &InvokeUserFlowListenerCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateInvokeUserFlowListenerCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInvokeUserFlowListenerCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvokeUserFlowListenerCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InvokeUserFlowListenerCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInvokeUserFlowListenerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InvokeUserFlowListenerable, len(val))
            for i, v := range val {
                res[i] = v.(InvokeUserFlowListenerable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *InvokeUserFlowListenerCollectionResponse) GetValue()([]InvokeUserFlowListenerable) {
    return m.value
}
// Serialize serializes information the current object
func (m *InvokeUserFlowListenerCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *InvokeUserFlowListenerCollectionResponse) SetValue(value []InvokeUserFlowListenerable)() {
    m.value = value
}
