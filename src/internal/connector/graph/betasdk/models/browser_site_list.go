package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSiteList a singleton entity which is used to specify IE mode site list metadata
type BrowserSiteList struct {
    Entity
    // The description of the site list.
    description *string
    // The name of the site list.
    displayName *string
    // The user who last modified the site list.
    lastModifiedBy IdentitySetable
    // The date and time when the site list was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user who published the site list.
    publishedBy IdentitySetable
    // The date and time when the site list was published.
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The current revision of the site list.
    revision *string
    // A collection of shared cookies defined for the site list.
    sharedCookies []BrowserSharedCookieable
    // A collection of sites defined for the site list.
    sites []BrowserSiteable
    // The status property
    status *BrowserSiteListStatus
}
// NewBrowserSiteList instantiates a new browserSiteList and sets the default values.
func NewBrowserSiteList()(*BrowserSiteList) {
    m := &BrowserSiteList{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBrowserSiteListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSiteListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSiteList(), nil
}
// GetDescription gets the description property value. The description of the site list.
func (m *BrowserSiteList) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the site list.
func (m *BrowserSiteList) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSiteList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["publishedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["publishedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    res["revision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevision(val)
        }
        return nil
    }
    res["sharedCookies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBrowserSharedCookieFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BrowserSharedCookieable, len(val))
            for i, v := range val {
                res[i] = v.(BrowserSharedCookieable)
            }
            m.SetSharedCookies(res)
        }
        return nil
    }
    res["sites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBrowserSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BrowserSiteable, len(val))
            for i, v := range val {
                res[i] = v.(BrowserSiteable)
            }
            m.SetSites(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteListStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*BrowserSiteListStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the site list.
func (m *BrowserSiteList) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the site list was last modified.
func (m *BrowserSiteList) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPublishedBy gets the publishedBy property value. The user who published the site list.
func (m *BrowserSiteList) GetPublishedBy()(IdentitySetable) {
    return m.publishedBy
}
// GetPublishedDateTime gets the publishedDateTime property value. The date and time when the site list was published.
func (m *BrowserSiteList) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.publishedDateTime
}
// GetRevision gets the revision property value. The current revision of the site list.
func (m *BrowserSiteList) GetRevision()(*string) {
    return m.revision
}
// GetSharedCookies gets the sharedCookies property value. A collection of shared cookies defined for the site list.
func (m *BrowserSiteList) GetSharedCookies()([]BrowserSharedCookieable) {
    return m.sharedCookies
}
// GetSites gets the sites property value. A collection of sites defined for the site list.
func (m *BrowserSiteList) GetSites()([]BrowserSiteable) {
    return m.sites
}
// GetStatus gets the status property value. The status property
func (m *BrowserSiteList) GetStatus()(*BrowserSiteListStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *BrowserSiteList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteObjectValue("publishedBy", m.GetPublishedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("revision", m.GetRevision())
        if err != nil {
            return err
        }
    }
    if m.GetSharedCookies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedCookies()))
        for i, v := range m.GetSharedCookies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sharedCookies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSites()))
        for i, v := range m.GetSites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sites", cast)
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
    return nil
}
// SetDescription sets the description property value. The description of the site list.
func (m *BrowserSiteList) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the site list.
func (m *BrowserSiteList) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the site list.
func (m *BrowserSiteList) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the site list was last modified.
func (m *BrowserSiteList) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPublishedBy sets the publishedBy property value. The user who published the site list.
func (m *BrowserSiteList) SetPublishedBy(value IdentitySetable)() {
    m.publishedBy = value
}
// SetPublishedDateTime sets the publishedDateTime property value. The date and time when the site list was published.
func (m *BrowserSiteList) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
// SetRevision sets the revision property value. The current revision of the site list.
func (m *BrowserSiteList) SetRevision(value *string)() {
    m.revision = value
}
// SetSharedCookies sets the sharedCookies property value. A collection of shared cookies defined for the site list.
func (m *BrowserSiteList) SetSharedCookies(value []BrowserSharedCookieable)() {
    m.sharedCookies = value
}
// SetSites sets the sites property value. A collection of sites defined for the site list.
func (m *BrowserSiteList) SetSites(value []BrowserSiteable)() {
    m.sites = value
}
// SetStatus sets the status property value. The status property
func (m *BrowserSiteList) SetStatus(value *BrowserSiteListStatus)() {
    m.status = value
}
