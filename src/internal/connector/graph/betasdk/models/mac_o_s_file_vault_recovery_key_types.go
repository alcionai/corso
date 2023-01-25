package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MacOSFileVaultRecoveryKeyTypes int

const (
    // Device default value, no intent.
    NOTCONFIGURED_MACOSFILEVAULTRECOVERYKEYTYPES MacOSFileVaultRecoveryKeyTypes = iota
    // An institutional recovery key is like a “master” recovery key that can be used to unlock any device whose password has been lost.
    INSTITUTIONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES
    // A personal recovery key is a unique code that can be used to unlock the user’s device, even if the password to the device is lost.
    PERSONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES
)

func (i MacOSFileVaultRecoveryKeyTypes) String() string {
    return []string{"notConfigured", "institutionalRecoveryKey", "personalRecoveryKey"}[i]
}
func ParseMacOSFileVaultRecoveryKeyTypes(v string) (interface{}, error) {
    result := NOTCONFIGURED_MACOSFILEVAULTRECOVERYKEYTYPES
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MACOSFILEVAULTRECOVERYKEYTYPES
        case "institutionalRecoveryKey":
            result = INSTITUTIONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES
        case "personalRecoveryKey":
            result = PERSONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES
        default:
            return 0, errors.New("Unknown MacOSFileVaultRecoveryKeyTypes value: " + v)
    }
    return &result, nil
}
func SerializeMacOSFileVaultRecoveryKeyTypes(values []MacOSFileVaultRecoveryKeyTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
