package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewRecurrenceSettings 
type AccessReviewRecurrenceSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The duration in days for recurrence.
    durationInDays *int32
    // The OdataType property
    odataType *string
    // The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
    recurrenceCount *int32
    // How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it is never, then there is no explicit end of the recurrence series. If it is endBy, then the recurrence ends at a certain date. If it is occurrences, then the series ends after recurrenceCount instances of the review have completed.
    recurrenceEndType *string
    // The recurrence interval. Possible vaules: onetime, weekly, monthly, quarterly, halfyearly or annual.
    recurrenceType *string
}
// NewAccessReviewRecurrenceSettings instantiates a new accessReviewRecurrenceSettings and sets the default values.
func NewAccessReviewRecurrenceSettings()(*AccessReviewRecurrenceSettings) {
    m := &AccessReviewRecurrenceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessReviewRecurrenceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewRecurrenceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewRecurrenceSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewRecurrenceSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDurationInDays gets the durationInDays property value. The duration in days for recurrence.
func (m *AccessReviewRecurrenceSettings) GetDurationInDays()(*int32) {
    return m.durationInDays
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewRecurrenceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["durationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
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
    res["recurrenceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceCount(val)
        }
        return nil
    }
    res["recurrenceEndType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceEndType(val)
        }
        return nil
    }
    res["recurrenceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewRecurrenceSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRecurrenceCount gets the recurrenceCount property value. The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceCount()(*int32) {
    return m.recurrenceCount
}
// GetRecurrenceEndType gets the recurrenceEndType property value. How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it is never, then there is no explicit end of the recurrence series. If it is endBy, then the recurrence ends at a certain date. If it is occurrences, then the series ends after recurrenceCount instances of the review have completed.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceEndType()(*string) {
    return m.recurrenceEndType
}
// GetRecurrenceType gets the recurrenceType property value. The recurrence interval. Possible vaules: onetime, weekly, monthly, quarterly, halfyearly or annual.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceType()(*string) {
    return m.recurrenceType
}
// Serialize serializes information the current object
func (m *AccessReviewRecurrenceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
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
        err := writer.WriteInt32Value("recurrenceCount", m.GetRecurrenceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceEndType", m.GetRecurrenceEndType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceType", m.GetRecurrenceType())
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
func (m *AccessReviewRecurrenceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDurationInDays sets the durationInDays property value. The duration in days for recurrence.
func (m *AccessReviewRecurrenceSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewRecurrenceSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecurrenceCount sets the recurrenceCount property value. The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceCount(value *int32)() {
    m.recurrenceCount = value
}
// SetRecurrenceEndType sets the recurrenceEndType property value. How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it is never, then there is no explicit end of the recurrence series. If it is endBy, then the recurrence ends at a certain date. If it is occurrences, then the series ends after recurrenceCount instances of the review have completed.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceEndType(value *string)() {
    m.recurrenceEndType = value
}
// SetRecurrenceType sets the recurrenceType property value. The recurrence interval. Possible vaules: onetime, weekly, monthly, quarterly, halfyearly or annual.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceType(value *string)() {
    m.recurrenceType = value
}
