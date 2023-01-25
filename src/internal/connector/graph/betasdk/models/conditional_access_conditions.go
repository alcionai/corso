package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ConditionalAccessConditions int

const (
    NONE_CONDITIONALACCESSCONDITIONS ConditionalAccessConditions = iota
    APPLICATION_CONDITIONALACCESSCONDITIONS
    USERS_CONDITIONALACCESSCONDITIONS
    DEVICEPLATFORM_CONDITIONALACCESSCONDITIONS
    LOCATION_CONDITIONALACCESSCONDITIONS
    CLIENTTYPE_CONDITIONALACCESSCONDITIONS
    SIGNINRISK_CONDITIONALACCESSCONDITIONS
    USERRISK_CONDITIONALACCESSCONDITIONS
    TIME_CONDITIONALACCESSCONDITIONS
    DEVICESTATE_CONDITIONALACCESSCONDITIONS
    CLIENT_CONDITIONALACCESSCONDITIONS
    IPADDRESSSEENBYAZUREAD_CONDITIONALACCESSCONDITIONS
    IPADDRESSSEENBYRESOURCEPROVIDER_CONDITIONALACCESSCONDITIONS
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSCONDITIONS
    SERVICEPRINCIPALS_CONDITIONALACCESSCONDITIONS
    SERVICEPRINCIPALRISK_CONDITIONALACCESSCONDITIONS
)

func (i ConditionalAccessConditions) String() string {
    return []string{"none", "application", "users", "devicePlatform", "location", "clientType", "signInRisk", "userRisk", "time", "deviceState", "client", "ipAddressSeenByAzureAD", "ipAddressSeenByResourceProvider", "unknownFutureValue", "servicePrincipals", "servicePrincipalRisk"}[i]
}
func ParseConditionalAccessConditions(v string) (interface{}, error) {
    result := NONE_CONDITIONALACCESSCONDITIONS
    switch v {
        case "none":
            result = NONE_CONDITIONALACCESSCONDITIONS
        case "application":
            result = APPLICATION_CONDITIONALACCESSCONDITIONS
        case "users":
            result = USERS_CONDITIONALACCESSCONDITIONS
        case "devicePlatform":
            result = DEVICEPLATFORM_CONDITIONALACCESSCONDITIONS
        case "location":
            result = LOCATION_CONDITIONALACCESSCONDITIONS
        case "clientType":
            result = CLIENTTYPE_CONDITIONALACCESSCONDITIONS
        case "signInRisk":
            result = SIGNINRISK_CONDITIONALACCESSCONDITIONS
        case "userRisk":
            result = USERRISK_CONDITIONALACCESSCONDITIONS
        case "time":
            result = TIME_CONDITIONALACCESSCONDITIONS
        case "deviceState":
            result = DEVICESTATE_CONDITIONALACCESSCONDITIONS
        case "client":
            result = CLIENT_CONDITIONALACCESSCONDITIONS
        case "ipAddressSeenByAzureAD":
            result = IPADDRESSSEENBYAZUREAD_CONDITIONALACCESSCONDITIONS
        case "ipAddressSeenByResourceProvider":
            result = IPADDRESSSEENBYRESOURCEPROVIDER_CONDITIONALACCESSCONDITIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSCONDITIONS
        case "servicePrincipals":
            result = SERVICEPRINCIPALS_CONDITIONALACCESSCONDITIONS
        case "servicePrincipalRisk":
            result = SERVICEPRINCIPALRISK_CONDITIONALACCESSCONDITIONS
        default:
            return 0, errors.New("Unknown ConditionalAccessConditions value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessConditions(values []ConditionalAccessConditions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
