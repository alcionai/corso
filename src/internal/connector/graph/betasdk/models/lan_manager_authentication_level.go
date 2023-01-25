package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type LanManagerAuthenticationLevel int

const (
    // Send LM & NTLM responses
    LMANDNLTM_LANMANAGERAUTHENTICATIONLEVEL LanManagerAuthenticationLevel = iota
    // Send LM & NTLM-use NTLMv2 session security if negotiated
    LMNTLMANDNTLMV2_LANMANAGERAUTHENTICATIONLEVEL
    // Send LM & NTLM responses only
    LMANDNTLMONLY_LANMANAGERAUTHENTICATIONLEVEL
    // Send LM & NTLMv2 responses only
    LMANDNTLMV2_LANMANAGERAUTHENTICATIONLEVEL
    // Send LM & NTLMv2 responses only. Refuse LM
    LMNTLMV2ANDNOTLM_LANMANAGERAUTHENTICATIONLEVEL
    // Send LM & NTLMv2 responses only. Refuse LM & NTLM
    LMNTLMV2ANDNOTLMORNTM_LANMANAGERAUTHENTICATIONLEVEL
)

func (i LanManagerAuthenticationLevel) String() string {
    return []string{"lmAndNltm", "lmNtlmAndNtlmV2", "lmAndNtlmOnly", "lmAndNtlmV2", "lmNtlmV2AndNotLm", "lmNtlmV2AndNotLmOrNtm"}[i]
}
func ParseLanManagerAuthenticationLevel(v string) (interface{}, error) {
    result := LMANDNLTM_LANMANAGERAUTHENTICATIONLEVEL
    switch v {
        case "lmAndNltm":
            result = LMANDNLTM_LANMANAGERAUTHENTICATIONLEVEL
        case "lmNtlmAndNtlmV2":
            result = LMNTLMANDNTLMV2_LANMANAGERAUTHENTICATIONLEVEL
        case "lmAndNtlmOnly":
            result = LMANDNTLMONLY_LANMANAGERAUTHENTICATIONLEVEL
        case "lmAndNtlmV2":
            result = LMANDNTLMV2_LANMANAGERAUTHENTICATIONLEVEL
        case "lmNtlmV2AndNotLm":
            result = LMNTLMV2ANDNOTLM_LANMANAGERAUTHENTICATIONLEVEL
        case "lmNtlmV2AndNotLmOrNtm":
            result = LMNTLMV2ANDNOTLMORNTM_LANMANAGERAUTHENTICATIONLEVEL
        default:
            return 0, errors.New("Unknown LanManagerAuthenticationLevel value: " + v)
    }
    return &result, nil
}
func SerializeLanManagerAuthenticationLevel(values []LanManagerAuthenticationLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
