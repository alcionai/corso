package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type IosKioskModeAppType int

const (
    // Device default value, no intent.
    NOTCONFIGURED_IOSKIOSKMODEAPPTYPE IosKioskModeAppType = iota
    // The app to be run comes from the app store.
    APPSTOREAPP_IOSKIOSKMODEAPPTYPE
    // The app to be run is built into the device.
    MANAGEDAPP_IOSKIOSKMODEAPPTYPE
    // The app to be run is a managed app.
    BUILTINAPP_IOSKIOSKMODEAPPTYPE
)

func (i IosKioskModeAppType) String() string {
    return []string{"notConfigured", "appStoreApp", "managedApp", "builtInApp"}[i]
}
func ParseIosKioskModeAppType(v string) (interface{}, error) {
    result := NOTCONFIGURED_IOSKIOSKMODEAPPTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_IOSKIOSKMODEAPPTYPE
        case "appStoreApp":
            result = APPSTOREAPP_IOSKIOSKMODEAPPTYPE
        case "managedApp":
            result = MANAGEDAPP_IOSKIOSKMODEAPPTYPE
        case "builtInApp":
            result = BUILTINAPP_IOSKIOSKMODEAPPTYPE
        default:
            return 0, errors.New("Unknown IosKioskModeAppType value: " + v)
    }
    return &result, nil
}
func SerializeIosKioskModeAppType(values []IosKioskModeAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
