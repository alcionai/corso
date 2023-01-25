package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationProgress 
type SynchronizationProgress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The numerator of a progress ratio; the number of units of changes already processed.
    completedUnits *int64
    // The OdataType property
    odataType *string
    // The time of a progress observation as an offset in minutes from UTC.
    progressObservationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The denominator of a progress ratio; a number of units of changes to be processed to accomplish synchronization.
    totalUnits *int64
    // An optional description of the units.
    units *string
}
// NewSynchronizationProgress instantiates a new synchronizationProgress and sets the default values.
func NewSynchronizationProgress()(*SynchronizationProgress) {
    m := &SynchronizationProgress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationProgressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationProgressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationProgress(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationProgress) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompletedUnits gets the completedUnits property value. The numerator of a progress ratio; the number of units of changes already processed.
func (m *SynchronizationProgress) GetCompletedUnits()(*int64) {
    return m.completedUnits
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationProgress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedUnits(val)
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
    res["progressObservationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgressObservationDateTime(val)
        }
        return nil
    }
    res["totalUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUnits(val)
        }
        return nil
    }
    res["units"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnits(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationProgress) GetOdataType()(*string) {
    return m.odataType
}
// GetProgressObservationDateTime gets the progressObservationDateTime property value. The time of a progress observation as an offset in minutes from UTC.
func (m *SynchronizationProgress) GetProgressObservationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.progressObservationDateTime
}
// GetTotalUnits gets the totalUnits property value. The denominator of a progress ratio; a number of units of changes to be processed to accomplish synchronization.
func (m *SynchronizationProgress) GetTotalUnits()(*int64) {
    return m.totalUnits
}
// GetUnits gets the units property value. An optional description of the units.
func (m *SynchronizationProgress) GetUnits()(*string) {
    return m.units
}
// Serialize serializes information the current object
func (m *SynchronizationProgress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("completedUnits", m.GetCompletedUnits())
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
        err := writer.WriteTimeValue("progressObservationDateTime", m.GetProgressObservationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalUnits", m.GetTotalUnits())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("units", m.GetUnits())
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
func (m *SynchronizationProgress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompletedUnits sets the completedUnits property value. The numerator of a progress ratio; the number of units of changes already processed.
func (m *SynchronizationProgress) SetCompletedUnits(value *int64)() {
    m.completedUnits = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationProgress) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProgressObservationDateTime sets the progressObservationDateTime property value. The time of a progress observation as an offset in minutes from UTC.
func (m *SynchronizationProgress) SetProgressObservationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.progressObservationDateTime = value
}
// SetTotalUnits sets the totalUnits property value. The denominator of a progress ratio; a number of units of changes to be processed to accomplish synchronization.
func (m *SynchronizationProgress) SetTotalUnits(value *int64)() {
    m.totalUnits = value
}
// SetUnits sets the units property value. An optional description of the units.
func (m *SynchronizationProgress) SetUnits(value *string)() {
    m.units = value
}
