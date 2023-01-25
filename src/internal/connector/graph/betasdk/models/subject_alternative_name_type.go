package models
import (
    "errors"
)
// Provides operations to call the add method.
type SubjectAlternativeNameType int

const (
    // No subject alternative name.
    NONE_SUBJECTALTERNATIVENAMETYPE SubjectAlternativeNameType = iota
    // Email address.
    EMAILADDRESS_SUBJECTALTERNATIVENAMETYPE
    // User Principal Name (UPN).
    USERPRINCIPALNAME_SUBJECTALTERNATIVENAMETYPE
    // Custom Azure AD Attribute.
    CUSTOMAZUREADATTRIBUTE_SUBJECTALTERNATIVENAMETYPE
    // Domain Name Service (DNS).
    DOMAINNAMESERVICE_SUBJECTALTERNATIVENAMETYPE
    // Universal Resource Identifier (URI).
    UNIVERSALRESOURCEIDENTIFIER_SUBJECTALTERNATIVENAMETYPE
)

func (i SubjectAlternativeNameType) String() string {
    return []string{"none", "emailAddress", "userPrincipalName", "customAzureADAttribute", "domainNameService", "universalResourceIdentifier"}[i]
}
func ParseSubjectAlternativeNameType(v string) (interface{}, error) {
    result := NONE_SUBJECTALTERNATIVENAMETYPE
    switch v {
        case "none":
            result = NONE_SUBJECTALTERNATIVENAMETYPE
        case "emailAddress":
            result = EMAILADDRESS_SUBJECTALTERNATIVENAMETYPE
        case "userPrincipalName":
            result = USERPRINCIPALNAME_SUBJECTALTERNATIVENAMETYPE
        case "customAzureADAttribute":
            result = CUSTOMAZUREADATTRIBUTE_SUBJECTALTERNATIVENAMETYPE
        case "domainNameService":
            result = DOMAINNAMESERVICE_SUBJECTALTERNATIVENAMETYPE
        case "universalResourceIdentifier":
            result = UNIVERSALRESOURCEIDENTIFIER_SUBJECTALTERNATIVENAMETYPE
        default:
            return 0, errors.New("Unknown SubjectAlternativeNameType value: " + v)
    }
    return &result, nil
}
func SerializeSubjectAlternativeNameType(values []SubjectAlternativeNameType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
