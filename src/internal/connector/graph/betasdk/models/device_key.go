package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceKey 
type DeviceKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceId property
    deviceId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The keyMaterial property
    keyMaterial []byte
    // The keyType property
    keyType *string
    // The OdataType property
    odataType *string
}
// NewDeviceKey instantiates a new deviceKey and sets the default values.
func NewDeviceKey()(*DeviceKey) {
    m := &DeviceKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceKey(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *DeviceKey) GetDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["keyMaterial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyMaterial(val)
        }
        return nil
    }
    res["keyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyType(val)
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
    return res
}
// GetKeyMaterial gets the keyMaterial property value. The keyMaterial property
func (m *DeviceKey) GetKeyMaterial()([]byte) {
    return m.keyMaterial
}
// GetKeyType gets the keyType property value. The keyType property
func (m *DeviceKey) GetKeyType()(*string) {
    return m.keyType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceKey) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("keyMaterial", m.GetKeyMaterial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("keyType", m.GetKeyType())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *DeviceKey) SetDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.deviceId = value
}
// SetKeyMaterial sets the keyMaterial property value. The keyMaterial property
func (m *DeviceKey) SetKeyMaterial(value []byte)() {
    m.keyMaterial = value
}
// SetKeyType sets the keyType property value. The keyType property
func (m *DeviceKey) SetKeyType(value *string)() {
    m.keyType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceKey) SetOdataType(value *string)() {
    m.odataType = value
}
