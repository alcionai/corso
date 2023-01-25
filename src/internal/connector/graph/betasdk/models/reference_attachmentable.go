package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferenceAttachmentable 
type ReferenceAttachmentable interface {
    Attachmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsFolder()(*bool)
    GetPermission()(*ReferenceAttachmentPermission)
    GetPreviewUrl()(*string)
    GetProviderType()(*ReferenceAttachmentProvider)
    GetSourceUrl()(*string)
    GetThumbnailUrl()(*string)
    SetIsFolder(value *bool)()
    SetPermission(value *ReferenceAttachmentPermission)()
    SetPreviewUrl(value *string)()
    SetProviderType(value *ReferenceAttachmentProvider)()
    SetSourceUrl(value *string)()
    SetThumbnailUrl(value *string)()
}
