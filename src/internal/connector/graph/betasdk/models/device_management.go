package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagement 
type DeviceManagement struct {
    Entity
    // The date & time when tenant data moved between scaleunits.
    accountMoveCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin consent information.
    adminConsent AdminConsentable
    // The summary state of ATP onboarding state for this account.
    advancedThreatProtectionOnboardingStateSummary AdvancedThreatProtectionOnboardingStateSummaryable
    // Android device owner enrollment profile entities.
    androidDeviceOwnerEnrollmentProfiles []AndroidDeviceOwnerEnrollmentProfileable
    // Android for Work app configuration schema entities.
    androidForWorkAppConfigurationSchemas []AndroidForWorkAppConfigurationSchemaable
    // Android for Work enrollment profile entities.
    androidForWorkEnrollmentProfiles []AndroidForWorkEnrollmentProfileable
    // The singleton Android for Work settings entity.
    androidForWorkSettings AndroidForWorkSettingsable
    // The singleton Android managed store account enterprise settings entity.
    androidManagedStoreAccountEnterpriseSettings AndroidManagedStoreAccountEnterpriseSettingsable
    // Android Enterprise app configuration schema entities.
    androidManagedStoreAppConfigurationSchemas []AndroidManagedStoreAppConfigurationSchemaable
    // Apple push notification certificate.
    applePushNotificationCertificate ApplePushNotificationCertificateable
    // Apple user initiated enrollment profiles
    appleUserInitiatedEnrollmentProfiles []AppleUserInitiatedEnrollmentProfileable
    // The list of assignment filters
    assignmentFilters []DeviceAndAppManagementAssignmentFilterable
    // The Audit Events
    auditEvents []AuditEventable
    // The list of autopilot events for the tenant.
    autopilotEvents []DeviceManagementAutopilotEventable
    // The Cart To Class Associations.
    cartToClassAssociations []CartToClassAssociationable
    // The available categories
    categories []DeviceManagementSettingCategoryable
    // Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
    certificateConnectorDetails []CertificateConnectorDetailsable
    // Collection of ChromeOSOnboardingSettings settings associated with account.
    chromeOSOnboardingSettings []ChromeOSOnboardingSettingsable
    // The list of CloudPC Connectivity Issue.
    cloudPCConnectivityIssues []CloudPCConnectivityIssueable
    // The list of co-managed devices report
    comanagedDevices []ManagedDeviceable
    // The list of co-management eligible devices report
    comanagementEligibleDevices []ComanagementEligibleDeviceable
    // List of all compliance categories
    complianceCategories []DeviceManagementConfigurationCategoryable
    // The list of Compliance Management Partners configured by the tenant.
    complianceManagementPartners []ComplianceManagementPartnerable
    // List of all compliance policies
    compliancePolicies []DeviceManagementCompliancePolicyable
    // List of all ComplianceSettings
    complianceSettings []DeviceManagementConfigurationSettingDefinitionable
    // The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
    conditionalAccessSettings OnPremisesConditionalAccessSettingsable
    // A list of ConfigManagerCollection
    configManagerCollections []ConfigManagerCollectionable
    // List of all Configuration Categories
    configurationCategories []DeviceManagementConfigurationCategoryable
    // List of all Configuration policies
    configurationPolicies []DeviceManagementConfigurationPolicyable
    // List of all templates
    configurationPolicyTemplates []DeviceManagementConfigurationPolicyTemplateable
    // List of all ConfigurationSettings
    configurationSettings []DeviceManagementConfigurationSettingDefinitionable
    // A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
    dataProcessorServiceForWindowsFeaturesOnboarding DataProcessorServiceForWindowsFeaturesOnboardingable
    // Data sharing consents.
    dataSharingConsents []DataSharingConsentable
    // This collections of multiple DEP tokens per-tenant.
    depOnboardingSettings []DepOnboardingSettingable
    // Collection of Derived credential settings associated with account.
    derivedCredentials []DeviceManagementDerivedCredentialSettingsable
    // The list of detected apps associated with a device.
    detectedApps []DetectedAppable
    // The list of device categories with the tenant.
    deviceCategories []DeviceCategoryable
    // The device compliance policies.
    deviceCompliancePolicies []DeviceCompliancePolicyable
    // The device compliance state summary for this account.
    deviceCompliancePolicyDeviceStateSummary DeviceCompliancePolicyDeviceStateSummaryable
    // The summary states of compliance policy settings for this account.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummaryable
    // The last requested time of device compliance reporting for this account. This property is read-only.
    deviceComplianceReportSummarizationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The list of device compliance scripts associated with the tenant.
    deviceComplianceScripts []DeviceComplianceScriptable
    // Summary of policies in conflict state for this account.
    deviceConfigurationConflictSummary []DeviceConfigurationConflictSummaryable
    // The device configuration device state summary for this account.
    deviceConfigurationDeviceStateSummaries DeviceConfigurationDeviceStateSummaryable
    // Restricted apps violations for this account.
    deviceConfigurationRestrictedAppsViolations []RestrictedAppsViolationable
    // The device configurations.
    deviceConfigurations []DeviceConfigurationable
    // Summary of all certificates for all devices.
    deviceConfigurationsAllManagedDeviceCertificateStates []ManagedAllDeviceCertificateStateable
    // The device configuration user state summary for this account.
    deviceConfigurationUserStateSummaries DeviceConfigurationUserStateSummaryable
    // The list of device custom attribute shell scripts associated with the tenant.
    deviceCustomAttributeShellScripts []DeviceCustomAttributeShellScriptable
    // The list of device enrollment configurations
    deviceEnrollmentConfigurations []DeviceEnrollmentConfigurationable
    // The list of device health scripts associated with the tenant.
    deviceHealthScripts []DeviceHealthScriptable
    // The list of Device Management Partners configured by the tenant.
    deviceManagementPartners []DeviceManagementPartnerable
    // The list of device management scripts associated with the tenant.
    deviceManagementScripts []DeviceManagementScriptable
    // Device protection overview.
    deviceProtectionOverview DeviceProtectionOverviewable
    // The list of device shell scripts associated with the tenant.
    deviceShellScripts []DeviceShellScriptable
    // A list of connector objects.
    domainJoinConnectors []DeviceManagementDomainJoinConnectorable
    // The embedded SIM activation code pools created by this account.
    embeddedSIMActivationCodePools []EmbeddedSIMActivationCodePoolable
    // The list of Exchange Connectors configured by the tenant.
    exchangeConnectors []DeviceManagementExchangeConnectorable
    // The list of Exchange On Premisis policies configured by the tenant.
    exchangeOnPremisesPolicies []DeviceManagementExchangeOnPremisesPolicyable
    // The policy which controls mobile device access to Exchange On Premises
    exchangeOnPremisesPolicy DeviceManagementExchangeOnPremisesPolicyable
    // The available group policy categories for this account.
    groupPolicyCategories []GroupPolicyCategoryable
    // The group policy configurations created by this account.
    groupPolicyConfigurations []GroupPolicyConfigurationable
    // The available group policy definition files for this account.
    groupPolicyDefinitionFiles []GroupPolicyDefinitionFileable
    // The available group policy definitions for this account.
    groupPolicyDefinitions []GroupPolicyDefinitionable
    // A list of Group Policy migration reports.
    groupPolicyMigrationReports []GroupPolicyMigrationReportable
    // A list of Group Policy Object files uploaded.
    groupPolicyObjectFiles []GroupPolicyObjectFileable
    // The available group policy uploaded definition files for this account.
    groupPolicyUploadedDefinitionFiles []GroupPolicyUploadedDefinitionFileable
    // The imported device identities.
    importedDeviceIdentities []ImportedDeviceIdentityable
    // Collection of imported Windows autopilot devices.
    importedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentityable
    // The device management intents
    intents []DeviceManagementIntentable
    // Intune Account ID for given tenant
    intuneAccountId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
    intuneBrand IntuneBrandable
    // Intune branding profiles targeted to AAD groups
    intuneBrandingProfiles []IntuneBrandingProfileable
    // The IOS software update installation statuses for this account.
    iosUpdateStatuses []IosUpdateDeviceStatusable
    // The last modified time of reporting for this account. This property is read-only.
    lastReportAggregationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
    legacyPcManangementEnabled *bool
    // The MacOS software update account summaries for this account.
    macOSSoftwareUpdateAccountSummaries []MacOSSoftwareUpdateAccountSummaryable
    // Device cleanup rule
    managedDeviceCleanupSettings ManagedDeviceCleanupSettingsable
    // Encryption report for devices in this account
    managedDeviceEncryptionStates []ManagedDeviceEncryptionStateable
    // Device overview
    managedDeviceOverview ManagedDeviceOverviewable
    // The list of managed devices.
    managedDevices []ManagedDeviceable
    // Maximum number of DEP tokens allowed per-tenant.
    maximumDepTokens *int32
    // Collection of MicrosoftTunnelConfiguration settings associated with account.
    microsoftTunnelConfigurations []MicrosoftTunnelConfigurationable
    // Collection of MicrosoftTunnelHealthThreshold settings associated with account.
    microsoftTunnelHealthThresholds []MicrosoftTunnelHealthThresholdable
    // Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
    microsoftTunnelServerLogCollectionResponses []MicrosoftTunnelServerLogCollectionResponseable
    // Collection of MicrosoftTunnelSite settings associated with account.
    microsoftTunnelSites []MicrosoftTunnelSiteable
    // The collection property of MobileAppTroubleshootingEvent.
    mobileAppTroubleshootingEvents []MobileAppTroubleshootingEventable
    // The list of Mobile threat Defense connectors configured by the tenant.
    mobileThreatDefenseConnectors []MobileThreatDefenseConnectorable
    // The collection of Ndes connectors for this account.
    ndesConnectors []NdesConnectorable
    // The Notification Message Templates.
    notificationMessageTemplates []NotificationMessageTemplateable
    // List of OEM Warranty Statuses
    oemWarrantyInformationOnboarding []OemWarrantyInformationOnboardingable
    // The list of device remote action audits with the tenant.
    remoteActionAudits []RemoteActionAuditable
    // The remote assist partners.
    remoteAssistancePartners []RemoteAssistancePartnerable
    // The remote assistance settings singleton
    remoteAssistanceSettings RemoteAssistanceSettingsable
    // Reports singleton
    reports DeviceManagementReportsable
    // Collection of resource access settings associated with account.
    resourceAccessProfiles []DeviceManagementResourceAccessProfileBaseable
    // The Resource Operations.
    resourceOperations []ResourceOperationable
    // List of all reusable settings that can be referred in a policy
    reusablePolicySettings []DeviceManagementReusablePolicySettingable
    // List of all reusable settings
    reusableSettings []DeviceManagementConfigurationSettingDefinitionable
    // The Role Assignments.
    roleAssignments []DeviceAndAppManagementRoleAssignmentable
    // The Role Definitions.
    roleDefinitions []RoleDefinitionable
    // The Role Scope Tags.
    roleScopeTags []RoleScopeTagable
    // The device management intent setting definitions
    settingDefinitions []DeviceManagementSettingDefinitionable
    // Account level settings.
    settings DeviceManagementSettingsable
    // The software update status summary.
    softwareUpdateStatusSummary SoftwareUpdateStatusSummaryable
    // Tenant mobile device management subscriptions.
    subscriptions *DeviceManagementSubscriptions
    // Tenant mobile device management subscription state.
    subscriptionState *DeviceManagementSubscriptionState
    // The telecom expense management partners.
    telecomExpenseManagementPartners []TelecomExpenseManagementPartnerable
    // The available templates
    templates []DeviceManagementTemplateable
    // List of all TemplateSettings
    templateSettings []DeviceManagementConfigurationSettingTemplateable
    // TenantAttach RBAC Enablement
    tenantAttachRBAC TenantAttachRBACable
    // The terms and conditions associated with device management of the company.
    termsAndConditions []TermsAndConditionsable
    // The list of troubleshooting events for the tenant.
    troubleshootingEvents []DeviceManagementTroubleshootingEventable
    // When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
    unlicensedAdminstratorsEnabled *bool
    // The user experience analytics anomaly entity contains anomaly details.
    userExperienceAnalyticsAnomaly []UserExperienceAnalyticsAnomalyable
    // The user experience analytics anomaly entity contains device details.
    userExperienceAnalyticsAnomalyDevice []UserExperienceAnalyticsAnomalyDeviceable
    // The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
    userExperienceAnalyticsAnomalySeverityOverview UserExperienceAnalyticsAnomalySeverityOverviewable
    // User experience analytics appHealth Application Performance
    userExperienceAnalyticsAppHealthApplicationPerformance []UserExperienceAnalyticsAppHealthApplicationPerformanceable
    // User experience analytics appHealth Application Performance by App Version
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable
    // User experience analytics appHealth Application Performance by App Version details
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable
    // User experience analytics appHealth Application Performance by App Version Device Id
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable
    // User experience analytics appHealth Application Performance by OS Version
    userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable
    // User experience analytics appHealth Model Performance
    userExperienceAnalyticsAppHealthDeviceModelPerformance []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable
    // User experience analytics appHealth Device Performance
    userExperienceAnalyticsAppHealthDevicePerformance []UserExperienceAnalyticsAppHealthDevicePerformanceable
    // User experience analytics device performance details
    userExperienceAnalyticsAppHealthDevicePerformanceDetails []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable
    // User experience analytics appHealth OS version Performance
    userExperienceAnalyticsAppHealthOSVersionPerformance []UserExperienceAnalyticsAppHealthOSVersionPerformanceable
    // User experience analytics appHealth overview
    userExperienceAnalyticsAppHealthOverview UserExperienceAnalyticsCategoryable
    // User experience analytics baselines
    userExperienceAnalyticsBaselines []UserExperienceAnalyticsBaselineable
    // User Experience Analytics Battery Health App Impact
    userExperienceAnalyticsBatteryHealthAppImpact []UserExperienceAnalyticsBatteryHealthAppImpactable
    // User Experience Analytics Battery Health Capacity Details
    userExperienceAnalyticsBatteryHealthCapacityDetails UserExperienceAnalyticsBatteryHealthCapacityDetailsable
    // User Experience Analytics Battery Health Device App Impact
    userExperienceAnalyticsBatteryHealthDeviceAppImpact []UserExperienceAnalyticsBatteryHealthDeviceAppImpactable
    // User Experience Analytics Battery Health Device Performance
    userExperienceAnalyticsBatteryHealthDevicePerformance []UserExperienceAnalyticsBatteryHealthDevicePerformanceable
    // User Experience Analytics Battery Health Device Runtime History
    userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable
    // User Experience Analytics Battery Health Model Performance
    userExperienceAnalyticsBatteryHealthModelPerformance []UserExperienceAnalyticsBatteryHealthModelPerformanceable
    // User Experience Analytics Battery Health Os Performance
    userExperienceAnalyticsBatteryHealthOsPerformance []UserExperienceAnalyticsBatteryHealthOsPerformanceable
    // User Experience Analytics Battery Health Runtime Details
    userExperienceAnalyticsBatteryHealthRuntimeDetails UserExperienceAnalyticsBatteryHealthRuntimeDetailsable
    // User experience analytics categories
    userExperienceAnalyticsCategories []UserExperienceAnalyticsCategoryable
    // User experience analytics device metric history
    userExperienceAnalyticsDeviceMetricHistory []UserExperienceAnalyticsMetricHistoryable
    // User experience analytics device performance
    userExperienceAnalyticsDevicePerformance []UserExperienceAnalyticsDevicePerformanceable
    // The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
    userExperienceAnalyticsDeviceScope UserExperienceAnalyticsDeviceScopeable
    // The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
    userExperienceAnalyticsDeviceScopes []UserExperienceAnalyticsDeviceScopeable
    // User experience analytics device scores
    userExperienceAnalyticsDeviceScores []UserExperienceAnalyticsDeviceScoresable
    // User experience analytics device Startup History
    userExperienceAnalyticsDeviceStartupHistory []UserExperienceAnalyticsDeviceStartupHistoryable
    // User experience analytics device Startup Processes
    userExperienceAnalyticsDeviceStartupProcesses []UserExperienceAnalyticsDeviceStartupProcessable
    // User experience analytics device Startup Process Performance
    userExperienceAnalyticsDeviceStartupProcessPerformance []UserExperienceAnalyticsDeviceStartupProcessPerformanceable
    // User experience analytics devices without cloud identity.
    userExperienceAnalyticsDevicesWithoutCloudIdentity []UserExperienceAnalyticsDeviceWithoutCloudIdentityable
    // The user experience analytics device events entity contains NRT device timeline events details.
    userExperienceAnalyticsDeviceTimelineEvents []UserExperienceAnalyticsDeviceTimelineEventsable
    // User experience analytics impacting process
    userExperienceAnalyticsImpactingProcess []UserExperienceAnalyticsImpactingProcessable
    // User experience analytics metric history
    userExperienceAnalyticsMetricHistory []UserExperienceAnalyticsMetricHistoryable
    // User experience analytics model scores
    userExperienceAnalyticsModelScores []UserExperienceAnalyticsModelScoresable
    // User experience analytics devices not Windows Autopilot ready.
    userExperienceAnalyticsNotAutopilotReadyDevice []UserExperienceAnalyticsNotAutopilotReadyDeviceable
    // User experience analytics overview
    userExperienceAnalyticsOverview UserExperienceAnalyticsOverviewable
    // User experience analytics remote connection
    userExperienceAnalyticsRemoteConnection []UserExperienceAnalyticsRemoteConnectionable
    // User experience analytics resource performance
    userExperienceAnalyticsResourcePerformance []UserExperienceAnalyticsResourcePerformanceable
    // User experience analytics device Startup Score History
    userExperienceAnalyticsScoreHistory []UserExperienceAnalyticsScoreHistoryable
    // User experience analytics device settings
    userExperienceAnalyticsSettings UserExperienceAnalyticsSettingsable
    // User experience analytics work from anywhere hardware readiness metrics.
    userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable
    // User experience analytics work from anywhere metrics.
    userExperienceAnalyticsWorkFromAnywhereMetrics []UserExperienceAnalyticsWorkFromAnywhereMetricable
    // The user experience analytics work from anywhere model performance
    userExperienceAnalyticsWorkFromAnywhereModelPerformance []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable
    // Collection of PFX certificates associated with a user.
    userPfxCertificates []UserPFXCertificateable
    // The virtualEndpoint property
    virtualEndpoint VirtualEndpointable
    // Windows auto pilot deployment profiles
    windowsAutopilotDeploymentProfiles []WindowsAutopilotDeploymentProfileable
    // The Windows autopilot device identities contained collection.
    windowsAutopilotDeviceIdentities []WindowsAutopilotDeviceIdentityable
    // The Windows autopilot account settings.
    windowsAutopilotSettings WindowsAutopilotSettingsable
    // A collection of windows driver update profiles
    windowsDriverUpdateProfiles []WindowsDriverUpdateProfileable
    // A collection of windows feature update profiles
    windowsFeatureUpdateProfiles []WindowsFeatureUpdateProfileable
    // The windows information protection app learning summaries.
    windowsInformationProtectionAppLearningSummaries []WindowsInformationProtectionAppLearningSummaryable
    // The windows information protection network learning summaries.
    windowsInformationProtectionNetworkLearningSummaries []WindowsInformationProtectionNetworkLearningSummaryable
    // The list of affected malware in the tenant.
    windowsMalwareInformation []WindowsMalwareInformationable
    // Malware overview for windows devices.
    windowsMalwareOverview WindowsMalwareOverviewable
    // A collection of windows quality update profiles
    windowsQualityUpdateProfiles []WindowsQualityUpdateProfileable
    // A collection of windows update catalog items (fetaure updates item , quality updates item)
    windowsUpdateCatalogItems []WindowsUpdateCatalogItemable
    // The Collection of ZebraFotaArtifacts.
    zebraFotaArtifacts []ZebraFotaArtifactable
    // The singleton ZebraFotaConnector associated with account.
    zebraFotaConnector ZebraFotaConnectorable
    // Collection of ZebraFotaDeployments associated with account.
    zebraFotaDeployments []ZebraFotaDeploymentable
}
// NewDeviceManagement instantiates a new DeviceManagement and sets the default values.
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagement(), nil
}
// GetAccountMoveCompletionDateTime gets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
func (m *DeviceManagement) GetAccountMoveCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.accountMoveCompletionDateTime
}
// GetAdminConsent gets the adminConsent property value. Admin consent information.
func (m *DeviceManagement) GetAdminConsent()(AdminConsentable) {
    return m.adminConsent
}
// GetAdvancedThreatProtectionOnboardingStateSummary gets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
func (m *DeviceManagement) GetAdvancedThreatProtectionOnboardingStateSummary()(AdvancedThreatProtectionOnboardingStateSummaryable) {
    return m.advancedThreatProtectionOnboardingStateSummary
}
// GetAndroidDeviceOwnerEnrollmentProfiles gets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
func (m *DeviceManagement) GetAndroidDeviceOwnerEnrollmentProfiles()([]AndroidDeviceOwnerEnrollmentProfileable) {
    return m.androidDeviceOwnerEnrollmentProfiles
}
// GetAndroidForWorkAppConfigurationSchemas gets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
func (m *DeviceManagement) GetAndroidForWorkAppConfigurationSchemas()([]AndroidForWorkAppConfigurationSchemaable) {
    return m.androidForWorkAppConfigurationSchemas
}
// GetAndroidForWorkEnrollmentProfiles gets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
func (m *DeviceManagement) GetAndroidForWorkEnrollmentProfiles()([]AndroidForWorkEnrollmentProfileable) {
    return m.androidForWorkEnrollmentProfiles
}
// GetAndroidForWorkSettings gets the androidForWorkSettings property value. The singleton Android for Work settings entity.
func (m *DeviceManagement) GetAndroidForWorkSettings()(AndroidForWorkSettingsable) {
    return m.androidForWorkSettings
}
// GetAndroidManagedStoreAccountEnterpriseSettings gets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
func (m *DeviceManagement) GetAndroidManagedStoreAccountEnterpriseSettings()(AndroidManagedStoreAccountEnterpriseSettingsable) {
    return m.androidManagedStoreAccountEnterpriseSettings
}
// GetAndroidManagedStoreAppConfigurationSchemas gets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
func (m *DeviceManagement) GetAndroidManagedStoreAppConfigurationSchemas()([]AndroidManagedStoreAppConfigurationSchemaable) {
    return m.androidManagedStoreAppConfigurationSchemas
}
// GetApplePushNotificationCertificate gets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable) {
    return m.applePushNotificationCertificate
}
// GetAppleUserInitiatedEnrollmentProfiles gets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
func (m *DeviceManagement) GetAppleUserInitiatedEnrollmentProfiles()([]AppleUserInitiatedEnrollmentProfileable) {
    return m.appleUserInitiatedEnrollmentProfiles
}
// GetAssignmentFilters gets the assignmentFilters property value. The list of assignment filters
func (m *DeviceManagement) GetAssignmentFilters()([]DeviceAndAppManagementAssignmentFilterable) {
    return m.assignmentFilters
}
// GetAuditEvents gets the auditEvents property value. The Audit Events
func (m *DeviceManagement) GetAuditEvents()([]AuditEventable) {
    return m.auditEvents
}
// GetAutopilotEvents gets the autopilotEvents property value. The list of autopilot events for the tenant.
func (m *DeviceManagement) GetAutopilotEvents()([]DeviceManagementAutopilotEventable) {
    return m.autopilotEvents
}
// GetCartToClassAssociations gets the cartToClassAssociations property value. The Cart To Class Associations.
func (m *DeviceManagement) GetCartToClassAssociations()([]CartToClassAssociationable) {
    return m.cartToClassAssociations
}
// GetCategories gets the categories property value. The available categories
func (m *DeviceManagement) GetCategories()([]DeviceManagementSettingCategoryable) {
    return m.categories
}
// GetCertificateConnectorDetails gets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagement) GetCertificateConnectorDetails()([]CertificateConnectorDetailsable) {
    return m.certificateConnectorDetails
}
// GetChromeOSOnboardingSettings gets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
func (m *DeviceManagement) GetChromeOSOnboardingSettings()([]ChromeOSOnboardingSettingsable) {
    return m.chromeOSOnboardingSettings
}
// GetCloudPCConnectivityIssues gets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
func (m *DeviceManagement) GetCloudPCConnectivityIssues()([]CloudPCConnectivityIssueable) {
    return m.cloudPCConnectivityIssues
}
// GetComanagedDevices gets the comanagedDevices property value. The list of co-managed devices report
func (m *DeviceManagement) GetComanagedDevices()([]ManagedDeviceable) {
    return m.comanagedDevices
}
// GetComanagementEligibleDevices gets the comanagementEligibleDevices property value. The list of co-management eligible devices report
func (m *DeviceManagement) GetComanagementEligibleDevices()([]ComanagementEligibleDeviceable) {
    return m.comanagementEligibleDevices
}
// GetComplianceCategories gets the complianceCategories property value. List of all compliance categories
func (m *DeviceManagement) GetComplianceCategories()([]DeviceManagementConfigurationCategoryable) {
    return m.complianceCategories
}
// GetComplianceManagementPartners gets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartnerable) {
    return m.complianceManagementPartners
}
// GetCompliancePolicies gets the compliancePolicies property value. List of all compliance policies
func (m *DeviceManagement) GetCompliancePolicies()([]DeviceManagementCompliancePolicyable) {
    return m.compliancePolicies
}
// GetComplianceSettings gets the complianceSettings property value. List of all ComplianceSettings
func (m *DeviceManagement) GetComplianceSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.complianceSettings
}
// GetConditionalAccessSettings gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable) {
    return m.conditionalAccessSettings
}
// GetConfigManagerCollections gets the configManagerCollections property value. A list of ConfigManagerCollection
func (m *DeviceManagement) GetConfigManagerCollections()([]ConfigManagerCollectionable) {
    return m.configManagerCollections
}
// GetConfigurationCategories gets the configurationCategories property value. List of all Configuration Categories
func (m *DeviceManagement) GetConfigurationCategories()([]DeviceManagementConfigurationCategoryable) {
    return m.configurationCategories
}
// GetConfigurationPolicies gets the configurationPolicies property value. List of all Configuration policies
func (m *DeviceManagement) GetConfigurationPolicies()([]DeviceManagementConfigurationPolicyable) {
    return m.configurationPolicies
}
// GetConfigurationPolicyTemplates gets the configurationPolicyTemplates property value. List of all templates
func (m *DeviceManagement) GetConfigurationPolicyTemplates()([]DeviceManagementConfigurationPolicyTemplateable) {
    return m.configurationPolicyTemplates
}
// GetConfigurationSettings gets the configurationSettings property value. List of all ConfigurationSettings
func (m *DeviceManagement) GetConfigurationSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.configurationSettings
}
// GetDataProcessorServiceForWindowsFeaturesOnboarding gets the dataProcessorServiceForWindowsFeaturesOnboarding property value. A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
func (m *DeviceManagement) GetDataProcessorServiceForWindowsFeaturesOnboarding()(DataProcessorServiceForWindowsFeaturesOnboardingable) {
    return m.dataProcessorServiceForWindowsFeaturesOnboarding
}
// GetDataSharingConsents gets the dataSharingConsents property value. Data sharing consents.
func (m *DeviceManagement) GetDataSharingConsents()([]DataSharingConsentable) {
    return m.dataSharingConsents
}
// GetDepOnboardingSettings gets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
func (m *DeviceManagement) GetDepOnboardingSettings()([]DepOnboardingSettingable) {
    return m.depOnboardingSettings
}
// GetDerivedCredentials gets the derivedCredentials property value. Collection of Derived credential settings associated with account.
func (m *DeviceManagement) GetDerivedCredentials()([]DeviceManagementDerivedCredentialSettingsable) {
    return m.derivedCredentials
}
// GetDetectedApps gets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) GetDetectedApps()([]DetectedAppable) {
    return m.detectedApps
}
// GetDeviceCategories gets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategoryable) {
    return m.deviceCategories
}
// GetDeviceCompliancePolicies gets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable) {
    return m.deviceCompliancePolicies
}
// GetDeviceCompliancePolicyDeviceStateSummary gets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable) {
    return m.deviceCompliancePolicyDeviceStateSummary
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    return m.deviceCompliancePolicySettingStateSummaries
}
// GetDeviceComplianceReportSummarizationDateTime gets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
func (m *DeviceManagement) GetDeviceComplianceReportSummarizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deviceComplianceReportSummarizationDateTime
}
// GetDeviceComplianceScripts gets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceComplianceScripts()([]DeviceComplianceScriptable) {
    return m.deviceComplianceScripts
}
// GetDeviceConfigurationConflictSummary gets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
func (m *DeviceManagement) GetDeviceConfigurationConflictSummary()([]DeviceConfigurationConflictSummaryable) {
    return m.deviceConfigurationConflictSummary
}
// GetDeviceConfigurationDeviceStateSummaries gets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable) {
    return m.deviceConfigurationDeviceStateSummaries
}
// GetDeviceConfigurationRestrictedAppsViolations gets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
func (m *DeviceManagement) GetDeviceConfigurationRestrictedAppsViolations()([]RestrictedAppsViolationable) {
    return m.deviceConfigurationRestrictedAppsViolations
}
// GetDeviceConfigurations gets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfigurationable) {
    return m.deviceConfigurations
}
// GetDeviceConfigurationsAllManagedDeviceCertificateStates gets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
func (m *DeviceManagement) GetDeviceConfigurationsAllManagedDeviceCertificateStates()([]ManagedAllDeviceCertificateStateable) {
    return m.deviceConfigurationsAllManagedDeviceCertificateStates
}
// GetDeviceConfigurationUserStateSummaries gets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationUserStateSummaries()(DeviceConfigurationUserStateSummaryable) {
    return m.deviceConfigurationUserStateSummaries
}
// GetDeviceCustomAttributeShellScripts gets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceCustomAttributeShellScripts()([]DeviceCustomAttributeShellScriptable) {
    return m.deviceCustomAttributeShellScripts
}
// GetDeviceEnrollmentConfigurations gets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable) {
    return m.deviceEnrollmentConfigurations
}
// GetDeviceHealthScripts gets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceHealthScripts()([]DeviceHealthScriptable) {
    return m.deviceHealthScripts
}
// GetDeviceManagementPartners gets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartnerable) {
    return m.deviceManagementPartners
}
// GetDeviceManagementScripts gets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceManagementScripts()([]DeviceManagementScriptable) {
    return m.deviceManagementScripts
}
// GetDeviceProtectionOverview gets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) GetDeviceProtectionOverview()(DeviceProtectionOverviewable) {
    return m.deviceProtectionOverview
}
// GetDeviceShellScripts gets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceShellScripts()([]DeviceShellScriptable) {
    return m.deviceShellScripts
}
// GetDomainJoinConnectors gets the domainJoinConnectors property value. A list of connector objects.
func (m *DeviceManagement) GetDomainJoinConnectors()([]DeviceManagementDomainJoinConnectorable) {
    return m.domainJoinConnectors
}
// GetEmbeddedSIMActivationCodePools gets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
func (m *DeviceManagement) GetEmbeddedSIMActivationCodePools()([]EmbeddedSIMActivationCodePoolable) {
    return m.embeddedSIMActivationCodePools
}
// GetExchangeConnectors gets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnectorable) {
    return m.exchangeConnectors
}
// GetExchangeOnPremisesPolicies gets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagement) GetExchangeOnPremisesPolicies()([]DeviceManagementExchangeOnPremisesPolicyable) {
    return m.exchangeOnPremisesPolicies
}
// GetExchangeOnPremisesPolicy gets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
func (m *DeviceManagement) GetExchangeOnPremisesPolicy()(DeviceManagementExchangeOnPremisesPolicyable) {
    return m.exchangeOnPremisesPolicy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountMoveCompletionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountMoveCompletionDateTime(val)
        }
        return nil
    }
    res["adminConsent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminConsentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsent(val.(AdminConsentable))
        }
        return nil
    }
    res["advancedThreatProtectionOnboardingStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOnboardingStateSummary(val.(AdvancedThreatProtectionOnboardingStateSummaryable))
        }
        return nil
    }
    res["androidDeviceOwnerEnrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerEnrollmentProfileable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidDeviceOwnerEnrollmentProfileable)
            }
            m.SetAndroidDeviceOwnerEnrollmentProfiles(res)
        }
        return nil
    }
    res["androidForWorkAppConfigurationSchemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidForWorkAppConfigurationSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidForWorkAppConfigurationSchemaable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidForWorkAppConfigurationSchemaable)
            }
            m.SetAndroidForWorkAppConfigurationSchemas(res)
        }
        return nil
    }
    res["androidForWorkEnrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidForWorkEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidForWorkEnrollmentProfileable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidForWorkEnrollmentProfileable)
            }
            m.SetAndroidForWorkEnrollmentProfiles(res)
        }
        return nil
    }
    res["androidForWorkSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidForWorkSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidForWorkSettings(val.(AndroidForWorkSettingsable))
        }
        return nil
    }
    res["androidManagedStoreAccountEnterpriseSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidManagedStoreAccountEnterpriseSettings(val.(AndroidManagedStoreAccountEnterpriseSettingsable))
        }
        return nil
    }
    res["androidManagedStoreAppConfigurationSchemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppConfigurationSchemaable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidManagedStoreAppConfigurationSchemaable)
            }
            m.SetAndroidManagedStoreAppConfigurationSchemas(res)
        }
        return nil
    }
    res["applePushNotificationCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplePushNotificationCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplePushNotificationCertificate(val.(ApplePushNotificationCertificateable))
        }
        return nil
    }
    res["appleUserInitiatedEnrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleUserInitiatedEnrollmentProfileable, len(val))
            for i, v := range val {
                res[i] = v.(AppleUserInitiatedEnrollmentProfileable)
            }
            m.SetAppleUserInitiatedEnrollmentProfiles(res)
        }
        return nil
    }
    res["assignmentFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementAssignmentFilterable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceAndAppManagementAssignmentFilterable)
            }
            m.SetAssignmentFilters(res)
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
    res["autopilotEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementAutopilotEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementAutopilotEventable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementAutopilotEventable)
            }
            m.SetAutopilotEvents(res)
        }
        return nil
    }
    res["cartToClassAssociations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCartToClassAssociationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CartToClassAssociationable, len(val))
            for i, v := range val {
                res[i] = v.(CartToClassAssociationable)
            }
            m.SetCartToClassAssociations(res)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingCategoryable)
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["certificateConnectorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateConnectorDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateConnectorDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(CertificateConnectorDetailsable)
            }
            m.SetCertificateConnectorDetails(res)
        }
        return nil
    }
    res["chromeOSOnboardingSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChromeOSOnboardingSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChromeOSOnboardingSettingsable, len(val))
            for i, v := range val {
                res[i] = v.(ChromeOSOnboardingSettingsable)
            }
            m.SetChromeOSOnboardingSettings(res)
        }
        return nil
    }
    res["cloudPCConnectivityIssues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPCConnectivityIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPCConnectivityIssueable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPCConnectivityIssueable)
            }
            m.SetCloudPCConnectivityIssues(res)
        }
        return nil
    }
    res["comanagedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceable)
            }
            m.SetComanagedDevices(res)
        }
        return nil
    }
    res["comanagementEligibleDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComanagementEligibleDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComanagementEligibleDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(ComanagementEligibleDeviceable)
            }
            m.SetComanagementEligibleDevices(res)
        }
        return nil
    }
    res["complianceCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationCategoryable)
            }
            m.SetComplianceCategories(res)
        }
        return nil
    }
    res["complianceManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceManagementPartnerable)
            }
            m.SetComplianceManagementPartners(res)
        }
        return nil
    }
    res["compliancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementCompliancePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementCompliancePolicyable)
            }
            m.SetCompliancePolicies(res)
        }
        return nil
    }
    res["complianceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
            }
            m.SetComplianceSettings(res)
        }
        return nil
    }
    res["conditionalAccessSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessSettings(val.(OnPremisesConditionalAccessSettingsable))
        }
        return nil
    }
    res["configManagerCollections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConfigManagerCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConfigManagerCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(ConfigManagerCollectionable)
            }
            m.SetConfigManagerCollections(res)
        }
        return nil
    }
    res["configurationCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationCategoryable)
            }
            m.SetConfigurationCategories(res)
        }
        return nil
    }
    res["configurationPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationPolicyable)
            }
            m.SetConfigurationPolicies(res)
        }
        return nil
    }
    res["configurationPolicyTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationPolicyTemplateable)
            }
            m.SetConfigurationPolicyTemplates(res)
        }
        return nil
    }
    res["configurationSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
            }
            m.SetConfigurationSettings(res)
        }
        return nil
    }
    res["dataProcessorServiceForWindowsFeaturesOnboarding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDataProcessorServiceForWindowsFeaturesOnboardingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataProcessorServiceForWindowsFeaturesOnboarding(val.(DataProcessorServiceForWindowsFeaturesOnboardingable))
        }
        return nil
    }
    res["dataSharingConsents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSharingConsentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSharingConsentable, len(val))
            for i, v := range val {
                res[i] = v.(DataSharingConsentable)
            }
            m.SetDataSharingConsents(res)
        }
        return nil
    }
    res["depOnboardingSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDepOnboardingSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DepOnboardingSettingable, len(val))
            for i, v := range val {
                res[i] = v.(DepOnboardingSettingable)
            }
            m.SetDepOnboardingSettings(res)
        }
        return nil
    }
    res["derivedCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementDerivedCredentialSettingsable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementDerivedCredentialSettingsable)
            }
            m.SetDerivedCredentials(res)
        }
        return nil
    }
    res["detectedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDetectedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DetectedAppable, len(val))
            for i, v := range val {
                res[i] = v.(DetectedAppable)
            }
            m.SetDetectedApps(res)
        }
        return nil
    }
    res["deviceCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCategoryable)
            }
            m.SetDeviceCategories(res)
        }
        return nil
    }
    res["deviceCompliancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCompliancePolicyable)
            }
            m.SetDeviceCompliancePolicies(res)
        }
        return nil
    }
    res["deviceCompliancePolicyDeviceStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyDeviceStateSummary(val.(DeviceCompliancePolicyDeviceStateSummaryable))
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
    res["deviceComplianceReportSummarizationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceReportSummarizationDateTime(val)
        }
        return nil
    }
    res["deviceComplianceScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptable)
            }
            m.SetDeviceComplianceScripts(res)
        }
        return nil
    }
    res["deviceConfigurationConflictSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationConflictSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationConflictSummaryable)
            }
            m.SetDeviceConfigurationConflictSummary(res)
        }
        return nil
    }
    res["deviceConfigurationDeviceStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationDeviceStateSummaries(val.(DeviceConfigurationDeviceStateSummaryable))
        }
        return nil
    }
    res["deviceConfigurationRestrictedAppsViolations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRestrictedAppsViolationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RestrictedAppsViolationable, len(val))
            for i, v := range val {
                res[i] = v.(RestrictedAppsViolationable)
            }
            m.SetDeviceConfigurationRestrictedAppsViolations(res)
        }
        return nil
    }
    res["deviceConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationable)
            }
            m.SetDeviceConfigurations(res)
        }
        return nil
    }
    res["deviceConfigurationsAllManagedDeviceCertificateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAllDeviceCertificateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAllDeviceCertificateStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAllDeviceCertificateStateable)
            }
            m.SetDeviceConfigurationsAllManagedDeviceCertificateStates(res)
        }
        return nil
    }
    res["deviceConfigurationUserStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationUserStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationUserStateSummaries(val.(DeviceConfigurationUserStateSummaryable))
        }
        return nil
    }
    res["deviceCustomAttributeShellScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCustomAttributeShellScriptable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCustomAttributeShellScriptable)
            }
            m.SetDeviceCustomAttributeShellScripts(res)
        }
        return nil
    }
    res["deviceEnrollmentConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceEnrollmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceEnrollmentConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceEnrollmentConfigurationable)
            }
            m.SetDeviceEnrollmentConfigurations(res)
        }
        return nil
    }
    res["deviceHealthScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceHealthScriptable)
            }
            m.SetDeviceHealthScripts(res)
        }
        return nil
    }
    res["deviceManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementPartnerable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementPartnerable)
            }
            m.SetDeviceManagementPartners(res)
        }
        return nil
    }
    res["deviceManagementScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementScriptable)
            }
            m.SetDeviceManagementScripts(res)
        }
        return nil
    }
    res["deviceProtectionOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceProtectionOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceProtectionOverview(val.(DeviceProtectionOverviewable))
        }
        return nil
    }
    res["deviceShellScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceShellScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceShellScriptable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceShellScriptable)
            }
            m.SetDeviceShellScripts(res)
        }
        return nil
    }
    res["domainJoinConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementDomainJoinConnectorable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementDomainJoinConnectorable)
            }
            m.SetDomainJoinConnectors(res)
        }
        return nil
    }
    res["embeddedSIMActivationCodePools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMActivationCodePoolable, len(val))
            for i, v := range val {
                res[i] = v.(EmbeddedSIMActivationCodePoolable)
            }
            m.SetEmbeddedSIMActivationCodePools(res)
        }
        return nil
    }
    res["exchangeConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExchangeConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeConnectorable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementExchangeConnectorable)
            }
            m.SetExchangeConnectors(res)
        }
        return nil
    }
    res["exchangeOnPremisesPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeOnPremisesPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementExchangeOnPremisesPolicyable)
            }
            m.SetExchangeOnPremisesPolicies(res)
        }
        return nil
    }
    res["exchangeOnPremisesPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeOnPremisesPolicy(val.(DeviceManagementExchangeOnPremisesPolicyable))
        }
        return nil
    }
    res["groupPolicyCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyCategoryable)
            }
            m.SetGroupPolicyCategories(res)
        }
        return nil
    }
    res["groupPolicyConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyConfigurationable)
            }
            m.SetGroupPolicyConfigurations(res)
        }
        return nil
    }
    res["groupPolicyDefinitionFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionFileable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyDefinitionFileable)
            }
            m.SetGroupPolicyDefinitionFiles(res)
        }
        return nil
    }
    res["groupPolicyDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyDefinitionable)
            }
            m.SetGroupPolicyDefinitions(res)
        }
        return nil
    }
    res["groupPolicyMigrationReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyMigrationReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyMigrationReportable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyMigrationReportable)
            }
            m.SetGroupPolicyMigrationReports(res)
        }
        return nil
    }
    res["groupPolicyObjectFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyObjectFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyObjectFileable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyObjectFileable)
            }
            m.SetGroupPolicyObjectFiles(res)
        }
        return nil
    }
    res["groupPolicyUploadedDefinitionFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyUploadedDefinitionFileable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyUploadedDefinitionFileable)
            }
            m.SetGroupPolicyUploadedDefinitionFiles(res)
        }
        return nil
    }
    res["importedDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ImportedDeviceIdentityable)
            }
            m.SetImportedDeviceIdentities(res)
        }
        return nil
    }
    res["importedWindowsAutopilotDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedWindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ImportedWindowsAutopilotDeviceIdentityable)
            }
            m.SetImportedWindowsAutopilotDeviceIdentities(res)
        }
        return nil
    }
    res["intents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementIntentable)
            }
            m.SetIntents(res)
        }
        return nil
    }
    res["intuneAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneAccountId(val)
        }
        return nil
    }
    res["intuneBrand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntuneBrandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneBrand(val.(IntuneBrandable))
        }
        return nil
    }
    res["intuneBrandingProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntuneBrandingProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntuneBrandingProfileable, len(val))
            for i, v := range val {
                res[i] = v.(IntuneBrandingProfileable)
            }
            m.SetIntuneBrandingProfiles(res)
        }
        return nil
    }
    res["iosUpdateStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosUpdateDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosUpdateDeviceStatusable, len(val))
            for i, v := range val {
                res[i] = v.(IosUpdateDeviceStatusable)
            }
            m.SetIosUpdateStatuses(res)
        }
        return nil
    }
    res["lastReportAggregationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportAggregationDateTime(val)
        }
        return nil
    }
    res["legacyPcManangementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegacyPcManangementEnabled(val)
        }
        return nil
    }
    res["macOSSoftwareUpdateAccountSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSoftwareUpdateAccountSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSSoftwareUpdateAccountSummaryable)
            }
            m.SetMacOSSoftwareUpdateAccountSummaries(res)
        }
        return nil
    }
    res["managedDeviceCleanupSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceCleanupSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceCleanupSettings(val.(ManagedDeviceCleanupSettingsable))
        }
        return nil
    }
    res["managedDeviceEncryptionStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceEncryptionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceEncryptionStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceEncryptionStateable)
            }
            m.SetManagedDeviceEncryptionStates(res)
        }
        return nil
    }
    res["managedDeviceOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOverview(val.(ManagedDeviceOverviewable))
        }
        return nil
    }
    res["managedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceable)
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    res["maximumDepTokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumDepTokens(val)
        }
        return nil
    }
    res["microsoftTunnelConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftTunnelConfigurationable)
            }
            m.SetMicrosoftTunnelConfigurations(res)
        }
        return nil
    }
    res["microsoftTunnelHealthThresholds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelHealthThresholdable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftTunnelHealthThresholdable)
            }
            m.SetMicrosoftTunnelHealthThresholds(res)
        }
        return nil
    }
    res["microsoftTunnelServerLogCollectionResponses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelServerLogCollectionResponseable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftTunnelServerLogCollectionResponseable)
            }
            m.SetMicrosoftTunnelServerLogCollectionResponses(res)
        }
        return nil
    }
    res["microsoftTunnelSites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelSiteable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftTunnelSiteable)
            }
            m.SetMicrosoftTunnelSites(res)
        }
        return nil
    }
    res["mobileAppTroubleshootingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppTroubleshootingEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppTroubleshootingEventable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppTroubleshootingEventable)
            }
            m.SetMobileAppTroubleshootingEvents(res)
        }
        return nil
    }
    res["mobileThreatDefenseConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileThreatDefenseConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileThreatDefenseConnectorable, len(val))
            for i, v := range val {
                res[i] = v.(MobileThreatDefenseConnectorable)
            }
            m.SetMobileThreatDefenseConnectors(res)
        }
        return nil
    }
    res["ndesConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNdesConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NdesConnectorable, len(val))
            for i, v := range val {
                res[i] = v.(NdesConnectorable)
            }
            m.SetNdesConnectors(res)
        }
        return nil
    }
    res["notificationMessageTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationMessageTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationMessageTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(NotificationMessageTemplateable)
            }
            m.SetNotificationMessageTemplates(res)
        }
        return nil
    }
    res["oemWarrantyInformationOnboarding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOemWarrantyInformationOnboardingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OemWarrantyInformationOnboardingable, len(val))
            for i, v := range val {
                res[i] = v.(OemWarrantyInformationOnboardingable)
            }
            m.SetOemWarrantyInformationOnboarding(res)
        }
        return nil
    }
    res["remoteActionAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRemoteActionAuditFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RemoteActionAuditable, len(val))
            for i, v := range val {
                res[i] = v.(RemoteActionAuditable)
            }
            m.SetRemoteActionAudits(res)
        }
        return nil
    }
    res["remoteAssistancePartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRemoteAssistancePartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RemoteAssistancePartnerable, len(val))
            for i, v := range val {
                res[i] = v.(RemoteAssistancePartnerable)
            }
            m.SetRemoteAssistancePartners(res)
        }
        return nil
    }
    res["remoteAssistanceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRemoteAssistanceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSettings(val.(RemoteAssistanceSettingsable))
        }
        return nil
    }
    res["reports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementReportsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReports(val.(DeviceManagementReportsable))
        }
        return nil
    }
    res["resourceAccessProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementResourceAccessProfileBaseable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementResourceAccessProfileBaseable)
            }
            m.SetResourceAccessProfiles(res)
        }
        return nil
    }
    res["resourceOperations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceOperationable, len(val))
            for i, v := range val {
                res[i] = v.(ResourceOperationable)
            }
            m.SetResourceOperations(res)
        }
        return nil
    }
    res["reusablePolicySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementReusablePolicySettingable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementReusablePolicySettingable)
            }
            m.SetReusablePolicySettings(res)
        }
        return nil
    }
    res["reusableSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
            }
            m.SetReusableSettings(res)
        }
        return nil
    }
    res["roleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementRoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceAndAppManagementRoleAssignmentable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(RoleDefinitionable)
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleScopeTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleScopeTagable, len(val))
            for i, v := range val {
                res[i] = v.(RoleScopeTagable)
            }
            m.SetRoleScopeTags(res)
        }
        return nil
    }
    res["settingDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingDefinitionable)
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(DeviceManagementSettingsable))
        }
        return nil
    }
    res["softwareUpdateStatusSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSoftwareUpdateStatusSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateStatusSummary(val.(SoftwareUpdateStatusSummaryable))
        }
        return nil
    }
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptions(val.(*DeviceManagementSubscriptions))
        }
        return nil
    }
    res["subscriptionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionState(val.(*DeviceManagementSubscriptionState))
        }
        return nil
    }
    res["telecomExpenseManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTelecomExpenseManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TelecomExpenseManagementPartnerable, len(val))
            for i, v := range val {
                res[i] = v.(TelecomExpenseManagementPartnerable)
            }
            m.SetTelecomExpenseManagementPartners(res)
        }
        return nil
    }
    res["templates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementTemplateable)
            }
            m.SetTemplates(res)
        }
        return nil
    }
    res["templateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingTemplateable)
            }
            m.SetTemplateSettings(res)
        }
        return nil
    }
    res["tenantAttachRBAC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantAttachRBACFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAttachRBAC(val.(TenantAttachRBACable))
        }
        return nil
    }
    res["termsAndConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTermsAndConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TermsAndConditionsable, len(val))
            for i, v := range val {
                res[i] = v.(TermsAndConditionsable)
            }
            m.SetTermsAndConditions(res)
        }
        return nil
    }
    res["troubleshootingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingEventable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementTroubleshootingEventable)
            }
            m.SetTroubleshootingEvents(res)
        }
        return nil
    }
    res["unlicensedAdminstratorsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlicensedAdminstratorsEnabled(val)
        }
        return nil
    }
    res["userExperienceAnalyticsAnomaly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAnomalyable)
            }
            m.SetUserExperienceAnalyticsAnomaly(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAnomalyDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAnomalyDeviceable)
            }
            m.SetUserExperienceAnalyticsAnomalyDevice(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAnomalySeverityOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsAnomalySeverityOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsAnomalySeverityOverview(val.(UserExperienceAnalyticsAnomalySeverityOverviewable))
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthApplicationPerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthApplicationPerformanceable)
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDeviceModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)
            }
            m.SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDevicePerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthDevicePerformanceable)
            }
            m.SetUserExperienceAnalyticsAppHealthDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformanceDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)
            }
            m.SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthOSVersionPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsAppHealthOSVersionPerformanceable)
            }
            m.SetUserExperienceAnalyticsAppHealthOSVersionPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsAppHealthOverview(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    res["userExperienceAnalyticsBaselines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBaselineable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsBaselineable)
            }
            m.SetUserExperienceAnalyticsBaselines(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthAppImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthAppImpactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthAppImpactable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsBatteryHealthAppImpactable)
            }
            m.SetUserExperienceAnalyticsBatteryHealthAppImpact(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthCapacityDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsBatteryHealthCapacityDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsBatteryHealthCapacityDetails(val.(UserExperienceAnalyticsBatteryHealthCapacityDetailsable))
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDeviceAppImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)
            }
            m.SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsBatteryHealthDevicePerformanceable)
            }
            m.SetUserExperienceAnalyticsBatteryHealthDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)
            }
            m.SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthModelPerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsBatteryHealthModelPerformanceable)
            }
            m.SetUserExperienceAnalyticsBatteryHealthModelPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthOsPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthOsPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthOsPerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsBatteryHealthOsPerformanceable)
            }
            m.SetUserExperienceAnalyticsBatteryHealthOsPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthRuntimeDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsBatteryHealthRuntimeDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(val.(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable))
        }
        return nil
    }
    res["userExperienceAnalyticsCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsCategoryable)
            }
            m.SetUserExperienceAnalyticsCategories(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceMetricHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetricHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsMetricHistoryable)
            }
            m.SetUserExperienceAnalyticsDeviceMetricHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDevicePerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDevicePerformanceable)
            }
            m.SetUserExperienceAnalyticsDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsDeviceScope(val.(UserExperienceAnalyticsDeviceScopeable))
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceScopeable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDeviceScopeable)
            }
            m.SetUserExperienceAnalyticsDeviceScopes(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceScoresable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDeviceScoresable)
            }
            m.SetUserExperienceAnalyticsDeviceScores(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDeviceStartupHistoryable)
            }
            m.SetUserExperienceAnalyticsDeviceStartupHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcesses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupProcessable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDeviceStartupProcessable)
            }
            m.SetUserExperienceAnalyticsDeviceStartupProcesses(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcessPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDeviceStartupProcessPerformanceable)
            }
            m.SetUserExperienceAnalyticsDeviceStartupProcessPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDevicesWithoutCloudIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDeviceWithoutCloudIdentityable)
            }
            m.SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceTimelineEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceTimelineEventsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceTimelineEventsable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsDeviceTimelineEventsable)
            }
            m.SetUserExperienceAnalyticsDeviceTimelineEvents(res)
        }
        return nil
    }
    res["userExperienceAnalyticsImpactingProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsImpactingProcessable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsImpactingProcessable)
            }
            m.SetUserExperienceAnalyticsImpactingProcess(res)
        }
        return nil
    }
    res["userExperienceAnalyticsMetricHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetricHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsMetricHistoryable)
            }
            m.SetUserExperienceAnalyticsMetricHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsModelScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsModelScoresable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsModelScoresable)
            }
            m.SetUserExperienceAnalyticsModelScores(res)
        }
        return nil
    }
    res["userExperienceAnalyticsNotAutopilotReadyDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsNotAutopilotReadyDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsNotAutopilotReadyDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsNotAutopilotReadyDeviceable)
            }
            m.SetUserExperienceAnalyticsNotAutopilotReadyDevice(res)
        }
        return nil
    }
    res["userExperienceAnalyticsOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsOverview(val.(UserExperienceAnalyticsOverviewable))
        }
        return nil
    }
    res["userExperienceAnalyticsRemoteConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsRemoteConnectionable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsRemoteConnectionable)
            }
            m.SetUserExperienceAnalyticsRemoteConnection(res)
        }
        return nil
    }
    res["userExperienceAnalyticsResourcePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsResourcePerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsResourcePerformanceable)
            }
            m.SetUserExperienceAnalyticsResourcePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsScoreHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsScoreHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsScoreHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsScoreHistoryable)
            }
            m.SetUserExperienceAnalyticsScoreHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsSettings(val.(UserExperienceAnalyticsSettingsable))
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable))
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereMetricable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsWorkFromAnywhereMetricable)
            }
            m.SetUserExperienceAnalyticsWorkFromAnywhereMetrics(res)
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)
            }
            m.SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(res)
        }
        return nil
    }
    res["userPfxCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserPFXCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserPFXCertificateable, len(val))
            for i, v := range val {
                res[i] = v.(UserPFXCertificateable)
            }
            m.SetUserPfxCertificates(res)
        }
        return nil
    }
    res["virtualEndpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualEndpoint(val.(VirtualEndpointable))
        }
        return nil
    }
    res["windowsAutopilotDeploymentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeploymentProfileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsAutopilotDeploymentProfileable)
            }
            m.SetWindowsAutopilotDeploymentProfiles(res)
        }
        return nil
    }
    res["windowsAutopilotDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsAutopilotDeviceIdentityable)
            }
            m.SetWindowsAutopilotDeviceIdentities(res)
        }
        return nil
    }
    res["windowsAutopilotSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsAutopilotSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsAutopilotSettings(val.(WindowsAutopilotSettingsable))
        }
        return nil
    }
    res["windowsDriverUpdateProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDriverUpdateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDriverUpdateProfileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDriverUpdateProfileable)
            }
            m.SetWindowsDriverUpdateProfiles(res)
        }
        return nil
    }
    res["windowsFeatureUpdateProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsFeatureUpdateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsFeatureUpdateProfileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsFeatureUpdateProfileable)
            }
            m.SetWindowsFeatureUpdateProfiles(res)
        }
        return nil
    }
    res["windowsInformationProtectionAppLearningSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionAppLearningSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionAppLearningSummaryable)
            }
            m.SetWindowsInformationProtectionAppLearningSummaries(res)
        }
        return nil
    }
    res["windowsInformationProtectionNetworkLearningSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionNetworkLearningSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionNetworkLearningSummaryable)
            }
            m.SetWindowsInformationProtectionNetworkLearningSummaries(res)
        }
        return nil
    }
    res["windowsMalwareInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsMalwareInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsMalwareInformationable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsMalwareInformationable)
            }
            m.SetWindowsMalwareInformation(res)
        }
        return nil
    }
    res["windowsMalwareOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMalwareOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMalwareOverview(val.(WindowsMalwareOverviewable))
        }
        return nil
    }
    res["windowsQualityUpdateProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsQualityUpdateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsQualityUpdateProfileable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsQualityUpdateProfileable)
            }
            m.SetWindowsQualityUpdateProfiles(res)
        }
        return nil
    }
    res["windowsUpdateCatalogItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsUpdateCatalogItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsUpdateCatalogItemable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsUpdateCatalogItemable)
            }
            m.SetWindowsUpdateCatalogItems(res)
        }
        return nil
    }
    res["zebraFotaArtifacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateZebraFotaArtifactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ZebraFotaArtifactable, len(val))
            for i, v := range val {
                res[i] = v.(ZebraFotaArtifactable)
            }
            m.SetZebraFotaArtifacts(res)
        }
        return nil
    }
    res["zebraFotaConnector"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateZebraFotaConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZebraFotaConnector(val.(ZebraFotaConnectorable))
        }
        return nil
    }
    res["zebraFotaDeployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateZebraFotaDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ZebraFotaDeploymentable, len(val))
            for i, v := range val {
                res[i] = v.(ZebraFotaDeploymentable)
            }
            m.SetZebraFotaDeployments(res)
        }
        return nil
    }
    return res
}
// GetGroupPolicyCategories gets the groupPolicyCategories property value. The available group policy categories for this account.
func (m *DeviceManagement) GetGroupPolicyCategories()([]GroupPolicyCategoryable) {
    return m.groupPolicyCategories
}
// GetGroupPolicyConfigurations gets the groupPolicyConfigurations property value. The group policy configurations created by this account.
func (m *DeviceManagement) GetGroupPolicyConfigurations()([]GroupPolicyConfigurationable) {
    return m.groupPolicyConfigurations
}
// GetGroupPolicyDefinitionFiles gets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitionFiles()([]GroupPolicyDefinitionFileable) {
    return m.groupPolicyDefinitionFiles
}
// GetGroupPolicyDefinitions gets the groupPolicyDefinitions property value. The available group policy definitions for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitions()([]GroupPolicyDefinitionable) {
    return m.groupPolicyDefinitions
}
// GetGroupPolicyMigrationReports gets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
func (m *DeviceManagement) GetGroupPolicyMigrationReports()([]GroupPolicyMigrationReportable) {
    return m.groupPolicyMigrationReports
}
// GetGroupPolicyObjectFiles gets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
func (m *DeviceManagement) GetGroupPolicyObjectFiles()([]GroupPolicyObjectFileable) {
    return m.groupPolicyObjectFiles
}
// GetGroupPolicyUploadedDefinitionFiles gets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
func (m *DeviceManagement) GetGroupPolicyUploadedDefinitionFiles()([]GroupPolicyUploadedDefinitionFileable) {
    return m.groupPolicyUploadedDefinitionFiles
}
// GetImportedDeviceIdentities gets the importedDeviceIdentities property value. The imported device identities.
func (m *DeviceManagement) GetImportedDeviceIdentities()([]ImportedDeviceIdentityable) {
    return m.importedDeviceIdentities
}
// GetImportedWindowsAutopilotDeviceIdentities gets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable) {
    return m.importedWindowsAutopilotDeviceIdentities
}
// GetIntents gets the intents property value. The device management intents
func (m *DeviceManagement) GetIntents()([]DeviceManagementIntentable) {
    return m.intents
}
// GetIntuneAccountId gets the intuneAccountId property value. Intune Account ID for given tenant
func (m *DeviceManagement) GetIntuneAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.intuneAccountId
}
// GetIntuneBrand gets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) GetIntuneBrand()(IntuneBrandable) {
    return m.intuneBrand
}
// GetIntuneBrandingProfiles gets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
func (m *DeviceManagement) GetIntuneBrandingProfiles()([]IntuneBrandingProfileable) {
    return m.intuneBrandingProfiles
}
// GetIosUpdateStatuses gets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatusable) {
    return m.iosUpdateStatuses
}
// GetLastReportAggregationDateTime gets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
func (m *DeviceManagement) GetLastReportAggregationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportAggregationDateTime
}
// GetLegacyPcManangementEnabled gets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
func (m *DeviceManagement) GetLegacyPcManangementEnabled()(*bool) {
    return m.legacyPcManangementEnabled
}
// GetMacOSSoftwareUpdateAccountSummaries gets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
func (m *DeviceManagement) GetMacOSSoftwareUpdateAccountSummaries()([]MacOSSoftwareUpdateAccountSummaryable) {
    return m.macOSSoftwareUpdateAccountSummaries
}
// GetManagedDeviceCleanupSettings gets the managedDeviceCleanupSettings property value. Device cleanup rule
func (m *DeviceManagement) GetManagedDeviceCleanupSettings()(ManagedDeviceCleanupSettingsable) {
    return m.managedDeviceCleanupSettings
}
// GetManagedDeviceEncryptionStates gets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
func (m *DeviceManagement) GetManagedDeviceEncryptionStates()([]ManagedDeviceEncryptionStateable) {
    return m.managedDeviceEncryptionStates
}
// GetManagedDeviceOverview gets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) GetManagedDeviceOverview()(ManagedDeviceOverviewable) {
    return m.managedDeviceOverview
}
// GetManagedDevices gets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) GetManagedDevices()([]ManagedDeviceable) {
    return m.managedDevices
}
// GetMaximumDepTokens gets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
func (m *DeviceManagement) GetMaximumDepTokens()(*int32) {
    return m.maximumDepTokens
}
// GetMicrosoftTunnelConfigurations gets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelConfigurations()([]MicrosoftTunnelConfigurationable) {
    return m.microsoftTunnelConfigurations
}
// GetMicrosoftTunnelHealthThresholds gets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelHealthThresholds()([]MicrosoftTunnelHealthThresholdable) {
    return m.microsoftTunnelHealthThresholds
}
// GetMicrosoftTunnelServerLogCollectionResponses gets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelServerLogCollectionResponses()([]MicrosoftTunnelServerLogCollectionResponseable) {
    return m.microsoftTunnelServerLogCollectionResponses
}
// GetMicrosoftTunnelSites gets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelSites()([]MicrosoftTunnelSiteable) {
    return m.microsoftTunnelSites
}
// GetMobileAppTroubleshootingEvents gets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable) {
    return m.mobileAppTroubleshootingEvents
}
// GetMobileThreatDefenseConnectors gets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable) {
    return m.mobileThreatDefenseConnectors
}
// GetNdesConnectors gets the ndesConnectors property value. The collection of Ndes connectors for this account.
func (m *DeviceManagement) GetNdesConnectors()([]NdesConnectorable) {
    return m.ndesConnectors
}
// GetNotificationMessageTemplates gets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplateable) {
    return m.notificationMessageTemplates
}
// GetOemWarrantyInformationOnboarding gets the oemWarrantyInformationOnboarding property value. List of OEM Warranty Statuses
func (m *DeviceManagement) GetOemWarrantyInformationOnboarding()([]OemWarrantyInformationOnboardingable) {
    return m.oemWarrantyInformationOnboarding
}
// GetRemoteActionAudits gets the remoteActionAudits property value. The list of device remote action audits with the tenant.
func (m *DeviceManagement) GetRemoteActionAudits()([]RemoteActionAuditable) {
    return m.remoteActionAudits
}
// GetRemoteAssistancePartners gets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartnerable) {
    return m.remoteAssistancePartners
}
// GetRemoteAssistanceSettings gets the remoteAssistanceSettings property value. The remote assistance settings singleton
func (m *DeviceManagement) GetRemoteAssistanceSettings()(RemoteAssistanceSettingsable) {
    return m.remoteAssistanceSettings
}
// GetReports gets the reports property value. Reports singleton
func (m *DeviceManagement) GetReports()(DeviceManagementReportsable) {
    return m.reports
}
// GetResourceAccessProfiles gets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
func (m *DeviceManagement) GetResourceAccessProfiles()([]DeviceManagementResourceAccessProfileBaseable) {
    return m.resourceAccessProfiles
}
// GetResourceOperations gets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperationable) {
    return m.resourceOperations
}
// GetReusablePolicySettings gets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
func (m *DeviceManagement) GetReusablePolicySettings()([]DeviceManagementReusablePolicySettingable) {
    return m.reusablePolicySettings
}
// GetReusableSettings gets the reusableSettings property value. List of all reusable settings
func (m *DeviceManagement) GetReusableSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.reusableSettings
}
// GetRoleAssignments gets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable) {
    return m.roleAssignments
}
// GetRoleDefinitions gets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinitionable) {
    return m.roleDefinitions
}
// GetRoleScopeTags gets the roleScopeTags property value. The Role Scope Tags.
func (m *DeviceManagement) GetRoleScopeTags()([]RoleScopeTagable) {
    return m.roleScopeTags
}
// GetSettingDefinitions gets the settingDefinitions property value. The device management intent setting definitions
func (m *DeviceManagement) GetSettingDefinitions()([]DeviceManagementSettingDefinitionable) {
    return m.settingDefinitions
}
// GetSettings gets the settings property value. Account level settings.
func (m *DeviceManagement) GetSettings()(DeviceManagementSettingsable) {
    return m.settings
}
// GetSoftwareUpdateStatusSummary gets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable) {
    return m.softwareUpdateStatusSummary
}
// GetSubscriptions gets the subscriptions property value. Tenant mobile device management subscriptions.
func (m *DeviceManagement) GetSubscriptions()(*DeviceManagementSubscriptions) {
    return m.subscriptions
}
// GetSubscriptionState gets the subscriptionState property value. Tenant mobile device management subscription state.
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    return m.subscriptionState
}
// GetTelecomExpenseManagementPartners gets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable) {
    return m.telecomExpenseManagementPartners
}
// GetTemplates gets the templates property value. The available templates
func (m *DeviceManagement) GetTemplates()([]DeviceManagementTemplateable) {
    return m.templates
}
// GetTemplateSettings gets the templateSettings property value. List of all TemplateSettings
func (m *DeviceManagement) GetTemplateSettings()([]DeviceManagementConfigurationSettingTemplateable) {
    return m.templateSettings
}
// GetTenantAttachRBAC gets the tenantAttachRBAC property value. TenantAttach RBAC Enablement
func (m *DeviceManagement) GetTenantAttachRBAC()(TenantAttachRBACable) {
    return m.tenantAttachRBAC
}
// GetTermsAndConditions gets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditionsable) {
    return m.termsAndConditions
}
// GetTroubleshootingEvents gets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable) {
    return m.troubleshootingEvents
}
// GetUnlicensedAdminstratorsEnabled gets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
func (m *DeviceManagement) GetUnlicensedAdminstratorsEnabled()(*bool) {
    return m.unlicensedAdminstratorsEnabled
}
// GetUserExperienceAnalyticsAnomaly gets the userExperienceAnalyticsAnomaly property value. The user experience analytics anomaly entity contains anomaly details.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomaly()([]UserExperienceAnalyticsAnomalyable) {
    return m.userExperienceAnalyticsAnomaly
}
// GetUserExperienceAnalyticsAnomalyDevice gets the userExperienceAnalyticsAnomalyDevice property value. The user experience analytics anomaly entity contains device details.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomalyDevice()([]UserExperienceAnalyticsAnomalyDeviceable) {
    return m.userExperienceAnalyticsAnomalyDevice
}
// GetUserExperienceAnalyticsAnomalySeverityOverview gets the userExperienceAnalyticsAnomalySeverityOverview property value. The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomalySeverityOverview()(UserExperienceAnalyticsAnomalySeverityOverviewable) {
    return m.userExperienceAnalyticsAnomalySeverityOverview
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformance gets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformanceable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformance
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion gets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion
}
// GetUserExperienceAnalyticsAppHealthDeviceModelPerformance gets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable) {
    return m.userExperienceAnalyticsAppHealthDeviceModelPerformance
}
// GetUserExperienceAnalyticsAppHealthDevicePerformance gets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformanceable) {
    return m.userExperienceAnalyticsAppHealthDevicePerformance
}
// GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails gets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable) {
    return m.userExperienceAnalyticsAppHealthDevicePerformanceDetails
}
// GetUserExperienceAnalyticsAppHealthOSVersionPerformance gets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable) {
    return m.userExperienceAnalyticsAppHealthOSVersionPerformance
}
// GetUserExperienceAnalyticsAppHealthOverview gets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOverview()(UserExperienceAnalyticsCategoryable) {
    return m.userExperienceAnalyticsAppHealthOverview
}
// GetUserExperienceAnalyticsBaselines gets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaselineable) {
    return m.userExperienceAnalyticsBaselines
}
// GetUserExperienceAnalyticsBatteryHealthAppImpact gets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthAppImpact()([]UserExperienceAnalyticsBatteryHealthAppImpactable) {
    return m.userExperienceAnalyticsBatteryHealthAppImpact
}
// GetUserExperienceAnalyticsBatteryHealthCapacityDetails gets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthCapacityDetails()(UserExperienceAnalyticsBatteryHealthCapacityDetailsable) {
    return m.userExperienceAnalyticsBatteryHealthCapacityDetails
}
// GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact gets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable) {
    return m.userExperienceAnalyticsBatteryHealthDeviceAppImpact
}
// GetUserExperienceAnalyticsBatteryHealthDevicePerformance gets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDevicePerformance()([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable) {
    return m.userExperienceAnalyticsBatteryHealthDevicePerformance
}
// GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory gets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable) {
    return m.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
}
// GetUserExperienceAnalyticsBatteryHealthModelPerformance gets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthModelPerformance()([]UserExperienceAnalyticsBatteryHealthModelPerformanceable) {
    return m.userExperienceAnalyticsBatteryHealthModelPerformance
}
// GetUserExperienceAnalyticsBatteryHealthOsPerformance gets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthOsPerformance()([]UserExperienceAnalyticsBatteryHealthOsPerformanceable) {
    return m.userExperienceAnalyticsBatteryHealthOsPerformance
}
// GetUserExperienceAnalyticsBatteryHealthRuntimeDetails gets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthRuntimeDetails()(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable) {
    return m.userExperienceAnalyticsBatteryHealthRuntimeDetails
}
// GetUserExperienceAnalyticsCategories gets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategoryable) {
    return m.userExperienceAnalyticsCategories
}
// GetUserExperienceAnalyticsDeviceMetricHistory gets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceMetricHistory()([]UserExperienceAnalyticsMetricHistoryable) {
    return m.userExperienceAnalyticsDeviceMetricHistory
}
// GetUserExperienceAnalyticsDevicePerformance gets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformanceable) {
    return m.userExperienceAnalyticsDevicePerformance
}
// GetUserExperienceAnalyticsDeviceScope gets the userExperienceAnalyticsDeviceScope property value. The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScope()(UserExperienceAnalyticsDeviceScopeable) {
    return m.userExperienceAnalyticsDeviceScope
}
// GetUserExperienceAnalyticsDeviceScopes gets the userExperienceAnalyticsDeviceScopes property value. The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScopes()([]UserExperienceAnalyticsDeviceScopeable) {
    return m.userExperienceAnalyticsDeviceScopes
}
// GetUserExperienceAnalyticsDeviceScores gets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScoresable) {
    return m.userExperienceAnalyticsDeviceScores
}
// GetUserExperienceAnalyticsDeviceStartupHistory gets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistoryable) {
    return m.userExperienceAnalyticsDeviceStartupHistory
}
// GetUserExperienceAnalyticsDeviceStartupProcesses gets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcessable) {
    return m.userExperienceAnalyticsDeviceStartupProcesses
}
// GetUserExperienceAnalyticsDeviceStartupProcessPerformance gets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable) {
    return m.userExperienceAnalyticsDeviceStartupProcessPerformance
}
// GetUserExperienceAnalyticsDevicesWithoutCloudIdentity gets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable) {
    return m.userExperienceAnalyticsDevicesWithoutCloudIdentity
}
// GetUserExperienceAnalyticsDeviceTimelineEvents gets the userExperienceAnalyticsDeviceTimelineEvents property value. The user experience analytics device events entity contains NRT device timeline events details.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceTimelineEvents()([]UserExperienceAnalyticsDeviceTimelineEventsable) {
    return m.userExperienceAnalyticsDeviceTimelineEvents
}
// GetUserExperienceAnalyticsImpactingProcess gets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
func (m *DeviceManagement) GetUserExperienceAnalyticsImpactingProcess()([]UserExperienceAnalyticsImpactingProcessable) {
    return m.userExperienceAnalyticsImpactingProcess
}
// GetUserExperienceAnalyticsMetricHistory gets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistoryable) {
    return m.userExperienceAnalyticsMetricHistory
}
// GetUserExperienceAnalyticsModelScores gets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScoresable) {
    return m.userExperienceAnalyticsModelScores
}
// GetUserExperienceAnalyticsNotAutopilotReadyDevice gets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
func (m *DeviceManagement) GetUserExperienceAnalyticsNotAutopilotReadyDevice()([]UserExperienceAnalyticsNotAutopilotReadyDeviceable) {
    return m.userExperienceAnalyticsNotAutopilotReadyDevice
}
// GetUserExperienceAnalyticsOverview gets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) GetUserExperienceAnalyticsOverview()(UserExperienceAnalyticsOverviewable) {
    return m.userExperienceAnalyticsOverview
}
// GetUserExperienceAnalyticsRemoteConnection gets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
func (m *DeviceManagement) GetUserExperienceAnalyticsRemoteConnection()([]UserExperienceAnalyticsRemoteConnectionable) {
    return m.userExperienceAnalyticsRemoteConnection
}
// GetUserExperienceAnalyticsResourcePerformance gets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
func (m *DeviceManagement) GetUserExperienceAnalyticsResourcePerformance()([]UserExperienceAnalyticsResourcePerformanceable) {
    return m.userExperienceAnalyticsResourcePerformance
}
// GetUserExperienceAnalyticsScoreHistory gets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistoryable) {
    return m.userExperienceAnalyticsScoreHistory
}
// GetUserExperienceAnalyticsSettings gets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) GetUserExperienceAnalyticsSettings()(UserExperienceAnalyticsSettingsable) {
    return m.userExperienceAnalyticsSettings
}
// GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric gets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable) {
    return m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
}
// GetUserExperienceAnalyticsWorkFromAnywhereMetrics gets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetricable) {
    return m.userExperienceAnalyticsWorkFromAnywhereMetrics
}
// GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance gets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable) {
    return m.userExperienceAnalyticsWorkFromAnywhereModelPerformance
}
// GetUserPfxCertificates gets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
func (m *DeviceManagement) GetUserPfxCertificates()([]UserPFXCertificateable) {
    return m.userPfxCertificates
}
// GetVirtualEndpoint gets the virtualEndpoint property value. The virtualEndpoint property
func (m *DeviceManagement) GetVirtualEndpoint()(VirtualEndpointable) {
    return m.virtualEndpoint
}
// GetWindowsAutopilotDeploymentProfiles gets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
func (m *DeviceManagement) GetWindowsAutopilotDeploymentProfiles()([]WindowsAutopilotDeploymentProfileable) {
    return m.windowsAutopilotDeploymentProfiles
}
// GetWindowsAutopilotDeviceIdentities gets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable) {
    return m.windowsAutopilotDeviceIdentities
}
// GetWindowsAutopilotSettings gets the windowsAutopilotSettings property value. The Windows autopilot account settings.
func (m *DeviceManagement) GetWindowsAutopilotSettings()(WindowsAutopilotSettingsable) {
    return m.windowsAutopilotSettings
}
// GetWindowsDriverUpdateProfiles gets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
func (m *DeviceManagement) GetWindowsDriverUpdateProfiles()([]WindowsDriverUpdateProfileable) {
    return m.windowsDriverUpdateProfiles
}
// GetWindowsFeatureUpdateProfiles gets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
func (m *DeviceManagement) GetWindowsFeatureUpdateProfiles()([]WindowsFeatureUpdateProfileable) {
    return m.windowsFeatureUpdateProfiles
}
// GetWindowsInformationProtectionAppLearningSummaries gets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable) {
    return m.windowsInformationProtectionAppLearningSummaries
}
// GetWindowsInformationProtectionNetworkLearningSummaries gets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable) {
    return m.windowsInformationProtectionNetworkLearningSummaries
}
// GetWindowsMalwareInformation gets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) GetWindowsMalwareInformation()([]WindowsMalwareInformationable) {
    return m.windowsMalwareInformation
}
// GetWindowsMalwareOverview gets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) GetWindowsMalwareOverview()(WindowsMalwareOverviewable) {
    return m.windowsMalwareOverview
}
// GetWindowsQualityUpdateProfiles gets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
func (m *DeviceManagement) GetWindowsQualityUpdateProfiles()([]WindowsQualityUpdateProfileable) {
    return m.windowsQualityUpdateProfiles
}
// GetWindowsUpdateCatalogItems gets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
func (m *DeviceManagement) GetWindowsUpdateCatalogItems()([]WindowsUpdateCatalogItemable) {
    return m.windowsUpdateCatalogItems
}
// GetZebraFotaArtifacts gets the zebraFotaArtifacts property value. The Collection of ZebraFotaArtifacts.
func (m *DeviceManagement) GetZebraFotaArtifacts()([]ZebraFotaArtifactable) {
    return m.zebraFotaArtifacts
}
// GetZebraFotaConnector gets the zebraFotaConnector property value. The singleton ZebraFotaConnector associated with account.
func (m *DeviceManagement) GetZebraFotaConnector()(ZebraFotaConnectorable) {
    return m.zebraFotaConnector
}
// GetZebraFotaDeployments gets the zebraFotaDeployments property value. Collection of ZebraFotaDeployments associated with account.
func (m *DeviceManagement) GetZebraFotaDeployments()([]ZebraFotaDeploymentable) {
    return m.zebraFotaDeployments
}
// Serialize serializes information the current object
func (m *DeviceManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("accountMoveCompletionDateTime", m.GetAccountMoveCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("adminConsent", m.GetAdminConsent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("advancedThreatProtectionOnboardingStateSummary", m.GetAdvancedThreatProtectionOnboardingStateSummary())
        if err != nil {
            return err
        }
    }
    if m.GetAndroidDeviceOwnerEnrollmentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidDeviceOwnerEnrollmentProfiles()))
        for i, v := range m.GetAndroidDeviceOwnerEnrollmentProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("androidDeviceOwnerEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAndroidForWorkAppConfigurationSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidForWorkAppConfigurationSchemas()))
        for i, v := range m.GetAndroidForWorkAppConfigurationSchemas() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("androidForWorkAppConfigurationSchemas", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAndroidForWorkEnrollmentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidForWorkEnrollmentProfiles()))
        for i, v := range m.GetAndroidForWorkEnrollmentProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("androidForWorkEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("androidForWorkSettings", m.GetAndroidForWorkSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("androidManagedStoreAccountEnterpriseSettings", m.GetAndroidManagedStoreAccountEnterpriseSettings())
        if err != nil {
            return err
        }
    }
    if m.GetAndroidManagedStoreAppConfigurationSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidManagedStoreAppConfigurationSchemas()))
        for i, v := range m.GetAndroidManagedStoreAppConfigurationSchemas() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("androidManagedStoreAppConfigurationSchemas", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("applePushNotificationCertificate", m.GetApplePushNotificationCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetAppleUserInitiatedEnrollmentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppleUserInitiatedEnrollmentProfiles()))
        for i, v := range m.GetAppleUserInitiatedEnrollmentProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appleUserInitiatedEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentFilters()))
        for i, v := range m.GetAssignmentFilters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentFilters", cast)
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
    if m.GetAutopilotEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAutopilotEvents()))
        for i, v := range m.GetAutopilotEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("autopilotEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCartToClassAssociations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCartToClassAssociations()))
        for i, v := range m.GetCartToClassAssociations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cartToClassAssociations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertificateConnectorDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateConnectorDetails()))
        for i, v := range m.GetCertificateConnectorDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certificateConnectorDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChromeOSOnboardingSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChromeOSOnboardingSettings()))
        for i, v := range m.GetChromeOSOnboardingSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("chromeOSOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPCConnectivityIssues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPCConnectivityIssues()))
        for i, v := range m.GetCloudPCConnectivityIssues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPCConnectivityIssues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComanagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComanagedDevices()))
        for i, v := range m.GetComanagedDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("comanagedDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComanagementEligibleDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComanagementEligibleDevices()))
        for i, v := range m.GetComanagementEligibleDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("comanagementEligibleDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceCategories()))
        for i, v := range m.GetComplianceCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("complianceCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceManagementPartners()))
        for i, v := range m.GetComplianceManagementPartners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("complianceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompliancePolicies()))
        for i, v := range m.GetCompliancePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("compliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceSettings()))
        for i, v := range m.GetComplianceSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("complianceSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditionalAccessSettings", m.GetConditionalAccessSettings())
        if err != nil {
            return err
        }
    }
    if m.GetConfigManagerCollections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigManagerCollections()))
        for i, v := range m.GetConfigManagerCollections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configManagerCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationCategories()))
        for i, v := range m.GetConfigurationCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configurationCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationPolicies()))
        for i, v := range m.GetConfigurationPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configurationPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationPolicyTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationPolicyTemplates()))
        for i, v := range m.GetConfigurationPolicyTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configurationPolicyTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationSettings()))
        for i, v := range m.GetConfigurationSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("configurationSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataProcessorServiceForWindowsFeaturesOnboarding", m.GetDataProcessorServiceForWindowsFeaturesOnboarding())
        if err != nil {
            return err
        }
    }
    if m.GetDataSharingConsents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDataSharingConsents()))
        for i, v := range m.GetDataSharingConsents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dataSharingConsents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDepOnboardingSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDepOnboardingSettings()))
        for i, v := range m.GetDepOnboardingSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("depOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDerivedCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDerivedCredentials()))
        for i, v := range m.GetDerivedCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("derivedCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetectedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectedApps()))
        for i, v := range m.GetDetectedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCategories()))
        for i, v := range m.GetDeviceCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicies()))
        for i, v := range m.GetDeviceCompliancePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCompliancePolicyDeviceStateSummary", m.GetDeviceCompliancePolicyDeviceStateSummary())
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
    if m.GetDeviceComplianceScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceComplianceScripts()))
        for i, v := range m.GetDeviceComplianceScripts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceComplianceScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationConflictSummary() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationConflictSummary()))
        for i, v := range m.GetDeviceConfigurationConflictSummary() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationConflictSummary", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceConfigurationDeviceStateSummaries", m.GetDeviceConfigurationDeviceStateSummaries())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationRestrictedAppsViolations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationRestrictedAppsViolations()))
        for i, v := range m.GetDeviceConfigurationRestrictedAppsViolations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationRestrictedAppsViolations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurations()))
        for i, v := range m.GetDeviceConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationsAllManagedDeviceCertificateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationsAllManagedDeviceCertificateStates()))
        for i, v := range m.GetDeviceConfigurationsAllManagedDeviceCertificateStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationsAllManagedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceConfigurationUserStateSummaries", m.GetDeviceConfigurationUserStateSummaries())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCustomAttributeShellScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCustomAttributeShellScripts()))
        for i, v := range m.GetDeviceCustomAttributeShellScripts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCustomAttributeShellScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceEnrollmentConfigurations()))
        for i, v := range m.GetDeviceEnrollmentConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceEnrollmentConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceHealthScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceHealthScripts()))
        for i, v := range m.GetDeviceHealthScripts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceHealthScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementPartners()))
        for i, v := range m.GetDeviceManagementPartners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementScripts()))
        for i, v := range m.GetDeviceManagementScripts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementScripts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceProtectionOverview", m.GetDeviceProtectionOverview())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceShellScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceShellScripts()))
        for i, v := range m.GetDeviceShellScripts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceShellScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomainJoinConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainJoinConnectors()))
        for i, v := range m.GetDomainJoinConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("domainJoinConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmbeddedSIMActivationCodePools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmbeddedSIMActivationCodePools()))
        for i, v := range m.GetEmbeddedSIMActivationCodePools() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("embeddedSIMActivationCodePools", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExchangeConnectors()))
        for i, v := range m.GetExchangeConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exchangeConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeOnPremisesPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExchangeOnPremisesPolicies()))
        for i, v := range m.GetExchangeOnPremisesPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exchangeOnPremisesPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("exchangeOnPremisesPolicy", m.GetExchangeOnPremisesPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyCategories()))
        for i, v := range m.GetGroupPolicyCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyConfigurations()))
        for i, v := range m.GetGroupPolicyConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyDefinitionFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyDefinitionFiles()))
        for i, v := range m.GetGroupPolicyDefinitionFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyDefinitions()))
        for i, v := range m.GetGroupPolicyDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyMigrationReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyMigrationReports()))
        for i, v := range m.GetGroupPolicyMigrationReports() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyMigrationReports", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyObjectFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyObjectFiles()))
        for i, v := range m.GetGroupPolicyObjectFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyObjectFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyUploadedDefinitionFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyUploadedDefinitionFiles()))
        for i, v := range m.GetGroupPolicyUploadedDefinitionFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyUploadedDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedDeviceIdentities()))
        for i, v := range m.GetImportedDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("importedDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedWindowsAutopilotDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetImportedWindowsAutopilotDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("importedWindowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntents()))
        for i, v := range m.GetIntents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("intents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("intuneAccountId", m.GetIntuneAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("intuneBrand", m.GetIntuneBrand())
        if err != nil {
            return err
        }
    }
    if m.GetIntuneBrandingProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntuneBrandingProfiles()))
        for i, v := range m.GetIntuneBrandingProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("intuneBrandingProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosUpdateStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosUpdateStatuses()))
        for i, v := range m.GetIosUpdateStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("iosUpdateStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMacOSSoftwareUpdateAccountSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMacOSSoftwareUpdateAccountSummaries()))
        for i, v := range m.GetMacOSSoftwareUpdateAccountSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("macOSSoftwareUpdateAccountSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceCleanupSettings", m.GetManagedDeviceCleanupSettings())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceEncryptionStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceEncryptionStates()))
        for i, v := range m.GetManagedDeviceEncryptionStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceEncryptionStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceOverview", m.GetManagedDeviceOverview())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumDepTokens", m.GetMaximumDepTokens())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelConfigurations()))
        for i, v := range m.GetMicrosoftTunnelConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelHealthThresholds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelHealthThresholds()))
        for i, v := range m.GetMicrosoftTunnelHealthThresholds() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelHealthThresholds", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelServerLogCollectionResponses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelServerLogCollectionResponses()))
        for i, v := range m.GetMicrosoftTunnelServerLogCollectionResponses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelServerLogCollectionResponses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelSites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelSites()))
        for i, v := range m.GetMicrosoftTunnelSites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelSites", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppTroubleshootingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppTroubleshootingEvents()))
        for i, v := range m.GetMobileAppTroubleshootingEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileThreatDefenseConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileThreatDefenseConnectors()))
        for i, v := range m.GetMobileThreatDefenseConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileThreatDefenseConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNdesConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNdesConnectors()))
        for i, v := range m.GetNdesConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("ndesConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationMessageTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationMessageTemplates()))
        for i, v := range m.GetNotificationMessageTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("notificationMessageTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOemWarrantyInformationOnboarding() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOemWarrantyInformationOnboarding()))
        for i, v := range m.GetOemWarrantyInformationOnboarding() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("oemWarrantyInformationOnboarding", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoteActionAudits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteActionAudits()))
        for i, v := range m.GetRemoteActionAudits() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("remoteActionAudits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAssistancePartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteAssistancePartners()))
        for i, v := range m.GetRemoteAssistancePartners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("remoteAssistancePartners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("remoteAssistanceSettings", m.GetRemoteAssistanceSettings())
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
    if m.GetResourceAccessProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceAccessProfiles()))
        for i, v := range m.GetResourceAccessProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceAccessProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceOperations()))
        for i, v := range m.GetResourceOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceOperations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReusablePolicySettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReusablePolicySettings()))
        for i, v := range m.GetReusablePolicySettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reusablePolicySettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReusableSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReusableSettings()))
        for i, v := range m.GetReusableSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reusableSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleScopeTags()))
        for i, v := range m.GetRoleScopeTags() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleScopeTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettingDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareUpdateStatusSummary", m.GetSoftwareUpdateStatusSummary())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := (*m.GetSubscriptions()).String()
        err = writer.WriteStringValue("subscriptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptionState() != nil {
        cast := (*m.GetSubscriptionState()).String()
        err = writer.WriteStringValue("subscriptionState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTelecomExpenseManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTelecomExpenseManagementPartners()))
        for i, v := range m.GetTelecomExpenseManagementPartners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("telecomExpenseManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplates()))
        for i, v := range m.GetTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateSettings()))
        for i, v := range m.GetTemplateSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("templateSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tenantAttachRBAC", m.GetTenantAttachRBAC())
        if err != nil {
            return err
        }
    }
    if m.GetTermsAndConditions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTermsAndConditions()))
        for i, v := range m.GetTermsAndConditions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("termsAndConditions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTroubleshootingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTroubleshootingEvents()))
        for i, v := range m.GetTroubleshootingEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("troubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAnomaly() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAnomaly()))
        for i, v := range m.GetUserExperienceAnalyticsAnomaly() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAnomaly", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAnomalyDevice() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAnomalyDevice()))
        for i, v := range m.GetUserExperienceAnalyticsAnomalyDevice() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAnomalyDevice", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsAnomalySeverityOverview", m.GetUserExperienceAnalyticsAnomalySeverityOverview())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDeviceModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformanceDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthOSVersionPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsAppHealthOverview", m.GetUserExperienceAnalyticsAppHealthOverview())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBaselines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBaselines()))
        for i, v := range m.GetUserExperienceAnalyticsBaselines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBaselines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthAppImpact() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthAppImpact()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthAppImpact() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthAppImpact", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsBatteryHealthCapacityDetails", m.GetUserExperienceAnalyticsBatteryHealthCapacityDetails())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceAppImpact", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthModelPerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthOsPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthOsPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthOsPerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthOsPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsBatteryHealthRuntimeDetails", m.GetUserExperienceAnalyticsBatteryHealthRuntimeDetails())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsCategories()))
        for i, v := range m.GetUserExperienceAnalyticsCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceMetricHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceMetricHistory()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceMetricHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDevicePerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsDeviceScope", m.GetUserExperienceAnalyticsDeviceScope())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceScopes()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceScopes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceScores()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceScores() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupHistory()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcesses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcesses()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcesses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcesses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcessPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()))
        for i, v := range m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicesWithoutCloudIdentity", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceTimelineEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceTimelineEvents()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceTimelineEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceTimelineEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsImpactingProcess() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsImpactingProcess()))
        for i, v := range m.GetUserExperienceAnalyticsImpactingProcess() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsImpactingProcess", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsMetricHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsMetricHistory()))
        for i, v := range m.GetUserExperienceAnalyticsMetricHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsModelScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsModelScores()))
        for i, v := range m.GetUserExperienceAnalyticsModelScores() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsModelScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsNotAutopilotReadyDevice() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsNotAutopilotReadyDevice()))
        for i, v := range m.GetUserExperienceAnalyticsNotAutopilotReadyDevice() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsNotAutopilotReadyDevice", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsOverview", m.GetUserExperienceAnalyticsOverview())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsRemoteConnection() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsRemoteConnection()))
        for i, v := range m.GetUserExperienceAnalyticsRemoteConnection() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsRemoteConnection", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsResourcePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsResourcePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsResourcePerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsResourcePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsScoreHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsScoreHistory()))
        for i, v := range m.GetUserExperienceAnalyticsScoreHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsScoreHistory", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsSettings", m.GetUserExperienceAnalyticsSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric", m.GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics()))
        for i, v := range m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereMetrics", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserPfxCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserPfxCertificates()))
        for i, v := range m.GetUserPfxCertificates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userPfxCertificates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("virtualEndpoint", m.GetVirtualEndpoint())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsAutopilotDeploymentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsAutopilotDeploymentProfiles()))
        for i, v := range m.GetWindowsAutopilotDeploymentProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeploymentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsAutopilotDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetWindowsAutopilotDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsAutopilotSettings", m.GetWindowsAutopilotSettings())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDriverUpdateProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsDriverUpdateProfiles()))
        for i, v := range m.GetWindowsDriverUpdateProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsDriverUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsFeatureUpdateProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsFeatureUpdateProfiles()))
        for i, v := range m.GetWindowsFeatureUpdateProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsFeatureUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionAppLearningSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionAppLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionAppLearningSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionAppLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionNetworkLearningSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionNetworkLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionNetworkLearningSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionNetworkLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsMalwareInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsMalwareInformation()))
        for i, v := range m.GetWindowsMalwareInformation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsMalwareInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsMalwareOverview", m.GetWindowsMalwareOverview())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsQualityUpdateProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsQualityUpdateProfiles()))
        for i, v := range m.GetWindowsQualityUpdateProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsQualityUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsUpdateCatalogItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsUpdateCatalogItems()))
        for i, v := range m.GetWindowsUpdateCatalogItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsUpdateCatalogItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetZebraFotaArtifacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetZebraFotaArtifacts()))
        for i, v := range m.GetZebraFotaArtifacts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("zebraFotaArtifacts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("zebraFotaConnector", m.GetZebraFotaConnector())
        if err != nil {
            return err
        }
    }
    if m.GetZebraFotaDeployments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetZebraFotaDeployments()))
        for i, v := range m.GetZebraFotaDeployments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("zebraFotaDeployments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountMoveCompletionDateTime sets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
func (m *DeviceManagement) SetAccountMoveCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.accountMoveCompletionDateTime = value
}
// SetAdminConsent sets the adminConsent property value. Admin consent information.
func (m *DeviceManagement) SetAdminConsent(value AdminConsentable)() {
    m.adminConsent = value
}
// SetAdvancedThreatProtectionOnboardingStateSummary sets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
func (m *DeviceManagement) SetAdvancedThreatProtectionOnboardingStateSummary(value AdvancedThreatProtectionOnboardingStateSummaryable)() {
    m.advancedThreatProtectionOnboardingStateSummary = value
}
// SetAndroidDeviceOwnerEnrollmentProfiles sets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
func (m *DeviceManagement) SetAndroidDeviceOwnerEnrollmentProfiles(value []AndroidDeviceOwnerEnrollmentProfileable)() {
    m.androidDeviceOwnerEnrollmentProfiles = value
}
// SetAndroidForWorkAppConfigurationSchemas sets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
func (m *DeviceManagement) SetAndroidForWorkAppConfigurationSchemas(value []AndroidForWorkAppConfigurationSchemaable)() {
    m.androidForWorkAppConfigurationSchemas = value
}
// SetAndroidForWorkEnrollmentProfiles sets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
func (m *DeviceManagement) SetAndroidForWorkEnrollmentProfiles(value []AndroidForWorkEnrollmentProfileable)() {
    m.androidForWorkEnrollmentProfiles = value
}
// SetAndroidForWorkSettings sets the androidForWorkSettings property value. The singleton Android for Work settings entity.
func (m *DeviceManagement) SetAndroidForWorkSettings(value AndroidForWorkSettingsable)() {
    m.androidForWorkSettings = value
}
// SetAndroidManagedStoreAccountEnterpriseSettings sets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
func (m *DeviceManagement) SetAndroidManagedStoreAccountEnterpriseSettings(value AndroidManagedStoreAccountEnterpriseSettingsable)() {
    m.androidManagedStoreAccountEnterpriseSettings = value
}
// SetAndroidManagedStoreAppConfigurationSchemas sets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
func (m *DeviceManagement) SetAndroidManagedStoreAppConfigurationSchemas(value []AndroidManagedStoreAppConfigurationSchemaable)() {
    m.androidManagedStoreAppConfigurationSchemas = value
}
// SetApplePushNotificationCertificate sets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)() {
    m.applePushNotificationCertificate = value
}
// SetAppleUserInitiatedEnrollmentProfiles sets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
func (m *DeviceManagement) SetAppleUserInitiatedEnrollmentProfiles(value []AppleUserInitiatedEnrollmentProfileable)() {
    m.appleUserInitiatedEnrollmentProfiles = value
}
// SetAssignmentFilters sets the assignmentFilters property value. The list of assignment filters
func (m *DeviceManagement) SetAssignmentFilters(value []DeviceAndAppManagementAssignmentFilterable)() {
    m.assignmentFilters = value
}
// SetAuditEvents sets the auditEvents property value. The Audit Events
func (m *DeviceManagement) SetAuditEvents(value []AuditEventable)() {
    m.auditEvents = value
}
// SetAutopilotEvents sets the autopilotEvents property value. The list of autopilot events for the tenant.
func (m *DeviceManagement) SetAutopilotEvents(value []DeviceManagementAutopilotEventable)() {
    m.autopilotEvents = value
}
// SetCartToClassAssociations sets the cartToClassAssociations property value. The Cart To Class Associations.
func (m *DeviceManagement) SetCartToClassAssociations(value []CartToClassAssociationable)() {
    m.cartToClassAssociations = value
}
// SetCategories sets the categories property value. The available categories
func (m *DeviceManagement) SetCategories(value []DeviceManagementSettingCategoryable)() {
    m.categories = value
}
// SetCertificateConnectorDetails sets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagement) SetCertificateConnectorDetails(value []CertificateConnectorDetailsable)() {
    m.certificateConnectorDetails = value
}
// SetChromeOSOnboardingSettings sets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
func (m *DeviceManagement) SetChromeOSOnboardingSettings(value []ChromeOSOnboardingSettingsable)() {
    m.chromeOSOnboardingSettings = value
}
// SetCloudPCConnectivityIssues sets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
func (m *DeviceManagement) SetCloudPCConnectivityIssues(value []CloudPCConnectivityIssueable)() {
    m.cloudPCConnectivityIssues = value
}
// SetComanagedDevices sets the comanagedDevices property value. The list of co-managed devices report
func (m *DeviceManagement) SetComanagedDevices(value []ManagedDeviceable)() {
    m.comanagedDevices = value
}
// SetComanagementEligibleDevices sets the comanagementEligibleDevices property value. The list of co-management eligible devices report
func (m *DeviceManagement) SetComanagementEligibleDevices(value []ComanagementEligibleDeviceable)() {
    m.comanagementEligibleDevices = value
}
// SetComplianceCategories sets the complianceCategories property value. List of all compliance categories
func (m *DeviceManagement) SetComplianceCategories(value []DeviceManagementConfigurationCategoryable)() {
    m.complianceCategories = value
}
// SetComplianceManagementPartners sets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartnerable)() {
    m.complianceManagementPartners = value
}
// SetCompliancePolicies sets the compliancePolicies property value. List of all compliance policies
func (m *DeviceManagement) SetCompliancePolicies(value []DeviceManagementCompliancePolicyable)() {
    m.compliancePolicies = value
}
// SetComplianceSettings sets the complianceSettings property value. List of all ComplianceSettings
func (m *DeviceManagement) SetComplianceSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.complianceSettings = value
}
// SetConditionalAccessSettings sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)() {
    m.conditionalAccessSettings = value
}
// SetConfigManagerCollections sets the configManagerCollections property value. A list of ConfigManagerCollection
func (m *DeviceManagement) SetConfigManagerCollections(value []ConfigManagerCollectionable)() {
    m.configManagerCollections = value
}
// SetConfigurationCategories sets the configurationCategories property value. List of all Configuration Categories
func (m *DeviceManagement) SetConfigurationCategories(value []DeviceManagementConfigurationCategoryable)() {
    m.configurationCategories = value
}
// SetConfigurationPolicies sets the configurationPolicies property value. List of all Configuration policies
func (m *DeviceManagement) SetConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)() {
    m.configurationPolicies = value
}
// SetConfigurationPolicyTemplates sets the configurationPolicyTemplates property value. List of all templates
func (m *DeviceManagement) SetConfigurationPolicyTemplates(value []DeviceManagementConfigurationPolicyTemplateable)() {
    m.configurationPolicyTemplates = value
}
// SetConfigurationSettings sets the configurationSettings property value. List of all ConfigurationSettings
func (m *DeviceManagement) SetConfigurationSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.configurationSettings = value
}
// SetDataProcessorServiceForWindowsFeaturesOnboarding sets the dataProcessorServiceForWindowsFeaturesOnboarding property value. A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
func (m *DeviceManagement) SetDataProcessorServiceForWindowsFeaturesOnboarding(value DataProcessorServiceForWindowsFeaturesOnboardingable)() {
    m.dataProcessorServiceForWindowsFeaturesOnboarding = value
}
// SetDataSharingConsents sets the dataSharingConsents property value. Data sharing consents.
func (m *DeviceManagement) SetDataSharingConsents(value []DataSharingConsentable)() {
    m.dataSharingConsents = value
}
// SetDepOnboardingSettings sets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
func (m *DeviceManagement) SetDepOnboardingSettings(value []DepOnboardingSettingable)() {
    m.depOnboardingSettings = value
}
// SetDerivedCredentials sets the derivedCredentials property value. Collection of Derived credential settings associated with account.
func (m *DeviceManagement) SetDerivedCredentials(value []DeviceManagementDerivedCredentialSettingsable)() {
    m.derivedCredentials = value
}
// SetDetectedApps sets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) SetDetectedApps(value []DetectedAppable)() {
    m.detectedApps = value
}
// SetDeviceCategories sets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategoryable)() {
    m.deviceCategories = value
}
// SetDeviceCompliancePolicies sets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)() {
    m.deviceCompliancePolicies = value
}
// SetDeviceCompliancePolicyDeviceStateSummary sets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)() {
    m.deviceCompliancePolicyDeviceStateSummary = value
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// SetDeviceComplianceReportSummarizationDateTime sets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
func (m *DeviceManagement) SetDeviceComplianceReportSummarizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deviceComplianceReportSummarizationDateTime = value
}
// SetDeviceComplianceScripts sets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceComplianceScripts(value []DeviceComplianceScriptable)() {
    m.deviceComplianceScripts = value
}
// SetDeviceConfigurationConflictSummary sets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
func (m *DeviceManagement) SetDeviceConfigurationConflictSummary(value []DeviceConfigurationConflictSummaryable)() {
    m.deviceConfigurationConflictSummary = value
}
// SetDeviceConfigurationDeviceStateSummaries sets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)() {
    m.deviceConfigurationDeviceStateSummaries = value
}
// SetDeviceConfigurationRestrictedAppsViolations sets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
func (m *DeviceManagement) SetDeviceConfigurationRestrictedAppsViolations(value []RestrictedAppsViolationable)() {
    m.deviceConfigurationRestrictedAppsViolations = value
}
// SetDeviceConfigurations sets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfigurationable)() {
    m.deviceConfigurations = value
}
// SetDeviceConfigurationsAllManagedDeviceCertificateStates sets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
func (m *DeviceManagement) SetDeviceConfigurationsAllManagedDeviceCertificateStates(value []ManagedAllDeviceCertificateStateable)() {
    m.deviceConfigurationsAllManagedDeviceCertificateStates = value
}
// SetDeviceConfigurationUserStateSummaries sets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
func (m *DeviceManagement) SetDeviceConfigurationUserStateSummaries(value DeviceConfigurationUserStateSummaryable)() {
    m.deviceConfigurationUserStateSummaries = value
}
// SetDeviceCustomAttributeShellScripts sets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceCustomAttributeShellScripts(value []DeviceCustomAttributeShellScriptable)() {
    m.deviceCustomAttributeShellScripts = value
}
// SetDeviceEnrollmentConfigurations sets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)() {
    m.deviceEnrollmentConfigurations = value
}
// SetDeviceHealthScripts sets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceHealthScripts(value []DeviceHealthScriptable)() {
    m.deviceHealthScripts = value
}
// SetDeviceManagementPartners sets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartnerable)() {
    m.deviceManagementPartners = value
}
// SetDeviceManagementScripts sets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceManagementScripts(value []DeviceManagementScriptable)() {
    m.deviceManagementScripts = value
}
// SetDeviceProtectionOverview sets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) SetDeviceProtectionOverview(value DeviceProtectionOverviewable)() {
    m.deviceProtectionOverview = value
}
// SetDeviceShellScripts sets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceShellScripts(value []DeviceShellScriptable)() {
    m.deviceShellScripts = value
}
// SetDomainJoinConnectors sets the domainJoinConnectors property value. A list of connector objects.
func (m *DeviceManagement) SetDomainJoinConnectors(value []DeviceManagementDomainJoinConnectorable)() {
    m.domainJoinConnectors = value
}
// SetEmbeddedSIMActivationCodePools sets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
func (m *DeviceManagement) SetEmbeddedSIMActivationCodePools(value []EmbeddedSIMActivationCodePoolable)() {
    m.embeddedSIMActivationCodePools = value
}
// SetExchangeConnectors sets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)() {
    m.exchangeConnectors = value
}
// SetExchangeOnPremisesPolicies sets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagement) SetExchangeOnPremisesPolicies(value []DeviceManagementExchangeOnPremisesPolicyable)() {
    m.exchangeOnPremisesPolicies = value
}
// SetExchangeOnPremisesPolicy sets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
func (m *DeviceManagement) SetExchangeOnPremisesPolicy(value DeviceManagementExchangeOnPremisesPolicyable)() {
    m.exchangeOnPremisesPolicy = value
}
// SetGroupPolicyCategories sets the groupPolicyCategories property value. The available group policy categories for this account.
func (m *DeviceManagement) SetGroupPolicyCategories(value []GroupPolicyCategoryable)() {
    m.groupPolicyCategories = value
}
// SetGroupPolicyConfigurations sets the groupPolicyConfigurations property value. The group policy configurations created by this account.
func (m *DeviceManagement) SetGroupPolicyConfigurations(value []GroupPolicyConfigurationable)() {
    m.groupPolicyConfigurations = value
}
// SetGroupPolicyDefinitionFiles sets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
func (m *DeviceManagement) SetGroupPolicyDefinitionFiles(value []GroupPolicyDefinitionFileable)() {
    m.groupPolicyDefinitionFiles = value
}
// SetGroupPolicyDefinitions sets the groupPolicyDefinitions property value. The available group policy definitions for this account.
func (m *DeviceManagement) SetGroupPolicyDefinitions(value []GroupPolicyDefinitionable)() {
    m.groupPolicyDefinitions = value
}
// SetGroupPolicyMigrationReports sets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
func (m *DeviceManagement) SetGroupPolicyMigrationReports(value []GroupPolicyMigrationReportable)() {
    m.groupPolicyMigrationReports = value
}
// SetGroupPolicyObjectFiles sets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
func (m *DeviceManagement) SetGroupPolicyObjectFiles(value []GroupPolicyObjectFileable)() {
    m.groupPolicyObjectFiles = value
}
// SetGroupPolicyUploadedDefinitionFiles sets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
func (m *DeviceManagement) SetGroupPolicyUploadedDefinitionFiles(value []GroupPolicyUploadedDefinitionFileable)() {
    m.groupPolicyUploadedDefinitionFiles = value
}
// SetImportedDeviceIdentities sets the importedDeviceIdentities property value. The imported device identities.
func (m *DeviceManagement) SetImportedDeviceIdentities(value []ImportedDeviceIdentityable)() {
    m.importedDeviceIdentities = value
}
// SetImportedWindowsAutopilotDeviceIdentities sets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
// SetIntents sets the intents property value. The device management intents
func (m *DeviceManagement) SetIntents(value []DeviceManagementIntentable)() {
    m.intents = value
}
// SetIntuneAccountId sets the intuneAccountId property value. Intune Account ID for given tenant
func (m *DeviceManagement) SetIntuneAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.intuneAccountId = value
}
// SetIntuneBrand sets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) SetIntuneBrand(value IntuneBrandable)() {
    m.intuneBrand = value
}
// SetIntuneBrandingProfiles sets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
func (m *DeviceManagement) SetIntuneBrandingProfiles(value []IntuneBrandingProfileable)() {
    m.intuneBrandingProfiles = value
}
// SetIosUpdateStatuses sets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)() {
    m.iosUpdateStatuses = value
}
// SetLastReportAggregationDateTime sets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
func (m *DeviceManagement) SetLastReportAggregationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportAggregationDateTime = value
}
// SetLegacyPcManangementEnabled sets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
func (m *DeviceManagement) SetLegacyPcManangementEnabled(value *bool)() {
    m.legacyPcManangementEnabled = value
}
// SetMacOSSoftwareUpdateAccountSummaries sets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
func (m *DeviceManagement) SetMacOSSoftwareUpdateAccountSummaries(value []MacOSSoftwareUpdateAccountSummaryable)() {
    m.macOSSoftwareUpdateAccountSummaries = value
}
// SetManagedDeviceCleanupSettings sets the managedDeviceCleanupSettings property value. Device cleanup rule
func (m *DeviceManagement) SetManagedDeviceCleanupSettings(value ManagedDeviceCleanupSettingsable)() {
    m.managedDeviceCleanupSettings = value
}
// SetManagedDeviceEncryptionStates sets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
func (m *DeviceManagement) SetManagedDeviceEncryptionStates(value []ManagedDeviceEncryptionStateable)() {
    m.managedDeviceEncryptionStates = value
}
// SetManagedDeviceOverview sets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) SetManagedDeviceOverview(value ManagedDeviceOverviewable)() {
    m.managedDeviceOverview = value
}
// SetManagedDevices sets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) SetManagedDevices(value []ManagedDeviceable)() {
    m.managedDevices = value
}
// SetMaximumDepTokens sets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
func (m *DeviceManagement) SetMaximumDepTokens(value *int32)() {
    m.maximumDepTokens = value
}
// SetMicrosoftTunnelConfigurations sets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelConfigurations(value []MicrosoftTunnelConfigurationable)() {
    m.microsoftTunnelConfigurations = value
}
// SetMicrosoftTunnelHealthThresholds sets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelHealthThresholds(value []MicrosoftTunnelHealthThresholdable)() {
    m.microsoftTunnelHealthThresholds = value
}
// SetMicrosoftTunnelServerLogCollectionResponses sets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelServerLogCollectionResponses(value []MicrosoftTunnelServerLogCollectionResponseable)() {
    m.microsoftTunnelServerLogCollectionResponses = value
}
// SetMicrosoftTunnelSites sets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelSites(value []MicrosoftTunnelSiteable)() {
    m.microsoftTunnelSites = value
}
// SetMobileAppTroubleshootingEvents sets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)() {
    m.mobileAppTroubleshootingEvents = value
}
// SetMobileThreatDefenseConnectors sets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)() {
    m.mobileThreatDefenseConnectors = value
}
// SetNdesConnectors sets the ndesConnectors property value. The collection of Ndes connectors for this account.
func (m *DeviceManagement) SetNdesConnectors(value []NdesConnectorable)() {
    m.ndesConnectors = value
}
// SetNotificationMessageTemplates sets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplateable)() {
    m.notificationMessageTemplates = value
}
// SetOemWarrantyInformationOnboarding sets the oemWarrantyInformationOnboarding property value. List of OEM Warranty Statuses
func (m *DeviceManagement) SetOemWarrantyInformationOnboarding(value []OemWarrantyInformationOnboardingable)() {
    m.oemWarrantyInformationOnboarding = value
}
// SetRemoteActionAudits sets the remoteActionAudits property value. The list of device remote action audits with the tenant.
func (m *DeviceManagement) SetRemoteActionAudits(value []RemoteActionAuditable)() {
    m.remoteActionAudits = value
}
// SetRemoteAssistancePartners sets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)() {
    m.remoteAssistancePartners = value
}
// SetRemoteAssistanceSettings sets the remoteAssistanceSettings property value. The remote assistance settings singleton
func (m *DeviceManagement) SetRemoteAssistanceSettings(value RemoteAssistanceSettingsable)() {
    m.remoteAssistanceSettings = value
}
// SetReports sets the reports property value. Reports singleton
func (m *DeviceManagement) SetReports(value DeviceManagementReportsable)() {
    m.reports = value
}
// SetResourceAccessProfiles sets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
func (m *DeviceManagement) SetResourceAccessProfiles(value []DeviceManagementResourceAccessProfileBaseable)() {
    m.resourceAccessProfiles = value
}
// SetResourceOperations sets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperationable)() {
    m.resourceOperations = value
}
// SetReusablePolicySettings sets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
func (m *DeviceManagement) SetReusablePolicySettings(value []DeviceManagementReusablePolicySettingable)() {
    m.reusablePolicySettings = value
}
// SetReusableSettings sets the reusableSettings property value. List of all reusable settings
func (m *DeviceManagement) SetReusableSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.reusableSettings = value
}
// SetRoleAssignments sets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)() {
    m.roleAssignments = value
}
// SetRoleDefinitions sets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinitionable)() {
    m.roleDefinitions = value
}
// SetRoleScopeTags sets the roleScopeTags property value. The Role Scope Tags.
func (m *DeviceManagement) SetRoleScopeTags(value []RoleScopeTagable)() {
    m.roleScopeTags = value
}
// SetSettingDefinitions sets the settingDefinitions property value. The device management intent setting definitions
func (m *DeviceManagement) SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)() {
    m.settingDefinitions = value
}
// SetSettings sets the settings property value. Account level settings.
func (m *DeviceManagement) SetSettings(value DeviceManagementSettingsable)() {
    m.settings = value
}
// SetSoftwareUpdateStatusSummary sets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)() {
    m.softwareUpdateStatusSummary = value
}
// SetSubscriptions sets the subscriptions property value. Tenant mobile device management subscriptions.
func (m *DeviceManagement) SetSubscriptions(value *DeviceManagementSubscriptions)() {
    m.subscriptions = value
}
// SetSubscriptionState sets the subscriptionState property value. Tenant mobile device management subscription state.
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    m.subscriptionState = value
}
// SetTelecomExpenseManagementPartners sets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)() {
    m.telecomExpenseManagementPartners = value
}
// SetTemplates sets the templates property value. The available templates
func (m *DeviceManagement) SetTemplates(value []DeviceManagementTemplateable)() {
    m.templates = value
}
// SetTemplateSettings sets the templateSettings property value. List of all TemplateSettings
func (m *DeviceManagement) SetTemplateSettings(value []DeviceManagementConfigurationSettingTemplateable)() {
    m.templateSettings = value
}
// SetTenantAttachRBAC sets the tenantAttachRBAC property value. TenantAttach RBAC Enablement
func (m *DeviceManagement) SetTenantAttachRBAC(value TenantAttachRBACable)() {
    m.tenantAttachRBAC = value
}
// SetTermsAndConditions sets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditionsable)() {
    m.termsAndConditions = value
}
// SetTroubleshootingEvents sets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)() {
    m.troubleshootingEvents = value
}
// SetUnlicensedAdminstratorsEnabled sets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
func (m *DeviceManagement) SetUnlicensedAdminstratorsEnabled(value *bool)() {
    m.unlicensedAdminstratorsEnabled = value
}
// SetUserExperienceAnalyticsAnomaly sets the userExperienceAnalyticsAnomaly property value. The user experience analytics anomaly entity contains anomaly details.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomaly(value []UserExperienceAnalyticsAnomalyable)() {
    m.userExperienceAnalyticsAnomaly = value
}
// SetUserExperienceAnalyticsAnomalyDevice sets the userExperienceAnalyticsAnomalyDevice property value. The user experience analytics anomaly entity contains device details.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomalyDevice(value []UserExperienceAnalyticsAnomalyDeviceable)() {
    m.userExperienceAnalyticsAnomalyDevice = value
}
// SetUserExperienceAnalyticsAnomalySeverityOverview sets the userExperienceAnalyticsAnomalySeverityOverview property value. The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomalySeverityOverview(value UserExperienceAnalyticsAnomalySeverityOverviewable)() {
    m.userExperienceAnalyticsAnomalySeverityOverview = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformance sets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformanceable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformance = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion sets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion = value
}
// SetUserExperienceAnalyticsAppHealthDeviceModelPerformance sets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)() {
    m.userExperienceAnalyticsAppHealthDeviceModelPerformance = value
}
// SetUserExperienceAnalyticsAppHealthDevicePerformance sets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformanceable)() {
    m.userExperienceAnalyticsAppHealthDevicePerformance = value
}
// SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails sets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)() {
    m.userExperienceAnalyticsAppHealthDevicePerformanceDetails = value
}
// SetUserExperienceAnalyticsAppHealthOSVersionPerformance sets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformanceable)() {
    m.userExperienceAnalyticsAppHealthOSVersionPerformance = value
}
// SetUserExperienceAnalyticsAppHealthOverview sets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOverview(value UserExperienceAnalyticsCategoryable)() {
    m.userExperienceAnalyticsAppHealthOverview = value
}
// SetUserExperienceAnalyticsBaselines sets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaselineable)() {
    m.userExperienceAnalyticsBaselines = value
}
// SetUserExperienceAnalyticsBatteryHealthAppImpact sets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthAppImpact(value []UserExperienceAnalyticsBatteryHealthAppImpactable)() {
    m.userExperienceAnalyticsBatteryHealthAppImpact = value
}
// SetUserExperienceAnalyticsBatteryHealthCapacityDetails sets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthCapacityDetails(value UserExperienceAnalyticsBatteryHealthCapacityDetailsable)() {
    m.userExperienceAnalyticsBatteryHealthCapacityDetails = value
}
// SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact sets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(value []UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)() {
    m.userExperienceAnalyticsBatteryHealthDeviceAppImpact = value
}
// SetUserExperienceAnalyticsBatteryHealthDevicePerformance sets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDevicePerformance(value []UserExperienceAnalyticsBatteryHealthDevicePerformanceable)() {
    m.userExperienceAnalyticsBatteryHealthDevicePerformance = value
}
// SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory sets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(value []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)() {
    m.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory = value
}
// SetUserExperienceAnalyticsBatteryHealthModelPerformance sets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthModelPerformance(value []UserExperienceAnalyticsBatteryHealthModelPerformanceable)() {
    m.userExperienceAnalyticsBatteryHealthModelPerformance = value
}
// SetUserExperienceAnalyticsBatteryHealthOsPerformance sets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthOsPerformance(value []UserExperienceAnalyticsBatteryHealthOsPerformanceable)() {
    m.userExperienceAnalyticsBatteryHealthOsPerformance = value
}
// SetUserExperienceAnalyticsBatteryHealthRuntimeDetails sets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(value UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)() {
    m.userExperienceAnalyticsBatteryHealthRuntimeDetails = value
}
// SetUserExperienceAnalyticsCategories sets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategoryable)() {
    m.userExperienceAnalyticsCategories = value
}
// SetUserExperienceAnalyticsDeviceMetricHistory sets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)() {
    m.userExperienceAnalyticsDeviceMetricHistory = value
}
// SetUserExperienceAnalyticsDevicePerformance sets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformanceable)() {
    m.userExperienceAnalyticsDevicePerformance = value
}
// SetUserExperienceAnalyticsDeviceScope sets the userExperienceAnalyticsDeviceScope property value. The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScope(value UserExperienceAnalyticsDeviceScopeable)() {
    m.userExperienceAnalyticsDeviceScope = value
}
// SetUserExperienceAnalyticsDeviceScopes sets the userExperienceAnalyticsDeviceScopes property value. The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScopes(value []UserExperienceAnalyticsDeviceScopeable)() {
    m.userExperienceAnalyticsDeviceScopes = value
}
// SetUserExperienceAnalyticsDeviceScores sets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScoresable)() {
    m.userExperienceAnalyticsDeviceScores = value
}
// SetUserExperienceAnalyticsDeviceStartupHistory sets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistoryable)() {
    m.userExperienceAnalyticsDeviceStartupHistory = value
}
// SetUserExperienceAnalyticsDeviceStartupProcesses sets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcessable)() {
    m.userExperienceAnalyticsDeviceStartupProcesses = value
}
// SetUserExperienceAnalyticsDeviceStartupProcessPerformance sets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformanceable)() {
    m.userExperienceAnalyticsDeviceStartupProcessPerformance = value
}
// SetUserExperienceAnalyticsDevicesWithoutCloudIdentity sets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(value []UserExperienceAnalyticsDeviceWithoutCloudIdentityable)() {
    m.userExperienceAnalyticsDevicesWithoutCloudIdentity = value
}
// SetUserExperienceAnalyticsDeviceTimelineEvents sets the userExperienceAnalyticsDeviceTimelineEvents property value. The user experience analytics device events entity contains NRT device timeline events details.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceTimelineEvents(value []UserExperienceAnalyticsDeviceTimelineEventsable)() {
    m.userExperienceAnalyticsDeviceTimelineEvents = value
}
// SetUserExperienceAnalyticsImpactingProcess sets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
func (m *DeviceManagement) SetUserExperienceAnalyticsImpactingProcess(value []UserExperienceAnalyticsImpactingProcessable)() {
    m.userExperienceAnalyticsImpactingProcess = value
}
// SetUserExperienceAnalyticsMetricHistory sets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)() {
    m.userExperienceAnalyticsMetricHistory = value
}
// SetUserExperienceAnalyticsModelScores sets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScoresable)() {
    m.userExperienceAnalyticsModelScores = value
}
// SetUserExperienceAnalyticsNotAutopilotReadyDevice sets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
func (m *DeviceManagement) SetUserExperienceAnalyticsNotAutopilotReadyDevice(value []UserExperienceAnalyticsNotAutopilotReadyDeviceable)() {
    m.userExperienceAnalyticsNotAutopilotReadyDevice = value
}
// SetUserExperienceAnalyticsOverview sets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) SetUserExperienceAnalyticsOverview(value UserExperienceAnalyticsOverviewable)() {
    m.userExperienceAnalyticsOverview = value
}
// SetUserExperienceAnalyticsRemoteConnection sets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
func (m *DeviceManagement) SetUserExperienceAnalyticsRemoteConnection(value []UserExperienceAnalyticsRemoteConnectionable)() {
    m.userExperienceAnalyticsRemoteConnection = value
}
// SetUserExperienceAnalyticsResourcePerformance sets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
func (m *DeviceManagement) SetUserExperienceAnalyticsResourcePerformance(value []UserExperienceAnalyticsResourcePerformanceable)() {
    m.userExperienceAnalyticsResourcePerformance = value
}
// SetUserExperienceAnalyticsScoreHistory sets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistoryable)() {
    m.userExperienceAnalyticsScoreHistory = value
}
// SetUserExperienceAnalyticsSettings sets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) SetUserExperienceAnalyticsSettings(value UserExperienceAnalyticsSettingsable)() {
    m.userExperienceAnalyticsSettings = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric sets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)() {
    m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereMetrics sets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetricable)() {
    m.userExperienceAnalyticsWorkFromAnywhereMetrics = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance sets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(value []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)() {
    m.userExperienceAnalyticsWorkFromAnywhereModelPerformance = value
}
// SetUserPfxCertificates sets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
func (m *DeviceManagement) SetUserPfxCertificates(value []UserPFXCertificateable)() {
    m.userPfxCertificates = value
}
// SetVirtualEndpoint sets the virtualEndpoint property value. The virtualEndpoint property
func (m *DeviceManagement) SetVirtualEndpoint(value VirtualEndpointable)() {
    m.virtualEndpoint = value
}
// SetWindowsAutopilotDeploymentProfiles sets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
func (m *DeviceManagement) SetWindowsAutopilotDeploymentProfiles(value []WindowsAutopilotDeploymentProfileable)() {
    m.windowsAutopilotDeploymentProfiles = value
}
// SetWindowsAutopilotDeviceIdentities sets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)() {
    m.windowsAutopilotDeviceIdentities = value
}
// SetWindowsAutopilotSettings sets the windowsAutopilotSettings property value. The Windows autopilot account settings.
func (m *DeviceManagement) SetWindowsAutopilotSettings(value WindowsAutopilotSettingsable)() {
    m.windowsAutopilotSettings = value
}
// SetWindowsDriverUpdateProfiles sets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
func (m *DeviceManagement) SetWindowsDriverUpdateProfiles(value []WindowsDriverUpdateProfileable)() {
    m.windowsDriverUpdateProfiles = value
}
// SetWindowsFeatureUpdateProfiles sets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
func (m *DeviceManagement) SetWindowsFeatureUpdateProfiles(value []WindowsFeatureUpdateProfileable)() {
    m.windowsFeatureUpdateProfiles = value
}
// SetWindowsInformationProtectionAppLearningSummaries sets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)() {
    m.windowsInformationProtectionAppLearningSummaries = value
}
// SetWindowsInformationProtectionNetworkLearningSummaries sets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)() {
    m.windowsInformationProtectionNetworkLearningSummaries = value
}
// SetWindowsMalwareInformation sets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) SetWindowsMalwareInformation(value []WindowsMalwareInformationable)() {
    m.windowsMalwareInformation = value
}
// SetWindowsMalwareOverview sets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) SetWindowsMalwareOverview(value WindowsMalwareOverviewable)() {
    m.windowsMalwareOverview = value
}
// SetWindowsQualityUpdateProfiles sets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
func (m *DeviceManagement) SetWindowsQualityUpdateProfiles(value []WindowsQualityUpdateProfileable)() {
    m.windowsQualityUpdateProfiles = value
}
// SetWindowsUpdateCatalogItems sets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
func (m *DeviceManagement) SetWindowsUpdateCatalogItems(value []WindowsUpdateCatalogItemable)() {
    m.windowsUpdateCatalogItems = value
}
// SetZebraFotaArtifacts sets the zebraFotaArtifacts property value. The Collection of ZebraFotaArtifacts.
func (m *DeviceManagement) SetZebraFotaArtifacts(value []ZebraFotaArtifactable)() {
    m.zebraFotaArtifacts = value
}
// SetZebraFotaConnector sets the zebraFotaConnector property value. The singleton ZebraFotaConnector associated with account.
func (m *DeviceManagement) SetZebraFotaConnector(value ZebraFotaConnectorable)() {
    m.zebraFotaConnector = value
}
// SetZebraFotaDeployments sets the zebraFotaDeployments property value. Collection of ZebraFotaDeployments associated with account.
func (m *DeviceManagement) SetZebraFotaDeployments(value []ZebraFotaDeploymentable)() {
    m.zebraFotaDeployments = value
}
