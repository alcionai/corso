package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ClassificationMethod int

const (
    PATTERNMATCH_CLASSIFICATIONMETHOD ClassificationMethod = iota
    EXACTDATAMATCH_CLASSIFICATIONMETHOD
    FINGERPRINT_CLASSIFICATIONMETHOD
    MACHINELEARNING_CLASSIFICATIONMETHOD
)

func (i ClassificationMethod) String() string {
    return []string{"patternMatch", "exactDataMatch", "fingerprint", "machineLearning"}[i]
}
func ParseClassificationMethod(v string) (interface{}, error) {
    result := PATTERNMATCH_CLASSIFICATIONMETHOD
    switch v {
        case "patternMatch":
            result = PATTERNMATCH_CLASSIFICATIONMETHOD
        case "exactDataMatch":
            result = EXACTDATAMATCH_CLASSIFICATIONMETHOD
        case "fingerprint":
            result = FINGERPRINT_CLASSIFICATIONMETHOD
        case "machineLearning":
            result = MACHINELEARNING_CLASSIFICATIONMETHOD
        default:
            return 0, errors.New("Unknown ClassificationMethod value: " + v)
    }
    return &result, nil
}
func SerializeClassificationMethod(values []ClassificationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
