package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DepTokenType int

const (
    // Token Type is None
    NONE_DEPTOKENTYPE DepTokenType = iota
    // Token Type is Dep.
    DEP_DEPTOKENTYPE
    // Token Type is Apple School Manager
    APPLESCHOOLMANAGER_DEPTOKENTYPE
)

func (i DepTokenType) String() string {
    return []string{"none", "dep", "appleSchoolManager"}[i]
}
func ParseDepTokenType(v string) (interface{}, error) {
    result := NONE_DEPTOKENTYPE
    switch v {
        case "none":
            result = NONE_DEPTOKENTYPE
        case "dep":
            result = DEP_DEPTOKENTYPE
        case "appleSchoolManager":
            result = APPLESCHOOLMANAGER_DEPTOKENTYPE
        default:
            return 0, errors.New("Unknown DepTokenType value: " + v)
    }
    return &result, nil
}
func SerializeDepTokenType(values []DepTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
