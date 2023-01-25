package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobilityManagementPolicy 
type MobilityManagementPolicy struct {
    Entity
    // Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
    appliesTo *PolicyScope
    // Compliance URL of the mobility management application.
    complianceUrl *string
    // Description of the mobility management application.
    description *string
    // Discovery URL of the mobility management application.
    discoveryUrl *string
    // Display name of the mobility management application.
    displayName *string
    // Azure AD groups under the scope of the mobility management application if appliesTo is selected
    includedGroups []Groupable
    // Whether policy is valid. Invalid policies may not be updated and should be deleted.
    isValid *bool
    // Terms of Use URL of the mobility management application.
    termsOfUseUrl *string
}
// NewMobilityManagementPolicy instantiates a new MobilityManagementPolicy and sets the default values.
func NewMobilityManagementPolicy()(*MobilityManagementPolicy) {
    m := &MobilityManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobilityManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobilityManagementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobilityManagementPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
func (m *MobilityManagementPolicy) GetAppliesTo()(*PolicyScope) {
    return m.appliesTo
}
// GetComplianceUrl gets the complianceUrl property value. Compliance URL of the mobility management application.
func (m *MobilityManagementPolicy) GetComplianceUrl()(*string) {
    return m.complianceUrl
}
// GetDescription gets the description property value. Description of the mobility management application.
func (m *MobilityManagementPolicy) GetDescription()(*string) {
    return m.description
}
// GetDiscoveryUrl gets the discoveryUrl property value. Discovery URL of the mobility management application.
func (m *MobilityManagementPolicy) GetDiscoveryUrl()(*string) {
    return m.discoveryUrl
}
// GetDisplayName gets the displayName property value. Display name of the mobility management application.
func (m *MobilityManagementPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobilityManagementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val.(*PolicyScope))
        }
        return nil
    }
    res["complianceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceUrl(val)
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
    res["discoveryUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoveryUrl(val)
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
    res["includedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Groupable, len(val))
            for i, v := range val {
                res[i] = v.(Groupable)
            }
            m.SetIncludedGroups(res)
        }
        return nil
    }
    res["isValid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValid(val)
        }
        return nil
    }
    res["termsOfUseUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsOfUseUrl(val)
        }
        return nil
    }
    return res
}
// GetIncludedGroups gets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *MobilityManagementPolicy) GetIncludedGroups()([]Groupable) {
    return m.includedGroups
}
// GetIsValid gets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
func (m *MobilityManagementPolicy) GetIsValid()(*bool) {
    return m.isValid
}
// GetTermsOfUseUrl gets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
func (m *MobilityManagementPolicy) GetTermsOfUseUrl()(*string) {
    return m.termsOfUseUrl
}
// Serialize serializes information the current object
func (m *MobilityManagementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err = writer.WriteStringValue("appliesTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("complianceUrl", m.GetComplianceUrl())
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
        err = writer.WriteStringValue("discoveryUrl", m.GetDiscoveryUrl())
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
    if m.GetIncludedGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludedGroups()))
        for i, v := range m.GetIncludedGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("includedGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValid", m.GetIsValid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("termsOfUseUrl", m.GetTermsOfUseUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
func (m *MobilityManagementPolicy) SetAppliesTo(value *PolicyScope)() {
    m.appliesTo = value
}
// SetComplianceUrl sets the complianceUrl property value. Compliance URL of the mobility management application.
func (m *MobilityManagementPolicy) SetComplianceUrl(value *string)() {
    m.complianceUrl = value
}
// SetDescription sets the description property value. Description of the mobility management application.
func (m *MobilityManagementPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDiscoveryUrl sets the discoveryUrl property value. Discovery URL of the mobility management application.
func (m *MobilityManagementPolicy) SetDiscoveryUrl(value *string)() {
    m.discoveryUrl = value
}
// SetDisplayName sets the displayName property value. Display name of the mobility management application.
func (m *MobilityManagementPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIncludedGroups sets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *MobilityManagementPolicy) SetIncludedGroups(value []Groupable)() {
    m.includedGroups = value
}
// SetIsValid sets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
func (m *MobilityManagementPolicy) SetIsValid(value *bool)() {
    m.isValid = value
}
// SetTermsOfUseUrl sets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
func (m *MobilityManagementPolicy) SetTermsOfUseUrl(value *string)() {
    m.termsOfUseUrl = value
}
