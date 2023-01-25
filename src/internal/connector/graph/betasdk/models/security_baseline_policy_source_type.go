package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SecurityBaselinePolicySourceType int

const (
    DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE SecurityBaselinePolicySourceType = iota
    DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE
)

func (i SecurityBaselinePolicySourceType) String() string {
    return []string{"deviceConfiguration", "deviceIntent"}[i]
}
func ParseSecurityBaselinePolicySourceType(v string) (interface{}, error) {
    result := DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE
    switch v {
        case "deviceConfiguration":
            result = DEVICECONFIGURATION_SECURITYBASELINEPOLICYSOURCETYPE
        case "deviceIntent":
            result = DEVICEINTENT_SECURITYBASELINEPOLICYSOURCETYPE
        default:
            return 0, errors.New("Unknown SecurityBaselinePolicySourceType value: " + v)
    }
    return &result, nil
}
func SerializeSecurityBaselinePolicySourceType(values []SecurityBaselinePolicySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
