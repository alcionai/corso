package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosRedirectSingleSignOnExtensionable 
type IosRedirectSingleSignOnExtensionable interface {
    IosSingleSignOnExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurations()([]KeyTypedValuePairable)
    GetExtensionIdentifier()(*string)
    GetTeamIdentifier()(*string)
    GetUrlPrefixes()([]string)
    SetConfigurations(value []KeyTypedValuePairable)()
    SetExtensionIdentifier(value *string)()
    SetTeamIdentifier(value *string)()
    SetUrlPrefixes(value []string)()
}
