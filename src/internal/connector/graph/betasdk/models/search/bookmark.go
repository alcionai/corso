package search

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Bookmark 
type Bookmark struct {
    SearchAnswer
    // Timestamp of when the bookmark will stop to appear as a search result. Set as null for always available.
    availabilityEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp of when the bookmark will start to appear as a search result. Set as null for always available.
    availabilityStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Categories commonly used to describe this bookmark. For example, IT and HR.
    categories []string
    // List of security groups able to view this bookmark.
    groupIds []string
    // True if this bookmark was suggested to the admin by a user or was mined and suggested by Microsoft. Read-only.
    isSuggested *bool
    // Keywords that trigger this bookmark to appear in search results.
    keywords AnswerKeywordable
    // A list of language names that are geographically specific and that this bookmark can be viewed in. Each language tag value follows the pattern {language}-{region}. As an example, en-us is English as used in the United States. See supported language tags for the list of possible values.
    languageTags []string
    // List of devices and operating systems able to view this bookmark. Possible values are: unknown, android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, androidASOP.
    platforms []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType
    // List of Power Apps associated with this bookmark. If users add existing Power Apps to a bookmark, they can complete tasks, such as to enter vacation time or to report expenses on the search results page.
    powerAppIds []string
    // The state property
    state *AnswerState
    // Variations of a bookmark for different countries or devices. Use when you need to show different content to users based on their device, country/region, or both. The date and group settings will apply to all variations.
    targetedVariations []AnswerVariantable
}
// NewBookmark instantiates a new Bookmark and sets the default values.
func NewBookmark()(*Bookmark) {
    m := &Bookmark{
        SearchAnswer: *NewSearchAnswer(),
    }
    return m
}
// CreateBookmarkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookmarkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookmark(), nil
}
// GetAvailabilityEndDateTime gets the availabilityEndDateTime property value. Timestamp of when the bookmark will stop to appear as a search result. Set as null for always available.
func (m *Bookmark) GetAvailabilityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.availabilityEndDateTime
}
// GetAvailabilityStartDateTime gets the availabilityStartDateTime property value. Timestamp of when the bookmark will start to appear as a search result. Set as null for always available.
func (m *Bookmark) GetAvailabilityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.availabilityStartDateTime
}
// GetCategories gets the categories property value. Categories commonly used to describe this bookmark. For example, IT and HR.
func (m *Bookmark) GetCategories()([]string) {
    return m.categories
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Bookmark) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SearchAnswer.GetFieldDeserializers()
    res["availabilityEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityEndDateTime(val)
        }
        return nil
    }
    res["availabilityStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityStartDateTime(val)
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
    res["groupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupIds(res)
        }
        return nil
    }
    res["isSuggested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSuggested(val)
        }
        return nil
    }
    res["keywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAnswerKeywordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeywords(val.(AnswerKeywordable))
        }
        return nil
    }
    res["languageTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLanguageTags(res)
        }
        return nil
    }
    res["platforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType, len(val))
            for i, v := range val {
                res[i] = *(v.(*ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType))
            }
            m.SetPlatforms(res)
        }
        return nil
    }
    res["powerAppIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPowerAppIds(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnswerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AnswerState))
        }
        return nil
    }
    res["targetedVariations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAnswerVariantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AnswerVariantable, len(val))
            for i, v := range val {
                res[i] = v.(AnswerVariantable)
            }
            m.SetTargetedVariations(res)
        }
        return nil
    }
    return res
}
// GetGroupIds gets the groupIds property value. List of security groups able to view this bookmark.
func (m *Bookmark) GetGroupIds()([]string) {
    return m.groupIds
}
// GetIsSuggested gets the isSuggested property value. True if this bookmark was suggested to the admin by a user or was mined and suggested by Microsoft. Read-only.
func (m *Bookmark) GetIsSuggested()(*bool) {
    return m.isSuggested
}
// GetKeywords gets the keywords property value. Keywords that trigger this bookmark to appear in search results.
func (m *Bookmark) GetKeywords()(AnswerKeywordable) {
    return m.keywords
}
// GetLanguageTags gets the languageTags property value. A list of language names that are geographically specific and that this bookmark can be viewed in. Each language tag value follows the pattern {language}-{region}. As an example, en-us is English as used in the United States. See supported language tags for the list of possible values.
func (m *Bookmark) GetLanguageTags()([]string) {
    return m.languageTags
}
// GetPlatforms gets the platforms property value. List of devices and operating systems able to view this bookmark. Possible values are: unknown, android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, androidASOP.
func (m *Bookmark) GetPlatforms()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType) {
    return m.platforms
}
// GetPowerAppIds gets the powerAppIds property value. List of Power Apps associated with this bookmark. If users add existing Power Apps to a bookmark, they can complete tasks, such as to enter vacation time or to report expenses on the search results page.
func (m *Bookmark) GetPowerAppIds()([]string) {
    return m.powerAppIds
}
// GetState gets the state property value. The state property
func (m *Bookmark) GetState()(*AnswerState) {
    return m.state
}
// GetTargetedVariations gets the targetedVariations property value. Variations of a bookmark for different countries or devices. Use when you need to show different content to users based on their device, country/region, or both. The date and group settings will apply to all variations.
func (m *Bookmark) GetTargetedVariations()([]AnswerVariantable) {
    return m.targetedVariations
}
// Serialize serializes information the current object
func (m *Bookmark) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SearchAnswer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("availabilityEndDateTime", m.GetAvailabilityEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("availabilityStartDateTime", m.GetAvailabilityStartDateTime())
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
    if m.GetGroupIds() != nil {
        err = writer.WriteCollectionOfStringValues("groupIds", m.GetGroupIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSuggested", m.GetIsSuggested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("keywords", m.GetKeywords())
        if err != nil {
            return err
        }
    }
    if m.GetLanguageTags() != nil {
        err = writer.WriteCollectionOfStringValues("languageTags", m.GetLanguageTags())
        if err != nil {
            return err
        }
    }
    if m.GetPlatforms() != nil {
        err = writer.WriteCollectionOfStringValues("platforms", ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.SerializeDevicePlatformType(m.GetPlatforms()))
        if err != nil {
            return err
        }
    }
    if m.GetPowerAppIds() != nil {
        err = writer.WriteCollectionOfStringValues("powerAppIds", m.GetPowerAppIds())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedVariations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedVariations()))
        for i, v := range m.GetTargetedVariations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("targetedVariations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityEndDateTime sets the availabilityEndDateTime property value. Timestamp of when the bookmark will stop to appear as a search result. Set as null for always available.
func (m *Bookmark) SetAvailabilityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.availabilityEndDateTime = value
}
// SetAvailabilityStartDateTime sets the availabilityStartDateTime property value. Timestamp of when the bookmark will start to appear as a search result. Set as null for always available.
func (m *Bookmark) SetAvailabilityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.availabilityStartDateTime = value
}
// SetCategories sets the categories property value. Categories commonly used to describe this bookmark. For example, IT and HR.
func (m *Bookmark) SetCategories(value []string)() {
    m.categories = value
}
// SetGroupIds sets the groupIds property value. List of security groups able to view this bookmark.
func (m *Bookmark) SetGroupIds(value []string)() {
    m.groupIds = value
}
// SetIsSuggested sets the isSuggested property value. True if this bookmark was suggested to the admin by a user or was mined and suggested by Microsoft. Read-only.
func (m *Bookmark) SetIsSuggested(value *bool)() {
    m.isSuggested = value
}
// SetKeywords sets the keywords property value. Keywords that trigger this bookmark to appear in search results.
func (m *Bookmark) SetKeywords(value AnswerKeywordable)() {
    m.keywords = value
}
// SetLanguageTags sets the languageTags property value. A list of language names that are geographically specific and that this bookmark can be viewed in. Each language tag value follows the pattern {language}-{region}. As an example, en-us is English as used in the United States. See supported language tags for the list of possible values.
func (m *Bookmark) SetLanguageTags(value []string)() {
    m.languageTags = value
}
// SetPlatforms sets the platforms property value. List of devices and operating systems able to view this bookmark. Possible values are: unknown, android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, androidASOP.
func (m *Bookmark) SetPlatforms(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType)() {
    m.platforms = value
}
// SetPowerAppIds sets the powerAppIds property value. List of Power Apps associated with this bookmark. If users add existing Power Apps to a bookmark, they can complete tasks, such as to enter vacation time or to report expenses on the search results page.
func (m *Bookmark) SetPowerAppIds(value []string)() {
    m.powerAppIds = value
}
// SetState sets the state property value. The state property
func (m *Bookmark) SetState(value *AnswerState)() {
    m.state = value
}
// SetTargetedVariations sets the targetedVariations property value. Variations of a bookmark for different countries or devices. Use when you need to show different content to users based on their device, country/region, or both. The date and group settings will apply to all variations.
func (m *Bookmark) SetTargetedVariations(value []AnswerVariantable)() {
    m.targetedVariations = value
}
