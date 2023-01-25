package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsAppStartLayoutTileSize int

const (
    // Hidden.
    HIDDEN_WINDOWSAPPSTARTLAYOUTTILESIZE WindowsAppStartLayoutTileSize = iota
    // Small 1x1.
    SMALL_WINDOWSAPPSTARTLAYOUTTILESIZE
    // Medium 2x2.
    MEDIUM_WINDOWSAPPSTARTLAYOUTTILESIZE
    // Wide 4x2.
    WIDE_WINDOWSAPPSTARTLAYOUTTILESIZE
    // Large 4x4.
    LARGE_WINDOWSAPPSTARTLAYOUTTILESIZE
)

func (i WindowsAppStartLayoutTileSize) String() string {
    return []string{"hidden", "small", "medium", "wide", "large"}[i]
}
func ParseWindowsAppStartLayoutTileSize(v string) (interface{}, error) {
    result := HIDDEN_WINDOWSAPPSTARTLAYOUTTILESIZE
    switch v {
        case "hidden":
            result = HIDDEN_WINDOWSAPPSTARTLAYOUTTILESIZE
        case "small":
            result = SMALL_WINDOWSAPPSTARTLAYOUTTILESIZE
        case "medium":
            result = MEDIUM_WINDOWSAPPSTARTLAYOUTTILESIZE
        case "wide":
            result = WIDE_WINDOWSAPPSTARTLAYOUTTILESIZE
        case "large":
            result = LARGE_WINDOWSAPPSTARTLAYOUTTILESIZE
        default:
            return 0, errors.New("Unknown WindowsAppStartLayoutTileSize value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAppStartLayoutTileSize(values []WindowsAppStartLayoutTileSize) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
