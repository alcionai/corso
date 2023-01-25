package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManagementDerivedCredentialIssuer int

const (
    // Intercede
    INTERCEDE_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER DeviceManagementDerivedCredentialIssuer = iota
    // Entrust
    ENTRUSTDATACARD_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
    // Purebred
    PUREBRED_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
    // XTec
    XTEC_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
)

func (i DeviceManagementDerivedCredentialIssuer) String() string {
    return []string{"intercede", "entrustDatacard", "purebred", "xTec"}[i]
}
func ParseDeviceManagementDerivedCredentialIssuer(v string) (interface{}, error) {
    result := INTERCEDE_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
    switch v {
        case "intercede":
            result = INTERCEDE_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        case "entrustDatacard":
            result = ENTRUSTDATACARD_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        case "purebred":
            result = PUREBRED_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        case "xTec":
            result = XTEC_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        default:
            return 0, errors.New("Unknown DeviceManagementDerivedCredentialIssuer value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementDerivedCredentialIssuer(values []DeviceManagementDerivedCredentialIssuer) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
