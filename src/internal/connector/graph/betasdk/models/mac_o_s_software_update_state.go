package models
import (
    "errors"
)
// Provides operations to call the add method.
type MacOSSoftwareUpdateState int

const (
    // The software update successfully installed
    SUCCESS_MACOSSOFTWAREUPDATESTATE MacOSSoftwareUpdateState = iota
    // The software update is being downloaded
    DOWNLOADING_MACOSSOFTWAREUPDATESTATE
    // The software update has been downloaded
    DOWNLOADED_MACOSSOFTWAREUPDATESTATE
    // The software update is being installed
    INSTALLING_MACOSSOFTWAREUPDATESTATE
    // No action is being taken on this software update
    IDLE_MACOSSOFTWAREUPDATESTATE
    // The software update is available on the device
    AVAILABLE_MACOSSOFTWAREUPDATESTATE
    // The software update has been scheduled on the device
    SCHEDULED_MACOSSOFTWAREUPDATESTATE
    // The software update download has failed
    DOWNLOADFAILED_MACOSSOFTWAREUPDATESTATE
    // There is not enough space to download the update
    DOWNLOADINSUFFICIENTSPACE_MACOSSOFTWAREUPDATESTATE
    // There is not enough power to download the update
    DOWNLOADINSUFFICIENTPOWER_MACOSSOFTWAREUPDATESTATE
    // There is insufficient network capacity to download the update
    DOWNLOADINSUFFICIENTNETWORK_MACOSSOFTWAREUPDATESTATE
    // There is not enough space to install the update
    INSTALLINSUFFICIENTSPACE_MACOSSOFTWAREUPDATESTATE
    // There is not enough power to install the update
    INSTALLINSUFFICIENTPOWER_MACOSSOFTWAREUPDATESTATE
    // Installation has failed for an unspecified reason
    INSTALLFAILED_MACOSSOFTWAREUPDATESTATE
    // The schedule update command has failed for an unspecified reason
    COMMANDFAILED_MACOSSOFTWAREUPDATESTATE
)

func (i MacOSSoftwareUpdateState) String() string {
    return []string{"success", "downloading", "downloaded", "installing", "idle", "available", "scheduled", "downloadFailed", "downloadInsufficientSpace", "downloadInsufficientPower", "downloadInsufficientNetwork", "installInsufficientSpace", "installInsufficientPower", "installFailed", "commandFailed"}[i]
}
func ParseMacOSSoftwareUpdateState(v string) (interface{}, error) {
    result := SUCCESS_MACOSSOFTWAREUPDATESTATE
    switch v {
        case "success":
            result = SUCCESS_MACOSSOFTWAREUPDATESTATE
        case "downloading":
            result = DOWNLOADING_MACOSSOFTWAREUPDATESTATE
        case "downloaded":
            result = DOWNLOADED_MACOSSOFTWAREUPDATESTATE
        case "installing":
            result = INSTALLING_MACOSSOFTWAREUPDATESTATE
        case "idle":
            result = IDLE_MACOSSOFTWAREUPDATESTATE
        case "available":
            result = AVAILABLE_MACOSSOFTWAREUPDATESTATE
        case "scheduled":
            result = SCHEDULED_MACOSSOFTWAREUPDATESTATE
        case "downloadFailed":
            result = DOWNLOADFAILED_MACOSSOFTWAREUPDATESTATE
        case "downloadInsufficientSpace":
            result = DOWNLOADINSUFFICIENTSPACE_MACOSSOFTWAREUPDATESTATE
        case "downloadInsufficientPower":
            result = DOWNLOADINSUFFICIENTPOWER_MACOSSOFTWAREUPDATESTATE
        case "downloadInsufficientNetwork":
            result = DOWNLOADINSUFFICIENTNETWORK_MACOSSOFTWAREUPDATESTATE
        case "installInsufficientSpace":
            result = INSTALLINSUFFICIENTSPACE_MACOSSOFTWAREUPDATESTATE
        case "installInsufficientPower":
            result = INSTALLINSUFFICIENTPOWER_MACOSSOFTWAREUPDATESTATE
        case "installFailed":
            result = INSTALLFAILED_MACOSSOFTWAREUPDATESTATE
        case "commandFailed":
            result = COMMANDFAILED_MACOSSOFTWAREUPDATESTATE
        default:
            return 0, errors.New("Unknown MacOSSoftwareUpdateState value: " + v)
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateState(values []MacOSSoftwareUpdateState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
