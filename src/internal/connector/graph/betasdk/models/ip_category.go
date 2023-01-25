package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpCategory 
type IpCategory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The description property
    description *string
    // The name property
    name *string
    // The OdataType property
    odataType *string
    // The vendor property
    vendor_escaped *string
}
// NewIpCategory instantiates a new ipCategory and sets the default values.
func NewIpCategory()(*IpCategory) {
    m := &IpCategory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIpCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpCategory(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IpCategory) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
func (m *IpCategory) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *IpCategory) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IpCategory) GetOdataType()(*string) {
    return m.odataType
}
// GetVendor gets the vendor property value. The vendor property
func (m *IpCategory) GetVendor()(*string) {
    return m.vendor_escaped
}
// Serialize serializes information the current object
func (m *IpCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendor", m.GetVendor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IpCategory) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *IpCategory) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name property
func (m *IpCategory) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IpCategory) SetOdataType(value *string)() {
    m.odataType = value
}
// SetVendor sets the vendor property value. The vendor property
func (m *IpCategory) SetVendor(value *string)() {
    m.vendor_escaped = value
}
