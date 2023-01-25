package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialSingleSignOnExtension 
type CredentialSingleSignOnExtension struct {
    SingleSignOnExtension
    // Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
    configurations []KeyTypedValuePairable
    // Gets or sets a list of hosts or domain names for which the app extension performs SSO.
    domains []string
    // Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
    extensionIdentifier *string
    // Gets or sets the case-sensitive realm name for this profile.
    realm *string
    // Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
    teamIdentifier *string
}
// NewCredentialSingleSignOnExtension instantiates a new CredentialSingleSignOnExtension and sets the default values.
func NewCredentialSingleSignOnExtension()(*CredentialSingleSignOnExtension) {
    m := &CredentialSingleSignOnExtension{
        SingleSignOnExtension: *NewSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.credentialSingleSignOnExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCredentialSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialSingleSignOnExtension(), nil
}
// GetConfigurations gets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *CredentialSingleSignOnExtension) GetConfigurations()([]KeyTypedValuePairable) {
    return m.configurations
}
// GetDomains gets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *CredentialSingleSignOnExtension) GetDomains()([]string) {
    return m.domains
}
// GetExtensionIdentifier gets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
func (m *CredentialSingleSignOnExtension) GetExtensionIdentifier()(*string) {
    return m.extensionIdentifier
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SingleSignOnExtension.GetFieldDeserializers()
    res["configurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyTypedValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyTypedValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyTypedValuePairable)
            }
            m.SetConfigurations(res)
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDomains(res)
        }
        return nil
    }
    res["extensionIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionIdentifier(val)
        }
        return nil
    }
    res["realm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealm(val)
        }
        return nil
    }
    res["teamIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamIdentifier(val)
        }
        return nil
    }
    return res
}
// GetRealm gets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *CredentialSingleSignOnExtension) GetRealm()(*string) {
    return m.realm
}
// GetTeamIdentifier gets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
func (m *CredentialSingleSignOnExtension) GetTeamIdentifier()(*string) {
    return m.teamIdentifier
}
// Serialize serializes information the current object
func (m *CredentialSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurations()))
        for i, v := range m.GetConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomains() != nil {
        err = writer.WriteCollectionOfStringValues("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("extensionIdentifier", m.GetExtensionIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("realm", m.GetRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamIdentifier", m.GetTeamIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurations sets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *CredentialSingleSignOnExtension) SetConfigurations(value []KeyTypedValuePairable)() {
    m.configurations = value
}
// SetDomains sets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *CredentialSingleSignOnExtension) SetDomains(value []string)() {
    m.domains = value
}
// SetExtensionIdentifier sets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
func (m *CredentialSingleSignOnExtension) SetExtensionIdentifier(value *string)() {
    m.extensionIdentifier = value
}
// SetRealm sets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *CredentialSingleSignOnExtension) SetRealm(value *string)() {
    m.realm = value
}
// SetTeamIdentifier sets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
func (m *CredentialSingleSignOnExtension) SetTeamIdentifier(value *string)() {
    m.teamIdentifier = value
}
