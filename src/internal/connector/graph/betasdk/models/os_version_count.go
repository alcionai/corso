package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OsVersionCount count of devices with malware for each OS version
type OsVersionCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Count of devices with malware for the OS version
    deviceCount *int32
    // The Timestamp of the last update for the device count in UTC
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // OS version
    osVersion *string
}
// NewOsVersionCount instantiates a new osVersionCount and sets the default values.
func NewOsVersionCount()(*OsVersionCount) {
    m := &OsVersionCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOsVersionCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOsVersionCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOsVersionCount(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OsVersionCount) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceCount gets the deviceCount property value. Count of devices with malware for the OS version
func (m *OsVersionCount) GetDeviceCount()(*int32) {
    return m.deviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OsVersionCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
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
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. The Timestamp of the last update for the device count in UTC
func (m *OsVersionCount) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdateDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OsVersionCount) GetOdataType()(*string) {
    return m.odataType
}
// GetOsVersion gets the osVersion property value. OS version
func (m *OsVersionCount) GetOsVersion()(*string) {
    return m.osVersion
}
// Serialize serializes information the current object
func (m *OsVersionCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
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
        err := writer.WriteStringValue("osVersion", m.GetOsVersion())
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
func (m *OsVersionCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceCount sets the deviceCount property value. Count of devices with malware for the OS version
func (m *OsVersionCount) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. The Timestamp of the last update for the device count in UTC
func (m *OsVersionCount) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OsVersionCount) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOsVersion sets the osVersion property value. OS version
func (m *OsVersionCount) SetOsVersion(value *string)() {
    m.osVersion = value
}
