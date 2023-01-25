package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeClockSettings 
type TimeClockSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The aprroved location of the timeClock.
    approvedLocation GeoCoordinatesable
    // The OdataType property
    odataType *string
}
// NewTimeClockSettings instantiates a new timeClockSettings and sets the default values.
func NewTimeClockSettings()(*TimeClockSettings) {
    m := &TimeClockSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTimeClockSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeClockSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeClockSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeClockSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApprovedLocation gets the approvedLocation property value. The aprroved location of the timeClock.
func (m *TimeClockSettings) GetApprovedLocation()(GeoCoordinatesable) {
    return m.approvedLocation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeClockSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["approvedLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovedLocation(val.(GeoCoordinatesable))
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
func (m *TimeClockSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TimeClockSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("approvedLocation", m.GetApprovedLocation())
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
func (m *TimeClockSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApprovedLocation sets the approvedLocation property value. The aprroved location of the timeClock.
func (m *TimeClockSettings) SetApprovedLocation(value GeoCoordinatesable)() {
    m.approvedLocation = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimeClockSettings) SetOdataType(value *string)() {
    m.odataType = value
}
