package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionWipeAction represents wipe requests issued by tenant admin for Bring-Your-Own-Device(BYOD) Windows devices.
type WindowsInformationProtectionWipeAction struct {
    Entity
    // Last checkin time of the device that was targeted by this wipe action.
    lastCheckInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *ActionState
    // Targeted device Mac address.
    targetedDeviceMacAddress *string
    // Targeted device name.
    targetedDeviceName *string
    // The DeviceRegistrationId being targeted by this wipe action.
    targetedDeviceRegistrationId *string
    // The UserId being targeted by this wipe action.
    targetedUserId *string
}
// NewWindowsInformationProtectionWipeAction instantiates a new windowsInformationProtectionWipeAction and sets the default values.
func NewWindowsInformationProtectionWipeAction()(*WindowsInformationProtectionWipeAction) {
    m := &WindowsInformationProtectionWipeAction{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionWipeAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionWipeAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastCheckInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckInDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ActionState))
        }
        return nil
    }
    res["targetedDeviceMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceMacAddress(val)
        }
        return nil
    }
    res["targetedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceName(val)
        }
        return nil
    }
    res["targetedDeviceRegistrationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceRegistrationId(val)
        }
        return nil
    }
    res["targetedUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedUserId(val)
        }
        return nil
    }
    return res
}
// GetLastCheckInDateTime gets the lastCheckInDateTime property value. Last checkin time of the device that was targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCheckInDateTime
}
// GetStatus gets the status property value. The status property
func (m *WindowsInformationProtectionWipeAction) GetStatus()(*ActionState) {
    return m.status
}
// GetTargetedDeviceMacAddress gets the targetedDeviceMacAddress property value. Targeted device Mac address.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceMacAddress()(*string) {
    return m.targetedDeviceMacAddress
}
// GetTargetedDeviceName gets the targetedDeviceName property value. Targeted device name.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceName()(*string) {
    return m.targetedDeviceName
}
// GetTargetedDeviceRegistrationId gets the targetedDeviceRegistrationId property value. The DeviceRegistrationId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceRegistrationId()(*string) {
    return m.targetedDeviceRegistrationId
}
// GetTargetedUserId gets the targetedUserId property value. The UserId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetTargetedUserId()(*string) {
    return m.targetedUserId
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionWipeAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastCheckInDateTime", m.GetLastCheckInDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceMacAddress", m.GetTargetedDeviceMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceName", m.GetTargetedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceRegistrationId", m.GetTargetedDeviceRegistrationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedUserId", m.GetTargetedUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastCheckInDateTime sets the lastCheckInDateTime property value. Last checkin time of the device that was targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckInDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *WindowsInformationProtectionWipeAction) SetStatus(value *ActionState)() {
    m.status = value
}
// SetTargetedDeviceMacAddress sets the targetedDeviceMacAddress property value. Targeted device Mac address.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceMacAddress(value *string)() {
    m.targetedDeviceMacAddress = value
}
// SetTargetedDeviceName sets the targetedDeviceName property value. Targeted device name.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceName(value *string)() {
    m.targetedDeviceName = value
}
// SetTargetedDeviceRegistrationId sets the targetedDeviceRegistrationId property value. The DeviceRegistrationId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceRegistrationId(value *string)() {
    m.targetedDeviceRegistrationId = value
}
// SetTargetedUserId sets the targetedUserId property value. The UserId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) SetTargetedUserId(value *string)() {
    m.targetedUserId = value
}
