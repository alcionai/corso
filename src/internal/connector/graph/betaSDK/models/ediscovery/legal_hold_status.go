package ediscovery
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type LegalHoldStatus int

const (
    PENDING_LEGALHOLDSTATUS LegalHoldStatus = iota
    ERROR_LEGALHOLDSTATUS
    SUCCESS_LEGALHOLDSTATUS
    UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS
)

func (i LegalHoldStatus) String() string {
    return []string{"Pending", "Error", "Success", "UnknownFutureValue"}[i]
}
func ParseLegalHoldStatus(v string) (interface{}, error) {
    result := PENDING_LEGALHOLDSTATUS
    switch v {
        case "Pending":
            result = PENDING_LEGALHOLDSTATUS
        case "Error":
            result = ERROR_LEGALHOLDSTATUS
        case "Success":
            result = SUCCESS_LEGALHOLDSTATUS
        case "UnknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS
        default:
            return 0, errors.New("Unknown LegalHoldStatus value: " + v)
    }
    return &result, nil
}
func SerializeLegalHoldStatus(values []LegalHoldStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
