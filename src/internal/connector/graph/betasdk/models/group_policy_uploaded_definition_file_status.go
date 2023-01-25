package models
import (
    "errors"
)
// Provides operations to call the add method.
type GroupPolicyUploadedDefinitionFileStatus int

const (
    // Group Policy uploaded definition file invalid upload status.
    NONE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS GroupPolicyUploadedDefinitionFileStatus = iota
    // Group Policy uploaded definition file upload in progress.
    UPLOADINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    // Group Policy uploaded definition file available.
    AVAILABLE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    // Group Policy uploaded definition file assigned to policy.
    ASSIGNED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    // Group Policy uploaded definition file removal in progress.
    REMOVALINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    // Group Policy uploaded definition file upload failed.
    UPLOADFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    // Group Policy uploaded definition file removal failed.
    REMOVALFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
)

func (i GroupPolicyUploadedDefinitionFileStatus) String() string {
    return []string{"none", "uploadInProgress", "available", "assigned", "removalInProgress", "uploadFailed", "removalFailed"}[i]
}
func ParseGroupPolicyUploadedDefinitionFileStatus(v string) (interface{}, error) {
    result := NONE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    switch v {
        case "none":
            result = NONE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "uploadInProgress":
            result = UPLOADINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "available":
            result = AVAILABLE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "assigned":
            result = ASSIGNED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "removalInProgress":
            result = REMOVALINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "uploadFailed":
            result = UPLOADFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "removalFailed":
            result = REMOVALFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        default:
            return 0, errors.New("Unknown GroupPolicyUploadedDefinitionFileStatus value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyUploadedDefinitionFileStatus(values []GroupPolicyUploadedDefinitionFileStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
