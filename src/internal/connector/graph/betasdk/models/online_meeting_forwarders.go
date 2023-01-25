package models
import (
    "errors"
)
// Provides operations to call the remove method.
type OnlineMeetingForwarders int

const (
    EVERYONE_ONLINEMEETINGFORWARDERS OnlineMeetingForwarders = iota
    ORGANIZER_ONLINEMEETINGFORWARDERS
    UNKNOWNFUTUREVALUE_ONLINEMEETINGFORWARDERS
)

func (i OnlineMeetingForwarders) String() string {
    return []string{"everyone", "organizer", "unknownFutureValue"}[i]
}
func ParseOnlineMeetingForwarders(v string) (interface{}, error) {
    result := EVERYONE_ONLINEMEETINGFORWARDERS
    switch v {
        case "everyone":
            result = EVERYONE_ONLINEMEETINGFORWARDERS
        case "organizer":
            result = ORGANIZER_ONLINEMEETINGFORWARDERS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONLINEMEETINGFORWARDERS
        default:
            return 0, errors.New("Unknown OnlineMeetingForwarders value: " + v)
    }
    return &result, nil
}
func SerializeOnlineMeetingForwarders(values []OnlineMeetingForwarders) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
