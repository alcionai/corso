package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentDeviceState entity that represents device state for an intent
type DeviceManagementIntentDeviceState struct {
    Entity
    // Device name that is being reported
    deviceDisplayName *string
    // Device id that is being reported
    deviceId *string
    // Last modified date time of an intent report
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The state property
    state *ComplianceStatus
    // The user name that is being reported on a device
    userName *string
    // The user principal name that is being reported on a device
    userPrincipalName *string
}
// NewDeviceManagementIntentDeviceState instantiates a new deviceManagementIntentDeviceState and sets the default values.
func NewDeviceManagementIntentDeviceState()(*DeviceManagementIntentDeviceState) {
    m := &DeviceManagementIntentDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentDeviceState(), nil
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Device name that is being reported
func (m *DeviceManagementIntentDeviceState) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetDeviceId gets the deviceId property value. Device id that is being reported
func (m *DeviceManagementIntentDeviceState) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
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
    res["lastReportedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Last modified date time of an intent report
func (m *DeviceManagementIntentDeviceState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetState gets the state property value. The state property
func (m *DeviceManagementIntentDeviceState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetUserName gets the userName property value. The user name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
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
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Device name that is being reported
func (m *DeviceManagementIntentDeviceState) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetDeviceId sets the deviceId property value. Device id that is being reported
func (m *DeviceManagementIntentDeviceState) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last modified date time of an intent report
func (m *DeviceManagementIntentDeviceState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetState sets the state property value. The state property
func (m *DeviceManagementIntentDeviceState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetUserName sets the userName property value. The user name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
