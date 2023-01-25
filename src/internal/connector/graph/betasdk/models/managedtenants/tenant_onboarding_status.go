package managedtenants
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type TenantOnboardingStatus int

const (
    INELIGIBLE_TENANTONBOARDINGSTATUS TenantOnboardingStatus = iota
    INPROCESS_TENANTONBOARDINGSTATUS
    ACTIVE_TENANTONBOARDINGSTATUS
    INACTIVE_TENANTONBOARDINGSTATUS
    UNKNOWNFUTUREVALUE_TENANTONBOARDINGSTATUS
)

func (i TenantOnboardingStatus) String() string {
    return []string{"ineligible", "inProcess", "active", "inactive", "unknownFutureValue"}[i]
}
func ParseTenantOnboardingStatus(v string) (interface{}, error) {
    result := INELIGIBLE_TENANTONBOARDINGSTATUS
    switch v {
        case "ineligible":
            result = INELIGIBLE_TENANTONBOARDINGSTATUS
        case "inProcess":
            result = INPROCESS_TENANTONBOARDINGSTATUS
        case "active":
            result = ACTIVE_TENANTONBOARDINGSTATUS
        case "inactive":
            result = INACTIVE_TENANTONBOARDINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TENANTONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown TenantOnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeTenantOnboardingStatus(values []TenantOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
