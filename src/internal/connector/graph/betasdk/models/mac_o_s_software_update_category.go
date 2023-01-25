package models
import (
    "errors"
)
// Provides operations to call the add method.
type MacOSSoftwareUpdateCategory int

const (
    // A critical update
    CRITICAL_MACOSSOFTWAREUPDATECATEGORY MacOSSoftwareUpdateCategory = iota
    // A configuration data file update
    CONFIGURATIONDATAFILE_MACOSSOFTWAREUPDATECATEGORY
    // A firmware update
    FIRMWARE_MACOSSOFTWAREUPDATECATEGORY
    // All other update types
    OTHER_MACOSSOFTWAREUPDATECATEGORY
)

func (i MacOSSoftwareUpdateCategory) String() string {
    return []string{"critical", "configurationDataFile", "firmware", "other"}[i]
}
func ParseMacOSSoftwareUpdateCategory(v string) (interface{}, error) {
    result := CRITICAL_MACOSSOFTWAREUPDATECATEGORY
    switch v {
        case "critical":
            result = CRITICAL_MACOSSOFTWAREUPDATECATEGORY
        case "configurationDataFile":
            result = CONFIGURATIONDATAFILE_MACOSSOFTWAREUPDATECATEGORY
        case "firmware":
            result = FIRMWARE_MACOSSOFTWAREUPDATECATEGORY
        case "other":
            result = OTHER_MACOSSOFTWAREUPDATECATEGORY
        default:
            return 0, errors.New("Unknown MacOSSoftwareUpdateCategory value: " + v)
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateCategory(values []MacOSSoftwareUpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
