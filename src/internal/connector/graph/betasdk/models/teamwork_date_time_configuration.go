package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDateTimeConfiguration 
type TeamworkDateTimeConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The date format for the device.
    dateFormat *string
    // The OdataType property
    odataType *string
    // The time of the day when the device is turned off.
    officeHoursEndTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The time of the day when the device is turned on.
    officeHoursStartTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The time format for the device.
    timeFormat *string
    // The time zone to which the office hours apply.
    timeZone *string
}
// NewTeamworkDateTimeConfiguration instantiates a new teamworkDateTimeConfiguration and sets the default values.
func NewTeamworkDateTimeConfiguration()(*TeamworkDateTimeConfiguration) {
    m := &TeamworkDateTimeConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkDateTimeConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDateTimeConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDateTimeConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDateTimeConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDateFormat gets the dateFormat property value. The date format for the device.
func (m *TeamworkDateTimeConfiguration) GetDateFormat()(*string) {
    return m.dateFormat
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDateTimeConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dateFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateFormat(val)
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
    res["officeHoursEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeHoursEndTime(val)
        }
        return nil
    }
    res["officeHoursStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeHoursStartTime(val)
        }
        return nil
    }
    res["timeFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkDateTimeConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetOfficeHoursEndTime gets the officeHoursEndTime property value. The time of the day when the device is turned off.
func (m *TeamworkDateTimeConfiguration) GetOfficeHoursEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.officeHoursEndTime
}
// GetOfficeHoursStartTime gets the officeHoursStartTime property value. The time of the day when the device is turned on.
func (m *TeamworkDateTimeConfiguration) GetOfficeHoursStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.officeHoursStartTime
}
// GetTimeFormat gets the timeFormat property value. The time format for the device.
func (m *TeamworkDateTimeConfiguration) GetTimeFormat()(*string) {
    return m.timeFormat
}
// GetTimeZone gets the timeZone property value. The time zone to which the office hours apply.
func (m *TeamworkDateTimeConfiguration) GetTimeZone()(*string) {
    return m.timeZone
}
// Serialize serializes information the current object
func (m *TeamworkDateTimeConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dateFormat", m.GetDateFormat())
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
        err := writer.WriteTimeOnlyValue("officeHoursEndTime", m.GetOfficeHoursEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("officeHoursStartTime", m.GetOfficeHoursStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeFormat", m.GetTimeFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
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
func (m *TeamworkDateTimeConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDateFormat sets the dateFormat property value. The date format for the device.
func (m *TeamworkDateTimeConfiguration) SetDateFormat(value *string)() {
    m.dateFormat = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDateTimeConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOfficeHoursEndTime sets the officeHoursEndTime property value. The time of the day when the device is turned off.
func (m *TeamworkDateTimeConfiguration) SetOfficeHoursEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.officeHoursEndTime = value
}
// SetOfficeHoursStartTime sets the officeHoursStartTime property value. The time of the day when the device is turned on.
func (m *TeamworkDateTimeConfiguration) SetOfficeHoursStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.officeHoursStartTime = value
}
// SetTimeFormat sets the timeFormat property value. The time format for the device.
func (m *TeamworkDateTimeConfiguration) SetTimeFormat(value *string)() {
    m.timeFormat = value
}
// SetTimeZone sets the timeZone property value. The time zone to which the office hours apply.
func (m *TeamworkDateTimeConfiguration) SetTimeZone(value *string)() {
    m.timeZone = value
}
