package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcOnPremisesConnectionType int

const (
    HYBRIDAZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE CloudPcOnPremisesConnectionType = iota
    AZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONTYPE
)

func (i CloudPcOnPremisesConnectionType) String() string {
    return []string{"hybridAzureADJoin", "azureADJoin", "unknownFutureValue"}[i]
}
func ParseCloudPcOnPremisesConnectionType(v string) (interface{}, error) {
    result := HYBRIDAZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
    switch v {
        case "hybridAzureADJoin":
            result = HYBRIDAZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
        case "azureADJoin":
            result = AZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown CloudPcOnPremisesConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcOnPremisesConnectionType(values []CloudPcOnPremisesConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
