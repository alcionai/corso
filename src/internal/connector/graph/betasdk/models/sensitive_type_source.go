package models
import (
    "errors"
)
// Provides operations to call the add method.
type SensitiveTypeSource int

const (
    OUTOFBOX_SENSITIVETYPESOURCE SensitiveTypeSource = iota
    TENANT_SENSITIVETYPESOURCE
)

func (i SensitiveTypeSource) String() string {
    return []string{"outOfBox", "tenant"}[i]
}
func ParseSensitiveTypeSource(v string) (interface{}, error) {
    result := OUTOFBOX_SENSITIVETYPESOURCE
    switch v {
        case "outOfBox":
            result = OUTOFBOX_SENSITIVETYPESOURCE
        case "tenant":
            result = TENANT_SENSITIVETYPESOURCE
        default:
            return 0, errors.New("Unknown SensitiveTypeSource value: " + v)
    }
    return &result, nil
}
func SerializeSensitiveTypeSource(values []SensitiveTypeSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
