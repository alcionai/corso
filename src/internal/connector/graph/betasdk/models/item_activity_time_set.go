package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActivityTimeSet 
type ItemActivityTimeSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The lastRecordedDateTime property
    lastRecordedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // When the activity was observed to take place.
    observedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // When the observation was recorded on the service.
    recordedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewItemActivityTimeSet instantiates a new itemActivityTimeSet and sets the default values.
func NewItemActivityTimeSet()(*ItemActivityTimeSet) {
    m := &ItemActivityTimeSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemActivityTimeSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActivityTimeSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActivityTimeSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActivityTimeSet) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivityTimeSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lastRecordedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRecordedDateTime(val)
        }
        return nil
    }
    res["observedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObservedDateTime(val)
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
    res["recordedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordedDateTime(val)
        }
        return nil
    }
    return res
}
// GetLastRecordedDateTime gets the lastRecordedDateTime property value. The lastRecordedDateTime property
func (m *ItemActivityTimeSet) GetLastRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRecordedDateTime
}
// GetObservedDateTime gets the observedDateTime property value. When the activity was observed to take place.
func (m *ItemActivityTimeSet) GetObservedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.observedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ItemActivityTimeSet) GetOdataType()(*string) {
    return m.odataType
}
// GetRecordedDateTime gets the recordedDateTime property value. When the observation was recorded on the service.
func (m *ItemActivityTimeSet) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.recordedDateTime
}
// Serialize serializes information the current object
func (m *ItemActivityTimeSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastRecordedDateTime", m.GetLastRecordedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("observedDateTime", m.GetObservedDateTime())
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
        err := writer.WriteTimeValue("recordedDateTime", m.GetRecordedDateTime())
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
func (m *ItemActivityTimeSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLastRecordedDateTime sets the lastRecordedDateTime property value. The lastRecordedDateTime property
func (m *ItemActivityTimeSet) SetLastRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRecordedDateTime = value
}
// SetObservedDateTime sets the observedDateTime property value. When the activity was observed to take place.
func (m *ItemActivityTimeSet) SetObservedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.observedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ItemActivityTimeSet) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecordedDateTime sets the recordedDateTime property value. When the observation was recorded on the service.
func (m *ItemActivityTimeSet) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recordedDateTime = value
}
