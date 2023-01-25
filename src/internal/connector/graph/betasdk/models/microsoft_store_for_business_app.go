package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftStoreForBusinessApp 
type MicrosoftStoreForBusinessApp struct {
    MobileApp
    // The collection of contained apps in a mobileApp acting as a package.
    containedApps []MobileContainedAppable
    // The licenseType property
    licenseType *MicrosoftStoreForBusinessLicenseType
    // The supported License Type.
    licensingType VppLicensingTypeable
    // The app package identifier
    packageIdentityName *string
    // The app product key
    productKey *string
    // The total number of Microsoft Store for Business licenses.
    totalLicenseCount *int32
    // The number of Microsoft Store for Business licenses in use.
    usedLicenseCount *int32
}
// NewMicrosoftStoreForBusinessApp instantiates a new MicrosoftStoreForBusinessApp and sets the default values.
func NewMicrosoftStoreForBusinessApp()(*MicrosoftStoreForBusinessApp) {
    m := &MicrosoftStoreForBusinessApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.microsoftStoreForBusinessApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftStoreForBusinessApp(), nil
}
// GetContainedApps gets the containedApps property value. The collection of contained apps in a mobileApp acting as a package.
func (m *MicrosoftStoreForBusinessApp) GetContainedApps()([]MobileContainedAppable) {
    return m.containedApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftStoreForBusinessApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["containedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileContainedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileContainedAppable, len(val))
            for i, v := range val {
                res[i] = v.(MobileContainedAppable)
            }
            m.SetContainedApps(res)
        }
        return nil
    }
    res["licenseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftStoreForBusinessLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val.(*MicrosoftStoreForBusinessLicenseType))
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
    res["packageIdentityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageIdentityName(val)
        }
        return nil
    }
    res["productKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
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
    return res
}
// GetLicenseType gets the licenseType property value. The licenseType property
func (m *MicrosoftStoreForBusinessApp) GetLicenseType()(*MicrosoftStoreForBusinessLicenseType) {
    return m.licenseType
}
// GetLicensingType gets the licensingType property value. The supported License Type.
func (m *MicrosoftStoreForBusinessApp) GetLicensingType()(VppLicensingTypeable) {
    return m.licensingType
}
// GetPackageIdentityName gets the packageIdentityName property value. The app package identifier
func (m *MicrosoftStoreForBusinessApp) GetPackageIdentityName()(*string) {
    return m.packageIdentityName
}
// GetProductKey gets the productKey property value. The app product key
func (m *MicrosoftStoreForBusinessApp) GetProductKey()(*string) {
    return m.productKey
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of Microsoft Store for Business licenses.
func (m *MicrosoftStoreForBusinessApp) GetTotalLicenseCount()(*int32) {
    return m.totalLicenseCount
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of Microsoft Store for Business licenses in use.
func (m *MicrosoftStoreForBusinessApp) GetUsedLicenseCount()(*int32) {
    return m.usedLicenseCount
}
// Serialize serializes information the current object
func (m *MicrosoftStoreForBusinessApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContainedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContainedApps()))
        for i, v := range m.GetContainedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("containedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseType() != nil {
        cast := (*m.GetLicenseType()).String()
        err = writer.WriteStringValue("licenseType", &cast)
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
        err = writer.WriteStringValue("packageIdentityName", m.GetPackageIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
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
    return nil
}
// SetContainedApps sets the containedApps property value. The collection of contained apps in a mobileApp acting as a package.
func (m *MicrosoftStoreForBusinessApp) SetContainedApps(value []MobileContainedAppable)() {
    m.containedApps = value
}
// SetLicenseType sets the licenseType property value. The licenseType property
func (m *MicrosoftStoreForBusinessApp) SetLicenseType(value *MicrosoftStoreForBusinessLicenseType)() {
    m.licenseType = value
}
// SetLicensingType sets the licensingType property value. The supported License Type.
func (m *MicrosoftStoreForBusinessApp) SetLicensingType(value VppLicensingTypeable)() {
    m.licensingType = value
}
// SetPackageIdentityName sets the packageIdentityName property value. The app package identifier
func (m *MicrosoftStoreForBusinessApp) SetPackageIdentityName(value *string)() {
    m.packageIdentityName = value
}
// SetProductKey sets the productKey property value. The app product key
func (m *MicrosoftStoreForBusinessApp) SetProductKey(value *string)() {
    m.productKey = value
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of Microsoft Store for Business licenses.
func (m *MicrosoftStoreForBusinessApp) SetTotalLicenseCount(value *int32)() {
    m.totalLicenseCount = value
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of Microsoft Store for Business licenses in use.
func (m *MicrosoftStoreForBusinessApp) SetUsedLicenseCount(value *int32)() {
    m.usedLicenseCount = value
}
