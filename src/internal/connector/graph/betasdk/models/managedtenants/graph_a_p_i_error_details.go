package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GraphAPIErrorDetails 
type GraphAPIErrorDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The code property
    code *string
    // The message property
    message *string
    // The OdataType property
    odataType *string
}
// NewGraphAPIErrorDetails instantiates a new graphAPIErrorDetails and sets the default values.
func NewGraphAPIErrorDetails()(*GraphAPIErrorDetails) {
    m := &GraphAPIErrorDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGraphAPIErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGraphAPIErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGraphAPIErrorDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GraphAPIErrorDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
func (m *GraphAPIErrorDetails) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GraphAPIErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
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
    return res
}
// GetMessage gets the message property value. The message property
func (m *GraphAPIErrorDetails) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GraphAPIErrorDetails) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GraphAPIErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GraphAPIErrorDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *GraphAPIErrorDetails) SetCode(value *string)() {
    m.code = value
}
// SetMessage sets the message property value. The message property
func (m *GraphAPIErrorDetails) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GraphAPIErrorDetails) SetOdataType(value *string)() {
    m.odataType = value
}
