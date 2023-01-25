package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type GroupPolicyOperationType int

const (
    // Group Policy invalid operation type.
    NONE_GROUPPOLICYOPERATIONTYPE GroupPolicyOperationType = iota
    // Group Policy upload operation type.
    UPLOAD_GROUPPOLICYOPERATIONTYPE
    // Group Policy upload new version operation type.
    UPLOADNEWVERSION_GROUPPOLICYOPERATIONTYPE
    // Group Policy add new language(ADML) files operation type.
    ADDLANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
    // Group Policy remove language(ADML) files operation type.
    REMOVELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
    // Group Policy update language(ADML) files operation type.
    UPDATELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
    // Group Policy remove uploaded file operation type.
    REMOVE_GROUPPOLICYOPERATIONTYPE
)

func (i GroupPolicyOperationType) String() string {
    return []string{"none", "upload", "uploadNewVersion", "addLanguageFiles", "removeLanguageFiles", "updateLanguageFiles", "remove"}[i]
}
func ParseGroupPolicyOperationType(v string) (interface{}, error) {
    result := NONE_GROUPPOLICYOPERATIONTYPE
    switch v {
        case "none":
            result = NONE_GROUPPOLICYOPERATIONTYPE
        case "upload":
            result = UPLOAD_GROUPPOLICYOPERATIONTYPE
        case "uploadNewVersion":
            result = UPLOADNEWVERSION_GROUPPOLICYOPERATIONTYPE
        case "addLanguageFiles":
            result = ADDLANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
        case "removeLanguageFiles":
            result = REMOVELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
        case "updateLanguageFiles":
            result = UPDATELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
        case "remove":
            result = REMOVE_GROUPPOLICYOPERATIONTYPE
        default:
            return 0, errors.New("Unknown GroupPolicyOperationType value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyOperationType(values []GroupPolicyOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
