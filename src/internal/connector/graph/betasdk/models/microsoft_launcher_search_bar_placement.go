package models
import (
    "errors"
)
// Provides operations to call the add method.
type MicrosoftLauncherSearchBarPlacement int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_MICROSOFTLAUNCHERSEARCHBARPLACEMENT MicrosoftLauncherSearchBarPlacement = iota
    // Indicates that the search bar will be displayed on the top of the device.
    TOP_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
    // Indicates that the search bar will be displayed on the bottom of the device.
    BOTTOM_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
    // Indicates that the search bar will be hidden on the device.
    HIDE_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
)

func (i MicrosoftLauncherSearchBarPlacement) String() string {
    return []string{"notConfigured", "top", "bottom", "hide"}[i]
}
func ParseMicrosoftLauncherSearchBarPlacement(v string) (interface{}, error) {
    result := NOTCONFIGURED_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
        case "top":
            result = TOP_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
        case "bottom":
            result = BOTTOM_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
        case "hide":
            result = HIDE_MICROSOFTLAUNCHERSEARCHBARPLACEMENT
        default:
            return 0, errors.New("Unknown MicrosoftLauncherSearchBarPlacement value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftLauncherSearchBarPlacement(values []MicrosoftLauncherSearchBarPlacement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
