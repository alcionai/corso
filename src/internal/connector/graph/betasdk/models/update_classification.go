package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UpdateClassification int

const (
    // User Defined, default value, no intent.
    USERDEFINED_UPDATECLASSIFICATION UpdateClassification = iota
    // Recommended and important.
    RECOMMENDEDANDIMPORTANT_UPDATECLASSIFICATION
    // Important.
    IMPORTANT_UPDATECLASSIFICATION
    // None.
    NONE_UPDATECLASSIFICATION
)

func (i UpdateClassification) String() string {
    return []string{"userDefined", "recommendedAndImportant", "important", "none"}[i]
}
func ParseUpdateClassification(v string) (interface{}, error) {
    result := USERDEFINED_UPDATECLASSIFICATION
    switch v {
        case "userDefined":
            result = USERDEFINED_UPDATECLASSIFICATION
        case "recommendedAndImportant":
            result = RECOMMENDEDANDIMPORTANT_UPDATECLASSIFICATION
        case "important":
            result = IMPORTANT_UPDATECLASSIFICATION
        case "none":
            result = NONE_UPDATECLASSIFICATION
        default:
            return 0, errors.New("Unknown UpdateClassification value: " + v)
    }
    return &result, nil
}
func SerializeUpdateClassification(values []UpdateClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
