package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptDeviceState contains properties for device run state of the device compliance script.
type DeviceComplianceScriptDeviceState struct {
    Entity
    // Indicates the type of execution status of the device management script.
    detectionState *RunState
    // The next timestamp of when the device compliance script is expected to execute
    expectedStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last timestamp of when the device compliance script executed
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time that Intune Managment Extension synced with Intune
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managed device on which the device compliance script executed
    managedDevice ManagedDeviceable
    // Error from the detection script
    scriptError *string
    // Output of the detection script
    scriptOutput *string
}
// NewDeviceComplianceScriptDeviceState instantiates a new deviceComplianceScriptDeviceState and sets the default values.
func NewDeviceComplianceScriptDeviceState()(*DeviceComplianceScriptDeviceState) {
    m := &DeviceComplianceScriptDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceComplianceScriptDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceScriptDeviceState(), nil
}
// GetDetectionState gets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceComplianceScriptDeviceState) GetDetectionState()(*RunState) {
    return m.detectionState
}
// GetExpectedStateUpdateDateTime gets the expectedStateUpdateDateTime property value. The next timestamp of when the device compliance script is expected to execute
func (m *DeviceComplianceScriptDeviceState) GetExpectedStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expectedStateUpdateDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["detectionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionState(val.(*RunState))
        }
        return nil
    }
    res["expectedStateUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedStateUpdateDateTime(val)
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
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
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
    res["scriptError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptError(val)
        }
        return nil
    }
    res["scriptOutput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptOutput(val)
        }
        return nil
    }
    return res
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the device compliance script executed
func (m *DeviceComplianceScriptDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastStateUpdateDateTime
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceComplianceScriptDeviceState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetManagedDevice gets the managedDevice property value. The managed device on which the device compliance script executed
func (m *DeviceComplianceScriptDeviceState) GetManagedDevice()(ManagedDeviceable) {
    return m.managedDevice
}
// GetScriptError gets the scriptError property value. Error from the detection script
func (m *DeviceComplianceScriptDeviceState) GetScriptError()(*string) {
    return m.scriptError
}
// GetScriptOutput gets the scriptOutput property value. Output of the detection script
func (m *DeviceComplianceScriptDeviceState) GetScriptOutput()(*string) {
    return m.scriptOutput
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetectionState() != nil {
        cast := (*m.GetDetectionState()).String()
        err = writer.WriteStringValue("detectionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expectedStateUpdateDateTime", m.GetExpectedStateUpdateDateTime())
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
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
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
        err = writer.WriteStringValue("scriptError", m.GetScriptError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scriptOutput", m.GetScriptOutput())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetectionState sets the detectionState property value. Indicates the type of execution status of the device management script.
func (m *DeviceComplianceScriptDeviceState) SetDetectionState(value *RunState)() {
    m.detectionState = value
}
// SetExpectedStateUpdateDateTime sets the expectedStateUpdateDateTime property value. The next timestamp of when the device compliance script is expected to execute
func (m *DeviceComplianceScriptDeviceState) SetExpectedStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expectedStateUpdateDateTime = value
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the device compliance script executed
func (m *DeviceComplianceScriptDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStateUpdateDateTime = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time that Intune Managment Extension synced with Intune
func (m *DeviceComplianceScriptDeviceState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetManagedDevice sets the managedDevice property value. The managed device on which the device compliance script executed
func (m *DeviceComplianceScriptDeviceState) SetManagedDevice(value ManagedDeviceable)() {
    m.managedDevice = value
}
// SetScriptError sets the scriptError property value. Error from the detection script
func (m *DeviceComplianceScriptDeviceState) SetScriptError(value *string)() {
    m.scriptError = value
}
// SetScriptOutput sets the scriptOutput property value. Output of the detection script
func (m *DeviceComplianceScriptDeviceState) SetScriptOutput(value *string)() {
    m.scriptOutput = value
}
