package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceGuardLocalSystemAuthorityCredentialGuardState int

const (
    // Running
    RUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE DeviceGuardLocalSystemAuthorityCredentialGuardState = iota
    // Reboot required
    REBOOTREQUIRED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    // Not licensed for Credential Guard
    NOTLICENSED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    // Not configured
    NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    // Virtualization Based security is not running
    VIRTUALIZATIONBASEDSECURITYNOTRUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
)

func (i DeviceGuardLocalSystemAuthorityCredentialGuardState) String() string {
    return []string{"running", "rebootRequired", "notLicensed", "notConfigured", "virtualizationBasedSecurityNotRunning"}[i]
}
func ParseDeviceGuardLocalSystemAuthorityCredentialGuardState(v string) (interface{}, error) {
    result := RUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    switch v {
        case "running":
            result = RUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "rebootRequired":
            result = REBOOTREQUIRED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "notLicensed":
            result = NOTLICENSED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "notConfigured":
            result = NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "virtualizationBasedSecurityNotRunning":
            result = VIRTUALIZATIONBASEDSECURITYNOTRUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        default:
            return 0, errors.New("Unknown DeviceGuardLocalSystemAuthorityCredentialGuardState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceGuardLocalSystemAuthorityCredentialGuardState(values []DeviceGuardLocalSystemAuthorityCredentialGuardState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
