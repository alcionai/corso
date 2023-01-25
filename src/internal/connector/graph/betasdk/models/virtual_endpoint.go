package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpoint 
type VirtualEndpoint struct {
    Entity
    // Cloud PC audit event.
    auditEvents []CloudPcAuditEventable
    // Cloud managed virtual desktops.
    cloudPCs []CloudPCable
    // Cloud PC organization mapping between public and US Government Community Cloud (GCC) organizations.
    crossCloudGovernmentOrganizationMapping CloudPcCrossCloudGovernmentOrganizationMappingable
    // The image resource on Cloud PC.
    deviceImages []CloudPcDeviceImageable
    // The external partner settings on a Cloud PC.
    externalPartnerSettings []CloudPcExternalPartnerSettingable
    // The gallery image resource on Cloud PC.
    galleryImages []CloudPcGalleryImageable
    // A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
    onPremisesConnections []CloudPcOnPremisesConnectionable
    // The Cloud PC organization settings for a tenant.
    organizationSettings CloudPcOrganizationSettingsable
    // Cloud PC provisioning policy.
    provisioningPolicies []CloudPcProvisioningPolicyable
    // Cloud PC related reports.
    reports CloudPcReportsable
    // Cloud PC service plans.
    servicePlans []CloudPcServicePlanable
    // The sharedUseServicePlans property
    sharedUseServicePlans []CloudPcSharedUseServicePlanable
    // Cloud PC snapshots.
    snapshots []CloudPcSnapshotable
    // Cloud PC supported regions.
    supportedRegions []CloudPcSupportedRegionable
    // Cloud PC user settings.
    userSettings []CloudPcUserSettingable
}
// NewVirtualEndpoint instantiates a new virtualEndpoint and sets the default values.
func NewVirtualEndpoint()(*VirtualEndpoint) {
    m := &VirtualEndpoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpoint(), nil
}
// GetAuditEvents gets the auditEvents property value. Cloud PC audit event.
func (m *VirtualEndpoint) GetAuditEvents()([]CloudPcAuditEventable) {
    return m.auditEvents
}
// GetCloudPCs gets the cloudPCs property value. Cloud managed virtual desktops.
func (m *VirtualEndpoint) GetCloudPCs()([]CloudPCable) {
    return m.cloudPCs
}
// GetCrossCloudGovernmentOrganizationMapping gets the crossCloudGovernmentOrganizationMapping property value. Cloud PC organization mapping between public and US Government Community Cloud (GCC) organizations.
func (m *VirtualEndpoint) GetCrossCloudGovernmentOrganizationMapping()(CloudPcCrossCloudGovernmentOrganizationMappingable) {
    return m.crossCloudGovernmentOrganizationMapping
}
// GetDeviceImages gets the deviceImages property value. The image resource on Cloud PC.
func (m *VirtualEndpoint) GetDeviceImages()([]CloudPcDeviceImageable) {
    return m.deviceImages
}
// GetExternalPartnerSettings gets the externalPartnerSettings property value. The external partner settings on a Cloud PC.
func (m *VirtualEndpoint) GetExternalPartnerSettings()([]CloudPcExternalPartnerSettingable) {
    return m.externalPartnerSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["auditEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcAuditEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcAuditEventable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcAuditEventable)
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["cloudPCs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPCFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPCable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPCable)
            }
            m.SetCloudPCs(res)
        }
        return nil
    }
    res["crossCloudGovernmentOrganizationMapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossCloudGovernmentOrganizationMapping(val.(CloudPcCrossCloudGovernmentOrganizationMappingable))
        }
        return nil
    }
    res["deviceImages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcDeviceImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDeviceImageable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcDeviceImageable)
            }
            m.SetDeviceImages(res)
        }
        return nil
    }
    res["externalPartnerSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcExternalPartnerSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcExternalPartnerSettingable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcExternalPartnerSettingable)
            }
            m.SetExternalPartnerSettings(res)
        }
        return nil
    }
    res["galleryImages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcGalleryImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcGalleryImageable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcGalleryImageable)
            }
            m.SetGalleryImages(res)
        }
        return nil
    }
    res["onPremisesConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcOnPremisesConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcOnPremisesConnectionable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcOnPremisesConnectionable)
            }
            m.SetOnPremisesConnections(res)
        }
        return nil
    }
    res["organizationSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcOrganizationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationSettings(val.(CloudPcOrganizationSettingsable))
        }
        return nil
    }
    res["provisioningPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcProvisioningPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcProvisioningPolicyable)
            }
            m.SetProvisioningPolicies(res)
        }
        return nil
    }
    res["reports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcReportsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReports(val.(CloudPcReportsable))
        }
        return nil
    }
    res["servicePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcServicePlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcServicePlanable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcServicePlanable)
            }
            m.SetServicePlans(res)
        }
        return nil
    }
    res["sharedUseServicePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcSharedUseServicePlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcSharedUseServicePlanable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcSharedUseServicePlanable)
            }
            m.SetSharedUseServicePlans(res)
        }
        return nil
    }
    res["snapshots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcSnapshotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcSnapshotable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcSnapshotable)
            }
            m.SetSnapshots(res)
        }
        return nil
    }
    res["supportedRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcSupportedRegionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcSupportedRegionable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcSupportedRegionable)
            }
            m.SetSupportedRegions(res)
        }
        return nil
    }
    res["userSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcUserSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcUserSettingable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcUserSettingable)
            }
            m.SetUserSettings(res)
        }
        return nil
    }
    return res
}
// GetGalleryImages gets the galleryImages property value. The gallery image resource on Cloud PC.
func (m *VirtualEndpoint) GetGalleryImages()([]CloudPcGalleryImageable) {
    return m.galleryImages
}
// GetOnPremisesConnections gets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *VirtualEndpoint) GetOnPremisesConnections()([]CloudPcOnPremisesConnectionable) {
    return m.onPremisesConnections
}
// GetOrganizationSettings gets the organizationSettings property value. The Cloud PC organization settings for a tenant.
func (m *VirtualEndpoint) GetOrganizationSettings()(CloudPcOrganizationSettingsable) {
    return m.organizationSettings
}
// GetProvisioningPolicies gets the provisioningPolicies property value. Cloud PC provisioning policy.
func (m *VirtualEndpoint) GetProvisioningPolicies()([]CloudPcProvisioningPolicyable) {
    return m.provisioningPolicies
}
// GetReports gets the reports property value. Cloud PC related reports.
func (m *VirtualEndpoint) GetReports()(CloudPcReportsable) {
    return m.reports
}
// GetServicePlans gets the servicePlans property value. Cloud PC service plans.
func (m *VirtualEndpoint) GetServicePlans()([]CloudPcServicePlanable) {
    return m.servicePlans
}
// GetSharedUseServicePlans gets the sharedUseServicePlans property value. The sharedUseServicePlans property
func (m *VirtualEndpoint) GetSharedUseServicePlans()([]CloudPcSharedUseServicePlanable) {
    return m.sharedUseServicePlans
}
// GetSnapshots gets the snapshots property value. Cloud PC snapshots.
func (m *VirtualEndpoint) GetSnapshots()([]CloudPcSnapshotable) {
    return m.snapshots
}
// GetSupportedRegions gets the supportedRegions property value. Cloud PC supported regions.
func (m *VirtualEndpoint) GetSupportedRegions()([]CloudPcSupportedRegionable) {
    return m.supportedRegions
}
// GetUserSettings gets the userSettings property value. Cloud PC user settings.
func (m *VirtualEndpoint) GetUserSettings()([]CloudPcUserSettingable) {
    return m.userSettings
}
// Serialize serializes information the current object
func (m *VirtualEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuditEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPCs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPCs()))
        for i, v := range m.GetCloudPCs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPCs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("crossCloudGovernmentOrganizationMapping", m.GetCrossCloudGovernmentOrganizationMapping())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceImages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceImages()))
        for i, v := range m.GetDeviceImages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceImages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExternalPartnerSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExternalPartnerSettings()))
        for i, v := range m.GetExternalPartnerSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("externalPartnerSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGalleryImages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGalleryImages()))
        for i, v := range m.GetGalleryImages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("galleryImages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesConnections()))
        for i, v := range m.GetOnPremisesConnections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesConnections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizationSettings", m.GetOrganizationSettings())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisioningPolicies()))
        for i, v := range m.GetProvisioningPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("provisioningPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reports", m.GetReports())
        if err != nil {
            return err
        }
    }
    if m.GetServicePlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePlans()))
        for i, v := range m.GetServicePlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("servicePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharedUseServicePlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedUseServicePlans()))
        for i, v := range m.GetSharedUseServicePlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sharedUseServicePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSnapshots() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSnapshots()))
        for i, v := range m.GetSnapshots() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("snapshots", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedRegions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSupportedRegions()))
        for i, v := range m.GetSupportedRegions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("supportedRegions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSettings()))
        for i, v := range m.GetUserSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuditEvents sets the auditEvents property value. Cloud PC audit event.
func (m *VirtualEndpoint) SetAuditEvents(value []CloudPcAuditEventable)() {
    m.auditEvents = value
}
// SetCloudPCs sets the cloudPCs property value. Cloud managed virtual desktops.
func (m *VirtualEndpoint) SetCloudPCs(value []CloudPCable)() {
    m.cloudPCs = value
}
// SetCrossCloudGovernmentOrganizationMapping sets the crossCloudGovernmentOrganizationMapping property value. Cloud PC organization mapping between public and US Government Community Cloud (GCC) organizations.
func (m *VirtualEndpoint) SetCrossCloudGovernmentOrganizationMapping(value CloudPcCrossCloudGovernmentOrganizationMappingable)() {
    m.crossCloudGovernmentOrganizationMapping = value
}
// SetDeviceImages sets the deviceImages property value. The image resource on Cloud PC.
func (m *VirtualEndpoint) SetDeviceImages(value []CloudPcDeviceImageable)() {
    m.deviceImages = value
}
// SetExternalPartnerSettings sets the externalPartnerSettings property value. The external partner settings on a Cloud PC.
func (m *VirtualEndpoint) SetExternalPartnerSettings(value []CloudPcExternalPartnerSettingable)() {
    m.externalPartnerSettings = value
}
// SetGalleryImages sets the galleryImages property value. The gallery image resource on Cloud PC.
func (m *VirtualEndpoint) SetGalleryImages(value []CloudPcGalleryImageable)() {
    m.galleryImages = value
}
// SetOnPremisesConnections sets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *VirtualEndpoint) SetOnPremisesConnections(value []CloudPcOnPremisesConnectionable)() {
    m.onPremisesConnections = value
}
// SetOrganizationSettings sets the organizationSettings property value. The Cloud PC organization settings for a tenant.
func (m *VirtualEndpoint) SetOrganizationSettings(value CloudPcOrganizationSettingsable)() {
    m.organizationSettings = value
}
// SetProvisioningPolicies sets the provisioningPolicies property value. Cloud PC provisioning policy.
func (m *VirtualEndpoint) SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)() {
    m.provisioningPolicies = value
}
// SetReports sets the reports property value. Cloud PC related reports.
func (m *VirtualEndpoint) SetReports(value CloudPcReportsable)() {
    m.reports = value
}
// SetServicePlans sets the servicePlans property value. Cloud PC service plans.
func (m *VirtualEndpoint) SetServicePlans(value []CloudPcServicePlanable)() {
    m.servicePlans = value
}
// SetSharedUseServicePlans sets the sharedUseServicePlans property value. The sharedUseServicePlans property
func (m *VirtualEndpoint) SetSharedUseServicePlans(value []CloudPcSharedUseServicePlanable)() {
    m.sharedUseServicePlans = value
}
// SetSnapshots sets the snapshots property value. Cloud PC snapshots.
func (m *VirtualEndpoint) SetSnapshots(value []CloudPcSnapshotable)() {
    m.snapshots = value
}
// SetSupportedRegions sets the supportedRegions property value. Cloud PC supported regions.
func (m *VirtualEndpoint) SetSupportedRegions(value []CloudPcSupportedRegionable)() {
    m.supportedRegions = value
}
// SetUserSettings sets the userSettings property value. Cloud PC user settings.
func (m *VirtualEndpoint) SetUserSettings(value []CloudPcUserSettingable)() {
    m.userSettings = value
}
