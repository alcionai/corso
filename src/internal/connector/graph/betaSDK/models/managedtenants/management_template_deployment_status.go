package managedtenants
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagementTemplateDeploymentStatus int

const (
    UNKNOWN_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS ManagementTemplateDeploymentStatus = iota
    INPROGRESS_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    COMPLETED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    FAILED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    INELIGIBLE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    UNKNOWNFUTUREVALUE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
)

func (i ManagementTemplateDeploymentStatus) String() string {
    return []string{"unknown", "inProgress", "completed", "failed", "ineligible", "unknownFutureValue"}[i]
}
func ParseManagementTemplateDeploymentStatus(v string) (interface{}, error) {
    result := UNKNOWN_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "inProgress":
            result = INPROGRESS_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "completed":
            result = COMPLETED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "failed":
            result = FAILED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "ineligible":
            result = INELIGIBLE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        default:
            return 0, errors.New("Unknown ManagementTemplateDeploymentStatus value: " + v)
    }
    return &result, nil
}
func SerializeManagementTemplateDeploymentStatus(values []ManagementTemplateDeploymentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
