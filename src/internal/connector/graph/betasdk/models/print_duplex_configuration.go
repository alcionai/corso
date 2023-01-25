package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type PrintDuplexConfiguration int

const (
    TWOSIDEDLONGEDGE_PRINTDUPLEXCONFIGURATION PrintDuplexConfiguration = iota
    TWOSIDEDSHORTEDGE_PRINTDUPLEXCONFIGURATION
    ONESIDED_PRINTDUPLEXCONFIGURATION
)

func (i PrintDuplexConfiguration) String() string {
    return []string{"twoSidedLongEdge", "twoSidedShortEdge", "oneSided"}[i]
}
func ParsePrintDuplexConfiguration(v string) (interface{}, error) {
    result := TWOSIDEDLONGEDGE_PRINTDUPLEXCONFIGURATION
    switch v {
        case "twoSidedLongEdge":
            result = TWOSIDEDLONGEDGE_PRINTDUPLEXCONFIGURATION
        case "twoSidedShortEdge":
            result = TWOSIDEDSHORTEDGE_PRINTDUPLEXCONFIGURATION
        case "oneSided":
            result = ONESIDED_PRINTDUPLEXCONFIGURATION
        default:
            return 0, errors.New("Unknown PrintDuplexConfiguration value: " + v)
    }
    return &result, nil
}
func SerializePrintDuplexConfiguration(values []PrintDuplexConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
