package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionPolicy 
type InformationProtectionPolicy struct {
    Entity
    // The labels property
    labels []InformationProtectionLabelable
}
// NewInformationProtectionPolicy instantiates a new informationProtectionPolicy and sets the default values.
func NewInformationProtectionPolicy()(*InformationProtectionPolicy) {
    m := &InformationProtectionPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInformationProtectionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInformationProtectionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionLabelable, len(val))
            for i, v := range val {
                res[i] = v.(InformationProtectionLabelable)
            }
            m.SetLabels(res)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. The labels property
func (m *InformationProtectionPolicy) GetLabels()([]InformationProtectionLabelable) {
    return m.labels
}
// Serialize serializes information the current object
func (m *InformationProtectionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabels sets the labels property value. The labels property
func (m *InformationProtectionPolicy) SetLabels(value []InformationProtectionLabelable)() {
    m.labels = value
}
