package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ImportedDeviceIdentityType int

const (
    // Unknown value of importedDeviceIdentityType.
    UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE ImportedDeviceIdentityType = iota
    // Device Identity is of type imei.
    IMEI_IMPORTEDDEVICEIDENTITYTYPE
    // Device Identity is of type serial number.
    SERIALNUMBER_IMPORTEDDEVICEIDENTITYTYPE
)

func (i ImportedDeviceIdentityType) String() string {
    return []string{"unknown", "imei", "serialNumber"}[i]
}
func ParseImportedDeviceIdentityType(v string) (interface{}, error) {
    result := UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE
        case "imei":
            result = IMEI_IMPORTEDDEVICEIDENTITYTYPE
        case "serialNumber":
            result = SERIALNUMBER_IMPORTEDDEVICEIDENTITYTYPE
        default:
            return 0, errors.New("Unknown ImportedDeviceIdentityType value: " + v)
    }
    return &result, nil
}
func SerializeImportedDeviceIdentityType(values []ImportedDeviceIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
