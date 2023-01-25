package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type MicrosoftStoreForBusinessPortalSelectionOptions int

const (
    // This option is not available for the account
    NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS MicrosoftStoreForBusinessPortalSelectionOptions = iota
    // Intune Company Portal only.
    COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
    // MSFB Private store only.
    PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
)

func (i MicrosoftStoreForBusinessPortalSelectionOptions) String() string {
    return []string{"none", "companyPortal", "privateStore"}[i]
}
func ParseMicrosoftStoreForBusinessPortalSelectionOptions(v string) (interface{}, error) {
    result := NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
    switch v {
        case "none":
            result = NONE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
        case "companyPortal":
            result = COMPANYPORTAL_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
        case "privateStore":
            result = PRIVATESTORE_MICROSOFTSTOREFORBUSINESSPORTALSELECTIONOPTIONS
        default:
            return 0, errors.New("Unknown MicrosoftStoreForBusinessPortalSelectionOptions value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftStoreForBusinessPortalSelectionOptions(values []MicrosoftStoreForBusinessPortalSelectionOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
