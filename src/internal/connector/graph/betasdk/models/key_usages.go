package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type KeyUsages int

const (
    // Key Encipherment Usage.
    KEYENCIPHERMENT_KEYUSAGES KeyUsages = iota
    // Digital Signature Usage.
    DIGITALSIGNATURE_KEYUSAGES
)

func (i KeyUsages) String() string {
    return []string{"keyEncipherment", "digitalSignature"}[i]
}
func ParseKeyUsages(v string) (interface{}, error) {
    result := KEYENCIPHERMENT_KEYUSAGES
    switch v {
        case "keyEncipherment":
            result = KEYENCIPHERMENT_KEYUSAGES
        case "digitalSignature":
            result = DIGITALSIGNATURE_KEYUSAGES
        default:
            return 0, errors.New("Unknown KeyUsages value: " + v)
    }
    return &result, nil
}
func SerializeKeyUsages(values []KeyUsages) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
