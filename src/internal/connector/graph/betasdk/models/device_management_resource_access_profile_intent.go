package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementResourceAccessProfileIntent int

const (
    // Apply the profile.
    APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT DeviceManagementResourceAccessProfileIntent = iota
    // Remove the profile from devices that have installed the profile.
    REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
)

func (i DeviceManagementResourceAccessProfileIntent) String() string {
    return []string{"apply", "remove"}[i]
}
func ParseDeviceManagementResourceAccessProfileIntent(v string) (interface{}, error) {
    result := APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
    switch v {
        case "apply":
            result = APPLY_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
        case "remove":
            result = REMOVE_DEVICEMANAGEMENTRESOURCEACCESSPROFILEINTENT
        default:
            return 0, errors.New("Unknown DeviceManagementResourceAccessProfileIntent value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementResourceAccessProfileIntent(values []DeviceManagementResourceAccessProfileIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
