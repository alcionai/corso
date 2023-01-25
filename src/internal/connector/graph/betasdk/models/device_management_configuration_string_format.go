package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementConfigurationStringFormat int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT DeviceManagementConfigurationStringFormat = iota
    EMAIL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    GUID_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    IP_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    BASE64_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    URL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    VERSION_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    XML_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    DATE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    TIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    BINARY_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    REGEX_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    JSON_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    DATETIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    SURFACEHUB_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
)

func (i DeviceManagementConfigurationStringFormat) String() string {
    return []string{"none", "email", "guid", "ip", "base64", "url", "version", "xml", "date", "time", "binary", "regEx", "json", "dateTime", "surfaceHub"}[i]
}
func ParseDeviceManagementConfigurationStringFormat(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "email":
            result = EMAIL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "guid":
            result = GUID_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "ip":
            result = IP_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "base64":
            result = BASE64_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "url":
            result = URL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "version":
            result = VERSION_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "xml":
            result = XML_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "date":
            result = DATE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "time":
            result = TIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "binary":
            result = BINARY_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "regEx":
            result = REGEX_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "json":
            result = JSON_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "dateTime":
            result = DATETIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "surfaceHub":
            result = SURFACEHUB_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationStringFormat value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationStringFormat(values []DeviceManagementConfigurationStringFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
