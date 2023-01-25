package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type PhysicalAddressType int

const (
    UNKNOWN_PHYSICALADDRESSTYPE PhysicalAddressType = iota
    HOME_PHYSICALADDRESSTYPE
    BUSINESS_PHYSICALADDRESSTYPE
    OTHER_PHYSICALADDRESSTYPE
)

func (i PhysicalAddressType) String() string {
    return []string{"unknown", "home", "business", "other"}[i]
}
func ParsePhysicalAddressType(v string) (interface{}, error) {
    result := UNKNOWN_PHYSICALADDRESSTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_PHYSICALADDRESSTYPE
        case "home":
            result = HOME_PHYSICALADDRESSTYPE
        case "business":
            result = BUSINESS_PHYSICALADDRESSTYPE
        case "other":
            result = OTHER_PHYSICALADDRESSTYPE
        default:
            return 0, errors.New("Unknown PhysicalAddressType value: " + v)
    }
    return &result, nil
}
func SerializePhysicalAddressType(values []PhysicalAddressType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
