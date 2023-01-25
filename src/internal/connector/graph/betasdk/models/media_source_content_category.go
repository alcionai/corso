package models
import (
    "errors"
)
// Provides operations to call the delta method.
type MediaSourceContentCategory int

const (
    MEETING_MEDIASOURCECONTENTCATEGORY MediaSourceContentCategory = iota
    LIVESTREAM_MEDIASOURCECONTENTCATEGORY
    PRESENTATION_MEDIASOURCECONTENTCATEGORY
    SCREENRECORDING_MEDIASOURCECONTENTCATEGORY
    UNKNOWNFUTUREVALUE_MEDIASOURCECONTENTCATEGORY
)

func (i MediaSourceContentCategory) String() string {
    return []string{"meeting", "liveStream", "presentation", "screenRecording", "unknownFutureValue"}[i]
}
func ParseMediaSourceContentCategory(v string) (interface{}, error) {
    result := MEETING_MEDIASOURCECONTENTCATEGORY
    switch v {
        case "meeting":
            result = MEETING_MEDIASOURCECONTENTCATEGORY
        case "liveStream":
            result = LIVESTREAM_MEDIASOURCECONTENTCATEGORY
        case "presentation":
            result = PRESENTATION_MEDIASOURCECONTENTCATEGORY
        case "screenRecording":
            result = SCREENRECORDING_MEDIASOURCECONTENTCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MEDIASOURCECONTENTCATEGORY
        default:
            return 0, errors.New("Unknown MediaSourceContentCategory value: " + v)
    }
    return &result, nil
}
func SerializeMediaSourceContentCategory(values []MediaSourceContentCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
