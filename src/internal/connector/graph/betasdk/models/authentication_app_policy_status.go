package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AuthenticationAppPolicyStatus int

const (
    UNKNOWN_AUTHENTICATIONAPPPOLICYSTATUS AuthenticationAppPolicyStatus = iota
    APPLOCKOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
    APPLOCKENABLED_AUTHENTICATIONAPPPOLICYSTATUS
    APPLOCKDISABLED_AUTHENTICATIONAPPPOLICYSTATUS
    APPCONTEXTOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
    APPCONTEXTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
    APPCONTEXTNOTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
    LOCATIONCONTEXTOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
    LOCATIONCONTEXTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
    LOCATIONCONTEXTNOTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
    NUMBERMATCHOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
    NUMBERMATCHCORRECTNUMBERENTERED_AUTHENTICATIONAPPPOLICYSTATUS
    NUMBERMATCHINCORRECTNUMBERENTERED_AUTHENTICATIONAPPPOLICYSTATUS
    NUMBERMATCHDENY_AUTHENTICATIONAPPPOLICYSTATUS
    TAMPERRESISTANTHARDWAREOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
    TAMPERRESISTANTHARDWAREUSED_AUTHENTICATIONAPPPOLICYSTATUS
    TAMPERRESISTANTHARDWARENOTUSED_AUTHENTICATIONAPPPOLICYSTATUS
    UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPPOLICYSTATUS
)

func (i AuthenticationAppPolicyStatus) String() string {
    return []string{"unknown", "appLockOutOfDate", "appLockEnabled", "appLockDisabled", "appContextOutOfDate", "appContextShown", "appContextNotShown", "locationContextOutOfDate", "locationContextShown", "locationContextNotShown", "numberMatchOutOfDate", "numberMatchCorrectNumberEntered", "numberMatchIncorrectNumberEntered", "numberMatchDeny", "tamperResistantHardwareOutOfDate", "tamperResistantHardwareUsed", "tamperResistantHardwareNotUsed", "unknownFutureValue"}[i]
}
func ParseAuthenticationAppPolicyStatus(v string) (interface{}, error) {
    result := UNKNOWN_AUTHENTICATIONAPPPOLICYSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_AUTHENTICATIONAPPPOLICYSTATUS
        case "appLockOutOfDate":
            result = APPLOCKOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
        case "appLockEnabled":
            result = APPLOCKENABLED_AUTHENTICATIONAPPPOLICYSTATUS
        case "appLockDisabled":
            result = APPLOCKDISABLED_AUTHENTICATIONAPPPOLICYSTATUS
        case "appContextOutOfDate":
            result = APPCONTEXTOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
        case "appContextShown":
            result = APPCONTEXTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
        case "appContextNotShown":
            result = APPCONTEXTNOTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
        case "locationContextOutOfDate":
            result = LOCATIONCONTEXTOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
        case "locationContextShown":
            result = LOCATIONCONTEXTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
        case "locationContextNotShown":
            result = LOCATIONCONTEXTNOTSHOWN_AUTHENTICATIONAPPPOLICYSTATUS
        case "numberMatchOutOfDate":
            result = NUMBERMATCHOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
        case "numberMatchCorrectNumberEntered":
            result = NUMBERMATCHCORRECTNUMBERENTERED_AUTHENTICATIONAPPPOLICYSTATUS
        case "numberMatchIncorrectNumberEntered":
            result = NUMBERMATCHINCORRECTNUMBERENTERED_AUTHENTICATIONAPPPOLICYSTATUS
        case "numberMatchDeny":
            result = NUMBERMATCHDENY_AUTHENTICATIONAPPPOLICYSTATUS
        case "tamperResistantHardwareOutOfDate":
            result = TAMPERRESISTANTHARDWAREOUTOFDATE_AUTHENTICATIONAPPPOLICYSTATUS
        case "tamperResistantHardwareUsed":
            result = TAMPERRESISTANTHARDWAREUSED_AUTHENTICATIONAPPPOLICYSTATUS
        case "tamperResistantHardwareNotUsed":
            result = TAMPERRESISTANTHARDWARENOTUSED_AUTHENTICATIONAPPPOLICYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPPOLICYSTATUS
        default:
            return 0, errors.New("Unknown AuthenticationAppPolicyStatus value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationAppPolicyStatus(values []AuthenticationAppPolicyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
