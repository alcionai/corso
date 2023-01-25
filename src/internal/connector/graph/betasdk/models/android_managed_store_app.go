package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreApp 
type AndroidManagedStoreApp struct {
    MobileApp
    // The Identity Name.
    appIdentifier *string
    // The Play for Work Store app URL.
    appStoreUrl *string
    // The tracks that are visible to this enterprise.
    appTracks []AndroidManagedStoreAppTrackable
    // Indicates whether the app is only available to a given enterprise's users.
    isPrivate *bool
    // Indicates whether the app is a preinstalled system app.
    isSystemApp *bool
    // The package identifier.
    packageId *string
    // Whether this app supports OEMConfig policy.
    supportsOemConfig *bool
    // The total number of VPP licenses.
    totalLicenseCount *int32
    // The number of VPP licenses in use.
    usedLicenseCount *int32
}
// NewAndroidManagedStoreApp instantiates a new AndroidManagedStoreApp and sets the default values.
func NewAndroidManagedStoreApp()(*AndroidManagedStoreApp) {
    m := &AndroidManagedStoreApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidManagedStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidManagedStoreWebApp":
                        return NewAndroidManagedStoreWebApp(), nil
                }
            }
        }
    }
    return NewAndroidManagedStoreApp(), nil
}
// GetAppIdentifier gets the appIdentifier property value. The Identity Name.
func (m *AndroidManagedStoreApp) GetAppIdentifier()(*string) {
    return m.appIdentifier
}
// GetAppStoreUrl gets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidManagedStoreApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetAppTracks gets the appTracks property value. The tracks that are visible to this enterprise.
func (m *AndroidManagedStoreApp) GetAppTracks()([]AndroidManagedStoreAppTrackable) {
    return m.appTracks
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["appTracks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedStoreAppTrackFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppTrackable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidManagedStoreAppTrackable)
            }
            m.SetAppTracks(res)
        }
        return nil
    }
    res["isPrivate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivate(val)
        }
        return nil
    }
    res["isSystemApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSystemApp(val)
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
    res["supportsOemConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsOemConfig(val)
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
// GetIsPrivate gets the isPrivate property value. Indicates whether the app is only available to a given enterprise's users.
func (m *AndroidManagedStoreApp) GetIsPrivate()(*bool) {
    return m.isPrivate
}
// GetIsSystemApp gets the isSystemApp property value. Indicates whether the app is a preinstalled system app.
func (m *AndroidManagedStoreApp) GetIsSystemApp()(*bool) {
    return m.isSystemApp
}
// GetPackageId gets the packageId property value. The package identifier.
func (m *AndroidManagedStoreApp) GetPackageId()(*string) {
    return m.packageId
}
// GetSupportsOemConfig gets the supportsOemConfig property value. Whether this app supports OEMConfig policy.
func (m *AndroidManagedStoreApp) GetSupportsOemConfig()(*bool) {
    return m.supportsOemConfig
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidManagedStoreApp) GetTotalLicenseCount()(*int32) {
    return m.totalLicenseCount
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidManagedStoreApp) GetUsedLicenseCount()(*int32) {
    return m.usedLicenseCount
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetAppTracks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppTracks()))
        for i, v := range m.GetAppTracks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appTracks", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPrivate", m.GetIsPrivate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSystemApp", m.GetIsSystemApp())
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
        err = writer.WriteBoolValue("supportsOemConfig", m.GetSupportsOemConfig())
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
func (m *AndroidManagedStoreApp) SetAppIdentifier(value *string)() {
    m.appIdentifier = value
}
// SetAppStoreUrl sets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidManagedStoreApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
// SetAppTracks sets the appTracks property value. The tracks that are visible to this enterprise.
func (m *AndroidManagedStoreApp) SetAppTracks(value []AndroidManagedStoreAppTrackable)() {
    m.appTracks = value
}
// SetIsPrivate sets the isPrivate property value. Indicates whether the app is only available to a given enterprise's users.
func (m *AndroidManagedStoreApp) SetIsPrivate(value *bool)() {
    m.isPrivate = value
}
// SetIsSystemApp sets the isSystemApp property value. Indicates whether the app is a preinstalled system app.
func (m *AndroidManagedStoreApp) SetIsSystemApp(value *bool)() {
    m.isSystemApp = value
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *AndroidManagedStoreApp) SetPackageId(value *string)() {
    m.packageId = value
}
// SetSupportsOemConfig sets the supportsOemConfig property value. Whether this app supports OEMConfig policy.
func (m *AndroidManagedStoreApp) SetSupportsOemConfig(value *bool)() {
    m.supportsOemConfig = value
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidManagedStoreApp) SetTotalLicenseCount(value *int32)() {
    m.totalLicenseCount = value
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidManagedStoreApp) SetUsedLicenseCount(value *int32)() {
    m.usedLicenseCount = value
}
