package security
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SubmissionClientSource int

const (
    MICROSOFT_SUBMISSIONCLIENTSOURCE SubmissionClientSource = iota
    OTHER_SUBMISSIONCLIENTSOURCE
    UNKNOWNFUTUREVALUE_SUBMISSIONCLIENTSOURCE
)

func (i SubmissionClientSource) String() string {
    return []string{"microsoft", "other", "unknownFutureValue"}[i]
}
func ParseSubmissionClientSource(v string) (interface{}, error) {
    result := MICROSOFT_SUBMISSIONCLIENTSOURCE
    switch v {
        case "microsoft":
            result = MICROSOFT_SUBMISSIONCLIENTSOURCE
        case "other":
            result = OTHER_SUBMISSIONCLIENTSOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBMISSIONCLIENTSOURCE
        default:
            return 0, errors.New("Unknown SubmissionClientSource value: " + v)
    }
    return &result, nil
}
func SerializeSubmissionClientSource(values []SubmissionClientSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
