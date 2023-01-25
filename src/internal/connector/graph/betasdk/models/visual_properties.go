package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VisualProperties 
type VisualProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The body of a visual user notification. Body is optional.
    body *string
    // The OdataType property
    odataType *string
    // The title of a visual user notification. This field is required for visual notification payloads.
    title *string
}
// NewVisualProperties instantiates a new visualProperties and sets the default values.
func NewVisualProperties()(*VisualProperties) {
    m := &VisualProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVisualPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVisualPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVisualProperties(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VisualProperties) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBody gets the body property value. The body of a visual user notification. Body is optional.
func (m *VisualProperties) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VisualProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VisualProperties) GetOdataType()(*string) {
    return m.odataType
}
// GetTitle gets the title property value. The title of a visual user notification. This field is required for visual notification payloads.
func (m *VisualProperties) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *VisualProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
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
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *VisualProperties) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBody sets the body property value. The body of a visual user notification. Body is optional.
func (m *VisualProperties) SetBody(value *string)() {
    m.body = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VisualProperties) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTitle sets the title property value. The title of a visual user notification. This field is required for visual notification payloads.
func (m *VisualProperties) SetTitle(value *string)() {
    m.title = value
}
