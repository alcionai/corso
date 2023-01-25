package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10AppsForceUpdateSchedule windows 10 force update schedule for Apps
type Windows10AppsForceUpdateSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Possible values for App update on Windows10 recurrence.
    recurrence *Windows10AppsUpdateRecurrence
    // If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
    runImmediatelyIfAfterStartDateTime *bool
    // The start time for the force restart.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewWindows10AppsForceUpdateSchedule instantiates a new windows10AppsForceUpdateSchedule and sets the default values.
func NewWindows10AppsForceUpdateSchedule()(*Windows10AppsForceUpdateSchedule) {
    m := &Windows10AppsForceUpdateSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindows10AppsForceUpdateScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10AppsForceUpdateScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10AppsForceUpdateSchedule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10AppsForceUpdateSchedule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10AppsForceUpdateSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10AppsUpdateRecurrence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(*Windows10AppsUpdateRecurrence))
        }
        return nil
    }
    res["runImmediatelyIfAfterStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunImmediatelyIfAfterStartDateTime(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Windows10AppsForceUpdateSchedule) GetOdataType()(*string) {
    return m.odataType
}
// GetRecurrence gets the recurrence property value. Possible values for App update on Windows10 recurrence.
func (m *Windows10AppsForceUpdateSchedule) GetRecurrence()(*Windows10AppsUpdateRecurrence) {
    return m.recurrence
}
// GetRunImmediatelyIfAfterStartDateTime gets the runImmediatelyIfAfterStartDateTime property value. If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
func (m *Windows10AppsForceUpdateSchedule) GetRunImmediatelyIfAfterStartDateTime()(*bool) {
    return m.runImmediatelyIfAfterStartDateTime
}
// GetStartDateTime gets the startDateTime property value. The start time for the force restart.
func (m *Windows10AppsForceUpdateSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *Windows10AppsForceUpdateSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRecurrence() != nil {
        cast := (*m.GetRecurrence()).String()
        err := writer.WriteStringValue("recurrence", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("runImmediatelyIfAfterStartDateTime", m.GetRunImmediatelyIfAfterStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
func (m *Windows10AppsForceUpdateSchedule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Windows10AppsForceUpdateSchedule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecurrence sets the recurrence property value. Possible values for App update on Windows10 recurrence.
func (m *Windows10AppsForceUpdateSchedule) SetRecurrence(value *Windows10AppsUpdateRecurrence)() {
    m.recurrence = value
}
// SetRunImmediatelyIfAfterStartDateTime sets the runImmediatelyIfAfterStartDateTime property value. If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
func (m *Windows10AppsForceUpdateSchedule) SetRunImmediatelyIfAfterStartDateTime(value *bool)() {
    m.runImmediatelyIfAfterStartDateTime = value
}
// SetStartDateTime sets the startDateTime property value. The start time for the force restart.
func (m *Windows10AppsForceUpdateSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
