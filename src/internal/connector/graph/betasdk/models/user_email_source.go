package models
import (
    "errors"
)
// Provides operations to call the add method.
type UserEmailSource int

const (
    // User principal name.
    USERPRINCIPALNAME_USEREMAILSOURCE UserEmailSource = iota
    // Primary SMTP address.
    PRIMARYSMTPADDRESS_USEREMAILSOURCE
)

func (i UserEmailSource) String() string {
    return []string{"userPrincipalName", "primarySmtpAddress"}[i]
}
func ParseUserEmailSource(v string) (interface{}, error) {
    result := USERPRINCIPALNAME_USEREMAILSOURCE
    switch v {
        case "userPrincipalName":
            result = USERPRINCIPALNAME_USEREMAILSOURCE
        case "primarySmtpAddress":
            result = PRIMARYSMTPADDRESS_USEREMAILSOURCE
        default:
            return 0, errors.New("Unknown UserEmailSource value: " + v)
    }
    return &result, nil
}
func SerializeUserEmailSource(values []UserEmailSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
