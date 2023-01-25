package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceComplianceLocalActionLockDeviceWithPasscode 
type AndroidDeviceComplianceLocalActionLockDeviceWithPasscode struct {
    AndroidDeviceComplianceLocalActionBase
    // Passcode to reset to Android device. This property is read-only.
    passcode *string
    // Number of sign in failures before wiping device, the value can be 4-11. Valid values 4 to 11
    passcodeSignInFailureCountBeforeWipe *int32
}
// NewAndroidDeviceComplianceLocalActionLockDeviceWithPasscode instantiates a new AndroidDeviceComplianceLocalActionLockDeviceWithPasscode and sets the default values.
func NewAndroidDeviceComplianceLocalActionLockDeviceWithPasscode()(*AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) {
    m := &AndroidDeviceComplianceLocalActionLockDeviceWithPasscode{
        AndroidDeviceComplianceLocalActionBase: *NewAndroidDeviceComplianceLocalActionBase(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceComplianceLocalActionLockDeviceWithPasscode";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceComplianceLocalActionLockDeviceWithPasscodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceComplianceLocalActionLockDeviceWithPasscodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceComplianceLocalActionLockDeviceWithPasscode(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceComplianceLocalActionBase.GetFieldDeserializers()
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
    res["passcodeSignInFailureCountBeforeWipe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeSignInFailureCountBeforeWipe(val)
        }
        return nil
    }
    return res
}
// GetPasscode gets the passcode property value. Passcode to reset to Android device. This property is read-only.
func (m *AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) GetPasscode()(*string) {
    return m.passcode
}
// GetPasscodeSignInFailureCountBeforeWipe gets the passcodeSignInFailureCountBeforeWipe property value. Number of sign in failures before wiping device, the value can be 4-11. Valid values 4 to 11
func (m *AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) GetPasscodeSignInFailureCountBeforeWipe()(*int32) {
    return m.passcodeSignInFailureCountBeforeWipe
}
// Serialize serializes information the current object
func (m *AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceComplianceLocalActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("passcodeSignInFailureCountBeforeWipe", m.GetPasscodeSignInFailureCountBeforeWipe())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPasscode sets the passcode property value. Passcode to reset to Android device. This property is read-only.
func (m *AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) SetPasscode(value *string)() {
    m.passcode = value
}
// SetPasscodeSignInFailureCountBeforeWipe sets the passcodeSignInFailureCountBeforeWipe property value. Number of sign in failures before wiping device, the value can be 4-11. Valid values 4 to 11
func (m *AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) SetPasscodeSignInFailureCountBeforeWipe(value *int32)() {
    m.passcodeSignInFailureCountBeforeWipe = value
}
