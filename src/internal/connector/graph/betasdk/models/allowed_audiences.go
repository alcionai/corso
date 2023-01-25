package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AllowedAudiences int

const (
    ME_ALLOWEDAUDIENCES AllowedAudiences = iota
    FAMILY_ALLOWEDAUDIENCES
    CONTACTS_ALLOWEDAUDIENCES
    GROUPMEMBERS_ALLOWEDAUDIENCES
    ORGANIZATION_ALLOWEDAUDIENCES
    FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES
    EVERYONE_ALLOWEDAUDIENCES
    UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES
)

func (i AllowedAudiences) String() string {
    return []string{"me", "family", "contacts", "groupMembers", "organization", "federatedOrganizations", "everyone", "unknownFutureValue"}[i]
}
func ParseAllowedAudiences(v string) (interface{}, error) {
    result := ME_ALLOWEDAUDIENCES
    switch v {
        case "me":
            result = ME_ALLOWEDAUDIENCES
        case "family":
            result = FAMILY_ALLOWEDAUDIENCES
        case "contacts":
            result = CONTACTS_ALLOWEDAUDIENCES
        case "groupMembers":
            result = GROUPMEMBERS_ALLOWEDAUDIENCES
        case "organization":
            result = ORGANIZATION_ALLOWEDAUDIENCES
        case "federatedOrganizations":
            result = FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES
        case "everyone":
            result = EVERYONE_ALLOWEDAUDIENCES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES
        default:
            return 0, errors.New("Unknown AllowedAudiences value: " + v)
    }
    return &result, nil
}
func SerializeAllowedAudiences(values []AllowedAudiences) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
