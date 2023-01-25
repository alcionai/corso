package windowsupdates
import (
    "errors"
)
// Provides operations to call the add method.
type ResourceConnectionState int

const (
    CONNECTED_RESOURCECONNECTIONSTATE ResourceConnectionState = iota
    NOTAUTHORIZED_RESOURCECONNECTIONSTATE
    NOTFOUND_RESOURCECONNECTIONSTATE
    UNKNOWNFUTUREVALUE_RESOURCECONNECTIONSTATE
)

func (i ResourceConnectionState) String() string {
    return []string{"connected", "notAuthorized", "notFound", "unknownFutureValue"}[i]
}
func ParseResourceConnectionState(v string) (interface{}, error) {
    result := CONNECTED_RESOURCECONNECTIONSTATE
    switch v {
        case "connected":
            result = CONNECTED_RESOURCECONNECTIONSTATE
        case "notAuthorized":
            result = NOTAUTHORIZED_RESOURCECONNECTIONSTATE
        case "notFound":
            result = NOTFOUND_RESOURCECONNECTIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RESOURCECONNECTIONSTATE
        default:
            return 0, errors.New("Unknown ResourceConnectionState value: " + v)
    }
    return &result, nil
}
func SerializeResourceConnectionState(values []ResourceConnectionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
