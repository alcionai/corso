package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CompanyPortalAction int

const (
    // Unknown device action
    UNKNOWN_COMPANYPORTALACTION CompanyPortalAction = iota
    // Remove device from Company Portal
    REMOVE_COMPANYPORTALACTION
    // Reset device enrolled in Company Portal
    RESET_COMPANYPORTALACTION
)

func (i CompanyPortalAction) String() string {
    return []string{"unknown", "remove", "reset"}[i]
}
func ParseCompanyPortalAction(v string) (interface{}, error) {
    result := UNKNOWN_COMPANYPORTALACTION
    switch v {
        case "unknown":
            result = UNKNOWN_COMPANYPORTALACTION
        case "remove":
            result = REMOVE_COMPANYPORTALACTION
        case "reset":
            result = RESET_COMPANYPORTALACTION
        default:
            return 0, errors.New("Unknown CompanyPortalAction value: " + v)
    }
    return &result, nil
}
func SerializeCompanyPortalAction(values []CompanyPortalAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
