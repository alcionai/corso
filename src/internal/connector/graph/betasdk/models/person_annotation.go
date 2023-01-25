package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonAnnotation 
type PersonAnnotation struct {
    ItemFacet
    // Contains the detail of the note itself.
    detail ItemBodyable
    // Contains a friendly name for the note.
    displayName *string
    // The thumbnailUrl property
    thumbnailUrl *string
}
// NewPersonAnnotation instantiates a new PersonAnnotation and sets the default values.
func NewPersonAnnotation()(*PersonAnnotation) {
    m := &PersonAnnotation{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personAnnotation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonAnnotationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonAnnotationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonAnnotation(), nil
}
// GetDetail gets the detail property value. Contains the detail of the note itself.
func (m *PersonAnnotation) GetDetail()(ItemBodyable) {
    return m.detail
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the note.
func (m *PersonAnnotation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonAnnotation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(ItemBodyable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *PersonAnnotation) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// Serialize serializes information the current object
func (m *PersonAnnotation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
// SetDetail sets the detail property value. Contains the detail of the note itself.
func (m *PersonAnnotation) SetDetail(value ItemBodyable)() {
    m.detail = value
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the note.
func (m *PersonAnnotation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *PersonAnnotation) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
