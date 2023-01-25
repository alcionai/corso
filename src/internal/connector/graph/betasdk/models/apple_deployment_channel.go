package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AppleDeploymentChannel int

const (
    // Send payload down over Device Channel.
    DEVICECHANNEL_APPLEDEPLOYMENTCHANNEL AppleDeploymentChannel = iota
    // Send payload down over User Channel.
    USERCHANNEL_APPLEDEPLOYMENTCHANNEL
)

func (i AppleDeploymentChannel) String() string {
    return []string{"deviceChannel", "userChannel"}[i]
}
func ParseAppleDeploymentChannel(v string) (interface{}, error) {
    result := DEVICECHANNEL_APPLEDEPLOYMENTCHANNEL
    switch v {
        case "deviceChannel":
            result = DEVICECHANNEL_APPLEDEPLOYMENTCHANNEL
        case "userChannel":
            result = USERCHANNEL_APPLEDEPLOYMENTCHANNEL
        default:
            return 0, errors.New("Unknown AppleDeploymentChannel value: " + v)
    }
    return &result, nil
}
func SerializeAppleDeploymentChannel(values []AppleDeploymentChannel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
