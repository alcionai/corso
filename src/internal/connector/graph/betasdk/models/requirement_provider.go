package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RequirementProvider int

const (
    USER_REQUIREMENTPROVIDER RequirementProvider = iota
    REQUEST_REQUIREMENTPROVIDER
    SERVICEPRINCIPAL_REQUIREMENTPROVIDER
    V1CONDITIONALACCESS_REQUIREMENTPROVIDER
    MULTICONDITIONALACCESS_REQUIREMENTPROVIDER
    TENANTSESSIONRISKPOLICY_REQUIREMENTPROVIDER
    ACCOUNTCOMPROMISEPOLICIES_REQUIREMENTPROVIDER
    V1CONDITIONALACCESSDEPENDENCY_REQUIREMENTPROVIDER
    V1CONDITIONALACCESSPOLICYIDREQUESTED_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYIDENTITYPROTECTIONPOLICY_REQUIREMENTPROVIDER
    BASELINEPROTECTION_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYBASELINEPROTECTION_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYMULTICONDITIONALACCESS_REQUIREMENTPROVIDER
    ENFORCEDFORCSPADMINS_REQUIREMENTPROVIDER
    SECURITYDEFAULTS_REQUIREMENTPROVIDER
    MFAREGISTRATIONREQUIREDBYSECURITYDEFAULTS_REQUIREMENTPROVIDER
    PROOFUPCODEREQUEST_REQUIREMENTPROVIDER
    CROSSTENANTOUTBOUNDRULE_REQUIREMENTPROVIDER
    GPSLOCATIONCONDITION_REQUIREMENTPROVIDER
    RISKBASEDPOLICY_REQUIREMENTPROVIDER
    UNKNOWNFUTUREVALUE_REQUIREMENTPROVIDER
)

func (i RequirementProvider) String() string {
    return []string{"user", "request", "servicePrincipal", "v1ConditionalAccess", "multiConditionalAccess", "tenantSessionRiskPolicy", "accountCompromisePolicies", "v1ConditionalAccessDependency", "v1ConditionalAccessPolicyIdRequested", "mfaRegistrationRequiredByIdentityProtectionPolicy", "baselineProtection", "mfaRegistrationRequiredByBaselineProtection", "mfaRegistrationRequiredByMultiConditionalAccess", "enforcedForCspAdmins", "securityDefaults", "mfaRegistrationRequiredBySecurityDefaults", "proofUpCodeRequest", "crossTenantOutboundRule", "gpsLocationCondition", "riskBasedPolicy", "unknownFutureValue"}[i]
}
func ParseRequirementProvider(v string) (interface{}, error) {
    result := USER_REQUIREMENTPROVIDER
    switch v {
        case "user":
            result = USER_REQUIREMENTPROVIDER
        case "request":
            result = REQUEST_REQUIREMENTPROVIDER
        case "servicePrincipal":
            result = SERVICEPRINCIPAL_REQUIREMENTPROVIDER
        case "v1ConditionalAccess":
            result = V1CONDITIONALACCESS_REQUIREMENTPROVIDER
        case "multiConditionalAccess":
            result = MULTICONDITIONALACCESS_REQUIREMENTPROVIDER
        case "tenantSessionRiskPolicy":
            result = TENANTSESSIONRISKPOLICY_REQUIREMENTPROVIDER
        case "accountCompromisePolicies":
            result = ACCOUNTCOMPROMISEPOLICIES_REQUIREMENTPROVIDER
        case "v1ConditionalAccessDependency":
            result = V1CONDITIONALACCESSDEPENDENCY_REQUIREMENTPROVIDER
        case "v1ConditionalAccessPolicyIdRequested":
            result = V1CONDITIONALACCESSPOLICYIDREQUESTED_REQUIREMENTPROVIDER
        case "mfaRegistrationRequiredByIdentityProtectionPolicy":
            result = MFAREGISTRATIONREQUIREDBYIDENTITYPROTECTIONPOLICY_REQUIREMENTPROVIDER
        case "baselineProtection":
            result = BASELINEPROTECTION_REQUIREMENTPROVIDER
        case "mfaRegistrationRequiredByBaselineProtection":
            result = MFAREGISTRATIONREQUIREDBYBASELINEPROTECTION_REQUIREMENTPROVIDER
        case "mfaRegistrationRequiredByMultiConditionalAccess":
            result = MFAREGISTRATIONREQUIREDBYMULTICONDITIONALACCESS_REQUIREMENTPROVIDER
        case "enforcedForCspAdmins":
            result = ENFORCEDFORCSPADMINS_REQUIREMENTPROVIDER
        case "securityDefaults":
            result = SECURITYDEFAULTS_REQUIREMENTPROVIDER
        case "mfaRegistrationRequiredBySecurityDefaults":
            result = MFAREGISTRATIONREQUIREDBYSECURITYDEFAULTS_REQUIREMENTPROVIDER
        case "proofUpCodeRequest":
            result = PROOFUPCODEREQUEST_REQUIREMENTPROVIDER
        case "crossTenantOutboundRule":
            result = CROSSTENANTOUTBOUNDRULE_REQUIREMENTPROVIDER
        case "gpsLocationCondition":
            result = GPSLOCATIONCONDITION_REQUIREMENTPROVIDER
        case "riskBasedPolicy":
            result = RISKBASEDPOLICY_REQUIREMENTPROVIDER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REQUIREMENTPROVIDER
        default:
            return 0, errors.New("Unknown RequirementProvider value: " + v)
    }
    return &result, nil
}
func SerializeRequirementProvider(values []RequirementProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
