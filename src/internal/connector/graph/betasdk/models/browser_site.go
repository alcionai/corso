package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSite singleton entity which is used to specify IE mode site metadata
type BrowserSite struct {
    Entity
    // Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
    allowRedirect *bool
    // The comment for the site.
    comment *string
    // The compatibilityMode property
    compatibilityMode *BrowserSiteCompatibilityMode
    // The date and time when the site was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when the site was deleted.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The history of modifications applied to the site.
    history []BrowserSiteHistoryable
    // The user who last modified the site.
    lastModifiedBy IdentitySetable
    // The date and time when the site was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The mergeType property
    mergeType *BrowserSiteMergeType
    // The status property
    status *BrowserSiteStatus
    // The targetEnvironment property
    targetEnvironment *BrowserSiteTargetEnvironment
    // The URL of the site.
    webUrl *string
}
// NewBrowserSite instantiates a new browserSite and sets the default values.
func NewBrowserSite()(*BrowserSite) {
    m := &BrowserSite{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBrowserSiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSite(), nil
}
// GetAllowRedirect gets the allowRedirect property value. Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
func (m *BrowserSite) GetAllowRedirect()(*bool) {
    return m.allowRedirect
}
// GetComment gets the comment property value. The comment for the site.
func (m *BrowserSite) GetComment()(*string) {
    return m.comment
}
// GetCompatibilityMode gets the compatibilityMode property value. The compatibilityMode property
func (m *BrowserSite) GetCompatibilityMode()(*BrowserSiteCompatibilityMode) {
    return m.compatibilityMode
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the site was created.
func (m *BrowserSite) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeletedDateTime gets the deletedDateTime property value. The date and time when the site was deleted.
func (m *BrowserSite) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deletedDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowRedirect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowRedirect(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["compatibilityMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteCompatibilityMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompatibilityMode(val.(*BrowserSiteCompatibilityMode))
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
    res["deletedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDateTime(val)
        }
        return nil
    }
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBrowserSiteHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BrowserSiteHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(BrowserSiteHistoryable)
            }
            m.SetHistory(res)
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
    res["mergeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteMergeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeType(val.(*BrowserSiteMergeType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*BrowserSiteStatus))
        }
        return nil
    }
    res["targetEnvironment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteTargetEnvironment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetEnvironment(val.(*BrowserSiteTargetEnvironment))
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
// GetHistory gets the history property value. The history of modifications applied to the site.
func (m *BrowserSite) GetHistory()([]BrowserSiteHistoryable) {
    return m.history
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the site.
func (m *BrowserSite) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the site was last modified.
func (m *BrowserSite) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMergeType gets the mergeType property value. The mergeType property
func (m *BrowserSite) GetMergeType()(*BrowserSiteMergeType) {
    return m.mergeType
}
// GetStatus gets the status property value. The status property
func (m *BrowserSite) GetStatus()(*BrowserSiteStatus) {
    return m.status
}
// GetTargetEnvironment gets the targetEnvironment property value. The targetEnvironment property
func (m *BrowserSite) GetTargetEnvironment()(*BrowserSiteTargetEnvironment) {
    return m.targetEnvironment
}
// GetWebUrl gets the webUrl property value. The URL of the site.
func (m *BrowserSite) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *BrowserSite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowRedirect", m.GetAllowRedirect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    if m.GetCompatibilityMode() != nil {
        cast := (*m.GetCompatibilityMode()).String()
        err = writer.WriteStringValue("compatibilityMode", &cast)
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
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
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
    if m.GetMergeType() != nil {
        cast := (*m.GetMergeType()).String()
        err = writer.WriteStringValue("mergeType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetEnvironment() != nil {
        cast := (*m.GetTargetEnvironment()).String()
        err = writer.WriteStringValue("targetEnvironment", &cast)
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
// SetAllowRedirect sets the allowRedirect property value. Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
func (m *BrowserSite) SetAllowRedirect(value *bool)() {
    m.allowRedirect = value
}
// SetComment sets the comment property value. The comment for the site.
func (m *BrowserSite) SetComment(value *string)() {
    m.comment = value
}
// SetCompatibilityMode sets the compatibilityMode property value. The compatibilityMode property
func (m *BrowserSite) SetCompatibilityMode(value *BrowserSiteCompatibilityMode)() {
    m.compatibilityMode = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the site was created.
func (m *BrowserSite) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeletedDateTime sets the deletedDateTime property value. The date and time when the site was deleted.
func (m *BrowserSite) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// SetHistory sets the history property value. The history of modifications applied to the site.
func (m *BrowserSite) SetHistory(value []BrowserSiteHistoryable)() {
    m.history = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the site.
func (m *BrowserSite) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the site was last modified.
func (m *BrowserSite) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMergeType sets the mergeType property value. The mergeType property
func (m *BrowserSite) SetMergeType(value *BrowserSiteMergeType)() {
    m.mergeType = value
}
// SetStatus sets the status property value. The status property
func (m *BrowserSite) SetStatus(value *BrowserSiteStatus)() {
    m.status = value
}
// SetTargetEnvironment sets the targetEnvironment property value. The targetEnvironment property
func (m *BrowserSite) SetTargetEnvironment(value *BrowserSiteTargetEnvironment)() {
    m.targetEnvironment = value
}
// SetWebUrl sets the webUrl property value. The URL of the site.
func (m *BrowserSite) SetWebUrl(value *string)() {
    m.webUrl = value
}
