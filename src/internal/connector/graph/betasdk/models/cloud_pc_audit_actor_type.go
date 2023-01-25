package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcAuditActorType int

const (
    ITPRO_CLOUDPCAUDITACTORTYPE CloudPcAuditActorType = iota
    APPLICATION_CLOUDPCAUDITACTORTYPE
    PARTNER_CLOUDPCAUDITACTORTYPE
    UNKNOWN_CLOUDPCAUDITACTORTYPE
)

func (i CloudPcAuditActorType) String() string {
    return []string{"itPro", "application", "partner", "unknown"}[i]
}
func ParseCloudPcAuditActorType(v string) (interface{}, error) {
    result := ITPRO_CLOUDPCAUDITACTORTYPE
    switch v {
        case "itPro":
            result = ITPRO_CLOUDPCAUDITACTORTYPE
        case "application":
            result = APPLICATION_CLOUDPCAUDITACTORTYPE
        case "partner":
            result = PARTNER_CLOUDPCAUDITACTORTYPE
        case "unknown":
            result = UNKNOWN_CLOUDPCAUDITACTORTYPE
        default:
            return 0, errors.New("Unknown CloudPcAuditActorType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditActorType(values []CloudPcAuditActorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
