package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PerfectForwardSecrecyGroup int

const (
    // PFS1
    PFS1_PERFECTFORWARDSECRECYGROUP PerfectForwardSecrecyGroup = iota
    // PFS2
    PFS2_PERFECTFORWARDSECRECYGROUP
    // PFS2048
    PFS2048_PERFECTFORWARDSECRECYGROUP
    // ECP256
    ECP256_PERFECTFORWARDSECRECYGROUP
    // ECP384
    ECP384_PERFECTFORWARDSECRECYGROUP
    // PFSMM
    PFSMM_PERFECTFORWARDSECRECYGROUP
    // PFS24
    PFS24_PERFECTFORWARDSECRECYGROUP
)

func (i PerfectForwardSecrecyGroup) String() string {
    return []string{"pfs1", "pfs2", "pfs2048", "ecp256", "ecp384", "pfsMM", "pfs24"}[i]
}
func ParsePerfectForwardSecrecyGroup(v string) (interface{}, error) {
    result := PFS1_PERFECTFORWARDSECRECYGROUP
    switch v {
        case "pfs1":
            result = PFS1_PERFECTFORWARDSECRECYGROUP
        case "pfs2":
            result = PFS2_PERFECTFORWARDSECRECYGROUP
        case "pfs2048":
            result = PFS2048_PERFECTFORWARDSECRECYGROUP
        case "ecp256":
            result = ECP256_PERFECTFORWARDSECRECYGROUP
        case "ecp384":
            result = ECP384_PERFECTFORWARDSECRECYGROUP
        case "pfsMM":
            result = PFSMM_PERFECTFORWARDSECRECYGROUP
        case "pfs24":
            result = PFS24_PERFECTFORWARDSECRECYGROUP
        default:
            return 0, errors.New("Unknown PerfectForwardSecrecyGroup value: " + v)
    }
    return &result, nil
}
func SerializePerfectForwardSecrecyGroup(values []PerfectForwardSecrecyGroup) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
