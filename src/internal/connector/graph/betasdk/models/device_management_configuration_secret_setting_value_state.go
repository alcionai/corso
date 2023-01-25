package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManagementConfigurationSecretSettingValueState int

const (
    // default invalid value
    INVALID_DEVICEMANAGEMENTCONFIGURATIONSECRETSETTINGVALUESTATE DeviceManagementConfigurationSecretSettingValueState = iota
    // secret value is not encrypted
    NOTENCRYPTED_DEVICEMANAGEMENTCONFIGURATIONSECRETSETTINGVALUESTATE
    // a token for the encrypted value is returned by the service
    ENCRYPTEDVALUETOKEN_DEVICEMANAGEMENTCONFIGURATIONSECRETSETTINGVALUESTATE
)

func (i DeviceManagementConfigurationSecretSettingValueState) String() string {
    return []string{"invalid", "notEncrypted", "encryptedValueToken"}[i]
}
func ParseDeviceManagementConfigurationSecretSettingValueState(v string) (interface{}, error) {
    result := INVALID_DEVICEMANAGEMENTCONFIGURATIONSECRETSETTINGVALUESTATE
    switch v {
        case "invalid":
            result = INVALID_DEVICEMANAGEMENTCONFIGURATIONSECRETSETTINGVALUESTATE
        case "notEncrypted":
            result = NOTENCRYPTED_DEVICEMANAGEMENTCONFIGURATIONSECRETSETTINGVALUESTATE
        case "encryptedValueToken":
            result = ENCRYPTEDVALUETOKEN_DEVICEMANAGEMENTCONFIGURATIONSECRETSETTINGVALUESTATE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationSecretSettingValueState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSecretSettingValueState(values []DeviceManagementConfigurationSecretSettingValueState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
