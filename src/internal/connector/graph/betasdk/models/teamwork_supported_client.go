package models
import (
    "errors"
)
// Provides operations to call the add method.
type TeamworkSupportedClient int

const (
    UNKNOWN_TEAMWORKSUPPORTEDCLIENT TeamworkSupportedClient = iota
    SKYPEDEFAULTANDTEAMS_TEAMWORKSUPPORTEDCLIENT
    TEAMSDEFAULTANDSKYPE_TEAMWORKSUPPORTEDCLIENT
    SKYPEONLY_TEAMWORKSUPPORTEDCLIENT
    TEAMSONLY_TEAMWORKSUPPORTEDCLIENT
    UNKNOWNFUTUREVALUE_TEAMWORKSUPPORTEDCLIENT
)

func (i TeamworkSupportedClient) String() string {
    return []string{"unknown", "skypeDefaultAndTeams", "teamsDefaultAndSkype", "skypeOnly", "teamsOnly", "unknownFutureValue"}[i]
}
func ParseTeamworkSupportedClient(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKSUPPORTEDCLIENT
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMWORKSUPPORTEDCLIENT
        case "skypeDefaultAndTeams":
            result = SKYPEDEFAULTANDTEAMS_TEAMWORKSUPPORTEDCLIENT
        case "teamsDefaultAndSkype":
            result = TEAMSDEFAULTANDSKYPE_TEAMWORKSUPPORTEDCLIENT
        case "skypeOnly":
            result = SKYPEONLY_TEAMWORKSUPPORTEDCLIENT
        case "teamsOnly":
            result = TEAMSONLY_TEAMWORKSUPPORTEDCLIENT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKSUPPORTEDCLIENT
        default:
            return 0, errors.New("Unknown TeamworkSupportedClient value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkSupportedClient(values []TeamworkSupportedClient) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
