package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UsernameSource int

const (
    // User principal name.
    USERPRINCIPALNAME_USERNAMESOURCE UsernameSource = iota
    // Primary SMTP address.
    PRIMARYSMTPADDRESS_USERNAMESOURCE
    // The user sam account name.
    SAMACCOUNTNAME_USERNAMESOURCE
)

func (i UsernameSource) String() string {
    return []string{"userPrincipalName", "primarySmtpAddress", "samAccountName"}[i]
}
func ParseUsernameSource(v string) (interface{}, error) {
    result := USERPRINCIPALNAME_USERNAMESOURCE
    switch v {
        case "userPrincipalName":
            result = USERPRINCIPALNAME_USERNAMESOURCE
        case "primarySmtpAddress":
            result = PRIMARYSMTPADDRESS_USERNAMESOURCE
        case "samAccountName":
            result = SAMACCOUNTNAME_USERNAMESOURCE
        default:
            return 0, errors.New("Unknown UsernameSource value: " + v)
    }
    return &result, nil
}
func SerializeUsernameSource(values []UsernameSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
