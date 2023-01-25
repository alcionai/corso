package models
import (
    "errors"
)
// Provides operations to call the add method.
type DelegatedAdminAccessContainerType int

const (
    SECURITYGROUP_DELEGATEDADMINACCESSCONTAINERTYPE DelegatedAdminAccessContainerType = iota
    UNKNOWNFUTUREVALUE_DELEGATEDADMINACCESSCONTAINERTYPE
)

func (i DelegatedAdminAccessContainerType) String() string {
    return []string{"securityGroup", "unknownFutureValue"}[i]
}
func ParseDelegatedAdminAccessContainerType(v string) (interface{}, error) {
    result := SECURITYGROUP_DELEGATEDADMINACCESSCONTAINERTYPE
    switch v {
        case "securityGroup":
            result = SECURITYGROUP_DELEGATEDADMINACCESSCONTAINERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINACCESSCONTAINERTYPE
        default:
            return 0, errors.New("Unknown DelegatedAdminAccessContainerType value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminAccessContainerType(values []DelegatedAdminAccessContainerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
