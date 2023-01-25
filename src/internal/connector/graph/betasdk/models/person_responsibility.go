package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonResponsibility 
type PersonResponsibility struct {
    ItemFacet
    // Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
    collaborationTags []string
    // Description of the responsibility.
    description *string
    // Contains a friendly name for the responsibility.
    displayName *string
    // The thumbnailUrl property
    thumbnailUrl *string
    // Contains a link to a web page or resource about the responsibility.
    webUrl *string
}
// NewPersonResponsibility instantiates a new PersonResponsibility and sets the default values.
func NewPersonResponsibility()(*PersonResponsibility) {
    m := &PersonResponsibility{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personResponsibility";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonResponsibilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonResponsibilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonResponsibility(), nil
}
// GetCollaborationTags gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *PersonResponsibility) GetCollaborationTags()([]string) {
    return m.collaborationTags
}
// GetDescription gets the description property value. Description of the responsibility.
func (m *PersonResponsibility) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the responsibility.
func (m *PersonResponsibility) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonResponsibility) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
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
func (m *PersonResponsibility) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. Contains a link to a web page or resource about the responsibility.
func (m *PersonResponsibility) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *PersonResponsibility) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
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
// SetCollaborationTags sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *PersonResponsibility) SetCollaborationTags(value []string)() {
    m.collaborationTags = value
}
// SetDescription sets the description property value. Description of the responsibility.
func (m *PersonResponsibility) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the responsibility.
func (m *PersonResponsibility) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *PersonResponsibility) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. Contains a link to a web page or resource about the responsibility.
func (m *PersonResponsibility) SetWebUrl(value *string)() {
    m.webUrl = value
}
