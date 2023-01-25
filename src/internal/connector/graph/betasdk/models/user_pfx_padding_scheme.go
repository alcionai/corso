package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type UserPfxPaddingScheme int

const (
    // Unknown padding Scheme.
    NONE_USERPFXPADDINGSCHEME UserPfxPaddingScheme = iota
    // Pkcs1 is no longer supported
    PKCS1_USERPFXPADDINGSCHEME
    // OaepSha1 is no longer supported
    OAEPSHA1_USERPFXPADDINGSCHEME
    // Use OAEP SHA-256 padding.
    OAEPSHA256_USERPFXPADDINGSCHEME
    // Use OAEP SHA-384 padding.
    OAEPSHA384_USERPFXPADDINGSCHEME
    // Use OAEP SHA-512 padding.
    OAEPSHA512_USERPFXPADDINGSCHEME
)

func (i UserPfxPaddingScheme) String() string {
    return []string{"none", "pkcs1", "oaepSha1", "oaepSha256", "oaepSha384", "oaepSha512"}[i]
}
func ParseUserPfxPaddingScheme(v string) (interface{}, error) {
    result := NONE_USERPFXPADDINGSCHEME
    switch v {
        case "none":
            result = NONE_USERPFXPADDINGSCHEME
        case "pkcs1":
            result = PKCS1_USERPFXPADDINGSCHEME
        case "oaepSha1":
            result = OAEPSHA1_USERPFXPADDINGSCHEME
        case "oaepSha256":
            result = OAEPSHA256_USERPFXPADDINGSCHEME
        case "oaepSha384":
            result = OAEPSHA384_USERPFXPADDINGSCHEME
        case "oaepSha512":
            result = OAEPSHA512_USERPFXPADDINGSCHEME
        default:
            return 0, errors.New("Unknown UserPfxPaddingScheme value: " + v)
    }
    return &result, nil
}
func SerializeUserPfxPaddingScheme(values []UserPfxPaddingScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
