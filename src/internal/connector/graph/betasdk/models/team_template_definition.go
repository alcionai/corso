package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplateDefinition 
type TeamTemplateDefinition struct {
    Entity
    // The audience property
    audience *TeamTemplateAudience
    // The categories property
    categories []string
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The iconUrl property
    iconUrl *string
    // The languageTag property
    languageTag *string
    // The lastModifiedBy property
    lastModifiedBy IdentitySetable
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The parentTemplateId property
    parentTemplateId *string
    // The publisherName property
    publisherName *string
    // The shortDescription property
    shortDescription *string
    // The teamDefinition property
    teamDefinition Teamable
}
// NewTeamTemplateDefinition instantiates a new teamTemplateDefinition and sets the default values.
func NewTeamTemplateDefinition()(*TeamTemplateDefinition) {
    m := &TeamTemplateDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamTemplateDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplateDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplateDefinition(), nil
}
// GetAudience gets the audience property value. The audience property
func (m *TeamTemplateDefinition) GetAudience()(*TeamTemplateAudience) {
    return m.audience
}
// GetCategories gets the categories property value. The categories property
func (m *TeamTemplateDefinition) GetCategories()([]string) {
    return m.categories
}
// GetDescription gets the description property value. The description property
func (m *TeamTemplateDefinition) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *TeamTemplateDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplateDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamTemplateAudience)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(*TeamTemplateAudience))
        }
        return nil
    }
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
    res["iconUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIconUrl(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["parentTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentTemplateId(val)
        }
        return nil
    }
    res["publisherName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    res["shortDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortDescription(val)
        }
        return nil
    }
    res["teamDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamDefinition(val.(Teamable))
        }
        return nil
    }
    return res
}
// GetIconUrl gets the iconUrl property value. The iconUrl property
func (m *TeamTemplateDefinition) GetIconUrl()(*string) {
    return m.iconUrl
}
// GetLanguageTag gets the languageTag property value. The languageTag property
func (m *TeamTemplateDefinition) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *TeamTemplateDefinition) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *TeamTemplateDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetParentTemplateId gets the parentTemplateId property value. The parentTemplateId property
func (m *TeamTemplateDefinition) GetParentTemplateId()(*string) {
    return m.parentTemplateId
}
// GetPublisherName gets the publisherName property value. The publisherName property
func (m *TeamTemplateDefinition) GetPublisherName()(*string) {
    return m.publisherName
}
// GetShortDescription gets the shortDescription property value. The shortDescription property
func (m *TeamTemplateDefinition) GetShortDescription()(*string) {
    return m.shortDescription
}
// GetTeamDefinition gets the teamDefinition property value. The teamDefinition property
func (m *TeamTemplateDefinition) GetTeamDefinition()(Teamable) {
    return m.teamDefinition
}
// Serialize serializes information the current object
func (m *TeamTemplateDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAudience() != nil {
        cast := (*m.GetAudience()).String()
        err = writer.WriteStringValue("audience", &cast)
        if err != nil {
            return err
        }
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
        err = writer.WriteStringValue("iconUrl", m.GetIconUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentTemplateId", m.GetParentTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortDescription", m.GetShortDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamDefinition", m.GetTeamDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudience sets the audience property value. The audience property
func (m *TeamTemplateDefinition) SetAudience(value *TeamTemplateAudience)() {
    m.audience = value
}
// SetCategories sets the categories property value. The categories property
func (m *TeamTemplateDefinition) SetCategories(value []string)() {
    m.categories = value
}
// SetDescription sets the description property value. The description property
func (m *TeamTemplateDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *TeamTemplateDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIconUrl sets the iconUrl property value. The iconUrl property
func (m *TeamTemplateDefinition) SetIconUrl(value *string)() {
    m.iconUrl = value
}
// SetLanguageTag sets the languageTag property value. The languageTag property
func (m *TeamTemplateDefinition) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *TeamTemplateDefinition) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *TeamTemplateDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetParentTemplateId sets the parentTemplateId property value. The parentTemplateId property
func (m *TeamTemplateDefinition) SetParentTemplateId(value *string)() {
    m.parentTemplateId = value
}
// SetPublisherName sets the publisherName property value. The publisherName property
func (m *TeamTemplateDefinition) SetPublisherName(value *string)() {
    m.publisherName = value
}
// SetShortDescription sets the shortDescription property value. The shortDescription property
func (m *TeamTemplateDefinition) SetShortDescription(value *string)() {
    m.shortDescription = value
}
// SetTeamDefinition sets the teamDefinition property value. The teamDefinition property
func (m *TeamTemplateDefinition) SetTeamDefinition(value Teamable)() {
    m.teamDefinition = value
}
