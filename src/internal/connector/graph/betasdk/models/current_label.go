package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CurrentLabel 
type CurrentLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The applicationMode property
    applicationMode *ApplicationMode
    // The id property
    id *string
    // The OdataType property
    odataType *string
}
// NewCurrentLabel instantiates a new currentLabel and sets the default values.
func NewCurrentLabel()(*CurrentLabel) {
    m := &CurrentLabel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCurrentLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCurrentLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCurrentLabel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CurrentLabel) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicationMode gets the applicationMode property value. The applicationMode property
func (m *CurrentLabel) GetApplicationMode()(*ApplicationMode) {
    return m.applicationMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CurrentLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationMode(val.(*ApplicationMode))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
// GetId gets the id property value. The id property
func (m *CurrentLabel) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CurrentLabel) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CurrentLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApplicationMode() != nil {
        cast := (*m.GetApplicationMode()).String()
        err := writer.WriteStringValue("applicationMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *CurrentLabel) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplicationMode sets the applicationMode property value. The applicationMode property
func (m *CurrentLabel) SetApplicationMode(value *ApplicationMode)() {
    m.applicationMode = value
}
// SetId sets the id property value. The id property
func (m *CurrentLabel) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CurrentLabel) SetOdataType(value *string)() {
    m.odataType = value
}
