package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSecurityProfile provides operations to manage the collection of site entities.
type UserSecurityProfile struct {
    Entity
    // The accounts property
    accounts []UserAccountable
    // The azureSubscriptionId property
    azureSubscriptionId *string
    // The azureTenantId property
    azureTenantId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The displayName property
    displayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The riskScore property
    riskScore *string
    // The tags property
    tags []string
    // The userPrincipalName property
    userPrincipalName *string
    // The vendorInformation property
    vendorInformation SecurityVendorInformationable
}
// NewUserSecurityProfile instantiates a new userSecurityProfile and sets the default values.
func NewUserSecurityProfile()(*UserSecurityProfile) {
    m := &UserSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSecurityProfile(), nil
}
// GetAccounts gets the accounts property value. The accounts property
func (m *UserSecurityProfile) GetAccounts()([]UserAccountable) {
    return m.accounts
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *UserSecurityProfile) GetAzureSubscriptionId()(*string) {
    return m.azureSubscriptionId
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *UserSecurityProfile) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *UserSecurityProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *UserSecurityProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserAccountable, len(val))
            for i, v := range val {
                res[i] = v.(UserAccountable)
            }
            m.SetAccounts(res)
        }
        return nil
    }
    res["azureSubscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
        }
        return nil
    }
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
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
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["vendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(SecurityVendorInformationable))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UserSecurityProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *UserSecurityProfile) GetRiskScore()(*string) {
    return m.riskScore
}
// GetTags gets the tags property value. The tags property
func (m *UserSecurityProfile) GetTags()([]string) {
    return m.tags
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *UserSecurityProfile) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *UserSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
}
// Serialize serializes information the current object
func (m *UserSecurityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accounts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. The accounts property
func (m *UserSecurityProfile) SetAccounts(value []UserAccountable)() {
    m.accounts = value
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *UserSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *UserSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *UserSecurityProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *UserSecurityProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UserSecurityProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *UserSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetTags sets the tags property value. The tags property
func (m *UserSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *UserSecurityProfile) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *UserSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
