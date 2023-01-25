package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppPowerShellScriptRequirement 
type Win32LobAppPowerShellScriptRequirement struct {
    Win32LobAppRequirement
    // Contains all supported Powershell Script output detection type.
    detectionType *Win32LobAppPowerShellScriptDetectionType
    // The unique display name for this rule
    displayName *string
    // A value indicating whether signature check is enforced
    enforceSignatureCheck *bool
    // A value indicating whether this script should run as 32-bit
    runAs32Bit *bool
    // Indicates the type of execution context the app runs in.
    runAsAccount *RunAsAccountType
    // The base64 encoded script content to detect Win32 Line of Business (LoB) app
    scriptContent *string
}
// NewWin32LobAppPowerShellScriptRequirement instantiates a new Win32LobAppPowerShellScriptRequirement and sets the default values.
func NewWin32LobAppPowerShellScriptRequirement()(*Win32LobAppPowerShellScriptRequirement) {
    m := &Win32LobAppPowerShellScriptRequirement{
        Win32LobAppRequirement: *NewWin32LobAppRequirement(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppPowerShellScriptRequirement";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWin32LobAppPowerShellScriptRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppPowerShellScriptRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppPowerShellScriptRequirement(), nil
}
// GetDetectionType gets the detectionType property value. Contains all supported Powershell Script output detection type.
func (m *Win32LobAppPowerShellScriptRequirement) GetDetectionType()(*Win32LobAppPowerShellScriptDetectionType) {
    return m.detectionType
}
// GetDisplayName gets the displayName property value. The unique display name for this rule
func (m *Win32LobAppPowerShellScriptRequirement) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. A value indicating whether signature check is enforced
func (m *Win32LobAppPowerShellScriptRequirement) GetEnforceSignatureCheck()(*bool) {
    return m.enforceSignatureCheck
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppPowerShellScriptRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppRequirement.GetFieldDeserializers()
    res["detectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppPowerShellScriptDetectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionType(val.(*Win32LobAppPowerShellScriptDetectionType))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
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
    res["runAsAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
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
func (m *Win32LobAppPowerShellScriptRequirement) GetRunAs32Bit()(*bool) {
    return m.runAs32Bit
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *Win32LobAppPowerShellScriptRequirement) GetRunAsAccount()(*RunAsAccountType) {
    return m.runAsAccount
}
// GetScriptContent gets the scriptContent property value. The base64 encoded script content to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppPowerShellScriptRequirement) GetScriptContent()(*string) {
    return m.scriptContent
}
// Serialize serializes information the current object
func (m *Win32LobAppPowerShellScriptRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetectionType() != nil {
        cast := (*m.GetDetectionType()).String()
        err = writer.WriteStringValue("detectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
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
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err = writer.WriteStringValue("runAsAccount", &cast)
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
// SetDetectionType sets the detectionType property value. Contains all supported Powershell Script output detection type.
func (m *Win32LobAppPowerShellScriptRequirement) SetDetectionType(value *Win32LobAppPowerShellScriptDetectionType)() {
    m.detectionType = value
}
// SetDisplayName sets the displayName property value. The unique display name for this rule
func (m *Win32LobAppPowerShellScriptRequirement) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. A value indicating whether signature check is enforced
func (m *Win32LobAppPowerShellScriptRequirement) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
// SetRunAs32Bit sets the runAs32Bit property value. A value indicating whether this script should run as 32-bit
func (m *Win32LobAppPowerShellScriptRequirement) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *Win32LobAppPowerShellScriptRequirement) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// SetScriptContent sets the scriptContent property value. The base64 encoded script content to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppPowerShellScriptRequirement) SetScriptContent(value *string)() {
    m.scriptContent = value
}
