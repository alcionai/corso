package models
import (
    "errors"
)
// Provides operations to call the add method.
type FolderProtectionType int

const (
    // Device default value, no intent.
    USERDEFINED_FOLDERPROTECTIONTYPE FolderProtectionType = iota
    // Block functionality.
    ENABLE_FOLDERPROTECTIONTYPE
    // Allow functionality but generate logs.
    AUDITMODE_FOLDERPROTECTIONTYPE
    // Block untrusted apps from writing to disk sectors.
    BLOCKDISKMODIFICATION_FOLDERPROTECTIONTYPE
    // Generate logs when untrusted apps write to disk sectors.
    AUDITDISKMODIFICATION_FOLDERPROTECTIONTYPE
)

func (i FolderProtectionType) String() string {
    return []string{"userDefined", "enable", "auditMode", "blockDiskModification", "auditDiskModification"}[i]
}
func ParseFolderProtectionType(v string) (interface{}, error) {
    result := USERDEFINED_FOLDERPROTECTIONTYPE
    switch v {
        case "userDefined":
            result = USERDEFINED_FOLDERPROTECTIONTYPE
        case "enable":
            result = ENABLE_FOLDERPROTECTIONTYPE
        case "auditMode":
            result = AUDITMODE_FOLDERPROTECTIONTYPE
        case "blockDiskModification":
            result = BLOCKDISKMODIFICATION_FOLDERPROTECTIONTYPE
        case "auditDiskModification":
            result = AUDITDISKMODIFICATION_FOLDERPROTECTIONTYPE
        default:
            return 0, errors.New("Unknown FolderProtectionType value: " + v)
    }
    return &result, nil
}
func SerializeFolderProtectionType(values []FolderProtectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
