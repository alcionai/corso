package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcExportJobStatus int

const (
    NOTSTARTED_CLOUDPCEXPORTJOBSTATUS CloudPcExportJobStatus = iota
    INPROGRESS_CLOUDPCEXPORTJOBSTATUS
    COMPLETED_CLOUDPCEXPORTJOBSTATUS
    FAILED_CLOUDPCEXPORTJOBSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCEXPORTJOBSTATUS
)

func (i CloudPcExportJobStatus) String() string {
    return []string{"notStarted", "inProgress", "completed", "failed", "unknownFutureValue"}[i]
}
func ParseCloudPcExportJobStatus(v string) (interface{}, error) {
    result := NOTSTARTED_CLOUDPCEXPORTJOBSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_CLOUDPCEXPORTJOBSTATUS
        case "inProgress":
            result = INPROGRESS_CLOUDPCEXPORTJOBSTATUS
        case "completed":
            result = COMPLETED_CLOUDPCEXPORTJOBSTATUS
        case "failed":
            result = FAILED_CLOUDPCEXPORTJOBSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCEXPORTJOBSTATUS
        default:
            return 0, errors.New("Unknown CloudPcExportJobStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcExportJobStatus(values []CloudPcExportJobStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
