package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type OwnerType int

const (
    // Unknown.
    UNKNOWN_OWNERTYPE OwnerType = iota
    // Owned by company.
    COMPANY_OWNERTYPE
    // Owned by person.
    PERSONAL_OWNERTYPE
)

func (i OwnerType) String() string {
    return []string{"unknown", "company", "personal"}[i]
}
func ParseOwnerType(v string) (interface{}, error) {
    result := UNKNOWN_OWNERTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_OWNERTYPE
        case "company":
            result = COMPANY_OWNERTYPE
        case "personal":
            result = PERSONAL_OWNERTYPE
        default:
            return 0, errors.New("Unknown OwnerType value: " + v)
    }
    return &result, nil
}
func SerializeOwnerType(values []OwnerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
