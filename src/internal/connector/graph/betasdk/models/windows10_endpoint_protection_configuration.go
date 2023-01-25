package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EndpointProtectionConfiguration 
type Windows10EndpointProtectionConfiguration struct {
    DeviceConfiguration
    // Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
    applicationGuardAllowCameraMicrophoneRedirection *bool
    // Allow users to download files from Edge in the application guard container and save them on the host file system
    applicationGuardAllowFileSaveOnHost *bool
    // Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
    applicationGuardAllowPersistence *bool
    // Allow printing to Local Printers from Container
    applicationGuardAllowPrintToLocalPrinters *bool
    // Allow printing to Network Printers from Container
    applicationGuardAllowPrintToNetworkPrinters *bool
    // Allow printing to PDF from Container
    applicationGuardAllowPrintToPDF *bool
    // Allow printing to XPS from Container
    applicationGuardAllowPrintToXPS *bool
    // Allow application guard to use virtual GPU
    applicationGuardAllowVirtualGPU *bool
    // Possible values for applicationGuardBlockClipboardSharingType
    applicationGuardBlockClipboardSharing *ApplicationGuardBlockClipboardSharingType
    // Possible values for applicationGuardBlockFileTransfer
    applicationGuardBlockFileTransfer *ApplicationGuardBlockFileTransferType
    // Block enterprise sites to load non-enterprise content, such as third party plug-ins
    applicationGuardBlockNonEnterpriseContent *bool
    // Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
    applicationGuardCertificateThumbprints []string
    // Enable Windows Defender Application Guard
    applicationGuardEnabled *bool
    // Possible values for ApplicationGuardEnabledOptions
    applicationGuardEnabledOptions *ApplicationGuardEnabledOptions
    // Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
    applicationGuardForceAuditing *bool
    // Possible values of AppLocker Application Control Types
    appLockerApplicationControl *AppLockerApplicationControlType
    // Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
    bitLockerAllowStandardUserEncryption *bool
    // Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
    bitLockerDisableWarningForOtherDiskEncryption *bool
    // Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
    bitLockerEnableStorageCardEncryptionOnMobile *bool
    // Allows the admin to require encryption to be turned on using BitLocker.
    bitLockerEncryptDevice *bool
    // BitLocker Fixed Drive Policy.
    bitLockerFixedDrivePolicy BitLockerFixedDrivePolicyable
    // BitLocker recovery password rotation type
    bitLockerRecoveryPasswordRotation *BitLockerRecoveryPasswordRotationType
    // BitLocker Removable Drive Policy.
    bitLockerRemovableDrivePolicy BitLockerRemovableDrivePolicyable
    // BitLocker System Drive Policy.
    bitLockerSystemDrivePolicy BitLockerSystemDrivePolicyable
    // List of folder paths to be added to the list of protected folders
    defenderAdditionalGuardedFolders []string
    // Possible values of Defender PUA Protection
    defenderAdobeReaderLaunchChildProcess *DefenderProtectionType
    // Possible values of Defender PUA Protection
    defenderAdvancedRansomewareProtectionType *DefenderProtectionType
    // Allows or disallows Windows Defender Behavior Monitoring functionality.
    defenderAllowBehaviorMonitoring *bool
    // To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
    defenderAllowCloudProtection *bool
    // Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
    defenderAllowEndUserAccess *bool
    // Allows or disallows Windows Defender Intrusion Prevention functionality.
    defenderAllowIntrusionPreventionSystem *bool
    // Allows or disallows Windows Defender On Access Protection functionality.
    defenderAllowOnAccessProtection *bool
    // Allows or disallows Windows Defender Realtime Monitoring functionality.
    defenderAllowRealTimeMonitoring *bool
    // Allows or disallows scanning of archives.
    defenderAllowScanArchiveFiles *bool
    // Allows or disallows Windows Defender IOAVP Protection functionality.
    defenderAllowScanDownloads *bool
    // Allows or disallows a scanning of network files.
    defenderAllowScanNetworkFiles *bool
    // Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
    defenderAllowScanRemovableDrivesDuringFullScan *bool
    // Allows or disallows Windows Defender Script Scanning functionality.
    defenderAllowScanScriptsLoadedInInternetExplorer *bool
    // List of exe files and folders to be excluded from attack surface reduction rules
    defenderAttackSurfaceReductionExcludedPaths []string
    // Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
    defenderBlockEndUserAccess *bool
    // Possible values of Defender Attack Surface Reduction Rules
    defenderBlockPersistenceThroughWmiType *DefenderAttackSurfaceType
    // This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
    defenderCheckForSignaturesBeforeRunningScan *bool
    // Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
    defenderCloudBlockLevel *DefenderCloudBlockLevelType
    // Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
    defenderCloudExtendedTimeoutInSeconds *int32
    // Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
    defenderDaysBeforeDeletingQuarantinedMalware *int32
    // Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
    defenderDetectedMalwareActions DefenderDetectedMalwareActionsable
    // Allows or disallows Windows Defender Behavior Monitoring functionality.
    defenderDisableBehaviorMonitoring *bool
    // This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
    defenderDisableCatchupFullScan *bool
    // This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
    defenderDisableCatchupQuickScan *bool
    // To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
    defenderDisableCloudProtection *bool
    // Allows or disallows Windows Defender Intrusion Prevention functionality.
    defenderDisableIntrusionPreventionSystem *bool
    // Allows or disallows Windows Defender On Access Protection functionality.
    defenderDisableOnAccessProtection *bool
    // Allows or disallows Windows Defender Realtime Monitoring functionality.
    defenderDisableRealTimeMonitoring *bool
    // Allows or disallows scanning of archives.
    defenderDisableScanArchiveFiles *bool
    // Allows or disallows Windows Defender IOAVP Protection functionality.
    defenderDisableScanDownloads *bool
    // Allows or disallows a scanning of network files.
    defenderDisableScanNetworkFiles *bool
    // Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
    defenderDisableScanRemovableDrivesDuringFullScan *bool
    // Allows or disallows Windows Defender Script Scanning functionality.
    defenderDisableScanScriptsLoadedInInternetExplorer *bool
    // Possible values of Defender PUA Protection
    defenderEmailContentExecution *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderEmailContentExecutionType *DefenderAttackSurfaceType
    // This policy setting allows you to enable or disable low CPU priority for scheduled scans.
    defenderEnableLowCpuPriority *bool
    // Allows or disallows scanning of email.
    defenderEnableScanIncomingMail *bool
    // Allows or disallows a full scan of mapped network drives.
    defenderEnableScanMappedNetworkDrivesDuringFullScan *bool
    // Xml content containing information regarding exploit protection details.
    defenderExploitProtectionXml []byte
    // Name of the file from which DefenderExploitProtectionXml was obtained.
    defenderExploitProtectionXmlFileName *string
    // File extensions to exclude from scans and real time protection.
    defenderFileExtensionsToExclude []string
    // Files and folder to exclude from scans and real time protection.
    defenderFilesAndFoldersToExclude []string
    // List of paths to exe that are allowed to access protected folders
    defenderGuardedFoldersAllowedAppPaths []string
    // Possible values of Folder Protection
    defenderGuardMyFoldersType *FolderProtectionType
    // Possible values of Defender PUA Protection
    defenderNetworkProtectionType *DefenderProtectionType
    // Possible values of Defender PUA Protection
    defenderOfficeAppsExecutableContentCreationOrLaunch *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderOfficeAppsExecutableContentCreationOrLaunchType *DefenderAttackSurfaceType
    // Possible values of Defender PUA Protection
    defenderOfficeAppsLaunchChildProcess *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderOfficeAppsLaunchChildProcessType *DefenderAttackSurfaceType
    // Possible values of Defender PUA Protection
    defenderOfficeAppsOtherProcessInjection *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderOfficeAppsOtherProcessInjectionType *DefenderAttackSurfaceType
    // Possible values of Defender PUA Protection
    defenderOfficeCommunicationAppsLaunchChildProcess *DefenderProtectionType
    // Possible values of Defender PUA Protection
    defenderOfficeMacroCodeAllowWin32Imports *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderOfficeMacroCodeAllowWin32ImportsType *DefenderAttackSurfaceType
    // Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderPotentiallyUnwantedAppAction *DefenderProtectionType
    // Possible values of Defender PUA Protection
    defenderPreventCredentialStealingType *DefenderProtectionType
    // Possible values of Defender PUA Protection
    defenderProcessCreation *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderProcessCreationType *DefenderAttackSurfaceType
    // Processes to exclude from scans and real time protection.
    defenderProcessesToExclude []string
    // Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
    defenderScanDirection *DefenderRealtimeScanDirection
    // Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
    defenderScanMaxCpuPercentage *int32
    // Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
    defenderScanType *DefenderScanType
    // Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
    defenderScheduledQuickScanTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
    defenderScheduledScanDay *WeeklySchedule
    // Selects the time of day that the Windows Defender scan should run.
    defenderScheduledScanTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Possible values of Defender PUA Protection
    defenderScriptDownloadedPayloadExecution *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderScriptDownloadedPayloadExecutionType *DefenderAttackSurfaceType
    // Possible values of Defender PUA Protection
    defenderScriptObfuscatedMacroCode *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderScriptObfuscatedMacroCodeType *DefenderAttackSurfaceType
    // Indicates whether or not to block user from overriding Exploit Protection settings.
    defenderSecurityCenterBlockExploitProtectionOverride *bool
    // Used to disable the display of the account protection area.
    defenderSecurityCenterDisableAccountUI *bool
    // Used to disable the display of the app and browser protection area.
    defenderSecurityCenterDisableAppBrowserUI *bool
    // Used to disable the display of the Clear TPM button.
    defenderSecurityCenterDisableClearTpmUI *bool
    // Used to disable the display of the family options area.
    defenderSecurityCenterDisableFamilyUI *bool
    // Used to disable the display of the hardware protection area.
    defenderSecurityCenterDisableHardwareUI *bool
    // Used to disable the display of the device performance and health area.
    defenderSecurityCenterDisableHealthUI *bool
    // Used to disable the display of the firewall and network protection area.
    defenderSecurityCenterDisableNetworkUI *bool
    // Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
    defenderSecurityCenterDisableNotificationAreaUI *bool
    // Used to disable the display of the ransomware protection area.
    defenderSecurityCenterDisableRansomwareUI *bool
    // Used to disable the display of the secure boot area under Device security.
    defenderSecurityCenterDisableSecureBootUI *bool
    // Used to disable the display of the security process troubleshooting under Device security.
    defenderSecurityCenterDisableTroubleshootingUI *bool
    // Used to disable the display of the virus and threat protection area.
    defenderSecurityCenterDisableVirusUI *bool
    // Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
    defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI *bool
    // The email address that is displayed to users.
    defenderSecurityCenterHelpEmail *string
    // The phone number or Skype ID that is displayed to users.
    defenderSecurityCenterHelpPhone *string
    // The help portal URL this is displayed to users.
    defenderSecurityCenterHelpURL *string
    // Possible values for defenderSecurityCenterITContactDisplay
    defenderSecurityCenterITContactDisplay *DefenderSecurityCenterITContactDisplayType
    // Possible values for defenderSecurityCenterNotificationsFromApp
    defenderSecurityCenterNotificationsFromApp *DefenderSecurityCenterNotificationsFromAppType
    // The company name that is displayed to the users.
    defenderSecurityCenterOrganizationDisplayName *string
    // Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
    defenderSignatureUpdateIntervalInHours *int32
    // Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
    defenderSubmitSamplesConsentType *DefenderSubmitSamplesConsentType
    // Possible values of Defender PUA Protection
    defenderUntrustedExecutable *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderUntrustedExecutableType *DefenderAttackSurfaceType
    // Possible values of Defender PUA Protection
    defenderUntrustedUSBProcess *DefenderProtectionType
    // Possible values of Defender Attack Surface Reduction Rules
    defenderUntrustedUSBProcessType *DefenderAttackSurfaceType
    // This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
    deviceGuardEnableSecureBootWithDMA *bool
    // Turns On Virtualization Based Security(VBS).
    deviceGuardEnableVirtualizationBasedSecurity *bool
    // Possible values of a property
    deviceGuardLaunchSystemGuard *Enablement
    // Possible values of Credential Guard settings.
    deviceGuardLocalSystemAuthorityCredentialGuardSettings *DeviceGuardLocalSystemAuthorityCredentialGuardType
    // Possible values of Secure Boot with DMA
    deviceGuardSecureBootWithDMA *SecureBootWithDMAType
    // Possible values of the DmaGuardDeviceEnumerationPolicy.
    dmaGuardDeviceEnumerationPolicy *DmaGuardDeviceEnumerationPolicyType
    // Blocks stateful FTP connections to the device
    firewallBlockStatefulFTP *bool
    // Possible values for firewallCertificateRevocationListCheckMethod
    firewallCertificateRevocationListCheckMethod *FirewallCertificateRevocationListCheckMethodType
    // Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
    firewallIdleTimeoutForSecurityAssociationInSeconds *int32
    // Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
    firewallIPSecExemptionsAllowDHCP *bool
    // Configures IPSec exemptions to allow ICMP
    firewallIPSecExemptionsAllowICMP *bool
    // Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
    firewallIPSecExemptionsAllowNeighborDiscovery *bool
    // Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
    firewallIPSecExemptionsAllowRouterDiscovery *bool
    // Configures IPSec exemptions to no exemptions
    firewallIPSecExemptionsNone *bool
    // If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
    firewallMergeKeyingModuleSettings *bool
    // Possible values for firewallPacketQueueingMethod
    firewallPacketQueueingMethod *FirewallPacketQueueingMethodType
    // Possible values for firewallPreSharedKeyEncodingMethod
    firewallPreSharedKeyEncodingMethod *FirewallPreSharedKeyEncodingMethodType
    // Configures the firewall profile settings for domain networks
    firewallProfileDomain WindowsFirewallNetworkProfileable
    // Configures the firewall profile settings for private networks
    firewallProfilePrivate WindowsFirewallNetworkProfileable
    // Configures the firewall profile settings for public networks
    firewallProfilePublic WindowsFirewallNetworkProfileable
    // Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
    firewallRules []WindowsFirewallRuleable
    // Possible values for LanManagerAuthenticationLevel
    lanManagerAuthenticationLevel *LanManagerAuthenticationLevel
    // If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
    lanManagerWorkstationDisableInsecureGuestLogons *bool
    // Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
    localSecurityOptionsAdministratorAccountName *string
    // Possible values for LocalSecurityOptionsAdministratorElevationPromptBehavior
    localSecurityOptionsAdministratorElevationPromptBehavior *LocalSecurityOptionsAdministratorElevationPromptBehaviorType
    // This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
    localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares *bool
    // Block PKU2U authentication requests to this device to use online identities.
    localSecurityOptionsAllowPKU2UAuthenticationRequests *bool
    // Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
    localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager *string
    // UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
    localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool *bool
    // This security setting determines whether a computer can be shut down without having to log on to Windows.
    localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn *bool
    // Allow UIAccess apps to prompt for elevation without using the secure desktop.
    localSecurityOptionsAllowUIAccessApplicationElevation *bool
    // Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
    localSecurityOptionsAllowUIAccessApplicationsForSecureLocations *bool
    // Prevent a portable computer from being undocked without having to log in.
    localSecurityOptionsAllowUndockWithoutHavingToLogon *bool
    // Prevent users from adding new Microsoft accounts to this computer.
    localSecurityOptionsBlockMicrosoftAccounts *bool
    // Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
    localSecurityOptionsBlockRemoteLogonWithBlankPassword *bool
    // Enabling this settings allows only interactively logged on user to access CD-ROM media.
    localSecurityOptionsBlockRemoteOpticalDriveAccess *bool
    // Restrict installing printer drivers as part of connecting to a shared printer to admins only.
    localSecurityOptionsBlockUsersInstallingPrinterDrivers *bool
    // This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
    localSecurityOptionsClearVirtualMemoryPageFile *bool
    // This security setting determines whether packet signing is required by the SMB client component.
    localSecurityOptionsClientDigitallySignCommunicationsAlways *bool
    // If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
    localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers *bool
    // App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
    localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation *bool
    // Determines whether the Local Administrator account is enabled or disabled.
    localSecurityOptionsDisableAdministratorAccount *bool
    // This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
    localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees *bool
    // Determines if the Guest account is enabled or disabled.
    localSecurityOptionsDisableGuestAccount *bool
    // This security setting determines whether packet signing is required by the SMB server component.
    localSecurityOptionsDisableServerDigitallySignCommunicationsAlways *bool
    // This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
    localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees *bool
    // This security setting determines what additional permissions will be granted for anonymous connections to the computer.
    localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts *bool
    // Require CTRL+ALT+DEL to be pressed before a user can log on.
    localSecurityOptionsDoNotRequireCtrlAltDel *bool
    // This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
    localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange *bool
    // Possible values for LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
    localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType
    // Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
    localSecurityOptionsGuestAccountName *string
    // Do not display the username of the last person who signed in on this device.
    localSecurityOptionsHideLastSignedInUser *bool
    // Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
    localSecurityOptionsHideUsernameAtSignIn *bool
    // Possible values for LocalSecurityOptionsInformationDisplayedOnLockScreen
    localSecurityOptionsInformationDisplayedOnLockScreen *LocalSecurityOptionsInformationDisplayedOnLockScreenType
    // Possible values for LocalSecurityOptionsInformationShownOnLockScreenType
    localSecurityOptionsInformationShownOnLockScreen *LocalSecurityOptionsInformationShownOnLockScreenType
    // Set message text for users attempting to log in.
    localSecurityOptionsLogOnMessageText *string
    // Set message title for users attempting to log in.
    localSecurityOptionsLogOnMessageTitle *string
    // Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
    localSecurityOptionsMachineInactivityLimit *int32
    // Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
    localSecurityOptionsMachineInactivityLimitInMinutes *int32
    // Possible values for LocalSecurityOptionsMinimumSessionSecurity
    localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients *LocalSecurityOptionsMinimumSessionSecurity
    // Possible values for LocalSecurityOptionsMinimumSessionSecurity
    localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers *LocalSecurityOptionsMinimumSessionSecurity
    // Enforce PKI certification path validation for a given executable file before it is permitted to run.
    localSecurityOptionsOnlyElevateSignedExecutables *bool
    // By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
    localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares *bool
    // Possible values for LocalSecurityOptionsSmartCardRemovalBehaviorType
    localSecurityOptionsSmartCardRemovalBehavior *LocalSecurityOptionsSmartCardRemovalBehaviorType
    // Possible values for LocalSecurityOptionsStandardUserElevationPromptBehavior
    localSecurityOptionsStandardUserElevationPromptBehavior *LocalSecurityOptionsStandardUserElevationPromptBehaviorType
    // Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
    localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation *bool
    // Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
    localSecurityOptionsUseAdminApprovalMode *bool
    // Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
    localSecurityOptionsUseAdminApprovalModeForAdministrators *bool
    // Virtualize file and registry write failures to per user locations
    localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations *bool
    // Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
    smartScreenBlockOverrideForFiles *bool
    // Allows IT Admins to configure SmartScreen for Windows.
    smartScreenEnableInShell *bool
    // This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
    userRightsAccessCredentialManagerAsTrustedCaller DeviceManagementUserRightsSettingable
    // This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
    userRightsActAsPartOfTheOperatingSystem DeviceManagementUserRightsSettingable
    // This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
    userRightsAllowAccessFromNetwork DeviceManagementUserRightsSettingable
    // This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
    userRightsBackupData DeviceManagementUserRightsSettingable
    // This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
    userRightsBlockAccessFromNetwork DeviceManagementUserRightsSettingable
    // This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
    userRightsChangeSystemTime DeviceManagementUserRightsSettingable
    // This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
    userRightsCreateGlobalObjects DeviceManagementUserRightsSettingable
    // This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
    userRightsCreatePageFile DeviceManagementUserRightsSettingable
    // This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
    userRightsCreatePermanentSharedObjects DeviceManagementUserRightsSettingable
    // This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
    userRightsCreateSymbolicLinks DeviceManagementUserRightsSettingable
    // This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
    userRightsCreateToken DeviceManagementUserRightsSettingable
    // This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
    userRightsDebugPrograms DeviceManagementUserRightsSettingable
    // This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
    userRightsDelegation DeviceManagementUserRightsSettingable
    // This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
    userRightsDenyLocalLogOn DeviceManagementUserRightsSettingable
    // This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
    userRightsGenerateSecurityAudits DeviceManagementUserRightsSettingable
    // Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
    userRightsImpersonateClient DeviceManagementUserRightsSettingable
    // This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
    userRightsIncreaseSchedulingPriority DeviceManagementUserRightsSettingable
    // This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
    userRightsLoadUnloadDrivers DeviceManagementUserRightsSettingable
    // This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
    userRightsLocalLogOn DeviceManagementUserRightsSettingable
    // This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
    userRightsLockMemory DeviceManagementUserRightsSettingable
    // This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
    userRightsManageAuditingAndSecurityLogs DeviceManagementUserRightsSettingable
    // This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
    userRightsManageVolumes DeviceManagementUserRightsSettingable
    // This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
    userRightsModifyFirmwareEnvironment DeviceManagementUserRightsSettingable
    // This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
    userRightsModifyObjectLabels DeviceManagementUserRightsSettingable
    // This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
    userRightsProfileSingleProcess DeviceManagementUserRightsSettingable
    // This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
    userRightsRemoteDesktopServicesLogOn DeviceManagementUserRightsSettingable
    // This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
    userRightsRemoteShutdown DeviceManagementUserRightsSettingable
    // This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
    userRightsRestoreData DeviceManagementUserRightsSettingable
    // This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
    userRightsTakeOwnership DeviceManagementUserRightsSettingable
    // Defender TamperProtection setting options
    windowsDefenderTamperProtection *WindowsDefenderTamperProtectionOptions
    // Possible values of xbox service start type
    xboxServicesAccessoryManagementServiceStartupMode *ServiceStartType
    // This setting determines whether xbox game save is enabled (1) or disabled (0).
    xboxServicesEnableXboxGameSaveTask *bool
    // Possible values of xbox service start type
    xboxServicesLiveAuthManagerServiceStartupMode *ServiceStartType
    // Possible values of xbox service start type
    xboxServicesLiveGameSaveServiceStartupMode *ServiceStartType
    // Possible values of xbox service start type
    xboxServicesLiveNetworkingServiceStartupMode *ServiceStartType
}
// NewWindows10EndpointProtectionConfiguration instantiates a new Windows10EndpointProtectionConfiguration and sets the default values.
func NewWindows10EndpointProtectionConfiguration()(*Windows10EndpointProtectionConfiguration) {
    m := &Windows10EndpointProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10EndpointProtectionConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EndpointProtectionConfiguration(), nil
}
// GetApplicationGuardAllowCameraMicrophoneRedirection gets the applicationGuardAllowCameraMicrophoneRedirection property value. Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowCameraMicrophoneRedirection()(*bool) {
    return m.applicationGuardAllowCameraMicrophoneRedirection
}
// GetApplicationGuardAllowFileSaveOnHost gets the applicationGuardAllowFileSaveOnHost property value. Allow users to download files from Edge in the application guard container and save them on the host file system
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowFileSaveOnHost()(*bool) {
    return m.applicationGuardAllowFileSaveOnHost
}
// GetApplicationGuardAllowPersistence gets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPersistence()(*bool) {
    return m.applicationGuardAllowPersistence
}
// GetApplicationGuardAllowPrintToLocalPrinters gets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToLocalPrinters()(*bool) {
    return m.applicationGuardAllowPrintToLocalPrinters
}
// GetApplicationGuardAllowPrintToNetworkPrinters gets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToNetworkPrinters()(*bool) {
    return m.applicationGuardAllowPrintToNetworkPrinters
}
// GetApplicationGuardAllowPrintToPDF gets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToPDF()(*bool) {
    return m.applicationGuardAllowPrintToPDF
}
// GetApplicationGuardAllowPrintToXPS gets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToXPS()(*bool) {
    return m.applicationGuardAllowPrintToXPS
}
// GetApplicationGuardAllowVirtualGPU gets the applicationGuardAllowVirtualGPU property value. Allow application guard to use virtual GPU
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowVirtualGPU()(*bool) {
    return m.applicationGuardAllowVirtualGPU
}
// GetApplicationGuardBlockClipboardSharing gets the applicationGuardBlockClipboardSharing property value. Possible values for applicationGuardBlockClipboardSharingType
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockClipboardSharing()(*ApplicationGuardBlockClipboardSharingType) {
    return m.applicationGuardBlockClipboardSharing
}
// GetApplicationGuardBlockFileTransfer gets the applicationGuardBlockFileTransfer property value. Possible values for applicationGuardBlockFileTransfer
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockFileTransfer()(*ApplicationGuardBlockFileTransferType) {
    return m.applicationGuardBlockFileTransfer
}
// GetApplicationGuardBlockNonEnterpriseContent gets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockNonEnterpriseContent()(*bool) {
    return m.applicationGuardBlockNonEnterpriseContent
}
// GetApplicationGuardCertificateThumbprints gets the applicationGuardCertificateThumbprints property value. Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardCertificateThumbprints()([]string) {
    return m.applicationGuardCertificateThumbprints
}
// GetApplicationGuardEnabled gets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardEnabled()(*bool) {
    return m.applicationGuardEnabled
}
// GetApplicationGuardEnabledOptions gets the applicationGuardEnabledOptions property value. Possible values for ApplicationGuardEnabledOptions
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardEnabledOptions()(*ApplicationGuardEnabledOptions) {
    return m.applicationGuardEnabledOptions
}
// GetApplicationGuardForceAuditing gets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardForceAuditing()(*bool) {
    return m.applicationGuardForceAuditing
}
// GetAppLockerApplicationControl gets the appLockerApplicationControl property value. Possible values of AppLocker Application Control Types
func (m *Windows10EndpointProtectionConfiguration) GetAppLockerApplicationControl()(*AppLockerApplicationControlType) {
    return m.appLockerApplicationControl
}
// GetBitLockerAllowStandardUserEncryption gets the bitLockerAllowStandardUserEncryption property value. Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerAllowStandardUserEncryption()(*bool) {
    return m.bitLockerAllowStandardUserEncryption
}
// GetBitLockerDisableWarningForOtherDiskEncryption gets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerDisableWarningForOtherDiskEncryption()(*bool) {
    return m.bitLockerDisableWarningForOtherDiskEncryption
}
// GetBitLockerEnableStorageCardEncryptionOnMobile gets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEnableStorageCardEncryptionOnMobile()(*bool) {
    return m.bitLockerEnableStorageCardEncryptionOnMobile
}
// GetBitLockerEncryptDevice gets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEncryptDevice()(*bool) {
    return m.bitLockerEncryptDevice
}
// GetBitLockerFixedDrivePolicy gets the bitLockerFixedDrivePolicy property value. BitLocker Fixed Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerFixedDrivePolicy()(BitLockerFixedDrivePolicyable) {
    return m.bitLockerFixedDrivePolicy
}
// GetBitLockerRecoveryPasswordRotation gets the bitLockerRecoveryPasswordRotation property value. BitLocker recovery password rotation type
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerRecoveryPasswordRotation()(*BitLockerRecoveryPasswordRotationType) {
    return m.bitLockerRecoveryPasswordRotation
}
// GetBitLockerRemovableDrivePolicy gets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerRemovableDrivePolicy()(BitLockerRemovableDrivePolicyable) {
    return m.bitLockerRemovableDrivePolicy
}
// GetBitLockerSystemDrivePolicy gets the bitLockerSystemDrivePolicy property value. BitLocker System Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerSystemDrivePolicy()(BitLockerSystemDrivePolicyable) {
    return m.bitLockerSystemDrivePolicy
}
// GetDefenderAdditionalGuardedFolders gets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdditionalGuardedFolders()([]string) {
    return m.defenderAdditionalGuardedFolders
}
// GetDefenderAdobeReaderLaunchChildProcess gets the defenderAdobeReaderLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdobeReaderLaunchChildProcess()(*DefenderProtectionType) {
    return m.defenderAdobeReaderLaunchChildProcess
}
// GetDefenderAdvancedRansomewareProtectionType gets the defenderAdvancedRansomewareProtectionType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdvancedRansomewareProtectionType()(*DefenderProtectionType) {
    return m.defenderAdvancedRansomewareProtectionType
}
// GetDefenderAllowBehaviorMonitoring gets the defenderAllowBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowBehaviorMonitoring()(*bool) {
    return m.defenderAllowBehaviorMonitoring
}
// GetDefenderAllowCloudProtection gets the defenderAllowCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowCloudProtection()(*bool) {
    return m.defenderAllowCloudProtection
}
// GetDefenderAllowEndUserAccess gets the defenderAllowEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowEndUserAccess()(*bool) {
    return m.defenderAllowEndUserAccess
}
// GetDefenderAllowIntrusionPreventionSystem gets the defenderAllowIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowIntrusionPreventionSystem()(*bool) {
    return m.defenderAllowIntrusionPreventionSystem
}
// GetDefenderAllowOnAccessProtection gets the defenderAllowOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowOnAccessProtection()(*bool) {
    return m.defenderAllowOnAccessProtection
}
// GetDefenderAllowRealTimeMonitoring gets the defenderAllowRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowRealTimeMonitoring()(*bool) {
    return m.defenderAllowRealTimeMonitoring
}
// GetDefenderAllowScanArchiveFiles gets the defenderAllowScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanArchiveFiles()(*bool) {
    return m.defenderAllowScanArchiveFiles
}
// GetDefenderAllowScanDownloads gets the defenderAllowScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanDownloads()(*bool) {
    return m.defenderAllowScanDownloads
}
// GetDefenderAllowScanNetworkFiles gets the defenderAllowScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanNetworkFiles()(*bool) {
    return m.defenderAllowScanNetworkFiles
}
// GetDefenderAllowScanRemovableDrivesDuringFullScan gets the defenderAllowScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanRemovableDrivesDuringFullScan()(*bool) {
    return m.defenderAllowScanRemovableDrivesDuringFullScan
}
// GetDefenderAllowScanScriptsLoadedInInternetExplorer gets the defenderAllowScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanScriptsLoadedInInternetExplorer()(*bool) {
    return m.defenderAllowScanScriptsLoadedInInternetExplorer
}
// GetDefenderAttackSurfaceReductionExcludedPaths gets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAttackSurfaceReductionExcludedPaths()([]string) {
    return m.defenderAttackSurfaceReductionExcludedPaths
}
// GetDefenderBlockEndUserAccess gets the defenderBlockEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderBlockEndUserAccess()(*bool) {
    return m.defenderBlockEndUserAccess
}
// GetDefenderBlockPersistenceThroughWmiType gets the defenderBlockPersistenceThroughWmiType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderBlockPersistenceThroughWmiType()(*DefenderAttackSurfaceType) {
    return m.defenderBlockPersistenceThroughWmiType
}
// GetDefenderCheckForSignaturesBeforeRunningScan gets the defenderCheckForSignaturesBeforeRunningScan property value. This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCheckForSignaturesBeforeRunningScan()(*bool) {
    return m.defenderCheckForSignaturesBeforeRunningScan
}
// GetDefenderCloudBlockLevel gets the defenderCloudBlockLevel property value. Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCloudBlockLevel()(*DefenderCloudBlockLevelType) {
    return m.defenderCloudBlockLevel
}
// GetDefenderCloudExtendedTimeoutInSeconds gets the defenderCloudExtendedTimeoutInSeconds property value. Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCloudExtendedTimeoutInSeconds()(*int32) {
    return m.defenderCloudExtendedTimeoutInSeconds
}
// GetDefenderDaysBeforeDeletingQuarantinedMalware gets the defenderDaysBeforeDeletingQuarantinedMalware property value. Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalware()(*int32) {
    return m.defenderDaysBeforeDeletingQuarantinedMalware
}
// GetDefenderDetectedMalwareActions gets the defenderDetectedMalwareActions property value. Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDetectedMalwareActions()(DefenderDetectedMalwareActionsable) {
    return m.defenderDetectedMalwareActions
}
// GetDefenderDisableBehaviorMonitoring gets the defenderDisableBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableBehaviorMonitoring()(*bool) {
    return m.defenderDisableBehaviorMonitoring
}
// GetDefenderDisableCatchupFullScan gets the defenderDisableCatchupFullScan property value. This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCatchupFullScan()(*bool) {
    return m.defenderDisableCatchupFullScan
}
// GetDefenderDisableCatchupQuickScan gets the defenderDisableCatchupQuickScan property value. This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCatchupQuickScan()(*bool) {
    return m.defenderDisableCatchupQuickScan
}
// GetDefenderDisableCloudProtection gets the defenderDisableCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCloudProtection()(*bool) {
    return m.defenderDisableCloudProtection
}
// GetDefenderDisableIntrusionPreventionSystem gets the defenderDisableIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableIntrusionPreventionSystem()(*bool) {
    return m.defenderDisableIntrusionPreventionSystem
}
// GetDefenderDisableOnAccessProtection gets the defenderDisableOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableOnAccessProtection()(*bool) {
    return m.defenderDisableOnAccessProtection
}
// GetDefenderDisableRealTimeMonitoring gets the defenderDisableRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableRealTimeMonitoring()(*bool) {
    return m.defenderDisableRealTimeMonitoring
}
// GetDefenderDisableScanArchiveFiles gets the defenderDisableScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanArchiveFiles()(*bool) {
    return m.defenderDisableScanArchiveFiles
}
// GetDefenderDisableScanDownloads gets the defenderDisableScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanDownloads()(*bool) {
    return m.defenderDisableScanDownloads
}
// GetDefenderDisableScanNetworkFiles gets the defenderDisableScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanNetworkFiles()(*bool) {
    return m.defenderDisableScanNetworkFiles
}
// GetDefenderDisableScanRemovableDrivesDuringFullScan gets the defenderDisableScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanRemovableDrivesDuringFullScan()(*bool) {
    return m.defenderDisableScanRemovableDrivesDuringFullScan
}
// GetDefenderDisableScanScriptsLoadedInInternetExplorer gets the defenderDisableScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanScriptsLoadedInInternetExplorer()(*bool) {
    return m.defenderDisableScanScriptsLoadedInInternetExplorer
}
// GetDefenderEmailContentExecution gets the defenderEmailContentExecution property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEmailContentExecution()(*DefenderProtectionType) {
    return m.defenderEmailContentExecution
}
// GetDefenderEmailContentExecutionType gets the defenderEmailContentExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEmailContentExecutionType()(*DefenderAttackSurfaceType) {
    return m.defenderEmailContentExecutionType
}
// GetDefenderEnableLowCpuPriority gets the defenderEnableLowCpuPriority property value. This policy setting allows you to enable or disable low CPU priority for scheduled scans.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableLowCpuPriority()(*bool) {
    return m.defenderEnableLowCpuPriority
}
// GetDefenderEnableScanIncomingMail gets the defenderEnableScanIncomingMail property value. Allows or disallows scanning of email.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableScanIncomingMail()(*bool) {
    return m.defenderEnableScanIncomingMail
}
// GetDefenderEnableScanMappedNetworkDrivesDuringFullScan gets the defenderEnableScanMappedNetworkDrivesDuringFullScan property value. Allows or disallows a full scan of mapped network drives.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableScanMappedNetworkDrivesDuringFullScan()(*bool) {
    return m.defenderEnableScanMappedNetworkDrivesDuringFullScan
}
// GetDefenderExploitProtectionXml gets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXml()([]byte) {
    return m.defenderExploitProtectionXml
}
// GetDefenderExploitProtectionXmlFileName gets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXmlFileName()(*string) {
    return m.defenderExploitProtectionXmlFileName
}
// GetDefenderFileExtensionsToExclude gets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderFileExtensionsToExclude()([]string) {
    return m.defenderFileExtensionsToExclude
}
// GetDefenderFilesAndFoldersToExclude gets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderFilesAndFoldersToExclude()([]string) {
    return m.defenderFilesAndFoldersToExclude
}
// GetDefenderGuardedFoldersAllowedAppPaths gets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
func (m *Windows10EndpointProtectionConfiguration) GetDefenderGuardedFoldersAllowedAppPaths()([]string) {
    return m.defenderGuardedFoldersAllowedAppPaths
}
// GetDefenderGuardMyFoldersType gets the defenderGuardMyFoldersType property value. Possible values of Folder Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderGuardMyFoldersType()(*FolderProtectionType) {
    return m.defenderGuardMyFoldersType
}
// GetDefenderNetworkProtectionType gets the defenderNetworkProtectionType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderNetworkProtectionType()(*DefenderProtectionType) {
    return m.defenderNetworkProtectionType
}
// GetDefenderOfficeAppsExecutableContentCreationOrLaunch gets the defenderOfficeAppsExecutableContentCreationOrLaunch property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsExecutableContentCreationOrLaunch()(*DefenderProtectionType) {
    return m.defenderOfficeAppsExecutableContentCreationOrLaunch
}
// GetDefenderOfficeAppsExecutableContentCreationOrLaunchType gets the defenderOfficeAppsExecutableContentCreationOrLaunchType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsExecutableContentCreationOrLaunchType()(*DefenderAttackSurfaceType) {
    return m.defenderOfficeAppsExecutableContentCreationOrLaunchType
}
// GetDefenderOfficeAppsLaunchChildProcess gets the defenderOfficeAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsLaunchChildProcess()(*DefenderProtectionType) {
    return m.defenderOfficeAppsLaunchChildProcess
}
// GetDefenderOfficeAppsLaunchChildProcessType gets the defenderOfficeAppsLaunchChildProcessType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsLaunchChildProcessType()(*DefenderAttackSurfaceType) {
    return m.defenderOfficeAppsLaunchChildProcessType
}
// GetDefenderOfficeAppsOtherProcessInjection gets the defenderOfficeAppsOtherProcessInjection property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsOtherProcessInjection()(*DefenderProtectionType) {
    return m.defenderOfficeAppsOtherProcessInjection
}
// GetDefenderOfficeAppsOtherProcessInjectionType gets the defenderOfficeAppsOtherProcessInjectionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsOtherProcessInjectionType()(*DefenderAttackSurfaceType) {
    return m.defenderOfficeAppsOtherProcessInjectionType
}
// GetDefenderOfficeCommunicationAppsLaunchChildProcess gets the defenderOfficeCommunicationAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeCommunicationAppsLaunchChildProcess()(*DefenderProtectionType) {
    return m.defenderOfficeCommunicationAppsLaunchChildProcess
}
// GetDefenderOfficeMacroCodeAllowWin32Imports gets the defenderOfficeMacroCodeAllowWin32Imports property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeMacroCodeAllowWin32Imports()(*DefenderProtectionType) {
    return m.defenderOfficeMacroCodeAllowWin32Imports
}
// GetDefenderOfficeMacroCodeAllowWin32ImportsType gets the defenderOfficeMacroCodeAllowWin32ImportsType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeMacroCodeAllowWin32ImportsType()(*DefenderAttackSurfaceType) {
    return m.defenderOfficeMacroCodeAllowWin32ImportsType
}
// GetDefenderPotentiallyUnwantedAppAction gets the defenderPotentiallyUnwantedAppAction property value. Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderPotentiallyUnwantedAppAction()(*DefenderProtectionType) {
    return m.defenderPotentiallyUnwantedAppAction
}
// GetDefenderPreventCredentialStealingType gets the defenderPreventCredentialStealingType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderPreventCredentialStealingType()(*DefenderProtectionType) {
    return m.defenderPreventCredentialStealingType
}
// GetDefenderProcessCreation gets the defenderProcessCreation property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessCreation()(*DefenderProtectionType) {
    return m.defenderProcessCreation
}
// GetDefenderProcessCreationType gets the defenderProcessCreationType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessCreationType()(*DefenderAttackSurfaceType) {
    return m.defenderProcessCreationType
}
// GetDefenderProcessesToExclude gets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessesToExclude()([]string) {
    return m.defenderProcessesToExclude
}
// GetDefenderScanDirection gets the defenderScanDirection property value. Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanDirection()(*DefenderRealtimeScanDirection) {
    return m.defenderScanDirection
}
// GetDefenderScanMaxCpuPercentage gets the defenderScanMaxCpuPercentage property value. Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanMaxCpuPercentage()(*int32) {
    return m.defenderScanMaxCpuPercentage
}
// GetDefenderScanType gets the defenderScanType property value. Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanType()(*DefenderScanType) {
    return m.defenderScanType
}
// GetDefenderScheduledQuickScanTime gets the defenderScheduledQuickScanTime property value. Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledQuickScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.defenderScheduledQuickScanTime
}
// GetDefenderScheduledScanDay gets the defenderScheduledScanDay property value. Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledScanDay()(*WeeklySchedule) {
    return m.defenderScheduledScanDay
}
// GetDefenderScheduledScanTime gets the defenderScheduledScanTime property value. Selects the time of day that the Windows Defender scan should run.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.defenderScheduledScanTime
}
// GetDefenderScriptDownloadedPayloadExecution gets the defenderScriptDownloadedPayloadExecution property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptDownloadedPayloadExecution()(*DefenderProtectionType) {
    return m.defenderScriptDownloadedPayloadExecution
}
// GetDefenderScriptDownloadedPayloadExecutionType gets the defenderScriptDownloadedPayloadExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptDownloadedPayloadExecutionType()(*DefenderAttackSurfaceType) {
    return m.defenderScriptDownloadedPayloadExecutionType
}
// GetDefenderScriptObfuscatedMacroCode gets the defenderScriptObfuscatedMacroCode property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptObfuscatedMacroCode()(*DefenderProtectionType) {
    return m.defenderScriptObfuscatedMacroCode
}
// GetDefenderScriptObfuscatedMacroCodeType gets the defenderScriptObfuscatedMacroCodeType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptObfuscatedMacroCodeType()(*DefenderAttackSurfaceType) {
    return m.defenderScriptObfuscatedMacroCodeType
}
// GetDefenderSecurityCenterBlockExploitProtectionOverride gets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterBlockExploitProtectionOverride()(*bool) {
    return m.defenderSecurityCenterBlockExploitProtectionOverride
}
// GetDefenderSecurityCenterDisableAccountUI gets the defenderSecurityCenterDisableAccountUI property value. Used to disable the display of the account protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableAccountUI()(*bool) {
    return m.defenderSecurityCenterDisableAccountUI
}
// GetDefenderSecurityCenterDisableAppBrowserUI gets the defenderSecurityCenterDisableAppBrowserUI property value. Used to disable the display of the app and browser protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableAppBrowserUI()(*bool) {
    return m.defenderSecurityCenterDisableAppBrowserUI
}
// GetDefenderSecurityCenterDisableClearTpmUI gets the defenderSecurityCenterDisableClearTpmUI property value. Used to disable the display of the Clear TPM button.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableClearTpmUI()(*bool) {
    return m.defenderSecurityCenterDisableClearTpmUI
}
// GetDefenderSecurityCenterDisableFamilyUI gets the defenderSecurityCenterDisableFamilyUI property value. Used to disable the display of the family options area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableFamilyUI()(*bool) {
    return m.defenderSecurityCenterDisableFamilyUI
}
// GetDefenderSecurityCenterDisableHardwareUI gets the defenderSecurityCenterDisableHardwareUI property value. Used to disable the display of the hardware protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableHardwareUI()(*bool) {
    return m.defenderSecurityCenterDisableHardwareUI
}
// GetDefenderSecurityCenterDisableHealthUI gets the defenderSecurityCenterDisableHealthUI property value. Used to disable the display of the device performance and health area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableHealthUI()(*bool) {
    return m.defenderSecurityCenterDisableHealthUI
}
// GetDefenderSecurityCenterDisableNetworkUI gets the defenderSecurityCenterDisableNetworkUI property value. Used to disable the display of the firewall and network protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableNetworkUI()(*bool) {
    return m.defenderSecurityCenterDisableNetworkUI
}
// GetDefenderSecurityCenterDisableNotificationAreaUI gets the defenderSecurityCenterDisableNotificationAreaUI property value. Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableNotificationAreaUI()(*bool) {
    return m.defenderSecurityCenterDisableNotificationAreaUI
}
// GetDefenderSecurityCenterDisableRansomwareUI gets the defenderSecurityCenterDisableRansomwareUI property value. Used to disable the display of the ransomware protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableRansomwareUI()(*bool) {
    return m.defenderSecurityCenterDisableRansomwareUI
}
// GetDefenderSecurityCenterDisableSecureBootUI gets the defenderSecurityCenterDisableSecureBootUI property value. Used to disable the display of the secure boot area under Device security.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableSecureBootUI()(*bool) {
    return m.defenderSecurityCenterDisableSecureBootUI
}
// GetDefenderSecurityCenterDisableTroubleshootingUI gets the defenderSecurityCenterDisableTroubleshootingUI property value. Used to disable the display of the security process troubleshooting under Device security.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableTroubleshootingUI()(*bool) {
    return m.defenderSecurityCenterDisableTroubleshootingUI
}
// GetDefenderSecurityCenterDisableVirusUI gets the defenderSecurityCenterDisableVirusUI property value. Used to disable the display of the virus and threat protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableVirusUI()(*bool) {
    return m.defenderSecurityCenterDisableVirusUI
}
// GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI gets the defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI property value. Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI()(*bool) {
    return m.defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI
}
// GetDefenderSecurityCenterHelpEmail gets the defenderSecurityCenterHelpEmail property value. The email address that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpEmail()(*string) {
    return m.defenderSecurityCenterHelpEmail
}
// GetDefenderSecurityCenterHelpPhone gets the defenderSecurityCenterHelpPhone property value. The phone number or Skype ID that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpPhone()(*string) {
    return m.defenderSecurityCenterHelpPhone
}
// GetDefenderSecurityCenterHelpURL gets the defenderSecurityCenterHelpURL property value. The help portal URL this is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpURL()(*string) {
    return m.defenderSecurityCenterHelpURL
}
// GetDefenderSecurityCenterITContactDisplay gets the defenderSecurityCenterITContactDisplay property value. Possible values for defenderSecurityCenterITContactDisplay
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterITContactDisplay()(*DefenderSecurityCenterITContactDisplayType) {
    return m.defenderSecurityCenterITContactDisplay
}
// GetDefenderSecurityCenterNotificationsFromApp gets the defenderSecurityCenterNotificationsFromApp property value. Possible values for defenderSecurityCenterNotificationsFromApp
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterNotificationsFromApp()(*DefenderSecurityCenterNotificationsFromAppType) {
    return m.defenderSecurityCenterNotificationsFromApp
}
// GetDefenderSecurityCenterOrganizationDisplayName gets the defenderSecurityCenterOrganizationDisplayName property value. The company name that is displayed to the users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterOrganizationDisplayName()(*string) {
    return m.defenderSecurityCenterOrganizationDisplayName
}
// GetDefenderSignatureUpdateIntervalInHours gets the defenderSignatureUpdateIntervalInHours property value. Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSignatureUpdateIntervalInHours()(*int32) {
    return m.defenderSignatureUpdateIntervalInHours
}
// GetDefenderSubmitSamplesConsentType gets the defenderSubmitSamplesConsentType property value. Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSubmitSamplesConsentType()(*DefenderSubmitSamplesConsentType) {
    return m.defenderSubmitSamplesConsentType
}
// GetDefenderUntrustedExecutable gets the defenderUntrustedExecutable property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedExecutable()(*DefenderProtectionType) {
    return m.defenderUntrustedExecutable
}
// GetDefenderUntrustedExecutableType gets the defenderUntrustedExecutableType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedExecutableType()(*DefenderAttackSurfaceType) {
    return m.defenderUntrustedExecutableType
}
// GetDefenderUntrustedUSBProcess gets the defenderUntrustedUSBProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedUSBProcess()(*DefenderProtectionType) {
    return m.defenderUntrustedUSBProcess
}
// GetDefenderUntrustedUSBProcessType gets the defenderUntrustedUSBProcessType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedUSBProcessType()(*DefenderAttackSurfaceType) {
    return m.defenderUntrustedUSBProcessType
}
// GetDeviceGuardEnableSecureBootWithDMA gets the deviceGuardEnableSecureBootWithDMA property value. This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardEnableSecureBootWithDMA()(*bool) {
    return m.deviceGuardEnableSecureBootWithDMA
}
// GetDeviceGuardEnableVirtualizationBasedSecurity gets the deviceGuardEnableVirtualizationBasedSecurity property value. Turns On Virtualization Based Security(VBS).
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardEnableVirtualizationBasedSecurity()(*bool) {
    return m.deviceGuardEnableVirtualizationBasedSecurity
}
// GetDeviceGuardLaunchSystemGuard gets the deviceGuardLaunchSystemGuard property value. Possible values of a property
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardLaunchSystemGuard()(*Enablement) {
    return m.deviceGuardLaunchSystemGuard
}
// GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings gets the deviceGuardLocalSystemAuthorityCredentialGuardSettings property value. Possible values of Credential Guard settings.
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings()(*DeviceGuardLocalSystemAuthorityCredentialGuardType) {
    return m.deviceGuardLocalSystemAuthorityCredentialGuardSettings
}
// GetDeviceGuardSecureBootWithDMA gets the deviceGuardSecureBootWithDMA property value. Possible values of Secure Boot with DMA
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardSecureBootWithDMA()(*SecureBootWithDMAType) {
    return m.deviceGuardSecureBootWithDMA
}
// GetDmaGuardDeviceEnumerationPolicy gets the dmaGuardDeviceEnumerationPolicy property value. Possible values of the DmaGuardDeviceEnumerationPolicy.
func (m *Windows10EndpointProtectionConfiguration) GetDmaGuardDeviceEnumerationPolicy()(*DmaGuardDeviceEnumerationPolicyType) {
    return m.dmaGuardDeviceEnumerationPolicy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10EndpointProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["applicationGuardAllowCameraMicrophoneRedirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowCameraMicrophoneRedirection(val)
        }
        return nil
    }
    res["applicationGuardAllowFileSaveOnHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowFileSaveOnHost(val)
        }
        return nil
    }
    res["applicationGuardAllowPersistence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPersistence(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToLocalPrinters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToLocalPrinters(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToNetworkPrinters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToNetworkPrinters(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToPDF"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToPDF(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToXPS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToXPS(val)
        }
        return nil
    }
    res["applicationGuardAllowVirtualGPU"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowVirtualGPU(val)
        }
        return nil
    }
    res["applicationGuardBlockClipboardSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardBlockClipboardSharingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockClipboardSharing(val.(*ApplicationGuardBlockClipboardSharingType))
        }
        return nil
    }
    res["applicationGuardBlockFileTransfer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardBlockFileTransferType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockFileTransfer(val.(*ApplicationGuardBlockFileTransferType))
        }
        return nil
    }
    res["applicationGuardBlockNonEnterpriseContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockNonEnterpriseContent(val)
        }
        return nil
    }
    res["applicationGuardCertificateThumbprints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetApplicationGuardCertificateThumbprints(res)
        }
        return nil
    }
    res["applicationGuardEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardEnabled(val)
        }
        return nil
    }
    res["applicationGuardEnabledOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardEnabledOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardEnabledOptions(val.(*ApplicationGuardEnabledOptions))
        }
        return nil
    }
    res["applicationGuardForceAuditing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardForceAuditing(val)
        }
        return nil
    }
    res["appLockerApplicationControl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLockerApplicationControlType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppLockerApplicationControl(val.(*AppLockerApplicationControlType))
        }
        return nil
    }
    res["bitLockerAllowStandardUserEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerAllowStandardUserEncryption(val)
        }
        return nil
    }
    res["bitLockerDisableWarningForOtherDiskEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerDisableWarningForOtherDiskEncryption(val)
        }
        return nil
    }
    res["bitLockerEnableStorageCardEncryptionOnMobile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEnableStorageCardEncryptionOnMobile(val)
        }
        return nil
    }
    res["bitLockerEncryptDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEncryptDevice(val)
        }
        return nil
    }
    res["bitLockerFixedDrivePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerFixedDrivePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerFixedDrivePolicy(val.(BitLockerFixedDrivePolicyable))
        }
        return nil
    }
    res["bitLockerRecoveryPasswordRotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBitLockerRecoveryPasswordRotationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerRecoveryPasswordRotation(val.(*BitLockerRecoveryPasswordRotationType))
        }
        return nil
    }
    res["bitLockerRemovableDrivePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerRemovableDrivePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerRemovableDrivePolicy(val.(BitLockerRemovableDrivePolicyable))
        }
        return nil
    }
    res["bitLockerSystemDrivePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerSystemDrivePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerSystemDrivePolicy(val.(BitLockerSystemDrivePolicyable))
        }
        return nil
    }
    res["defenderAdditionalGuardedFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderAdditionalGuardedFolders(res)
        }
        return nil
    }
    res["defenderAdobeReaderLaunchChildProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAdobeReaderLaunchChildProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderAdvancedRansomewareProtectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAdvancedRansomewareProtectionType(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderAllowBehaviorMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowBehaviorMonitoring(val)
        }
        return nil
    }
    res["defenderAllowCloudProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowCloudProtection(val)
        }
        return nil
    }
    res["defenderAllowEndUserAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowEndUserAccess(val)
        }
        return nil
    }
    res["defenderAllowIntrusionPreventionSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowIntrusionPreventionSystem(val)
        }
        return nil
    }
    res["defenderAllowOnAccessProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowOnAccessProtection(val)
        }
        return nil
    }
    res["defenderAllowRealTimeMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowRealTimeMonitoring(val)
        }
        return nil
    }
    res["defenderAllowScanArchiveFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanArchiveFiles(val)
        }
        return nil
    }
    res["defenderAllowScanDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanDownloads(val)
        }
        return nil
    }
    res["defenderAllowScanNetworkFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanNetworkFiles(val)
        }
        return nil
    }
    res["defenderAllowScanRemovableDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanRemovableDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderAllowScanScriptsLoadedInInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanScriptsLoadedInInternetExplorer(val)
        }
        return nil
    }
    res["defenderAttackSurfaceReductionExcludedPaths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderAttackSurfaceReductionExcludedPaths(res)
        }
        return nil
    }
    res["defenderBlockEndUserAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderBlockEndUserAccess(val)
        }
        return nil
    }
    res["defenderBlockPersistenceThroughWmiType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderBlockPersistenceThroughWmiType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderCheckForSignaturesBeforeRunningScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderCheckForSignaturesBeforeRunningScan(val)
        }
        return nil
    }
    res["defenderCloudBlockLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderCloudBlockLevelType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderCloudBlockLevel(val.(*DefenderCloudBlockLevelType))
        }
        return nil
    }
    res["defenderCloudExtendedTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderCloudExtendedTimeoutInSeconds(val)
        }
        return nil
    }
    res["defenderDaysBeforeDeletingQuarantinedMalware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDaysBeforeDeletingQuarantinedMalware(val)
        }
        return nil
    }
    res["defenderDetectedMalwareActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDefenderDetectedMalwareActionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDetectedMalwareActions(val.(DefenderDetectedMalwareActionsable))
        }
        return nil
    }
    res["defenderDisableBehaviorMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableBehaviorMonitoring(val)
        }
        return nil
    }
    res["defenderDisableCatchupFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableCatchupFullScan(val)
        }
        return nil
    }
    res["defenderDisableCatchupQuickScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableCatchupQuickScan(val)
        }
        return nil
    }
    res["defenderDisableCloudProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableCloudProtection(val)
        }
        return nil
    }
    res["defenderDisableIntrusionPreventionSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableIntrusionPreventionSystem(val)
        }
        return nil
    }
    res["defenderDisableOnAccessProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableOnAccessProtection(val)
        }
        return nil
    }
    res["defenderDisableRealTimeMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableRealTimeMonitoring(val)
        }
        return nil
    }
    res["defenderDisableScanArchiveFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanArchiveFiles(val)
        }
        return nil
    }
    res["defenderDisableScanDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanDownloads(val)
        }
        return nil
    }
    res["defenderDisableScanNetworkFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanNetworkFiles(val)
        }
        return nil
    }
    res["defenderDisableScanRemovableDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanRemovableDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderDisableScanScriptsLoadedInInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanScriptsLoadedInInternetExplorer(val)
        }
        return nil
    }
    res["defenderEmailContentExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEmailContentExecution(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderEmailContentExecutionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEmailContentExecutionType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderEnableLowCpuPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEnableLowCpuPriority(val)
        }
        return nil
    }
    res["defenderEnableScanIncomingMail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEnableScanIncomingMail(val)
        }
        return nil
    }
    res["defenderEnableScanMappedNetworkDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEnableScanMappedNetworkDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderExploitProtectionXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderExploitProtectionXml(val)
        }
        return nil
    }
    res["defenderExploitProtectionXmlFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderExploitProtectionXmlFileName(val)
        }
        return nil
    }
    res["defenderFileExtensionsToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderFileExtensionsToExclude(res)
        }
        return nil
    }
    res["defenderFilesAndFoldersToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderFilesAndFoldersToExclude(res)
        }
        return nil
    }
    res["defenderGuardedFoldersAllowedAppPaths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderGuardedFoldersAllowedAppPaths(res)
        }
        return nil
    }
    res["defenderGuardMyFoldersType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFolderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderGuardMyFoldersType(val.(*FolderProtectionType))
        }
        return nil
    }
    res["defenderNetworkProtectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderNetworkProtectionType(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsExecutableContentCreationOrLaunch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsExecutableContentCreationOrLaunch(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsExecutableContentCreationOrLaunchType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsExecutableContentCreationOrLaunchType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderOfficeAppsLaunchChildProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsLaunchChildProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsLaunchChildProcessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsLaunchChildProcessType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderOfficeAppsOtherProcessInjection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsOtherProcessInjection(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsOtherProcessInjectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsOtherProcessInjectionType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderOfficeCommunicationAppsLaunchChildProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeCommunicationAppsLaunchChildProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeMacroCodeAllowWin32Imports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeMacroCodeAllowWin32Imports(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeMacroCodeAllowWin32ImportsType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeMacroCodeAllowWin32ImportsType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderPotentiallyUnwantedAppAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderPotentiallyUnwantedAppAction(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderPreventCredentialStealingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderPreventCredentialStealingType(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderProcessCreation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderProcessCreation(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderProcessCreationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderProcessCreationType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderProcessesToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderProcessesToExclude(res)
        }
        return nil
    }
    res["defenderScanDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderRealtimeScanDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanDirection(val.(*DefenderRealtimeScanDirection))
        }
        return nil
    }
    res["defenderScanMaxCpuPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanMaxCpuPercentage(val)
        }
        return nil
    }
    res["defenderScanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderScanType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanType(val.(*DefenderScanType))
        }
        return nil
    }
    res["defenderScheduledQuickScanTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledQuickScanTime(val)
        }
        return nil
    }
    res["defenderScheduledScanDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeeklySchedule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledScanDay(val.(*WeeklySchedule))
        }
        return nil
    }
    res["defenderScheduledScanTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledScanTime(val)
        }
        return nil
    }
    res["defenderScriptDownloadedPayloadExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptDownloadedPayloadExecution(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderScriptDownloadedPayloadExecutionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptDownloadedPayloadExecutionType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderScriptObfuscatedMacroCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptObfuscatedMacroCode(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderScriptObfuscatedMacroCodeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptObfuscatedMacroCodeType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderSecurityCenterBlockExploitProtectionOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterBlockExploitProtectionOverride(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableAccountUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableAccountUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableAppBrowserUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableAppBrowserUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableClearTpmUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableClearTpmUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableFamilyUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableFamilyUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableHardwareUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableHardwareUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableHealthUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableHealthUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableNetworkUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableNetworkUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableNotificationAreaUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableNotificationAreaUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableRansomwareUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableRansomwareUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableSecureBootUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableSecureBootUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableTroubleshootingUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableTroubleshootingUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableVirusUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableVirusUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterHelpEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterHelpEmail(val)
        }
        return nil
    }
    res["defenderSecurityCenterHelpPhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterHelpPhone(val)
        }
        return nil
    }
    res["defenderSecurityCenterHelpURL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterHelpURL(val)
        }
        return nil
    }
    res["defenderSecurityCenterITContactDisplay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderSecurityCenterITContactDisplayType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterITContactDisplay(val.(*DefenderSecurityCenterITContactDisplayType))
        }
        return nil
    }
    res["defenderSecurityCenterNotificationsFromApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderSecurityCenterNotificationsFromAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterNotificationsFromApp(val.(*DefenderSecurityCenterNotificationsFromAppType))
        }
        return nil
    }
    res["defenderSecurityCenterOrganizationDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterOrganizationDisplayName(val)
        }
        return nil
    }
    res["defenderSignatureUpdateIntervalInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSignatureUpdateIntervalInHours(val)
        }
        return nil
    }
    res["defenderSubmitSamplesConsentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderSubmitSamplesConsentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSubmitSamplesConsentType(val.(*DefenderSubmitSamplesConsentType))
        }
        return nil
    }
    res["defenderUntrustedExecutable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedExecutable(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderUntrustedExecutableType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedExecutableType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderUntrustedUSBProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedUSBProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderUntrustedUSBProcessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedUSBProcessType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["deviceGuardEnableSecureBootWithDMA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardEnableSecureBootWithDMA(val)
        }
        return nil
    }
    res["deviceGuardEnableVirtualizationBasedSecurity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardEnableVirtualizationBasedSecurity(val)
        }
        return nil
    }
    res["deviceGuardLaunchSystemGuard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardLaunchSystemGuard(val.(*Enablement))
        }
        return nil
    }
    res["deviceGuardLocalSystemAuthorityCredentialGuardSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardLocalSystemAuthorityCredentialGuardType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings(val.(*DeviceGuardLocalSystemAuthorityCredentialGuardType))
        }
        return nil
    }
    res["deviceGuardSecureBootWithDMA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecureBootWithDMAType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardSecureBootWithDMA(val.(*SecureBootWithDMAType))
        }
        return nil
    }
    res["dmaGuardDeviceEnumerationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDmaGuardDeviceEnumerationPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDmaGuardDeviceEnumerationPolicy(val.(*DmaGuardDeviceEnumerationPolicyType))
        }
        return nil
    }
    res["firewallBlockStatefulFTP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallBlockStatefulFTP(val)
        }
        return nil
    }
    res["firewallCertificateRevocationListCheckMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallCertificateRevocationListCheckMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallCertificateRevocationListCheckMethod(val.(*FirewallCertificateRevocationListCheckMethodType))
        }
        return nil
    }
    res["firewallIdleTimeoutForSecurityAssociationInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIdleTimeoutForSecurityAssociationInSeconds(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowDHCP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowDHCP(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowICMP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowICMP(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowNeighborDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowNeighborDiscovery(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowRouterDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowRouterDiscovery(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsNone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsNone(val)
        }
        return nil
    }
    res["firewallMergeKeyingModuleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallMergeKeyingModuleSettings(val)
        }
        return nil
    }
    res["firewallPacketQueueingMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallPacketQueueingMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallPacketQueueingMethod(val.(*FirewallPacketQueueingMethodType))
        }
        return nil
    }
    res["firewallPreSharedKeyEncodingMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallPreSharedKeyEncodingMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallPreSharedKeyEncodingMethod(val.(*FirewallPreSharedKeyEncodingMethodType))
        }
        return nil
    }
    res["firewallProfileDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfileDomain(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallProfilePrivate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfilePrivate(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallProfilePublic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfilePublic(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsFirewallRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsFirewallRuleable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsFirewallRuleable)
            }
            m.SetFirewallRules(res)
        }
        return nil
    }
    res["lanManagerAuthenticationLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanManagerAuthenticationLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanManagerAuthenticationLevel(val.(*LanManagerAuthenticationLevel))
        }
        return nil
    }
    res["lanManagerWorkstationDisableInsecureGuestLogons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanManagerWorkstationDisableInsecureGuestLogons(val)
        }
        return nil
    }
    res["localSecurityOptionsAdministratorAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAdministratorAccountName(val)
        }
        return nil
    }
    res["localSecurityOptionsAdministratorElevationPromptBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsAdministratorElevationPromptBehaviorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAdministratorElevationPromptBehavior(val.(*LocalSecurityOptionsAdministratorElevationPromptBehaviorType))
        }
        return nil
    }
    res["localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowPKU2UAuthenticationRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowUIAccessApplicationElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowUIAccessApplicationElevation(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowUIAccessApplicationsForSecureLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowUndockWithoutHavingToLogon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockMicrosoftAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockMicrosoftAccounts(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockRemoteLogonWithBlankPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockRemoteOpticalDriveAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockUsersInstallingPrinterDrivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers(val)
        }
        return nil
    }
    res["localSecurityOptionsClearVirtualMemoryPageFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsClearVirtualMemoryPageFile(val)
        }
        return nil
    }
    res["localSecurityOptionsClientDigitallySignCommunicationsAlways"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways(val)
        }
        return nil
    }
    res["localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers(val)
        }
        return nil
    }
    res["localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableAdministratorAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableAdministratorAccount(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableGuestAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableGuestAccount(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableServerDigitallySignCommunicationsAlways"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees(val)
        }
        return nil
    }
    res["localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts(val)
        }
        return nil
    }
    res["localSecurityOptionsDoNotRequireCtrlAltDel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDoNotRequireCtrlAltDel(val)
        }
        return nil
    }
    res["localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange(val)
        }
        return nil
    }
    res["localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(val.(*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType))
        }
        return nil
    }
    res["localSecurityOptionsGuestAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsGuestAccountName(val)
        }
        return nil
    }
    res["localSecurityOptionsHideLastSignedInUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsHideLastSignedInUser(val)
        }
        return nil
    }
    res["localSecurityOptionsHideUsernameAtSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsHideUsernameAtSignIn(val)
        }
        return nil
    }
    res["localSecurityOptionsInformationDisplayedOnLockScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsInformationDisplayedOnLockScreenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsInformationDisplayedOnLockScreen(val.(*LocalSecurityOptionsInformationDisplayedOnLockScreenType))
        }
        return nil
    }
    res["localSecurityOptionsInformationShownOnLockScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsInformationShownOnLockScreenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsInformationShownOnLockScreen(val.(*LocalSecurityOptionsInformationShownOnLockScreenType))
        }
        return nil
    }
    res["localSecurityOptionsLogOnMessageText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsLogOnMessageText(val)
        }
        return nil
    }
    res["localSecurityOptionsLogOnMessageTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsLogOnMessageTitle(val)
        }
        return nil
    }
    res["localSecurityOptionsMachineInactivityLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMachineInactivityLimit(val)
        }
        return nil
    }
    res["localSecurityOptionsMachineInactivityLimitInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMachineInactivityLimitInMinutes(val)
        }
        return nil
    }
    res["localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsMinimumSessionSecurity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(val.(*LocalSecurityOptionsMinimumSessionSecurity))
        }
        return nil
    }
    res["localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsMinimumSessionSecurity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(val.(*LocalSecurityOptionsMinimumSessionSecurity))
        }
        return nil
    }
    res["localSecurityOptionsOnlyElevateSignedExecutables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsOnlyElevateSignedExecutables(val)
        }
        return nil
    }
    res["localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares(val)
        }
        return nil
    }
    res["localSecurityOptionsSmartCardRemovalBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsSmartCardRemovalBehaviorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsSmartCardRemovalBehavior(val.(*LocalSecurityOptionsSmartCardRemovalBehaviorType))
        }
        return nil
    }
    res["localSecurityOptionsStandardUserElevationPromptBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsStandardUserElevationPromptBehaviorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsStandardUserElevationPromptBehavior(val.(*LocalSecurityOptionsStandardUserElevationPromptBehaviorType))
        }
        return nil
    }
    res["localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation(val)
        }
        return nil
    }
    res["localSecurityOptionsUseAdminApprovalMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsUseAdminApprovalMode(val)
        }
        return nil
    }
    res["localSecurityOptionsUseAdminApprovalModeForAdministrators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators(val)
        }
        return nil
    }
    res["localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations(val)
        }
        return nil
    }
    res["smartScreenBlockOverrideForFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenBlockOverrideForFiles(val)
        }
        return nil
    }
    res["smartScreenEnableInShell"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenEnableInShell(val)
        }
        return nil
    }
    res["userRightsAccessCredentialManagerAsTrustedCaller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsAccessCredentialManagerAsTrustedCaller(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsActAsPartOfTheOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsActAsPartOfTheOperatingSystem(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsAllowAccessFromNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsAllowAccessFromNetwork(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsBackupData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsBackupData(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsBlockAccessFromNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsBlockAccessFromNetwork(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsChangeSystemTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsChangeSystemTime(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreateGlobalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreateGlobalObjects(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreatePageFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreatePageFile(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreatePermanentSharedObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreatePermanentSharedObjects(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreateSymbolicLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreateSymbolicLinks(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreateToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreateToken(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsDebugPrograms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsDebugPrograms(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsDelegation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsDelegation(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsDenyLocalLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsDenyLocalLogOn(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsGenerateSecurityAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsGenerateSecurityAudits(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsImpersonateClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsImpersonateClient(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsIncreaseSchedulingPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsIncreaseSchedulingPriority(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsLoadUnloadDrivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsLoadUnloadDrivers(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsLocalLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsLocalLogOn(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsLockMemory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsLockMemory(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsManageAuditingAndSecurityLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsManageAuditingAndSecurityLogs(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsManageVolumes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsManageVolumes(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsModifyFirmwareEnvironment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsModifyFirmwareEnvironment(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsModifyObjectLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsModifyObjectLabels(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsProfileSingleProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsProfileSingleProcess(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsRemoteDesktopServicesLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsRemoteDesktopServicesLogOn(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsRemoteShutdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsRemoteShutdown(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsRestoreData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsRestoreData(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsTakeOwnership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsTakeOwnership(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["windowsDefenderTamperProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderTamperProtectionOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsDefenderTamperProtection(val.(*WindowsDefenderTamperProtectionOptions))
        }
        return nil
    }
    res["xboxServicesAccessoryManagementServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesAccessoryManagementServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    res["xboxServicesEnableXboxGameSaveTask"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesEnableXboxGameSaveTask(val)
        }
        return nil
    }
    res["xboxServicesLiveAuthManagerServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesLiveAuthManagerServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    res["xboxServicesLiveGameSaveServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesLiveGameSaveServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    res["xboxServicesLiveNetworkingServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesLiveNetworkingServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    return res
}
// GetFirewallBlockStatefulFTP gets the firewallBlockStatefulFTP property value. Blocks stateful FTP connections to the device
func (m *Windows10EndpointProtectionConfiguration) GetFirewallBlockStatefulFTP()(*bool) {
    return m.firewallBlockStatefulFTP
}
// GetFirewallCertificateRevocationListCheckMethod gets the firewallCertificateRevocationListCheckMethod property value. Possible values for firewallCertificateRevocationListCheckMethod
func (m *Windows10EndpointProtectionConfiguration) GetFirewallCertificateRevocationListCheckMethod()(*FirewallCertificateRevocationListCheckMethodType) {
    return m.firewallCertificateRevocationListCheckMethod
}
// GetFirewallIdleTimeoutForSecurityAssociationInSeconds gets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIdleTimeoutForSecurityAssociationInSeconds()(*int32) {
    return m.firewallIdleTimeoutForSecurityAssociationInSeconds
}
// GetFirewallIPSecExemptionsAllowDHCP gets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowDHCP()(*bool) {
    return m.firewallIPSecExemptionsAllowDHCP
}
// GetFirewallIPSecExemptionsAllowICMP gets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowICMP()(*bool) {
    return m.firewallIPSecExemptionsAllowICMP
}
// GetFirewallIPSecExemptionsAllowNeighborDiscovery gets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowNeighborDiscovery()(*bool) {
    return m.firewallIPSecExemptionsAllowNeighborDiscovery
}
// GetFirewallIPSecExemptionsAllowRouterDiscovery gets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowRouterDiscovery()(*bool) {
    return m.firewallIPSecExemptionsAllowRouterDiscovery
}
// GetFirewallIPSecExemptionsNone gets the firewallIPSecExemptionsNone property value. Configures IPSec exemptions to no exemptions
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsNone()(*bool) {
    return m.firewallIPSecExemptionsNone
}
// GetFirewallMergeKeyingModuleSettings gets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
func (m *Windows10EndpointProtectionConfiguration) GetFirewallMergeKeyingModuleSettings()(*bool) {
    return m.firewallMergeKeyingModuleSettings
}
// GetFirewallPacketQueueingMethod gets the firewallPacketQueueingMethod property value. Possible values for firewallPacketQueueingMethod
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPacketQueueingMethod()(*FirewallPacketQueueingMethodType) {
    return m.firewallPacketQueueingMethod
}
// GetFirewallPreSharedKeyEncodingMethod gets the firewallPreSharedKeyEncodingMethod property value. Possible values for firewallPreSharedKeyEncodingMethod
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPreSharedKeyEncodingMethod()(*FirewallPreSharedKeyEncodingMethodType) {
    return m.firewallPreSharedKeyEncodingMethod
}
// GetFirewallProfileDomain gets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfileDomain()(WindowsFirewallNetworkProfileable) {
    return m.firewallProfileDomain
}
// GetFirewallProfilePrivate gets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePrivate()(WindowsFirewallNetworkProfileable) {
    return m.firewallProfilePrivate
}
// GetFirewallProfilePublic gets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePublic()(WindowsFirewallNetworkProfileable) {
    return m.firewallProfilePublic
}
// GetFirewallRules gets the firewallRules property value. Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
func (m *Windows10EndpointProtectionConfiguration) GetFirewallRules()([]WindowsFirewallRuleable) {
    return m.firewallRules
}
// GetLanManagerAuthenticationLevel gets the lanManagerAuthenticationLevel property value. Possible values for LanManagerAuthenticationLevel
func (m *Windows10EndpointProtectionConfiguration) GetLanManagerAuthenticationLevel()(*LanManagerAuthenticationLevel) {
    return m.lanManagerAuthenticationLevel
}
// GetLanManagerWorkstationDisableInsecureGuestLogons gets the lanManagerWorkstationDisableInsecureGuestLogons property value. If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
func (m *Windows10EndpointProtectionConfiguration) GetLanManagerWorkstationDisableInsecureGuestLogons()(*bool) {
    return m.lanManagerWorkstationDisableInsecureGuestLogons
}
// GetLocalSecurityOptionsAdministratorAccountName gets the localSecurityOptionsAdministratorAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAdministratorAccountName()(*string) {
    return m.localSecurityOptionsAdministratorAccountName
}
// GetLocalSecurityOptionsAdministratorElevationPromptBehavior gets the localSecurityOptionsAdministratorElevationPromptBehavior property value. Possible values for LocalSecurityOptionsAdministratorElevationPromptBehavior
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAdministratorElevationPromptBehavior()(*LocalSecurityOptionsAdministratorElevationPromptBehaviorType) {
    return m.localSecurityOptionsAdministratorElevationPromptBehavior
}
// GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares gets the localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares property value. This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares()(*bool) {
    return m.localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares
}
// GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests gets the localSecurityOptionsAllowPKU2UAuthenticationRequests property value. Block PKU2U authentication requests to this device to use online identities.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests()(*bool) {
    return m.localSecurityOptionsAllowPKU2UAuthenticationRequests
}
// GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager gets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager property value. Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager()(*string) {
    return m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager
}
// GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool gets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool property value. UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool()(*bool) {
    return m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool
}
// GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn gets the localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn property value. This security setting determines whether a computer can be shut down without having to log on to Windows.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn()(*bool) {
    return m.localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn
}
// GetLocalSecurityOptionsAllowUIAccessApplicationElevation gets the localSecurityOptionsAllowUIAccessApplicationElevation property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUIAccessApplicationElevation()(*bool) {
    return m.localSecurityOptionsAllowUIAccessApplicationElevation
}
// GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations gets the localSecurityOptionsAllowUIAccessApplicationsForSecureLocations property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations()(*bool) {
    return m.localSecurityOptionsAllowUIAccessApplicationsForSecureLocations
}
// GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon gets the localSecurityOptionsAllowUndockWithoutHavingToLogon property value. Prevent a portable computer from being undocked without having to log in.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon()(*bool) {
    return m.localSecurityOptionsAllowUndockWithoutHavingToLogon
}
// GetLocalSecurityOptionsBlockMicrosoftAccounts gets the localSecurityOptionsBlockMicrosoftAccounts property value. Prevent users from adding new Microsoft accounts to this computer.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockMicrosoftAccounts()(*bool) {
    return m.localSecurityOptionsBlockMicrosoftAccounts
}
// GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword gets the localSecurityOptionsBlockRemoteLogonWithBlankPassword property value. Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword()(*bool) {
    return m.localSecurityOptionsBlockRemoteLogonWithBlankPassword
}
// GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess gets the localSecurityOptionsBlockRemoteOpticalDriveAccess property value. Enabling this settings allows only interactively logged on user to access CD-ROM media.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess()(*bool) {
    return m.localSecurityOptionsBlockRemoteOpticalDriveAccess
}
// GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers gets the localSecurityOptionsBlockUsersInstallingPrinterDrivers property value. Restrict installing printer drivers as part of connecting to a shared printer to admins only.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers()(*bool) {
    return m.localSecurityOptionsBlockUsersInstallingPrinterDrivers
}
// GetLocalSecurityOptionsClearVirtualMemoryPageFile gets the localSecurityOptionsClearVirtualMemoryPageFile property value. This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClearVirtualMemoryPageFile()(*bool) {
    return m.localSecurityOptionsClearVirtualMemoryPageFile
}
// GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways gets the localSecurityOptionsClientDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB client component.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways()(*bool) {
    return m.localSecurityOptionsClientDigitallySignCommunicationsAlways
}
// GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers gets the localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers property value. If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers()(*bool) {
    return m.localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers
}
// GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation gets the localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation property value. App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation()(*bool) {
    return m.localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation
}
// GetLocalSecurityOptionsDisableAdministratorAccount gets the localSecurityOptionsDisableAdministratorAccount property value. Determines whether the Local Administrator account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableAdministratorAccount()(*bool) {
    return m.localSecurityOptionsDisableAdministratorAccount
}
// GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees gets the localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees property value. This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees()(*bool) {
    return m.localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees
}
// GetLocalSecurityOptionsDisableGuestAccount gets the localSecurityOptionsDisableGuestAccount property value. Determines if the Guest account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableGuestAccount()(*bool) {
    return m.localSecurityOptionsDisableGuestAccount
}
// GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways gets the localSecurityOptionsDisableServerDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB server component.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways()(*bool) {
    return m.localSecurityOptionsDisableServerDigitallySignCommunicationsAlways
}
// GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees gets the localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees property value. This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees()(*bool) {
    return m.localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees
}
// GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts gets the localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts property value. This security setting determines what additional permissions will be granted for anonymous connections to the computer.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts()(*bool) {
    return m.localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts
}
// GetLocalSecurityOptionsDoNotRequireCtrlAltDel gets the localSecurityOptionsDoNotRequireCtrlAltDel property value. Require CTRL+ALT+DEL to be pressed before a user can log on.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotRequireCtrlAltDel()(*bool) {
    return m.localSecurityOptionsDoNotRequireCtrlAltDel
}
// GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange gets the localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange property value. This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange()(*bool) {
    return m.localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange
}
// GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser gets the localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser property value. Possible values for LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser()(*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType) {
    return m.localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
}
// GetLocalSecurityOptionsGuestAccountName gets the localSecurityOptionsGuestAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsGuestAccountName()(*string) {
    return m.localSecurityOptionsGuestAccountName
}
// GetLocalSecurityOptionsHideLastSignedInUser gets the localSecurityOptionsHideLastSignedInUser property value. Do not display the username of the last person who signed in on this device.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsHideLastSignedInUser()(*bool) {
    return m.localSecurityOptionsHideLastSignedInUser
}
// GetLocalSecurityOptionsHideUsernameAtSignIn gets the localSecurityOptionsHideUsernameAtSignIn property value. Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsHideUsernameAtSignIn()(*bool) {
    return m.localSecurityOptionsHideUsernameAtSignIn
}
// GetLocalSecurityOptionsInformationDisplayedOnLockScreen gets the localSecurityOptionsInformationDisplayedOnLockScreen property value. Possible values for LocalSecurityOptionsInformationDisplayedOnLockScreen
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsInformationDisplayedOnLockScreen()(*LocalSecurityOptionsInformationDisplayedOnLockScreenType) {
    return m.localSecurityOptionsInformationDisplayedOnLockScreen
}
// GetLocalSecurityOptionsInformationShownOnLockScreen gets the localSecurityOptionsInformationShownOnLockScreen property value. Possible values for LocalSecurityOptionsInformationShownOnLockScreenType
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsInformationShownOnLockScreen()(*LocalSecurityOptionsInformationShownOnLockScreenType) {
    return m.localSecurityOptionsInformationShownOnLockScreen
}
// GetLocalSecurityOptionsLogOnMessageText gets the localSecurityOptionsLogOnMessageText property value. Set message text for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsLogOnMessageText()(*string) {
    return m.localSecurityOptionsLogOnMessageText
}
// GetLocalSecurityOptionsLogOnMessageTitle gets the localSecurityOptionsLogOnMessageTitle property value. Set message title for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsLogOnMessageTitle()(*string) {
    return m.localSecurityOptionsLogOnMessageTitle
}
// GetLocalSecurityOptionsMachineInactivityLimit gets the localSecurityOptionsMachineInactivityLimit property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMachineInactivityLimit()(*int32) {
    return m.localSecurityOptionsMachineInactivityLimit
}
// GetLocalSecurityOptionsMachineInactivityLimitInMinutes gets the localSecurityOptionsMachineInactivityLimitInMinutes property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMachineInactivityLimitInMinutes()(*int32) {
    return m.localSecurityOptionsMachineInactivityLimitInMinutes
}
// GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients gets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients()(*LocalSecurityOptionsMinimumSessionSecurity) {
    return m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients
}
// GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers gets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers()(*LocalSecurityOptionsMinimumSessionSecurity) {
    return m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers
}
// GetLocalSecurityOptionsOnlyElevateSignedExecutables gets the localSecurityOptionsOnlyElevateSignedExecutables property value. Enforce PKI certification path validation for a given executable file before it is permitted to run.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsOnlyElevateSignedExecutables()(*bool) {
    return m.localSecurityOptionsOnlyElevateSignedExecutables
}
// GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares gets the localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares property value. By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares()(*bool) {
    return m.localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares
}
// GetLocalSecurityOptionsSmartCardRemovalBehavior gets the localSecurityOptionsSmartCardRemovalBehavior property value. Possible values for LocalSecurityOptionsSmartCardRemovalBehaviorType
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsSmartCardRemovalBehavior()(*LocalSecurityOptionsSmartCardRemovalBehaviorType) {
    return m.localSecurityOptionsSmartCardRemovalBehavior
}
// GetLocalSecurityOptionsStandardUserElevationPromptBehavior gets the localSecurityOptionsStandardUserElevationPromptBehavior property value. Possible values for LocalSecurityOptionsStandardUserElevationPromptBehavior
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsStandardUserElevationPromptBehavior()(*LocalSecurityOptionsStandardUserElevationPromptBehaviorType) {
    return m.localSecurityOptionsStandardUserElevationPromptBehavior
}
// GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation gets the localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation property value. Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation()(*bool) {
    return m.localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation
}
// GetLocalSecurityOptionsUseAdminApprovalMode gets the localSecurityOptionsUseAdminApprovalMode property value. Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsUseAdminApprovalMode()(*bool) {
    return m.localSecurityOptionsUseAdminApprovalMode
}
// GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators gets the localSecurityOptionsUseAdminApprovalModeForAdministrators property value. Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators()(*bool) {
    return m.localSecurityOptionsUseAdminApprovalModeForAdministrators
}
// GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations gets the localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations property value. Virtualize file and registry write failures to per user locations
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations()(*bool) {
    return m.localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations
}
// GetSmartScreenBlockOverrideForFiles gets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenBlockOverrideForFiles()(*bool) {
    return m.smartScreenBlockOverrideForFiles
}
// GetSmartScreenEnableInShell gets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenEnableInShell()(*bool) {
    return m.smartScreenEnableInShell
}
// GetUserRightsAccessCredentialManagerAsTrustedCaller gets the userRightsAccessCredentialManagerAsTrustedCaller property value. This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsAccessCredentialManagerAsTrustedCaller()(DeviceManagementUserRightsSettingable) {
    return m.userRightsAccessCredentialManagerAsTrustedCaller
}
// GetUserRightsActAsPartOfTheOperatingSystem gets the userRightsActAsPartOfTheOperatingSystem property value. This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsActAsPartOfTheOperatingSystem()(DeviceManagementUserRightsSettingable) {
    return m.userRightsActAsPartOfTheOperatingSystem
}
// GetUserRightsAllowAccessFromNetwork gets the userRightsAllowAccessFromNetwork property value. This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsAllowAccessFromNetwork()(DeviceManagementUserRightsSettingable) {
    return m.userRightsAllowAccessFromNetwork
}
// GetUserRightsBackupData gets the userRightsBackupData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsBackupData()(DeviceManagementUserRightsSettingable) {
    return m.userRightsBackupData
}
// GetUserRightsBlockAccessFromNetwork gets the userRightsBlockAccessFromNetwork property value. This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsBlockAccessFromNetwork()(DeviceManagementUserRightsSettingable) {
    return m.userRightsBlockAccessFromNetwork
}
// GetUserRightsChangeSystemTime gets the userRightsChangeSystemTime property value. This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsChangeSystemTime()(DeviceManagementUserRightsSettingable) {
    return m.userRightsChangeSystemTime
}
// GetUserRightsCreateGlobalObjects gets the userRightsCreateGlobalObjects property value. This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateGlobalObjects()(DeviceManagementUserRightsSettingable) {
    return m.userRightsCreateGlobalObjects
}
// GetUserRightsCreatePageFile gets the userRightsCreatePageFile property value. This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreatePageFile()(DeviceManagementUserRightsSettingable) {
    return m.userRightsCreatePageFile
}
// GetUserRightsCreatePermanentSharedObjects gets the userRightsCreatePermanentSharedObjects property value. This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreatePermanentSharedObjects()(DeviceManagementUserRightsSettingable) {
    return m.userRightsCreatePermanentSharedObjects
}
// GetUserRightsCreateSymbolicLinks gets the userRightsCreateSymbolicLinks property value. This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateSymbolicLinks()(DeviceManagementUserRightsSettingable) {
    return m.userRightsCreateSymbolicLinks
}
// GetUserRightsCreateToken gets the userRightsCreateToken property value. This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateToken()(DeviceManagementUserRightsSettingable) {
    return m.userRightsCreateToken
}
// GetUserRightsDebugPrograms gets the userRightsDebugPrograms property value. This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDebugPrograms()(DeviceManagementUserRightsSettingable) {
    return m.userRightsDebugPrograms
}
// GetUserRightsDelegation gets the userRightsDelegation property value. This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDelegation()(DeviceManagementUserRightsSettingable) {
    return m.userRightsDelegation
}
// GetUserRightsDenyLocalLogOn gets the userRightsDenyLocalLogOn property value. This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDenyLocalLogOn()(DeviceManagementUserRightsSettingable) {
    return m.userRightsDenyLocalLogOn
}
// GetUserRightsGenerateSecurityAudits gets the userRightsGenerateSecurityAudits property value. This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsGenerateSecurityAudits()(DeviceManagementUserRightsSettingable) {
    return m.userRightsGenerateSecurityAudits
}
// GetUserRightsImpersonateClient gets the userRightsImpersonateClient property value. Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsImpersonateClient()(DeviceManagementUserRightsSettingable) {
    return m.userRightsImpersonateClient
}
// GetUserRightsIncreaseSchedulingPriority gets the userRightsIncreaseSchedulingPriority property value. This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsIncreaseSchedulingPriority()(DeviceManagementUserRightsSettingable) {
    return m.userRightsIncreaseSchedulingPriority
}
// GetUserRightsLoadUnloadDrivers gets the userRightsLoadUnloadDrivers property value. This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLoadUnloadDrivers()(DeviceManagementUserRightsSettingable) {
    return m.userRightsLoadUnloadDrivers
}
// GetUserRightsLocalLogOn gets the userRightsLocalLogOn property value. This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLocalLogOn()(DeviceManagementUserRightsSettingable) {
    return m.userRightsLocalLogOn
}
// GetUserRightsLockMemory gets the userRightsLockMemory property value. This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLockMemory()(DeviceManagementUserRightsSettingable) {
    return m.userRightsLockMemory
}
// GetUserRightsManageAuditingAndSecurityLogs gets the userRightsManageAuditingAndSecurityLogs property value. This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsManageAuditingAndSecurityLogs()(DeviceManagementUserRightsSettingable) {
    return m.userRightsManageAuditingAndSecurityLogs
}
// GetUserRightsManageVolumes gets the userRightsManageVolumes property value. This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsManageVolumes()(DeviceManagementUserRightsSettingable) {
    return m.userRightsManageVolumes
}
// GetUserRightsModifyFirmwareEnvironment gets the userRightsModifyFirmwareEnvironment property value. This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsModifyFirmwareEnvironment()(DeviceManagementUserRightsSettingable) {
    return m.userRightsModifyFirmwareEnvironment
}
// GetUserRightsModifyObjectLabels gets the userRightsModifyObjectLabels property value. This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsModifyObjectLabels()(DeviceManagementUserRightsSettingable) {
    return m.userRightsModifyObjectLabels
}
// GetUserRightsProfileSingleProcess gets the userRightsProfileSingleProcess property value. This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsProfileSingleProcess()(DeviceManagementUserRightsSettingable) {
    return m.userRightsProfileSingleProcess
}
// GetUserRightsRemoteDesktopServicesLogOn gets the userRightsRemoteDesktopServicesLogOn property value. This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRemoteDesktopServicesLogOn()(DeviceManagementUserRightsSettingable) {
    return m.userRightsRemoteDesktopServicesLogOn
}
// GetUserRightsRemoteShutdown gets the userRightsRemoteShutdown property value. This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRemoteShutdown()(DeviceManagementUserRightsSettingable) {
    return m.userRightsRemoteShutdown
}
// GetUserRightsRestoreData gets the userRightsRestoreData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRestoreData()(DeviceManagementUserRightsSettingable) {
    return m.userRightsRestoreData
}
// GetUserRightsTakeOwnership gets the userRightsTakeOwnership property value. This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsTakeOwnership()(DeviceManagementUserRightsSettingable) {
    return m.userRightsTakeOwnership
}
// GetWindowsDefenderTamperProtection gets the windowsDefenderTamperProtection property value. Defender TamperProtection setting options
func (m *Windows10EndpointProtectionConfiguration) GetWindowsDefenderTamperProtection()(*WindowsDefenderTamperProtectionOptions) {
    return m.windowsDefenderTamperProtection
}
// GetXboxServicesAccessoryManagementServiceStartupMode gets the xboxServicesAccessoryManagementServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesAccessoryManagementServiceStartupMode()(*ServiceStartType) {
    return m.xboxServicesAccessoryManagementServiceStartupMode
}
// GetXboxServicesEnableXboxGameSaveTask gets the xboxServicesEnableXboxGameSaveTask property value. This setting determines whether xbox game save is enabled (1) or disabled (0).
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesEnableXboxGameSaveTask()(*bool) {
    return m.xboxServicesEnableXboxGameSaveTask
}
// GetXboxServicesLiveAuthManagerServiceStartupMode gets the xboxServicesLiveAuthManagerServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveAuthManagerServiceStartupMode()(*ServiceStartType) {
    return m.xboxServicesLiveAuthManagerServiceStartupMode
}
// GetXboxServicesLiveGameSaveServiceStartupMode gets the xboxServicesLiveGameSaveServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveGameSaveServiceStartupMode()(*ServiceStartType) {
    return m.xboxServicesLiveGameSaveServiceStartupMode
}
// GetXboxServicesLiveNetworkingServiceStartupMode gets the xboxServicesLiveNetworkingServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveNetworkingServiceStartupMode()(*ServiceStartType) {
    return m.xboxServicesLiveNetworkingServiceStartupMode
}
// Serialize serializes information the current object
func (m *Windows10EndpointProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowCameraMicrophoneRedirection", m.GetApplicationGuardAllowCameraMicrophoneRedirection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowFileSaveOnHost", m.GetApplicationGuardAllowFileSaveOnHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPersistence", m.GetApplicationGuardAllowPersistence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToLocalPrinters", m.GetApplicationGuardAllowPrintToLocalPrinters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToNetworkPrinters", m.GetApplicationGuardAllowPrintToNetworkPrinters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToPDF", m.GetApplicationGuardAllowPrintToPDF())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToXPS", m.GetApplicationGuardAllowPrintToXPS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowVirtualGPU", m.GetApplicationGuardAllowVirtualGPU())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardBlockClipboardSharing() != nil {
        cast := (*m.GetApplicationGuardBlockClipboardSharing()).String()
        err = writer.WriteStringValue("applicationGuardBlockClipboardSharing", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardBlockFileTransfer() != nil {
        cast := (*m.GetApplicationGuardBlockFileTransfer()).String()
        err = writer.WriteStringValue("applicationGuardBlockFileTransfer", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardBlockNonEnterpriseContent", m.GetApplicationGuardBlockNonEnterpriseContent())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardCertificateThumbprints() != nil {
        err = writer.WriteCollectionOfStringValues("applicationGuardCertificateThumbprints", m.GetApplicationGuardCertificateThumbprints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardEnabled", m.GetApplicationGuardEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardEnabledOptions() != nil {
        cast := (*m.GetApplicationGuardEnabledOptions()).String()
        err = writer.WriteStringValue("applicationGuardEnabledOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardForceAuditing", m.GetApplicationGuardForceAuditing())
        if err != nil {
            return err
        }
    }
    if m.GetAppLockerApplicationControl() != nil {
        cast := (*m.GetAppLockerApplicationControl()).String()
        err = writer.WriteStringValue("appLockerApplicationControl", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerAllowStandardUserEncryption", m.GetBitLockerAllowStandardUserEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerDisableWarningForOtherDiskEncryption", m.GetBitLockerDisableWarningForOtherDiskEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerEnableStorageCardEncryptionOnMobile", m.GetBitLockerEnableStorageCardEncryptionOnMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerEncryptDevice", m.GetBitLockerEncryptDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bitLockerFixedDrivePolicy", m.GetBitLockerFixedDrivePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetBitLockerRecoveryPasswordRotation() != nil {
        cast := (*m.GetBitLockerRecoveryPasswordRotation()).String()
        err = writer.WriteStringValue("bitLockerRecoveryPasswordRotation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bitLockerRemovableDrivePolicy", m.GetBitLockerRemovableDrivePolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bitLockerSystemDrivePolicy", m.GetBitLockerSystemDrivePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAdditionalGuardedFolders() != nil {
        err = writer.WriteCollectionOfStringValues("defenderAdditionalGuardedFolders", m.GetDefenderAdditionalGuardedFolders())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAdobeReaderLaunchChildProcess() != nil {
        cast := (*m.GetDefenderAdobeReaderLaunchChildProcess()).String()
        err = writer.WriteStringValue("defenderAdobeReaderLaunchChildProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAdvancedRansomewareProtectionType() != nil {
        cast := (*m.GetDefenderAdvancedRansomewareProtectionType()).String()
        err = writer.WriteStringValue("defenderAdvancedRansomewareProtectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowBehaviorMonitoring", m.GetDefenderAllowBehaviorMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowCloudProtection", m.GetDefenderAllowCloudProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowEndUserAccess", m.GetDefenderAllowEndUserAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowIntrusionPreventionSystem", m.GetDefenderAllowIntrusionPreventionSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowOnAccessProtection", m.GetDefenderAllowOnAccessProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowRealTimeMonitoring", m.GetDefenderAllowRealTimeMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanArchiveFiles", m.GetDefenderAllowScanArchiveFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanDownloads", m.GetDefenderAllowScanDownloads())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanNetworkFiles", m.GetDefenderAllowScanNetworkFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanRemovableDrivesDuringFullScan", m.GetDefenderAllowScanRemovableDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanScriptsLoadedInInternetExplorer", m.GetDefenderAllowScanScriptsLoadedInInternetExplorer())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAttackSurfaceReductionExcludedPaths() != nil {
        err = writer.WriteCollectionOfStringValues("defenderAttackSurfaceReductionExcludedPaths", m.GetDefenderAttackSurfaceReductionExcludedPaths())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderBlockEndUserAccess", m.GetDefenderBlockEndUserAccess())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderBlockPersistenceThroughWmiType() != nil {
        cast := (*m.GetDefenderBlockPersistenceThroughWmiType()).String()
        err = writer.WriteStringValue("defenderBlockPersistenceThroughWmiType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderCheckForSignaturesBeforeRunningScan", m.GetDefenderCheckForSignaturesBeforeRunningScan())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderCloudBlockLevel() != nil {
        cast := (*m.GetDefenderCloudBlockLevel()).String()
        err = writer.WriteStringValue("defenderCloudBlockLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderCloudExtendedTimeoutInSeconds", m.GetDefenderCloudExtendedTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderDaysBeforeDeletingQuarantinedMalware", m.GetDefenderDaysBeforeDeletingQuarantinedMalware())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defenderDetectedMalwareActions", m.GetDefenderDetectedMalwareActions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableBehaviorMonitoring", m.GetDefenderDisableBehaviorMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableCatchupFullScan", m.GetDefenderDisableCatchupFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableCatchupQuickScan", m.GetDefenderDisableCatchupQuickScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableCloudProtection", m.GetDefenderDisableCloudProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableIntrusionPreventionSystem", m.GetDefenderDisableIntrusionPreventionSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableOnAccessProtection", m.GetDefenderDisableOnAccessProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableRealTimeMonitoring", m.GetDefenderDisableRealTimeMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanArchiveFiles", m.GetDefenderDisableScanArchiveFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanDownloads", m.GetDefenderDisableScanDownloads())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanNetworkFiles", m.GetDefenderDisableScanNetworkFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanRemovableDrivesDuringFullScan", m.GetDefenderDisableScanRemovableDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanScriptsLoadedInInternetExplorer", m.GetDefenderDisableScanScriptsLoadedInInternetExplorer())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderEmailContentExecution() != nil {
        cast := (*m.GetDefenderEmailContentExecution()).String()
        err = writer.WriteStringValue("defenderEmailContentExecution", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderEmailContentExecutionType() != nil {
        cast := (*m.GetDefenderEmailContentExecutionType()).String()
        err = writer.WriteStringValue("defenderEmailContentExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderEnableLowCpuPriority", m.GetDefenderEnableLowCpuPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderEnableScanIncomingMail", m.GetDefenderEnableScanIncomingMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderEnableScanMappedNetworkDrivesDuringFullScan", m.GetDefenderEnableScanMappedNetworkDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("defenderExploitProtectionXml", m.GetDefenderExploitProtectionXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderExploitProtectionXmlFileName", m.GetDefenderExploitProtectionXmlFileName())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderFileExtensionsToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderFileExtensionsToExclude", m.GetDefenderFileExtensionsToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderFilesAndFoldersToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderFilesAndFoldersToExclude", m.GetDefenderFilesAndFoldersToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderGuardedFoldersAllowedAppPaths() != nil {
        err = writer.WriteCollectionOfStringValues("defenderGuardedFoldersAllowedAppPaths", m.GetDefenderGuardedFoldersAllowedAppPaths())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderGuardMyFoldersType() != nil {
        cast := (*m.GetDefenderGuardMyFoldersType()).String()
        err = writer.WriteStringValue("defenderGuardMyFoldersType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderNetworkProtectionType() != nil {
        cast := (*m.GetDefenderNetworkProtectionType()).String()
        err = writer.WriteStringValue("defenderNetworkProtectionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsExecutableContentCreationOrLaunch() != nil {
        cast := (*m.GetDefenderOfficeAppsExecutableContentCreationOrLaunch()).String()
        err = writer.WriteStringValue("defenderOfficeAppsExecutableContentCreationOrLaunch", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsExecutableContentCreationOrLaunchType() != nil {
        cast := (*m.GetDefenderOfficeAppsExecutableContentCreationOrLaunchType()).String()
        err = writer.WriteStringValue("defenderOfficeAppsExecutableContentCreationOrLaunchType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsLaunchChildProcess() != nil {
        cast := (*m.GetDefenderOfficeAppsLaunchChildProcess()).String()
        err = writer.WriteStringValue("defenderOfficeAppsLaunchChildProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsLaunchChildProcessType() != nil {
        cast := (*m.GetDefenderOfficeAppsLaunchChildProcessType()).String()
        err = writer.WriteStringValue("defenderOfficeAppsLaunchChildProcessType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsOtherProcessInjection() != nil {
        cast := (*m.GetDefenderOfficeAppsOtherProcessInjection()).String()
        err = writer.WriteStringValue("defenderOfficeAppsOtherProcessInjection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsOtherProcessInjectionType() != nil {
        cast := (*m.GetDefenderOfficeAppsOtherProcessInjectionType()).String()
        err = writer.WriteStringValue("defenderOfficeAppsOtherProcessInjectionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeCommunicationAppsLaunchChildProcess() != nil {
        cast := (*m.GetDefenderOfficeCommunicationAppsLaunchChildProcess()).String()
        err = writer.WriteStringValue("defenderOfficeCommunicationAppsLaunchChildProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeMacroCodeAllowWin32Imports() != nil {
        cast := (*m.GetDefenderOfficeMacroCodeAllowWin32Imports()).String()
        err = writer.WriteStringValue("defenderOfficeMacroCodeAllowWin32Imports", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeMacroCodeAllowWin32ImportsType() != nil {
        cast := (*m.GetDefenderOfficeMacroCodeAllowWin32ImportsType()).String()
        err = writer.WriteStringValue("defenderOfficeMacroCodeAllowWin32ImportsType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderPotentiallyUnwantedAppAction() != nil {
        cast := (*m.GetDefenderPotentiallyUnwantedAppAction()).String()
        err = writer.WriteStringValue("defenderPotentiallyUnwantedAppAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderPreventCredentialStealingType() != nil {
        cast := (*m.GetDefenderPreventCredentialStealingType()).String()
        err = writer.WriteStringValue("defenderPreventCredentialStealingType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderProcessCreation() != nil {
        cast := (*m.GetDefenderProcessCreation()).String()
        err = writer.WriteStringValue("defenderProcessCreation", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderProcessCreationType() != nil {
        cast := (*m.GetDefenderProcessCreationType()).String()
        err = writer.WriteStringValue("defenderProcessCreationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderProcessesToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderProcessesToExclude", m.GetDefenderProcessesToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScanDirection() != nil {
        cast := (*m.GetDefenderScanDirection()).String()
        err = writer.WriteStringValue("defenderScanDirection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderScanMaxCpuPercentage", m.GetDefenderScanMaxCpuPercentage())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScanType() != nil {
        cast := (*m.GetDefenderScanType()).String()
        err = writer.WriteStringValue("defenderScanType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("defenderScheduledQuickScanTime", m.GetDefenderScheduledQuickScanTime())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScheduledScanDay() != nil {
        cast := (*m.GetDefenderScheduledScanDay()).String()
        err = writer.WriteStringValue("defenderScheduledScanDay", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("defenderScheduledScanTime", m.GetDefenderScheduledScanTime())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptDownloadedPayloadExecution() != nil {
        cast := (*m.GetDefenderScriptDownloadedPayloadExecution()).String()
        err = writer.WriteStringValue("defenderScriptDownloadedPayloadExecution", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptDownloadedPayloadExecutionType() != nil {
        cast := (*m.GetDefenderScriptDownloadedPayloadExecutionType()).String()
        err = writer.WriteStringValue("defenderScriptDownloadedPayloadExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptObfuscatedMacroCode() != nil {
        cast := (*m.GetDefenderScriptObfuscatedMacroCode()).String()
        err = writer.WriteStringValue("defenderScriptObfuscatedMacroCode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptObfuscatedMacroCodeType() != nil {
        cast := (*m.GetDefenderScriptObfuscatedMacroCodeType()).String()
        err = writer.WriteStringValue("defenderScriptObfuscatedMacroCodeType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterBlockExploitProtectionOverride", m.GetDefenderSecurityCenterBlockExploitProtectionOverride())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableAccountUI", m.GetDefenderSecurityCenterDisableAccountUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableAppBrowserUI", m.GetDefenderSecurityCenterDisableAppBrowserUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableClearTpmUI", m.GetDefenderSecurityCenterDisableClearTpmUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableFamilyUI", m.GetDefenderSecurityCenterDisableFamilyUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableHardwareUI", m.GetDefenderSecurityCenterDisableHardwareUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableHealthUI", m.GetDefenderSecurityCenterDisableHealthUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableNetworkUI", m.GetDefenderSecurityCenterDisableNetworkUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableNotificationAreaUI", m.GetDefenderSecurityCenterDisableNotificationAreaUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableRansomwareUI", m.GetDefenderSecurityCenterDisableRansomwareUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableSecureBootUI", m.GetDefenderSecurityCenterDisableSecureBootUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableTroubleshootingUI", m.GetDefenderSecurityCenterDisableTroubleshootingUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableVirusUI", m.GetDefenderSecurityCenterDisableVirusUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI", m.GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterHelpEmail", m.GetDefenderSecurityCenterHelpEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterHelpPhone", m.GetDefenderSecurityCenterHelpPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterHelpURL", m.GetDefenderSecurityCenterHelpURL())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderSecurityCenterITContactDisplay() != nil {
        cast := (*m.GetDefenderSecurityCenterITContactDisplay()).String()
        err = writer.WriteStringValue("defenderSecurityCenterITContactDisplay", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderSecurityCenterNotificationsFromApp() != nil {
        cast := (*m.GetDefenderSecurityCenterNotificationsFromApp()).String()
        err = writer.WriteStringValue("defenderSecurityCenterNotificationsFromApp", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterOrganizationDisplayName", m.GetDefenderSecurityCenterOrganizationDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderSignatureUpdateIntervalInHours", m.GetDefenderSignatureUpdateIntervalInHours())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderSubmitSamplesConsentType() != nil {
        cast := (*m.GetDefenderSubmitSamplesConsentType()).String()
        err = writer.WriteStringValue("defenderSubmitSamplesConsentType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedExecutable() != nil {
        cast := (*m.GetDefenderUntrustedExecutable()).String()
        err = writer.WriteStringValue("defenderUntrustedExecutable", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedExecutableType() != nil {
        cast := (*m.GetDefenderUntrustedExecutableType()).String()
        err = writer.WriteStringValue("defenderUntrustedExecutableType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedUSBProcess() != nil {
        cast := (*m.GetDefenderUntrustedUSBProcess()).String()
        err = writer.WriteStringValue("defenderUntrustedUSBProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedUSBProcessType() != nil {
        cast := (*m.GetDefenderUntrustedUSBProcessType()).String()
        err = writer.WriteStringValue("defenderUntrustedUSBProcessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceGuardEnableSecureBootWithDMA", m.GetDeviceGuardEnableSecureBootWithDMA())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceGuardEnableVirtualizationBasedSecurity", m.GetDeviceGuardEnableVirtualizationBasedSecurity())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardLaunchSystemGuard() != nil {
        cast := (*m.GetDeviceGuardLaunchSystemGuard()).String()
        err = writer.WriteStringValue("deviceGuardLaunchSystemGuard", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings() != nil {
        cast := (*m.GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings()).String()
        err = writer.WriteStringValue("deviceGuardLocalSystemAuthorityCredentialGuardSettings", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardSecureBootWithDMA() != nil {
        cast := (*m.GetDeviceGuardSecureBootWithDMA()).String()
        err = writer.WriteStringValue("deviceGuardSecureBootWithDMA", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDmaGuardDeviceEnumerationPolicy() != nil {
        cast := (*m.GetDmaGuardDeviceEnumerationPolicy()).String()
        err = writer.WriteStringValue("dmaGuardDeviceEnumerationPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallBlockStatefulFTP", m.GetFirewallBlockStatefulFTP())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallCertificateRevocationListCheckMethod() != nil {
        cast := (*m.GetFirewallCertificateRevocationListCheckMethod()).String()
        err = writer.WriteStringValue("firewallCertificateRevocationListCheckMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("firewallIdleTimeoutForSecurityAssociationInSeconds", m.GetFirewallIdleTimeoutForSecurityAssociationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowDHCP", m.GetFirewallIPSecExemptionsAllowDHCP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowICMP", m.GetFirewallIPSecExemptionsAllowICMP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowNeighborDiscovery", m.GetFirewallIPSecExemptionsAllowNeighborDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowRouterDiscovery", m.GetFirewallIPSecExemptionsAllowRouterDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsNone", m.GetFirewallIPSecExemptionsNone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallMergeKeyingModuleSettings", m.GetFirewallMergeKeyingModuleSettings())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallPacketQueueingMethod() != nil {
        cast := (*m.GetFirewallPacketQueueingMethod()).String()
        err = writer.WriteStringValue("firewallPacketQueueingMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFirewallPreSharedKeyEncodingMethod() != nil {
        cast := (*m.GetFirewallPreSharedKeyEncodingMethod()).String()
        err = writer.WriteStringValue("firewallPreSharedKeyEncodingMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfileDomain", m.GetFirewallProfileDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfilePrivate", m.GetFirewallProfilePrivate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfilePublic", m.GetFirewallProfilePublic())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFirewallRules()))
        for i, v := range m.GetFirewallRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("firewallRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLanManagerAuthenticationLevel() != nil {
        cast := (*m.GetLanManagerAuthenticationLevel()).String()
        err = writer.WriteStringValue("lanManagerAuthenticationLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lanManagerWorkstationDisableInsecureGuestLogons", m.GetLanManagerWorkstationDisableInsecureGuestLogons())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsAdministratorAccountName", m.GetLocalSecurityOptionsAdministratorAccountName())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsAdministratorElevationPromptBehavior() != nil {
        cast := (*m.GetLocalSecurityOptionsAdministratorElevationPromptBehavior()).String()
        err = writer.WriteStringValue("localSecurityOptionsAdministratorElevationPromptBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares", m.GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowPKU2UAuthenticationRequests", m.GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager", m.GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool", m.GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn", m.GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowUIAccessApplicationElevation", m.GetLocalSecurityOptionsAllowUIAccessApplicationElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowUIAccessApplicationsForSecureLocations", m.GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowUndockWithoutHavingToLogon", m.GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockMicrosoftAccounts", m.GetLocalSecurityOptionsBlockMicrosoftAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockRemoteLogonWithBlankPassword", m.GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockRemoteOpticalDriveAccess", m.GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockUsersInstallingPrinterDrivers", m.GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsClearVirtualMemoryPageFile", m.GetLocalSecurityOptionsClearVirtualMemoryPageFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsClientDigitallySignCommunicationsAlways", m.GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers", m.GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation", m.GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableAdministratorAccount", m.GetLocalSecurityOptionsDisableAdministratorAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees", m.GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableGuestAccount", m.GetLocalSecurityOptionsDisableGuestAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableServerDigitallySignCommunicationsAlways", m.GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees", m.GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts", m.GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDoNotRequireCtrlAltDel", m.GetLocalSecurityOptionsDoNotRequireCtrlAltDel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange", m.GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser() != nil {
        cast := (*m.GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser()).String()
        err = writer.WriteStringValue("localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsGuestAccountName", m.GetLocalSecurityOptionsGuestAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsHideLastSignedInUser", m.GetLocalSecurityOptionsHideLastSignedInUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsHideUsernameAtSignIn", m.GetLocalSecurityOptionsHideUsernameAtSignIn())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsInformationDisplayedOnLockScreen() != nil {
        cast := (*m.GetLocalSecurityOptionsInformationDisplayedOnLockScreen()).String()
        err = writer.WriteStringValue("localSecurityOptionsInformationDisplayedOnLockScreen", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsInformationShownOnLockScreen() != nil {
        cast := (*m.GetLocalSecurityOptionsInformationShownOnLockScreen()).String()
        err = writer.WriteStringValue("localSecurityOptionsInformationShownOnLockScreen", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsLogOnMessageText", m.GetLocalSecurityOptionsLogOnMessageText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsLogOnMessageTitle", m.GetLocalSecurityOptionsLogOnMessageTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("localSecurityOptionsMachineInactivityLimit", m.GetLocalSecurityOptionsMachineInactivityLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("localSecurityOptionsMachineInactivityLimitInMinutes", m.GetLocalSecurityOptionsMachineInactivityLimitInMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients() != nil {
        cast := (*m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients()).String()
        err = writer.WriteStringValue("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers() != nil {
        cast := (*m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers()).String()
        err = writer.WriteStringValue("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsOnlyElevateSignedExecutables", m.GetLocalSecurityOptionsOnlyElevateSignedExecutables())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares", m.GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsSmartCardRemovalBehavior() != nil {
        cast := (*m.GetLocalSecurityOptionsSmartCardRemovalBehavior()).String()
        err = writer.WriteStringValue("localSecurityOptionsSmartCardRemovalBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsStandardUserElevationPromptBehavior() != nil {
        cast := (*m.GetLocalSecurityOptionsStandardUserElevationPromptBehavior()).String()
        err = writer.WriteStringValue("localSecurityOptionsStandardUserElevationPromptBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation", m.GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsUseAdminApprovalMode", m.GetLocalSecurityOptionsUseAdminApprovalMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsUseAdminApprovalModeForAdministrators", m.GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations", m.GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenBlockOverrideForFiles", m.GetSmartScreenBlockOverrideForFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenEnableInShell", m.GetSmartScreenEnableInShell())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsAccessCredentialManagerAsTrustedCaller", m.GetUserRightsAccessCredentialManagerAsTrustedCaller())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsActAsPartOfTheOperatingSystem", m.GetUserRightsActAsPartOfTheOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsAllowAccessFromNetwork", m.GetUserRightsAllowAccessFromNetwork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsBackupData", m.GetUserRightsBackupData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsBlockAccessFromNetwork", m.GetUserRightsBlockAccessFromNetwork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsChangeSystemTime", m.GetUserRightsChangeSystemTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreateGlobalObjects", m.GetUserRightsCreateGlobalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreatePageFile", m.GetUserRightsCreatePageFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreatePermanentSharedObjects", m.GetUserRightsCreatePermanentSharedObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreateSymbolicLinks", m.GetUserRightsCreateSymbolicLinks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreateToken", m.GetUserRightsCreateToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsDebugPrograms", m.GetUserRightsDebugPrograms())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsDelegation", m.GetUserRightsDelegation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsDenyLocalLogOn", m.GetUserRightsDenyLocalLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsGenerateSecurityAudits", m.GetUserRightsGenerateSecurityAudits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsImpersonateClient", m.GetUserRightsImpersonateClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsIncreaseSchedulingPriority", m.GetUserRightsIncreaseSchedulingPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsLoadUnloadDrivers", m.GetUserRightsLoadUnloadDrivers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsLocalLogOn", m.GetUserRightsLocalLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsLockMemory", m.GetUserRightsLockMemory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsManageAuditingAndSecurityLogs", m.GetUserRightsManageAuditingAndSecurityLogs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsManageVolumes", m.GetUserRightsManageVolumes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsModifyFirmwareEnvironment", m.GetUserRightsModifyFirmwareEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsModifyObjectLabels", m.GetUserRightsModifyObjectLabels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsProfileSingleProcess", m.GetUserRightsProfileSingleProcess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsRemoteDesktopServicesLogOn", m.GetUserRightsRemoteDesktopServicesLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsRemoteShutdown", m.GetUserRightsRemoteShutdown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsRestoreData", m.GetUserRightsRestoreData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsTakeOwnership", m.GetUserRightsTakeOwnership())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDefenderTamperProtection() != nil {
        cast := (*m.GetWindowsDefenderTamperProtection()).String()
        err = writer.WriteStringValue("windowsDefenderTamperProtection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesAccessoryManagementServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesAccessoryManagementServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesAccessoryManagementServiceStartupMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("xboxServicesEnableXboxGameSaveTask", m.GetXboxServicesEnableXboxGameSaveTask())
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesLiveAuthManagerServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesLiveAuthManagerServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesLiveAuthManagerServiceStartupMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesLiveGameSaveServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesLiveGameSaveServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesLiveGameSaveServiceStartupMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesLiveNetworkingServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesLiveNetworkingServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesLiveNetworkingServiceStartupMode", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationGuardAllowCameraMicrophoneRedirection sets the applicationGuardAllowCameraMicrophoneRedirection property value. Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowCameraMicrophoneRedirection(value *bool)() {
    m.applicationGuardAllowCameraMicrophoneRedirection = value
}
// SetApplicationGuardAllowFileSaveOnHost sets the applicationGuardAllowFileSaveOnHost property value. Allow users to download files from Edge in the application guard container and save them on the host file system
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowFileSaveOnHost(value *bool)() {
    m.applicationGuardAllowFileSaveOnHost = value
}
// SetApplicationGuardAllowPersistence sets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPersistence(value *bool)() {
    m.applicationGuardAllowPersistence = value
}
// SetApplicationGuardAllowPrintToLocalPrinters sets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToLocalPrinters(value *bool)() {
    m.applicationGuardAllowPrintToLocalPrinters = value
}
// SetApplicationGuardAllowPrintToNetworkPrinters sets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToNetworkPrinters(value *bool)() {
    m.applicationGuardAllowPrintToNetworkPrinters = value
}
// SetApplicationGuardAllowPrintToPDF sets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToPDF(value *bool)() {
    m.applicationGuardAllowPrintToPDF = value
}
// SetApplicationGuardAllowPrintToXPS sets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToXPS(value *bool)() {
    m.applicationGuardAllowPrintToXPS = value
}
// SetApplicationGuardAllowVirtualGPU sets the applicationGuardAllowVirtualGPU property value. Allow application guard to use virtual GPU
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowVirtualGPU(value *bool)() {
    m.applicationGuardAllowVirtualGPU = value
}
// SetApplicationGuardBlockClipboardSharing sets the applicationGuardBlockClipboardSharing property value. Possible values for applicationGuardBlockClipboardSharingType
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockClipboardSharing(value *ApplicationGuardBlockClipboardSharingType)() {
    m.applicationGuardBlockClipboardSharing = value
}
// SetApplicationGuardBlockFileTransfer sets the applicationGuardBlockFileTransfer property value. Possible values for applicationGuardBlockFileTransfer
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockFileTransfer(value *ApplicationGuardBlockFileTransferType)() {
    m.applicationGuardBlockFileTransfer = value
}
// SetApplicationGuardBlockNonEnterpriseContent sets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockNonEnterpriseContent(value *bool)() {
    m.applicationGuardBlockNonEnterpriseContent = value
}
// SetApplicationGuardCertificateThumbprints sets the applicationGuardCertificateThumbprints property value. Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardCertificateThumbprints(value []string)() {
    m.applicationGuardCertificateThumbprints = value
}
// SetApplicationGuardEnabled sets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardEnabled(value *bool)() {
    m.applicationGuardEnabled = value
}
// SetApplicationGuardEnabledOptions sets the applicationGuardEnabledOptions property value. Possible values for ApplicationGuardEnabledOptions
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardEnabledOptions(value *ApplicationGuardEnabledOptions)() {
    m.applicationGuardEnabledOptions = value
}
// SetApplicationGuardForceAuditing sets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardForceAuditing(value *bool)() {
    m.applicationGuardForceAuditing = value
}
// SetAppLockerApplicationControl sets the appLockerApplicationControl property value. Possible values of AppLocker Application Control Types
func (m *Windows10EndpointProtectionConfiguration) SetAppLockerApplicationControl(value *AppLockerApplicationControlType)() {
    m.appLockerApplicationControl = value
}
// SetBitLockerAllowStandardUserEncryption sets the bitLockerAllowStandardUserEncryption property value. Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerAllowStandardUserEncryption(value *bool)() {
    m.bitLockerAllowStandardUserEncryption = value
}
// SetBitLockerDisableWarningForOtherDiskEncryption sets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerDisableWarningForOtherDiskEncryption(value *bool)() {
    m.bitLockerDisableWarningForOtherDiskEncryption = value
}
// SetBitLockerEnableStorageCardEncryptionOnMobile sets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEnableStorageCardEncryptionOnMobile(value *bool)() {
    m.bitLockerEnableStorageCardEncryptionOnMobile = value
}
// SetBitLockerEncryptDevice sets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEncryptDevice(value *bool)() {
    m.bitLockerEncryptDevice = value
}
// SetBitLockerFixedDrivePolicy sets the bitLockerFixedDrivePolicy property value. BitLocker Fixed Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerFixedDrivePolicy(value BitLockerFixedDrivePolicyable)() {
    m.bitLockerFixedDrivePolicy = value
}
// SetBitLockerRecoveryPasswordRotation sets the bitLockerRecoveryPasswordRotation property value. BitLocker recovery password rotation type
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerRecoveryPasswordRotation(value *BitLockerRecoveryPasswordRotationType)() {
    m.bitLockerRecoveryPasswordRotation = value
}
// SetBitLockerRemovableDrivePolicy sets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerRemovableDrivePolicy(value BitLockerRemovableDrivePolicyable)() {
    m.bitLockerRemovableDrivePolicy = value
}
// SetBitLockerSystemDrivePolicy sets the bitLockerSystemDrivePolicy property value. BitLocker System Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerSystemDrivePolicy(value BitLockerSystemDrivePolicyable)() {
    m.bitLockerSystemDrivePolicy = value
}
// SetDefenderAdditionalGuardedFolders sets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdditionalGuardedFolders(value []string)() {
    m.defenderAdditionalGuardedFolders = value
}
// SetDefenderAdobeReaderLaunchChildProcess sets the defenderAdobeReaderLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdobeReaderLaunchChildProcess(value *DefenderProtectionType)() {
    m.defenderAdobeReaderLaunchChildProcess = value
}
// SetDefenderAdvancedRansomewareProtectionType sets the defenderAdvancedRansomewareProtectionType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdvancedRansomewareProtectionType(value *DefenderProtectionType)() {
    m.defenderAdvancedRansomewareProtectionType = value
}
// SetDefenderAllowBehaviorMonitoring sets the defenderAllowBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowBehaviorMonitoring(value *bool)() {
    m.defenderAllowBehaviorMonitoring = value
}
// SetDefenderAllowCloudProtection sets the defenderAllowCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowCloudProtection(value *bool)() {
    m.defenderAllowCloudProtection = value
}
// SetDefenderAllowEndUserAccess sets the defenderAllowEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowEndUserAccess(value *bool)() {
    m.defenderAllowEndUserAccess = value
}
// SetDefenderAllowIntrusionPreventionSystem sets the defenderAllowIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowIntrusionPreventionSystem(value *bool)() {
    m.defenderAllowIntrusionPreventionSystem = value
}
// SetDefenderAllowOnAccessProtection sets the defenderAllowOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowOnAccessProtection(value *bool)() {
    m.defenderAllowOnAccessProtection = value
}
// SetDefenderAllowRealTimeMonitoring sets the defenderAllowRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowRealTimeMonitoring(value *bool)() {
    m.defenderAllowRealTimeMonitoring = value
}
// SetDefenderAllowScanArchiveFiles sets the defenderAllowScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanArchiveFiles(value *bool)() {
    m.defenderAllowScanArchiveFiles = value
}
// SetDefenderAllowScanDownloads sets the defenderAllowScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanDownloads(value *bool)() {
    m.defenderAllowScanDownloads = value
}
// SetDefenderAllowScanNetworkFiles sets the defenderAllowScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanNetworkFiles(value *bool)() {
    m.defenderAllowScanNetworkFiles = value
}
// SetDefenderAllowScanRemovableDrivesDuringFullScan sets the defenderAllowScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanRemovableDrivesDuringFullScan(value *bool)() {
    m.defenderAllowScanRemovableDrivesDuringFullScan = value
}
// SetDefenderAllowScanScriptsLoadedInInternetExplorer sets the defenderAllowScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanScriptsLoadedInInternetExplorer(value *bool)() {
    m.defenderAllowScanScriptsLoadedInInternetExplorer = value
}
// SetDefenderAttackSurfaceReductionExcludedPaths sets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAttackSurfaceReductionExcludedPaths(value []string)() {
    m.defenderAttackSurfaceReductionExcludedPaths = value
}
// SetDefenderBlockEndUserAccess sets the defenderBlockEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderBlockEndUserAccess(value *bool)() {
    m.defenderBlockEndUserAccess = value
}
// SetDefenderBlockPersistenceThroughWmiType sets the defenderBlockPersistenceThroughWmiType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderBlockPersistenceThroughWmiType(value *DefenderAttackSurfaceType)() {
    m.defenderBlockPersistenceThroughWmiType = value
}
// SetDefenderCheckForSignaturesBeforeRunningScan sets the defenderCheckForSignaturesBeforeRunningScan property value. This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCheckForSignaturesBeforeRunningScan(value *bool)() {
    m.defenderCheckForSignaturesBeforeRunningScan = value
}
// SetDefenderCloudBlockLevel sets the defenderCloudBlockLevel property value. Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCloudBlockLevel(value *DefenderCloudBlockLevelType)() {
    m.defenderCloudBlockLevel = value
}
// SetDefenderCloudExtendedTimeoutInSeconds sets the defenderCloudExtendedTimeoutInSeconds property value. Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCloudExtendedTimeoutInSeconds(value *int32)() {
    m.defenderCloudExtendedTimeoutInSeconds = value
}
// SetDefenderDaysBeforeDeletingQuarantinedMalware sets the defenderDaysBeforeDeletingQuarantinedMalware property value. Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalware(value *int32)() {
    m.defenderDaysBeforeDeletingQuarantinedMalware = value
}
// SetDefenderDetectedMalwareActions sets the defenderDetectedMalwareActions property value. Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDetectedMalwareActions(value DefenderDetectedMalwareActionsable)() {
    m.defenderDetectedMalwareActions = value
}
// SetDefenderDisableBehaviorMonitoring sets the defenderDisableBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableBehaviorMonitoring(value *bool)() {
    m.defenderDisableBehaviorMonitoring = value
}
// SetDefenderDisableCatchupFullScan sets the defenderDisableCatchupFullScan property value. This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCatchupFullScan(value *bool)() {
    m.defenderDisableCatchupFullScan = value
}
// SetDefenderDisableCatchupQuickScan sets the defenderDisableCatchupQuickScan property value. This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCatchupQuickScan(value *bool)() {
    m.defenderDisableCatchupQuickScan = value
}
// SetDefenderDisableCloudProtection sets the defenderDisableCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCloudProtection(value *bool)() {
    m.defenderDisableCloudProtection = value
}
// SetDefenderDisableIntrusionPreventionSystem sets the defenderDisableIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableIntrusionPreventionSystem(value *bool)() {
    m.defenderDisableIntrusionPreventionSystem = value
}
// SetDefenderDisableOnAccessProtection sets the defenderDisableOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableOnAccessProtection(value *bool)() {
    m.defenderDisableOnAccessProtection = value
}
// SetDefenderDisableRealTimeMonitoring sets the defenderDisableRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableRealTimeMonitoring(value *bool)() {
    m.defenderDisableRealTimeMonitoring = value
}
// SetDefenderDisableScanArchiveFiles sets the defenderDisableScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanArchiveFiles(value *bool)() {
    m.defenderDisableScanArchiveFiles = value
}
// SetDefenderDisableScanDownloads sets the defenderDisableScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanDownloads(value *bool)() {
    m.defenderDisableScanDownloads = value
}
// SetDefenderDisableScanNetworkFiles sets the defenderDisableScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanNetworkFiles(value *bool)() {
    m.defenderDisableScanNetworkFiles = value
}
// SetDefenderDisableScanRemovableDrivesDuringFullScan sets the defenderDisableScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanRemovableDrivesDuringFullScan(value *bool)() {
    m.defenderDisableScanRemovableDrivesDuringFullScan = value
}
// SetDefenderDisableScanScriptsLoadedInInternetExplorer sets the defenderDisableScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanScriptsLoadedInInternetExplorer(value *bool)() {
    m.defenderDisableScanScriptsLoadedInInternetExplorer = value
}
// SetDefenderEmailContentExecution sets the defenderEmailContentExecution property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEmailContentExecution(value *DefenderProtectionType)() {
    m.defenderEmailContentExecution = value
}
// SetDefenderEmailContentExecutionType sets the defenderEmailContentExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEmailContentExecutionType(value *DefenderAttackSurfaceType)() {
    m.defenderEmailContentExecutionType = value
}
// SetDefenderEnableLowCpuPriority sets the defenderEnableLowCpuPriority property value. This policy setting allows you to enable or disable low CPU priority for scheduled scans.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableLowCpuPriority(value *bool)() {
    m.defenderEnableLowCpuPriority = value
}
// SetDefenderEnableScanIncomingMail sets the defenderEnableScanIncomingMail property value. Allows or disallows scanning of email.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableScanIncomingMail(value *bool)() {
    m.defenderEnableScanIncomingMail = value
}
// SetDefenderEnableScanMappedNetworkDrivesDuringFullScan sets the defenderEnableScanMappedNetworkDrivesDuringFullScan property value. Allows or disallows a full scan of mapped network drives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableScanMappedNetworkDrivesDuringFullScan(value *bool)() {
    m.defenderEnableScanMappedNetworkDrivesDuringFullScan = value
}
// SetDefenderExploitProtectionXml sets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXml(value []byte)() {
    m.defenderExploitProtectionXml = value
}
// SetDefenderExploitProtectionXmlFileName sets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXmlFileName(value *string)() {
    m.defenderExploitProtectionXmlFileName = value
}
// SetDefenderFileExtensionsToExclude sets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderFileExtensionsToExclude(value []string)() {
    m.defenderFileExtensionsToExclude = value
}
// SetDefenderFilesAndFoldersToExclude sets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderFilesAndFoldersToExclude(value []string)() {
    m.defenderFilesAndFoldersToExclude = value
}
// SetDefenderGuardedFoldersAllowedAppPaths sets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderGuardedFoldersAllowedAppPaths(value []string)() {
    m.defenderGuardedFoldersAllowedAppPaths = value
}
// SetDefenderGuardMyFoldersType sets the defenderGuardMyFoldersType property value. Possible values of Folder Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderGuardMyFoldersType(value *FolderProtectionType)() {
    m.defenderGuardMyFoldersType = value
}
// SetDefenderNetworkProtectionType sets the defenderNetworkProtectionType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderNetworkProtectionType(value *DefenderProtectionType)() {
    m.defenderNetworkProtectionType = value
}
// SetDefenderOfficeAppsExecutableContentCreationOrLaunch sets the defenderOfficeAppsExecutableContentCreationOrLaunch property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsExecutableContentCreationOrLaunch(value *DefenderProtectionType)() {
    m.defenderOfficeAppsExecutableContentCreationOrLaunch = value
}
// SetDefenderOfficeAppsExecutableContentCreationOrLaunchType sets the defenderOfficeAppsExecutableContentCreationOrLaunchType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsExecutableContentCreationOrLaunchType(value *DefenderAttackSurfaceType)() {
    m.defenderOfficeAppsExecutableContentCreationOrLaunchType = value
}
// SetDefenderOfficeAppsLaunchChildProcess sets the defenderOfficeAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsLaunchChildProcess(value *DefenderProtectionType)() {
    m.defenderOfficeAppsLaunchChildProcess = value
}
// SetDefenderOfficeAppsLaunchChildProcessType sets the defenderOfficeAppsLaunchChildProcessType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsLaunchChildProcessType(value *DefenderAttackSurfaceType)() {
    m.defenderOfficeAppsLaunchChildProcessType = value
}
// SetDefenderOfficeAppsOtherProcessInjection sets the defenderOfficeAppsOtherProcessInjection property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsOtherProcessInjection(value *DefenderProtectionType)() {
    m.defenderOfficeAppsOtherProcessInjection = value
}
// SetDefenderOfficeAppsOtherProcessInjectionType sets the defenderOfficeAppsOtherProcessInjectionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsOtherProcessInjectionType(value *DefenderAttackSurfaceType)() {
    m.defenderOfficeAppsOtherProcessInjectionType = value
}
// SetDefenderOfficeCommunicationAppsLaunchChildProcess sets the defenderOfficeCommunicationAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeCommunicationAppsLaunchChildProcess(value *DefenderProtectionType)() {
    m.defenderOfficeCommunicationAppsLaunchChildProcess = value
}
// SetDefenderOfficeMacroCodeAllowWin32Imports sets the defenderOfficeMacroCodeAllowWin32Imports property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeMacroCodeAllowWin32Imports(value *DefenderProtectionType)() {
    m.defenderOfficeMacroCodeAllowWin32Imports = value
}
// SetDefenderOfficeMacroCodeAllowWin32ImportsType sets the defenderOfficeMacroCodeAllowWin32ImportsType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeMacroCodeAllowWin32ImportsType(value *DefenderAttackSurfaceType)() {
    m.defenderOfficeMacroCodeAllowWin32ImportsType = value
}
// SetDefenderPotentiallyUnwantedAppAction sets the defenderPotentiallyUnwantedAppAction property value. Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderPotentiallyUnwantedAppAction(value *DefenderProtectionType)() {
    m.defenderPotentiallyUnwantedAppAction = value
}
// SetDefenderPreventCredentialStealingType sets the defenderPreventCredentialStealingType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderPreventCredentialStealingType(value *DefenderProtectionType)() {
    m.defenderPreventCredentialStealingType = value
}
// SetDefenderProcessCreation sets the defenderProcessCreation property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessCreation(value *DefenderProtectionType)() {
    m.defenderProcessCreation = value
}
// SetDefenderProcessCreationType sets the defenderProcessCreationType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessCreationType(value *DefenderAttackSurfaceType)() {
    m.defenderProcessCreationType = value
}
// SetDefenderProcessesToExclude sets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessesToExclude(value []string)() {
    m.defenderProcessesToExclude = value
}
// SetDefenderScanDirection sets the defenderScanDirection property value. Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanDirection(value *DefenderRealtimeScanDirection)() {
    m.defenderScanDirection = value
}
// SetDefenderScanMaxCpuPercentage sets the defenderScanMaxCpuPercentage property value. Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanMaxCpuPercentage(value *int32)() {
    m.defenderScanMaxCpuPercentage = value
}
// SetDefenderScanType sets the defenderScanType property value. Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanType(value *DefenderScanType)() {
    m.defenderScanType = value
}
// SetDefenderScheduledQuickScanTime sets the defenderScheduledQuickScanTime property value. Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledQuickScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.defenderScheduledQuickScanTime = value
}
// SetDefenderScheduledScanDay sets the defenderScheduledScanDay property value. Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledScanDay(value *WeeklySchedule)() {
    m.defenderScheduledScanDay = value
}
// SetDefenderScheduledScanTime sets the defenderScheduledScanTime property value. Selects the time of day that the Windows Defender scan should run.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.defenderScheduledScanTime = value
}
// SetDefenderScriptDownloadedPayloadExecution sets the defenderScriptDownloadedPayloadExecution property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptDownloadedPayloadExecution(value *DefenderProtectionType)() {
    m.defenderScriptDownloadedPayloadExecution = value
}
// SetDefenderScriptDownloadedPayloadExecutionType sets the defenderScriptDownloadedPayloadExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptDownloadedPayloadExecutionType(value *DefenderAttackSurfaceType)() {
    m.defenderScriptDownloadedPayloadExecutionType = value
}
// SetDefenderScriptObfuscatedMacroCode sets the defenderScriptObfuscatedMacroCode property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptObfuscatedMacroCode(value *DefenderProtectionType)() {
    m.defenderScriptObfuscatedMacroCode = value
}
// SetDefenderScriptObfuscatedMacroCodeType sets the defenderScriptObfuscatedMacroCodeType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptObfuscatedMacroCodeType(value *DefenderAttackSurfaceType)() {
    m.defenderScriptObfuscatedMacroCodeType = value
}
// SetDefenderSecurityCenterBlockExploitProtectionOverride sets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterBlockExploitProtectionOverride(value *bool)() {
    m.defenderSecurityCenterBlockExploitProtectionOverride = value
}
// SetDefenderSecurityCenterDisableAccountUI sets the defenderSecurityCenterDisableAccountUI property value. Used to disable the display of the account protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableAccountUI(value *bool)() {
    m.defenderSecurityCenterDisableAccountUI = value
}
// SetDefenderSecurityCenterDisableAppBrowserUI sets the defenderSecurityCenterDisableAppBrowserUI property value. Used to disable the display of the app and browser protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableAppBrowserUI(value *bool)() {
    m.defenderSecurityCenterDisableAppBrowserUI = value
}
// SetDefenderSecurityCenterDisableClearTpmUI sets the defenderSecurityCenterDisableClearTpmUI property value. Used to disable the display of the Clear TPM button.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableClearTpmUI(value *bool)() {
    m.defenderSecurityCenterDisableClearTpmUI = value
}
// SetDefenderSecurityCenterDisableFamilyUI sets the defenderSecurityCenterDisableFamilyUI property value. Used to disable the display of the family options area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableFamilyUI(value *bool)() {
    m.defenderSecurityCenterDisableFamilyUI = value
}
// SetDefenderSecurityCenterDisableHardwareUI sets the defenderSecurityCenterDisableHardwareUI property value. Used to disable the display of the hardware protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableHardwareUI(value *bool)() {
    m.defenderSecurityCenterDisableHardwareUI = value
}
// SetDefenderSecurityCenterDisableHealthUI sets the defenderSecurityCenterDisableHealthUI property value. Used to disable the display of the device performance and health area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableHealthUI(value *bool)() {
    m.defenderSecurityCenterDisableHealthUI = value
}
// SetDefenderSecurityCenterDisableNetworkUI sets the defenderSecurityCenterDisableNetworkUI property value. Used to disable the display of the firewall and network protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableNetworkUI(value *bool)() {
    m.defenderSecurityCenterDisableNetworkUI = value
}
// SetDefenderSecurityCenterDisableNotificationAreaUI sets the defenderSecurityCenterDisableNotificationAreaUI property value. Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableNotificationAreaUI(value *bool)() {
    m.defenderSecurityCenterDisableNotificationAreaUI = value
}
// SetDefenderSecurityCenterDisableRansomwareUI sets the defenderSecurityCenterDisableRansomwareUI property value. Used to disable the display of the ransomware protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableRansomwareUI(value *bool)() {
    m.defenderSecurityCenterDisableRansomwareUI = value
}
// SetDefenderSecurityCenterDisableSecureBootUI sets the defenderSecurityCenterDisableSecureBootUI property value. Used to disable the display of the secure boot area under Device security.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableSecureBootUI(value *bool)() {
    m.defenderSecurityCenterDisableSecureBootUI = value
}
// SetDefenderSecurityCenterDisableTroubleshootingUI sets the defenderSecurityCenterDisableTroubleshootingUI property value. Used to disable the display of the security process troubleshooting under Device security.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableTroubleshootingUI(value *bool)() {
    m.defenderSecurityCenterDisableTroubleshootingUI = value
}
// SetDefenderSecurityCenterDisableVirusUI sets the defenderSecurityCenterDisableVirusUI property value. Used to disable the display of the virus and threat protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableVirusUI(value *bool)() {
    m.defenderSecurityCenterDisableVirusUI = value
}
// SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI sets the defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI property value. Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI(value *bool)() {
    m.defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI = value
}
// SetDefenderSecurityCenterHelpEmail sets the defenderSecurityCenterHelpEmail property value. The email address that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpEmail(value *string)() {
    m.defenderSecurityCenterHelpEmail = value
}
// SetDefenderSecurityCenterHelpPhone sets the defenderSecurityCenterHelpPhone property value. The phone number or Skype ID that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpPhone(value *string)() {
    m.defenderSecurityCenterHelpPhone = value
}
// SetDefenderSecurityCenterHelpURL sets the defenderSecurityCenterHelpURL property value. The help portal URL this is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpURL(value *string)() {
    m.defenderSecurityCenterHelpURL = value
}
// SetDefenderSecurityCenterITContactDisplay sets the defenderSecurityCenterITContactDisplay property value. Possible values for defenderSecurityCenterITContactDisplay
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterITContactDisplay(value *DefenderSecurityCenterITContactDisplayType)() {
    m.defenderSecurityCenterITContactDisplay = value
}
// SetDefenderSecurityCenterNotificationsFromApp sets the defenderSecurityCenterNotificationsFromApp property value. Possible values for defenderSecurityCenterNotificationsFromApp
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterNotificationsFromApp(value *DefenderSecurityCenterNotificationsFromAppType)() {
    m.defenderSecurityCenterNotificationsFromApp = value
}
// SetDefenderSecurityCenterOrganizationDisplayName sets the defenderSecurityCenterOrganizationDisplayName property value. The company name that is displayed to the users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterOrganizationDisplayName(value *string)() {
    m.defenderSecurityCenterOrganizationDisplayName = value
}
// SetDefenderSignatureUpdateIntervalInHours sets the defenderSignatureUpdateIntervalInHours property value. Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSignatureUpdateIntervalInHours(value *int32)() {
    m.defenderSignatureUpdateIntervalInHours = value
}
// SetDefenderSubmitSamplesConsentType sets the defenderSubmitSamplesConsentType property value. Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSubmitSamplesConsentType(value *DefenderSubmitSamplesConsentType)() {
    m.defenderSubmitSamplesConsentType = value
}
// SetDefenderUntrustedExecutable sets the defenderUntrustedExecutable property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedExecutable(value *DefenderProtectionType)() {
    m.defenderUntrustedExecutable = value
}
// SetDefenderUntrustedExecutableType sets the defenderUntrustedExecutableType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedExecutableType(value *DefenderAttackSurfaceType)() {
    m.defenderUntrustedExecutableType = value
}
// SetDefenderUntrustedUSBProcess sets the defenderUntrustedUSBProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedUSBProcess(value *DefenderProtectionType)() {
    m.defenderUntrustedUSBProcess = value
}
// SetDefenderUntrustedUSBProcessType sets the defenderUntrustedUSBProcessType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedUSBProcessType(value *DefenderAttackSurfaceType)() {
    m.defenderUntrustedUSBProcessType = value
}
// SetDeviceGuardEnableSecureBootWithDMA sets the deviceGuardEnableSecureBootWithDMA property value. This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardEnableSecureBootWithDMA(value *bool)() {
    m.deviceGuardEnableSecureBootWithDMA = value
}
// SetDeviceGuardEnableVirtualizationBasedSecurity sets the deviceGuardEnableVirtualizationBasedSecurity property value. Turns On Virtualization Based Security(VBS).
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardEnableVirtualizationBasedSecurity(value *bool)() {
    m.deviceGuardEnableVirtualizationBasedSecurity = value
}
// SetDeviceGuardLaunchSystemGuard sets the deviceGuardLaunchSystemGuard property value. Possible values of a property
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardLaunchSystemGuard(value *Enablement)() {
    m.deviceGuardLaunchSystemGuard = value
}
// SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings sets the deviceGuardLocalSystemAuthorityCredentialGuardSettings property value. Possible values of Credential Guard settings.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings(value *DeviceGuardLocalSystemAuthorityCredentialGuardType)() {
    m.deviceGuardLocalSystemAuthorityCredentialGuardSettings = value
}
// SetDeviceGuardSecureBootWithDMA sets the deviceGuardSecureBootWithDMA property value. Possible values of Secure Boot with DMA
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardSecureBootWithDMA(value *SecureBootWithDMAType)() {
    m.deviceGuardSecureBootWithDMA = value
}
// SetDmaGuardDeviceEnumerationPolicy sets the dmaGuardDeviceEnumerationPolicy property value. Possible values of the DmaGuardDeviceEnumerationPolicy.
func (m *Windows10EndpointProtectionConfiguration) SetDmaGuardDeviceEnumerationPolicy(value *DmaGuardDeviceEnumerationPolicyType)() {
    m.dmaGuardDeviceEnumerationPolicy = value
}
// SetFirewallBlockStatefulFTP sets the firewallBlockStatefulFTP property value. Blocks stateful FTP connections to the device
func (m *Windows10EndpointProtectionConfiguration) SetFirewallBlockStatefulFTP(value *bool)() {
    m.firewallBlockStatefulFTP = value
}
// SetFirewallCertificateRevocationListCheckMethod sets the firewallCertificateRevocationListCheckMethod property value. Possible values for firewallCertificateRevocationListCheckMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallCertificateRevocationListCheckMethod(value *FirewallCertificateRevocationListCheckMethodType)() {
    m.firewallCertificateRevocationListCheckMethod = value
}
// SetFirewallIdleTimeoutForSecurityAssociationInSeconds sets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIdleTimeoutForSecurityAssociationInSeconds(value *int32)() {
    m.firewallIdleTimeoutForSecurityAssociationInSeconds = value
}
// SetFirewallIPSecExemptionsAllowDHCP sets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowDHCP(value *bool)() {
    m.firewallIPSecExemptionsAllowDHCP = value
}
// SetFirewallIPSecExemptionsAllowICMP sets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowICMP(value *bool)() {
    m.firewallIPSecExemptionsAllowICMP = value
}
// SetFirewallIPSecExemptionsAllowNeighborDiscovery sets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowNeighborDiscovery(value *bool)() {
    m.firewallIPSecExemptionsAllowNeighborDiscovery = value
}
// SetFirewallIPSecExemptionsAllowRouterDiscovery sets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowRouterDiscovery(value *bool)() {
    m.firewallIPSecExemptionsAllowRouterDiscovery = value
}
// SetFirewallIPSecExemptionsNone sets the firewallIPSecExemptionsNone property value. Configures IPSec exemptions to no exemptions
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsNone(value *bool)() {
    m.firewallIPSecExemptionsNone = value
}
// SetFirewallMergeKeyingModuleSettings sets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
func (m *Windows10EndpointProtectionConfiguration) SetFirewallMergeKeyingModuleSettings(value *bool)() {
    m.firewallMergeKeyingModuleSettings = value
}
// SetFirewallPacketQueueingMethod sets the firewallPacketQueueingMethod property value. Possible values for firewallPacketQueueingMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPacketQueueingMethod(value *FirewallPacketQueueingMethodType)() {
    m.firewallPacketQueueingMethod = value
}
// SetFirewallPreSharedKeyEncodingMethod sets the firewallPreSharedKeyEncodingMethod property value. Possible values for firewallPreSharedKeyEncodingMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPreSharedKeyEncodingMethod(value *FirewallPreSharedKeyEncodingMethodType)() {
    m.firewallPreSharedKeyEncodingMethod = value
}
// SetFirewallProfileDomain sets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfileDomain(value WindowsFirewallNetworkProfileable)() {
    m.firewallProfileDomain = value
}
// SetFirewallProfilePrivate sets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePrivate(value WindowsFirewallNetworkProfileable)() {
    m.firewallProfilePrivate = value
}
// SetFirewallProfilePublic sets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePublic(value WindowsFirewallNetworkProfileable)() {
    m.firewallProfilePublic = value
}
// SetFirewallRules sets the firewallRules property value. Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
func (m *Windows10EndpointProtectionConfiguration) SetFirewallRules(value []WindowsFirewallRuleable)() {
    m.firewallRules = value
}
// SetLanManagerAuthenticationLevel sets the lanManagerAuthenticationLevel property value. Possible values for LanManagerAuthenticationLevel
func (m *Windows10EndpointProtectionConfiguration) SetLanManagerAuthenticationLevel(value *LanManagerAuthenticationLevel)() {
    m.lanManagerAuthenticationLevel = value
}
// SetLanManagerWorkstationDisableInsecureGuestLogons sets the lanManagerWorkstationDisableInsecureGuestLogons property value. If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
func (m *Windows10EndpointProtectionConfiguration) SetLanManagerWorkstationDisableInsecureGuestLogons(value *bool)() {
    m.lanManagerWorkstationDisableInsecureGuestLogons = value
}
// SetLocalSecurityOptionsAdministratorAccountName sets the localSecurityOptionsAdministratorAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAdministratorAccountName(value *string)() {
    m.localSecurityOptionsAdministratorAccountName = value
}
// SetLocalSecurityOptionsAdministratorElevationPromptBehavior sets the localSecurityOptionsAdministratorElevationPromptBehavior property value. Possible values for LocalSecurityOptionsAdministratorElevationPromptBehavior
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAdministratorElevationPromptBehavior(value *LocalSecurityOptionsAdministratorElevationPromptBehaviorType)() {
    m.localSecurityOptionsAdministratorElevationPromptBehavior = value
}
// SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares sets the localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares property value. This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares(value *bool)() {
    m.localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares = value
}
// SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests sets the localSecurityOptionsAllowPKU2UAuthenticationRequests property value. Block PKU2U authentication requests to this device to use online identities.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests(value *bool)() {
    m.localSecurityOptionsAllowPKU2UAuthenticationRequests = value
}
// SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager sets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager property value. Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager(value *string)() {
    m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager = value
}
// SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool sets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool property value. UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool(value *bool)() {
    m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool = value
}
// SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn sets the localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn property value. This security setting determines whether a computer can be shut down without having to log on to Windows.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn(value *bool)() {
    m.localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn = value
}
// SetLocalSecurityOptionsAllowUIAccessApplicationElevation sets the localSecurityOptionsAllowUIAccessApplicationElevation property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUIAccessApplicationElevation(value *bool)() {
    m.localSecurityOptionsAllowUIAccessApplicationElevation = value
}
// SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations sets the localSecurityOptionsAllowUIAccessApplicationsForSecureLocations property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations(value *bool)() {
    m.localSecurityOptionsAllowUIAccessApplicationsForSecureLocations = value
}
// SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon sets the localSecurityOptionsAllowUndockWithoutHavingToLogon property value. Prevent a portable computer from being undocked without having to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon(value *bool)() {
    m.localSecurityOptionsAllowUndockWithoutHavingToLogon = value
}
// SetLocalSecurityOptionsBlockMicrosoftAccounts sets the localSecurityOptionsBlockMicrosoftAccounts property value. Prevent users from adding new Microsoft accounts to this computer.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockMicrosoftAccounts(value *bool)() {
    m.localSecurityOptionsBlockMicrosoftAccounts = value
}
// SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword sets the localSecurityOptionsBlockRemoteLogonWithBlankPassword property value. Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword(value *bool)() {
    m.localSecurityOptionsBlockRemoteLogonWithBlankPassword = value
}
// SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess sets the localSecurityOptionsBlockRemoteOpticalDriveAccess property value. Enabling this settings allows only interactively logged on user to access CD-ROM media.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess(value *bool)() {
    m.localSecurityOptionsBlockRemoteOpticalDriveAccess = value
}
// SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers sets the localSecurityOptionsBlockUsersInstallingPrinterDrivers property value. Restrict installing printer drivers as part of connecting to a shared printer to admins only.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers(value *bool)() {
    m.localSecurityOptionsBlockUsersInstallingPrinterDrivers = value
}
// SetLocalSecurityOptionsClearVirtualMemoryPageFile sets the localSecurityOptionsClearVirtualMemoryPageFile property value. This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClearVirtualMemoryPageFile(value *bool)() {
    m.localSecurityOptionsClearVirtualMemoryPageFile = value
}
// SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways sets the localSecurityOptionsClientDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB client component.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways(value *bool)() {
    m.localSecurityOptionsClientDigitallySignCommunicationsAlways = value
}
// SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers sets the localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers property value. If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers(value *bool)() {
    m.localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers = value
}
// SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation sets the localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation property value. App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation(value *bool)() {
    m.localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation = value
}
// SetLocalSecurityOptionsDisableAdministratorAccount sets the localSecurityOptionsDisableAdministratorAccount property value. Determines whether the Local Administrator account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableAdministratorAccount(value *bool)() {
    m.localSecurityOptionsDisableAdministratorAccount = value
}
// SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees sets the localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees property value. This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees(value *bool)() {
    m.localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees = value
}
// SetLocalSecurityOptionsDisableGuestAccount sets the localSecurityOptionsDisableGuestAccount property value. Determines if the Guest account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableGuestAccount(value *bool)() {
    m.localSecurityOptionsDisableGuestAccount = value
}
// SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways sets the localSecurityOptionsDisableServerDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB server component.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways(value *bool)() {
    m.localSecurityOptionsDisableServerDigitallySignCommunicationsAlways = value
}
// SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees sets the localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees property value. This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees(value *bool)() {
    m.localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees = value
}
// SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts sets the localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts property value. This security setting determines what additional permissions will be granted for anonymous connections to the computer.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts(value *bool)() {
    m.localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts = value
}
// SetLocalSecurityOptionsDoNotRequireCtrlAltDel sets the localSecurityOptionsDoNotRequireCtrlAltDel property value. Require CTRL+ALT+DEL to be pressed before a user can log on.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotRequireCtrlAltDel(value *bool)() {
    m.localSecurityOptionsDoNotRequireCtrlAltDel = value
}
// SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange sets the localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange property value. This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange(value *bool)() {
    m.localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange = value
}
// SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser sets the localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser property value. Possible values for LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(value *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)() {
    m.localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = value
}
// SetLocalSecurityOptionsGuestAccountName sets the localSecurityOptionsGuestAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsGuestAccountName(value *string)() {
    m.localSecurityOptionsGuestAccountName = value
}
// SetLocalSecurityOptionsHideLastSignedInUser sets the localSecurityOptionsHideLastSignedInUser property value. Do not display the username of the last person who signed in on this device.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsHideLastSignedInUser(value *bool)() {
    m.localSecurityOptionsHideLastSignedInUser = value
}
// SetLocalSecurityOptionsHideUsernameAtSignIn sets the localSecurityOptionsHideUsernameAtSignIn property value. Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsHideUsernameAtSignIn(value *bool)() {
    m.localSecurityOptionsHideUsernameAtSignIn = value
}
// SetLocalSecurityOptionsInformationDisplayedOnLockScreen sets the localSecurityOptionsInformationDisplayedOnLockScreen property value. Possible values for LocalSecurityOptionsInformationDisplayedOnLockScreen
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsInformationDisplayedOnLockScreen(value *LocalSecurityOptionsInformationDisplayedOnLockScreenType)() {
    m.localSecurityOptionsInformationDisplayedOnLockScreen = value
}
// SetLocalSecurityOptionsInformationShownOnLockScreen sets the localSecurityOptionsInformationShownOnLockScreen property value. Possible values for LocalSecurityOptionsInformationShownOnLockScreenType
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsInformationShownOnLockScreen(value *LocalSecurityOptionsInformationShownOnLockScreenType)() {
    m.localSecurityOptionsInformationShownOnLockScreen = value
}
// SetLocalSecurityOptionsLogOnMessageText sets the localSecurityOptionsLogOnMessageText property value. Set message text for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsLogOnMessageText(value *string)() {
    m.localSecurityOptionsLogOnMessageText = value
}
// SetLocalSecurityOptionsLogOnMessageTitle sets the localSecurityOptionsLogOnMessageTitle property value. Set message title for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsLogOnMessageTitle(value *string)() {
    m.localSecurityOptionsLogOnMessageTitle = value
}
// SetLocalSecurityOptionsMachineInactivityLimit sets the localSecurityOptionsMachineInactivityLimit property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMachineInactivityLimit(value *int32)() {
    m.localSecurityOptionsMachineInactivityLimit = value
}
// SetLocalSecurityOptionsMachineInactivityLimitInMinutes sets the localSecurityOptionsMachineInactivityLimitInMinutes property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMachineInactivityLimitInMinutes(value *int32)() {
    m.localSecurityOptionsMachineInactivityLimitInMinutes = value
}
// SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients sets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(value *LocalSecurityOptionsMinimumSessionSecurity)() {
    m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients = value
}
// SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers sets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(value *LocalSecurityOptionsMinimumSessionSecurity)() {
    m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers = value
}
// SetLocalSecurityOptionsOnlyElevateSignedExecutables sets the localSecurityOptionsOnlyElevateSignedExecutables property value. Enforce PKI certification path validation for a given executable file before it is permitted to run.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsOnlyElevateSignedExecutables(value *bool)() {
    m.localSecurityOptionsOnlyElevateSignedExecutables = value
}
// SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares sets the localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares property value. By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares(value *bool)() {
    m.localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares = value
}
// SetLocalSecurityOptionsSmartCardRemovalBehavior sets the localSecurityOptionsSmartCardRemovalBehavior property value. Possible values for LocalSecurityOptionsSmartCardRemovalBehaviorType
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsSmartCardRemovalBehavior(value *LocalSecurityOptionsSmartCardRemovalBehaviorType)() {
    m.localSecurityOptionsSmartCardRemovalBehavior = value
}
// SetLocalSecurityOptionsStandardUserElevationPromptBehavior sets the localSecurityOptionsStandardUserElevationPromptBehavior property value. Possible values for LocalSecurityOptionsStandardUserElevationPromptBehavior
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsStandardUserElevationPromptBehavior(value *LocalSecurityOptionsStandardUserElevationPromptBehaviorType)() {
    m.localSecurityOptionsStandardUserElevationPromptBehavior = value
}
// SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation sets the localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation property value. Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation(value *bool)() {
    m.localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation = value
}
// SetLocalSecurityOptionsUseAdminApprovalMode sets the localSecurityOptionsUseAdminApprovalMode property value. Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsUseAdminApprovalMode(value *bool)() {
    m.localSecurityOptionsUseAdminApprovalMode = value
}
// SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators sets the localSecurityOptionsUseAdminApprovalModeForAdministrators property value. Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators(value *bool)() {
    m.localSecurityOptionsUseAdminApprovalModeForAdministrators = value
}
// SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations sets the localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations property value. Virtualize file and registry write failures to per user locations
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations(value *bool)() {
    m.localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations = value
}
// SetSmartScreenBlockOverrideForFiles sets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenBlockOverrideForFiles(value *bool)() {
    m.smartScreenBlockOverrideForFiles = value
}
// SetSmartScreenEnableInShell sets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenEnableInShell(value *bool)() {
    m.smartScreenEnableInShell = value
}
// SetUserRightsAccessCredentialManagerAsTrustedCaller sets the userRightsAccessCredentialManagerAsTrustedCaller property value. This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsAccessCredentialManagerAsTrustedCaller(value DeviceManagementUserRightsSettingable)() {
    m.userRightsAccessCredentialManagerAsTrustedCaller = value
}
// SetUserRightsActAsPartOfTheOperatingSystem sets the userRightsActAsPartOfTheOperatingSystem property value. This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsActAsPartOfTheOperatingSystem(value DeviceManagementUserRightsSettingable)() {
    m.userRightsActAsPartOfTheOperatingSystem = value
}
// SetUserRightsAllowAccessFromNetwork sets the userRightsAllowAccessFromNetwork property value. This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsAllowAccessFromNetwork(value DeviceManagementUserRightsSettingable)() {
    m.userRightsAllowAccessFromNetwork = value
}
// SetUserRightsBackupData sets the userRightsBackupData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsBackupData(value DeviceManagementUserRightsSettingable)() {
    m.userRightsBackupData = value
}
// SetUserRightsBlockAccessFromNetwork sets the userRightsBlockAccessFromNetwork property value. This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsBlockAccessFromNetwork(value DeviceManagementUserRightsSettingable)() {
    m.userRightsBlockAccessFromNetwork = value
}
// SetUserRightsChangeSystemTime sets the userRightsChangeSystemTime property value. This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsChangeSystemTime(value DeviceManagementUserRightsSettingable)() {
    m.userRightsChangeSystemTime = value
}
// SetUserRightsCreateGlobalObjects sets the userRightsCreateGlobalObjects property value. This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateGlobalObjects(value DeviceManagementUserRightsSettingable)() {
    m.userRightsCreateGlobalObjects = value
}
// SetUserRightsCreatePageFile sets the userRightsCreatePageFile property value. This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreatePageFile(value DeviceManagementUserRightsSettingable)() {
    m.userRightsCreatePageFile = value
}
// SetUserRightsCreatePermanentSharedObjects sets the userRightsCreatePermanentSharedObjects property value. This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreatePermanentSharedObjects(value DeviceManagementUserRightsSettingable)() {
    m.userRightsCreatePermanentSharedObjects = value
}
// SetUserRightsCreateSymbolicLinks sets the userRightsCreateSymbolicLinks property value. This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateSymbolicLinks(value DeviceManagementUserRightsSettingable)() {
    m.userRightsCreateSymbolicLinks = value
}
// SetUserRightsCreateToken sets the userRightsCreateToken property value. This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateToken(value DeviceManagementUserRightsSettingable)() {
    m.userRightsCreateToken = value
}
// SetUserRightsDebugPrograms sets the userRightsDebugPrograms property value. This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDebugPrograms(value DeviceManagementUserRightsSettingable)() {
    m.userRightsDebugPrograms = value
}
// SetUserRightsDelegation sets the userRightsDelegation property value. This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDelegation(value DeviceManagementUserRightsSettingable)() {
    m.userRightsDelegation = value
}
// SetUserRightsDenyLocalLogOn sets the userRightsDenyLocalLogOn property value. This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDenyLocalLogOn(value DeviceManagementUserRightsSettingable)() {
    m.userRightsDenyLocalLogOn = value
}
// SetUserRightsGenerateSecurityAudits sets the userRightsGenerateSecurityAudits property value. This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsGenerateSecurityAudits(value DeviceManagementUserRightsSettingable)() {
    m.userRightsGenerateSecurityAudits = value
}
// SetUserRightsImpersonateClient sets the userRightsImpersonateClient property value. Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsImpersonateClient(value DeviceManagementUserRightsSettingable)() {
    m.userRightsImpersonateClient = value
}
// SetUserRightsIncreaseSchedulingPriority sets the userRightsIncreaseSchedulingPriority property value. This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsIncreaseSchedulingPriority(value DeviceManagementUserRightsSettingable)() {
    m.userRightsIncreaseSchedulingPriority = value
}
// SetUserRightsLoadUnloadDrivers sets the userRightsLoadUnloadDrivers property value. This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLoadUnloadDrivers(value DeviceManagementUserRightsSettingable)() {
    m.userRightsLoadUnloadDrivers = value
}
// SetUserRightsLocalLogOn sets the userRightsLocalLogOn property value. This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLocalLogOn(value DeviceManagementUserRightsSettingable)() {
    m.userRightsLocalLogOn = value
}
// SetUserRightsLockMemory sets the userRightsLockMemory property value. This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLockMemory(value DeviceManagementUserRightsSettingable)() {
    m.userRightsLockMemory = value
}
// SetUserRightsManageAuditingAndSecurityLogs sets the userRightsManageAuditingAndSecurityLogs property value. This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsManageAuditingAndSecurityLogs(value DeviceManagementUserRightsSettingable)() {
    m.userRightsManageAuditingAndSecurityLogs = value
}
// SetUserRightsManageVolumes sets the userRightsManageVolumes property value. This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsManageVolumes(value DeviceManagementUserRightsSettingable)() {
    m.userRightsManageVolumes = value
}
// SetUserRightsModifyFirmwareEnvironment sets the userRightsModifyFirmwareEnvironment property value. This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsModifyFirmwareEnvironment(value DeviceManagementUserRightsSettingable)() {
    m.userRightsModifyFirmwareEnvironment = value
}
// SetUserRightsModifyObjectLabels sets the userRightsModifyObjectLabels property value. This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsModifyObjectLabels(value DeviceManagementUserRightsSettingable)() {
    m.userRightsModifyObjectLabels = value
}
// SetUserRightsProfileSingleProcess sets the userRightsProfileSingleProcess property value. This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsProfileSingleProcess(value DeviceManagementUserRightsSettingable)() {
    m.userRightsProfileSingleProcess = value
}
// SetUserRightsRemoteDesktopServicesLogOn sets the userRightsRemoteDesktopServicesLogOn property value. This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRemoteDesktopServicesLogOn(value DeviceManagementUserRightsSettingable)() {
    m.userRightsRemoteDesktopServicesLogOn = value
}
// SetUserRightsRemoteShutdown sets the userRightsRemoteShutdown property value. This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRemoteShutdown(value DeviceManagementUserRightsSettingable)() {
    m.userRightsRemoteShutdown = value
}
// SetUserRightsRestoreData sets the userRightsRestoreData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRestoreData(value DeviceManagementUserRightsSettingable)() {
    m.userRightsRestoreData = value
}
// SetUserRightsTakeOwnership sets the userRightsTakeOwnership property value. This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsTakeOwnership(value DeviceManagementUserRightsSettingable)() {
    m.userRightsTakeOwnership = value
}
// SetWindowsDefenderTamperProtection sets the windowsDefenderTamperProtection property value. Defender TamperProtection setting options
func (m *Windows10EndpointProtectionConfiguration) SetWindowsDefenderTamperProtection(value *WindowsDefenderTamperProtectionOptions)() {
    m.windowsDefenderTamperProtection = value
}
// SetXboxServicesAccessoryManagementServiceStartupMode sets the xboxServicesAccessoryManagementServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesAccessoryManagementServiceStartupMode(value *ServiceStartType)() {
    m.xboxServicesAccessoryManagementServiceStartupMode = value
}
// SetXboxServicesEnableXboxGameSaveTask sets the xboxServicesEnableXboxGameSaveTask property value. This setting determines whether xbox game save is enabled (1) or disabled (0).
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesEnableXboxGameSaveTask(value *bool)() {
    m.xboxServicesEnableXboxGameSaveTask = value
}
// SetXboxServicesLiveAuthManagerServiceStartupMode sets the xboxServicesLiveAuthManagerServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveAuthManagerServiceStartupMode(value *ServiceStartType)() {
    m.xboxServicesLiveAuthManagerServiceStartupMode = value
}
// SetXboxServicesLiveGameSaveServiceStartupMode sets the xboxServicesLiveGameSaveServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveGameSaveServiceStartupMode(value *ServiceStartType)() {
    m.xboxServicesLiveGameSaveServiceStartupMode = value
}
// SetXboxServicesLiveNetworkingServiceStartupMode sets the xboxServicesLiveNetworkingServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveNetworkingServiceStartupMode(value *ServiceStartType)() {
    m.xboxServicesLiveNetworkingServiceStartupMode = value
}
