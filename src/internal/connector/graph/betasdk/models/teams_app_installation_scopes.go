package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type TeamsAppInstallationScopes int

const (
    TEAM_TEAMSAPPINSTALLATIONSCOPES TeamsAppInstallationScopes = iota
    GROUPCHAT_TEAMSAPPINSTALLATIONSCOPES
    PERSONAL_TEAMSAPPINSTALLATIONSCOPES
    UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES
)

func (i TeamsAppInstallationScopes) String() string {
    return []string{"team", "groupChat", "personal", "unknownFutureValue"}[i]
}
func ParseTeamsAppInstallationScopes(v string) (interface{}, error) {
    result := TEAM_TEAMSAPPINSTALLATIONSCOPES
    switch v {
        case "team":
            result = TEAM_TEAMSAPPINSTALLATIONSCOPES
        case "groupChat":
            result = GROUPCHAT_TEAMSAPPINSTALLATIONSCOPES
        case "personal":
            result = PERSONAL_TEAMSAPPINSTALLATIONSCOPES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES
        default:
            return 0, errors.New("Unknown TeamsAppInstallationScopes value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAppInstallationScopes(values []TeamsAppInstallationScopes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
