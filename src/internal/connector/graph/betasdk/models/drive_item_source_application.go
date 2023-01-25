package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DriveItemSourceApplication int

const (
    TEAMS_DRIVEITEMSOURCEAPPLICATION DriveItemSourceApplication = iota
    YAMMER_DRIVEITEMSOURCEAPPLICATION
    SHAREPOINT_DRIVEITEMSOURCEAPPLICATION
    ONEDRIVE_DRIVEITEMSOURCEAPPLICATION
    STREAM_DRIVEITEMSOURCEAPPLICATION
    POWERPOINT_DRIVEITEMSOURCEAPPLICATION
    OFFICE_DRIVEITEMSOURCEAPPLICATION
    UNKNOWNFUTUREVALUE_DRIVEITEMSOURCEAPPLICATION
)

func (i DriveItemSourceApplication) String() string {
    return []string{"teams", "yammer", "sharePoint", "oneDrive", "stream", "powerPoint", "office", "unknownFutureValue"}[i]
}
func ParseDriveItemSourceApplication(v string) (interface{}, error) {
    result := TEAMS_DRIVEITEMSOURCEAPPLICATION
    switch v {
        case "teams":
            result = TEAMS_DRIVEITEMSOURCEAPPLICATION
        case "yammer":
            result = YAMMER_DRIVEITEMSOURCEAPPLICATION
        case "sharePoint":
            result = SHAREPOINT_DRIVEITEMSOURCEAPPLICATION
        case "oneDrive":
            result = ONEDRIVE_DRIVEITEMSOURCEAPPLICATION
        case "stream":
            result = STREAM_DRIVEITEMSOURCEAPPLICATION
        case "powerPoint":
            result = POWERPOINT_DRIVEITEMSOURCEAPPLICATION
        case "office":
            result = OFFICE_DRIVEITEMSOURCEAPPLICATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DRIVEITEMSOURCEAPPLICATION
        default:
            return 0, errors.New("Unknown DriveItemSourceApplication value: " + v)
    }
    return &result, nil
}
func SerializeDriveItemSourceApplication(values []DriveItemSourceApplication) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
