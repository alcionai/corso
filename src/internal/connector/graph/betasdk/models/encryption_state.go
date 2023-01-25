package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type EncryptionState int

const (
    // Not encrypted
    NOTENCRYPTED_ENCRYPTIONSTATE EncryptionState = iota
    // Encrypted
    ENCRYPTED_ENCRYPTIONSTATE
)

func (i EncryptionState) String() string {
    return []string{"notEncrypted", "encrypted"}[i]
}
func ParseEncryptionState(v string) (interface{}, error) {
    result := NOTENCRYPTED_ENCRYPTIONSTATE
    switch v {
        case "notEncrypted":
            result = NOTENCRYPTED_ENCRYPTIONSTATE
        case "encrypted":
            result = ENCRYPTED_ENCRYPTIONSTATE
        default:
            return 0, errors.New("Unknown EncryptionState value: " + v)
    }
    return &result, nil
}
func SerializeEncryptionState(values []EncryptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
