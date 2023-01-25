package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AndroidForWorkAppConfigurationSchemaItemDataType int

const (
    BOOL_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE AndroidForWorkAppConfigurationSchemaItemDataType = iota
    INTEGER_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    STRING_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    CHOICE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    MULTISELECT_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLEARRAY_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    HIDDEN_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
)

func (i AndroidForWorkAppConfigurationSchemaItemDataType) String() string {
    return []string{"bool", "integer", "string", "choice", "multiselect", "bundle", "bundleArray", "hidden"}[i]
}
func ParseAndroidForWorkAppConfigurationSchemaItemDataType(v string) (interface{}, error) {
    result := BOOL_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    switch v {
        case "bool":
            result = BOOL_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "integer":
            result = INTEGER_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "string":
            result = STRING_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "choice":
            result = CHOICE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "multiselect":
            result = MULTISELECT_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "bundle":
            result = BUNDLE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "bundleArray":
            result = BUNDLEARRAY_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "hidden":
            result = HIDDEN_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        default:
            return 0, errors.New("Unknown AndroidForWorkAppConfigurationSchemaItemDataType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkAppConfigurationSchemaItemDataType(values []AndroidForWorkAppConfigurationSchemaItemDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
