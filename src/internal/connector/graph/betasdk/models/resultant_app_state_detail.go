package models
import (
    "errors"
)
// Provides operations to call the add method.
type ResultantAppStateDetail int

const (
    // Device architecture (e.g. x86/amd64) is not applicable for the application.
    PROCESSORARCHITECTURENOTAPPLICABLE_RESULTANTAPPSTATEDETAIL ResultantAppStateDetail = iota
    // Available disk space on the target device is less than the configured minimum.
    MINIMUMDISKSPACENOTMET_RESULTANTAPPSTATEDETAIL
    // OS version on the target device is less than the configured minimum.
    MINIMUMOSVERSIONNOTMET_RESULTANTAPPSTATEDETAIL
    // Amount of RAM on the target device is less than the configured minimum.
    MINIMUMPHYSICALMEMORYNOTMET_RESULTANTAPPSTATEDETAIL
    // Count of logical processors on the target device is less than the configured minimum.
    MINIMUMLOGICALPROCESSORCOUNTNOTMET_RESULTANTAPPSTATEDETAIL
    // CPU speed on the target device is less than the configured minimum.
    MINIMUMCPUSPEEDNOTMET_RESULTANTAPPSTATEDETAIL
    // Application is not applicable to this platform. (e.g. Android app targeted to IOS)
    PLATFORMNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
    // File system requirement rule is not met
    FILESYSTEMREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
    // Registry requirement rule is not met
    REGISTRYREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
    // PowerShell script requirement rule is not met
    POWERSHELLSCRIPTREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
    // All targeted, superseding apps are not applicable.
    SUPERSEDINGAPPSNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
    // No additional details are available.
    NOADDITIONALDETAILS_RESULTANTAPPSTATEDETAIL
    // One or more of the application's dependencies failed to install.
    DEPENDENCYFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
    // One or more of the application's dependencies have requirements which are not met.
    DEPENDENCYWITHREQUIREMENTSNOTMET_RESULTANTAPPSTATEDETAIL
    // One or more of the application's dependencies require a device reboot to complete installation.
    DEPENDENCYPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    // One or more of the application's dependencies are configured to not automatically install.
    DEPENDENCYWITHAUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
    // A superseded app failed to uninstall.
    SUPERSEDEDAPPUNINSTALLFAILED_RESULTANTAPPSTATEDETAIL
    // A superseded app requires a reboot to complete uninstall.
    SUPERSEDEDAPPUNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    // Superseded apps are being removed.
    REMOVINGSUPERSEDEDAPPS_RESULTANTAPPSTATEDETAIL
    // The latest version of the app failed to update from an earlier version.
    IOSAPPSTOREUPDATEFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
    // An update is available.
    VPPAPPHASUPDATEAVAILABLE_RESULTANTAPPSTATEDETAIL
    // The user rejected the app update.
    USERREJECTEDUPDATE_RESULTANTAPPSTATEDETAIL
    // To complete the removal of the app, the device must be rebooted.
    UNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    // Superseding applications are detected.
    SUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
    // Superseded applications are detected.
    SUPERSEDEDAPPSDETECTED_RESULTANTAPPSTATEDETAIL
    // Application failed to install. See error code property for more details.
    SEEINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
    // Application is configured to not be automatically installed.
    AUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
    // The app is managed but no longer installed.
    MANAGEDAPPNOLONGERPRESENT_RESULTANTAPPSTATEDETAIL
    // The user rejected the app install.
    USERREJECTEDINSTALL_RESULTANTAPPSTATEDETAIL
    // The user must log into the App Store to install app.
    USERISNOTLOGGEDINTOAPPSTORE_RESULTANTAPPSTATEDETAIL
    // App cannot be installed. An untargeted, superseding app was detected, which created a conflict.
    UNTARGETEDSUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
    // App was removed in order to install a superseding app.
    APPREMOVEDBYSUPERSEDENCE_RESULTANTAPPSTATEDETAIL
    // Application failed to uninstall. See error code property for more details.
    SEEUNINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
    // Device must be rebooted to complete installation of the application.
    PENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    // One or more of the application's dependencies are installing.
    INSTALLINGDEPENDENCIES_RESULTANTAPPSTATEDETAIL
    // Application content was downloaded to the device.
    CONTENTDOWNLOADED_RESULTANTAPPSTATEDETAIL
)

