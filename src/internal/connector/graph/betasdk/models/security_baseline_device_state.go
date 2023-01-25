package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineDeviceState the security baseline compliance state summary of the security baseline for a device.
type SecurityBaselineDeviceState struct {
    Entity
    // Display name of the device
    deviceDisplayName *string
    // Last modified date time of the policy report
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Intune device id
    managedDeviceId *string
    // Security Baseline Compliance State
    state *SecurityBaselineComplianceState
    // User Principal Name
    userPrincipalName *string
}
// NewSecurityBaselineDeviceState instantiates a new securityBaselineDeviceState and sets the default values.
func NewSecurityBaselineDeviceState()(*SecurityBaselineDeviceState) {
    m := &SecurityBaselineDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityBaselineDeviceStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineDeviceStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineDeviceState(), nil
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Display name of the device
func (m *SecurityBaselineDeviceState) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineDeviceState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityBaselineComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*SecurityBaselineComplianceState))
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
// GetLastReportedDateTime gets the lastReportedDateTime property value. Last modified date time of the policy report
func (m *SecurityBaselineDeviceState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. Intune device id
func (m *SecurityBaselineDeviceState) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetState gets the state property value. Security Baseline Compliance State
func (m *SecurityBaselineDeviceState) GetState()(*SecurityBaselineComplianceState) {
    return m.state
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name
func (m *SecurityBaselineDeviceState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *SecurityBaselineDeviceState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
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
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Display name of the device
func (m *SecurityBaselineDeviceState) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last modified date time of the policy report
func (m *SecurityBaselineDeviceState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. Intune device id
func (m *SecurityBaselineDeviceState) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetState sets the state property value. Security Baseline Compliance State
func (m *SecurityBaselineDeviceState) SetState(value *SecurityBaselineComplianceState)() {
    m.state = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name
func (m *SecurityBaselineDeviceState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
