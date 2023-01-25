package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsDriverUpdateProfileInventorySyncState int

const (
    // Pending sync.
    PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE WindowsDriverUpdateProfileInventorySyncState = iota
    // Successful sync.
    SUCCESS_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
    // Failed sync.
    FAILURE_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
)

func (i WindowsDriverUpdateProfileInventorySyncState) String() string {
    return []string{"pending", "success", "failure"}[i]
}
func ParseWindowsDriverUpdateProfileInventorySyncState(v string) (interface{}, error) {
    result := PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
    switch v {
        case "pending":
            result = PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        case "success":
            result = SUCCESS_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        case "failure":
            result = FAILURE_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        default:
            return 0, errors.New("Unknown WindowsDriverUpdateProfileInventorySyncState value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDriverUpdateProfileInventorySyncState(values []WindowsDriverUpdateProfileInventorySyncState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
