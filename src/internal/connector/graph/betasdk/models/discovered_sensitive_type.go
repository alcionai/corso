package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DiscoveredSensitiveType 
type DiscoveredSensitiveType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The classificationAttributes property
    classificationAttributes []ClassificationAttributeable
    // The confidence property
    confidence *int32
    // The count property
    count *int32
    // The id property
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The OdataType property
    odataType *string
}
// NewDiscoveredSensitiveType instantiates a new discoveredSensitiveType and sets the default values.
func NewDiscoveredSensitiveType()(*DiscoveredSensitiveType) {
    m := &DiscoveredSensitiveType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDiscoveredSensitiveTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDiscoveredSensitiveTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDiscoveredSensitiveType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DiscoveredSensitiveType) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClassificationAttributes gets the classificationAttributes property value. The classificationAttributes property
func (m *DiscoveredSensitiveType) GetClassificationAttributes()([]ClassificationAttributeable) {
    return m.classificationAttributes
}
// GetConfidence gets the confidence property value. The confidence property
func (m *DiscoveredSensitiveType) GetConfidence()(*int32) {
    return m.confidence
}
// GetCount gets the count property value. The count property
func (m *DiscoveredSensitiveType) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DiscoveredSensitiveType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classificationAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassificationAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationAttributeable, len(val))
            for i, v := range val {
                res[i] = v.(ClassificationAttributeable)
            }
            m.SetClassificationAttributes(res)
        }
        return nil
    }
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
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
func (m *DiscoveredSensitiveType) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DiscoveredSensitiveType) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DiscoveredSensitiveType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClassificationAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassificationAttributes()))
        for i, v := range m.GetClassificationAttributes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("classificationAttributes", cast)
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteUUIDValue("id", m.GetId())
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
func (m *DiscoveredSensitiveType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClassificationAttributes sets the classificationAttributes property value. The classificationAttributes property
func (m *DiscoveredSensitiveType) SetClassificationAttributes(value []ClassificationAttributeable)() {
    m.classificationAttributes = value
}
// SetConfidence sets the confidence property value. The confidence property
func (m *DiscoveredSensitiveType) SetConfidence(value *int32)() {
    m.confidence = value
}
// SetCount sets the count property value. The count property
func (m *DiscoveredSensitiveType) SetCount(value *int32)() {
    m.count = value
}
// SetId sets the id property value. The id property
func (m *DiscoveredSensitiveType) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DiscoveredSensitiveType) SetOdataType(value *string)() {
    m.odataType = value
}
