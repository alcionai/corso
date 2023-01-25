package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFrameworkKeySet 
type TrustFrameworkKeySet struct {
    Entity
    // A collection of the keys.
    keys []TrustFrameworkKeyable
}
// NewTrustFrameworkKeySet instantiates a new TrustFrameworkKeySet and sets the default values.
func NewTrustFrameworkKeySet()(*TrustFrameworkKeySet) {
    m := &TrustFrameworkKeySet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrustFrameworkKeySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkKeySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrustFrameworkKeySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkKeySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkKeyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkKeyable, len(val))
            for i, v := range val {
                res[i] = v.(TrustFrameworkKeyable)
            }
            m.SetKeys(res)
        }
        return nil
    }
    return res
}
// GetKeys gets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) GetKeys()([]TrustFrameworkKeyable) {
    return m.keys
}
// Serialize serializes information the current object
func (m *TrustFrameworkKeySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetKeys() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeys()))
        for i, v := range m.GetKeys() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("keys", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeys sets the keys property value. A collection of the keys.
func (m *TrustFrameworkKeySet) SetKeys(value []TrustFrameworkKeyable)() {
    m.keys = value
}
