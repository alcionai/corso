package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MobileThreatDefensePartnerPriority int

const (
    // Indicates use of Microsoft Defender Endpoint over 3rd party MTD connectors
    DEFENDEROVERTHIRDPARTYPARTNER_MOBILETHREATDEFENSEPARTNERPRIORITY MobileThreatDefensePartnerPriority = iota
    // Indicates use of a 3rd party MTD connector over Microsoft Defender Endpoint
    THIRDPARTYPARTNEROVERDEFENDER_MOBILETHREATDEFENSEPARTNERPRIORITY
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MOBILETHREATDEFENSEPARTNERPRIORITY
)

func (i MobileThreatDefensePartnerPriority) String() string {
    return []string{"defenderOverThirdPartyPartner", "thirdPartyPartnerOverDefender", "unknownFutureValue"}[i]
}
func ParseMobileThreatDefensePartnerPriority(v string) (interface{}, error) {
    result := DEFENDEROVERTHIRDPARTYPARTNER_MOBILETHREATDEFENSEPARTNERPRIORITY
    switch v {
        case "defenderOverThirdPartyPartner":
            result = DEFENDEROVERTHIRDPARTYPARTNER_MOBILETHREATDEFENSEPARTNERPRIORITY
        case "thirdPartyPartnerOverDefender":
            result = THIRDPARTYPARTNEROVERDEFENDER_MOBILETHREATDEFENSEPARTNERPRIORITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MOBILETHREATDEFENSEPARTNERPRIORITY
        default:
            return 0, errors.New("Unknown MobileThreatDefensePartnerPriority value: " + v)
    }
    return &result, nil
}
func SerializeMobileThreatDefensePartnerPriority(values []MobileThreatDefensePartnerPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
