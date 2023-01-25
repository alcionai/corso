package models
import (
    "errors"
)
// Provides operations to call the add method.
type TeamworkDeviceType int

const (
    UNKNOWN_TEAMWORKDEVICETYPE TeamworkDeviceType = iota
    IPPHONE_TEAMWORKDEVICETYPE
    TEAMSROOM_TEAMWORKDEVICETYPE
    SURFACEHUB_TEAMWORKDEVICETYPE
    COLLABORATIONBAR_TEAMWORKDEVICETYPE
    TEAMSDISPLAY_TEAMWORKDEVICETYPE
    TOUCHCONSOLE_TEAMWORKDEVICETYPE
    LOWCOSTPHONE_TEAMWORKDEVICETYPE
    TEAMSPANEL_TEAMWORKDEVICETYPE
    SIP_TEAMWORKDEVICETYPE
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICETYPE
)

func (i TeamworkDeviceType) String() string {
    return []string{"unknown", "ipPhone", "teamsRoom", "surfaceHub", "collaborationBar", "teamsDisplay", "touchConsole", "lowCostPhone", "teamsPanel", "sip", "unknownFutureValue"}[i]
}
func ParseTeamworkDeviceType(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKDEVICETYPE
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMWORKDEVICETYPE
        case "ipPhone":
            result = IPPHONE_TEAMWORKDEVICETYPE
        case "teamsRoom":
            result = TEAMSROOM_TEAMWORKDEVICETYPE
        case "surfaceHub":
            result = SURFACEHUB_TEAMWORKDEVICETYPE
        case "collaborationBar":
            result = COLLABORATIONBAR_TEAMWORKDEVICETYPE
        case "teamsDisplay":
            result = TEAMSDISPLAY_TEAMWORKDEVICETYPE
        case "touchConsole":
            result = TOUCHCONSOLE_TEAMWORKDEVICETYPE
        case "lowCostPhone":
            result = LOWCOSTPHONE_TEAMWORKDEVICETYPE
        case "teamsPanel":
            result = TEAMSPANEL_TEAMWORKDEVICETYPE
        case "sip":
            result = SIP_TEAMWORKDEVICETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICETYPE
        default:
            return 0, errors.New("Unknown TeamworkDeviceType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkDeviceType(values []TeamworkDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
