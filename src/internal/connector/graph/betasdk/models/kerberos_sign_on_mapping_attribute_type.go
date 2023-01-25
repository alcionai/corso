package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type KerberosSignOnMappingAttributeType int

const (
    USERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE KerberosSignOnMappingAttributeType = iota
    ONPREMISESUSERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    USERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    ONPREMISESUSERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    ONPREMISESSAMACCOUNTNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
)

func (i KerberosSignOnMappingAttributeType) String() string {
    return []string{"userPrincipalName", "onPremisesUserPrincipalName", "userPrincipalUsername", "onPremisesUserPrincipalUsername", "onPremisesSAMAccountName"}[i]
}
func ParseKerberosSignOnMappingAttributeType(v string) (interface{}, error) {
    result := USERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    switch v {
        case "userPrincipalName":
            result = USERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "onPremisesUserPrincipalName":
            result = ONPREMISESUSERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "userPrincipalUsername":
            result = USERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "onPremisesUserPrincipalUsername":
            result = ONPREMISESUSERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "onPremisesSAMAccountName":
            result = ONPREMISESSAMACCOUNTNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        default:
            return 0, errors.New("Unknown KerberosSignOnMappingAttributeType value: " + v)
    }
    return &result, nil
}
func SerializeKerberosSignOnMappingAttributeType(values []KerberosSignOnMappingAttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
