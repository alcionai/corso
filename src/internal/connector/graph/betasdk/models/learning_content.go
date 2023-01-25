package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LearningContent provides operations to call the add method.
type LearningContent struct {
    Entity
    // Keywords, topics, and other tags associated with the learning content. Optional.
    additionalTags []string
    // The content web URL for the learning content. Required.
    contentWebUrl *string
    // The authors, creators, or contributors of the learning content. Optional.
    contributors []string
    // The date when the learning content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description or summary for the learning content. Optional.
    description *string
    // The duration of the learning content in seconds. The value is represented in ISO 8601 format for durations. Optional.
    duration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Unique external content ID for the learning content. Required.
    externalId *string
    // The format of the learning content. For example, Course, Video, Book, Book Summary, Audiobook Summary. Optional.
    format *string
    // Indicates whether the content is active or not. Inactive content will not show up in the UI. The default value is true. Optional.
    isActive *bool
    // Indicates whether the learning content requires the user to sign-in on the learning provider platform or not. The default value is false. Optional.
    isPremium *bool
    // Indicates whether the learning content is searchable or not. The default value is true. Optional.
    isSearchable *bool
    // The language of the learning content, for example, en-us or fr-fr. Required.
    languageTag *string
    // The date when the learning content was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of pages of the learning content, for example, 9. Optional.
    numberOfPages *int32
    // The skills tags associated with the learning content. Optional.
    skillTags []string
    // The source name of the learning content, such as LinkedIn Learning or Coursera. Optional.
    sourceName *string
    // The URL of learning content thumbnail image. Optional.
    thumbnailWebUrl *string
    // The title of the learning content. Required.
    title *string
}
// NewLearningContent instantiates a new learningContent and sets the default values.
func NewLearningContent()(*LearningContent) {
    m := &LearningContent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLearningContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLearningContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLearningContent(), nil
}
// GetAdditionalTags gets the additionalTags property value. Keywords, topics, and other tags associated with the learning content. Optional.
func (m *LearningContent) GetAdditionalTags()([]string) {
    return m.additionalTags
}
// GetContentWebUrl gets the contentWebUrl property value. The content web URL for the learning content. Required.
func (m *LearningContent) GetContentWebUrl()(*string) {
    return m.contentWebUrl
}
// GetContributors gets the contributors property value. The authors, creators, or contributors of the learning content. Optional.
func (m *LearningContent) GetContributors()([]string) {
    return m.contributors
}
// GetCreatedDateTime gets the createdDateTime property value. The date when the learning content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description or summary for the learning content. Optional.
func (m *LearningContent) GetDescription()(*string) {
    return m.description
}
// GetDuration gets the duration property value. The duration of the learning content in seconds. The value is represented in ISO 8601 format for durations. Optional.
func (m *LearningContent) GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.duration
}
// GetExternalId gets the externalId property value. Unique external content ID for the learning content. Required.
func (m *LearningContent) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LearningContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAdditionalTags(res)
        }
        return nil
    }
    res["contentWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentWebUrl(val)
        }
        return nil
    }
    res["contributors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContributors(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["isPremium"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPremium(val)
        }
        return nil
    }
    res["isSearchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
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
    res["numberOfPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfPages(val)
        }
        return nil
    }
    res["skillTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSkillTags(res)
        }
        return nil
    }
    res["sourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceName(val)
        }
        return nil
    }
    res["thumbnailWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailWebUrl(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. The format of the learning content. For example, Course, Video, Book, Book Summary, Audiobook Summary. Optional.
func (m *LearningContent) GetFormat()(*string) {
    return m.format
}
// GetIsActive gets the isActive property value. Indicates whether the content is active or not. Inactive content will not show up in the UI. The default value is true. Optional.
func (m *LearningContent) GetIsActive()(*bool) {
    return m.isActive
}
// GetIsPremium gets the isPremium property value. Indicates whether the learning content requires the user to sign-in on the learning provider platform or not. The default value is false. Optional.
func (m *LearningContent) GetIsPremium()(*bool) {
    return m.isPremium
}
// GetIsSearchable gets the isSearchable property value. Indicates whether the learning content is searchable or not. The default value is true. Optional.
func (m *LearningContent) GetIsSearchable()(*bool) {
    return m.isSearchable
}
// GetLanguageTag gets the languageTag property value. The language of the learning content, for example, en-us or fr-fr. Required.
func (m *LearningContent) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date when the learning content was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumberOfPages gets the numberOfPages property value. The number of pages of the learning content, for example, 9. Optional.
func (m *LearningContent) GetNumberOfPages()(*int32) {
    return m.numberOfPages
}
// GetSkillTags gets the skillTags property value. The skills tags associated with the learning content. Optional.
func (m *LearningContent) GetSkillTags()([]string) {
    return m.skillTags
}
// GetSourceName gets the sourceName property value. The source name of the learning content, such as LinkedIn Learning or Coursera. Optional.
func (m *LearningContent) GetSourceName()(*string) {
    return m.sourceName
}
// GetThumbnailWebUrl gets the thumbnailWebUrl property value. The URL of learning content thumbnail image. Optional.
func (m *LearningContent) GetThumbnailWebUrl()(*string) {
    return m.thumbnailWebUrl
}
// GetTitle gets the title property value. The title of the learning content. Required.
func (m *LearningContent) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *LearningContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalTags() != nil {
        err = writer.WriteCollectionOfStringValues("additionalTags", m.GetAdditionalTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentWebUrl", m.GetContentWebUrl())
        if err != nil {
            return err
        }
    }
    if m.GetContributors() != nil {
        err = writer.WriteCollectionOfStringValues("contributors", m.GetContributors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteISODurationValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPremium", m.GetIsPremium())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfPages", m.GetNumberOfPages())
        if err != nil {
            return err
        }
    }
    if m.GetSkillTags() != nil {
        err = writer.WriteCollectionOfStringValues("skillTags", m.GetSkillTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceName", m.GetSourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailWebUrl", m.GetThumbnailWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalTags sets the additionalTags property value. Keywords, topics, and other tags associated with the learning content. Optional.
func (m *LearningContent) SetAdditionalTags(value []string)() {
    m.additionalTags = value
}
// SetContentWebUrl sets the contentWebUrl property value. The content web URL for the learning content. Required.
func (m *LearningContent) SetContentWebUrl(value *string)() {
    m.contentWebUrl = value
}
// SetContributors sets the contributors property value. The authors, creators, or contributors of the learning content. Optional.
func (m *LearningContent) SetContributors(value []string)() {
    m.contributors = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date when the learning content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description or summary for the learning content. Optional.
func (m *LearningContent) SetDescription(value *string)() {
    m.description = value
}
// SetDuration sets the duration property value. The duration of the learning content in seconds. The value is represented in ISO 8601 format for durations. Optional.
func (m *LearningContent) SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.duration = value
}
// SetExternalId sets the externalId property value. Unique external content ID for the learning content. Required.
func (m *LearningContent) SetExternalId(value *string)() {
    m.externalId = value
}
// SetFormat sets the format property value. The format of the learning content. For example, Course, Video, Book, Book Summary, Audiobook Summary. Optional.
func (m *LearningContent) SetFormat(value *string)() {
    m.format = value
}
// SetIsActive sets the isActive property value. Indicates whether the content is active or not. Inactive content will not show up in the UI. The default value is true. Optional.
func (m *LearningContent) SetIsActive(value *bool)() {
    m.isActive = value
}
// SetIsPremium sets the isPremium property value. Indicates whether the learning content requires the user to sign-in on the learning provider platform or not. The default value is false. Optional.
func (m *LearningContent) SetIsPremium(value *bool)() {
    m.isPremium = value
}
// SetIsSearchable sets the isSearchable property value. Indicates whether the learning content is searchable or not. The default value is true. Optional.
func (m *LearningContent) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// SetLanguageTag sets the languageTag property value. The language of the learning content, for example, en-us or fr-fr. Required.
func (m *LearningContent) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date when the learning content was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumberOfPages sets the numberOfPages property value. The number of pages of the learning content, for example, 9. Optional.
func (m *LearningContent) SetNumberOfPages(value *int32)() {
    m.numberOfPages = value
}
// SetSkillTags sets the skillTags property value. The skills tags associated with the learning content. Optional.
func (m *LearningContent) SetSkillTags(value []string)() {
    m.skillTags = value
}
// SetSourceName sets the sourceName property value. The source name of the learning content, such as LinkedIn Learning or Coursera. Optional.
func (m *LearningContent) SetSourceName(value *string)() {
    m.sourceName = value
}
// SetThumbnailWebUrl sets the thumbnailWebUrl property value. The URL of learning content thumbnail image. Optional.
func (m *LearningContent) SetThumbnailWebUrl(value *string)() {
    m.thumbnailWebUrl = value
}
// SetTitle sets the title property value. The title of the learning content. Required.
func (m *LearningContent) SetTitle(value *string)() {
    m.title = value
}
