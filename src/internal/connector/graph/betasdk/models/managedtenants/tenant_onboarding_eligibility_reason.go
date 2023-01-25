package managedtenants
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type TenantOnboardingEligibilityReason int

const (
    NONE_TENANTONBOARDINGELIGIBILITYREASON TenantOnboardingEligibilityReason = iota
    CONTRACTTYPE_TENANTONBOARDINGELIGIBILITYREASON
    DELEGATEDADMINPRIVILEGES_TENANTONBOARDINGELIGIBILITYREASON
    USERSCOUNT_TENANTONBOARDINGELIGIBILITYREASON
    LICENSE_TENANTONBOARDINGELIGIBILITYREASON
    UNKNOWNFUTUREVALUE_TENANTONBOARDINGELIGIBILITYREASON
)

func (i TenantOnboardingEligibilityReason) String() string {
    return []string{"none", "contractType", "delegatedAdminPrivileges", "usersCount", "license", "unknownFutureValue"}[i]
}
func ParseTenantOnboardingEligibilityReason(v string) (interface{}, error) {
    result := NONE_TENANTONBOARDINGELIGIBILITYREASON
    switch v {
        case "none":
            result = NONE_TENANTONBOARDINGELIGIBILITYREASON
        case "contractType":
            result = CONTRACTTYPE_TENANTONBOARDINGELIGIBILITYREASON
        case "delegatedAdminPrivileges":
            result = DELEGATEDADMINPRIVILEGES_TENANTONBOARDINGELIGIBILITYREASON
        case "usersCount":
            result = USERSCOUNT_TENANTONBOARDINGELIGIBILITYREASON
        case "license":
            result = LICENSE_TENANTONBOARDINGELIGIBILITYREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TENANTONBOARDINGELIGIBILITYREASON
        default:
            return 0, errors.New("Unknown TenantOnboardingEligibilityReason value: " + v)
    }
    return &result, nil
}
func SerializeTenantOnboardingEligibilityReason(values []TenantOnboardingEligibilityReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
