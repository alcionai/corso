package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSharedCookie provides operations to call the add method.
type BrowserSharedCookie struct {
    Entity
    // The comment for the shared cookie.
    comment *string
    // The date and time when the shared cookie was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when the shared cookie was deleted.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the cookie.
    displayName *string
    // The history of modifications applied to the cookie.
    history []BrowserSharedCookieHistoryable
    // Controls whether a cookie is a host-only or domain cookie.
    hostOnly *bool
    // The URL of the cookie.
    hostOrDomain *string
    // The user who last modified the cookie.
    lastModifiedBy IdentitySetable
    // The date and time when the cookie was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The path of the cookie.
    path *string
    // The sourceEnvironment property
    sourceEnvironment *BrowserSharedCookieSourceEnvironment
    // The status property
    status *BrowserSharedCookieStatus
}
// NewBrowserSharedCookie instantiates a new browserSharedCookie and sets the default values.
func NewBrowserSharedCookie()(*BrowserSharedCookie) {
    m := &BrowserSharedCookie{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBrowserSharedCookieFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSharedCookieFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSharedCookie(), nil
}
// GetComment gets the comment property value. The comment for the shared cookie.
func (m *BrowserSharedCookie) GetComment()(*string) {
    return m.comment
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the shared cookie was created.
func (m *BrowserSharedCookie) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeletedDateTime gets the deletedDateTime property value. The date and time when the shared cookie was deleted.
func (m *BrowserSharedCookie) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deletedDateTime
}
// GetDisplayName gets the displayName property value. The name of the cookie.
func (m *BrowserSharedCookie) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSharedCookie) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBrowserSharedCookieHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BrowserSharedCookieHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(BrowserSharedCookieHistoryable)
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["hostOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostOnly(val)
        }
        return nil
    }
    res["hostOrDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostOrDomain(val)
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["sourceEnvironment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSharedCookieSourceEnvironment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceEnvironment(val.(*BrowserSharedCookieSourceEnvironment))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSharedCookieStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*BrowserSharedCookieStatus))
        }
        return nil
    }
    return res
}
// GetHistory gets the history property value. The history of modifications applied to the cookie.
func (m *BrowserSharedCookie) GetHistory()([]BrowserSharedCookieHistoryable) {
    return m.history
}
// GetHostOnly gets the hostOnly property value. Controls whether a cookie is a host-only or domain cookie.
func (m *BrowserSharedCookie) GetHostOnly()(*bool) {
    return m.hostOnly
}
// GetHostOrDomain gets the hostOrDomain property value. The URL of the cookie.
func (m *BrowserSharedCookie) GetHostOrDomain()(*string) {
    return m.hostOrDomain
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the cookie.
func (m *BrowserSharedCookie) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the cookie was last modified.
func (m *BrowserSharedCookie) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPath gets the path property value. The path of the cookie.
func (m *BrowserSharedCookie) GetPath()(*string) {
    return m.path
}
// GetSourceEnvironment gets the sourceEnvironment property value. The sourceEnvironment property
func (m *BrowserSharedCookie) GetSourceEnvironment()(*BrowserSharedCookieSourceEnvironment) {
    return m.sourceEnvironment
}
// GetStatus gets the status property value. The status property
func (m *BrowserSharedCookie) GetStatus()(*BrowserSharedCookieStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *BrowserSharedCookie) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
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
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteBoolValue("hostOnly", m.GetHostOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hostOrDomain", m.GetHostOrDomain())
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
        err = writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    if m.GetSourceEnvironment() != nil {
        cast := (*m.GetSourceEnvironment()).String()
        err = writer.WriteStringValue("sourceEnvironment", &cast)
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
// SetComment sets the comment property value. The comment for the shared cookie.
func (m *BrowserSharedCookie) SetComment(value *string)() {
    m.comment = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the shared cookie was created.
func (m *BrowserSharedCookie) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeletedDateTime sets the deletedDateTime property value. The date and time when the shared cookie was deleted.
func (m *BrowserSharedCookie) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the cookie.
func (m *BrowserSharedCookie) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHistory sets the history property value. The history of modifications applied to the cookie.
func (m *BrowserSharedCookie) SetHistory(value []BrowserSharedCookieHistoryable)() {
    m.history = value
}
// SetHostOnly sets the hostOnly property value. Controls whether a cookie is a host-only or domain cookie.
func (m *BrowserSharedCookie) SetHostOnly(value *bool)() {
    m.hostOnly = value
}
// SetHostOrDomain sets the hostOrDomain property value. The URL of the cookie.
func (m *BrowserSharedCookie) SetHostOrDomain(value *string)() {
    m.hostOrDomain = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the cookie.
func (m *BrowserSharedCookie) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the cookie was last modified.
func (m *BrowserSharedCookie) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPath sets the path property value. The path of the cookie.
func (m *BrowserSharedCookie) SetPath(value *string)() {
    m.path = value
}
// SetSourceEnvironment sets the sourceEnvironment property value. The sourceEnvironment property
func (m *BrowserSharedCookie) SetSourceEnvironment(value *BrowserSharedCookieSourceEnvironment)() {
    m.sourceEnvironment = value
}
// SetStatus sets the status property value. The status property
func (m *BrowserSharedCookie) SetStatus(value *BrowserSharedCookieStatus)() {
    m.status = value
}
