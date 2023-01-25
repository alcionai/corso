package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagedTenant 
type ManagedTenant struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Aggregate view of device compliance policies across managed tenants.
    aggregatedPolicyCompliances []AggregatedPolicyComplianceable
    // The collection of audit events across managed tenants.
    auditEvents []AuditEventable
    // The collection of cloud PC connections across managed tenants.
    cloudPcConnections []CloudPcConnectionable
    // The collection of cloud PC devices across managed tenants.
    cloudPcDevices []CloudPcDeviceable
    // Overview of cloud PC information across managed tenants.
    cloudPcsOverview []CloudPcOverviewable
    // Aggregate view of conditional access policy coverage across managed tenants.
    conditionalAccessPolicyCoverages []ConditionalAccessPolicyCoverageable
    // Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
    credentialUserRegistrationsSummaries []CredentialUserRegistrationsSummaryable
    // Summary information for device compliance policy setting states across managed tenants.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummaryable
    // The collection of compliance for managed devices across managed tenants.
    managedDeviceCompliances []ManagedDeviceComplianceable
    // Trend insights for device compliance across managed tenants.
    managedDeviceComplianceTrends []ManagedDeviceComplianceTrendable
    // The managedTenantAlertLogs property
    managedTenantAlertLogs []ManagedTenantAlertLogable
    // The managedTenantAlertRuleDefinitions property
    managedTenantAlertRuleDefinitions []ManagedTenantAlertRuleDefinitionable
    // The managedTenantAlertRules property
    managedTenantAlertRules []ManagedTenantAlertRuleable
    // The managedTenantAlerts property
    managedTenantAlerts []ManagedTenantAlertable
    // The managedTenantApiNotifications property
    managedTenantApiNotifications []ManagedTenantApiNotificationable
    // The managedTenantEmailNotifications property
    managedTenantEmailNotifications []ManagedTenantEmailNotificationable
    // The managedTenantTicketingEndpoints property
    managedTenantTicketingEndpoints []ManagedTenantTicketingEndpointable
    // The collection of baseline management actions across managed tenants.
    managementActions []ManagementActionable
    // The tenant level status of management actions across managed tenants.
    managementActionTenantDeploymentStatuses []ManagementActionTenantDeploymentStatusable
    // The collection of baseline management intents across managed tenants.
    managementIntents []ManagementIntentable
    // The managementTemplateCollections property
    managementTemplateCollections []ManagementTemplateCollectionable
    // The managementTemplateCollectionTenantSummaries property
    managementTemplateCollectionTenantSummaries []ManagementTemplateCollectionTenantSummaryable
    // The collection of baseline management templates across managed tenants.
    managementTemplates []ManagementTemplateable
    // The managementTemplateSteps property
    managementTemplateSteps []ManagementTemplateStepable
    // The managementTemplateStepTenantSummaries property
    managementTemplateStepTenantSummaries []ManagementTemplateStepTenantSummaryable
    // The managementTemplateStepVersions property
    managementTemplateStepVersions []ManagementTemplateStepVersionable
    // The collection of role assignments to a signed-in user for a managed tenant.
    myRoles []MyRoleable
    // The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
    tenantGroups []TenantGroupable
    // The collection of tenants associated with the managing entity.
    tenants []Tenantable
    // The collection of tenant level customized information across managed tenants.
    tenantsCustomizedInformation []TenantCustomizedInformationable
    // The collection tenant level detailed information across managed tenants.
    tenantsDetailedInformation []TenantDetailedInformationable
    // The collection of tenant tags across managed tenants.
    tenantTags []TenantTagable
    // The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
    windowsDeviceMalwareStates []WindowsDeviceMalwareStateable
    // The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
    windowsProtectionStates []WindowsProtectionStateable
}
// NewManagedTenant instantiates a new ManagedTenant and sets the default values.
func NewManagedTenant()(*ManagedTenant) {
    m := &ManagedTenant{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagedTenantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenant(), nil
}
// GetAggregatedPolicyCompliances gets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) GetAggregatedPolicyCompliances()([]AggregatedPolicyComplianceable) {
    return m.aggregatedPolicyCompliances
}
// GetAuditEvents gets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) GetAuditEvents()([]AuditEventable) {
    return m.auditEvents
}
// GetCloudPcConnections gets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) GetCloudPcConnections()([]CloudPcConnectionable) {
    return m.cloudPcConnections
}
// GetCloudPcDevices gets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) GetCloudPcDevices()([]CloudPcDeviceable) {
    return m.cloudPcDevices
}
// GetCloudPcsOverview gets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) GetCloudPcsOverview()([]CloudPcOverviewable) {
    return m.cloudPcsOverview
}
// GetConditionalAccessPolicyCoverages gets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) GetConditionalAccessPolicyCoverages()([]ConditionalAccessPolicyCoverageable) {
    return m.conditionalAccessPolicyCoverages
}
// GetCredentialUserRegistrationsSummaries gets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) GetCredentialUserRegistrationsSummaries()([]CredentialUserRegistrationsSummaryable) {
    return m.credentialUserRegistrationsSummaries
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    return m.deviceCompliancePolicySettingStateSummaries
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedPolicyCompliances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAggregatedPolicyComplianceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AggregatedPolicyComplianceable, len(val))
            for i, v := range val {
                res[i] = v.(AggregatedPolicyComplianceable)
            }
            m.SetAggregatedPolicyCompliances(res)
        }
        return nil
    }
    res["auditEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditEventable, len(val))
            for i, v := range val {
                res[i] = v.(AuditEventable)
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["cloudPcConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcConnectionable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcConnectionable)
            }
            m.SetCloudPcConnections(res)
        }
        return nil
    }
    res["cloudPcDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcDeviceable)
            }
            m.SetCloudPcDevices(res)
        }
        return nil
    }
    res["cloudPcsOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcOverviewable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcOverviewable)
            }
            m.SetCloudPcsOverview(res)
        }
        return nil
    }
    res["conditionalAccessPolicyCoverages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyCoverageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyCoverageable, len(val))
            for i, v := range val {
                res[i] = v.(ConditionalAccessPolicyCoverageable)
            }
            m.SetConditionalAccessPolicyCoverages(res)
        }
        return nil
    }
    res["credentialUserRegistrationsSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CredentialUserRegistrationsSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(CredentialUserRegistrationsSummaryable)
            }
            m.SetCredentialUserRegistrationsSummaries(res)
        }
        return nil
    }
    res["deviceCompliancePolicySettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicySettingStateSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCompliancePolicySettingStateSummaryable)
            }
            m.SetDeviceCompliancePolicySettingStateSummaries(res)
        }
        return nil
    }
    res["managedDeviceCompliances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceComplianceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceComplianceable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceComplianceable)
            }
            m.SetManagedDeviceCompliances(res)
        }
        return nil
    }
    res["managedDeviceComplianceTrends"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceComplianceTrendFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceComplianceTrendable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceComplianceTrendable)
            }
            m.SetManagedDeviceComplianceTrends(res)
        }
        return nil
    }
    res["managedTenantAlertLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertLogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertLogable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertLogable)
            }
            m.SetManagedTenantAlertLogs(res)
        }
        return nil
    }
    res["managedTenantAlertRuleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertRuleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertRuleDefinitionable)
            }
            m.SetManagedTenantAlertRuleDefinitions(res)
        }
        return nil
    }
    res["managedTenantAlertRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertRuleable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertRuleable)
            }
            m.SetManagedTenantAlertRules(res)
        }
        return nil
    }
    res["managedTenantAlerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertable)
            }
            m.SetManagedTenantAlerts(res)
        }
        return nil
    }
    res["managedTenantApiNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantApiNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantApiNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantApiNotificationable)
            }
            m.SetManagedTenantApiNotifications(res)
        }
        return nil
    }
    res["managedTenantEmailNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantEmailNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantEmailNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantEmailNotificationable)
            }
            m.SetManagedTenantEmailNotifications(res)
        }
        return nil
    }
    res["managedTenantTicketingEndpoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantTicketingEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantTicketingEndpointable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantTicketingEndpointable)
            }
            m.SetManagedTenantTicketingEndpoints(res)
        }
        return nil
    }
    res["managementActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionable)
            }
            m.SetManagementActions(res)
        }
        return nil
    }
    res["managementActionTenantDeploymentStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionTenantDeploymentStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionTenantDeploymentStatusable)
            }
            m.SetManagementActionTenantDeploymentStatuses(res)
        }
        return nil
    }
    res["managementIntents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementIntentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementIntentable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementIntentable)
            }
            m.SetManagementIntents(res)
        }
        return nil
    }
    res["managementTemplateCollections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateCollectionable)
            }
            m.SetManagementTemplateCollections(res)
        }
        return nil
    }
    res["managementTemplateCollectionTenantSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateCollectionTenantSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollectionTenantSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateCollectionTenantSummaryable)
            }
            m.SetManagementTemplateCollectionTenantSummaries(res)
        }
        return nil
    }
    res["managementTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateable)
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    res["managementTemplateSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepable)
            }
            m.SetManagementTemplateSteps(res)
        }
        return nil
    }
    res["managementTemplateStepTenantSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepTenantSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepTenantSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepTenantSummaryable)
            }
            m.SetManagementTemplateStepTenantSummaries(res)
        }
        return nil
    }
    res["managementTemplateStepVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepVersionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepVersionable)
            }
            m.SetManagementTemplateStepVersions(res)
        }
        return nil
    }
    res["myRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMyRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MyRoleable, len(val))
            for i, v := range val {
                res[i] = v.(MyRoleable)
            }
            m.SetMyRoles(res)
        }
        return nil
    }
    res["tenantGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantGroupable, len(val))
            for i, v := range val {
                res[i] = v.(TenantGroupable)
            }
            m.SetTenantGroups(res)
        }
        return nil
    }
    res["tenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Tenantable, len(val))
            for i, v := range val {
                res[i] = v.(Tenantable)
            }
            m.SetTenants(res)
        }
        return nil
    }
    res["tenantsCustomizedInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantCustomizedInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantCustomizedInformationable, len(val))
            for i, v := range val {
                res[i] = v.(TenantCustomizedInformationable)
            }
            m.SetTenantsCustomizedInformation(res)
        }
        return nil
    }
    res["tenantsDetailedInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantDetailedInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantDetailedInformationable, len(val))
            for i, v := range val {
                res[i] = v.(TenantDetailedInformationable)
            }
            m.SetTenantsDetailedInformation(res)
        }
        return nil
    }
    res["tenantTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantTagable, len(val))
            for i, v := range val {
                res[i] = v.(TenantTagable)
            }
            m.SetTenantTags(res)
        }
        return nil
    }
    res["windowsDeviceMalwareStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDeviceMalwareStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDeviceMalwareStateable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDeviceMalwareStateable)
            }
            m.SetWindowsDeviceMalwareStates(res)
        }
        return nil
    }
    res["windowsProtectionStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsProtectionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsProtectionStateable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsProtectionStateable)
            }
            m.SetWindowsProtectionStates(res)
        }
        return nil
    }
    return res
}
// GetManagedDeviceCompliances gets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) GetManagedDeviceCompliances()([]ManagedDeviceComplianceable) {
    return m.managedDeviceCompliances
}
// GetManagedDeviceComplianceTrends gets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) GetManagedDeviceComplianceTrends()([]ManagedDeviceComplianceTrendable) {
    return m.managedDeviceComplianceTrends
}
// GetManagedTenantAlertLogs gets the managedTenantAlertLogs property value. The managedTenantAlertLogs property
func (m *ManagedTenant) GetManagedTenantAlertLogs()([]ManagedTenantAlertLogable) {
    return m.managedTenantAlertLogs
}
// GetManagedTenantAlertRuleDefinitions gets the managedTenantAlertRuleDefinitions property value. The managedTenantAlertRuleDefinitions property
func (m *ManagedTenant) GetManagedTenantAlertRuleDefinitions()([]ManagedTenantAlertRuleDefinitionable) {
    return m.managedTenantAlertRuleDefinitions
}
// GetManagedTenantAlertRules gets the managedTenantAlertRules property value. The managedTenantAlertRules property
func (m *ManagedTenant) GetManagedTenantAlertRules()([]ManagedTenantAlertRuleable) {
    return m.managedTenantAlertRules
}
// GetManagedTenantAlerts gets the managedTenantAlerts property value. The managedTenantAlerts property
func (m *ManagedTenant) GetManagedTenantAlerts()([]ManagedTenantAlertable) {
    return m.managedTenantAlerts
}
// GetManagedTenantApiNotifications gets the managedTenantApiNotifications property value. The managedTenantApiNotifications property
func (m *ManagedTenant) GetManagedTenantApiNotifications()([]ManagedTenantApiNotificationable) {
    return m.managedTenantApiNotifications
}
// GetManagedTenantEmailNotifications gets the managedTenantEmailNotifications property value. The managedTenantEmailNotifications property
func (m *ManagedTenant) GetManagedTenantEmailNotifications()([]ManagedTenantEmailNotificationable) {
    return m.managedTenantEmailNotifications
}
// GetManagedTenantTicketingEndpoints gets the managedTenantTicketingEndpoints property value. The managedTenantTicketingEndpoints property
func (m *ManagedTenant) GetManagedTenantTicketingEndpoints()([]ManagedTenantTicketingEndpointable) {
    return m.managedTenantTicketingEndpoints
}
// GetManagementActions gets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) GetManagementActions()([]ManagementActionable) {
    return m.managementActions
}
// GetManagementActionTenantDeploymentStatuses gets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) GetManagementActionTenantDeploymentStatuses()([]ManagementActionTenantDeploymentStatusable) {
    return m.managementActionTenantDeploymentStatuses
}
// GetManagementIntents gets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) GetManagementIntents()([]ManagementIntentable) {
    return m.managementIntents
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) GetManagementTemplateCollections()([]ManagementTemplateCollectionable) {
    return m.managementTemplateCollections
}
// GetManagementTemplateCollectionTenantSummaries gets the managementTemplateCollectionTenantSummaries property value. The managementTemplateCollectionTenantSummaries property
func (m *ManagedTenant) GetManagementTemplateCollectionTenantSummaries()([]ManagementTemplateCollectionTenantSummaryable) {
    return m.managementTemplateCollectionTenantSummaries
}
// GetManagementTemplates gets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) GetManagementTemplates()([]ManagementTemplateable) {
    return m.managementTemplates
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) GetManagementTemplateSteps()([]ManagementTemplateStepable) {
    return m.managementTemplateSteps
}
// GetManagementTemplateStepTenantSummaries gets the managementTemplateStepTenantSummaries property value. The managementTemplateStepTenantSummaries property
func (m *ManagedTenant) GetManagementTemplateStepTenantSummaries()([]ManagementTemplateStepTenantSummaryable) {
    return m.managementTemplateStepTenantSummaries
}
// GetManagementTemplateStepVersions gets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) GetManagementTemplateStepVersions()([]ManagementTemplateStepVersionable) {
    return m.managementTemplateStepVersions
}
// GetMyRoles gets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) GetMyRoles()([]MyRoleable) {
    return m.myRoles
}
// GetTenantGroups gets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) GetTenantGroups()([]TenantGroupable) {
    return m.tenantGroups
}
// GetTenants gets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) GetTenants()([]Tenantable) {
    return m.tenants
}
// GetTenantsCustomizedInformation gets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) GetTenantsCustomizedInformation()([]TenantCustomizedInformationable) {
    return m.tenantsCustomizedInformation
}
// GetTenantsDetailedInformation gets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) GetTenantsDetailedInformation()([]TenantDetailedInformationable) {
    return m.tenantsDetailedInformation
}
// GetTenantTags gets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) GetTenantTags()([]TenantTagable) {
    return m.tenantTags
}
// GetWindowsDeviceMalwareStates gets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsDeviceMalwareStates()([]WindowsDeviceMalwareStateable) {
    return m.windowsDeviceMalwareStates
}
// GetWindowsProtectionStates gets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsProtectionStates()([]WindowsProtectionStateable) {
    return m.windowsProtectionStates
}
// Serialize serializes information the current object
func (m *ManagedTenant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAggregatedPolicyCompliances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAggregatedPolicyCompliances()))
        for i, v := range m.GetAggregatedPolicyCompliances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("aggregatedPolicyCompliances", cast)
        if err != nil {
            return err
        }
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
    if m.GetCloudPcConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcConnections()))
        for i, v := range m.GetCloudPcConnections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcDevices()))
        for i, v := range m.GetCloudPcDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcsOverview() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcsOverview()))
        for i, v := range m.GetCloudPcsOverview() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcsOverview", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicyCoverages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConditionalAccessPolicyCoverages()))
        for i, v := range m.GetConditionalAccessPolicyCoverages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicyCoverages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCredentialUserRegistrationsSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredentialUserRegistrationsSummaries()))
        for i, v := range m.GetCredentialUserRegistrationsSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationsSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicySettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicySettingStateSummaries()))
        for i, v := range m.GetDeviceCompliancePolicySettingStateSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCompliances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceCompliances()))
        for i, v := range m.GetManagedDeviceCompliances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCompliances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceComplianceTrends() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceComplianceTrends()))
        for i, v := range m.GetManagedDeviceComplianceTrends() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceComplianceTrends", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantAlertLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlertLogs()))
        for i, v := range m.GetManagedTenantAlertLogs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlertLogs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantAlertRuleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlertRuleDefinitions()))
        for i, v := range m.GetManagedTenantAlertRuleDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlertRuleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantAlertRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlertRules()))
        for i, v := range m.GetManagedTenantAlertRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlertRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlerts()))
        for i, v := range m.GetManagedTenantAlerts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlerts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantApiNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantApiNotifications()))
        for i, v := range m.GetManagedTenantApiNotifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantApiNotifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantEmailNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantEmailNotifications()))
        for i, v := range m.GetManagedTenantEmailNotifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantEmailNotifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantTicketingEndpoints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantTicketingEndpoints()))
        for i, v := range m.GetManagedTenantTicketingEndpoints() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantTicketingEndpoints", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementActions()))
        for i, v := range m.GetManagementActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActionTenantDeploymentStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementActionTenantDeploymentStatuses()))
        for i, v := range m.GetManagementActionTenantDeploymentStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementActionTenantDeploymentStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementIntents()))
        for i, v := range m.GetManagementIntents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateCollections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateCollections()))
        for i, v := range m.GetManagementTemplateCollections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateCollectionTenantSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateCollectionTenantSummaries()))
        for i, v := range m.GetManagementTemplateCollectionTenantSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollectionTenantSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateSteps()))
        for i, v := range m.GetManagementTemplateSteps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateStepTenantSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateStepTenantSummaries()))
        for i, v := range m.GetManagementTemplateStepTenantSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateStepTenantSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateStepVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateStepVersions()))
        for i, v := range m.GetManagementTemplateStepVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateStepVersions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMyRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMyRoles()))
        for i, v := range m.GetMyRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("myRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantGroups()))
        for i, v := range m.GetTenantGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsCustomizedInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantsCustomizedInformation()))
        for i, v := range m.GetTenantsCustomizedInformation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantsCustomizedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsDetailedInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantsDetailedInformation()))
        for i, v := range m.GetTenantsDetailedInformation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantsDetailedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantTags()))
        for i, v := range m.GetTenantTags() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDeviceMalwareStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsDeviceMalwareStates()))
        for i, v := range m.GetWindowsDeviceMalwareStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsDeviceMalwareStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsProtectionStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsProtectionStates()))
        for i, v := range m.GetWindowsProtectionStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsProtectionStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAggregatedPolicyCompliances sets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) SetAggregatedPolicyCompliances(value []AggregatedPolicyComplianceable)() {
    m.aggregatedPolicyCompliances = value
}
// SetAuditEvents sets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) SetAuditEvents(value []AuditEventable)() {
    m.auditEvents = value
}
// SetCloudPcConnections sets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) SetCloudPcConnections(value []CloudPcConnectionable)() {
    m.cloudPcConnections = value
}
// SetCloudPcDevices sets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) SetCloudPcDevices(value []CloudPcDeviceable)() {
    m.cloudPcDevices = value
}
// SetCloudPcsOverview sets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) SetCloudPcsOverview(value []CloudPcOverviewable)() {
    m.cloudPcsOverview = value
}
// SetConditionalAccessPolicyCoverages sets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) SetConditionalAccessPolicyCoverages(value []ConditionalAccessPolicyCoverageable)() {
    m.conditionalAccessPolicyCoverages = value
}
// SetCredentialUserRegistrationsSummaries sets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) SetCredentialUserRegistrationsSummaries(value []CredentialUserRegistrationsSummaryable)() {
    m.credentialUserRegistrationsSummaries = value
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// SetManagedDeviceCompliances sets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) SetManagedDeviceCompliances(value []ManagedDeviceComplianceable)() {
    m.managedDeviceCompliances = value
}
// SetManagedDeviceComplianceTrends sets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) SetManagedDeviceComplianceTrends(value []ManagedDeviceComplianceTrendable)() {
    m.managedDeviceComplianceTrends = value
}
// SetManagedTenantAlertLogs sets the managedTenantAlertLogs property value. The managedTenantAlertLogs property
func (m *ManagedTenant) SetManagedTenantAlertLogs(value []ManagedTenantAlertLogable)() {
    m.managedTenantAlertLogs = value
}
// SetManagedTenantAlertRuleDefinitions sets the managedTenantAlertRuleDefinitions property value. The managedTenantAlertRuleDefinitions property
func (m *ManagedTenant) SetManagedTenantAlertRuleDefinitions(value []ManagedTenantAlertRuleDefinitionable)() {
    m.managedTenantAlertRuleDefinitions = value
}
// SetManagedTenantAlertRules sets the managedTenantAlertRules property value. The managedTenantAlertRules property
func (m *ManagedTenant) SetManagedTenantAlertRules(value []ManagedTenantAlertRuleable)() {
    m.managedTenantAlertRules = value
}
// SetManagedTenantAlerts sets the managedTenantAlerts property value. The managedTenantAlerts property
func (m *ManagedTenant) SetManagedTenantAlerts(value []ManagedTenantAlertable)() {
    m.managedTenantAlerts = value
}
// SetManagedTenantApiNotifications sets the managedTenantApiNotifications property value. The managedTenantApiNotifications property
func (m *ManagedTenant) SetManagedTenantApiNotifications(value []ManagedTenantApiNotificationable)() {
    m.managedTenantApiNotifications = value
}
// SetManagedTenantEmailNotifications sets the managedTenantEmailNotifications property value. The managedTenantEmailNotifications property
func (m *ManagedTenant) SetManagedTenantEmailNotifications(value []ManagedTenantEmailNotificationable)() {
    m.managedTenantEmailNotifications = value
}
// SetManagedTenantTicketingEndpoints sets the managedTenantTicketingEndpoints property value. The managedTenantTicketingEndpoints property
func (m *ManagedTenant) SetManagedTenantTicketingEndpoints(value []ManagedTenantTicketingEndpointable)() {
    m.managedTenantTicketingEndpoints = value
}
// SetManagementActions sets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) SetManagementActions(value []ManagementActionable)() {
    m.managementActions = value
}
// SetManagementActionTenantDeploymentStatuses sets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) SetManagementActionTenantDeploymentStatuses(value []ManagementActionTenantDeploymentStatusable)() {
    m.managementActionTenantDeploymentStatuses = value
}
// SetManagementIntents sets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) SetManagementIntents(value []ManagementIntentable)() {
    m.managementIntents = value
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) SetManagementTemplateCollections(value []ManagementTemplateCollectionable)() {
    m.managementTemplateCollections = value
}
// SetManagementTemplateCollectionTenantSummaries sets the managementTemplateCollectionTenantSummaries property value. The managementTemplateCollectionTenantSummaries property
func (m *ManagedTenant) SetManagementTemplateCollectionTenantSummaries(value []ManagementTemplateCollectionTenantSummaryable)() {
    m.managementTemplateCollectionTenantSummaries = value
}
// SetManagementTemplates sets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) SetManagementTemplates(value []ManagementTemplateable)() {
    m.managementTemplates = value
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) SetManagementTemplateSteps(value []ManagementTemplateStepable)() {
    m.managementTemplateSteps = value
}
// SetManagementTemplateStepTenantSummaries sets the managementTemplateStepTenantSummaries property value. The managementTemplateStepTenantSummaries property
func (m *ManagedTenant) SetManagementTemplateStepTenantSummaries(value []ManagementTemplateStepTenantSummaryable)() {
    m.managementTemplateStepTenantSummaries = value
}
// SetManagementTemplateStepVersions sets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) SetManagementTemplateStepVersions(value []ManagementTemplateStepVersionable)() {
    m.managementTemplateStepVersions = value
}
// SetMyRoles sets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) SetMyRoles(value []MyRoleable)() {
    m.myRoles = value
}
// SetTenantGroups sets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) SetTenantGroups(value []TenantGroupable)() {
    m.tenantGroups = value
}
// SetTenants sets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) SetTenants(value []Tenantable)() {
    m.tenants = value
}
// SetTenantsCustomizedInformation sets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) SetTenantsCustomizedInformation(value []TenantCustomizedInformationable)() {
    m.tenantsCustomizedInformation = value
}
// SetTenantsDetailedInformation sets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) SetTenantsDetailedInformation(value []TenantDetailedInformationable)() {
    m.tenantsDetailedInformation = value
}
// SetTenantTags sets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) SetTenantTags(value []TenantTagable)() {
    m.tenantTags = value
}
// SetWindowsDeviceMalwareStates sets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsDeviceMalwareStates(value []WindowsDeviceMalwareStateable)() {
    m.windowsDeviceMalwareStates = value
}
// SetWindowsProtectionStates sets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsProtectionStates(value []WindowsProtectionStateable)() {
    m.windowsProtectionStates = value
}
