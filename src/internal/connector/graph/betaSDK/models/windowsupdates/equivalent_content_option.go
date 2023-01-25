package windowsupdates
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EquivalentContentOption int

const (
    NONE_EQUIVALENTCONTENTOPTION EquivalentContentOption = iota
    LATESTSECURITY_EQUIVALENTCONTENTOPTION
    UNKNOWNFUTUREVALUE_EQUIVALENTCONTENTOPTION
)

func (i EquivalentContentOption) String() string {
    return []string{"none", "latestSecurity", "unknownFutureValue"}[i]
}
func ParseEquivalentContentOption(v string) (interface{}, error) {
    result := NONE_EQUIVALENTCONTENTOPTION
    switch v {
        case "none":
            result = NONE_EQUIVALENTCONTENTOPTION
        case "latestSecurity":
            result = LATESTSECURITY_EQUIVALENTCONTENTOPTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EQUIVALENTCONTENTOPTION
        default:
            return 0, errors.New("Unknown EquivalentContentOption value: " + v)
    }
    return &result, nil
}
func SerializeEquivalentContentOption(values []EquivalentContentOption) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
