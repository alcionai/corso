package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppPowerShellScriptDetection 
type Win32LobAppPowerShellScriptDetection struct {
    Win32LobAppDetection
    // A value indicating whether signature check is enforced
    enforceSignatureCheck *bool
    // A value indicating whether this script should run as 32-bit
    runAs32Bit *bool
    // The base64 encoded script content to detect Win32 Line of Business (LoB) app
    scriptContent *string
}
// NewWin32LobAppPowerShellScriptDetection instantiates a new Win32LobAppPowerShellScriptDetection and sets the default values.
func NewWin32LobAppPowerShellScriptDetection()(*Win32LobAppPowerShellScriptDetection) {
    m := &Win32LobAppPowerShellScriptDetection{
        Win32LobAppDetection: *NewWin32LobAppDetection(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppPowerShellScriptDetection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWin32LobAppPowerShellScriptDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppPowerShellScriptDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppPowerShellScriptDetection(), nil
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. A value indicating whether signature check is enforced
func (m *Win32LobAppPowerShellScriptDetection) GetEnforceSignatureCheck()(*bool) {
    return m.enforceSignatureCheck
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppPowerShellScriptDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppDetection.GetFieldDeserializers()
    res["enforceSignatureCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceSignatureCheck(val)
        }
        return nil
    }
    res["runAs32Bit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAs32Bit(val)
        }
        return nil
    }
    res["scriptContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptContent(val)
        }
        return nil
    }
    return res
}
// GetRunAs32Bit gets the runAs32Bit property value. A value indicating whether this script should run as 32-bit
func (m *Win32LobAppPowerShellScriptDetection) GetRunAs32Bit()(*bool) {
    return m.runAs32Bit
}
// GetScriptContent gets the scriptContent property value. The base64 encoded script content to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppPowerShellScriptDetection) GetScriptContent()(*string) {
    return m.scriptContent
}
// Serialize serializes information the current object
func (m *Win32LobAppPowerShellScriptDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppDetection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enforceSignatureCheck", m.GetEnforceSignatureCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("runAs32Bit", m.GetRunAs32Bit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scriptContent", m.GetScriptContent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. A value indicating whether signature check is enforced
func (m *Win32LobAppPowerShellScriptDetection) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
// SetRunAs32Bit sets the runAs32Bit property value. A value indicating whether this script should run as 32-bit
func (m *Win32LobAppPowerShellScriptDetection) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
// SetScriptContent sets the scriptContent property value. The base64 encoded script content to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppPowerShellScriptDetection) SetScriptContent(value *string)() {
    m.scriptContent = value
}
