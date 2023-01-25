package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ManagedAppFlaggedReason int

const (
    // No issue.
    NONE_MANAGEDAPPFLAGGEDREASON ManagedAppFlaggedReason = iota
    // The app registration is running on a rooted/unlocked device.
    ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON
    // The app registration is running on an Android device on which the bootloader is unlocked.
    ANDROIDBOOTLOADERUNLOCKED_MANAGEDAPPFLAGGEDREASON
    // The app registration is running on an Android device on which the factory ROM has been modified.
    ANDROIDFACTORYROMMODIFIED_MANAGEDAPPFLAGGEDREASON
)

func (i ManagedAppFlaggedReason) String() string {
    return []string{"none", "rootedDevice", "androidBootloaderUnlocked", "androidFactoryRomModified"}[i]
}
func ParseManagedAppFlaggedReason(v string) (interface{}, error) {
    result := NONE_MANAGEDAPPFLAGGEDREASON
    switch v {
        case "none":
            result = NONE_MANAGEDAPPFLAGGEDREASON
        case "rootedDevice":
            result = ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON
        case "androidBootloaderUnlocked":
            result = ANDROIDBOOTLOADERUNLOCKED_MANAGEDAPPFLAGGEDREASON
        case "androidFactoryRomModified":
            result = ANDROIDFACTORYROMMODIFIED_MANAGEDAPPFLAGGEDREASON
        default:
            return 0, errors.New("Unknown ManagedAppFlaggedReason value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppFlaggedReason(values []ManagedAppFlaggedReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
