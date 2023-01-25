package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerVpnConfiguration 
type AndroidDeviceOwnerVpnConfiguration struct {
    VpnConfiguration
    // Whether or not to enable always-on VPN connection.
    alwaysOn *bool
    // If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
    alwaysOnLockdown *bool
    // Android VPN connection type.
    connectionType *AndroidVpnConnectionType
    // Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
    customData []KeyValueable
    // Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
    customKeyValueData []KeyValuePairable
    // Tenant level settings for the Derived Credentials to be used for authentication.
    derivedCredentialSettings DeviceManagementDerivedCredentialSettingsable
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate AndroidDeviceOwnerCertificateProfileBaseable
    // Microsoft Tunnel site ID.
    microsoftTunnelSiteId *string
    // Proxy server.
    proxyServer VpnProxyServerable
    // Targeted mobile apps. This collection can contain a maximum of 500 elements.
    targetedMobileApps []AppListItemable
    // Targeted App package IDs.
    targetedPackageIds []string
}
// NewAndroidDeviceOwnerVpnConfiguration instantiates a new AndroidDeviceOwnerVpnConfiguration and sets the default values.
func NewAndroidDeviceOwnerVpnConfiguration()(*AndroidDeviceOwnerVpnConfiguration) {
    m := &AndroidDeviceOwnerVpnConfiguration{
        VpnConfiguration: *NewVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerVpnConfiguration(), nil
}
// GetAlwaysOn gets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidDeviceOwnerVpnConfiguration) GetAlwaysOn()(*bool) {
    return m.alwaysOn
}
// GetAlwaysOnLockdown gets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidDeviceOwnerVpnConfiguration) GetAlwaysOnLockdown()(*bool) {
    return m.alwaysOnLockdown
}
// GetConnectionType gets the connectionType property value. Android VPN connection type.
func (m *AndroidDeviceOwnerVpnConfiguration) GetConnectionType()(*AndroidVpnConnectionType) {
    return m.connectionType
}
// GetCustomData gets the customData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) GetCustomData()([]KeyValueable) {
    return m.customData
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    return m.customKeyValueData
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *AndroidDeviceOwnerVpnConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    return m.derivedCredentialSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VpnConfiguration.GetFieldDeserializers()
    res["alwaysOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlwaysOn(val)
        }
        return nil
    }
    res["alwaysOnLockdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlwaysOnLockdown(val)
        }
        return nil
    }
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*AndroidVpnConnectionType))
        }
        return nil
    }
    res["customData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValueable)
            }
            m.SetCustomData(res)
        }
        return nil
    }
    res["customKeyValueData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetCustomKeyValueData(res)
        }
        return nil
    }
    res["derivedCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialSettings(val.(DeviceManagementDerivedCredentialSettingsable))
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidDeviceOwnerCertificateProfileBaseable))
        }
        return nil
    }
    res["microsoftTunnelSiteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftTunnelSiteId(val)
        }
        return nil
    }
    res["proxyServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVpnProxyServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyServer(val.(VpnProxyServerable))
        }
        return nil
    }
    res["targetedMobileApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetTargetedMobileApps(res)
        }
        return nil
    }
    res["targetedPackageIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTargetedPackageIds(res)
        }
        return nil
    }
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidDeviceOwnerVpnConfiguration) GetIdentityCertificate()(AndroidDeviceOwnerCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidDeviceOwnerVpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    return m.microsoftTunnelSiteId
}
// GetProxyServer gets the proxyServer property value. Proxy server.
func (m *AndroidDeviceOwnerVpnConfiguration) GetProxyServer()(VpnProxyServerable) {
    return m.proxyServer
}
// GetTargetedMobileApps gets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) GetTargetedMobileApps()([]AppListItemable) {
    return m.targetedMobileApps
}
// GetTargetedPackageIds gets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidDeviceOwnerVpnConfiguration) GetTargetedPackageIds()([]string) {
    return m.targetedPackageIds
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("alwaysOn", m.GetAlwaysOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("alwaysOnLockdown", m.GetAlwaysOnLockdown())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomData()))
        for i, v := range m.GetCustomData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customData", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomKeyValueData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomKeyValueData()))
        for i, v := range m.GetCustomKeyValueData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customKeyValueData", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("derivedCredentialSettings", m.GetDerivedCredentialSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftTunnelSiteId", m.GetMicrosoftTunnelSiteId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("proxyServer", m.GetProxyServer())
        if err != nil {
            return err
        }
    }
    if m.GetTargetedMobileApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedMobileApps()))
        for i, v := range m.GetTargetedMobileApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("targetedMobileApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedPackageIds() != nil {
        err = writer.WriteCollectionOfStringValues("targetedPackageIds", m.GetTargetedPackageIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlwaysOn sets the alwaysOn property value. Whether or not to enable always-on VPN connection.
func (m *AndroidDeviceOwnerVpnConfiguration) SetAlwaysOn(value *bool)() {
    m.alwaysOn = value
}
// SetAlwaysOnLockdown sets the alwaysOnLockdown property value. If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidDeviceOwnerVpnConfiguration) SetAlwaysOnLockdown(value *bool)() {
    m.alwaysOnLockdown = value
}
// SetConnectionType sets the connectionType property value. Android VPN connection type.
func (m *AndroidDeviceOwnerVpnConfiguration) SetConnectionType(value *AndroidVpnConnectionType)() {
    m.connectionType = value
}
// SetCustomData sets the customData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) SetCustomData(value []KeyValueable)() {
    m.customData = value
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    m.customKeyValueData = value
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *AndroidDeviceOwnerVpnConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    m.derivedCredentialSettings = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidDeviceOwnerVpnConfiguration) SetIdentityCertificate(value AndroidDeviceOwnerCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. Microsoft Tunnel site ID.
func (m *AndroidDeviceOwnerVpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    m.microsoftTunnelSiteId = value
}
// SetProxyServer sets the proxyServer property value. Proxy server.
func (m *AndroidDeviceOwnerVpnConfiguration) SetProxyServer(value VpnProxyServerable)() {
    m.proxyServer = value
}
// SetTargetedMobileApps sets the targetedMobileApps property value. Targeted mobile apps. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerVpnConfiguration) SetTargetedMobileApps(value []AppListItemable)() {
    m.targetedMobileApps = value
}
// SetTargetedPackageIds sets the targetedPackageIds property value. Targeted App package IDs.
func (m *AndroidDeviceOwnerVpnConfiguration) SetTargetedPackageIds(value []string)() {
    m.targetedPackageIds = value
}
