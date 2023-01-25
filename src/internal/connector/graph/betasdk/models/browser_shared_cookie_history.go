package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSharedCookieHistory 
type BrowserSharedCookieHistory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The comment for the shared cookie.
    comment *string
    // The name of the cookie.
    displayName *string
    // Controls whether a cookie is a host-only or domain cookie.
    hostOnly *bool
    // The URL of the cookie.
    hostOrDomain *string
    // The lastModifiedBy property
    lastModifiedBy IdentitySetable
    // The OdataType property
    odataType *string
    // The path of the cookie.
    path *string
    // The date and time when the cookie was last published.
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies how the cookies are shared between Microsoft Edge and Internet Explorer. The possible values are: microsoftEdge, internetExplorer11, both, unknownFutureValue.
    sourceEnvironment *BrowserSharedCookieSourceEnvironment
}
// NewBrowserSharedCookieHistory instantiates a new browserSharedCookieHistory and sets the default values.
func NewBrowserSharedCookieHistory()(*BrowserSharedCookieHistory) {
    m := &BrowserSharedCookieHistory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBrowserSharedCookieHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSharedCookieHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSharedCookieHistory(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BrowserSharedCookieHistory) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComment gets the comment property value. The comment for the shared cookie.
func (m *BrowserSharedCookieHistory) GetComment()(*string) {
    return m.comment
}
// GetDisplayName gets the displayName property value. The name of the cookie.
func (m *BrowserSharedCookieHistory) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSharedCookieHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    return res
}
// GetHostOnly gets the hostOnly property value. Controls whether a cookie is a host-only or domain cookie.
func (m *BrowserSharedCookieHistory) GetHostOnly()(*bool) {
    return m.hostOnly
}
// GetHostOrDomain gets the hostOrDomain property value. The URL of the cookie.
func (m *BrowserSharedCookieHistory) GetHostOrDomain()(*string) {
    return m.hostOrDomain
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *BrowserSharedCookieHistory) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BrowserSharedCookieHistory) GetOdataType()(*string) {
    return m.odataType
}
// GetPath gets the path property value. The path of the cookie.
func (m *BrowserSharedCookieHistory) GetPath()(*string) {
    return m.path
}
// GetPublishedDateTime gets the publishedDateTime property value. The date and time when the cookie was last published.
func (m *BrowserSharedCookieHistory) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.publishedDateTime
}
// GetSourceEnvironment gets the sourceEnvironment property value. Specifies how the cookies are shared between Microsoft Edge and Internet Explorer. The possible values are: microsoftEdge, internetExplorer11, both, unknownFutureValue.
func (m *BrowserSharedCookieHistory) GetSourceEnvironment()(*BrowserSharedCookieSourceEnvironment) {
    return m.sourceEnvironment
}
// Serialize serializes information the current object
func (m *BrowserSharedCookieHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hostOnly", m.GetHostOnly())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostOrDomain", m.GetHostOrDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSourceEnvironment() != nil {
        cast := (*m.GetSourceEnvironment()).String()
        err := writer.WriteStringValue("sourceEnvironment", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BrowserSharedCookieHistory) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComment sets the comment property value. The comment for the shared cookie.
func (m *BrowserSharedCookieHistory) SetComment(value *string)() {
    m.comment = value
}
// SetDisplayName sets the displayName property value. The name of the cookie.
func (m *BrowserSharedCookieHistory) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHostOnly sets the hostOnly property value. Controls whether a cookie is a host-only or domain cookie.
func (m *BrowserSharedCookieHistory) SetHostOnly(value *bool)() {
    m.hostOnly = value
}
// SetHostOrDomain sets the hostOrDomain property value. The URL of the cookie.
func (m *BrowserSharedCookieHistory) SetHostOrDomain(value *string)() {
    m.hostOrDomain = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *BrowserSharedCookieHistory) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BrowserSharedCookieHistory) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPath sets the path property value. The path of the cookie.
func (m *BrowserSharedCookieHistory) SetPath(value *string)() {
    m.path = value
}
// SetPublishedDateTime sets the publishedDateTime property value. The date and time when the cookie was last published.
func (m *BrowserSharedCookieHistory) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
// SetSourceEnvironment sets the sourceEnvironment property value. Specifies how the cookies are shared between Microsoft Edge and Internet Explorer. The possible values are: microsoftEdge, internetExplorer11, both, unknownFutureValue.
func (m *BrowserSharedCookieHistory) SetSourceEnvironment(value *BrowserSharedCookieSourceEnvironment)() {
    m.sourceEnvironment = value
}
