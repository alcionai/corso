package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppApp 
type MacOsVppApp struct {
    MobileApp
    // The store URL.
    appStoreUrl *string
    // The licenses assigned to this app.
    assignedLicenses []MacOsVppAppAssignedLicenseable
    // The Identity Name.
    bundleId *string
    // The supported License Type.
    licensingType VppLicensingTypeable
    // The VPP application release date and time.
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Results of revoke license actions on this app.
    revokeLicenseActionResults []MacOsVppAppRevokeLicensesActionResultable
    // The total number of VPP licenses.
    totalLicenseCount *int32
    // The number of VPP licenses in use.
    usedLicenseCount *int32
    // Possible types of an Apple Volume Purchase Program token.
    vppTokenAccountType *VppTokenAccountType
    // The Apple Id associated with the given Apple Volume Purchase Program Token.
    vppTokenAppleId *string
    // Identifier of the VPP token associated with this app.
    vppTokenId *string
    // The organization associated with the Apple Volume Purchase Program Token
    vppTokenOrganizationName *string
}
// NewMacOsVppApp instantiates a new MacOsVppApp and sets the default values.
func NewMacOsVppApp()(*MacOsVppApp) {
    m := &MacOsVppApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.macOsVppApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOsVppAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOsVppAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOsVppApp(), nil
}
// GetAppStoreUrl gets the appStoreUrl property value. The store URL.
func (m *MacOsVppApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses assigned to this app.
func (m *MacOsVppApp) GetAssignedLicenses()([]MacOsVppAppAssignedLicenseable) {
    return m.assignedLicenses
}
// GetBundleId gets the bundleId property value. The Identity Name.
func (m *MacOsVppApp) GetBundleId()(*string) {
    return m.bundleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOsVppApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
    res["assignedLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOsVppAppAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOsVppAppAssignedLicenseable, len(val))
            for i, v := range val {
                res[i] = v.(MacOsVppAppAssignedLicenseable)
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
    res["licensingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVppLicensingTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicensingType(val.(VppLicensingTypeable))
        }
        return nil
    }
    res["releaseDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    res["revokeLicenseActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOsVppAppRevokeLicensesActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOsVppAppRevokeLicensesActionResultable, len(val))
            for i, v := range val {
                res[i] = v.(MacOsVppAppRevokeLicensesActionResultable)
            }
            m.SetRevokeLicenseActionResults(res)
        }
        return nil
    }
    res["totalLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenseCount(val)
        }
        return nil
    }
    res["usedLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedLicenseCount(val)
        }
        return nil
    }
    res["vppTokenAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenAccountType(val.(*VppTokenAccountType))
        }
        return nil
    }
    res["vppTokenAppleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenAppleId(val)
        }
        return nil
    }
    res["vppTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenId(val)
        }
        return nil
    }
    res["vppTokenOrganizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenOrganizationName(val)
        }
        return nil
    }
    return res
}
// GetLicensingType gets the licensingType property value. The supported License Type.
func (m *MacOsVppApp) GetLicensingType()(VppLicensingTypeable) {
    return m.licensingType
}
// GetReleaseDateTime gets the releaseDateTime property value. The VPP application release date and time.
func (m *MacOsVppApp) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.releaseDateTime
}
// GetRevokeLicenseActionResults gets the revokeLicenseActionResults property value. Results of revoke license actions on this app.
func (m *MacOsVppApp) GetRevokeLicenseActionResults()([]MacOsVppAppRevokeLicensesActionResultable) {
    return m.revokeLicenseActionResults
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of VPP licenses.
func (m *MacOsVppApp) GetTotalLicenseCount()(*int32) {
    return m.totalLicenseCount
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *MacOsVppApp) GetUsedLicenseCount()(*int32) {
    return m.usedLicenseCount
}
// GetVppTokenAccountType gets the vppTokenAccountType property value. Possible types of an Apple Volume Purchase Program token.
func (m *MacOsVppApp) GetVppTokenAccountType()(*VppTokenAccountType) {
    return m.vppTokenAccountType
}
// GetVppTokenAppleId gets the vppTokenAppleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *MacOsVppApp) GetVppTokenAppleId()(*string) {
    return m.vppTokenAppleId
}
// GetVppTokenId gets the vppTokenId property value. Identifier of the VPP token associated with this app.
func (m *MacOsVppApp) GetVppTokenId()(*string) {
    return m.vppTokenId
}
// GetVppTokenOrganizationName gets the vppTokenOrganizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *MacOsVppApp) GetVppTokenOrganizationName()(*string) {
    return m.vppTokenOrganizationName
}
// Serialize serializes information the current object
func (m *MacOsVppApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAssignedLicenses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bundleId", m.GetBundleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licensingType", m.GetLicensingType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRevokeLicenseActionResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRevokeLicenseActionResults()))
        for i, v := range m.GetRevokeLicenseActionResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("revokeLicenseActionResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicenseCount", m.GetTotalLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usedLicenseCount", m.GetUsedLicenseCount())
        if err != nil {
            return err
        }
    }
    if m.GetVppTokenAccountType() != nil {
        cast := (*m.GetVppTokenAccountType()).String()
        err = writer.WriteStringValue("vppTokenAccountType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vppTokenAppleId", m.GetVppTokenAppleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vppTokenId", m.GetVppTokenId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vppTokenOrganizationName", m.GetVppTokenOrganizationName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppStoreUrl sets the appStoreUrl property value. The store URL.
func (m *MacOsVppApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses assigned to this app.
func (m *MacOsVppApp) SetAssignedLicenses(value []MacOsVppAppAssignedLicenseable)() {
    m.assignedLicenses = value
}
// SetBundleId sets the bundleId property value. The Identity Name.
func (m *MacOsVppApp) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetLicensingType sets the licensingType property value. The supported License Type.
func (m *MacOsVppApp) SetLicensingType(value VppLicensingTypeable)() {
    m.licensingType = value
}
// SetReleaseDateTime sets the releaseDateTime property value. The VPP application release date and time.
func (m *MacOsVppApp) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releaseDateTime = value
}
// SetRevokeLicenseActionResults sets the revokeLicenseActionResults property value. Results of revoke license actions on this app.
func (m *MacOsVppApp) SetRevokeLicenseActionResults(value []MacOsVppAppRevokeLicensesActionResultable)() {
    m.revokeLicenseActionResults = value
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of VPP licenses.
func (m *MacOsVppApp) SetTotalLicenseCount(value *int32)() {
    m.totalLicenseCount = value
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *MacOsVppApp) SetUsedLicenseCount(value *int32)() {
    m.usedLicenseCount = value
}
// SetVppTokenAccountType sets the vppTokenAccountType property value. Possible types of an Apple Volume Purchase Program token.
func (m *MacOsVppApp) SetVppTokenAccountType(value *VppTokenAccountType)() {
    m.vppTokenAccountType = value
}
// SetVppTokenAppleId sets the vppTokenAppleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *MacOsVppApp) SetVppTokenAppleId(value *string)() {
    m.vppTokenAppleId = value
}
// SetVppTokenId sets the vppTokenId property value. Identifier of the VPP token associated with this app.
func (m *MacOsVppApp) SetVppTokenId(value *string)() {
    m.vppTokenId = value
}
// SetVppTokenOrganizationName sets the vppTokenOrganizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *MacOsVppApp) SetVppTokenOrganizationName(value *string)() {
    m.vppTokenOrganizationName = value
}
