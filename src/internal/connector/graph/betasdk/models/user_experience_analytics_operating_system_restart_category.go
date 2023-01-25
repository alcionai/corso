package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type UserExperienceAnalyticsOperatingSystemRestartCategory int

const (
    // Unknown
    UNKNOWN_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY UserExperienceAnalyticsOperatingSystemRestartCategory = iota
    // Restart with update
    RESTARTWITHUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    // Restart without update
    RESTARTWITHOUTUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    // Blue screen restart
    BLUESCREEN_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    // Shutdown with update
    SHUTDOWNWITHUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    // Shutdown without update
    SHUTDOWNWITHOUTUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    // Long power button press
    LONGPOWERBUTTONPRESS_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    // Boot error
    BOOTERROR_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    // Update
    UPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
)

func (i UserExperienceAnalyticsOperatingSystemRestartCategory) String() string {
    return []string{"unknown", "restartWithUpdate", "restartWithoutUpdate", "blueScreen", "shutdownWithUpdate", "shutdownWithoutUpdate", "longPowerButtonPress", "bootError", "update"}[i]
}
func ParseUserExperienceAnalyticsOperatingSystemRestartCategory(v string) (interface{}, error) {
    result := UNKNOWN_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
    switch v {
        case "unknown":
            result = UNKNOWN_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "restartWithUpdate":
            result = RESTARTWITHUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "restartWithoutUpdate":
            result = RESTARTWITHOUTUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "blueScreen":
            result = BLUESCREEN_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "shutdownWithUpdate":
            result = SHUTDOWNWITHUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "shutdownWithoutUpdate":
            result = SHUTDOWNWITHOUTUPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "longPowerButtonPress":
            result = LONGPOWERBUTTONPRESS_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "bootError":
            result = BOOTERROR_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        case "update":
            result = UPDATE_USEREXPERIENCEANALYTICSOPERATINGSYSTEMRESTARTCATEGORY
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsOperatingSystemRestartCategory value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsOperatingSystemRestartCategory(values []UserExperienceAnalyticsOperatingSystemRestartCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
