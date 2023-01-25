package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ConfigurationManagerClientState int

const (
    // Configuration manager agent is older than 1806 or not installed or this device has not checked into Intune for over 30 days.
    UNKNOWN_CONFIGURATIONMANAGERCLIENTSTATE ConfigurationManagerClientState = iota
    // The configuration manager agent is installed but may not be showing up in the configuration manager console yet. Wait a few hours for it to refresh.
    INSTALLED_CONFIGURATIONMANAGERCLIENTSTATE
    // This device was able to check in with the configuration manager service successfully.
    HEALTHY_CONFIGURATIONMANAGERCLIENTSTATE
    // The configuration manager agent failed to install.
    INSTALLFAILED_CONFIGURATIONMANAGERCLIENTSTATE
    // The update from version x to version y of the configuration manager agent failed. 
    UPDATEFAILED_CONFIGURATIONMANAGERCLIENTSTATE
    // The configuration manager agent was able to reach the configuration manager service in the past but is now no longer able to. 
    COMMUNICATIONERROR_CONFIGURATIONMANAGERCLIENTSTATE
)

func (i ConfigurationManagerClientState) String() string {
    return []string{"unknown", "installed", "healthy", "installFailed", "updateFailed", "communicationError"}[i]
}
func ParseConfigurationManagerClientState(v string) (interface{}, error) {
    result := UNKNOWN_CONFIGURATIONMANAGERCLIENTSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_CONFIGURATIONMANAGERCLIENTSTATE
        case "installed":
            result = INSTALLED_CONFIGURATIONMANAGERCLIENTSTATE
        case "healthy":
            result = HEALTHY_CONFIGURATIONMANAGERCLIENTSTATE
        case "installFailed":
            result = INSTALLFAILED_CONFIGURATIONMANAGERCLIENTSTATE
        case "updateFailed":
            result = UPDATEFAILED_CONFIGURATIONMANAGERCLIENTSTATE
        case "communicationError":
            result = COMMUNICATIONERROR_CONFIGURATIONMANAGERCLIENTSTATE
        default:
            return 0, errors.New("Unknown ConfigurationManagerClientState value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationManagerClientState(values []ConfigurationManagerClientState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
