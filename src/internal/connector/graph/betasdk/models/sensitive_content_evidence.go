package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitiveContentEvidence 
type SensitiveContentEvidence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The length property
    length *int32
    // The match property
    match *string
    // The OdataType property
    odataType *string
    // The offset property
    offset *int32
}
// NewSensitiveContentEvidence instantiates a new sensitiveContentEvidence and sets the default values.
func NewSensitiveContentEvidence()(*SensitiveContentEvidence) {
    m := &SensitiveContentEvidence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSensitiveContentEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitiveContentEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitiveContentEvidence(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveContentEvidence) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitiveContentEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
        }
        return nil
    }
    res["match"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatch(val)
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
    res["offset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
        }
        return nil
    }
    return res
}
// GetLength gets the length property value. The length property
func (m *SensitiveContentEvidence) GetLength()(*int32) {
    return m.length
}
// GetMatch gets the match property value. The match property
func (m *SensitiveContentEvidence) GetMatch()(*string) {
    return m.match
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SensitiveContentEvidence) GetOdataType()(*string) {
    return m.odataType
}
// GetOffset gets the offset property value. The offset property
func (m *SensitiveContentEvidence) GetOffset()(*int32) {
    return m.offset
}
// Serialize serializes information the current object
func (m *SensitiveContentEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("length", m.GetLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("match", m.GetMatch())
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
        err := writer.WriteInt32Value("offset", m.GetOffset())
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
func (m *SensitiveContentEvidence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLength sets the length property value. The length property
func (m *SensitiveContentEvidence) SetLength(value *int32)() {
    m.length = value
}
// SetMatch sets the match property value. The match property
func (m *SensitiveContentEvidence) SetMatch(value *string)() {
    m.match = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SensitiveContentEvidence) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffset sets the offset property value. The offset property
func (m *SensitiveContentEvidence) SetOffset(value *int32)() {
    m.offset = value
}
