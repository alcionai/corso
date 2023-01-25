package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnmanagedDevice unmanaged device discovered in the network.
type UnmanagedDevice struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Device name.
    deviceName *string
    // Domain.
    domain *string
    // IP address.
    ipAddress *string
    // Last logged on user.
    lastLoggedOnUser *string
    // Last seen date and time.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Location.
    location *string
    // MAC address.
    macAddress *string
    // Manufacturer.
    manufacturer *string
    // Model.
    model *string
    // The OdataType property
    odataType *string
    // Operating system.
    os *string
    // Operating system version.
    osVersion *string
}
// NewUnmanagedDevice instantiates a new unmanagedDevice and sets the default values.
func NewUnmanagedDevice()(*UnmanagedDevice) {
    m := &UnmanagedDevice{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnmanagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnmanagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnmanagedDevice(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnmanagedDevice) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceName gets the deviceName property value. Device name.
func (m *UnmanagedDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDomain gets the domain property value. Domain.
func (m *UnmanagedDevice) GetDomain()(*string) {
    return m.domain
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnmanagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["lastLoggedOnUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastLoggedOnUser(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["macAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
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
    res["os"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
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
// GetIpAddress gets the ipAddress property value. IP address.
func (m *UnmanagedDevice) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetLastLoggedOnUser gets the lastLoggedOnUser property value. Last logged on user.
func (m *UnmanagedDevice) GetLastLoggedOnUser()(*string) {
    return m.lastLoggedOnUser
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. Last seen date and time.
func (m *UnmanagedDevice) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetLocation gets the location property value. Location.
func (m *UnmanagedDevice) GetLocation()(*string) {
    return m.location
}
// GetMacAddress gets the macAddress property value. MAC address.
func (m *UnmanagedDevice) GetMacAddress()(*string) {
    return m.macAddress
}
// GetManufacturer gets the manufacturer property value. Manufacturer.
func (m *UnmanagedDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. Model.
func (m *UnmanagedDevice) GetModel()(*string) {
    return m.model
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UnmanagedDevice) GetOdataType()(*string) {
    return m.odataType
}
// GetOs gets the os property value. Operating system.
func (m *UnmanagedDevice) GetOs()(*string) {
    return m.os
}
// GetOsVersion gets the osVersion property value. Operating system version.
func (m *UnmanagedDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// Serialize serializes information the current object
func (m *UnmanagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastLoggedOnUser", m.GetLastLoggedOnUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
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
        err := writer.WriteStringValue("os", m.GetOs())
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
func (m *UnmanagedDevice) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceName sets the deviceName property value. Device name.
func (m *UnmanagedDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDomain sets the domain property value. Domain.
func (m *UnmanagedDevice) SetDomain(value *string)() {
    m.domain = value
}
// SetIpAddress sets the ipAddress property value. IP address.
func (m *UnmanagedDevice) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetLastLoggedOnUser sets the lastLoggedOnUser property value. Last logged on user.
func (m *UnmanagedDevice) SetLastLoggedOnUser(value *string)() {
    m.lastLoggedOnUser = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. Last seen date and time.
func (m *UnmanagedDevice) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetLocation sets the location property value. Location.
func (m *UnmanagedDevice) SetLocation(value *string)() {
    m.location = value
}
// SetMacAddress sets the macAddress property value. MAC address.
func (m *UnmanagedDevice) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer.
func (m *UnmanagedDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. Model.
func (m *UnmanagedDevice) SetModel(value *string)() {
    m.model = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UnmanagedDevice) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOs sets the os property value. Operating system.
func (m *UnmanagedDevice) SetOs(value *string)() {
    m.os = value
}
// SetOsVersion sets the osVersion property value. Operating system version.
func (m *UnmanagedDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
