package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type EmbeddedSIMDeviceStateValue int

const (
    // Designates that the embedded SIM activation code is free and available to be assigned to a device.
    NOTEVALUATED_EMBEDDEDSIMDEVICESTATEVALUE EmbeddedSIMDeviceStateValue = iota
    // Designates that Intune Service failed to deliver this profile to a device.
    FAILED_EMBEDDEDSIMDEVICESTATEVALUE
    // Designates that the embedded SIM activation code has been assigned to a device and the device is installing the token.
    INSTALLING_EMBEDDEDSIMDEVICESTATEVALUE
    // Designates that the embedded SIM activation code has been successfully installed on the target device.
    INSTALLED_EMBEDDEDSIMDEVICESTATEVALUE
    // Designates that Intune Service is trying to delete the profile from the device.
    DELETING_EMBEDDEDSIMDEVICESTATEVALUE
    // Designates that there is error with this profile.
    ERROR_EMBEDDEDSIMDEVICESTATEVALUE
    // Designates that the profile is deleted from the device.
    DELETED_EMBEDDEDSIMDEVICESTATEVALUE
    // Designates that the profile is removed from the device by the user
    REMOVEDBYUSER_EMBEDDEDSIMDEVICESTATEVALUE
)

func (i EmbeddedSIMDeviceStateValue) String() string {
    return []string{"notEvaluated", "failed", "installing", "installed", "deleting", "error", "deleted", "removedByUser"}[i]
}
func ParseEmbeddedSIMDeviceStateValue(v string) (interface{}, error) {
    result := NOTEVALUATED_EMBEDDEDSIMDEVICESTATEVALUE
    switch v {
        case "notEvaluated":
            result = NOTEVALUATED_EMBEDDEDSIMDEVICESTATEVALUE
        case "failed":
            result = FAILED_EMBEDDEDSIMDEVICESTATEVALUE
        case "installing":
            result = INSTALLING_EMBEDDEDSIMDEVICESTATEVALUE
        case "installed":
            result = INSTALLED_EMBEDDEDSIMDEVICESTATEVALUE
        case "deleting":
            result = DELETING_EMBEDDEDSIMDEVICESTATEVALUE
        case "error":
            result = ERROR_EMBEDDEDSIMDEVICESTATEVALUE
        case "deleted":
            result = DELETED_EMBEDDEDSIMDEVICESTATEVALUE
        case "removedByUser":
            result = REMOVEDBYUSER_EMBEDDEDSIMDEVICESTATEVALUE
        default:
            return 0, errors.New("Unknown EmbeddedSIMDeviceStateValue value: " + v)
    }
    return &result, nil
}
func SerializeEmbeddedSIMDeviceStateValue(values []EmbeddedSIMDeviceStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
