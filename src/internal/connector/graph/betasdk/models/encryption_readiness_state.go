package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EncryptionReadinessState int

const (
    // Not ready
    NOTREADY_ENCRYPTIONREADINESSSTATE EncryptionReadinessState = iota
    // Ready
    READY_ENCRYPTIONREADINESSSTATE
)

func (i EncryptionReadinessState) String() string {
    return []string{"notReady", "ready"}[i]
}
func ParseEncryptionReadinessState(v string) (interface{}, error) {
    result := NOTREADY_ENCRYPTIONREADINESSSTATE
    switch v {
        case "notReady":
            result = NOTREADY_ENCRYPTIONREADINESSSTATE
        case "ready":
            result = READY_ENCRYPTIONREADINESSSTATE
        default:
            return 0, errors.New("Unknown EncryptionReadinessState value: " + v)
    }
    return &result, nil
}
func SerializeEncryptionReadinessState(values []EncryptionReadinessState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
