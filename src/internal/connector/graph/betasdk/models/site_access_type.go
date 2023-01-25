package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SiteAccessType int

const (
    BLOCK_SITEACCESSTYPE SiteAccessType = iota
    FULL_SITEACCESSTYPE
    LIMITED_SITEACCESSTYPE
)

func (i SiteAccessType) String() string {
    return []string{"block", "full", "limited"}[i]
}
func ParseSiteAccessType(v string) (interface{}, error) {
    result := BLOCK_SITEACCESSTYPE
    switch v {
        case "block":
            result = BLOCK_SITEACCESSTYPE
        case "full":
            result = FULL_SITEACCESSTYPE
        case "limited":
            result = LIMITED_SITEACCESSTYPE
        default:
            return 0, errors.New("Unknown SiteAccessType value: " + v)
    }
    return &result, nil
}
func SerializeSiteAccessType(values []SiteAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
