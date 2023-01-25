package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MacOSSoftwareUpdateBehavior int

const (
    // Not configured.
    NOTCONFIGURED_MACOSSOFTWAREUPDATEBEHAVIOR MacOSSoftwareUpdateBehavior = iota
    // Download and/or install the software update, depending on the current device state.
    DEFAULT_ESCAPED_MACOSSOFTWAREUPDATEBEHAVIOR
    // Download the software update without installing it.
    DOWNLOADONLY_MACOSSOFTWAREUPDATEBEHAVIOR
    // Install an already downloaded software update.
    INSTALLASAP_MACOSSOFTWAREUPDATEBEHAVIOR
    // Download the software update and notify the user via the App Store.
    NOTIFYONLY_MACOSSOFTWAREUPDATEBEHAVIOR
    // Download the software update and install it at a later time.
    INSTALLLATER_MACOSSOFTWAREUPDATEBEHAVIOR
)

func (i MacOSSoftwareUpdateBehavior) String() string {
    return []string{"notConfigured", "default", "downloadOnly", "installASAP", "notifyOnly", "installLater"}[i]
}
func ParseMacOSSoftwareUpdateBehavior(v string) (interface{}, error) {
    result := NOTCONFIGURED_MACOSSOFTWAREUPDATEBEHAVIOR
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MACOSSOFTWAREUPDATEBEHAVIOR
        case "default":
            result = DEFAULT_ESCAPED_MACOSSOFTWAREUPDATEBEHAVIOR
        case "downloadOnly":
            result = DOWNLOADONLY_MACOSSOFTWAREUPDATEBEHAVIOR
        case "installASAP":
            result = INSTALLASAP_MACOSSOFTWAREUPDATEBEHAVIOR
        case "notifyOnly":
            result = NOTIFYONLY_MACOSSOFTWAREUPDATEBEHAVIOR
        case "installLater":
            result = INSTALLLATER_MACOSSOFTWAREUPDATEBEHAVIOR
        default:
            return 0, errors.New("Unknown MacOSSoftwareUpdateBehavior value: " + v)
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateBehavior(values []MacOSSoftwareUpdateBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
