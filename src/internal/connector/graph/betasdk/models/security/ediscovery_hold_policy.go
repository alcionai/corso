package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryHoldPolicy 
type EdiscoveryHoldPolicy struct {
    PolicyBase
    // KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
    contentQuery *string
    // Lists any errors that happened while placing the hold.
    errors []string
    // Indicates whether the hold is enabled and actively holding content.
    isEnabled *bool
    // Data sources that represent SharePoint sites.
    siteSources []SiteSourceable
    // Data sources that represent Exchange mailboxes.
    userSources []UserSourceable
}
// NewEdiscoveryHoldPolicy instantiates a new EdiscoveryHoldPolicy and sets the default values.
func NewEdiscoveryHoldPolicy()(*EdiscoveryHoldPolicy) {
    m := &EdiscoveryHoldPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.security.ediscoveryHoldPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEdiscoveryHoldPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryHoldPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryHoldPolicy(), nil
}
// GetContentQuery gets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *EdiscoveryHoldPolicy) GetContentQuery()(*string) {
    return m.contentQuery
}
// GetErrors gets the errors property value. Lists any errors that happened while placing the hold.
func (m *EdiscoveryHoldPolicy) GetErrors()([]string) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryHoldPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["contentQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["siteSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SiteSourceable, len(val))
            for i, v := range val {
                res[i] = v.(SiteSourceable)
            }
            m.SetSiteSources(res)
        }
        return nil
    }
    res["userSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSourceable, len(val))
            for i, v := range val {
                res[i] = v.(UserSourceable)
            }
            m.SetUserSources(res)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *EdiscoveryHoldPolicy) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetSiteSources gets the siteSources property value. Data sources that represent SharePoint sites.
func (m *EdiscoveryHoldPolicy) GetSiteSources()([]SiteSourceable) {
    return m.siteSources
}
// GetUserSources gets the userSources property value. Data sources that represent Exchange mailboxes.
func (m *EdiscoveryHoldPolicy) GetUserSources()([]UserSourceable) {
    return m.userSources
}
// Serialize serializes information the current object
func (m *EdiscoveryHoldPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
        if err != nil {
            return err
        }
    }
    if m.GetErrors() != nil {
        err = writer.WriteCollectionOfStringValues("errors", m.GetErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetSiteSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSiteSources()))
        for i, v := range m.GetSiteSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("siteSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSources()))
        for i, v := range m.GetUserSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentQuery sets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *EdiscoveryHoldPolicy) SetContentQuery(value *string)() {
    m.contentQuery = value
}
// SetErrors sets the errors property value. Lists any errors that happened while placing the hold.
func (m *EdiscoveryHoldPolicy) SetErrors(value []string)() {
    m.errors = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *EdiscoveryHoldPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetSiteSources sets the siteSources property value. Data sources that represent SharePoint sites.
func (m *EdiscoveryHoldPolicy) SetSiteSources(value []SiteSourceable)() {
    m.siteSources = value
}
// SetUserSources sets the userSources property value. Data sources that represent Exchange mailboxes.
func (m *EdiscoveryHoldPolicy) SetUserSources(value []UserSourceable)() {
    m.userSources = value
}
