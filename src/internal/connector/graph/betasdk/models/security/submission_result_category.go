package security
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SubmissionResultCategory int

const (
    NOTJUNK_SUBMISSIONRESULTCATEGORY SubmissionResultCategory = iota
    SPAM_SUBMISSIONRESULTCATEGORY
    PHISHING_SUBMISSIONRESULTCATEGORY
    MALWARE_SUBMISSIONRESULTCATEGORY
    ALLOWEDBYPOLICY_SUBMISSIONRESULTCATEGORY
    BLOCKEDBYPOLICY_SUBMISSIONRESULTCATEGORY
    SPOOF_SUBMISSIONRESULTCATEGORY
    UNKNOWN_SUBMISSIONRESULTCATEGORY
    NORESULTAVAILABLE_SUBMISSIONRESULTCATEGORY
    UNKNOWNFUTUREVALUE_SUBMISSIONRESULTCATEGORY
)

func (i SubmissionResultCategory) String() string {
    return []string{"notJunk", "spam", "phishing", "malware", "allowedByPolicy", "blockedByPolicy", "spoof", "unknown", "noResultAvailable", "unknownFutureValue"}[i]
}
func ParseSubmissionResultCategory(v string) (interface{}, error) {
    result := NOTJUNK_SUBMISSIONRESULTCATEGORY
    switch v {
        case "notJunk":
            result = NOTJUNK_SUBMISSIONRESULTCATEGORY
        case "spam":
            result = SPAM_SUBMISSIONRESULTCATEGORY
        case "phishing":
            result = PHISHING_SUBMISSIONRESULTCATEGORY
        case "malware":
            result = MALWARE_SUBMISSIONRESULTCATEGORY
        case "allowedByPolicy":
            result = ALLOWEDBYPOLICY_SUBMISSIONRESULTCATEGORY
        case "blockedByPolicy":
            result = BLOCKEDBYPOLICY_SUBMISSIONRESULTCATEGORY
        case "spoof":
            result = SPOOF_SUBMISSIONRESULTCATEGORY
        case "unknown":
            result = UNKNOWN_SUBMISSIONRESULTCATEGORY
        case "noResultAvailable":
            result = NORESULTAVAILABLE_SUBMISSIONRESULTCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBMISSIONRESULTCATEGORY
        default:
            return 0, errors.New("Unknown SubmissionResultCategory value: " + v)
    }
    return &result, nil
}
func SerializeSubmissionResultCategory(values []SubmissionResultCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
