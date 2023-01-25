package windowsupdates
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RequestedDeploymentStateValue int

const (
    NONE_REQUESTEDDEPLOYMENTSTATEVALUE RequestedDeploymentStateValue = iota
    PAUSED_REQUESTEDDEPLOYMENTSTATEVALUE
    ARCHIVED_REQUESTEDDEPLOYMENTSTATEVALUE
    UNKNOWNFUTUREVALUE_REQUESTEDDEPLOYMENTSTATEVALUE
)

func (i RequestedDeploymentStateValue) String() string {
    return []string{"none", "paused", "archived", "unknownFutureValue"}[i]
}
func ParseRequestedDeploymentStateValue(v string) (interface{}, error) {
    result := NONE_REQUESTEDDEPLOYMENTSTATEVALUE
    switch v {
        case "none":
            result = NONE_REQUESTEDDEPLOYMENTSTATEVALUE
        case "paused":
            result = PAUSED_REQUESTEDDEPLOYMENTSTATEVALUE
        case "archived":
            result = ARCHIVED_REQUESTEDDEPLOYMENTSTATEVALUE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REQUESTEDDEPLOYMENTSTATEVALUE
        default:
            return 0, errors.New("Unknown RequestedDeploymentStateValue value: " + v)
    }
    return &result, nil
}
func SerializeRequestedDeploymentStateValue(values []RequestedDeploymentStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
