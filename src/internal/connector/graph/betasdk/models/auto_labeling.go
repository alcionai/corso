package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AutoLabeling 
type AutoLabeling struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The message property
    message *string
    // The OdataType property
    odataType *string
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
}
// NewAutoLabeling instantiates a new autoLabeling and sets the default values.
func NewAutoLabeling()(*AutoLabeling) {
    m := &AutoLabeling{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAutoLabelingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAutoLabelingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAutoLabeling(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutoLabeling) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AutoLabeling) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["sensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSensitiveTypeIds(res)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *AutoLabeling) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AutoLabeling) GetOdataType()(*string) {
    return m.odataType
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *AutoLabeling) GetSensitiveTypeIds()([]string) {
    return m.sensitiveTypeIds
}
// Serialize serializes information the current object
func (m *AutoLabeling) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSensitiveTypeIds() != nil {
        err := writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
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
func (m *AutoLabeling) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMessage sets the message property value. The message property
func (m *AutoLabeling) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AutoLabeling) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *AutoLabeling) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
