package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentClassification 
type ContentClassification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The confidence property
    confidence *int32
    // The matches property
    matches []MatchLocationable
    // The OdataType property
    odataType *string
    // The sensitiveTypeId property
    sensitiveTypeId *string
    // The uniqueCount property
    uniqueCount *int32
}
// NewContentClassification instantiates a new contentClassification and sets the default values.
func NewContentClassification()(*ContentClassification) {
    m := &ContentClassification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateContentClassificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentClassificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentClassification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentClassification) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfidence gets the confidence property value. The confidence property
func (m *ContentClassification) GetConfidence()(*int32) {
    return m.confidence
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentClassification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMatchLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MatchLocationable, len(val))
            for i, v := range val {
                res[i] = v.(MatchLocationable)
            }
            m.SetMatches(res)
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
    res["uniqueCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueCount(val)
        }
        return nil
    }
    return res
}
// GetMatches gets the matches property value. The matches property
func (m *ContentClassification) GetMatches()([]MatchLocationable) {
    return m.matches
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentClassification) GetOdataType()(*string) {
    return m.odataType
}
// GetSensitiveTypeId gets the sensitiveTypeId property value. The sensitiveTypeId property
func (m *ContentClassification) GetSensitiveTypeId()(*string) {
    return m.sensitiveTypeId
}
// GetUniqueCount gets the uniqueCount property value. The uniqueCount property
func (m *ContentClassification) GetUniqueCount()(*int32) {
    return m.uniqueCount
}
// Serialize serializes information the current object
func (m *ContentClassification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    if m.GetMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("matches", cast)
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
        err := writer.WriteInt32Value("uniqueCount", m.GetUniqueCount())
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
func (m *ContentClassification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfidence sets the confidence property value. The confidence property
func (m *ContentClassification) SetConfidence(value *int32)() {
    m.confidence = value
}
// SetMatches sets the matches property value. The matches property
func (m *ContentClassification) SetMatches(value []MatchLocationable)() {
    m.matches = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentClassification) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSensitiveTypeId sets the sensitiveTypeId property value. The sensitiveTypeId property
func (m *ContentClassification) SetSensitiveTypeId(value *string)() {
    m.sensitiveTypeId = value
}
// SetUniqueCount sets the uniqueCount property value. The uniqueCount property
func (m *ContentClassification) SetUniqueCount(value *int32)() {
    m.uniqueCount = value
}
