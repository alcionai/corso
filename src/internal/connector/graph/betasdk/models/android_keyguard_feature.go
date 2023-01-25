package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidKeyguardFeature int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDKEYGUARDFEATURE AndroidKeyguardFeature = iota
    // Camera usage when on secure keyguard screens.
    CAMERA_ANDROIDKEYGUARDFEATURE
    // Showing notifications when on secure keyguard screens.
    NOTIFICATIONS_ANDROIDKEYGUARDFEATURE
    // Showing unredacted notifications when on secure keyguard screens.
    UNREDACTEDNOTIFICATIONS_ANDROIDKEYGUARDFEATURE
    // Trust agent state when on secure keyguard screens.
    TRUSTAGENTS_ANDROIDKEYGUARDFEATURE
    // Fingerprint sensor usage when on secure keyguard screens.
    FINGERPRINT_ANDROIDKEYGUARDFEATURE
    // Notification text entry when on secure keyguard screens.
    REMOTEINPUT_ANDROIDKEYGUARDFEATURE
    // All keyguard features when on secure keyguard screens.
    ALLFEATURES_ANDROIDKEYGUARDFEATURE
    // Face authentication on secure keyguard screens.
    FACE_ANDROIDKEYGUARDFEATURE
    // Iris authentication on secure keyguard screens.
    IRIS_ANDROIDKEYGUARDFEATURE
    // All biometric authentication on secure keyguard screens.
    BIOMETRICS_ANDROIDKEYGUARDFEATURE
)

func (i AndroidKeyguardFeature) String() string {
    return []string{"notConfigured", "camera", "notifications", "unredactedNotifications", "trustAgents", "fingerprint", "remoteInput", "allFeatures", "face", "iris", "biometrics"}[i]
}
func ParseAndroidKeyguardFeature(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDKEYGUARDFEATURE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDKEYGUARDFEATURE
        case "camera":
            result = CAMERA_ANDROIDKEYGUARDFEATURE
        case "notifications":
            result = NOTIFICATIONS_ANDROIDKEYGUARDFEATURE
        case "unredactedNotifications":
            result = UNREDACTEDNOTIFICATIONS_ANDROIDKEYGUARDFEATURE
        case "trustAgents":
            result = TRUSTAGENTS_ANDROIDKEYGUARDFEATURE
        case "fingerprint":
            result = FINGERPRINT_ANDROIDKEYGUARDFEATURE
        case "remoteInput":
            result = REMOTEINPUT_ANDROIDKEYGUARDFEATURE
        case "allFeatures":
            result = ALLFEATURES_ANDROIDKEYGUARDFEATURE
        case "face":
            result = FACE_ANDROIDKEYGUARDFEATURE
        case "iris":
            result = IRIS_ANDROIDKEYGUARDFEATURE
        case "biometrics":
            result = BIOMETRICS_ANDROIDKEYGUARDFEATURE
        default:
            return 0, errors.New("Unknown AndroidKeyguardFeature value: " + v)
    }
    return &result, nil
}
func SerializeAndroidKeyguardFeature(values []AndroidKeyguardFeature) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
