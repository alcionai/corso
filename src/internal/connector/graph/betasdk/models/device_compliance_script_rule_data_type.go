package models
import (
    "errors"
)
// Provides operations to manage the columns property of the microsoft.graph.site entity.
type DeviceComplianceScriptRuleDataType int

const (
    // None data type.
    NONE_DEVICECOMPLIANCESCRIPTRULEDATATYPE DeviceComplianceScriptRuleDataType = iota
    // Boolean data type.
    BOOLEAN_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Int64 data type.
    INT64_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Double data type.
    DOUBLE_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // String data type.
    STRING_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // DateTime data type.
    DATETIME_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Version data type.
    VERSION_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Base64 data type.
    BASE64_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Xml data type.
    XML_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Boolean array data type.
    BOOLEANARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Int64 array data type.
    INT64ARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Double array data type.
    DOUBLEARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // String array data type.
    STRINGARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // DateTime array data type.
    DATETIMEARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    // Version array data type.
    VERSIONARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
)

func (i DeviceComplianceScriptRuleDataType) String() string {
    return []string{"none", "boolean", "int64", "double", "string", "dateTime", "version", "base64", "xml", "booleanArray", "int64Array", "doubleArray", "stringArray", "dateTimeArray", "versionArray"}[i]
}
func ParseDeviceComplianceScriptRuleDataType(v string) (interface{}, error) {
    result := NONE_DEVICECOMPLIANCESCRIPTRULEDATATYPE
    switch v {
        case "none":
            result = NONE_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "boolean":
            result = BOOLEAN_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "int64":
            result = INT64_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "double":
            result = DOUBLE_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "string":
            result = STRING_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "dateTime":
            result = DATETIME_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "version":
            result = VERSION_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "base64":
            result = BASE64_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "xml":
            result = XML_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "booleanArray":
            result = BOOLEANARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "int64Array":
            result = INT64ARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "doubleArray":
            result = DOUBLEARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "stringArray":
            result = STRINGARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "dateTimeArray":
            result = DATETIMEARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        case "versionArray":
            result = VERSIONARRAY_DEVICECOMPLIANCESCRIPTRULEDATATYPE
        default:
            return 0, errors.New("Unknown DeviceComplianceScriptRuleDataType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceComplianceScriptRuleDataType(values []DeviceComplianceScriptRuleDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
