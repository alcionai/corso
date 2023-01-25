package security
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type SubmissionContentType int

const (
    EMAIL_SUBMISSIONCONTENTTYPE SubmissionContentType = iota
    URL_SUBMISSIONCONTENTTYPE
    FILE_SUBMISSIONCONTENTTYPE
    APP_SUBMISSIONCONTENTTYPE
    UNKNOWNFUTUREVALUE_SUBMISSIONCONTENTTYPE
)

func (i SubmissionContentType) String() string {
    return []string{"email", "url", "file", "app", "unknownFutureValue"}[i]
}
func ParseSubmissionContentType(v string) (interface{}, error) {
    result := EMAIL_SUBMISSIONCONTENTTYPE
    switch v {
        case "email":
            result = EMAIL_SUBMISSIONCONTENTTYPE
        case "url":
            result = URL_SUBMISSIONCONTENTTYPE
        case "file":
            result = FILE_SUBMISSIONCONTENTTYPE
        case "app":
            result = APP_SUBMISSIONCONTENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBMISSIONCONTENTTYPE
        default:
            return 0, errors.New("Unknown SubmissionContentType value: " + v)
    }
    return &result, nil
}
func SerializeSubmissionContentType(values []SubmissionContentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
