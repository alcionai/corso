package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonInterest 
type PersonInterest struct {
    ItemFacet
    // Contains categories a user has associated with the interest (for example, personal, recipies).
    categories []string
    // Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
    collaborationTags []string
    // Contains a description of the interest.
    description *string
    // Contains a friendly name for the interest.
    displayName *string
    // The thumbnailUrl property
    thumbnailUrl *string
    // Contains a link to a web page or resource about the interest.
    webUrl *string
}
// NewPersonInterest instantiates a new PersonInterest and sets the default values.
func NewPersonInterest()(*PersonInterest) {
    m := &PersonInterest{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personInterest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonInterestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonInterestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonInterest(), nil
}
// GetCategories gets the categories property value. Contains categories a user has associated with the interest (for example, personal, recipies).
func (m *PersonInterest) GetCategories()([]string) {
    return m.categories
}
// GetCollaborationTags gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *PersonInterest) GetCollaborationTags()([]string) {
    return m.collaborationTags
}
// GetDescription gets the description property value. Contains a description of the interest.
func (m *PersonInterest) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the interest.
func (m *PersonInterest) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonInterest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["collaborationTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCollaborationTags(res)
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
func (m *PersonInterest) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. Contains a link to a web page or resource about the interest.
func (m *PersonInterest) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *PersonInterest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetCollaborationTags() != nil {
        err = writer.WriteCollectionOfStringValues("collaborationTags", m.GetCollaborationTags())
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
// SetCategories sets the categories property value. Contains categories a user has associated with the interest (for example, personal, recipies).
func (m *PersonInterest) SetCategories(value []string)() {
    m.categories = value
}
// SetCollaborationTags sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *PersonInterest) SetCollaborationTags(value []string)() {
    m.collaborationTags = value
}
// SetDescription sets the description property value. Contains a description of the interest.
func (m *PersonInterest) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the interest.
func (m *PersonInterest) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *PersonInterest) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. Contains a link to a web page or resource about the interest.
func (m *PersonInterest) SetWebUrl(value *string)() {
    m.webUrl = value
}
