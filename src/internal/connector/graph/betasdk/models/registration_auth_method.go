package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RegistrationAuthMethod int

const (
    EMAIL_REGISTRATIONAUTHMETHOD RegistrationAuthMethod = iota
    MOBILEPHONE_REGISTRATIONAUTHMETHOD
    OFFICEPHONE_REGISTRATIONAUTHMETHOD
    SECURITYQUESTION_REGISTRATIONAUTHMETHOD
    APPNOTIFICATION_REGISTRATIONAUTHMETHOD
    APPCODE_REGISTRATIONAUTHMETHOD
    ALTERNATEMOBILEPHONE_REGISTRATIONAUTHMETHOD
    FIDO_REGISTRATIONAUTHMETHOD
    APPPASSWORD_REGISTRATIONAUTHMETHOD
    UNKNOWNFUTUREVALUE_REGISTRATIONAUTHMETHOD
)

func (i RegistrationAuthMethod) String() string {
    return []string{"email", "mobilePhone", "officePhone", "securityQuestion", "appNotification", "appCode", "alternateMobilePhone", "fido", "appPassword", "unknownFutureValue"}[i]
}
func ParseRegistrationAuthMethod(v string) (interface{}, error) {
    result := EMAIL_REGISTRATIONAUTHMETHOD
    switch v {
        case "email":
            result = EMAIL_REGISTRATIONAUTHMETHOD
        case "mobilePhone":
            result = MOBILEPHONE_REGISTRATIONAUTHMETHOD
        case "officePhone":
            result = OFFICEPHONE_REGISTRATIONAUTHMETHOD
        case "securityQuestion":
            result = SECURITYQUESTION_REGISTRATIONAUTHMETHOD
        case "appNotification":
            result = APPNOTIFICATION_REGISTRATIONAUTHMETHOD
        case "appCode":
            result = APPCODE_REGISTRATIONAUTHMETHOD
        case "alternateMobilePhone":
            result = ALTERNATEMOBILEPHONE_REGISTRATIONAUTHMETHOD
        case "fido":
            result = FIDO_REGISTRATIONAUTHMETHOD
        case "appPassword":
            result = APPPASSWORD_REGISTRATIONAUTHMETHOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REGISTRATIONAUTHMETHOD
        default:
            return 0, errors.New("Unknown RegistrationAuthMethod value: " + v)
    }
    return &result, nil
}
func SerializeRegistrationAuthMethod(values []RegistrationAuthMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
