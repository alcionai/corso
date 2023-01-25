package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonWebsite 
type PersonWebsite struct {
    ItemFacet
    // Contains categories a user has associated with the website (for example, personal, recipes).
    categories []string
    // Contains a description of the website.
    description *string
    // Contains a friendly name for the website.
    displayName *string
    // The thumbnailUrl property
    thumbnailUrl *string
    // Contains a link to the website itself.
    webUrl *string
}
// NewPersonWebsite instantiates a new PersonWebsite and sets the default values.
func NewPersonWebsite()(*PersonWebsite) {
    m := &PersonWebsite{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personWebsite";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonWebsiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonWebsiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonWebsite(), nil
}
// GetCategories gets the categories property value. Contains categories a user has associated with the website (for example, personal, recipes).
func (m *PersonWebsite) GetCategories()([]string) {
    return m.categories
}
// GetDescription gets the description property value. Contains a description of the website.
func (m *PersonWebsite) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the website.
func (m *PersonWebsite) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonWebsite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
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
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *PersonWebsite) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. Contains a link to the website itself.
func (m *PersonWebsite) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *PersonWebsite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
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
// SetCategories sets the categories property value. Contains categories a user has associated with the website (for example, personal, recipes).
func (m *PersonWebsite) SetCategories(value []string)() {
    m.categories = value
}
// SetDescription sets the description property value. Contains a description of the website.
func (m *PersonWebsite) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the website.
func (m *PersonWebsite) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *PersonWebsite) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. Contains a link to the website itself.
func (m *PersonWebsite) SetWebUrl(value *string)() {
    m.webUrl = value
}
