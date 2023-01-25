package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnsupportedDeviceConfigurationDetail a description of why an entity is unsupported.
type UnsupportedDeviceConfigurationDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A message explaining why an entity is unsupported.
    message *string
    // The OdataType property
    odataType *string
    // If message is related to a specific property in the original entity, then the name of that property.
    propertyName *string
}
// NewUnsupportedDeviceConfigurationDetail instantiates a new unsupportedDeviceConfigurationDetail and sets the default values.
func NewUnsupportedDeviceConfigurationDetail()(*UnsupportedDeviceConfigurationDetail) {
    m := &UnsupportedDeviceConfigurationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnsupportedDeviceConfigurationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnsupportedDeviceConfigurationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnsupportedDeviceConfigurationDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnsupportedDeviceConfigurationDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnsupportedDeviceConfigurationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
    res["propertyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyName(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. A message explaining why an entity is unsupported.
func (m *UnsupportedDeviceConfigurationDetail) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UnsupportedDeviceConfigurationDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetPropertyName gets the propertyName property value. If message is related to a specific property in the original entity, then the name of that property.
func (m *UnsupportedDeviceConfigurationDetail) GetPropertyName()(*string) {
    return m.propertyName
}
// Serialize serializes information the current object
func (m *UnsupportedDeviceConfigurationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
        err := writer.WriteStringValue("propertyName", m.GetPropertyName())
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
func (m *UnsupportedDeviceConfigurationDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMessage sets the message property value. A message explaining why an entity is unsupported.
func (m *UnsupportedDeviceConfigurationDetail) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UnsupportedDeviceConfigurationDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPropertyName sets the propertyName property value. If message is related to a specific property in the original entity, then the name of that property.
func (m *UnsupportedDeviceConfigurationDetail) SetPropertyName(value *string)() {
    m.propertyName = value
}
