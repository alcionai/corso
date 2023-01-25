package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcPartnerAgentInstallStatus int

const (
    INSTALLED_CLOUDPCPARTNERAGENTINSTALLSTATUS CloudPcPartnerAgentInstallStatus = iota
    INSTALLFAILED_CLOUDPCPARTNERAGENTINSTALLSTATUS
    INSTALLING_CLOUDPCPARTNERAGENTINSTALLSTATUS
    UNINSTALLING_CLOUDPCPARTNERAGENTINSTALLSTATUS
    UNINSTALLFAILED_CLOUDPCPARTNERAGENTINSTALLSTATUS
    LICENSED_CLOUDPCPARTNERAGENTINSTALLSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCPARTNERAGENTINSTALLSTATUS
)

func (i CloudPcPartnerAgentInstallStatus) String() string {
    return []string{"installed", "installFailed", "installing", "uninstalling", "uninstallFailed", "licensed", "unknownFutureValue"}[i]
}
func ParseCloudPcPartnerAgentInstallStatus(v string) (interface{}, error) {
    result := INSTALLED_CLOUDPCPARTNERAGENTINSTALLSTATUS
    switch v {
        case "installed":
            result = INSTALLED_CLOUDPCPARTNERAGENTINSTALLSTATUS
        case "installFailed":
            result = INSTALLFAILED_CLOUDPCPARTNERAGENTINSTALLSTATUS
        case "installing":
            result = INSTALLING_CLOUDPCPARTNERAGENTINSTALLSTATUS
        case "uninstalling":
            result = UNINSTALLING_CLOUDPCPARTNERAGENTINSTALLSTATUS
        case "uninstallFailed":
            result = UNINSTALLFAILED_CLOUDPCPARTNERAGENTINSTALLSTATUS
        case "licensed":
            result = LICENSED_CLOUDPCPARTNERAGENTINSTALLSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPARTNERAGENTINSTALLSTATUS
        default:
            return 0, errors.New("Unknown CloudPcPartnerAgentInstallStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcPartnerAgentInstallStatus(values []CloudPcPartnerAgentInstallStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
