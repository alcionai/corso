package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferenceAttachment 
type ReferenceAttachment struct {
    Attachment
    // Specifies whether the attachment is a link to a folder. Must set this to true if sourceUrl is a link to a folder. Optional.
    isFolder *bool
    // Specifies the permissions granted for the attachment by the type of provider in providerType. Possible values are: other, view, edit, anonymousView, anonymousEdit, organizationView, organizationEdit. Optional.
    permission *ReferenceAttachmentPermission
    // Applies to only a reference attachment of an image - URL to get a preview image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
    previewUrl *string
    // The type of provider that supports an attachment of this contentType. Possible values are: other, oneDriveBusiness, oneDriveConsumer, dropbox. Optional.
    providerType *ReferenceAttachmentProvider
    // URL to get the attachment content. If this is a URL to a folder, then for the folder to be displayed correctly in Outlook or Outlook on the web, set isFolder to true. Required.
    sourceUrl *string
    // Applies to only a reference attachment of an image - URL to get a thumbnail image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
    thumbnailUrl *string
}
// NewReferenceAttachment instantiates a new ReferenceAttachment and sets the default values.
func NewReferenceAttachment()(*ReferenceAttachment) {
    m := &ReferenceAttachment{
        Attachment: *NewAttachment(),
    }
    odataTypeValue := "#microsoft.graph.referenceAttachment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateReferenceAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReferenceAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferenceAttachment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReferenceAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Attachment.GetFieldDeserializers()
    res["isFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFolder(val)
        }
        return nil
    }
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReferenceAttachmentPermission)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val.(*ReferenceAttachmentPermission))
        }
        return nil
    }
    res["previewUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewUrl(val)
        }
        return nil
    }
    res["providerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReferenceAttachmentProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderType(val.(*ReferenceAttachmentProvider))
        }
        return nil
    }
    res["sourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceUrl(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    return res
}
// GetIsFolder gets the isFolder property value. Specifies whether the attachment is a link to a folder. Must set this to true if sourceUrl is a link to a folder. Optional.
func (m *ReferenceAttachment) GetIsFolder()(*bool) {
    return m.isFolder
}
// GetPermission gets the permission property value. Specifies the permissions granted for the attachment by the type of provider in providerType. Possible values are: other, view, edit, anonymousView, anonymousEdit, organizationView, organizationEdit. Optional.
func (m *ReferenceAttachment) GetPermission()(*ReferenceAttachmentPermission) {
    return m.permission
}
// GetPreviewUrl gets the previewUrl property value. Applies to only a reference attachment of an image - URL to get a preview image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) GetPreviewUrl()(*string) {
    return m.previewUrl
}
// GetProviderType gets the providerType property value. The type of provider that supports an attachment of this contentType. Possible values are: other, oneDriveBusiness, oneDriveConsumer, dropbox. Optional.
func (m *ReferenceAttachment) GetProviderType()(*ReferenceAttachmentProvider) {
    return m.providerType
}
// GetSourceUrl gets the sourceUrl property value. URL to get the attachment content. If this is a URL to a folder, then for the folder to be displayed correctly in Outlook or Outlook on the web, set isFolder to true. Required.
func (m *ReferenceAttachment) GetSourceUrl()(*string) {
    return m.sourceUrl
}
// GetThumbnailUrl gets the thumbnailUrl property value. Applies to only a reference attachment of an image - URL to get a thumbnail image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// Serialize serializes information the current object
func (m *ReferenceAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Attachment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isFolder", m.GetIsFolder())
        if err != nil {
            return err
        }
    }
    if m.GetPermission() != nil {
        cast := (*m.GetPermission()).String()
        err = writer.WriteStringValue("permission", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("previewUrl", m.GetPreviewUrl())
        if err != nil {
            return err
        }
    }
    if m.GetProviderType() != nil {
        cast := (*m.GetProviderType()).String()
        err = writer.WriteStringValue("providerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceUrl", m.GetSourceUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsFolder sets the isFolder property value. Specifies whether the attachment is a link to a folder. Must set this to true if sourceUrl is a link to a folder. Optional.
func (m *ReferenceAttachment) SetIsFolder(value *bool)() {
    m.isFolder = value
}
// SetPermission sets the permission property value. Specifies the permissions granted for the attachment by the type of provider in providerType. Possible values are: other, view, edit, anonymousView, anonymousEdit, organizationView, organizationEdit. Optional.
func (m *ReferenceAttachment) SetPermission(value *ReferenceAttachmentPermission)() {
    m.permission = value
}
// SetPreviewUrl sets the previewUrl property value. Applies to only a reference attachment of an image - URL to get a preview image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) SetPreviewUrl(value *string)() {
    m.previewUrl = value
}
// SetProviderType sets the providerType property value. The type of provider that supports an attachment of this contentType. Possible values are: other, oneDriveBusiness, oneDriveConsumer, dropbox. Optional.
func (m *ReferenceAttachment) SetProviderType(value *ReferenceAttachmentProvider)() {
    m.providerType = value
}
// SetSourceUrl sets the sourceUrl property value. URL to get the attachment content. If this is a URL to a folder, then for the folder to be displayed correctly in Outlook or Outlook on the web, set isFolder to true. Required.
func (m *ReferenceAttachment) SetSourceUrl(value *string)() {
    m.sourceUrl = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. Applies to only a reference attachment of an image - URL to get a thumbnail image. Use thumbnailUrl and previewUrl only when sourceUrl identifies an image file. Optional.
func (m *ReferenceAttachment) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
