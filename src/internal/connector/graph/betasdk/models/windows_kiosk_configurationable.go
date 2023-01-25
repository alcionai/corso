package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskConfigurationable 
type WindowsKioskConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEdgeKioskEnablePublicBrowsing()(*bool)
    GetKioskBrowserBlockedUrlExceptions()([]string)
    GetKioskBrowserBlockedURLs()([]string)
    GetKioskBrowserDefaultUrl()(*string)
    GetKioskBrowserEnableEndSessionButton()(*bool)
    GetKioskBrowserEnableHomeButton()(*bool)
    GetKioskBrowserEnableNavigationButtons()(*bool)
    GetKioskBrowserRestartOnIdleTimeInMinutes()(*int32)
    GetKioskProfiles()([]WindowsKioskProfileable)
    GetWindowsKioskForceUpdateSchedule()(WindowsKioskForceUpdateScheduleable)
    SetEdgeKioskEnablePublicBrowsing(value *bool)()
    SetKioskBrowserBlockedUrlExceptions(value []string)()
    SetKioskBrowserBlockedURLs(value []string)()
    SetKioskBrowserDefaultUrl(value *string)()
    SetKioskBrowserEnableEndSessionButton(value *bool)()
    SetKioskBrowserEnableHomeButton(value *bool)()
    SetKioskBrowserEnableNavigationButtons(value *bool)()
    SetKioskBrowserRestartOnIdleTimeInMinutes(value *int32)()
    SetKioskProfiles(value []WindowsKioskProfileable)()
    SetWindowsKioskForceUpdateSchedule(value WindowsKioskForceUpdateScheduleable)()
}
