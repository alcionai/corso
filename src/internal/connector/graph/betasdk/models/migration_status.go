package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type MigrationStatus int

const (
    READY_MIGRATIONSTATUS MigrationStatus = iota
    NEEDSREVIEW_MIGRATIONSTATUS
    ADDITIONALSTEPSREQUIRED_MIGRATIONSTATUS
    UNKNOWNFUTUREVALUE_MIGRATIONSTATUS
)

func (i MigrationStatus) String() string {
    return []string{"ready", "needsReview", "additionalStepsRequired", "unknownFutureValue"}[i]
}
func ParseMigrationStatus(v string) (interface{}, error) {
    result := READY_MIGRATIONSTATUS
    switch v {
        case "ready":
            result = READY_MIGRATIONSTATUS
        case "needsReview":
            result = NEEDSREVIEW_MIGRATIONSTATUS
        case "additionalStepsRequired":
            result = ADDITIONALSTEPSREQUIRED_MIGRATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MIGRATIONSTATUS
        default:
            return 0, errors.New("Unknown MigrationStatus value: " + v)
    }
    return &result, nil
}
func SerializeMigrationStatus(values []MigrationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
