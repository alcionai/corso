package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCardEntry 
type TimeCardEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of breaks associated with the timeCard.
    breaks []TimeCardBreakable
    // The clock-in event of the timeCard.
    clockInEvent TimeCardEventable
    // The clock-out event of the timeCard.
    clockOutEvent TimeCardEventable
    // The OdataType property
    odataType *string
}
// NewTimeCardEntry instantiates a new timeCardEntry and sets the default values.
func NewTimeCardEntry()(*TimeCardEntry) {
    m := &TimeCardEntry{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTimeCardEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeCardEntry(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEntry) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBreaks gets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCardEntry) GetBreaks()([]TimeCardBreakable) {
    return m.breaks
}
// GetClockInEvent gets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCardEntry) GetClockInEvent()(TimeCardEventable) {
    return m.clockInEvent
}
// GetClockOutEvent gets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCardEntry) GetClockOutEvent()(TimeCardEventable) {
    return m.clockOutEvent
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCardEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["breaks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeCardBreakFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeCardBreakable, len(val))
            for i, v := range val {
                res[i] = v.(TimeCardBreakable)
            }
            m.SetBreaks(res)
        }
        return nil
    }
    res["clockInEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockInEvent(val.(TimeCardEventable))
        }
        return nil
    }
    res["clockOutEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockOutEvent(val.(TimeCardEventable))
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
func (m *TimeCardEntry) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TimeCardEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBreaks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBreaks()))
        for i, v := range m.GetBreaks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("breaks", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clockInEvent", m.GetClockInEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clockOutEvent", m.GetClockOutEvent())
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
func (m *TimeCardEntry) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBreaks sets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCardEntry) SetBreaks(value []TimeCardBreakable)() {
    m.breaks = value
}
// SetClockInEvent sets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCardEntry) SetClockInEvent(value TimeCardEventable)() {
    m.clockInEvent = value
}
// SetClockOutEvent sets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCardEntry) SetClockOutEvent(value TimeCardEventable)() {
    m.clockOutEvent = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimeCardEntry) SetOdataType(value *string)() {
    m.odataType = value
}
