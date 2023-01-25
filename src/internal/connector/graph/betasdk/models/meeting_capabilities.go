package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MeetingCapabilities int

const (
    QUESTIONANDANSWER_MEETINGCAPABILITIES MeetingCapabilities = iota
    UNKNOWNFUTUREVALUE_MEETINGCAPABILITIES
)

func (i MeetingCapabilities) String() string {
    return []string{"questionAndAnswer", "unknownFutureValue"}[i]
}
func ParseMeetingCapabilities(v string) (interface{}, error) {
    result := QUESTIONANDANSWER_MEETINGCAPABILITIES
    switch v {
        case "questionAndAnswer":
            result = QUESTIONANDANSWER_MEETINGCAPABILITIES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MEETINGCAPABILITIES
        default:
            return 0, errors.New("Unknown MeetingCapabilities value: " + v)
    }
    return &result, nil
}
func SerializeMeetingCapabilities(values []MeetingCapabilities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
