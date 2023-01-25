package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingAppPolicyCreationHistory 
type MobileAppTroubleshootingAppPolicyCreationHistory struct {
    MobileAppTroubleshootingHistoryItem
    // Error code for the failure, empty if no failure.
    errorCode *string
    // Indicates the type of execution status of the device management script.
    runState *RunState
}
// NewMobileAppTroubleshootingAppPolicyCreationHistory instantiates a new MobileAppTroubleshootingAppPolicyCreationHistory and sets the default values.
func NewMobileAppTroubleshootingAppPolicyCreationHistory()(*MobileAppTroubleshootingAppPolicyCreationHistory) {
    m := &MobileAppTroubleshootingAppPolicyCreationHistory{
        MobileAppTroubleshootingHistoryItem: *NewMobileAppTroubleshootingHistoryItem(),
    }
    return m
}
// CreateMobileAppTroubleshootingAppPolicyCreationHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppTroubleshootingAppPolicyCreationHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppTroubleshootingAppPolicyCreationHistory(), nil
}
// GetErrorCode gets the errorCode property value. Error code for the failure, empty if no failure.
func (m *MobileAppTroubleshootingAppPolicyCreationHistory) GetErrorCode()(*string) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppTroubleshootingAppPolicyCreationHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppTroubleshootingHistoryItem.GetFieldDeserializers()
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
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
// GetRunState gets the runState property value. Indicates the type of execution status of the device management script.
func (m *MobileAppTroubleshootingAppPolicyCreationHistory) GetRunState()(*RunState) {
    return m.runState
}
// Serialize serializes information the current object
func (m *MobileAppTroubleshootingAppPolicyCreationHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppTroubleshootingHistoryItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
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
// SetErrorCode sets the errorCode property value. Error code for the failure, empty if no failure.
func (m *MobileAppTroubleshootingAppPolicyCreationHistory) SetErrorCode(value *string)() {
    m.errorCode = value
}
// SetRunState sets the runState property value. Indicates the type of execution status of the device management script.
func (m *MobileAppTroubleshootingAppPolicyCreationHistory) SetRunState(value *RunState)() {
    m.runState = value
}
