package ediscovery
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SourceType int

const (
    MAILBOX_SOURCETYPE SourceType = iota
    SITE_SOURCETYPE
)

func (i SourceType) String() string {
    return []string{"mailbox", "site"}[i]
}
func ParseSourceType(v string) (interface{}, error) {
    result := MAILBOX_SOURCETYPE
    switch v {
        case "mailbox":
            result = MAILBOX_SOURCETYPE
        case "site":
            result = SITE_SOURCETYPE
        default:
            return 0, errors.New("Unknown SourceType value: " + v)
    }
    return &result, nil
}
func SerializeSourceType(values []SourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
