package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TiAction int

const (
    UNKNOWN_TIACTION TiAction = iota
    ALLOW_TIACTION
    BLOCK_TIACTION
    ALERT_TIACTION
    UNKNOWNFUTUREVALUE_TIACTION
)

func (i TiAction) String() string {
    return []string{"unknown", "allow", "block", "alert", "unknownFutureValue"}[i]
}
func ParseTiAction(v string) (interface{}, error) {
    result := UNKNOWN_TIACTION
    switch v {
        case "unknown":
            result = UNKNOWN_TIACTION
        case "allow":
            result = ALLOW_TIACTION
        case "block":
            result = BLOCK_TIACTION
        case "alert":
            result = ALERT_TIACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TIACTION
        default:
            return 0, errors.New("Unknown TiAction value: " + v)
    }
    return &result, nil
}
func SerializeTiAction(values []TiAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
