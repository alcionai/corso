package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpointable 
type VirtualEndpointable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuditEvents()([]CloudPcAuditEventable)
    GetCloudPCs()([]CloudPCable)
    GetCrossCloudGovernmentOrganizationMapping()(CloudPcCrossCloudGovernmentOrganizationMappingable)
    GetDeviceImages()([]CloudPcDeviceImageable)
    GetExternalPartnerSettings()([]CloudPcExternalPartnerSettingable)
    GetGalleryImages()([]CloudPcGalleryImageable)
    GetOnPremisesConnections()([]CloudPcOnPremisesConnectionable)
    GetOrganizationSettings()(CloudPcOrganizationSettingsable)
    GetProvisioningPolicies()([]CloudPcProvisioningPolicyable)
    GetReports()(CloudPcReportsable)
    GetServicePlans()([]CloudPcServicePlanable)
    GetSharedUseServicePlans()([]CloudPcSharedUseServicePlanable)
    GetSnapshots()([]CloudPcSnapshotable)
    GetSupportedRegions()([]CloudPcSupportedRegionable)
    GetUserSettings()([]CloudPcUserSettingable)
    SetAuditEvents(value []CloudPcAuditEventable)()
    SetCloudPCs(value []CloudPCable)()
    SetCrossCloudGovernmentOrganizationMapping(value CloudPcCrossCloudGovernmentOrganizationMappingable)()
    SetDeviceImages(value []CloudPcDeviceImageable)()
    SetExternalPartnerSettings(value []CloudPcExternalPartnerSettingable)()
    SetGalleryImages(value []CloudPcGalleryImageable)()
    SetOnPremisesConnections(value []CloudPcOnPremisesConnectionable)()
    SetOrganizationSettings(value CloudPcOrganizationSettingsable)()
    SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)()
    SetReports(value CloudPcReportsable)()
    SetServicePlans(value []CloudPcServicePlanable)()
    SetSharedUseServicePlans(value []CloudPcSharedUseServicePlanable)()
    SetSnapshots(value []CloudPcSnapshotable)()
    SetSupportedRegions(value []CloudPcSupportedRegionable)()
    SetUserSettings(value []CloudPcUserSettingable)()
}
