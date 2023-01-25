package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NumberRange number Range definition.
type NumberRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Lower number.
    lowerNumber *int32
    // The OdataType property
    odataType *string
    // Upper number.
    upperNumber *int32
}
// NewNumberRange instantiates a new numberRange and sets the default values.
func NewNumberRange()(*NumberRange) {
    m := &NumberRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateNumberRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNumberRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNumberRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NumberRange) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NumberRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lowerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowerNumber(val)
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
    res["upperNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpperNumber(val)
        }
        return nil
    }
    return res
}
// GetLowerNumber gets the lowerNumber property value. Lower number.
func (m *NumberRange) GetLowerNumber()(*int32) {
    return m.lowerNumber
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NumberRange) GetOdataType()(*string) {
    return m.odataType
}
// GetUpperNumber gets the upperNumber property value. Upper number.
func (m *NumberRange) GetUpperNumber()(*int32) {
    return m.upperNumber
}
// Serialize serializes information the current object
func (m *NumberRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("lowerNumber", m.GetLowerNumber())
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
        err := writer.WriteInt32Value("upperNumber", m.GetUpperNumber())
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
func (m *NumberRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLowerNumber sets the lowerNumber property value. Lower number.
func (m *NumberRange) SetLowerNumber(value *int32)() {
    m.lowerNumber = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NumberRange) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUpperNumber sets the upperNumber property value. Upper number.
func (m *NumberRange) SetUpperNumber(value *int32)() {
    m.upperNumber = value
}