func (i ResultantAppStateDetail) String() string {
    return []string{"processorArchitectureNotApplicable", "minimumDiskSpaceNotMet", "minimumOsVersionNotMet", "minimumPhysicalMemoryNotMet", "minimumLogicalProcessorCountNotMet", "minimumCpuSpeedNotMet", "platformNotApplicable", "fileSystemRequirementNotMet", "registryRequirementNotMet", "powerShellScriptRequirementNotMet", "supersedingAppsNotApplicable", "noAdditionalDetails", "dependencyFailedToInstall", "dependencyWithRequirementsNotMet", "dependencyPendingReboot", "dependencyWithAutoInstallDisabled", "supersededAppUninstallFailed", "supersededAppUninstallPendingReboot", "removingSupersededApps", "iosAppStoreUpdateFailedToInstall", "vppAppHasUpdateAvailable", "userRejectedUpdate", "uninstallPendingReboot", "supersedingAppsDetected", "supersededAppsDetected", "seeInstallErrorCode", "autoInstallDisabled", "managedAppNoLongerPresent", "userRejectedInstall", "userIsNotLoggedIntoAppStore", "untargetedSupersedingAppsDetected", "appRemovedBySupersedence", "seeUninstallErrorCode", "pendingReboot", "installingDependencies", "contentDownloaded"}[i]
}
func ParseResultantAppStateDetail(v string) (interface{}, error) {
    result := PROCESSORARCHITECTURENOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
    switch v {
        case "processorArchitectureNotApplicable":
            result = PROCESSORARCHITECTURENOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
        case "minimumDiskSpaceNotMet":
            result = MINIMUMDISKSPACENOTMET_RESULTANTAPPSTATEDETAIL
        case "minimumOsVersionNotMet":
            result = MINIMUMOSVERSIONNOTMET_RESULTANTAPPSTATEDETAIL
        case "minimumPhysicalMemoryNotMet":
            result = MINIMUMPHYSICALMEMORYNOTMET_RESULTANTAPPSTATEDETAIL
        case "minimumLogicalProcessorCountNotMet":
            result = MINIMUMLOGICALPROCESSORCOUNTNOTMET_RESULTANTAPPSTATEDETAIL
        case "minimumCpuSpeedNotMet":
            result = MINIMUMCPUSPEEDNOTMET_RESULTANTAPPSTATEDETAIL
        case "platformNotApplicable":
            result = PLATFORMNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
        case "fileSystemRequirementNotMet":
            result = FILESYSTEMREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
        case "registryRequirementNotMet":
            result = REGISTRYREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
        case "powerShellScriptRequirementNotMet":
            result = POWERSHELLSCRIPTREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
        case "supersedingAppsNotApplicable":
            result = SUPERSEDINGAPPSNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
        case "noAdditionalDetails":
            result = NOADDITIONALDETAILS_RESULTANTAPPSTATEDETAIL
        case "dependencyFailedToInstall":
            result = DEPENDENCYFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
        case "dependencyWithRequirementsNotMet":
            result = DEPENDENCYWITHREQUIREMENTSNOTMET_RESULTANTAPPSTATEDETAIL
        case "dependencyPendingReboot":
            result = DEPENDENCYPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "dependencyWithAutoInstallDisabled":
            result = DEPENDENCYWITHAUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
        case "supersededAppUninstallFailed":
            result = SUPERSEDEDAPPUNINSTALLFAILED_RESULTANTAPPSTATEDETAIL
        case "supersededAppUninstallPendingReboot":
            result = SUPERSEDEDAPPUNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "removingSupersededApps":
            result = REMOVINGSUPERSEDEDAPPS_RESULTANTAPPSTATEDETAIL
        case "iosAppStoreUpdateFailedToInstall":
            result = IOSAPPSTOREUPDATEFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
        case "vppAppHasUpdateAvailable":
            result = VPPAPPHASUPDATEAVAILABLE_RESULTANTAPPSTATEDETAIL
        case "userRejectedUpdate":
            result = USERREJECTEDUPDATE_RESULTANTAPPSTATEDETAIL
        case "uninstallPendingReboot":
            result = UNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "supersedingAppsDetected":
            result = SUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
        case "supersededAppsDetected":
            result = SUPERSEDEDAPPSDETECTED_RESULTANTAPPSTATEDETAIL
        case "seeInstallErrorCode":
            result = SEEINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
        case "autoInstallDisabled":
            result = AUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
        case "managedAppNoLongerPresent":
            result = MANAGEDAPPNOLONGERPRESENT_RESULTANTAPPSTATEDETAIL
        case "userRejectedInstall":
            result = USERREJECTEDINSTALL_RESULTANTAPPSTATEDETAIL
        case "userIsNotLoggedIntoAppStore":
            result = USERISNOTLOGGEDINTOAPPSTORE_RESULTANTAPPSTATEDETAIL
        case "untargetedSupersedingAppsDetected":
            result = UNTARGETEDSUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
        case "appRemovedBySupersedence":
            result = APPREMOVEDBYSUPERSEDENCE_RESULTANTAPPSTATEDETAIL
        case "seeUninstallErrorCode":
            result = SEEUNINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
        case "pendingReboot":
            result = PENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "installingDependencies":
            result = INSTALLINGDEPENDENCIES_RESULTANTAPPSTATEDETAIL
        case "contentDownloaded":
            result = CONTENTDOWNLOADED_RESULTANTAPPSTATEDETAIL
        default:
            return 0, errors.New("Unknown ResultantAppStateDetail value: " + v)
    }
    return &result, nil
}
func SerializeResultantAppStateDetail(values []ResultantAppStateDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
