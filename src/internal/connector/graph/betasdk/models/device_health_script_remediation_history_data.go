package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptRemediationHistoryData the number of devices remediated by a device health script on a given date.
type DeviceHealthScriptRemediationHistoryData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The date on which devices were remediated by the device health script.
    date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of devices that were found to have no issue by the device health script.
    noIssueDeviceCount *int32
    // The OdataType property
    odataType *string
    // The number of devices remediated by the device health script.
    remediatedDeviceCount *int32
}
// NewDeviceHealthScriptRemediationHistoryData instantiates a new deviceHealthScriptRemediationHistoryData and sets the default values.
func NewDeviceHealthScriptRemediationHistoryData()(*DeviceHealthScriptRemediationHistoryData) {
    m := &DeviceHealthScriptRemediationHistoryData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceHealthScriptRemediationHistoryDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRemediationHistoryDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptRemediationHistoryData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationHistoryData) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDate gets the date property value. The date on which devices were remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.date
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRemediationHistoryData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["noIssueDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoIssueDeviceCount(val)
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
    res["remediatedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetNoIssueDeviceCount gets the noIssueDeviceCount property value. The number of devices that were found to have no issue by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetNoIssueDeviceCount()(*int32) {
    return m.noIssueDeviceCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRemediationHistoryData) GetOdataType()(*string) {
    return m.odataType
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. The number of devices remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) GetRemediatedDeviceCount()(*int32) {
    return m.remediatedDeviceCount
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRemediationHistoryData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("noIssueDeviceCount", m.GetNoIssueDeviceCount())
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
        err := writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
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
func (m *DeviceHealthScriptRemediationHistoryData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDate sets the date property value. The date on which devices were remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) SetDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.date = value
}
// SetNoIssueDeviceCount sets the noIssueDeviceCount property value. The number of devices that were found to have no issue by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) SetNoIssueDeviceCount(value *int32)() {
    m.noIssueDeviceCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRemediationHistoryData) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. The number of devices remediated by the device health script.
func (m *DeviceHealthScriptRemediationHistoryData) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
