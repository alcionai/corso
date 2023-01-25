package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type GroupPolicySettingType int

const (
    // GroupPolicySettingType unknown
    UNKNOWN_GROUPPOLICYSETTINGTYPE GroupPolicySettingType = iota
    // Policy setting type
    POLICY_GROUPPOLICYSETTINGTYPE
    // Account setting type
    ACCOUNT_GROUPPOLICYSETTINGTYPE
    // SecurityOptions setting type
    SECURITYOPTIONS_GROUPPOLICYSETTINGTYPE
    // UserRightsAssignment setting type
    USERRIGHTSASSIGNMENT_GROUPPOLICYSETTINGTYPE
    // AuditSetting setting type
    AUDITSETTING_GROUPPOLICYSETTINGTYPE
    // WindowsFirewallSettings setting type
    WINDOWSFIREWALLSETTINGS_GROUPPOLICYSETTINGTYPE
    // AppLockerRuleCollection setting type
    APPLOCKERRULECOLLECTION_GROUPPOLICYSETTINGTYPE
    // DataSourcesSettings setting type
    DATASOURCESSETTINGS_GROUPPOLICYSETTINGTYPE
    // DevicesSettings setting type
    DEVICESSETTINGS_GROUPPOLICYSETTINGTYPE
    // DriveMapSettings setting type
    DRIVEMAPSETTINGS_GROUPPOLICYSETTINGTYPE
    // EnvironmentVariables setting type
    ENVIRONMENTVARIABLES_GROUPPOLICYSETTINGTYPE
    // FilesSettings setting type
    FILESSETTINGS_GROUPPOLICYSETTINGTYPE
    // FolderOptions setting type
    FOLDEROPTIONS_GROUPPOLICYSETTINGTYPE
    // Folders setting type
    FOLDERS_GROUPPOLICYSETTINGTYPE
    // IniFiles setting type
    INIFILES_GROUPPOLICYSETTINGTYPE
    // InternetOptions setting type
    INTERNETOPTIONS_GROUPPOLICYSETTINGTYPE
    // LocalUsersAndGroups setting type
    LOCALUSERSANDGROUPS_GROUPPOLICYSETTINGTYPE
    // NetworkOptions setting type
    NETWORKOPTIONS_GROUPPOLICYSETTINGTYPE
    // NetworkShares setting type
    NETWORKSHARES_GROUPPOLICYSETTINGTYPE
    // NTServices setting type
    NTSERVICES_GROUPPOLICYSETTINGTYPE
    // PowerOptions setting type
    POWEROPTIONS_GROUPPOLICYSETTINGTYPE
    // Printers setting type
    PRINTERS_GROUPPOLICYSETTINGTYPE
    // RegionalOptionsSettings setting type
    REGIONALOPTIONSSETTINGS_GROUPPOLICYSETTINGTYPE
    // RegistrySettings setting type
    REGISTRYSETTINGS_GROUPPOLICYSETTINGTYPE
    // ScheduledTasks setting type
    SCHEDULEDTASKS_GROUPPOLICYSETTINGTYPE
    // ShortcutSettings setting type
    SHORTCUTSETTINGS_GROUPPOLICYSETTINGTYPE
    // StartMenuSettings setting type
    STARTMENUSETTINGS_GROUPPOLICYSETTINGTYPE
)

func (i GroupPolicySettingType) String() string {
    return []string{"unknown", "policy", "account", "securityOptions", "userRightsAssignment", "auditSetting", "windowsFirewallSettings", "appLockerRuleCollection", "dataSourcesSettings", "devicesSettings", "driveMapSettings", "environmentVariables", "filesSettings", "folderOptions", "folders", "iniFiles", "internetOptions", "localUsersAndGroups", "networkOptions", "networkShares", "ntServices", "powerOptions", "printers", "regionalOptionsSettings", "registrySettings", "scheduledTasks", "shortcutSettings", "startMenuSettings"}[i]
}
func ParseGroupPolicySettingType(v string) (interface{}, error) {
    result := UNKNOWN_GROUPPOLICYSETTINGTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_GROUPPOLICYSETTINGTYPE
        case "policy":
            result = POLICY_GROUPPOLICYSETTINGTYPE
        case "account":
            result = ACCOUNT_GROUPPOLICYSETTINGTYPE
        case "securityOptions":
            result = SECURITYOPTIONS_GROUPPOLICYSETTINGTYPE
        case "userRightsAssignment":
            result = USERRIGHTSASSIGNMENT_GROUPPOLICYSETTINGTYPE
        case "auditSetting":
            result = AUDITSETTING_GROUPPOLICYSETTINGTYPE
        case "windowsFirewallSettings":
            result = WINDOWSFIREWALLSETTINGS_GROUPPOLICYSETTINGTYPE
        case "appLockerRuleCollection":
            result = APPLOCKERRULECOLLECTION_GROUPPOLICYSETTINGTYPE
        case "dataSourcesSettings":
            result = DATASOURCESSETTINGS_GROUPPOLICYSETTINGTYPE
        case "devicesSettings":
            result = DEVICESSETTINGS_GROUPPOLICYSETTINGTYPE
        case "driveMapSettings":
            result = DRIVEMAPSETTINGS_GROUPPOLICYSETTINGTYPE
        case "environmentVariables":
            result = ENVIRONMENTVARIABLES_GROUPPOLICYSETTINGTYPE
        case "filesSettings":
            result = FILESSETTINGS_GROUPPOLICYSETTINGTYPE
        case "folderOptions":
            result = FOLDEROPTIONS_GROUPPOLICYSETTINGTYPE
        case "folders":
            result = FOLDERS_GROUPPOLICYSETTINGTYPE
        case "iniFiles":
            result = INIFILES_GROUPPOLICYSETTINGTYPE
        case "internetOptions":
            result = INTERNETOPTIONS_GROUPPOLICYSETTINGTYPE
        case "localUsersAndGroups":
            result = LOCALUSERSANDGROUPS_GROUPPOLICYSETTINGTYPE
        case "networkOptions":
            result = NETWORKOPTIONS_GROUPPOLICYSETTINGTYPE
        case "networkShares":
            result = NETWORKSHARES_GROUPPOLICYSETTINGTYPE
        case "ntServices":
            result = NTSERVICES_GROUPPOLICYSETTINGTYPE
        case "powerOptions":
            result = POWEROPTIONS_GROUPPOLICYSETTINGTYPE
        case "printers":
            result = PRINTERS_GROUPPOLICYSETTINGTYPE
        case "regionalOptionsSettings":
            result = REGIONALOPTIONSSETTINGS_GROUPPOLICYSETTINGTYPE
        case "registrySettings":
            result = REGISTRYSETTINGS_GROUPPOLICYSETTINGTYPE
        case "scheduledTasks":
            result = SCHEDULEDTASKS_GROUPPOLICYSETTINGTYPE
        case "shortcutSettings":
            result = SHORTCUTSETTINGS_GROUPPOLICYSETTINGTYPE
        case "startMenuSettings":
            result = STARTMENUSETTINGS_GROUPPOLICYSETTINGTYPE
        default:
            return 0, errors.New("Unknown GroupPolicySettingType value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicySettingType(values []GroupPolicySettingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
