package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceManagementConfigurationWindowsSkus int

const (
    UNKNOWN_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS DeviceManagementConfigurationWindowsSkus = iota
    WINDOWSHOME_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    WINDOWSPROFESSIONAL_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    WINDOWSENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    WINDOWSEDUCATION_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    WINDOWSMOBILE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    WINDOWSMOBILEENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    WINDOWSTEAMSURFACE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    IOT_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    IOTENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    HOLOLENS_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    HOLOLENSENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    HOLOGRAPHICFORBUSINESS_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    WINDOWSMULTISESSION_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    SURFACEHUB_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
)

func (i DeviceManagementConfigurationWindowsSkus) String() string {
    return []string{"unknown", "windowsHome", "windowsProfessional", "windowsEnterprise", "windowsEducation", "windowsMobile", "windowsMobileEnterprise", "windowsTeamSurface", "iot", "iotEnterprise", "holoLens", "holoLensEnterprise", "holographicForBusiness", "windowsMultiSession", "surfaceHub"}[i]
}
func ParseDeviceManagementConfigurationWindowsSkus(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsHome":
            result = WINDOWSHOME_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsProfessional":
            result = WINDOWSPROFESSIONAL_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsEnterprise":
            result = WINDOWSENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsEducation":
            result = WINDOWSEDUCATION_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsMobile":
            result = WINDOWSMOBILE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsMobileEnterprise":
            result = WINDOWSMOBILEENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsTeamSurface":
            result = WINDOWSTEAMSURFACE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "iot":
            result = IOT_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "iotEnterprise":
            result = IOTENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "holoLens":
            result = HOLOLENS_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "holoLensEnterprise":
            result = HOLOLENSENTERPRISE_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "holographicForBusiness":
            result = HOLOGRAPHICFORBUSINESS_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "windowsMultiSession":
            result = WINDOWSMULTISESSION_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        case "surfaceHub":
            result = SURFACEHUB_DEVICEMANAGEMENTCONFIGURATIONWINDOWSSKUS
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationWindowsSkus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationWindowsSkus(values []DeviceManagementConfigurationWindowsSkus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
