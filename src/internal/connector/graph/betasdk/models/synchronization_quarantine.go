package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationQuarantine 
type SynchronizationQuarantine struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Date and time when the quarantine was last evaluated and imposed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    currentBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Describes the error(s) that occurred when putting the synchronization job into quarantine.
    error SynchronizationErrorable
    // Date and time when the next attempt to re-evaluate the quarantine will be made. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    nextAttempt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The reason property
    reason *QuarantineReason
    // Date and time when the quarantine was first imposed in this series (a series starts when a quarantine is first imposed, and is reset as soon as the quarantine is lifted). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    seriesBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of times in this series the quarantine was re-evaluated and left in effect (a series starts when quarantine is first imposed, and is reset as soon as quarantine is lifted).
    seriesCount *int64
}
// NewSynchronizationQuarantine instantiates a new synchronizationQuarantine and sets the default values.
func NewSynchronizationQuarantine()(*SynchronizationQuarantine) {
    m := &SynchronizationQuarantine{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationQuarantineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationQuarantineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationQuarantine(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationQuarantine) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCurrentBegan gets the currentBegan property value. Date and time when the quarantine was last evaluated and imposed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) GetCurrentBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.currentBegan
}
// GetError gets the error property value. Describes the error(s) that occurred when putting the synchronization job into quarantine.
func (m *SynchronizationQuarantine) GetError()(SynchronizationErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationQuarantine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentBegan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentBegan(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(SynchronizationErrorable))
        }
        return nil
    }
    res["nextAttempt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextAttempt(val)
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
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQuarantineReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*QuarantineReason))
        }
        return nil
    }
    res["seriesBegan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesBegan(val)
        }
        return nil
    }
    res["seriesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesCount(val)
        }
        return nil
    }
    return res
}
// GetNextAttempt gets the nextAttempt property value. Date and time when the next attempt to re-evaluate the quarantine will be made. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) GetNextAttempt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nextAttempt
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationQuarantine) GetOdataType()(*string) {
    return m.odataType
}
// GetReason gets the reason property value. The reason property
func (m *SynchronizationQuarantine) GetReason()(*QuarantineReason) {
    return m.reason
}
// GetSeriesBegan gets the seriesBegan property value. Date and time when the quarantine was first imposed in this series (a series starts when a quarantine is first imposed, and is reset as soon as the quarantine is lifted). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) GetSeriesBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.seriesBegan
}
// GetSeriesCount gets the seriesCount property value. Number of times in this series the quarantine was re-evaluated and left in effect (a series starts when quarantine is first imposed, and is reset as soon as quarantine is lifted).
func (m *SynchronizationQuarantine) GetSeriesCount()(*int64) {
    return m.seriesCount
}
// Serialize serializes information the current object
func (m *SynchronizationQuarantine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("currentBegan", m.GetCurrentBegan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("nextAttempt", m.GetNextAttempt())
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
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err := writer.WriteStringValue("reason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("seriesBegan", m.GetSeriesBegan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("seriesCount", m.GetSeriesCount())
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
func (m *SynchronizationQuarantine) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCurrentBegan sets the currentBegan property value. Date and time when the quarantine was last evaluated and imposed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) SetCurrentBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.currentBegan = value
}
// SetError sets the error property value. Describes the error(s) that occurred when putting the synchronization job into quarantine.
func (m *SynchronizationQuarantine) SetError(value SynchronizationErrorable)() {
    m.error = value
}
// SetNextAttempt sets the nextAttempt property value. Date and time when the next attempt to re-evaluate the quarantine will be made. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) SetNextAttempt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextAttempt = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationQuarantine) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReason sets the reason property value. The reason property
func (m *SynchronizationQuarantine) SetReason(value *QuarantineReason)() {
    m.reason = value
}
// SetSeriesBegan sets the seriesBegan property value. Date and time when the quarantine was first imposed in this series (a series starts when a quarantine is first imposed, and is reset as soon as the quarantine is lifted). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationQuarantine) SetSeriesBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.seriesBegan = value
}
// SetSeriesCount sets the seriesCount property value. Number of times in this series the quarantine was re-evaluated and left in effect (a series starts when quarantine is first imposed, and is reset as soon as quarantine is lifted).
func (m *SynchronizationQuarantine) SetSeriesCount(value *int64)() {
    m.seriesCount = value
}
