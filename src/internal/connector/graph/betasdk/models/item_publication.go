package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPublication 
type ItemPublication struct {
    ItemFacet
    // Description of the publication.
    description *string
    // Title of the publication.
    displayName *string
    // The date that the publication was published.
    publishedDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Publication or publisher for the publication.
    publisher *string
    // URL referencing a thumbnail of the publication.
    thumbnailUrl *string
    // URL referencing the publication.
    webUrl *string
}
// NewItemPublication instantiates a new ItemPublication and sets the default values.
func NewItemPublication()(*ItemPublication) {
    m := &ItemPublication{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.itemPublication";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateItemPublicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPublicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPublication(), nil
}
// GetDescription gets the description property value. Description of the publication.
func (m *ItemPublication) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Title of the publication.
func (m *ItemPublication) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPublication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["publishedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDate(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
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
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetPublishedDate gets the publishedDate property value. The date that the publication was published.
func (m *ItemPublication) GetPublishedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.publishedDate
}
// GetPublisher gets the publisher property value. Publication or publisher for the publication.
func (m *ItemPublication) GetPublisher()(*string) {
    return m.publisher
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL referencing a thumbnail of the publication.
func (m *ItemPublication) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. URL referencing the publication.
func (m *ItemPublication) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *ItemPublication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteDateOnlyValue("publishedDate", m.GetPublishedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
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
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the publication.
func (m *ItemPublication) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Title of the publication.
func (m *ItemPublication) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPublishedDate sets the publishedDate property value. The date that the publication was published.
func (m *ItemPublication) SetPublishedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.publishedDate = value
}
// SetPublisher sets the publisher property value. Publication or publisher for the publication.
func (m *ItemPublication) SetPublisher(value *string)() {
    m.publisher = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL referencing a thumbnail of the publication.
func (m *ItemPublication) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. URL referencing the publication.
func (m *ItemPublication) SetWebUrl(value *string)() {
    m.webUrl = value
}
