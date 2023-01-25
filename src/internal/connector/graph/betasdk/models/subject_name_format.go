package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SubjectNameFormat int

const (
    // Common name.
    COMMONNAME_SUBJECTNAMEFORMAT SubjectNameFormat = iota
    // Common Name Including Email.
    COMMONNAMEINCLUDINGEMAIL_SUBJECTNAMEFORMAT
    // Common Name As Email.
    COMMONNAMEASEMAIL_SUBJECTNAMEFORMAT
    // Custom subject name format.
    CUSTOM_SUBJECTNAMEFORMAT
    // Common Name As IMEI.
    COMMONNAMEASIMEI_SUBJECTNAMEFORMAT
    // Common Name As Serial Number.
    COMMONNAMEASSERIALNUMBER_SUBJECTNAMEFORMAT
    // Common Name As Serial Number.
    COMMONNAMEASAADDEVICEID_SUBJECTNAMEFORMAT
    // Common Name As Serial Number.
    COMMONNAMEASINTUNEDEVICEID_SUBJECTNAMEFORMAT
    // Common Name As Serial Number.
    COMMONNAMEASDURABLEDEVICEID_SUBJECTNAMEFORMAT
)

func (i SubjectNameFormat) String() string {
    return []string{"commonName", "commonNameIncludingEmail", "commonNameAsEmail", "custom", "commonNameAsIMEI", "commonNameAsSerialNumber", "commonNameAsAadDeviceId", "commonNameAsIntuneDeviceId", "commonNameAsDurableDeviceId"}[i]
}
func ParseSubjectNameFormat(v string) (interface{}, error) {
    result := COMMONNAME_SUBJECTNAMEFORMAT
    switch v {
        case "commonName":
            result = COMMONNAME_SUBJECTNAMEFORMAT
        case "commonNameIncludingEmail":
            result = COMMONNAMEINCLUDINGEMAIL_SUBJECTNAMEFORMAT
        case "commonNameAsEmail":
            result = COMMONNAMEASEMAIL_SUBJECTNAMEFORMAT
        case "custom":
            result = CUSTOM_SUBJECTNAMEFORMAT
        case "commonNameAsIMEI":
            result = COMMONNAMEASIMEI_SUBJECTNAMEFORMAT
        case "commonNameAsSerialNumber":
            result = COMMONNAMEASSERIALNUMBER_SUBJECTNAMEFORMAT
        case "commonNameAsAadDeviceId":
            result = COMMONNAMEASAADDEVICEID_SUBJECTNAMEFORMAT
        case "commonNameAsIntuneDeviceId":
            result = COMMONNAMEASINTUNEDEVICEID_SUBJECTNAMEFORMAT
        case "commonNameAsDurableDeviceId":
            result = COMMONNAMEASDURABLEDEVICEID_SUBJECTNAMEFORMAT
        default:
            return 0, errors.New("Unknown SubjectNameFormat value: " + v)
    }
    return &result, nil
}
func SerializeSubjectNameFormat(values []SubjectNameFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
