package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ZebraFotaScheduleMode int

const (
    // Instructs the device to install the update as soon as it is received.
    INSTALLNOW_ZEBRAFOTASCHEDULEMODE ZebraFotaScheduleMode = iota
    // Schedule an update to be installed at a specified date and time.
    SCHEDULED_ZEBRAFOTASCHEDULEMODE
    // Unknown future enum value.
    UNKNOWNFUTUREVALUE_ZEBRAFOTASCHEDULEMODE
)

func (i ZebraFotaScheduleMode) String() string {
    return []string{"installNow", "scheduled", "unknownFutureValue"}[i]
}
func ParseZebraFotaScheduleMode(v string) (interface{}, error) {
    result := INSTALLNOW_ZEBRAFOTASCHEDULEMODE
    switch v {
        case "installNow":
            result = INSTALLNOW_ZEBRAFOTASCHEDULEMODE
        case "scheduled":
            result = SCHEDULED_ZEBRAFOTASCHEDULEMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTASCHEDULEMODE
        default:
            return 0, errors.New("Unknown ZebraFotaScheduleMode value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaScheduleMode(values []ZebraFotaScheduleMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
