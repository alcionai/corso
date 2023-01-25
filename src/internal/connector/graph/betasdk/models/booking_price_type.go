package models
import (
    "errors"
)
// Provides operations to call the add method.
type BookingPriceType int

const (
    // The price of the service is not defined.
    UNDEFINED_BOOKINGPRICETYPE BookingPriceType = iota
    // The price of the service is fixed.
    FIXEDPRICE_BOOKINGPRICETYPE
    // The price of the service starts with a particular value, but can be higher based on the final services performed.
    STARTINGAT_BOOKINGPRICETYPE
    // The price of the service depends on the number of hours a staff member works on the service.
    HOURLY_BOOKINGPRICETYPE
    // The service is free.
    FREE_BOOKINGPRICETYPE
    // The price of the service varies.
    PRICEVARIES_BOOKINGPRICETYPE
    // The price of the service is not listed.
    CALLUS_BOOKINGPRICETYPE
    // The price of the service is not set.
    NOTSET_BOOKINGPRICETYPE
)

func (i BookingPriceType) String() string {
    return []string{"undefined", "fixedPrice", "startingAt", "hourly", "free", "priceVaries", "callUs", "notSet"}[i]
}
func ParseBookingPriceType(v string) (interface{}, error) {
    result := UNDEFINED_BOOKINGPRICETYPE
    switch v {
        case "undefined":
            result = UNDEFINED_BOOKINGPRICETYPE
        case "fixedPrice":
            result = FIXEDPRICE_BOOKINGPRICETYPE
        case "startingAt":
            result = STARTINGAT_BOOKINGPRICETYPE
        case "hourly":
            result = HOURLY_BOOKINGPRICETYPE
        case "free":
            result = FREE_BOOKINGPRICETYPE
        case "priceVaries":
            result = PRICEVARIES_BOOKINGPRICETYPE
        case "callUs":
            result = CALLUS_BOOKINGPRICETYPE
        case "notSet":
            result = NOTSET_BOOKINGPRICETYPE
        default:
            return 0, errors.New("Unknown BookingPriceType value: " + v)
    }
    return &result, nil
}
func SerializeBookingPriceType(values []BookingPriceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
