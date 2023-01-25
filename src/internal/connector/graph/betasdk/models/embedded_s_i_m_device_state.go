package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmbeddedSIMDeviceState describes the embedded SIM activation code deployment state in relation to a device.
type EmbeddedSIMDeviceState struct {
    Entity
    // The time the embedded SIM device status was created. Generated service side.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device name to which the subscription was provisioned e.g. DESKTOP-JOE
    deviceName *string
    // The time the embedded SIM device last checked in. Updated service side.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time the embedded SIM device status was last modified. Updated service side.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Describes the various states for an embedded SIM activation code.
    state *EmbeddedSIMDeviceStateValue
    // String description of the provisioning state.
    stateDetails *string
    // The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be deployed.
    universalIntegratedCircuitCardIdentifier *string
    // Username which the subscription was provisioned to e.g. joe@contoso.com
    userName *string
}
// NewEmbeddedSIMDeviceState instantiates a new embeddedSIMDeviceState and sets the default values.
func NewEmbeddedSIMDeviceState()(*EmbeddedSIMDeviceState) {
    m := &EmbeddedSIMDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEmbeddedSIMDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmbeddedSIMDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMDeviceState(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The time the embedded SIM device status was created. Generated service side.
func (m *EmbeddedSIMDeviceState) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeviceName gets the deviceName property value. Device name to which the subscription was provisioned e.g. DESKTOP-JOE
func (m *EmbeddedSIMDeviceState) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmbeddedSIMDeviceStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*EmbeddedSIMDeviceStateValue))
        }
        return nil
    }
    res["stateDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateDetails(val)
        }
        return nil
    }
    res["universalIntegratedCircuitCardIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniversalIntegratedCircuitCardIdentifier(val)
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
    return res
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The time the embedded SIM device last checked in. Updated service side.
func (m *EmbeddedSIMDeviceState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the embedded SIM device status was last modified. Updated service side.
func (m *EmbeddedSIMDeviceState) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetState gets the state property value. Describes the various states for an embedded SIM activation code.
func (m *EmbeddedSIMDeviceState) GetState()(*EmbeddedSIMDeviceStateValue) {
    return m.state
}
// GetStateDetails gets the stateDetails property value. String description of the provisioning state.
func (m *EmbeddedSIMDeviceState) GetStateDetails()(*string) {
    return m.stateDetails
}
// GetUniversalIntegratedCircuitCardIdentifier gets the universalIntegratedCircuitCardIdentifier property value. The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be deployed.
func (m *EmbeddedSIMDeviceState) GetUniversalIntegratedCircuitCardIdentifier()(*string) {
    return m.universalIntegratedCircuitCardIdentifier
}
// GetUserName gets the userName property value. Username which the subscription was provisioned to e.g. joe@contoso.com
func (m *EmbeddedSIMDeviceState) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *EmbeddedSIMDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
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
        err = writer.WriteStringValue("stateDetails", m.GetStateDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("universalIntegratedCircuitCardIdentifier", m.GetUniversalIntegratedCircuitCardIdentifier())
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
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The time the embedded SIM device status was created. Generated service side.
func (m *EmbeddedSIMDeviceState) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeviceName sets the deviceName property value. Device name to which the subscription was provisioned e.g. DESKTOP-JOE
func (m *EmbeddedSIMDeviceState) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The time the embedded SIM device last checked in. Updated service side.
func (m *EmbeddedSIMDeviceState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the embedded SIM device status was last modified. Updated service side.
func (m *EmbeddedSIMDeviceState) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetState sets the state property value. Describes the various states for an embedded SIM activation code.
func (m *EmbeddedSIMDeviceState) SetState(value *EmbeddedSIMDeviceStateValue)() {
    m.state = value
}
// SetStateDetails sets the stateDetails property value. String description of the provisioning state.
func (m *EmbeddedSIMDeviceState) SetStateDetails(value *string)() {
    m.stateDetails = value
}
// SetUniversalIntegratedCircuitCardIdentifier sets the universalIntegratedCircuitCardIdentifier property value. The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be deployed.
func (m *EmbeddedSIMDeviceState) SetUniversalIntegratedCircuitCardIdentifier(value *string)() {
    m.universalIntegratedCircuitCardIdentifier = value
}
// SetUserName sets the userName property value. Username which the subscription was provisioned to e.g. joe@contoso.com
func (m *EmbeddedSIMDeviceState) SetUserName(value *string)() {
    m.userName = value
}
