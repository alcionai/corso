package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MatchLocation 
type MatchLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The length property
    length *int32
    // The OdataType property
    odataType *string
    // The offset property
    offset *int32
}
// NewMatchLocation instantiates a new matchLocation and sets the default values.
func NewMatchLocation()(*MatchLocation) {
    m := &MatchLocation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMatchLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMatchLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMatchLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchLocation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MatchLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetLength gets the length property value. The length property
func (m *MatchLocation) GetLength()(*int32) {
    return m.length
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MatchLocation) GetOdataType()(*string) {
    return m.odataType
}
// GetOffset gets the offset property value. The offset property
func (m *MatchLocation) GetOffset()(*int32) {
    return m.offset
}
// Serialize serializes information the current object
func (m *MatchLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchLocation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLength sets the length property value. The length property
func (m *MatchLocation) SetLength(value *int32)() {
    m.length = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MatchLocation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffset sets the offset property value. The offset property
func (m *MatchLocation) SetOffset(value *int32)() {
    m.offset = value
}
