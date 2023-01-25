package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type PersonAnnualEventType int

const (
    BIRTHDAY_PERSONANNUALEVENTTYPE PersonAnnualEventType = iota
    WEDDING_PERSONANNUALEVENTTYPE
    WORK_PERSONANNUALEVENTTYPE
    OTHER_PERSONANNUALEVENTTYPE
    UNKNOWNFUTUREVALUE_PERSONANNUALEVENTTYPE
)

func (i PersonAnnualEventType) String() string {
    return []string{"birthday", "wedding", "work", "other", "unknownFutureValue"}[i]
}
func ParsePersonAnnualEventType(v string) (interface{}, error) {
    result := BIRTHDAY_PERSONANNUALEVENTTYPE
    switch v {
        case "birthday":
            result = BIRTHDAY_PERSONANNUALEVENTTYPE
        case "wedding":
            result = WEDDING_PERSONANNUALEVENTTYPE
        case "work":
            result = WORK_PERSONANNUALEVENTTYPE
        case "other":
            result = OTHER_PERSONANNUALEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERSONANNUALEVENTTYPE
        default:
            return 0, errors.New("Unknown PersonAnnualEventType value: " + v)
    }
    return &result, nil
}
func SerializePersonAnnualEventType(values []PersonAnnualEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
