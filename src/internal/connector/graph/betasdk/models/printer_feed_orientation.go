package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PrinterFeedOrientation int

const (
    LONGEDGEFIRST_PRINTERFEEDORIENTATION PrinterFeedOrientation = iota
    SHORTEDGEFIRST_PRINTERFEEDORIENTATION
)

func (i PrinterFeedOrientation) String() string {
    return []string{"longEdgeFirst", "shortEdgeFirst"}[i]
}
func ParsePrinterFeedOrientation(v string) (interface{}, error) {
    result := LONGEDGEFIRST_PRINTERFEEDORIENTATION
    switch v {
        case "longEdgeFirst":
            result = LONGEDGEFIRST_PRINTERFEEDORIENTATION
        case "shortEdgeFirst":
            result = SHORTEDGEFIRST_PRINTERFEEDORIENTATION
        default:
            return 0, errors.New("Unknown PrinterFeedOrientation value: " + v)
    }
    return &result, nil
}
func SerializePrinterFeedOrientation(values []PrinterFeedOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
