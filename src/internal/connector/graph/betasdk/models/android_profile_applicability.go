package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidProfileApplicability int

const (
    DEFAULT_ESCAPED_ANDROIDPROFILEAPPLICABILITY AndroidProfileApplicability = iota
    ANDROIDWORKPROFILE_ANDROIDPROFILEAPPLICABILITY
    ANDROIDDEVICEOWNER_ANDROIDPROFILEAPPLICABILITY
)

func (i AndroidProfileApplicability) String() string {
    return []string{"default", "androidWorkProfile", "androidDeviceOwner"}[i]
}
func ParseAndroidProfileApplicability(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_ANDROIDPROFILEAPPLICABILITY
    switch v {
        case "default":
            result = DEFAULT_ESCAPED_ANDROIDPROFILEAPPLICABILITY
        case "androidWorkProfile":
            result = ANDROIDWORKPROFILE_ANDROIDPROFILEAPPLICABILITY
        case "androidDeviceOwner":
            result = ANDROIDDEVICEOWNER_ANDROIDPROFILEAPPLICABILITY
        default:
            return 0, errors.New("Unknown AndroidProfileApplicability value: " + v)
    }
    return &result, nil
}
func SerializeAndroidProfileApplicability(values []AndroidProfileApplicability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
