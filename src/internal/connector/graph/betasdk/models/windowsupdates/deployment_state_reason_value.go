package windowsupdates
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeploymentStateReasonValue int

const (
    SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE DeploymentStateReasonValue = iota
    OFFERINGBYREQUEST_DEPLOYMENTSTATEREASONVALUE
    PAUSEDBYREQUEST_DEPLOYMENTSTATEREASONVALUE
    PAUSEDBYMONITORING_DEPLOYMENTSTATEREASONVALUE
    UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEREASONVALUE
    FAULTEDBYCONTENTOUTDATED_DEPLOYMENTSTATEREASONVALUE
)

func (i DeploymentStateReasonValue) String() string {
    return []string{"scheduledByOfferWindow", "offeringByRequest", "pausedByRequest", "pausedByMonitoring", "unknownFutureValue", "faultedByContentOutdated"}[i]
}
func ParseDeploymentStateReasonValue(v string) (interface{}, error) {
    result := SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE
    switch v {
        case "scheduledByOfferWindow":
            result = SCHEDULEDBYOFFERWINDOW_DEPLOYMENTSTATEREASONVALUE
        case "offeringByRequest":
            result = OFFERINGBYREQUEST_DEPLOYMENTSTATEREASONVALUE
        case "pausedByRequest":
            result = PAUSEDBYREQUEST_DEPLOYMENTSTATEREASONVALUE
        case "pausedByMonitoring":
            result = PAUSEDBYMONITORING_DEPLOYMENTSTATEREASONVALUE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEPLOYMENTSTATEREASONVALUE
        case "faultedByContentOutdated":
            result = FAULTEDBYCONTENTOUTDATED_DEPLOYMENTSTATEREASONVALUE
        default:
            return 0, errors.New("Unknown DeploymentStateReasonValue value: " + v)
    }
    return &result, nil
}
func SerializeDeploymentStateReasonValue(values []DeploymentStateReasonValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
