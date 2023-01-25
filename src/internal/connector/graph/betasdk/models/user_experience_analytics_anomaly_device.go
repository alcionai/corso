package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalyDevice the user experience analytics anomaly entity contains device details.
type UserExperienceAnalyticsAnomalyDevice struct {
    Entity
    // The unique identifier of the anomaly.
    anomalyId *string
    // Indicates the first occurance date and time for the anomaly on the device.
    anomalyOnDeviceFirstOccurrenceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates the latest occurance date and time for the anomaly on the device.
    anomalyOnDeviceLatestOccurrenceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the device.
    deviceId *string
    // The manufacturer name of the device.
    deviceManufacturer *string
    // The model name of the device.
    deviceModel *string
    // The name of the device.
    deviceName *string
    // The name of the OS installed on the device.
    osName *string
    // The OS version installed on the device.
    osVersion *string
}
// NewUserExperienceAnalyticsAnomalyDevice instantiates a new userExperienceAnalyticsAnomalyDevice and sets the default values.
func NewUserExperienceAnalyticsAnomalyDevice()(*UserExperienceAnalyticsAnomalyDevice) {
    m := &UserExperienceAnalyticsAnomalyDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomalyDevice(), nil
}
// GetAnomalyId gets the anomalyId property value. The unique identifier of the anomaly.
func (m *UserExperienceAnalyticsAnomalyDevice) GetAnomalyId()(*string) {
    return m.anomalyId
}
// GetAnomalyOnDeviceFirstOccurrenceDateTime gets the anomalyOnDeviceFirstOccurrenceDateTime property value. Indicates the first occurance date and time for the anomaly on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetAnomalyOnDeviceFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.anomalyOnDeviceFirstOccurrenceDateTime
}
// GetAnomalyOnDeviceLatestOccurrenceDateTime gets the anomalyOnDeviceLatestOccurrenceDateTime property value. Indicates the latest occurance date and time for the anomaly on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetAnomalyOnDeviceLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.anomalyOnDeviceLatestOccurrenceDateTime
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceManufacturer gets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceManufacturer()(*string) {
    return m.deviceManufacturer
}
// GetDeviceModel gets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetDeviceName gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAnomalyDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["anomalyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyId(val)
        }
        return nil
    }
    res["anomalyOnDeviceFirstOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyOnDeviceFirstOccurrenceDateTime(val)
        }
        return nil
    }
    res["anomalyOnDeviceLatestOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyOnDeviceLatestOccurrenceDateTime(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManufacturer(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["osName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsName(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    return res
}
// GetOsName gets the osName property value. The name of the OS installed on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetOsName()(*string) {
    return m.osName
}
// GetOsVersion gets the osVersion property value. The OS version installed on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomalyDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("anomalyId", m.GetAnomalyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("anomalyOnDeviceFirstOccurrenceDateTime", m.GetAnomalyOnDeviceFirstOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("anomalyOnDeviceLatestOccurrenceDateTime", m.GetAnomalyOnDeviceLatestOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceManufacturer", m.GetDeviceManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osName", m.GetOsName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnomalyId sets the anomalyId property value. The unique identifier of the anomaly.
func (m *UserExperienceAnalyticsAnomalyDevice) SetAnomalyId(value *string)() {
    m.anomalyId = value
}
// SetAnomalyOnDeviceFirstOccurrenceDateTime sets the anomalyOnDeviceFirstOccurrenceDateTime property value. Indicates the first occurance date and time for the anomaly on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetAnomalyOnDeviceFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.anomalyOnDeviceFirstOccurrenceDateTime = value
}
// SetAnomalyOnDeviceLatestOccurrenceDateTime sets the anomalyOnDeviceLatestOccurrenceDateTime property value. Indicates the latest occurance date and time for the anomaly on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetAnomalyOnDeviceLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.anomalyOnDeviceLatestOccurrenceDateTime = value
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceManufacturer sets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceManufacturer(value *string)() {
    m.deviceManufacturer = value
}
// SetDeviceModel sets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetOsName sets the osName property value. The name of the OS installed on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetOsName(value *string)() {
    m.osName = value
}
// SetOsVersion sets the osVersion property value. The OS version installed on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
