package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OemWarrantyInformationOnboarding warranty status entity for a given OEM
type OemWarrantyInformationOnboarding struct {
    Entity
    // Specifies whether warranty API is available. This property is read-only.
    available *bool
    // Specifies whether warranty query is enabled for given OEM. This property is read-only.
    enabled *bool
    // OEM name. This property is read-only.
    oemName *string
}
// NewOemWarrantyInformationOnboarding instantiates a new oemWarrantyInformationOnboarding and sets the default values.
func NewOemWarrantyInformationOnboarding()(*OemWarrantyInformationOnboarding) {
    m := &OemWarrantyInformationOnboarding{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOemWarrantyInformationOnboardingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOemWarrantyInformationOnboardingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOemWarrantyInformationOnboarding(), nil
}
// GetAvailable gets the available property value. Specifies whether warranty API is available. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetAvailable()(*bool) {
    return m.available
}
// GetEnabled gets the enabled property value. Specifies whether warranty query is enabled for given OEM. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OemWarrantyInformationOnboarding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["available"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailable(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["oemName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOemName(val)
        }
        return nil
    }
    return res
}
// GetOemName gets the oemName property value. OEM name. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetOemName()(*string) {
    return m.oemName
}
// Serialize serializes information the current object
func (m *OemWarrantyInformationOnboarding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetAvailable sets the available property value. Specifies whether warranty API is available. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetAvailable(value *bool)() {
    m.available = value
}
// SetEnabled sets the enabled property value. Specifies whether warranty query is enabled for given OEM. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetOemName sets the oemName property value. OEM name. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetOemName(value *string)() {
    m.oemName = value
}
