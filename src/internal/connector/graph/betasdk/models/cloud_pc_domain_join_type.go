package models
import (
    "errors"
)
// Provides operations to call the add method.
type CloudPcDomainJoinType int

const (
    AZUREADJOIN_CLOUDPCDOMAINJOINTYPE CloudPcDomainJoinType = iota
    HYBRIDAZUREADJOIN_CLOUDPCDOMAINJOINTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCDOMAINJOINTYPE
)

func (i CloudPcDomainJoinType) String() string {
    return []string{"azureADJoin", "hybridAzureADJoin", "unknownFutureValue"}[i]
}
func ParseCloudPcDomainJoinType(v string) (interface{}, error) {
    result := AZUREADJOIN_CLOUDPCDOMAINJOINTYPE
    switch v {
        case "azureADJoin":
            result = AZUREADJOIN_CLOUDPCDOMAINJOINTYPE
        case "hybridAzureADJoin":
            result = HYBRIDAZUREADJOIN_CLOUDPCDOMAINJOINTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDOMAINJOINTYPE
        default:
            return 0, errors.New("Unknown CloudPcDomainJoinType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDomainJoinType(values []CloudPcDomainJoinType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
