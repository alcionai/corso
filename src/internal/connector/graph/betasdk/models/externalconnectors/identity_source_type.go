package externalconnectors
import (
    "errors"
)
// Provides operations to call the add method.
type IdentitySourceType int

const (
    AZUREACTIVEDIRECTORY_IDENTITYSOURCETYPE IdentitySourceType = iota
    EXTERNAL_IDENTITYSOURCETYPE
    UNKNOWNFUTUREVALUE_IDENTITYSOURCETYPE
)

func (i IdentitySourceType) String() string {
    return []string{"azureActiveDirectory", "external", "unknownFutureValue"}[i]
}
func ParseIdentitySourceType(v string) (interface{}, error) {
    result := AZUREACTIVEDIRECTORY_IDENTITYSOURCETYPE
    switch v {
        case "azureActiveDirectory":
            result = AZUREACTIVEDIRECTORY_IDENTITYSOURCETYPE
        case "external":
            result = EXTERNAL_IDENTITYSOURCETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IDENTITYSOURCETYPE
        default:
            return 0, errors.New("Unknown IdentitySourceType value: " + v)
    }
    return &result, nil
}
func SerializeIdentitySourceType(values []IdentitySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
