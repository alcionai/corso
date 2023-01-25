package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppPowerShellScriptDetectionable 
type Win32LobAppPowerShellScriptDetectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppDetectionable
    GetEnforceSignatureCheck()(*bool)
    GetRunAs32Bit()(*bool)
    GetScriptContent()(*string)
    SetEnforceSignatureCheck(value *bool)()
    SetRunAs32Bit(value *bool)()
    SetScriptContent(value *string)()
}
