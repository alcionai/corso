package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ConfigurationManagerActionDeliveryStatus int

const (
    // Pending to deliver the action to ConfigurationManager
    UNKNOWN_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS ConfigurationManagerActionDeliveryStatus = iota
    // Pending to deliver the action to ConfigurationManager
    PENDINGDELIVERY_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
    // Action is sent to ConfigurationManager Connector service (cloud)
    DELIVEREDTOCONNECTORSERVICE_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
    // Failed to send the action to ConfigurationManager Connector service (cloud)
    FAILEDTODELIVERTOCONNECTORSERVICE_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
    // Action is delivered to ConfigurationManager on-prem server
    DELIVEREDTOONPREMISESSERVER_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
)

func (i ConfigurationManagerActionDeliveryStatus) String() string {
    return []string{"unknown", "pendingDelivery", "deliveredToConnectorService", "failedToDeliverToConnectorService", "deliveredToOnPremisesServer"}[i]
}
func ParseConfigurationManagerActionDeliveryStatus(v string) (interface{}, error) {
    result := UNKNOWN_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
        case "pendingDelivery":
            result = PENDINGDELIVERY_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
        case "deliveredToConnectorService":
            result = DELIVEREDTOCONNECTORSERVICE_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
        case "failedToDeliverToConnectorService":
            result = FAILEDTODELIVERTOCONNECTORSERVICE_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
        case "deliveredToOnPremisesServer":
            result = DELIVEREDTOONPREMISESSERVER_CONFIGURATIONMANAGERACTIONDELIVERYSTATUS
        default:
            return 0, errors.New("Unknown ConfigurationManagerActionDeliveryStatus value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationManagerActionDeliveryStatus(values []ConfigurationManagerActionDeliveryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
