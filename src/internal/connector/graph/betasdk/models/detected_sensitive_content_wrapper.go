package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedSensitiveContentWrapper 
type DetectedSensitiveContentWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The classification property
    classification []DetectedSensitiveContentable
    // The OdataType property
    odataType *string
}
// NewDetectedSensitiveContentWrapper instantiates a new detectedSensitiveContentWrapper and sets the default values.
func NewDetectedSensitiveContentWrapper()(*DetectedSensitiveContentWrapper) {
    m := &DetectedSensitiveContentWrapper{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDetectedSensitiveContentWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDetectedSensitiveContentWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDetectedSensitiveContentWrapper(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetectedSensitiveContentWrapper) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClassification gets the classification property value. The classification property
func (m *DetectedSensitiveContentWrapper) GetClassification()([]DetectedSensitiveContentable) {
    return m.classification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedSensitiveContentWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDetectedSensitiveContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DetectedSensitiveContentable, len(val))
            for i, v := range val {
                res[i] = v.(DetectedSensitiveContentable)
            }
            m.SetClassification(res)
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
func (m *DetectedSensitiveContentWrapper) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DetectedSensitiveContentWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClassification() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassification()))
        for i, v := range m.GetClassification() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("classification", cast)
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
func (m *DetectedSensitiveContentWrapper) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClassification sets the classification property value. The classification property
func (m *DetectedSensitiveContentWrapper) SetClassification(value []DetectedSensitiveContentable)() {
    m.classification = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DetectedSensitiveContentWrapper) SetOdataType(value *string)() {
    m.odataType = value
}
