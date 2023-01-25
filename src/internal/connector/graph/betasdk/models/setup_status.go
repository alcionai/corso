package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type SetupStatus int

const (
    UNKNOWN_SETUPSTATUS SetupStatus = iota
    NOTREGISTEREDYET_SETUPSTATUS
    REGISTEREDSETUPNOTSTARTED_SETUPSTATUS
    REGISTEREDSETUPINPROGRESS_SETUPSTATUS
    REGISTRATIONANDSETUPCOMPLETED_SETUPSTATUS
    REGISTRATIONFAILED_SETUPSTATUS
    REGISTRATIONTIMEDOUT_SETUPSTATUS
    DISABLED_SETUPSTATUS
)

func (i SetupStatus) String() string {
    return []string{"unknown", "notRegisteredYet", "registeredSetupNotStarted", "registeredSetupInProgress", "registrationAndSetupCompleted", "registrationFailed", "registrationTimedOut", "disabled"}[i]
}
func ParseSetupStatus(v string) (interface{}, error) {
    result := UNKNOWN_SETUPSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_SETUPSTATUS
        case "notRegisteredYet":
            result = NOTREGISTEREDYET_SETUPSTATUS
        case "registeredSetupNotStarted":
            result = REGISTEREDSETUPNOTSTARTED_SETUPSTATUS
        case "registeredSetupInProgress":
            result = REGISTEREDSETUPINPROGRESS_SETUPSTATUS
        case "registrationAndSetupCompleted":
            result = REGISTRATIONANDSETUPCOMPLETED_SETUPSTATUS
        case "registrationFailed":
            result = REGISTRATIONFAILED_SETUPSTATUS
        case "registrationTimedOut":
            result = REGISTRATIONTIMEDOUT_SETUPSTATUS
        case "disabled":
            result = DISABLED_SETUPSTATUS
        default:
            return 0, errors.New("Unknown SetupStatus value: " + v)
    }
    return &result, nil
}
func SerializeSetupStatus(values []SetupStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
