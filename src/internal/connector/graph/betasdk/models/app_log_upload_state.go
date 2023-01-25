package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AppLogUploadState int

const (
    // Request is waiting to be processed or under processing
    PENDING_APPLOGUPLOADSTATE AppLogUploadState = iota
    // Request is completed with file uploaded to Azure blob for download.
    COMPLETED_APPLOGUPLOADSTATE
    // Request finished processing and in error state.
    FAILED_APPLOGUPLOADSTATE
)

func (i AppLogUploadState) String() string {
    return []string{"pending", "completed", "failed"}[i]
}
func ParseAppLogUploadState(v string) (interface{}, error) {
    result := PENDING_APPLOGUPLOADSTATE
    switch v {
        case "pending":
            result = PENDING_APPLOGUPLOADSTATE
        case "completed":
            result = COMPLETED_APPLOGUPLOADSTATE
        case "failed":
            result = FAILED_APPLOGUPLOADSTATE
        default:
            return 0, errors.New("Unknown AppLogUploadState value: " + v)
    }
    return &result, nil
}
func SerializeAppLogUploadState(values []AppLogUploadState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
