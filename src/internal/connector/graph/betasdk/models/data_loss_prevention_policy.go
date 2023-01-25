package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataLossPreventionPolicy provides operations to call the remove method.
type DataLossPreventionPolicy struct {
    Entity
    // The name property
    name *string
}
// NewDataLossPreventionPolicy instantiates a new dataLossPreventionPolicy and sets the default values.
func NewDataLossPreventionPolicy()(*DataLossPreventionPolicy) {
    m := &DataLossPreventionPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDataLossPreventionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataLossPreventionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataLossPreventionPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataLossPreventionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *DataLossPreventionPolicy) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *DataLossPreventionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name property
func (m *DataLossPreventionPolicy) SetName(value *string)() {
    m.name = value
}
