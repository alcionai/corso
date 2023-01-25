package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFrameworkPolicy 
type TrustFrameworkPolicy struct {
    Entity
}
// NewTrustFrameworkPolicy instantiates a new TrustFrameworkPolicy and sets the default values.
func NewTrustFrameworkPolicy()(*TrustFrameworkPolicy) {
    m := &TrustFrameworkPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrustFrameworkPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TrustFrameworkPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
