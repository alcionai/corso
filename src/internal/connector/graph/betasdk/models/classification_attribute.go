package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassificationAttribute 
type ClassificationAttribute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The confidence property
    confidence *int32
    // The count property
    count *int32
    // The OdataType property
    odataType *string
}
// NewClassificationAttribute instantiates a new classificationAttribute and sets the default values.
func NewClassificationAttribute()(*ClassificationAttribute) {
    m := &ClassificationAttribute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassificationAttributeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationAttributeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassificationAttribute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationAttribute) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfidence gets the confidence property value. The confidence property
func (m *ClassificationAttribute) GetConfidence()(*int32) {
    return m.confidence
}
// GetCount gets the count property value. The count property
func (m *ClassificationAttribute) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationAttribute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confidence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidence(val)
        }
        return nil
    }
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ClassificationAttribute) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ClassificationAttribute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("count", m.GetCount())
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
func (m *ClassificationAttribute) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfidence sets the confidence property value. The confidence property
func (m *ClassificationAttribute) SetConfidence(value *int32)() {
    m.confidence = value
}
// SetCount sets the count property value. The count property
func (m *ClassificationAttribute) SetCount(value *int32)() {
    m.count = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ClassificationAttribute) SetOdataType(value *string)() {
    m.odataType = value
}
