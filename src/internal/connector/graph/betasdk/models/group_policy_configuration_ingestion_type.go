package models
import (
    "errors"
)
// Provides operations to call the add method.
type GroupPolicyConfigurationIngestionType int

const (
    // Unknown policy configuration ingestion type
    UNKNOWN_GROUPPOLICYCONFIGURATIONINGESTIONTYPE GroupPolicyConfigurationIngestionType = iota
    // Indicates policy created have definitions ingested by IT admin with sufficient permissions through custom ingestion process
    CUSTOM_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
    // Indicates policy created have definitions ingested through system ingestion process
    BUILTIN_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
    // Indicated atleast 1 tenant admin & system ingested definitions configured for this policy
    MIXED_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
    // Unknown future enum value
    UNKNOWNFUTUREVALUE_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
)

func (i GroupPolicyConfigurationIngestionType) String() string {
    return []string{"unknown", "custom", "builtIn", "mixed", "unknownFutureValue"}[i]
}
func ParseGroupPolicyConfigurationIngestionType(v string) (interface{}, error) {
    result := UNKNOWN_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
        case "custom":
            result = CUSTOM_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
        case "builtIn":
            result = BUILTIN_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
        case "mixed":
            result = MIXED_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_GROUPPOLICYCONFIGURATIONINGESTIONTYPE
        default:
            return 0, errors.New("Unknown GroupPolicyConfigurationIngestionType value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyConfigurationIngestionType(values []GroupPolicyConfigurationIngestionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
