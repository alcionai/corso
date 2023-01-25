package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ZebraFotaDeploymentState int

const (
    // Deployment is created but Zebra has not confirmed its creation.
    PENDINGCREATION_ZEBRAFOTADEPLOYMENTSTATE ZebraFotaDeploymentState = iota
    // Deployment was not successfully created with Zebra.
    CREATEFAILED_ZEBRAFOTADEPLOYMENTSTATE
    // Deployment has been created but has not been deployed yet.
    CREATED_ZEBRAFOTADEPLOYMENTSTATE
    // Deployment has started but not completed.
    INPROGRESS_ZEBRAFOTADEPLOYMENTSTATE
    // Deployment has completed or end date has passed.
    COMPLETED_ZEBRAFOTADEPLOYMENTSTATE
    // Admin has requested to cancel a deployment but Zebra has not confirmed cancellation.
    PENDINGCANCEL_ZEBRAFOTADEPLOYMENTSTATE
    // Deployment has been successfully canceled by Zebra.
    CANCELED_ZEBRAFOTADEPLOYMENTSTATE
    // Unknown future enum value.
    UNKNOWNFUTUREVALUE_ZEBRAFOTADEPLOYMENTSTATE
)

func (i ZebraFotaDeploymentState) String() string {
    return []string{"pendingCreation", "createFailed", "created", "inProgress", "completed", "pendingCancel", "canceled", "unknownFutureValue"}[i]
}
func ParseZebraFotaDeploymentState(v string) (interface{}, error) {
    result := PENDINGCREATION_ZEBRAFOTADEPLOYMENTSTATE
    switch v {
        case "pendingCreation":
            result = PENDINGCREATION_ZEBRAFOTADEPLOYMENTSTATE
        case "createFailed":
            result = CREATEFAILED_ZEBRAFOTADEPLOYMENTSTATE
        case "created":
            result = CREATED_ZEBRAFOTADEPLOYMENTSTATE
        case "inProgress":
            result = INPROGRESS_ZEBRAFOTADEPLOYMENTSTATE
        case "completed":
            result = COMPLETED_ZEBRAFOTADEPLOYMENTSTATE
        case "pendingCancel":
            result = PENDINGCANCEL_ZEBRAFOTADEPLOYMENTSTATE
        case "canceled":
            result = CANCELED_ZEBRAFOTADEPLOYMENTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTADEPLOYMENTSTATE
        default:
            return 0, errors.New("Unknown ZebraFotaDeploymentState value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaDeploymentState(values []ZebraFotaDeploymentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
