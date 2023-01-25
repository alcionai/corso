package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InferenceData 
type InferenceData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Confidence score reflecting the accuracy of the data inferred about the user.
    confidenceScore *float64
    // The OdataType property
    odataType *string
    // Records if the user has confirmed this inference as being True or False.
    userHasVerifiedAccuracy *bool
}
// NewInferenceData instantiates a new inferenceData and sets the default values.
func NewInferenceData()(*InferenceData) {
    m := &InferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInferenceDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInferenceDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInferenceData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InferenceData) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfidenceScore gets the confidenceScore property value. Confidence score reflecting the accuracy of the data inferred about the user.
func (m *InferenceData) GetConfidenceScore()(*float64) {
    return m.confidenceScore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InferenceData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confidenceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidenceScore(val)
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
    res["userHasVerifiedAccuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserHasVerifiedAccuracy(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InferenceData) GetOdataType()(*string) {
    return m.odataType
}
// GetUserHasVerifiedAccuracy gets the userHasVerifiedAccuracy property value. Records if the user has confirmed this inference as being True or False.
func (m *InferenceData) GetUserHasVerifiedAccuracy()(*bool) {
    return m.userHasVerifiedAccuracy
}
// Serialize serializes information the current object
func (m *InferenceData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("confidenceScore", m.GetConfidenceScore())
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
        err := writer.WriteBoolValue("userHasVerifiedAccuracy", m.GetUserHasVerifiedAccuracy())
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
func (m *InferenceData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfidenceScore sets the confidenceScore property value. Confidence score reflecting the accuracy of the data inferred about the user.
func (m *InferenceData) SetConfidenceScore(value *float64)() {
    m.confidenceScore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InferenceData) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserHasVerifiedAccuracy sets the userHasVerifiedAccuracy property value. Records if the user has confirmed this inference as being True or False.
func (m *InferenceData) SetUserHasVerifiedAccuracy(value *bool)() {
    m.userHasVerifiedAccuracy = value
}
