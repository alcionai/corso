package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidDeviceOwnerKioskModeFolderIcon int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON AndroidDeviceOwnerKioskModeFolderIcon = iota
    // Folder icon appears as dark square.
    DARKSQUARE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
    // Folder icon appears as dark circle.
    DARKCIRCLE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
    // Folder icon appears as light square.
    LIGHTSQUARE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
    // Folder icon appears as light circle  .
    LIGHTCIRCLE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
)

func (i AndroidDeviceOwnerKioskModeFolderIcon) String() string {
    return []string{"notConfigured", "darkSquare", "darkCircle", "lightSquare", "lightCircle"}[i]
}
func ParseAndroidDeviceOwnerKioskModeFolderIcon(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
        case "darkSquare":
            result = DARKSQUARE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
        case "darkCircle":
            result = DARKCIRCLE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
        case "lightSquare":
            result = LIGHTSQUARE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
        case "lightCircle":
            result = LIGHTCIRCLE_ANDROIDDEVICEOWNERKIOSKMODEFOLDERICON
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerKioskModeFolderIcon value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerKioskModeFolderIcon(values []AndroidDeviceOwnerKioskModeFolderIcon) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
