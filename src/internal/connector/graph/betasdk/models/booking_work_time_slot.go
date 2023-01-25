package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingWorkTimeSlot 
type BookingWorkTimeSlot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The time of the day when work stops. For example, 17:00:00.0000000.
    end *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The OdataType property
    odataType *string
    // The time of the day when work starts. For example, 08:00:00.0000000.
    start *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
}
// NewBookingWorkTimeSlot instantiates a new bookingWorkTimeSlot and sets the default values.
func NewBookingWorkTimeSlot()(*BookingWorkTimeSlot) {
    m := &BookingWorkTimeSlot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBookingWorkTimeSlotFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingWorkTimeSlotFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingWorkTimeSlot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkTimeSlot) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnd gets the end property value. The time of the day when work stops. For example, 17:00:00.0000000.
func (m *BookingWorkTimeSlot) GetEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.end
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingWorkTimeSlot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
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
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BookingWorkTimeSlot) GetOdataType()(*string) {
    return m.odataType
}
// GetStart gets the start property value. The time of the day when work starts. For example, 08:00:00.0000000.
func (m *BookingWorkTimeSlot) GetStart()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.start
}
// Serialize serializes information the current object
func (m *BookingWorkTimeSlot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeOnlyValue("end", m.GetEnd())
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
        err := writer.WriteTimeOnlyValue("start", m.GetStart())
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
func (m *BookingWorkTimeSlot) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnd sets the end property value. The time of the day when work stops. For example, 17:00:00.0000000.
func (m *BookingWorkTimeSlot) SetEnd(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.end = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingWorkTimeSlot) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStart sets the start property value. The time of the day when work starts. For example, 08:00:00.0000000.
func (m *BookingWorkTimeSlot) SetStart(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.start = value
}
