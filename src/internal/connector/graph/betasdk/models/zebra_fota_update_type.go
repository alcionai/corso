package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ZebraFotaUpdateType int

const (
    // Custom update where the user selects specific BSP, OS version, and patch number to update to.
    CUSTOM_ZEBRAFOTAUPDATETYPE ZebraFotaUpdateType = iota
    // The latest released update becomes the target OS. Latest may update the device to a new Android version.
    LATEST_ZEBRAFOTAUPDATETYPE
    // The device always looks for the latest package available in the repo and tries to update whenever a new package is available. This continues until the admin cancels the auto update.
    AUTO_ZEBRAFOTAUPDATETYPE
    // Unknown future enum value.
    UNKNOWNFUTUREVALUE_ZEBRAFOTAUPDATETYPE
)

func (i ZebraFotaUpdateType) String() string {
    return []string{"custom", "latest", "auto", "unknownFutureValue"}[i]
}
func ParseZebraFotaUpdateType(v string) (interface{}, error) {
    result := CUSTOM_ZEBRAFOTAUPDATETYPE
    switch v {
        case "custom":
            result = CUSTOM_ZEBRAFOTAUPDATETYPE
        case "latest":
            result = LATEST_ZEBRAFOTAUPDATETYPE
        case "auto":
            result = AUTO_ZEBRAFOTAUPDATETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTAUPDATETYPE
        default:
            return 0, errors.New("Unknown ZebraFotaUpdateType value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaUpdateType(values []ZebraFotaUpdateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
