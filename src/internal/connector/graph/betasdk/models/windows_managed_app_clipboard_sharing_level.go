package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsManagedAppClipboardSharingLevel int

const (
    // Org users can paste data from and cut/copy data to any account, document, location or application.
    ANYDESTINATIONANYSOURCE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL WindowsManagedAppClipboardSharingLevel = iota
    // Org users cannot cut, copy or paste data to or from external accounts, documents, locations or applications from or into the org context.
    NONE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
)

func (i WindowsManagedAppClipboardSharingLevel) String() string {
    return []string{"anyDestinationAnySource", "none"}[i]
}
func ParseWindowsManagedAppClipboardSharingLevel(v string) (interface{}, error) {
    result := ANYDESTINATIONANYSOURCE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
    switch v {
        case "anyDestinationAnySource":
            result = ANYDESTINATIONANYSOURCE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
        case "none":
            result = NONE_WINDOWSMANAGEDAPPCLIPBOARDSHARINGLEVEL
        default:
            return 0, errors.New("Unknown WindowsManagedAppClipboardSharingLevel value: " + v)
    }
    return &result, nil
}
func SerializeWindowsManagedAppClipboardSharingLevel(values []WindowsManagedAppClipboardSharingLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
