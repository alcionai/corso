package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ReferenceAttachmentPermission int

const (
    OTHER_REFERENCEATTACHMENTPERMISSION ReferenceAttachmentPermission = iota
    VIEW_REFERENCEATTACHMENTPERMISSION
    EDIT_REFERENCEATTACHMENTPERMISSION
    ANONYMOUSVIEW_REFERENCEATTACHMENTPERMISSION
    ANONYMOUSEDIT_REFERENCEATTACHMENTPERMISSION
    ORGANIZATIONVIEW_REFERENCEATTACHMENTPERMISSION
    ORGANIZATIONEDIT_REFERENCEATTACHMENTPERMISSION
)

func (i ReferenceAttachmentPermission) String() string {
    return []string{"other", "view", "edit", "anonymousView", "anonymousEdit", "organizationView", "organizationEdit"}[i]
}
func ParseReferenceAttachmentPermission(v string) (interface{}, error) {
    result := OTHER_REFERENCEATTACHMENTPERMISSION
    switch v {
        case "other":
            result = OTHER_REFERENCEATTACHMENTPERMISSION
        case "view":
            result = VIEW_REFERENCEATTACHMENTPERMISSION
        case "edit":
            result = EDIT_REFERENCEATTACHMENTPERMISSION
        case "anonymousView":
            result = ANONYMOUSVIEW_REFERENCEATTACHMENTPERMISSION
        case "anonymousEdit":
            result = ANONYMOUSEDIT_REFERENCEATTACHMENTPERMISSION
        case "organizationView":
            result = ORGANIZATIONVIEW_REFERENCEATTACHMENTPERMISSION
        case "organizationEdit":
            result = ORGANIZATIONEDIT_REFERENCEATTACHMENTPERMISSION
        default:
            return 0, errors.New("Unknown ReferenceAttachmentPermission value: " + v)
    }
    return &result, nil
}
func SerializeReferenceAttachmentPermission(values []ReferenceAttachmentPermission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
