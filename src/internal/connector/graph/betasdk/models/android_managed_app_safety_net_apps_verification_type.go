package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AndroidManagedAppSafetyNetAppsVerificationType int

const (
    // no requirement set
    NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE AndroidManagedAppSafetyNetAppsVerificationType = iota
    // require that Android device has SafetyNet Apps Verification enabled
    ENABLED_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
)

func (i AndroidManagedAppSafetyNetAppsVerificationType) String() string {
    return []string{"none", "enabled"}[i]
}
func ParseAndroidManagedAppSafetyNetAppsVerificationType(v string) (interface{}, error) {
    result := NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
    switch v {
        case "none":
            result = NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
        case "enabled":
            result = ENABLED_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
        default:
            return 0, errors.New("Unknown AndroidManagedAppSafetyNetAppsVerificationType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedAppSafetyNetAppsVerificationType(values []AndroidManagedAppSafetyNetAppsVerificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
