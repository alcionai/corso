package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type OfficeProductId int

const (
    O365PROPLUSRETAIL_OFFICEPRODUCTID OfficeProductId = iota
    O365BUSINESSRETAIL_OFFICEPRODUCTID
    VISIOPRORETAIL_OFFICEPRODUCTID
    PROJECTPRORETAIL_OFFICEPRODUCTID
)

func (i OfficeProductId) String() string {
    return []string{"o365ProPlusRetail", "o365BusinessRetail", "visioProRetail", "projectProRetail"}[i]
}
func ParseOfficeProductId(v string) (interface{}, error) {
    result := O365PROPLUSRETAIL_OFFICEPRODUCTID
    switch v {
        case "o365ProPlusRetail":
            result = O365PROPLUSRETAIL_OFFICEPRODUCTID
        case "o365BusinessRetail":
            result = O365BUSINESSRETAIL_OFFICEPRODUCTID
        case "visioProRetail":
            result = VISIOPRORETAIL_OFFICEPRODUCTID
        case "projectProRetail":
            result = PROJECTPRORETAIL_OFFICEPRODUCTID
        default:
            return 0, errors.New("Unknown OfficeProductId value: " + v)
    }
    return &result, nil
}
func SerializeOfficeProductId(values []OfficeProductId) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
