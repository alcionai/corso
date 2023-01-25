package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AppleSubjectNameFormat int

const (
    // Common name.
    COMMONNAME_APPLESUBJECTNAMEFORMAT AppleSubjectNameFormat = iota
    // Common name as email.
    COMMONNAMEASEMAIL_APPLESUBJECTNAMEFORMAT
    // Custom subject name format.
    CUSTOM_APPLESUBJECTNAMEFORMAT
    // Common Name Including Email.
    COMMONNAMEINCLUDINGEMAIL_APPLESUBJECTNAMEFORMAT
    // Common Name As IMEI.
    COMMONNAMEASIMEI_APPLESUBJECTNAMEFORMAT
    // Common Name As Serial Number.
    COMMONNAMEASSERIALNUMBER_APPLESUBJECTNAMEFORMAT
)

func (i AppleSubjectNameFormat) String() string {
    return []string{"commonName", "commonNameAsEmail", "custom", "commonNameIncludingEmail", "commonNameAsIMEI", "commonNameAsSerialNumber"}[i]
}
func ParseAppleSubjectNameFormat(v string) (interface{}, error) {
    result := COMMONNAME_APPLESUBJECTNAMEFORMAT
    switch v {
        case "commonName":
            result = COMMONNAME_APPLESUBJECTNAMEFORMAT
        case "commonNameAsEmail":
            result = COMMONNAMEASEMAIL_APPLESUBJECTNAMEFORMAT
        case "custom":
            result = CUSTOM_APPLESUBJECTNAMEFORMAT
        case "commonNameIncludingEmail":
            result = COMMONNAMEINCLUDINGEMAIL_APPLESUBJECTNAMEFORMAT
        case "commonNameAsIMEI":
            result = COMMONNAMEASIMEI_APPLESUBJECTNAMEFORMAT
        case "commonNameAsSerialNumber":
            result = COMMONNAMEASSERIALNUMBER_APPLESUBJECTNAMEFORMAT
        default:
            return 0, errors.New("Unknown AppleSubjectNameFormat value: " + v)
    }
    return &result, nil
}
func SerializeAppleSubjectNameFormat(values []AppleSubjectNameFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
