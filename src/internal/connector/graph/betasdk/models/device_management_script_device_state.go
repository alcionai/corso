package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptDeviceState contains properties for device run state of the device management script.
type DeviceManagementScriptDeviceState struct {
    Entity
    // Error code corresponding to erroneous execution of the device management script.
    errorCode *int32
    // Error description corresponding to erroneous execution of the device management script.
    errorDescription *string
    // Latest time the device management script executes.
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managed devices that executes the device management script.
    managedDevice ManagedDeviceable
    // Details of execution output.
    resultMessage *string
    // Indicates the type of execution status of the device management script.
    runState *RunState
}
// NewDeviceManagementScriptDeviceState instantiates a new deviceManagementScriptDeviceState and sets the default values.
func NewDeviceManagementScriptDeviceState()(*DeviceManagementScriptDeviceState) {
    m := &DeviceManagementScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptDeviceState(), nil
}
// GetErrorCode gets the errorCode property value. Error code corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetErrorDescription gets the errorDescription property value. Error description corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) GetErrorDescription()(*string) {
    return m.errorDescription
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDescription(val)
        }
        return nil
    }
    res["lastStateUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastStateUpdateDateTime(val)
        }
        return nil
    }
    res["managedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDevice(val.(ManagedDeviceable))
        }
        return nil
    }
    res["resultMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultMessage(val)
        }
        return nil
    }
    res["runState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunState(val.(*RunState))
        }
        return nil
    }
    return res
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. Latest time the device management script executes.
func (m *DeviceManagementScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastStateUpdateDateTime
}
// GetManagedDevice gets the managedDevice property value. The managed devices that executes the device management script.
func (m *DeviceManagementScriptDeviceState) GetManagedDevice()(ManagedDeviceable) {
    return m.managedDevice
}
// GetResultMessage gets the resultMessage property value. Details of execution output.
func (m *DeviceManagementScriptDeviceState) GetResultMessage()(*string) {
    return m.resultMessage
}
// GetRunState gets the runState property value. Indicates the type of execution status of the device management script.
func (m *DeviceManagementScriptDeviceState) GetRunState()(*RunState) {
    return m.runState
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorDescription", m.GetErrorDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastStateUpdateDateTime", m.GetLastStateUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDevice", m.GetManagedDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resultMessage", m.GetResultMessage())
        if err != nil {
            return err
        }
    }
    if m.GetRunState() != nil {
        cast := (*m.GetRunState()).String()
        err = writer.WriteStringValue("runState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorCode sets the errorCode property value. Error code corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetErrorDescription sets the errorDescription property value. Error description corresponding to erroneous execution of the device management script.
func (m *DeviceManagementScriptDeviceState) SetErrorDescription(value *string)() {
    m.errorDescription = value
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. Latest time the device management script executes.
func (m *DeviceManagementScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStateUpdateDateTime = value
}
// SetManagedDevice sets the managedDevice property value. The managed devices that executes the device management script.
func (m *DeviceManagementScriptDeviceState) SetManagedDevice(value ManagedDeviceable)() {
    m.managedDevice = value
}
// SetResultMessage sets the resultMessage property value. Details of execution output.
func (m *DeviceManagementScriptDeviceState) SetResultMessage(value *string)() {
    m.resultMessage = value
}
// SetRunState sets the runState property value. Indicates the type of execution status of the device management script.
func (m *DeviceManagementScriptDeviceState) SetRunState(value *RunState)() {
    m.runState = value
}
