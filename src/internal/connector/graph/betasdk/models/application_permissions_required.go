package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ApplicationPermissionsRequired int

const (
    UNKNOWN_APPLICATIONPERMISSIONSREQUIRED ApplicationPermissionsRequired = iota
    ANONYMOUS_APPLICATIONPERMISSIONSREQUIRED
    GUEST_APPLICATIONPERMISSIONSREQUIRED
    USER_APPLICATIONPERMISSIONSREQUIRED
    ADMINISTRATOR_APPLICATIONPERMISSIONSREQUIRED
    SYSTEM_APPLICATIONPERMISSIONSREQUIRED
    UNKNOWNFUTUREVALUE_APPLICATIONPERMISSIONSREQUIRED
)

func (i ApplicationPermissionsRequired) String() string {
    return []string{"unknown", "anonymous", "guest", "user", "administrator", "system", "unknownFutureValue"}[i]
}
func ParseApplicationPermissionsRequired(v string) (interface{}, error) {
    result := UNKNOWN_APPLICATIONPERMISSIONSREQUIRED
    switch v {
        case "unknown":
            result = UNKNOWN_APPLICATIONPERMISSIONSREQUIRED
        case "anonymous":
            result = ANONYMOUS_APPLICATIONPERMISSIONSREQUIRED
        case "guest":
            result = GUEST_APPLICATIONPERMISSIONSREQUIRED
        case "user":
            result = USER_APPLICATIONPERMISSIONSREQUIRED
        case "administrator":
            result = ADMINISTRATOR_APPLICATIONPERMISSIONSREQUIRED
        case "system":
            result = SYSTEM_APPLICATIONPERMISSIONSREQUIRED
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPLICATIONPERMISSIONSREQUIRED
        default:
            return 0, errors.New("Unknown ApplicationPermissionsRequired value: " + v)
    }
    return &result, nil
}
func SerializeApplicationPermissionsRequired(values []ApplicationPermissionsRequired) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
