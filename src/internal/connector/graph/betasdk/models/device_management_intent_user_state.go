package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentUserState entity that represents user state for an intent
type DeviceManagementIntentUserState struct {
    Entity
    // Count of Devices that belongs to a user for an intent
    deviceCount *int32
    // Last modified date time of an intent report
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The state property
    state *ComplianceStatus
    // The user name that is being reported on a device
    userName *string
    // The user principal name that is being reported on a device
    userPrincipalName *string
}
// NewDeviceManagementIntentUserState instantiates a new deviceManagementIntentUserState and sets the default values.
func NewDeviceManagementIntentUserState()(*DeviceManagementIntentUserState) {
    m := &DeviceManagementIntentUserState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentUserStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentUserStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentUserState(), nil
}
// GetDeviceCount gets the deviceCount property value. Count of Devices that belongs to a user for an intent
func (m *DeviceManagementIntentUserState) GetDeviceCount()(*int32) {
    return m.deviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentUserState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
func (m *DeviceManagementIntentUserState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetState gets the state property value. The state property
func (m *DeviceManagementIntentUserState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetUserName gets the userName property value. The user name that is being reported on a device
func (m *DeviceManagementIntentUserState) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name that is being reported on a device
func (m *DeviceManagementIntentUserState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentUserState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
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
// SetDeviceCount sets the deviceCount property value. Count of Devices that belongs to a user for an intent
func (m *DeviceManagementIntentUserState) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last modified date time of an intent report
func (m *DeviceManagementIntentUserState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetState sets the state property value. The state property
func (m *DeviceManagementIntentUserState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetUserName sets the userName property value. The user name that is being reported on a device
func (m *DeviceManagementIntentUserState) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name that is being reported on a device
func (m *DeviceManagementIntentUserState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
