package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PrintColorConfiguration int

const (
    BLACKANDWHITE_PRINTCOLORCONFIGURATION PrintColorConfiguration = iota
    GRAYSCALE_PRINTCOLORCONFIGURATION
    COLOR_PRINTCOLORCONFIGURATION
    AUTO_PRINTCOLORCONFIGURATION
)

func (i PrintColorConfiguration) String() string {
    return []string{"blackAndWhite", "grayscale", "color", "auto"}[i]
}
func ParsePrintColorConfiguration(v string) (interface{}, error) {
    result := BLACKANDWHITE_PRINTCOLORCONFIGURATION
    switch v {
        case "blackAndWhite":
            result = BLACKANDWHITE_PRINTCOLORCONFIGURATION
        case "grayscale":
            result = GRAYSCALE_PRINTCOLORCONFIGURATION
        case "color":
            result = COLOR_PRINTCOLORCONFIGURATION
        case "auto":
            result = AUTO_PRINTCOLORCONFIGURATION
        default:
            return 0, errors.New("Unknown PrintColorConfiguration value: " + v)
    }
    return &result, nil
}
func SerializePrintColorConfiguration(values []PrintColorConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
