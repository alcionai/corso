package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365ActiveUserDetail 
type Office365ActiveUserDetail struct {
    Entity
    // All the products assigned for the user.
    assignedProducts []string
    // The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
    deletedDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates.
    displayName *string
    // The date when user last read or sent email.
    exchangeLastActivityDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The last date when the user was assigned an Exchange license.
    exchangeLicenseAssignDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Whether the user has been assigned an Exchange license.
    hasExchangeLicense *bool
    // Whether the user has been assigned a OneDrive license.
    hasOneDriveLicense *bool
    // Whether the user has been assigned a SharePoint license.
    hasSharePointLicense *bool
    // Whether the user has been assigned a Skype For Business license.
    hasSkypeForBusinessLicense *bool
    // Whether the user has been assigned a Teams license.
    hasTeamsLicense *bool
    // Whether the user has been assigned a Yammer license.
    hasYammerLicense *bool
    // Whether this user has been deleted or soft deleted.
    isDeleted *bool
    // The date when user last viewed or edited files, shared files internally or externally, or synced files.
    oneDriveLastActivityDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The last date when the user was assigned a OneDrive license.
    oneDriveLicenseAssignDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The latest date of the content.
    reportRefreshDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
    sharePointLastActivityDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The last date when the user was assigned a SharePoint license.
    sharePointLicenseAssignDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
    skypeForBusinessLastActivityDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The last date when the user was assigned a Skype For Business license.
    skypeForBusinessLicenseAssignDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
    teamsLastActivityDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The last date when the user was assigned a Teams license.
    teamsLicenseAssignDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
    userPrincipalName *string
    // The date when user last posted, read, or liked message.
    yammerLastActivityDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The last date when the user was assigned a Yammer license.
    yammerLicenseAssignDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewOffice365ActiveUserDetail instantiates a new Office365ActiveUserDetail and sets the default values.
func NewOffice365ActiveUserDetail()(*Office365ActiveUserDetail) {
    m := &Office365ActiveUserDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365ActiveUserDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365ActiveUserDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365ActiveUserDetail(), nil
}
// GetAssignedProducts gets the assignedProducts property value. All the products assigned for the user.
func (m *Office365ActiveUserDetail) GetAssignedProducts()([]string) {
    return m.assignedProducts
}
// GetDeletedDate gets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
func (m *Office365ActiveUserDetail) GetDeletedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.deletedDate
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates.
func (m *Office365ActiveUserDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetExchangeLastActivityDate gets the exchangeLastActivityDate property value. The date when user last read or sent email.
func (m *Office365ActiveUserDetail) GetExchangeLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.exchangeLastActivityDate
}
// GetExchangeLicenseAssignDate gets the exchangeLicenseAssignDate property value. The last date when the user was assigned an Exchange license.
func (m *Office365ActiveUserDetail) GetExchangeLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.exchangeLicenseAssignDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365ActiveUserDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignedProducts(res)
        }
        return nil
    }
    res["deletedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDate(val)
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
    res["exchangeLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastActivityDate(val)
        }
        return nil
    }
    res["exchangeLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLicenseAssignDate(val)
        }
        return nil
    }
    res["hasExchangeLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasExchangeLicense(val)
        }
        return nil
    }
    res["hasOneDriveLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasOneDriveLicense(val)
        }
        return nil
    }
    res["hasSharePointLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasSharePointLicense(val)
        }
        return nil
    }
    res["hasSkypeForBusinessLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasSkypeForBusinessLicense(val)
        }
        return nil
    }
    res["hasTeamsLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasTeamsLicense(val)
        }
        return nil
    }
    res["hasYammerLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasYammerLicense(val)
        }
        return nil
    }
    res["isDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["oneDriveLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveLastActivityDate(val)
        }
        return nil
    }
    res["oneDriveLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveLicenseAssignDate(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePointLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointLastActivityDate(val)
        }
        return nil
    }
    res["sharePointLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointLicenseAssignDate(val)
        }
        return nil
    }
    res["skypeForBusinessLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessLastActivityDate(val)
        }
        return nil
    }
    res["skypeForBusinessLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessLicenseAssignDate(val)
        }
        return nil
    }
    res["teamsLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsLastActivityDate(val)
        }
        return nil
    }
    res["teamsLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsLicenseAssignDate(val)
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
    res["yammerLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLastActivityDate(val)
        }
        return nil
    }
    res["yammerLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLicenseAssignDate(val)
        }
        return nil
    }
    return res
}
// GetHasExchangeLicense gets the hasExchangeLicense property value. Whether the user has been assigned an Exchange license.
func (m *Office365ActiveUserDetail) GetHasExchangeLicense()(*bool) {
    return m.hasExchangeLicense
}
// GetHasOneDriveLicense gets the hasOneDriveLicense property value. Whether the user has been assigned a OneDrive license.
func (m *Office365ActiveUserDetail) GetHasOneDriveLicense()(*bool) {
    return m.hasOneDriveLicense
}
// GetHasSharePointLicense gets the hasSharePointLicense property value. Whether the user has been assigned a SharePoint license.
func (m *Office365ActiveUserDetail) GetHasSharePointLicense()(*bool) {
    return m.hasSharePointLicense
}
// GetHasSkypeForBusinessLicense gets the hasSkypeForBusinessLicense property value. Whether the user has been assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) GetHasSkypeForBusinessLicense()(*bool) {
    return m.hasSkypeForBusinessLicense
}
// GetHasTeamsLicense gets the hasTeamsLicense property value. Whether the user has been assigned a Teams license.
func (m *Office365ActiveUserDetail) GetHasTeamsLicense()(*bool) {
    return m.hasTeamsLicense
}
// GetHasYammerLicense gets the hasYammerLicense property value. Whether the user has been assigned a Yammer license.
func (m *Office365ActiveUserDetail) GetHasYammerLicense()(*bool) {
    return m.hasYammerLicense
}
// GetIsDeleted gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365ActiveUserDetail) GetIsDeleted()(*bool) {
    return m.isDeleted
}
// GetOneDriveLastActivityDate gets the oneDriveLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, or synced files.
func (m *Office365ActiveUserDetail) GetOneDriveLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.oneDriveLastActivityDate
}
// GetOneDriveLicenseAssignDate gets the oneDriveLicenseAssignDate property value. The last date when the user was assigned a OneDrive license.
func (m *Office365ActiveUserDetail) GetOneDriveLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.oneDriveLicenseAssignDate
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365ActiveUserDetail) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportRefreshDate
}
// GetSharePointLastActivityDate gets the sharePointLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
func (m *Office365ActiveUserDetail) GetSharePointLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.sharePointLastActivityDate
}
// GetSharePointLicenseAssignDate gets the sharePointLicenseAssignDate property value. The last date when the user was assigned a SharePoint license.
func (m *Office365ActiveUserDetail) GetSharePointLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.sharePointLicenseAssignDate
}
// GetSkypeForBusinessLastActivityDate gets the skypeForBusinessLastActivityDate property value. The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
func (m *Office365ActiveUserDetail) GetSkypeForBusinessLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.skypeForBusinessLastActivityDate
}
// GetSkypeForBusinessLicenseAssignDate gets the skypeForBusinessLicenseAssignDate property value. The last date when the user was assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) GetSkypeForBusinessLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.skypeForBusinessLicenseAssignDate
}
// GetTeamsLastActivityDate gets the teamsLastActivityDate property value. The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
func (m *Office365ActiveUserDetail) GetTeamsLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.teamsLastActivityDate
}
// GetTeamsLicenseAssignDate gets the teamsLicenseAssignDate property value. The last date when the user was assigned a Teams license.
func (m *Office365ActiveUserDetail) GetTeamsLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.teamsLicenseAssignDate
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *Office365ActiveUserDetail) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetYammerLastActivityDate gets the yammerLastActivityDate property value. The date when user last posted, read, or liked message.
func (m *Office365ActiveUserDetail) GetYammerLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.yammerLastActivityDate
}
// GetYammerLicenseAssignDate gets the yammerLicenseAssignDate property value. The last date when the user was assigned a Yammer license.
func (m *Office365ActiveUserDetail) GetYammerLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.yammerLicenseAssignDate
}
// Serialize serializes information the current object
func (m *Office365ActiveUserDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedProducts() != nil {
        err = writer.WriteCollectionOfStringValues("assignedProducts", m.GetAssignedProducts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("deletedDate", m.GetDeletedDate())
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
        err = writer.WriteDateOnlyValue("exchangeLastActivityDate", m.GetExchangeLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("exchangeLicenseAssignDate", m.GetExchangeLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasExchangeLicense", m.GetHasExchangeLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasOneDriveLicense", m.GetHasOneDriveLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasSharePointLicense", m.GetHasSharePointLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasSkypeForBusinessLicense", m.GetHasSkypeForBusinessLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasTeamsLicense", m.GetHasTeamsLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasYammerLicense", m.GetHasYammerLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("oneDriveLastActivityDate", m.GetOneDriveLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("oneDriveLicenseAssignDate", m.GetOneDriveLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("sharePointLastActivityDate", m.GetSharePointLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("sharePointLicenseAssignDate", m.GetSharePointLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("skypeForBusinessLastActivityDate", m.GetSkypeForBusinessLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("skypeForBusinessLicenseAssignDate", m.GetSkypeForBusinessLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("teamsLastActivityDate", m.GetTeamsLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("teamsLicenseAssignDate", m.GetTeamsLicenseAssignDate())
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
        err = writer.WriteDateOnlyValue("yammerLastActivityDate", m.GetYammerLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("yammerLicenseAssignDate", m.GetYammerLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedProducts sets the assignedProducts property value. All the products assigned for the user.
func (m *Office365ActiveUserDetail) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
// SetDeletedDate sets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
func (m *Office365ActiveUserDetail) SetDeletedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.deletedDate = value
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates.
func (m *Office365ActiveUserDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExchangeLastActivityDate sets the exchangeLastActivityDate property value. The date when user last read or sent email.
func (m *Office365ActiveUserDetail) SetExchangeLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.exchangeLastActivityDate = value
}
// SetExchangeLicenseAssignDate sets the exchangeLicenseAssignDate property value. The last date when the user was assigned an Exchange license.
func (m *Office365ActiveUserDetail) SetExchangeLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.exchangeLicenseAssignDate = value
}
// SetHasExchangeLicense sets the hasExchangeLicense property value. Whether the user has been assigned an Exchange license.
func (m *Office365ActiveUserDetail) SetHasExchangeLicense(value *bool)() {
    m.hasExchangeLicense = value
}
// SetHasOneDriveLicense sets the hasOneDriveLicense property value. Whether the user has been assigned a OneDrive license.
func (m *Office365ActiveUserDetail) SetHasOneDriveLicense(value *bool)() {
    m.hasOneDriveLicense = value
}
// SetHasSharePointLicense sets the hasSharePointLicense property value. Whether the user has been assigned a SharePoint license.
func (m *Office365ActiveUserDetail) SetHasSharePointLicense(value *bool)() {
    m.hasSharePointLicense = value
}
// SetHasSkypeForBusinessLicense sets the hasSkypeForBusinessLicense property value. Whether the user has been assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) SetHasSkypeForBusinessLicense(value *bool)() {
    m.hasSkypeForBusinessLicense = value
}
// SetHasTeamsLicense sets the hasTeamsLicense property value. Whether the user has been assigned a Teams license.
func (m *Office365ActiveUserDetail) SetHasTeamsLicense(value *bool)() {
    m.hasTeamsLicense = value
}
// SetHasYammerLicense sets the hasYammerLicense property value. Whether the user has been assigned a Yammer license.
func (m *Office365ActiveUserDetail) SetHasYammerLicense(value *bool)() {
    m.hasYammerLicense = value
}
// SetIsDeleted sets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365ActiveUserDetail) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetOneDriveLastActivityDate sets the oneDriveLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, or synced files.
func (m *Office365ActiveUserDetail) SetOneDriveLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.oneDriveLastActivityDate = value
}
// SetOneDriveLicenseAssignDate sets the oneDriveLicenseAssignDate property value. The last date when the user was assigned a OneDrive license.
func (m *Office365ActiveUserDetail) SetOneDriveLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.oneDriveLicenseAssignDate = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365ActiveUserDetail) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportRefreshDate = value
}
// SetSharePointLastActivityDate sets the sharePointLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
func (m *Office365ActiveUserDetail) SetSharePointLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.sharePointLastActivityDate = value
}
// SetSharePointLicenseAssignDate sets the sharePointLicenseAssignDate property value. The last date when the user was assigned a SharePoint license.
func (m *Office365ActiveUserDetail) SetSharePointLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.sharePointLicenseAssignDate = value
}
// SetSkypeForBusinessLastActivityDate sets the skypeForBusinessLastActivityDate property value. The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
func (m *Office365ActiveUserDetail) SetSkypeForBusinessLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.skypeForBusinessLastActivityDate = value
}
// SetSkypeForBusinessLicenseAssignDate sets the skypeForBusinessLicenseAssignDate property value. The last date when the user was assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) SetSkypeForBusinessLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.skypeForBusinessLicenseAssignDate = value
}
// SetTeamsLastActivityDate sets the teamsLastActivityDate property value. The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
func (m *Office365ActiveUserDetail) SetTeamsLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.teamsLastActivityDate = value
}
// SetTeamsLicenseAssignDate sets the teamsLicenseAssignDate property value. The last date when the user was assigned a Teams license.
func (m *Office365ActiveUserDetail) SetTeamsLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.teamsLicenseAssignDate = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *Office365ActiveUserDetail) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetYammerLastActivityDate sets the yammerLastActivityDate property value. The date when user last posted, read, or liked message.
func (m *Office365ActiveUserDetail) SetYammerLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.yammerLastActivityDate = value
}
// SetYammerLicenseAssignDate sets the yammerLicenseAssignDate property value. The last date when the user was assigned a Yammer license.
func (m *Office365ActiveUserDetail) SetYammerLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.yammerLicenseAssignDate = value
}
