package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsPrivacyDataCategory int

const (
    // No access level specified, no intents. Device may behave either as in UserInControl or ForceAllow. It may depend on the privacy data been accessed, Windows versions and other factors.
    NOTCONFIGURED_WINDOWSPRIVACYDATACATEGORY WindowsPrivacyDataCategory = iota
    // Let apps access user’s name, picture and other account information created in Microsoft account. Added in Windows 10, version 1607.
    ACCOUNTINFO_WINDOWSPRIVACYDATACATEGORY
    // Allow apps to receive information, send notifications, and stay up-to-date, even when the user is not using them. Be aware that when disabling communication apps (Email, Voice, etc) from background access these apps may or may not function as they are with the background access. Added in Windows 10, version 1703.
    APPSRUNINBACKGROUND_WINDOWSPRIVACYDATACATEGORY
    // Let apps access user’s calendar. Added in Windows 10, version 1607.
    CALENDAR_WINDOWSPRIVACYDATACATEGORY
    // Let apps access user’s call history. Added in Windows 10, version 1607.
    CALLHISTORY_WINDOWSPRIVACYDATACATEGORY
    // Let apps access the camera on user’s device. Added in Windows 10, version 1607.
    CAMERA_WINDOWSPRIVACYDATACATEGORY
    // Let apps access user’s contact information. Added in Windows 10, version 1607.
    CONTACTS_WINDOWSPRIVACYDATACATEGORY
    // Let apps access diagnostic information about other running apps. Added in Windows 10, version 1703.
    DIAGNOSTICSINFO_WINDOWSPRIVACYDATACATEGORY
    // Let apps access and send email. Added in Windows 10, version 1607.
    EMAIL_WINDOWSPRIVACYDATACATEGORY
    // Let apps access the precise location data of device user. Added in Windows 10, version 1607.
    LOCATION_WINDOWSPRIVACYDATACATEGORY
    // Let apps read or send messages, text or MMS. Added in Windows 10, version 1607.
    MESSAGING_WINDOWSPRIVACYDATACATEGORY
    // Let apps use microphone on the user device. Added in Windows 10, version 1607.
    MICROPHONE_WINDOWSPRIVACYDATACATEGORY
    // Let apps use motion data generated on the device user. Added in Windows 10, version 1607.
    MOTION_WINDOWSPRIVACYDATACATEGORY
    // Let apps access user’s notifications. Added in Windows 10, version 1607.
    NOTIFICATIONS_WINDOWSPRIVACYDATACATEGORY
    // Let apps access phone data and make phone calls. Added in Windows 10, version 1607.
    PHONE_WINDOWSPRIVACYDATACATEGORY
    // Let apps use radios, including Bluetooth, to send and receive data. Added in Windows 10, version 1607.
    RADIOS_WINDOWSPRIVACYDATACATEGORY
    // Let apps access Task Scheduler. Added in Windows 10, version 1703.
    TASKS_WINDOWSPRIVACYDATACATEGORY
    // Let apps automatically share and sync info with wireless devices that don’t explicitly pair with user’s device. Added in Windows 10, version 1607.
    SYNCWITHDEVICES_WINDOWSPRIVACYDATACATEGORY
    // Let apps access trusted devices. Added in Windows 10, version 1607.
    TRUSTEDDEVICES_WINDOWSPRIVACYDATACATEGORY
)

func (i WindowsPrivacyDataCategory) String() string {
    return []string{"notConfigured", "accountInfo", "appsRunInBackground", "calendar", "callHistory", "camera", "contacts", "diagnosticsInfo", "email", "location", "messaging", "microphone", "motion", "notifications", "phone", "radios", "tasks", "syncWithDevices", "trustedDevices"}[i]
}
func ParseWindowsPrivacyDataCategory(v string) (interface{}, error) {
    result := NOTCONFIGURED_WINDOWSPRIVACYDATACATEGORY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWSPRIVACYDATACATEGORY
        case "accountInfo":
            result = ACCOUNTINFO_WINDOWSPRIVACYDATACATEGORY
        case "appsRunInBackground":
            result = APPSRUNINBACKGROUND_WINDOWSPRIVACYDATACATEGORY
        case "calendar":
            result = CALENDAR_WINDOWSPRIVACYDATACATEGORY
        case "callHistory":
            result = CALLHISTORY_WINDOWSPRIVACYDATACATEGORY
        case "camera":
            result = CAMERA_WINDOWSPRIVACYDATACATEGORY
        case "contacts":
            result = CONTACTS_WINDOWSPRIVACYDATACATEGORY
        case "diagnosticsInfo":
            result = DIAGNOSTICSINFO_WINDOWSPRIVACYDATACATEGORY
        case "email":
            result = EMAIL_WINDOWSPRIVACYDATACATEGORY
        case "location":
            result = LOCATION_WINDOWSPRIVACYDATACATEGORY
        case "messaging":
            result = MESSAGING_WINDOWSPRIVACYDATACATEGORY
        case "microphone":
            result = MICROPHONE_WINDOWSPRIVACYDATACATEGORY
        case "motion":
            result = MOTION_WINDOWSPRIVACYDATACATEGORY
        case "notifications":
            result = NOTIFICATIONS_WINDOWSPRIVACYDATACATEGORY
        case "phone":
            result = PHONE_WINDOWSPRIVACYDATACATEGORY
        case "radios":
            result = RADIOS_WINDOWSPRIVACYDATACATEGORY
        case "tasks":
            result = TASKS_WINDOWSPRIVACYDATACATEGORY
        case "syncWithDevices":
            result = SYNCWITHDEVICES_WINDOWSPRIVACYDATACATEGORY
        case "trustedDevices":
            result = TRUSTEDDEVICES_WINDOWSPRIVACYDATACATEGORY
        default:
            return 0, errors.New("Unknown WindowsPrivacyDataCategory value: " + v)
    }
    return &result, nil
}
func SerializeWindowsPrivacyDataCategory(values []WindowsPrivacyDataCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
