package models
import (
    "errors"
)
// Provides operations to call the add method.
type DomainNameSource int

const (
    // Full domain name.
    FULLDOMAINNAME_DOMAINNAMESOURCE DomainNameSource = iota
    // net bios domain name.
    NETBIOSDOMAINNAME_DOMAINNAMESOURCE
)

func (i DomainNameSource) String() string {
    return []string{"fullDomainName", "netBiosDomainName"}[i]
}
func ParseDomainNameSource(v string) (interface{}, error) {
    result := FULLDOMAINNAME_DOMAINNAMESOURCE
    switch v {
        case "fullDomainName":
            result = FULLDOMAINNAME_DOMAINNAMESOURCE
        case "netBiosDomainName":
            result = NETBIOSDOMAINNAME_DOMAINNAMESOURCE
        default:
            return 0, errors.New("Unknown DomainNameSource value: " + v)
    }
    return &result, nil
}
func SerializeDomainNameSource(values []DomainNameSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
