package models
import (
    "errors"
)
// Provides operations to call the add method.
type Win32LobAppFileSystemDetectionType int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPFILESYSTEMDETECTIONTYPE Win32LobAppFileSystemDetectionType = iota
    // Whether the specified file or folder exists.
    EXISTS_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Last modified date.
    MODIFIEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Created date.
    CREATEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Version value type.
    VERSION_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Size detection type.
    SIZEINMB_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // The specified file or folder does not exist.
    DOESNOTEXIST_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
)

func (i Win32LobAppFileSystemDetectionType) String() string {
    return []string{"notConfigured", "exists", "modifiedDate", "createdDate", "version", "sizeInMB", "doesNotExist"}[i]
}
func ParseWin32LobAppFileSystemDetectionType(v string) (interface{}, error) {
    result := NOTCONFIGURED_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "exists":
            result = EXISTS_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "modifiedDate":
            result = MODIFIEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "createdDate":
            result = CREATEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "version":
            result = VERSION_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "sizeInMB":
            result = SIZEINMB_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "doesNotExist":
            result = DOESNOTEXIST_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        default:
            return 0, errors.New("Unknown Win32LobAppFileSystemDetectionType value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppFileSystemDetectionType(values []Win32LobAppFileSystemDetectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
