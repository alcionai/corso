package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassifcationErrorBase 
type ClassifcationErrorBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The code property
    code *string
    // The innerError property
    innerError ClassificationInnerErrorable
    // The message property
    message *string
    // The OdataType property
    odataType *string
    // The target property
    target *string
}
// NewClassifcationErrorBase instantiates a new classifcationErrorBase and sets the default values.
func NewClassifcationErrorBase()(*ClassifcationErrorBase) {
    m := &ClassifcationErrorBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassifcationErrorBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassifcationErrorBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.classificationError":
                        return NewClassificationError(), nil
                }
            }
        }
    }
    return NewClassifcationErrorBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifcationErrorBase) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
func (m *ClassifcationErrorBase) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifcationErrorBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["innerError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClassificationInnerErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerError(val.(ClassificationInnerErrorable))
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
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetInnerError gets the innerError property value. The innerError property
func (m *ClassifcationErrorBase) GetInnerError()(ClassificationInnerErrorable) {
    return m.innerError
}
// GetMessage gets the message property value. The message property
func (m *ClassifcationErrorBase) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ClassifcationErrorBase) GetOdataType()(*string) {
    return m.odataType
}
// GetTarget gets the target property value. The target property
func (m *ClassifcationErrorBase) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *ClassifcationErrorBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("innerError", m.GetInnerError())
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
        err := writer.WriteStringValue("target", m.GetTarget())
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
func (m *ClassifcationErrorBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *ClassifcationErrorBase) SetCode(value *string)() {
    m.code = value
}
// SetInnerError sets the innerError property value. The innerError property
func (m *ClassifcationErrorBase) SetInnerError(value ClassificationInnerErrorable)() {
    m.innerError = value
}
// SetMessage sets the message property value. The message property
func (m *ClassifcationErrorBase) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ClassifcationErrorBase) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTarget sets the target property value. The target property
func (m *ClassifcationErrorBase) SetTarget(value *string)() {
    m.target = value
}
