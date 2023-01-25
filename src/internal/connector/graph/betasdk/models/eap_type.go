package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EapType int

const (
    // EAP-Transport Layer Security (EAP-TLS).
    EAPTLS_EAPTYPE EapType = iota
    // Lightweight Extensible Authentication Protocol (LEAP).
    LEAP_EAPTYPE
    // EAP for GSM Subscriber Identity Module (EAP-SIM).
    EAPSIM_EAPTYPE
    // EAP-Tunneled Transport Layer Security (EAP-TTLS).
    EAPTTLS_EAPTYPE
    // Protected Extensible Authentication Protocol (PEAP).
    PEAP_EAPTYPE
    // EAP-Flexible Authentication via Secure Tunneling (EAP-FAST).
    EAPFAST_EAPTYPE
    // Tunnel Extensible Authentication Protocol (TEAP).
    TEAP_EAPTYPE
)

func (i EapType) String() string {
    return []string{"eapTls", "leap", "eapSim", "eapTtls", "peap", "eapFast", "teap"}[i]
}
func ParseEapType(v string) (interface{}, error) {
    result := EAPTLS_EAPTYPE
    switch v {
        case "eapTls":
            result = EAPTLS_EAPTYPE
        case "leap":
            result = LEAP_EAPTYPE
        case "eapSim":
            result = EAPSIM_EAPTYPE
        case "eapTtls":
            result = EAPTTLS_EAPTYPE
        case "peap":
            result = PEAP_EAPTYPE
        case "eapFast":
            result = EAPFAST_EAPTYPE
        case "teap":
            result = TEAP_EAPTYPE
        default:
            return 0, errors.New("Unknown EapType value: " + v)
    }
    return &result, nil
}
func SerializeEapType(values []EapType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
