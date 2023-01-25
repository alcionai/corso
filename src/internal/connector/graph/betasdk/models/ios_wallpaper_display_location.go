package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type IosWallpaperDisplayLocation int

const (
    // No location specified for wallpaper display.
    NOTCONFIGURED_IOSWALLPAPERDISPLAYLOCATION IosWallpaperDisplayLocation = iota
    // A configured wallpaper image is displayed on Lock screen.
    LOCKSCREEN_IOSWALLPAPERDISPLAYLOCATION
    // A configured wallpaper image is displayed on Home (icon list) screen.
    HOMESCREEN_IOSWALLPAPERDISPLAYLOCATION
    // A configured wallpaper image is displayed on Lock screen and Home screen.
    LOCKANDHOMESCREENS_IOSWALLPAPERDISPLAYLOCATION
)

func (i IosWallpaperDisplayLocation) String() string {
    return []string{"notConfigured", "lockScreen", "homeScreen", "lockAndHomeScreens"}[i]
}
func ParseIosWallpaperDisplayLocation(v string) (interface{}, error) {
    result := NOTCONFIGURED_IOSWALLPAPERDISPLAYLOCATION
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_IOSWALLPAPERDISPLAYLOCATION
        case "lockScreen":
            result = LOCKSCREEN_IOSWALLPAPERDISPLAYLOCATION
        case "homeScreen":
            result = HOMESCREEN_IOSWALLPAPERDISPLAYLOCATION
        case "lockAndHomeScreens":
            result = LOCKANDHOMESCREENS_IOSWALLPAPERDISPLAYLOCATION
        default:
            return 0, errors.New("Unknown IosWallpaperDisplayLocation value: " + v)
    }
    return &result, nil
}
func SerializeIosWallpaperDisplayLocation(values []IosWallpaperDisplayLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
