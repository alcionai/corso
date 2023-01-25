package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// B2cIdentityUserFlow 
type B2cIdentityUserFlow struct {
    IdentityUserFlow
    // Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
    apiConnectorConfiguration UserFlowApiConnectorConfigurationable
    // Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
    defaultLanguageTag *string
    // The identityProviders property
    identityProviders []IdentityProviderable
    // The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
    isLanguageCustomizationEnabled *bool
    // The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
    languages []UserFlowLanguageConfigurationable
    // The user attribute assignments included in the user flow.
    userAttributeAssignments []IdentityUserFlowAttributeAssignmentable
    // The userFlowIdentityProviders property
    userFlowIdentityProviders []IdentityProviderBaseable
}
// NewB2cIdentityUserFlow instantiates a new B2cIdentityUserFlow and sets the default values.
func NewB2cIdentityUserFlow()(*B2cIdentityUserFlow) {
    m := &B2cIdentityUserFlow{
        IdentityUserFlow: *NewIdentityUserFlow(),
    }
    return m
}
// CreateB2cIdentityUserFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cIdentityUserFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cIdentityUserFlow(), nil
}
// GetApiConnectorConfiguration gets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2cIdentityUserFlow) GetApiConnectorConfiguration()(UserFlowApiConnectorConfigurationable) {
    return m.apiConnectorConfiguration
}
// GetDefaultLanguageTag gets the defaultLanguageTag property value. Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
func (m *B2cIdentityUserFlow) GetDefaultLanguageTag()(*string) {
    return m.defaultLanguageTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cIdentityUserFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityUserFlow.GetFieldDeserializers()
    res["apiConnectorConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFlowApiConnectorConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiConnectorConfiguration(val.(UserFlowApiConnectorConfigurationable))
        }
        return nil
    }
    res["defaultLanguageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguageTag(val)
        }
        return nil
    }
    res["identityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderable)
            }
            m.SetIdentityProviders(res)
        }
        return nil
    }
    res["isLanguageCustomizationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLanguageCustomizationEnabled(val)
        }
        return nil
    }
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFlowLanguageConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguageConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(UserFlowLanguageConfigurationable)
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["userAttributeAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttributeAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityUserFlowAttributeAssignmentable)
            }
            m.SetUserAttributeAssignments(res)
        }
        return nil
    }
    res["userFlowIdentityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderBaseable)
            }
            m.SetUserFlowIdentityProviders(res)
        }
        return nil
    }
    return res
}
// GetIdentityProviders gets the identityProviders property value. The identityProviders property
func (m *B2cIdentityUserFlow) GetIdentityProviders()([]IdentityProviderable) {
    return m.identityProviders
}
// GetIsLanguageCustomizationEnabled gets the isLanguageCustomizationEnabled property value. The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
func (m *B2cIdentityUserFlow) GetIsLanguageCustomizationEnabled()(*bool) {
    return m.isLanguageCustomizationEnabled
}
// GetLanguages gets the languages property value. The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *B2cIdentityUserFlow) GetLanguages()([]UserFlowLanguageConfigurationable) {
    return m.languages
}
// GetUserAttributeAssignments gets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2cIdentityUserFlow) GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignmentable) {
    return m.userAttributeAssignments
}
// GetUserFlowIdentityProviders gets the userFlowIdentityProviders property value. The userFlowIdentityProviders property
func (m *B2cIdentityUserFlow) GetUserFlowIdentityProviders()([]IdentityProviderBaseable) {
    return m.userFlowIdentityProviders
}
// Serialize serializes information the current object
func (m *B2cIdentityUserFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityUserFlow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("apiConnectorConfiguration", m.GetApiConnectorConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultLanguageTag", m.GetDefaultLanguageTag())
        if err != nil {
            return err
        }
    }
    if m.GetIdentityProviders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIdentityProviders()))
        for i, v := range m.GetIdentityProviders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("identityProviders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLanguageCustomizationEnabled", m.GetIsLanguageCustomizationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetLanguages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserAttributeAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserAttributeAssignments()))
        for i, v := range m.GetUserAttributeAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserFlowIdentityProviders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserFlowIdentityProviders()))
        for i, v := range m.GetUserFlowIdentityProviders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userFlowIdentityProviders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiConnectorConfiguration sets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2cIdentityUserFlow) SetApiConnectorConfiguration(value UserFlowApiConnectorConfigurationable)() {
    m.apiConnectorConfiguration = value
}
// SetDefaultLanguageTag sets the defaultLanguageTag property value. Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
func (m *B2cIdentityUserFlow) SetDefaultLanguageTag(value *string)() {
    m.defaultLanguageTag = value
}
// SetIdentityProviders sets the identityProviders property value. The identityProviders property
func (m *B2cIdentityUserFlow) SetIdentityProviders(value []IdentityProviderable)() {
    m.identityProviders = value
}
// SetIsLanguageCustomizationEnabled sets the isLanguageCustomizationEnabled property value. The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
func (m *B2cIdentityUserFlow) SetIsLanguageCustomizationEnabled(value *bool)() {
    m.isLanguageCustomizationEnabled = value
}
// SetLanguages sets the languages property value. The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *B2cIdentityUserFlow) SetLanguages(value []UserFlowLanguageConfigurationable)() {
    m.languages = value
}
// SetUserAttributeAssignments sets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2cIdentityUserFlow) SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignmentable)() {
    m.userAttributeAssignments = value
}
// SetUserFlowIdentityProviders sets the userFlowIdentityProviders property value. The userFlowIdentityProviders property
func (m *B2cIdentityUserFlow) SetUserFlowIdentityProviders(value []IdentityProviderBaseable)() {
    m.userFlowIdentityProviders = value
}
