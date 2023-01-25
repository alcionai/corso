package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DiffieHellmanGroup int

const (
    // Group1
    GROUP1_DIFFIEHELLMANGROUP DiffieHellmanGroup = iota
    // Group2
    GROUP2_DIFFIEHELLMANGROUP
    // Group14
    GROUP14_DIFFIEHELLMANGROUP
    // ECP256
    ECP256_DIFFIEHELLMANGROUP
    // ECP384
    ECP384_DIFFIEHELLMANGROUP
    // Group24
    GROUP24_DIFFIEHELLMANGROUP
)

func (i DiffieHellmanGroup) String() string {
    return []string{"group1", "group2", "group14", "ecp256", "ecp384", "group24"}[i]
}
func ParseDiffieHellmanGroup(v string) (interface{}, error) {
    result := GROUP1_DIFFIEHELLMANGROUP
    switch v {
        case "group1":
            result = GROUP1_DIFFIEHELLMANGROUP
        case "group2":
            result = GROUP2_DIFFIEHELLMANGROUP
        case "group14":
            result = GROUP14_DIFFIEHELLMANGROUP
        case "ecp256":
            result = ECP256_DIFFIEHELLMANGROUP
        case "ecp384":
            result = ECP384_DIFFIEHELLMANGROUP
        case "group24":
            result = GROUP24_DIFFIEHELLMANGROUP
        default:
            return 0, errors.New("Unknown DiffieHellmanGroup value: " + v)
    }
    return &result, nil
}
func SerializeDiffieHellmanGroup(values []DiffieHellmanGroup) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
