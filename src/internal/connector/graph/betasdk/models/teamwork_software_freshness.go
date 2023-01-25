package models
import (
    "errors"
)
// Provides operations to call the add method.
type TeamworkSoftwareFreshness int

const (
    UNKNOWN_TEAMWORKSOFTWAREFRESHNESS TeamworkSoftwareFreshness = iota
    LATEST_TEAMWORKSOFTWAREFRESHNESS
    UPDATEAVAILABLE_TEAMWORKSOFTWAREFRESHNESS
    UNKNOWNFUTUREVALUE_TEAMWORKSOFTWAREFRESHNESS
)

func (i TeamworkSoftwareFreshness) String() string {
    return []string{"unknown", "latest", "updateAvailable", "unknownFutureValue"}[i]
}
func ParseTeamworkSoftwareFreshness(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKSOFTWAREFRESHNESS
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMWORKSOFTWAREFRESHNESS
        case "latest":
            result = LATEST_TEAMWORKSOFTWAREFRESHNESS
        case "updateAvailable":
            result = UPDATEAVAILABLE_TEAMWORKSOFTWAREFRESHNESS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKSOFTWAREFRESHNESS
        default:
            return 0, errors.New("Unknown TeamworkSoftwareFreshness value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkSoftwareFreshness(values []TeamworkSoftwareFreshness) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
