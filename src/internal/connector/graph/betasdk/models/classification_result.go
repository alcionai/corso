package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassificationResult 
type ClassificationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The confidence level, 0 to 100, of the result.
    confidenceLevel *int32
    // The number of instances of the specific information type in the input.
    count *int32
    // The OdataType property
    odataType *string
    // The GUID of the discovered sensitive information type.
    sensitiveTypeId *string
}
// NewClassificationResult instantiates a new classificationResult and sets the default values.
func NewClassificationResult()(*ClassificationResult) {
    m := &ClassificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassificationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassificationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfidenceLevel gets the confidenceLevel property value. The confidence level, 0 to 100, of the result.
func (m *ClassificationResult) GetConfidenceLevel()(*int32) {
    return m.confidenceLevel
}
// GetCount gets the count property value. The number of instances of the specific information type in the input.
func (m *ClassificationResult) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confidenceLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidenceLevel(val)
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
    res["sensitiveTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitiveTypeId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ClassificationResult) GetOdataType()(*string) {
    return m.odataType
}
// GetSensitiveTypeId gets the sensitiveTypeId property value. The GUID of the discovered sensitive information type.
func (m *ClassificationResult) GetSensitiveTypeId()(*string) {
    return m.sensitiveTypeId
}
// Serialize serializes information the current object
func (m *ClassificationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidenceLevel", m.GetConfidenceLevel())
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
        err := writer.WriteStringValue("sensitiveTypeId", m.GetSensitiveTypeId())
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
func (m *ClassificationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfidenceLevel sets the confidenceLevel property value. The confidence level, 0 to 100, of the result.
func (m *ClassificationResult) SetConfidenceLevel(value *int32)() {
    m.confidenceLevel = value
}
// SetCount sets the count property value. The number of instances of the specific information type in the input.
func (m *ClassificationResult) SetCount(value *int32)() {
    m.count = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ClassificationResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSensitiveTypeId sets the sensitiveTypeId property value. The GUID of the discovered sensitive information type.
func (m *ClassificationResult) SetSensitiveTypeId(value *string)() {
    m.sensitiveTypeId = value
}
