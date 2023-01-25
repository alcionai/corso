package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosDeviceFeaturesConfiguration 
type IosDeviceFeaturesConfiguration struct {
    AppleDeviceFeaturesConfigurationBase
    // Asset tag information for the device, displayed on the login window and lock screen.
    assetTagTemplate *string
    // Gets or sets iOS Web Content Filter settings, supervised mode only
    contentFilterSettings IosWebContentFilterBaseable
    // A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements.
    homeScreenDockIcons []IosHomeScreenItemable
    // Gets or sets the number of rows to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridWidth must be configured as well.
    homeScreenGridHeight *int32
    // Gets or sets the number of columns to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridHeight must be configured as well.
    homeScreenGridWidth *int32
    // A list of pages on the Home Screen. This collection can contain a maximum of 500 elements.
    homeScreenPages []IosHomeScreenPageable
    // Identity Certificate for the renewal of Kerberos ticket used in single sign-on settings.
    identityCertificateForClientAuthentication IosCertificateProfileBaseable
    // Gets or sets a single sign-on extension profile.
    iosSingleSignOnExtension IosSingleSignOnExtensionable
    // A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later.
    lockScreenFootnote *string
    // Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This collection can contain a maximum of 500 elements.
    notificationSettings []IosNotificationSettingsable
    // Gets or sets a single sign-on extension profile. Deprecated: use IOSSingleSignOnExtension instead.
    singleSignOnExtension SingleSignOnExtensionable
    // PKINIT Certificate for the authentication with single sign-on extension settings.
    singleSignOnExtensionPkinitCertificate IosCertificateProfileBaseable
    // The Kerberos login settings that enable apps on receiving devices to authenticate smoothly.
    singleSignOnSettings IosSingleSignOnSettingsable
    // An enum type for wallpaper display location specifier.
    wallpaperDisplayLocation *IosWallpaperDisplayLocation
    // A wallpaper image must be in either PNG or JPEG format. It requires a supervised device with iOS 8 or later version.
    wallpaperImage MimeContentable
}
// NewIosDeviceFeaturesConfiguration instantiates a new IosDeviceFeaturesConfiguration and sets the default values.
func NewIosDeviceFeaturesConfiguration()(*IosDeviceFeaturesConfiguration) {
    m := &IosDeviceFeaturesConfiguration{
        AppleDeviceFeaturesConfigurationBase: *NewAppleDeviceFeaturesConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.iosDeviceFeaturesConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosDeviceFeaturesConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosDeviceFeaturesConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosDeviceFeaturesConfiguration(), nil
}
// GetAssetTagTemplate gets the assetTagTemplate property value. Asset tag information for the device, displayed on the login window and lock screen.
func (m *IosDeviceFeaturesConfiguration) GetAssetTagTemplate()(*string) {
    return m.assetTagTemplate
}
// GetContentFilterSettings gets the contentFilterSettings property value. Gets or sets iOS Web Content Filter settings, supervised mode only
func (m *IosDeviceFeaturesConfiguration) GetContentFilterSettings()(IosWebContentFilterBaseable) {
    return m.contentFilterSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosDeviceFeaturesConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleDeviceFeaturesConfigurationBase.GetFieldDeserializers()
    res["assetTagTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetTagTemplate(val)
        }
        return nil
    }
    res["contentFilterSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosWebContentFilterBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentFilterSettings(val.(IosWebContentFilterBaseable))
        }
        return nil
    }
    res["homeScreenDockIcons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosHomeScreenItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosHomeScreenItemable, len(val))
            for i, v := range val {
                res[i] = v.(IosHomeScreenItemable)
            }
            m.SetHomeScreenDockIcons(res)
        }
        return nil
    }
    res["homeScreenGridHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeScreenGridHeight(val)
        }
        return nil
    }
    res["homeScreenGridWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeScreenGridWidth(val)
        }
        return nil
    }
    res["homeScreenPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosHomeScreenPageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosHomeScreenPageable, len(val))
            for i, v := range val {
                res[i] = v.(IosHomeScreenPageable)
            }
            m.SetHomeScreenPages(res)
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(IosCertificateProfileBaseable))
        }
        return nil
    }
    res["iosSingleSignOnExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosSingleSignOnExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosSingleSignOnExtension(val.(IosSingleSignOnExtensionable))
        }
        return nil
    }
    res["lockScreenFootnote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenFootnote(val)
        }
        return nil
    }
    res["notificationSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosNotificationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosNotificationSettingsable, len(val))
            for i, v := range val {
                res[i] = v.(IosNotificationSettingsable)
            }
            m.SetNotificationSettings(res)
        }
        return nil
    }
    res["singleSignOnExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSingleSignOnExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnExtension(val.(SingleSignOnExtensionable))
        }
        return nil
    }
    res["singleSignOnExtensionPkinitCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnExtensionPkinitCertificate(val.(IosCertificateProfileBaseable))
        }
        return nil
    }
    res["singleSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosSingleSignOnSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnSettings(val.(IosSingleSignOnSettingsable))
        }
        return nil
    }
    res["wallpaperDisplayLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIosWallpaperDisplayLocation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWallpaperDisplayLocation(val.(*IosWallpaperDisplayLocation))
        }
        return nil
    }
    res["wallpaperImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWallpaperImage(val.(MimeContentable))
        }
        return nil
    }
    return res
}
// GetHomeScreenDockIcons gets the homeScreenDockIcons property value. A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenDockIcons()([]IosHomeScreenItemable) {
    return m.homeScreenDockIcons
}
// GetHomeScreenGridHeight gets the homeScreenGridHeight property value. Gets or sets the number of rows to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridWidth must be configured as well.
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenGridHeight()(*int32) {
    return m.homeScreenGridHeight
}
// GetHomeScreenGridWidth gets the homeScreenGridWidth property value. Gets or sets the number of columns to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridHeight must be configured as well.
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenGridWidth()(*int32) {
    return m.homeScreenGridWidth
}
// GetHomeScreenPages gets the homeScreenPages property value. A list of pages on the Home Screen. This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenPages()([]IosHomeScreenPageable) {
    return m.homeScreenPages
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for the renewal of Kerberos ticket used in single sign-on settings.
func (m *IosDeviceFeaturesConfiguration) GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable) {
    return m.identityCertificateForClientAuthentication
}
// GetIosSingleSignOnExtension gets the iosSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
func (m *IosDeviceFeaturesConfiguration) GetIosSingleSignOnExtension()(IosSingleSignOnExtensionable) {
    return m.iosSingleSignOnExtension
}
// GetLockScreenFootnote gets the lockScreenFootnote property value. A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later.
func (m *IosDeviceFeaturesConfiguration) GetLockScreenFootnote()(*string) {
    return m.lockScreenFootnote
}
// GetNotificationSettings gets the notificationSettings property value. Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) GetNotificationSettings()([]IosNotificationSettingsable) {
    return m.notificationSettings
}
// GetSingleSignOnExtension gets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use IOSSingleSignOnExtension instead.
func (m *IosDeviceFeaturesConfiguration) GetSingleSignOnExtension()(SingleSignOnExtensionable) {
    return m.singleSignOnExtension
}
// GetSingleSignOnExtensionPkinitCertificate gets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extension settings.
func (m *IosDeviceFeaturesConfiguration) GetSingleSignOnExtensionPkinitCertificate()(IosCertificateProfileBaseable) {
    return m.singleSignOnExtensionPkinitCertificate
}
// GetSingleSignOnSettings gets the singleSignOnSettings property value. The Kerberos login settings that enable apps on receiving devices to authenticate smoothly.
func (m *IosDeviceFeaturesConfiguration) GetSingleSignOnSettings()(IosSingleSignOnSettingsable) {
    return m.singleSignOnSettings
}
// GetWallpaperDisplayLocation gets the wallpaperDisplayLocation property value. An enum type for wallpaper display location specifier.
func (m *IosDeviceFeaturesConfiguration) GetWallpaperDisplayLocation()(*IosWallpaperDisplayLocation) {
    return m.wallpaperDisplayLocation
}
// GetWallpaperImage gets the wallpaperImage property value. A wallpaper image must be in either PNG or JPEG format. It requires a supervised device with iOS 8 or later version.
func (m *IosDeviceFeaturesConfiguration) GetWallpaperImage()(MimeContentable) {
    return m.wallpaperImage
}
// Serialize serializes information the current object
func (m *IosDeviceFeaturesConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppleDeviceFeaturesConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assetTagTemplate", m.GetAssetTagTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentFilterSettings", m.GetContentFilterSettings())
        if err != nil {
            return err
        }
    }
    if m.GetHomeScreenDockIcons() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeScreenDockIcons()))
        for i, v := range m.GetHomeScreenDockIcons() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("homeScreenDockIcons", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("homeScreenGridHeight", m.GetHomeScreenGridHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("homeScreenGridWidth", m.GetHomeScreenGridWidth())
        if err != nil {
            return err
        }
    }
    if m.GetHomeScreenPages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeScreenPages()))
        for i, v := range m.GetHomeScreenPages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("homeScreenPages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificateForClientAuthentication", m.GetIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("iosSingleSignOnExtension", m.GetIosSingleSignOnExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lockScreenFootnote", m.GetLockScreenFootnote())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationSettings()))
        for i, v := range m.GetNotificationSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("notificationSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnExtension", m.GetSingleSignOnExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnExtensionPkinitCertificate", m.GetSingleSignOnExtensionPkinitCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnSettings", m.GetSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetWallpaperDisplayLocation() != nil {
        cast := (*m.GetWallpaperDisplayLocation()).String()
        err = writer.WriteStringValue("wallpaperDisplayLocation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("wallpaperImage", m.GetWallpaperImage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssetTagTemplate sets the assetTagTemplate property value. Asset tag information for the device, displayed on the login window and lock screen.
func (m *IosDeviceFeaturesConfiguration) SetAssetTagTemplate(value *string)() {
    m.assetTagTemplate = value
}
// SetContentFilterSettings sets the contentFilterSettings property value. Gets or sets iOS Web Content Filter settings, supervised mode only
func (m *IosDeviceFeaturesConfiguration) SetContentFilterSettings(value IosWebContentFilterBaseable)() {
    m.contentFilterSettings = value
}
// SetHomeScreenDockIcons sets the homeScreenDockIcons property value. A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenDockIcons(value []IosHomeScreenItemable)() {
    m.homeScreenDockIcons = value
}
// SetHomeScreenGridHeight sets the homeScreenGridHeight property value. Gets or sets the number of rows to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridWidth must be configured as well.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenGridHeight(value *int32)() {
    m.homeScreenGridHeight = value
}
// SetHomeScreenGridWidth sets the homeScreenGridWidth property value. Gets or sets the number of columns to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridHeight must be configured as well.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenGridWidth(value *int32)() {
    m.homeScreenGridWidth = value
}
// SetHomeScreenPages sets the homeScreenPages property value. A list of pages on the Home Screen. This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenPages(value []IosHomeScreenPageable)() {
    m.homeScreenPages = value
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for the renewal of Kerberos ticket used in single sign-on settings.
func (m *IosDeviceFeaturesConfiguration) SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)() {
    m.identityCertificateForClientAuthentication = value
}
// SetIosSingleSignOnExtension sets the iosSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
func (m *IosDeviceFeaturesConfiguration) SetIosSingleSignOnExtension(value IosSingleSignOnExtensionable)() {
    m.iosSingleSignOnExtension = value
}
// SetLockScreenFootnote sets the lockScreenFootnote property value. A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later.
func (m *IosDeviceFeaturesConfiguration) SetLockScreenFootnote(value *string)() {
    m.lockScreenFootnote = value
}
// SetNotificationSettings sets the notificationSettings property value. Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) SetNotificationSettings(value []IosNotificationSettingsable)() {
    m.notificationSettings = value
}
// SetSingleSignOnExtension sets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use IOSSingleSignOnExtension instead.
func (m *IosDeviceFeaturesConfiguration) SetSingleSignOnExtension(value SingleSignOnExtensionable)() {
    m.singleSignOnExtension = value
}
// SetSingleSignOnExtensionPkinitCertificate sets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extension settings.
func (m *IosDeviceFeaturesConfiguration) SetSingleSignOnExtensionPkinitCertificate(value IosCertificateProfileBaseable)() {
    m.singleSignOnExtensionPkinitCertificate = value
}
// SetSingleSignOnSettings sets the singleSignOnSettings property value. The Kerberos login settings that enable apps on receiving devices to authenticate smoothly.
func (m *IosDeviceFeaturesConfiguration) SetSingleSignOnSettings(value IosSingleSignOnSettingsable)() {
    m.singleSignOnSettings = value
}
// SetWallpaperDisplayLocation sets the wallpaperDisplayLocation property value. An enum type for wallpaper display location specifier.
func (m *IosDeviceFeaturesConfiguration) SetWallpaperDisplayLocation(value *IosWallpaperDisplayLocation)() {
    m.wallpaperDisplayLocation = value
}
// SetWallpaperImage sets the wallpaperImage property value. A wallpaper image must be in either PNG or JPEG format. It requires a supervised device with iOS 8 or later version.
func (m *IosDeviceFeaturesConfiguration) SetWallpaperImage(value MimeContentable)() {
    m.wallpaperImage = value
}
