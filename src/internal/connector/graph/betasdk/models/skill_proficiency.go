package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SkillProficiency 
type SkillProficiency struct {
    ItemFacet
    // Contains categories a user has associated with the skill (for example, personal, professional, hobby).
    categories []string
    // Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
    collaborationTags []string
    // Contains a friendly name for the skill.
    displayName *string
    // Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
    proficiency *SkillProficiencyLevel
    // The thumbnailUrl property
    thumbnailUrl *string
    // Contains a link to an information source about the skill.
    webUrl *string
}
// NewSkillProficiency instantiates a new SkillProficiency and sets the default values.
func NewSkillProficiency()(*SkillProficiency) {
    m := &SkillProficiency{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.skillProficiency";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSkillProficiencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSkillProficiencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSkillProficiency(), nil
}
// GetCategories gets the categories property value. Contains categories a user has associated with the skill (for example, personal, professional, hobby).
func (m *SkillProficiency) GetCategories()([]string) {
    return m.categories
}
// GetCollaborationTags gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *SkillProficiency) GetCollaborationTags()([]string) {
    return m.collaborationTags
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the skill.
func (m *SkillProficiency) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SkillProficiency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["proficiency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSkillProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProficiency(val.(*SkillProficiencyLevel))
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
// GetProficiency gets the proficiency property value. Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
func (m *SkillProficiency) GetProficiency()(*SkillProficiencyLevel) {
    return m.proficiency
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *SkillProficiency) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. Contains a link to an information source about the skill.
func (m *SkillProficiency) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *SkillProficiency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetProficiency() != nil {
        cast := (*m.GetProficiency()).String()
        err = writer.WriteStringValue("proficiency", &cast)
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
// SetCategories sets the categories property value. Contains categories a user has associated with the skill (for example, personal, professional, hobby).
func (m *SkillProficiency) SetCategories(value []string)() {
    m.categories = value
}
// SetCollaborationTags sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *SkillProficiency) SetCollaborationTags(value []string)() {
    m.collaborationTags = value
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the skill.
func (m *SkillProficiency) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetProficiency sets the proficiency property value. Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
func (m *SkillProficiency) SetProficiency(value *SkillProficiencyLevel)() {
    m.proficiency = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *SkillProficiency) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. Contains a link to an information source about the skill.
func (m *SkillProficiency) SetWebUrl(value *string)() {
    m.webUrl = value
}
