package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AutoAdmittedUsersType int

const (
    EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE AutoAdmittedUsersType = iota
    EVERYONE_AUTOADMITTEDUSERSTYPE
)

func (i AutoAdmittedUsersType) String() string {
    return []string{"everyoneInCompany", "everyone"}[i]
}
func ParseAutoAdmittedUsersType(v string) (interface{}, error) {
    result := EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE
    switch v {
        case "everyoneInCompany":
            result = EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE
        case "everyone":
            result = EVERYONE_AUTOADMITTEDUSERSTYPE
        default:
            return 0, errors.New("Unknown AutoAdmittedUsersType value: " + v)
    }
    return &result, nil
}
func SerializeAutoAdmittedUsersType(values []AutoAdmittedUsersType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
