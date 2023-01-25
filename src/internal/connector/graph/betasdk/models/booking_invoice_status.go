package models
import (
    "errors"
)
// Provides operations to call the add method.
type BookingInvoiceStatus int

const (
    DRAFT_BOOKINGINVOICESTATUS BookingInvoiceStatus = iota
    REVIEWING_BOOKINGINVOICESTATUS
    OPEN_BOOKINGINVOICESTATUS
    CANCELED_BOOKINGINVOICESTATUS
    PAID_BOOKINGINVOICESTATUS
    CORRECTIVE_BOOKINGINVOICESTATUS
)

func (i BookingInvoiceStatus) String() string {
    return []string{"draft", "reviewing", "open", "canceled", "paid", "corrective"}[i]
}
func ParseBookingInvoiceStatus(v string) (interface{}, error) {
    result := DRAFT_BOOKINGINVOICESTATUS
    switch v {
        case "draft":
            result = DRAFT_BOOKINGINVOICESTATUS
        case "reviewing":
            result = REVIEWING_BOOKINGINVOICESTATUS
        case "open":
            result = OPEN_BOOKINGINVOICESTATUS
        case "canceled":
            result = CANCELED_BOOKINGINVOICESTATUS
        case "paid":
            result = PAID_BOOKINGINVOICESTATUS
        case "corrective":
            result = CORRECTIVE_BOOKINGINVOICESTATUS
        default:
            return 0, errors.New("Unknown BookingInvoiceStatus value: " + v)
    }
    return &result, nil
}
func SerializeBookingInvoiceStatus(values []BookingInvoiceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
