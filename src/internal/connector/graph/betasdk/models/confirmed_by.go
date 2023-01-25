package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ConfirmedBy int

const (
    NONE_CONFIRMEDBY ConfirmedBy = iota
    USER_CONFIRMEDBY
    MANAGER_CONFIRMEDBY
    UNKNOWNFUTUREVALUE_CONFIRMEDBY
)

func (i ConfirmedBy) String() string {
    return []string{"none", "user", "manager", "unknownFutureValue"}[i]
}
func ParseConfirmedBy(v string) (interface{}, error) {
    result := NONE_CONFIRMEDBY
    switch v {
        case "none":
            result = NONE_CONFIRMEDBY
        case "user":
            result = USER_CONFIRMEDBY
        case "manager":
            result = MANAGER_CONFIRMEDBY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONFIRMEDBY
        default:
            return 0, errors.New("Unknown ConfirmedBy value: " + v)
    }
    return &result, nil
}
func SerializeConfirmedBy(values []ConfirmedBy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
