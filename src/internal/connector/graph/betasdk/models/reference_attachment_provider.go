package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ReferenceAttachmentProvider int

const (
    OTHER_REFERENCEATTACHMENTPROVIDER ReferenceAttachmentProvider = iota
    ONEDRIVEBUSINESS_REFERENCEATTACHMENTPROVIDER
    ONEDRIVECONSUMER_REFERENCEATTACHMENTPROVIDER
    DROPBOX_REFERENCEATTACHMENTPROVIDER
)

func (i ReferenceAttachmentProvider) String() string {
    return []string{"other", "oneDriveBusiness", "oneDriveConsumer", "dropbox"}[i]
}
func ParseReferenceAttachmentProvider(v string) (interface{}, error) {
    result := OTHER_REFERENCEATTACHMENTPROVIDER
    switch v {
        case "other":
            result = OTHER_REFERENCEATTACHMENTPROVIDER
        case "oneDriveBusiness":
            result = ONEDRIVEBUSINESS_REFERENCEATTACHMENTPROVIDER
        case "oneDriveConsumer":
            result = ONEDRIVECONSUMER_REFERENCEATTACHMENTPROVIDER
        case "dropbox":
            result = DROPBOX_REFERENCEATTACHMENTPROVIDER
        default:
            return 0, errors.New("Unknown ReferenceAttachmentProvider value: " + v)
    }
    return &result, nil
}
func SerializeReferenceAttachmentProvider(values []ReferenceAttachmentProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
