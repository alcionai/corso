package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcPartnerAgentName int

const (
    CITRIX_CLOUDPCPARTNERAGENTNAME CloudPcPartnerAgentName = iota
    UNKNOWNFUTUREVALUE_CLOUDPCPARTNERAGENTNAME
)

func (i CloudPcPartnerAgentName) String() string {
    return []string{"citrix", "unknownFutureValue"}[i]
}
func ParseCloudPcPartnerAgentName(v string) (interface{}, error) {
    result := CITRIX_CLOUDPCPARTNERAGENTNAME
    switch v {
        case "citrix":
            result = CITRIX_CLOUDPCPARTNERAGENTNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPARTNERAGENTNAME
        default:
            return 0, errors.New("Unknown CloudPcPartnerAgentName value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcPartnerAgentName(values []CloudPcPartnerAgentName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
