package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppPowerShellScriptRequirementable 
type Win32LobAppPowerShellScriptRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRequirementable
    GetDetectionType()(*Win32LobAppPowerShellScriptDetectionType)
    GetDisplayName()(*string)
    GetEnforceSignatureCheck()(*bool)
    GetRunAs32Bit()(*bool)
    GetRunAsAccount()(*RunAsAccountType)
    GetScriptContent()(*string)
    SetDetectionType(value *Win32LobAppPowerShellScriptDetectionType)()
    SetDisplayName(value *string)()
    SetEnforceSignatureCheck(value *bool)()
    SetRunAs32Bit(value *bool)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetScriptContent(value *string)()
}
