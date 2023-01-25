package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type PrintMediaType int

const (
    STATIONERY_PRINTMEDIATYPE PrintMediaType = iota
    TRANSPARENCY_PRINTMEDIATYPE
    ENVELOPE_PRINTMEDIATYPE
    ENVELOPEPLAIN_PRINTMEDIATYPE
    CONTINUOUS_PRINTMEDIATYPE
    SCREEN_PRINTMEDIATYPE
    SCREENPAGED_PRINTMEDIATYPE
    CONTINUOUSLONG_PRINTMEDIATYPE
    CONTINUOUSSHORT_PRINTMEDIATYPE
    ENVELOPEWINDOW_PRINTMEDIATYPE
    MULTIPARTFORM_PRINTMEDIATYPE
    MULTILAYER_PRINTMEDIATYPE
    LABELS_PRINTMEDIATYPE
)

func (i PrintMediaType) String() string {
    return []string{"stationery", "transparency", "envelope", "envelopePlain", "continuous", "screen", "screenPaged", "continuousLong", "continuousShort", "envelopeWindow", "multiPartForm", "multiLayer", "labels"}[i]
}
func ParsePrintMediaType(v string) (interface{}, error) {
    result := STATIONERY_PRINTMEDIATYPE
    switch v {
        case "stationery":
            result = STATIONERY_PRINTMEDIATYPE
        case "transparency":
            result = TRANSPARENCY_PRINTMEDIATYPE
        case "envelope":
            result = ENVELOPE_PRINTMEDIATYPE
        case "envelopePlain":
            result = ENVELOPEPLAIN_PRINTMEDIATYPE
        case "continuous":
            result = CONTINUOUS_PRINTMEDIATYPE
        case "screen":
            result = SCREEN_PRINTMEDIATYPE
        case "screenPaged":
            result = SCREENPAGED_PRINTMEDIATYPE
        case "continuousLong":
            result = CONTINUOUSLONG_PRINTMEDIATYPE
        case "continuousShort":
            result = CONTINUOUSSHORT_PRINTMEDIATYPE
        case "envelopeWindow":
            result = ENVELOPEWINDOW_PRINTMEDIATYPE
        case "multiPartForm":
            result = MULTIPARTFORM_PRINTMEDIATYPE
        case "multiLayer":
            result = MULTILAYER_PRINTMEDIATYPE
        case "labels":
            result = LABELS_PRINTMEDIATYPE
        default:
            return 0, errors.New("Unknown PrintMediaType value: " + v)
    }
    return &result, nil
}
func SerializePrintMediaType(values []PrintMediaType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
