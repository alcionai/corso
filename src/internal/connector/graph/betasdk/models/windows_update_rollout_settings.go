package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateRolloutSettings a complex type to store the windows update rollout settings including offer start date time, offer end date time, and days between each set of offers.
type WindowsUpdateRolloutSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The feature update's ending  of release date and time to be set, update, and displayed for a feature Update profile for example: 2020-06-09T10:00:00Z.
    offerEndDateTimeInUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of day(s) between each set of offers to be set, updated, and displayed for a feature update profile, for example: if OfferStartDateTimeInUTC is 2020-06-09T10:00:00Z, and OfferIntervalInDays is 1, then the next two sets of offers will be made consecutively on 2020-06-10T10:00:00Z (next day at the same specified time) and 2020-06-11T10:00:00Z (next next day at the same specified time) with 1 day in between each set of offers.
    offerIntervalInDays *int32
    // The feature update's starting date and time to be set, update, and displayed for a feature Update profile for example: 2020-06-09T10:00:00Z.
    offerStartDateTimeInUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewWindowsUpdateRolloutSettings instantiates a new windowsUpdateRolloutSettings and sets the default values.
func NewWindowsUpdateRolloutSettings()(*WindowsUpdateRolloutSettings) {
    m := &WindowsUpdateRolloutSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsUpdateRolloutSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateRolloutSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdateRolloutSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdateRolloutSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateRolloutSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["offerEndDateTimeInUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferEndDateTimeInUTC(val)
        }
        return nil
    }
    res["offerIntervalInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferIntervalInDays(val)
        }
        return nil
    }
    res["offerStartDateTimeInUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferStartDateTimeInUTC(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsUpdateRolloutSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetOfferEndDateTimeInUTC gets the offerEndDateTimeInUTC property value. The feature update's ending  of release date and time to be set, update, and displayed for a feature Update profile for example: 2020-06-09T10:00:00Z.
func (m *WindowsUpdateRolloutSettings) GetOfferEndDateTimeInUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.offerEndDateTimeInUTC
}
// GetOfferIntervalInDays gets the offerIntervalInDays property value. The number of day(s) between each set of offers to be set, updated, and displayed for a feature update profile, for example: if OfferStartDateTimeInUTC is 2020-06-09T10:00:00Z, and OfferIntervalInDays is 1, then the next two sets of offers will be made consecutively on 2020-06-10T10:00:00Z (next day at the same specified time) and 2020-06-11T10:00:00Z (next next day at the same specified time) with 1 day in between each set of offers.
func (m *WindowsUpdateRolloutSettings) GetOfferIntervalInDays()(*int32) {
    return m.offerIntervalInDays
}
// GetOfferStartDateTimeInUTC gets the offerStartDateTimeInUTC property value. The feature update's starting date and time to be set, update, and displayed for a feature Update profile for example: 2020-06-09T10:00:00Z.
func (m *WindowsUpdateRolloutSettings) GetOfferStartDateTimeInUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.offerStartDateTimeInUTC
}
// Serialize serializes information the current object
func (m *WindowsUpdateRolloutSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("offerEndDateTimeInUTC", m.GetOfferEndDateTimeInUTC())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("offerIntervalInDays", m.GetOfferIntervalInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("offerStartDateTimeInUTC", m.GetOfferStartDateTimeInUTC())
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
func (m *WindowsUpdateRolloutSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsUpdateRolloutSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOfferEndDateTimeInUTC sets the offerEndDateTimeInUTC property value. The feature update's ending  of release date and time to be set, update, and displayed for a feature Update profile for example: 2020-06-09T10:00:00Z.
func (m *WindowsUpdateRolloutSettings) SetOfferEndDateTimeInUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.offerEndDateTimeInUTC = value
}
// SetOfferIntervalInDays sets the offerIntervalInDays property value. The number of day(s) between each set of offers to be set, updated, and displayed for a feature update profile, for example: if OfferStartDateTimeInUTC is 2020-06-09T10:00:00Z, and OfferIntervalInDays is 1, then the next two sets of offers will be made consecutively on 2020-06-10T10:00:00Z (next day at the same specified time) and 2020-06-11T10:00:00Z (next next day at the same specified time) with 1 day in between each set of offers.
func (m *WindowsUpdateRolloutSettings) SetOfferIntervalInDays(value *int32)() {
    m.offerIntervalInDays = value
}
// SetOfferStartDateTimeInUTC sets the offerStartDateTimeInUTC property value. The feature update's starting date and time to be set, update, and displayed for a feature Update profile for example: 2020-06-09T10:00:00Z.
func (m *WindowsUpdateRolloutSettings) SetOfferStartDateTimeInUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.offerStartDateTimeInUTC = value
}
