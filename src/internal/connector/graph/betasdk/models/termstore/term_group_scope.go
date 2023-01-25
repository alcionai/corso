package termstore
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TermGroupScope int

const (
    GLOBAL_TERMGROUPSCOPE TermGroupScope = iota
    SYSTEM_TERMGROUPSCOPE
    SITECOLLECTION_TERMGROUPSCOPE
)

func (i TermGroupScope) String() string {
    return []string{"global", "system", "siteCollection"}[i]
}
func ParseTermGroupScope(v string) (interface{}, error) {
    result := GLOBAL_TERMGROUPSCOPE
    switch v {
        case "global":
            result = GLOBAL_TERMGROUPSCOPE
        case "system":
            result = SYSTEM_TERMGROUPSCOPE
        case "siteCollection":
            result = SITECOLLECTION_TERMGROUPSCOPE
        default:
            return 0, errors.New("Unknown TermGroupScope value: " + v)
    }
    return &result, nil
}
func SerializeTermGroupScope(values []TermGroupScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
