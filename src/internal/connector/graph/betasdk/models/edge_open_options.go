package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EdgeOpenOptions int

const (
    // Not configured.
    NOTCONFIGURED_EDGEOPENOPTIONS EdgeOpenOptions = iota
    // StartPage.
    STARTPAGE_EDGEOPENOPTIONS
    // NewTabPage.
    NEWTABPAGE_EDGEOPENOPTIONS
    // PreviousPages.
    PREVIOUSPAGES_EDGEOPENOPTIONS
    // SpecificPages.
    SPECIFICPAGES_EDGEOPENOPTIONS
)

func (i EdgeOpenOptions) String() string {
    return []string{"notConfigured", "startPage", "newTabPage", "previousPages", "specificPages"}[i]
}
func ParseEdgeOpenOptions(v string) (interface{}, error) {
    result := NOTCONFIGURED_EDGEOPENOPTIONS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_EDGEOPENOPTIONS
        case "startPage":
            result = STARTPAGE_EDGEOPENOPTIONS
        case "newTabPage":
            result = NEWTABPAGE_EDGEOPENOPTIONS
        case "previousPages":
            result = PREVIOUSPAGES_EDGEOPENOPTIONS
        case "specificPages":
            result = SPECIFICPAGES_EDGEOPENOPTIONS
        default:
            return 0, errors.New("Unknown EdgeOpenOptions value: " + v)
    }
    return &result, nil
}
func SerializeEdgeOpenOptions(values []EdgeOpenOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
