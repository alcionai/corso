package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManagementCertificationAuthority int

const (
    // Not configured.
    NOTCONFIGURED_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY DeviceManagementCertificationAuthority = iota
    // Microsoft Certification Authority type.
    MICROSOFT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
    // DigiCert Certification Authority type.
    DIGICERT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
)

func (i DeviceManagementCertificationAuthority) String() string {
    return []string{"notConfigured", "microsoft", "digiCert"}[i]
}
func ParseDeviceManagementCertificationAuthority(v string) (interface{}, error) {
    result := NOTCONFIGURED_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
        case "microsoft":
            result = MICROSOFT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
        case "digiCert":
            result = DIGICERT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
        default:
            return 0, errors.New("Unknown DeviceManagementCertificationAuthority value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementCertificationAuthority(values []DeviceManagementCertificationAuthority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
