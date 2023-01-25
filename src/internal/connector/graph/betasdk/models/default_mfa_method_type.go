package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DefaultMfaMethodType int

const (
    NONE_DEFAULTMFAMETHODTYPE DefaultMfaMethodType = iota
    MOBILEPHONE_DEFAULTMFAMETHODTYPE
    ALTERNATEMOBILEPHONE_DEFAULTMFAMETHODTYPE
    OFFICEPHONE_DEFAULTMFAMETHODTYPE
    MICROSOFTAUTHENTICATORPUSH_DEFAULTMFAMETHODTYPE
    SOFTWAREONETIMEPASSCODE_DEFAULTMFAMETHODTYPE
    UNKNOWNFUTUREVALUE_DEFAULTMFAMETHODTYPE
)

func (i DefaultMfaMethodType) String() string {
    return []string{"none", "mobilePhone", "alternateMobilePhone", "officePhone", "microsoftAuthenticatorPush", "softwareOneTimePasscode", "unknownFutureValue"}[i]
}
func ParseDefaultMfaMethodType(v string) (interface{}, error) {
    result := NONE_DEFAULTMFAMETHODTYPE
    switch v {
        case "none":
            result = NONE_DEFAULTMFAMETHODTYPE
        case "mobilePhone":
            result = MOBILEPHONE_DEFAULTMFAMETHODTYPE
        case "alternateMobilePhone":
            result = ALTERNATEMOBILEPHONE_DEFAULTMFAMETHODTYPE
        case "officePhone":
            result = OFFICEPHONE_DEFAULTMFAMETHODTYPE
        case "microsoftAuthenticatorPush":
            result = MICROSOFTAUTHENTICATORPUSH_DEFAULTMFAMETHODTYPE
        case "softwareOneTimePasscode":
            result = SOFTWAREONETIMEPASSCODE_DEFAULTMFAMETHODTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEFAULTMFAMETHODTYPE
        default:
            return 0, errors.New("Unknown DefaultMfaMethodType value: " + v)
    }
    return &result, nil
}
func SerializeDefaultMfaMethodType(values []DefaultMfaMethodType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
