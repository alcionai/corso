package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFrameworkPolicyCollectionResponse 
type TrustFrameworkPolicyCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []TrustFrameworkPolicyable
}
// NewTrustFrameworkPolicyCollectionResponse instantiates a new TrustFrameworkPolicyCollectionResponse and sets the default values.
func NewTrustFrameworkPolicyCollectionResponse()(*TrustFrameworkPolicyCollectionResponse) {
    m := &TrustFrameworkPolicyCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateTrustFrameworkPolicyCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkPolicyCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkPolicyCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkPolicyCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TrustFrameworkPolicyable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *TrustFrameworkPolicyCollectionResponse) GetValue()([]TrustFrameworkPolicyable) {
    return m.value
}
// Serialize serializes information the current object
func (m *TrustFrameworkPolicyCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TrustFrameworkPolicyCollectionResponse) SetValue(value []TrustFrameworkPolicyable)() {
    m.value = value
}
