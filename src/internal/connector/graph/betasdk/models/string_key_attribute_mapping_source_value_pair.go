package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StringKeyAttributeMappingSourceValuePair 
type StringKeyAttributeMappingSourceValuePair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the parameter.
    key *string
    // The OdataType property
    odataType *string
    // The value of the parameter.
    value AttributeMappingSourceable
}
// NewStringKeyAttributeMappingSourceValuePair instantiates a new stringKeyAttributeMappingSourceValuePair and sets the default values.
func NewStringKeyAttributeMappingSourceValuePair()(*StringKeyAttributeMappingSourceValuePair) {
    m := &StringKeyAttributeMappingSourceValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStringKeyAttributeMappingSourceValuePair(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StringKeyAttributeMappingSourceValuePair) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyAttributeMappingSourceValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttributeMappingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(AttributeMappingSourceable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) GetKey()(*string) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *StringKeyAttributeMappingSourceValuePair) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The value of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) GetValue()(AttributeMappingSourceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *StringKeyAttributeMappingSourceValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
        err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *StringKeyAttributeMappingSourceValuePair) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKey sets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) SetKey(value *string)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *StringKeyAttributeMappingSourceValuePair) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The value of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) SetValue(value AttributeMappingSourceable)() {
    m.value = value
}
