package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// B2cIdentityUserFlowCollectionResponse 
type B2cIdentityUserFlowCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []B2cIdentityUserFlowable
}
// NewB2cIdentityUserFlowCollectionResponse instantiates a new B2cIdentityUserFlowCollectionResponse and sets the default values.
func NewB2cIdentityUserFlowCollectionResponse()(*B2cIdentityUserFlowCollectionResponse) {
    m := &B2cIdentityUserFlowCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateB2cIdentityUserFlowCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cIdentityUserFlowCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cIdentityUserFlowCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cIdentityUserFlowCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateB2cIdentityUserFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]B2cIdentityUserFlowable, len(val))
            for i, v := range val {
                res[i] = v.(B2cIdentityUserFlowable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *B2cIdentityUserFlowCollectionResponse) GetValue()([]B2cIdentityUserFlowable) {
    return m.value
}
// Serialize serializes information the current object
func (m *B2cIdentityUserFlowCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *B2cIdentityUserFlowCollectionResponse) SetValue(value []B2cIdentityUserFlowable)() {
    m.value = value
}
