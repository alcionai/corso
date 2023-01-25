package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RecommendationType int

const (
    ADFSAPPSMIGRATION_RECOMMENDATIONTYPE RecommendationType = iota
    ENABLEDESKTOPSSO_RECOMMENDATIONTYPE
    ENABLEPHS_RECOMMENDATIONTYPE
    ENABLEPROVISIONING_RECOMMENDATIONTYPE
    SWITCHFROMPERUSERMFA_RECOMMENDATIONTYPE
    TENANTMFA_RECOMMENDATIONTYPE
    THIRDPARTYAPPS_RECOMMENDATIONTYPE
    TURNOFFPERUSERMFA_RECOMMENDATIONTYPE
    USEAUTHENTICATORAPP_RECOMMENDATIONTYPE
    USEMYAPPS_RECOMMENDATIONTYPE
    STALEAPPS_RECOMMENDATIONTYPE
    STALEAPPCREDS_RECOMMENDATIONTYPE
    APPLICATIONCREDENTIALEXPIRY_RECOMMENDATIONTYPE
    SERVICEPRINCIPALKEYEXPIRY_RECOMMENDATIONTYPE
    ADMINMFAV2_RECOMMENDATIONTYPE
    BLOCKLEGACYAUTHENTICATION_RECOMMENDATIONTYPE
    INTEGRATEDAPPS_RECOMMENDATIONTYPE
    MFAREGISTRATIONV2_RECOMMENDATIONTYPE
    PWAGEPOLICYNEW_RECOMMENDATIONTYPE
    PASSWORDHASHSYNC_RECOMMENDATIONTYPE
    ONEADMIN_RECOMMENDATIONTYPE
    ROLEOVERLAP_RECOMMENDATIONTYPE
    SELFSERVICEPASSWORDRESET_RECOMMENDATIONTYPE
    SIGNINRISKPOLICY_RECOMMENDATIONTYPE
    USERRISKPOLICY_RECOMMENDATIONTYPE
    VERIFYAPPPUBLISHER_RECOMMENDATIONTYPE
    PRIVATELINKFORAAD_RECOMMENDATIONTYPE
    APPROLEASSIGNMENTSGROUPS_RECOMMENDATIONTYPE
    APPROLEASSIGNMENTSUSERS_RECOMMENDATIONTYPE
    MANAGEDIDENTITY_RECOMMENDATIONTYPE
    OVERPRIVILEGEDAPPS_RECOMMENDATIONTYPE
    UNKNOWNFUTUREVALUE_RECOMMENDATIONTYPE
)

func (i RecommendationType) String() string {
    return []string{"adfsAppsMigration", "enableDesktopSSO", "enablePHS", "enableProvisioning", "switchFromPerUserMFA", "tenantMFA", "thirdPartyApps", "turnOffPerUserMFA", "useAuthenticatorApp", "useMyApps", "staleApps", "staleAppCreds", "applicationCredentialExpiry", "servicePrincipalKeyExpiry", "adminMFAV2", "blockLegacyAuthentication", "integratedApps", "mfaRegistrationV2", "pwagePolicyNew", "passwordHashSync", "oneAdmin", "roleOverlap", "selfServicePasswordReset", "signinRiskPolicy", "userRiskPolicy", "verifyAppPublisher", "privateLinkForAAD", "appRoleAssignmentsGroups", "appRoleAssignmentsUsers", "managedIdentity", "overprivilegedApps", "unknownFutureValue"}[i]
}
func ParseRecommendationType(v string) (interface{}, error) {
    result := ADFSAPPSMIGRATION_RECOMMENDATIONTYPE
    switch v {
        case "adfsAppsMigration":
            result = ADFSAPPSMIGRATION_RECOMMENDATIONTYPE
        case "enableDesktopSSO":
            result = ENABLEDESKTOPSSO_RECOMMENDATIONTYPE
        case "enablePHS":
            result = ENABLEPHS_RECOMMENDATIONTYPE
        case "enableProvisioning":
            result = ENABLEPROVISIONING_RECOMMENDATIONTYPE
        case "switchFromPerUserMFA":
            result = SWITCHFROMPERUSERMFA_RECOMMENDATIONTYPE
        case "tenantMFA":
            result = TENANTMFA_RECOMMENDATIONTYPE
        case "thirdPartyApps":
            result = THIRDPARTYAPPS_RECOMMENDATIONTYPE
        case "turnOffPerUserMFA":
            result = TURNOFFPERUSERMFA_RECOMMENDATIONTYPE
        case "useAuthenticatorApp":
            result = USEAUTHENTICATORAPP_RECOMMENDATIONTYPE
        case "useMyApps":
            result = USEMYAPPS_RECOMMENDATIONTYPE
        case "staleApps":
            result = STALEAPPS_RECOMMENDATIONTYPE
        case "staleAppCreds":
            result = STALEAPPCREDS_RECOMMENDATIONTYPE
        case "applicationCredentialExpiry":
            result = APPLICATIONCREDENTIALEXPIRY_RECOMMENDATIONTYPE
        case "servicePrincipalKeyExpiry":
            result = SERVICEPRINCIPALKEYEXPIRY_RECOMMENDATIONTYPE
        case "adminMFAV2":
            result = ADMINMFAV2_RECOMMENDATIONTYPE
        case "blockLegacyAuthentication":
            result = BLOCKLEGACYAUTHENTICATION_RECOMMENDATIONTYPE
        case "integratedApps":
            result = INTEGRATEDAPPS_RECOMMENDATIONTYPE
        case "mfaRegistrationV2":
            result = MFAREGISTRATIONV2_RECOMMENDATIONTYPE
        case "pwagePolicyNew":
            result = PWAGEPOLICYNEW_RECOMMENDATIONTYPE
        case "passwordHashSync":
            result = PASSWORDHASHSYNC_RECOMMENDATIONTYPE
        case "oneAdmin":
            result = ONEADMIN_RECOMMENDATIONTYPE
        case "roleOverlap":
            result = ROLEOVERLAP_RECOMMENDATIONTYPE
        case "selfServicePasswordReset":
            result = SELFSERVICEPASSWORDRESET_RECOMMENDATIONTYPE
        case "signinRiskPolicy":
            result = SIGNINRISKPOLICY_RECOMMENDATIONTYPE
        case "userRiskPolicy":
            result = USERRISKPOLICY_RECOMMENDATIONTYPE
        case "verifyAppPublisher":
            result = VERIFYAPPPUBLISHER_RECOMMENDATIONTYPE
        case "privateLinkForAAD":
            result = PRIVATELINKFORAAD_RECOMMENDATIONTYPE
        case "appRoleAssignmentsGroups":
            result = APPROLEASSIGNMENTSGROUPS_RECOMMENDATIONTYPE
        case "appRoleAssignmentsUsers":
            result = APPROLEASSIGNMENTSUSERS_RECOMMENDATIONTYPE
        case "managedIdentity":
            result = MANAGEDIDENTITY_RECOMMENDATIONTYPE
        case "overprivilegedApps":
            result = OVERPRIVILEGEDAPPS_RECOMMENDATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RECOMMENDATIONTYPE
        default:
            return 0, errors.New("Unknown RecommendationType value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationType(values []RecommendationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
