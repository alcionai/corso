package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AndroidManagedStoreAppConfigurationSchemaItemDataType int

const (
    BOOL_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE AndroidManagedStoreAppConfigurationSchemaItemDataType = iota
    INTEGER_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    STRING_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    CHOICE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    MULTISELECT_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLEARRAY_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    HIDDEN_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
)

func (i AndroidManagedStoreAppConfigurationSchemaItemDataType) String() string {
    return []string{"bool", "integer", "string", "choice", "multiselect", "bundle", "bundleArray", "hidden"}[i]
}
func ParseAndroidManagedStoreAppConfigurationSchemaItemDataType(v string) (interface{}, error) {
    result := BOOL_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    switch v {
        case "bool":
            result = BOOL_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "integer":
            result = INTEGER_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "string":
            result = STRING_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "choice":
            result = CHOICE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "multiselect":
            result = MULTISELECT_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "bundle":
            result = BUNDLE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "bundleArray":
            result = BUNDLEARRAY_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "hidden":
            result = HIDDEN_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAppConfigurationSchemaItemDataType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAppConfigurationSchemaItemDataType(values []AndroidManagedStoreAppConfigurationSchemaItemDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
