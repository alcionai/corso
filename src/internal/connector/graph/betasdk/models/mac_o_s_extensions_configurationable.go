package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSExtensionsConfigurationable 
type MacOSExtensionsConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKernelExtensionAllowedTeamIdentifiers()([]string)
    GetKernelExtensionOverridesAllowed()(*bool)
    GetKernelExtensionsAllowed()([]MacOSKernelExtensionable)
    GetSystemExtensionsAllowed()([]MacOSSystemExtensionable)
    GetSystemExtensionsAllowedTeamIdentifiers()([]string)
    GetSystemExtensionsAllowedTypes()([]MacOSSystemExtensionTypeMappingable)
    GetSystemExtensionsBlockOverride()(*bool)
    SetKernelExtensionAllowedTeamIdentifiers(value []string)()
    SetKernelExtensionOverridesAllowed(value *bool)()
    SetKernelExtensionsAllowed(value []MacOSKernelExtensionable)()
    SetSystemExtensionsAllowed(value []MacOSSystemExtensionable)()
    SetSystemExtensionsAllowedTeamIdentifiers(value []string)()
    SetSystemExtensionsAllowedTypes(value []MacOSSystemExtensionTypeMappingable)()
    SetSystemExtensionsBlockOverride(value *bool)()
}
