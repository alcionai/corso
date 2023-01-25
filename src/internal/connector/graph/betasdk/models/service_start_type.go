package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ServiceStartType int

const (
    // Manual service start type(default)
    MANUAL_SERVICESTARTTYPE ServiceStartType = iota
    // Automatic service start type
    AUTOMATIC_SERVICESTARTTYPE
    // Service start type disabled
    DISABLED_SERVICESTARTTYPE
)

func (i ServiceStartType) String() string {
    return []string{"manual", "automatic", "disabled"}[i]
}
func ParseServiceStartType(v string) (interface{}, error) {
    result := MANUAL_SERVICESTARTTYPE
    switch v {
        case "manual":
            result = MANUAL_SERVICESTARTTYPE
        case "automatic":
            result = AUTOMATIC_SERVICESTARTTYPE
        case "disabled":
            result = DISABLED_SERVICESTARTTYPE
        default:
            return 0, errors.New("Unknown ServiceStartType value: " + v)
    }
    return &result, nil
}
func SerializeServiceStartType(values []ServiceStartType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
