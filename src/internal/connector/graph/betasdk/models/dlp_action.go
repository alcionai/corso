package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DlpAction int

const (
    NOTIFYUSER_DLPACTION DlpAction = iota
    BLOCKACCESS_DLPACTION
    DEVICERESTRICTION_DLPACTION
)

func (i DlpAction) String() string {
    return []string{"notifyUser", "blockAccess", "deviceRestriction"}[i]
}
func ParseDlpAction(v string) (interface{}, error) {
    result := NOTIFYUSER_DLPACTION
    switch v {
        case "notifyUser":
            result = NOTIFYUSER_DLPACTION
        case "blockAccess":
            result = BLOCKACCESS_DLPACTION
        case "deviceRestriction":
            result = DEVICERESTRICTION_DLPACTION
        default:
            return 0, errors.New("Unknown DlpAction value: " + v)
    }
    return &result, nil
}
func SerializeDlpAction(values []DlpAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
