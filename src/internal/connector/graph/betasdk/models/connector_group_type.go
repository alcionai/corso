package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ConnectorGroupType int

const (
    APPLICATIONPROXY_CONNECTORGROUPTYPE ConnectorGroupType = iota
)

func (i ConnectorGroupType) String() string {
    return []string{"applicationProxy"}[i]
}
func ParseConnectorGroupType(v string) (interface{}, error) {
    result := APPLICATIONPROXY_CONNECTORGROUPTYPE
    switch v {
        case "applicationProxy":
            result = APPLICATIONPROXY_CONNECTORGROUPTYPE
        default:
            return 0, errors.New("Unknown ConnectorGroupType value: " + v)
    }
    return &result, nil
}
func SerializeConnectorGroupType(values []ConnectorGroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
