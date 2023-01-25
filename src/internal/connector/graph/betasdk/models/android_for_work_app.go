package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkApp 
type AndroidForWorkApp struct {
    MobileApp
    // The Identity Name.
    appIdentifier *string
    // The Play for Work Store app URL.
    appStoreUrl *string
    // The package identifier.
    packageId *string
    // The total number of VPP licenses.
    totalLicenseCount *int32
    // The number of VPP licenses in use.
    usedLicenseCount *int32
}
// NewAndroidForWorkApp instantiates a new AndroidForWorkApp and sets the default values.
func NewAndroidForWorkApp()(*AndroidForWorkApp) {
    m := &AndroidForWorkApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidForWorkAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkApp(), nil
}
// GetAppIdentifier gets the appIdentifier property value. The Identity Name.
func (m *AndroidForWorkApp) GetAppIdentifier()(*string) {
    return m.appIdentifier
}
// GetAppStoreUrl gets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidForWorkApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppIdentifier(val)
        }
        return nil
    }
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
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
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
// GetPackageId gets the packageId property value. The package identifier.
func (m *AndroidForWorkApp) GetPackageId()(*string) {
    return m.packageId
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidForWorkApp) GetTotalLicenseCount()(*int32) {
    return m.totalLicenseCount
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidForWorkApp) GetUsedLicenseCount()(*int32) {
    return m.usedLicenseCount
}
// Serialize serializes information the current object
func (m *AndroidForWorkApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appIdentifier", m.GetAppIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageId", m.GetPackageId())
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
// SetAppIdentifier sets the appIdentifier property value. The Identity Name.
func (m *AndroidForWorkApp) SetAppIdentifier(value *string)() {
    m.appIdentifier = value
}
// SetAppStoreUrl sets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidForWorkApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *AndroidForWorkApp) SetPackageId(value *string)() {
    m.packageId = value
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidForWorkApp) SetTotalLicenseCount(value *int32)() {
    m.totalLicenseCount = value
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidForWorkApp) SetUsedLicenseCount(value *int32)() {
    m.usedLicenseCount = value
}
