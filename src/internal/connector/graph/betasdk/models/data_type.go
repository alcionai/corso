package models
import (
    "errors"
)
// Provides operations to manage the columns property of the microsoft.graph.site entity.
type DataType int

const (
    // None data type.
    NONE_DATATYPE DataType = iota
    // Boolean data type.
    BOOLEAN_DATATYPE
    // Int64 data type.
    INT64_DATATYPE
    // Double data type.
    DOUBLE_DATATYPE
    // String data type.
    STRING_DATATYPE
    // DateTime data type.
    DATETIME_DATATYPE
    // Version data type.
    VERSION_DATATYPE
    // Base64 data type.
    BASE64_DATATYPE
    // Xml data type.
    XML_DATATYPE
    // Boolean array data type.
    BOOLEANARRAY_DATATYPE
    // Int64 array data type.
    INT64ARRAY_DATATYPE
    // Double array data type.
    DOUBLEARRAY_DATATYPE
    // String array data type.
    STRINGARRAY_DATATYPE
    // DateTime array data type.
    DATETIMEARRAY_DATATYPE
    // Version array data type.
    VERSIONARRAY_DATATYPE
)

func (i DataType) String() string {
    return []string{"none", "boolean", "int64", "double", "string", "dateTime", "version", "base64", "xml", "booleanArray", "int64Array", "doubleArray", "stringArray", "dateTimeArray", "versionArray"}[i]
}
func ParseDataType(v string) (interface{}, error) {
    result := NONE_DATATYPE
    switch v {
        case "none":
            result = NONE_DATATYPE
        case "boolean":
            result = BOOLEAN_DATATYPE
        case "int64":
            result = INT64_DATATYPE
        case "double":
            result = DOUBLE_DATATYPE
        case "string":
            result = STRING_DATATYPE
        case "dateTime":
            result = DATETIME_DATATYPE
        case "version":
            result = VERSION_DATATYPE
        case "base64":
            result = BASE64_DATATYPE
        case "xml":
            result = XML_DATATYPE
        case "booleanArray":
            result = BOOLEANARRAY_DATATYPE
        case "int64Array":
            result = INT64ARRAY_DATATYPE
        case "doubleArray":
            result = DOUBLEARRAY_DATATYPE
        case "stringArray":
            result = STRINGARRAY_DATATYPE
        case "dateTimeArray":
            result = DATETIMEARRAY_DATATYPE
        case "versionArray":
            result = VERSIONARRAY_DATATYPE
        default:
            return 0, errors.New("Unknown DataType value: " + v)
    }
    return &result, nil
}
func SerializeDataType(values []DataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
