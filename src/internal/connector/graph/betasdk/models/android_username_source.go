package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidUsernameSource int

const (
    // The username.
    USERNAME_ANDROIDUSERNAMESOURCE AndroidUsernameSource = iota
    // The user principal name.
    USERPRINCIPALNAME_ANDROIDUSERNAMESOURCE
    // The user sam account name.
    SAMACCOUNTNAME_ANDROIDUSERNAMESOURCE
    // Primary SMTP address.
    PRIMARYSMTPADDRESS_ANDROIDUSERNAMESOURCE
)

func (i AndroidUsernameSource) String() string {
    return []string{"username", "userPrincipalName", "samAccountName", "primarySmtpAddress"}[i]
}
func ParseAndroidUsernameSource(v string) (interface{}, error) {
    result := USERNAME_ANDROIDUSERNAMESOURCE
    switch v {
        case "username":
            result = USERNAME_ANDROIDUSERNAMESOURCE
        case "userPrincipalName":
            result = USERPRINCIPALNAME_ANDROIDUSERNAMESOURCE
        case "samAccountName":
            result = SAMACCOUNTNAME_ANDROIDUSERNAMESOURCE
        case "primarySmtpAddress":
            result = PRIMARYSMTPADDRESS_ANDROIDUSERNAMESOURCE
        default:
            return 0, errors.New("Unknown AndroidUsernameSource value: " + v)
    }
    return &result, nil
}
func SerializeAndroidUsernameSource(values []AndroidUsernameSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
