package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagedAppNotificationRestriction int

const (
    // Share all notifications.
    ALLOW_MANAGEDAPPNOTIFICATIONRESTRICTION ManagedAppNotificationRestriction = iota
    // Do not share Orgnizational data in notifications.
    BLOCKORGANIZATIONALDATA_MANAGEDAPPNOTIFICATIONRESTRICTION
    // Do not share notifications.
    BLOCK_MANAGEDAPPNOTIFICATIONRESTRICTION
)

func (i ManagedAppNotificationRestriction) String() string {
    return []string{"allow", "blockOrganizationalData", "block"}[i]
}
func ParseManagedAppNotificationRestriction(v string) (interface{}, error) {
    result := ALLOW_MANAGEDAPPNOTIFICATIONRESTRICTION
    switch v {
        case "allow":
            result = ALLOW_MANAGEDAPPNOTIFICATIONRESTRICTION
        case "blockOrganizationalData":
            result = BLOCKORGANIZATIONALDATA_MANAGEDAPPNOTIFICATIONRESTRICTION
        case "block":
            result = BLOCK_MANAGEDAPPNOTIFICATIONRESTRICTION
        default:
            return 0, errors.New("Unknown ManagedAppNotificationRestriction value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppNotificationRestriction(values []ManagedAppNotificationRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
