package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CoachmarkLocation 
type CoachmarkLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Length of coachmark.
    length *int32
    // The OdataType property
    odataType *string
    // Offset of coachmark.
    offset *int32
    // Type of coachmark location. The possible values are: unknown, fromEmail, subject, externalTag, displayName, messageBody, unknownFutureValue.
    type_escaped *CoachmarkLocationType
}
// NewCoachmarkLocation instantiates a new coachmarkLocation and sets the default values.
func NewCoachmarkLocation()(*CoachmarkLocation) {
    m := &CoachmarkLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCoachmarkLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCoachmarkLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCoachmarkLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CoachmarkLocation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CoachmarkLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCoachmarkLocationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*CoachmarkLocationType))
        }
        return nil
    }
    return res
}
// GetLength gets the length property value. Length of coachmark.
func (m *CoachmarkLocation) GetLength()(*int32) {
    return m.length
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CoachmarkLocation) GetOdataType()(*string) {
    return m.odataType
}
// GetOffset gets the offset property value. Offset of coachmark.
func (m *CoachmarkLocation) GetOffset()(*int32) {
    return m.offset
}
// GetType gets the type property value. Type of coachmark location. The possible values are: unknown, fromEmail, subject, externalTag, displayName, messageBody, unknownFutureValue.
func (m *CoachmarkLocation) GetType()(*CoachmarkLocationType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *CoachmarkLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("length", m.GetLength())
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *CoachmarkLocation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLength sets the length property value. Length of coachmark.
func (m *CoachmarkLocation) SetLength(value *int32)() {
    m.length = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CoachmarkLocation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffset sets the offset property value. Offset of coachmark.
func (m *CoachmarkLocation) SetOffset(value *int32)() {
    m.offset = value
}
// SetType sets the type property value. Type of coachmark location. The possible values are: unknown, fromEmail, subject, externalTag, displayName, messageBody, unknownFutureValue.
func (m *CoachmarkLocation) SetType(value *CoachmarkLocationType)() {
    m.type_escaped = value
}
