package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementDomainJoinConnectorState int

const (
    // Connector is actively pinging Intune.
    ACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE DeviceManagementDomainJoinConnectorState = iota
    // There is no heart-beat from connector from last one hour.
    ERROR_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
    // There is no heart-beat from connector from last 5 days.
    INACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
)

func (i DeviceManagementDomainJoinConnectorState) String() string {
    return []string{"active", "error", "inactive"}[i]
}
func ParseDeviceManagementDomainJoinConnectorState(v string) (interface{}, error) {
    result := ACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
    switch v {
        case "active":
            result = ACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
        case "error":
            result = ERROR_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
        case "inactive":
            result = INACTIVE_DEVICEMANAGEMENTDOMAINJOINCONNECTORSTATE
        default:
            return 0, errors.New("Unknown DeviceManagementDomainJoinConnectorState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementDomainJoinConnectorState(values []DeviceManagementDomainJoinConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
