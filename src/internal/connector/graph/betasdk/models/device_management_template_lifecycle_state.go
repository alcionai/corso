package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementTemplateLifecycleState int

const (
    // Invalid
    INVALID_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE DeviceManagementTemplateLifecycleState = iota
    // Draft
    DRAFT_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    // Active
    ACTIVE_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    // Superseded
    SUPERSEDED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    // Deprecated
    DEPRECATED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    // Retired
    RETIRED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
)

func (i DeviceManagementTemplateLifecycleState) String() string {
    return []string{"invalid", "draft", "active", "superseded", "deprecated", "retired"}[i]
}
func ParseDeviceManagementTemplateLifecycleState(v string) (interface{}, error) {
    result := INVALID_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
    switch v {
        case "invalid":
            result = INVALID_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "draft":
            result = DRAFT_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "active":
            result = ACTIVE_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "superseded":
            result = SUPERSEDED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "deprecated":
            result = DEPRECATED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        case "retired":
            result = RETIRED_DEVICEMANAGEMENTTEMPLATELIFECYCLESTATE
        default:
            return 0, errors.New("Unknown DeviceManagementTemplateLifecycleState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementTemplateLifecycleState(values []DeviceManagementTemplateLifecycleState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
