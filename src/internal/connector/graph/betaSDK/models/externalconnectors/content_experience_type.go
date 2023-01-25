package externalconnectors
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ContentExperienceType int

const (
    SEARCH_CONTENTEXPERIENCETYPE ContentExperienceType = iota
    COMPLIANCE_CONTENTEXPERIENCETYPE
    UNKNOWNFUTUREVALUE_CONTENTEXPERIENCETYPE
)

func (i ContentExperienceType) String() string {
    return []string{"search", "compliance", "unknownFutureValue"}[i]
}
func ParseContentExperienceType(v string) (interface{}, error) {
    result := SEARCH_CONTENTEXPERIENCETYPE
    switch v {
        case "search":
            result = SEARCH_CONTENTEXPERIENCETYPE
        case "compliance":
            result = COMPLIANCE_CONTENTEXPERIENCETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONTENTEXPERIENCETYPE
        default:
            return 0, errors.New("Unknown ContentExperienceType value: " + v)
    }
    return &result, nil
}
func SerializeContentExperienceType(values []ContentExperienceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
