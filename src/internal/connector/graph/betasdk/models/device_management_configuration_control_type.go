package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementConfigurationControlType int

const (
    // Don’t override default
    DEFAULT_ESCAPED_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE DeviceManagementConfigurationControlType = iota
    // Display Choice in dropdown
    DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Display text input in small text input
    SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Display text input in large text input
    LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Allow for toggle control type
    TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Allow for multiheader grid control type
    MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Allow for context pane control type
    CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
)

func (i DeviceManagementConfigurationControlType) String() string {
    return []string{"default", "dropdown", "smallTextBox", "largeTextBox", "toggle", "multiheaderGrid", "contextPane"}[i]
}
func ParseDeviceManagementConfigurationControlType(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    switch v {
        case "default":
            result = DEFAULT_ESCAPED_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "dropdown":
            result = DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "smallTextBox":
            result = SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "largeTextBox":
            result = LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "toggle":
            result = TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "multiheaderGrid":
            result = MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "contextPane":
            result = CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationControlType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationControlType(values []DeviceManagementConfigurationControlType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
