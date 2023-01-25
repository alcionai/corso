package models
import (
    "errors"
)
// Provides operations to call the add method.
type LocalSecurityOptionsMinimumSessionSecurity int

const (
    // Send LM & NTLM responses
    NONE_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY LocalSecurityOptionsMinimumSessionSecurity = iota
    // Send LM & NTLM-use NTLMv2 session security if negotiated
    REQUIRENTMLV2SESSIONSECURITY_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
    // Send LM & NTLM responses only
    REQUIRE128BITENCRYPTION_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
    // Send LM & NTLMv2 responses only
    NTLMV2AND128BITENCRYPTION_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
)

func (i LocalSecurityOptionsMinimumSessionSecurity) String() string {
    return []string{"none", "requireNtmlV2SessionSecurity", "require128BitEncryption", "ntlmV2And128BitEncryption"}[i]
}
func ParseLocalSecurityOptionsMinimumSessionSecurity(v string) (interface{}, error) {
    result := NONE_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
    switch v {
        case "none":
            result = NONE_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
        case "requireNtmlV2SessionSecurity":
            result = REQUIRENTMLV2SESSIONSECURITY_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
        case "require128BitEncryption":
            result = REQUIRE128BITENCRYPTION_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
        case "ntlmV2And128BitEncryption":
            result = NTLMV2AND128BITENCRYPTION_LOCALSECURITYOPTIONSMINIMUMSESSIONSECURITY
        default:
            return 0, errors.New("Unknown LocalSecurityOptionsMinimumSessionSecurity value: " + v)
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsMinimumSessionSecurity(values []LocalSecurityOptionsMinimumSessionSecurity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
