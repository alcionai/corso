package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkPeripheral 
type TeamworkPeripheral struct {
    Entity
    // Display name for the peripheral.
    displayName *string
    // The product ID of the device. Each product from a vendor has its own ID.
    productId *string
    // The unique identifier for the vendor of the device. Each vendor has a unique ID.
    vendorId *string
}
// NewTeamworkPeripheral instantiates a new TeamworkPeripheral and sets the default values.
func NewTeamworkPeripheral()(*TeamworkPeripheral) {
    m := &TeamworkPeripheral{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkPeripheralFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkPeripheralFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkPeripheral(), nil
}
// GetDisplayName gets the displayName property value. Display name for the peripheral.
func (m *TeamworkPeripheral) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheral) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["productId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductId(val)
        }
        return nil
    }
    res["vendorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorId(val)
        }
        return nil
    }
    return res
}
// GetProductId gets the productId property value. The product ID of the device. Each product from a vendor has its own ID.
func (m *TeamworkPeripheral) GetProductId()(*string) {
    return m.productId
}
// GetVendorId gets the vendorId property value. The unique identifier for the vendor of the device. Each vendor has a unique ID.
func (m *TeamworkPeripheral) GetVendorId()(*string) {
    return m.vendorId
}
// Serialize serializes information the current object
func (m *TeamworkPeripheral) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productId", m.GetProductId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorId", m.GetVendorId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Display name for the peripheral.
func (m *TeamworkPeripheral) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetProductId sets the productId property value. The product ID of the device. Each product from a vendor has its own ID.
func (m *TeamworkPeripheral) SetProductId(value *string)() {
    m.productId = value
}
// SetVendorId sets the vendorId property value. The unique identifier for the vendor of the device. Each vendor has a unique ID.
func (m *TeamworkPeripheral) SetVendorId(value *string)() {
    m.vendorId = value
}
