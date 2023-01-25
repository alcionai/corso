package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResetPasscodeActionResult 
type ResetPasscodeActionResult struct {
    DeviceActionResult
    // RotateBitLockerKeys action error code. Valid values 0 to 2147483647
    errorCode *int32
    // Newly generated passcode for the device
    passcode *string
}
// NewResetPasscodeActionResult instantiates a new ResetPasscodeActionResult and sets the default values.
func NewResetPasscodeActionResult()(*ResetPasscodeActionResult) {
    m := &ResetPasscodeActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateResetPasscodeActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResetPasscodeActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResetPasscodeActionResult(), nil
}
// GetErrorCode gets the errorCode property value. RotateBitLockerKeys action error code. Valid values 0 to 2147483647
func (m *ResetPasscodeActionResult) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResetPasscodeActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
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
    res["passcode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscode(val)
        }
        return nil
    }
    return res
}
// GetPasscode gets the passcode property value. Newly generated passcode for the device
func (m *ResetPasscodeActionResult) GetPasscode()(*string) {
    return m.passcode
}
// Serialize serializes information the current object
func (m *ResetPasscodeActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
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
        err = writer.WriteStringValue("passcode", m.GetPasscode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorCode sets the errorCode property value. RotateBitLockerKeys action error code. Valid values 0 to 2147483647
func (m *ResetPasscodeActionResult) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetPasscode sets the passcode property value. Newly generated passcode for the device
func (m *ResetPasscodeActionResult) SetPasscode(value *string)() {
    m.passcode = value
}
