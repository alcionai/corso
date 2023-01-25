package models
import (
    "errors"
)
// Provides operations to call the add method.
type ITunesPairingMode int

const (
    // Pairing is not allowed
    DISALLOW_ITUNESPAIRINGMODE ITunesPairingMode = iota
    // Pairing allowed
    ALLOW_ITUNESPAIRINGMODE
    // Certificate required to pair with iTunes
    REQUIRESCERTIFICATE_ITUNESPAIRINGMODE
)

func (i ITunesPairingMode) String() string {
    return []string{"disallow", "allow", "requiresCertificate"}[i]
}
func ParseITunesPairingMode(v string) (interface{}, error) {
    result := DISALLOW_ITUNESPAIRINGMODE
    switch v {
        case "disallow":
            result = DISALLOW_ITUNESPAIRINGMODE
        case "allow":
            result = ALLOW_ITUNESPAIRINGMODE
        case "requiresCertificate":
            result = REQUIRESCERTIFICATE_ITUNESPAIRINGMODE
        default:
            return 0, errors.New("Unknown ITunesPairingMode value: " + v)
    }
    return &result, nil
}
func SerializeITunesPairingMode(values []ITunesPairingMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
