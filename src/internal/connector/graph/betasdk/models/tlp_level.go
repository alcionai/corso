package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TlpLevel int

const (
    UNKNOWN_TLPLEVEL TlpLevel = iota
    WHITE_TLPLEVEL
    GREEN_TLPLEVEL
    AMBER_TLPLEVEL
    RED_TLPLEVEL
    UNKNOWNFUTUREVALUE_TLPLEVEL
)

func (i TlpLevel) String() string {
    return []string{"unknown", "white", "green", "amber", "red", "unknownFutureValue"}[i]
}
func ParseTlpLevel(v string) (interface{}, error) {
    result := UNKNOWN_TLPLEVEL
    switch v {
        case "unknown":
            result = UNKNOWN_TLPLEVEL
        case "white":
            result = WHITE_TLPLEVEL
        case "green":
            result = GREEN_TLPLEVEL
        case "amber":
            result = AMBER_TLPLEVEL
        case "red":
            result = RED_TLPLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TLPLEVEL
        default:
            return 0, errors.New("Unknown TlpLevel value: " + v)
    }
    return &result, nil
}
func SerializeTlpLevel(values []TlpLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
