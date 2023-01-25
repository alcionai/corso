package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosDeviceFeaturesConfigurationable 
type IosDeviceFeaturesConfigurationable interface {
    AppleDeviceFeaturesConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssetTagTemplate()(*string)
    GetContentFilterSettings()(IosWebContentFilterBaseable)
    GetHomeScreenDockIcons()([]IosHomeScreenItemable)
    GetHomeScreenGridHeight()(*int32)
    GetHomeScreenGridWidth()(*int32)
    GetHomeScreenPages()([]IosHomeScreenPageable)
    GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable)
    GetIosSingleSignOnExtension()(IosSingleSignOnExtensionable)
    GetLockScreenFootnote()(*string)
    GetNotificationSettings()([]IosNotificationSettingsable)
    GetSingleSignOnExtension()(SingleSignOnExtensionable)
    GetSingleSignOnExtensionPkinitCertificate()(IosCertificateProfileBaseable)
    GetSingleSignOnSettings()(IosSingleSignOnSettingsable)
    GetWallpaperDisplayLocation()(*IosWallpaperDisplayLocation)
    GetWallpaperImage()(MimeContentable)
    SetAssetTagTemplate(value *string)()
    SetContentFilterSettings(value IosWebContentFilterBaseable)()
    SetHomeScreenDockIcons(value []IosHomeScreenItemable)()
    SetHomeScreenGridHeight(value *int32)()
    SetHomeScreenGridWidth(value *int32)()
    SetHomeScreenPages(value []IosHomeScreenPageable)()
    SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)()
    SetIosSingleSignOnExtension(value IosSingleSignOnExtensionable)()
    SetLockScreenFootnote(value *string)()
    SetNotificationSettings(value []IosNotificationSettingsable)()
    SetSingleSignOnExtension(value SingleSignOnExtensionable)()
    SetSingleSignOnExtensionPkinitCertificate(value IosCertificateProfileBaseable)()
    SetSingleSignOnSettings(value IosSingleSignOnSettingsable)()
    SetWallpaperDisplayLocation(value *IosWallpaperDisplayLocation)()
    SetWallpaperImage(value MimeContentable)()
}
