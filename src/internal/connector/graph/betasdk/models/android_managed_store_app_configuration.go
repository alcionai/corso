package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppConfiguration 
type AndroidManagedStoreAppConfiguration struct {
    ManagedDeviceMobileAppConfiguration
    // Whether or not this AppConfig is an OEMConfig policy.
    appSupportsOemConfig *bool
    // Setting to specify whether to allow ConnectedApps experience for this app.
    connectedAppsEnabled *bool
    // Android Enterprise app configuration package id.
    packageId *string
    // Android Enterprise app configuration JSON payload.
    payloadJson *string
    // List of Android app permissions and corresponding permission actions.
    permissionActions []AndroidPermissionActionable
    // Android profile applicability
    profileApplicability *AndroidProfileApplicability
}
// NewAndroidManagedStoreAppConfiguration instantiates a new AndroidManagedStoreAppConfiguration and sets the default values.
func NewAndroidManagedStoreAppConfiguration()(*AndroidManagedStoreAppConfiguration) {
    m := &AndroidManagedStoreAppConfiguration{
        ManagedDeviceMobileAppConfiguration: *NewManagedDeviceMobileAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreAppConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidManagedStoreAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAppConfiguration(), nil
}
// GetAppSupportsOemConfig gets the appSupportsOemConfig property value. Whether or not this AppConfig is an OEMConfig policy.
func (m *AndroidManagedStoreAppConfiguration) GetAppSupportsOemConfig()(*bool) {
    return m.appSupportsOemConfig
}
// GetConnectedAppsEnabled gets the connectedAppsEnabled property value. Setting to specify whether to allow ConnectedApps experience for this app.
func (m *AndroidManagedStoreAppConfiguration) GetConnectedAppsEnabled()(*bool) {
    return m.connectedAppsEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedDeviceMobileAppConfiguration.GetFieldDeserializers()
    res["appSupportsOemConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppSupportsOemConfig(val)
        }
        return nil
    }
    res["connectedAppsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedAppsEnabled(val)
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
    res["payloadJson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadJson(val)
        }
        return nil
    }
    res["permissionActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidPermissionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidPermissionActionable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidPermissionActionable)
            }
            m.SetPermissionActions(res)
        }
        return nil
    }
    res["profileApplicability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidProfileApplicability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileApplicability(val.(*AndroidProfileApplicability))
        }
        return nil
    }
    return res
}
// GetPackageId gets the packageId property value. Android Enterprise app configuration package id.
func (m *AndroidManagedStoreAppConfiguration) GetPackageId()(*string) {
    return m.packageId
}
// GetPayloadJson gets the payloadJson property value. Android Enterprise app configuration JSON payload.
func (m *AndroidManagedStoreAppConfiguration) GetPayloadJson()(*string) {
    return m.payloadJson
}
// GetPermissionActions gets the permissionActions property value. List of Android app permissions and corresponding permission actions.
func (m *AndroidManagedStoreAppConfiguration) GetPermissionActions()([]AndroidPermissionActionable) {
    return m.permissionActions
}
// GetProfileApplicability gets the profileApplicability property value. Android profile applicability
func (m *AndroidManagedStoreAppConfiguration) GetProfileApplicability()(*AndroidProfileApplicability) {
    return m.profileApplicability
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAppConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedDeviceMobileAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appSupportsOemConfig", m.GetAppSupportsOemConfig())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connectedAppsEnabled", m.GetConnectedAppsEnabled())
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
        err = writer.WriteStringValue("payloadJson", m.GetPayloadJson())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionActions()))
        for i, v := range m.GetPermissionActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("permissionActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProfileApplicability() != nil {
        cast := (*m.GetProfileApplicability()).String()
        err = writer.WriteStringValue("profileApplicability", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppSupportsOemConfig sets the appSupportsOemConfig property value. Whether or not this AppConfig is an OEMConfig policy.
func (m *AndroidManagedStoreAppConfiguration) SetAppSupportsOemConfig(value *bool)() {
    m.appSupportsOemConfig = value
}
// SetConnectedAppsEnabled sets the connectedAppsEnabled property value. Setting to specify whether to allow ConnectedApps experience for this app.
func (m *AndroidManagedStoreAppConfiguration) SetConnectedAppsEnabled(value *bool)() {
    m.connectedAppsEnabled = value
}
// SetPackageId sets the packageId property value. Android Enterprise app configuration package id.
func (m *AndroidManagedStoreAppConfiguration) SetPackageId(value *string)() {
    m.packageId = value
}
// SetPayloadJson sets the payloadJson property value. Android Enterprise app configuration JSON payload.
func (m *AndroidManagedStoreAppConfiguration) SetPayloadJson(value *string)() {
    m.payloadJson = value
}
// SetPermissionActions sets the permissionActions property value. List of Android app permissions and corresponding permission actions.
func (m *AndroidManagedStoreAppConfiguration) SetPermissionActions(value []AndroidPermissionActionable)() {
    m.permissionActions = value
}
// SetProfileApplicability sets the profileApplicability property value. Android profile applicability
func (m *AndroidManagedStoreAppConfiguration) SetProfileApplicability(value *AndroidProfileApplicability)() {
    m.profileApplicability = value
}
