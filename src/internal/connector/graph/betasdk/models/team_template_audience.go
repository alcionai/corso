package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TeamTemplateAudience int

const (
    ORGANIZATION_TEAMTEMPLATEAUDIENCE TeamTemplateAudience = iota
    USER_TEAMTEMPLATEAUDIENCE
    PUBLIC_TEAMTEMPLATEAUDIENCE
    UNKNOWNFUTUREVALUE_TEAMTEMPLATEAUDIENCE
)

func (i TeamTemplateAudience) String() string {
    return []string{"organization", "user", "public", "unknownFutureValue"}[i]
}
func ParseTeamTemplateAudience(v string) (interface{}, error) {
    result := ORGANIZATION_TEAMTEMPLATEAUDIENCE
    switch v {
        case "organization":
            result = ORGANIZATION_TEAMTEMPLATEAUDIENCE
        case "user":
            result = USER_TEAMTEMPLATEAUDIENCE
        case "public":
            result = PUBLIC_TEAMTEMPLATEAUDIENCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMTEMPLATEAUDIENCE
        default:
            return 0, errors.New("Unknown TeamTemplateAudience value: " + v)
    }
    return &result, nil
}
func SerializeTeamTemplateAudience(values []TeamTemplateAudience) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
