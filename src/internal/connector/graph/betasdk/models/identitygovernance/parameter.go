package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Parameter 
type Parameter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the parameter.
    name *string
    // The OdataType property
    odataType *string
    // The values of the parameter.
    values []string
    // The valueType property
    valueType *ValueType
}
// NewParameter instantiates a new parameter and sets the default values.
func NewParameter()(*Parameter) {
    m := &Parameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParameter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Parameter) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Parameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValues(res)
        }
        return nil
    }
    res["valueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*ValueType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the parameter.
func (m *Parameter) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Parameter) GetOdataType()(*string) {
    return m.odataType
}
// GetValues gets the values property value. The values of the parameter.
func (m *Parameter) GetValues()([]string) {
    return m.values
}
// GetValueType gets the valueType property value. The valueType property
func (m *Parameter) GetValueType()(*ValueType) {
    return m.valueType
}
// Serialize serializes information the current object
func (m *Parameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetValues() != nil {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := (*m.GetValueType()).String()
        err := writer.WriteStringValue("valueType", &cast)
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
func (m *Parameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. The name of the parameter.
func (m *Parameter) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Parameter) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValues sets the values property value. The values of the parameter.
func (m *Parameter) SetValues(value []string)() {
    m.values = value
}
// SetValueType sets the valueType property value. The valueType property
func (m *Parameter) SetValueType(value *ValueType)() {
    m.valueType = value
}
