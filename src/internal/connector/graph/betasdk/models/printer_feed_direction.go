package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type PrinterFeedDirection int

const (
    LONGEDGEFIRST_PRINTERFEEDDIRECTION PrinterFeedDirection = iota
    SHORTEDGEFIRST_PRINTERFEEDDIRECTION
)

func (i PrinterFeedDirection) String() string {
    return []string{"longEdgeFirst", "shortEdgeFirst"}[i]
}
func ParsePrinterFeedDirection(v string) (interface{}, error) {
    result := LONGEDGEFIRST_PRINTERFEEDDIRECTION
    switch v {
        case "longEdgeFirst":
            result = LONGEDGEFIRST_PRINTERFEEDDIRECTION
        case "shortEdgeFirst":
            result = SHORTEDGEFIRST_PRINTERFEEDDIRECTION
        default:
            return 0, errors.New("Unknown PrinterFeedDirection value: " + v)
    }
    return &result, nil
}
func SerializePrinterFeedDirection(values []PrinterFeedDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
