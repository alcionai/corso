package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingAppointment 
type BookingAppointment struct {
    Entity
    // Additional information that is sent to the customer when an appointment is confirmed.
    additionalInformation *string
    // The URL of the meeting to join anonymously.
    anonymousJoinWebUrl *string
    // The SMTP address of the bookingCustomer who is booking the appointment.
    customerEmailAddress *string
    // The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
    customerId *string
    // Represents location information for the bookingCustomer who is booking the appointment.
    customerLocation Locationable
    // The customer's name.
    customerName *string
    // Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID.  You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by customerId.
    customerNotes *string
    // The customer's phone number.
    customerPhone *string
    // A collection of the customer properties for an appointment. An appointment will contain a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
    customers []BookingCustomerInformationBaseable
    // The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
    customerTimeZone *string
    // The length of the appointment, denoted in ISO8601 format.
    duration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The end property
    end DateTimeTimeZoneable
    // The current number of customers in the appointment.
    filledAttendeesCount *int32
    // The billed amount on the invoice.
    invoiceAmount *float64
    // The date, time, and time zone of the invoice for this appointment.
    invoiceDate DateTimeTimeZoneable
    // The ID of the invoice.
    invoiceId *string
    // The invoiceStatus property
    invoiceStatus *BookingInvoiceStatus
    // The URL of the invoice in Microsoft Bookings.
    invoiceUrl *string
    // True indicates that the appointment will be held online. Default value is false.
    isLocationOnline *bool
    // The URL of the online meeting for the appointment.
    joinWebUrl *string
    // The maximum number of customers allowed in an appointment. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
    maximumAttendeesCount *int32
    // The onlineMeetingUrl property
    onlineMeetingUrl *string
    // True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
    optOutOfCustomerEmail *bool
    // The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
    postBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
    preBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The regular price for an appointment for the specified bookingService.
    price *float64
    // Represents the type of pricing of a booking service.
    priceType *BookingPriceType
    // The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
    reminders []BookingReminderable
    // An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
    selfServiceAppointmentId *string
    // The ID of the bookingService associated with this appointment.
    serviceId *string
    // The location where the service is delivered.
    serviceLocation Locationable
    // The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
    serviceName *string
    // Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
    serviceNotes *string
    // True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
    smsNotificationsEnabled *bool
    // The ID of each bookingStaffMember who is scheduled in this appointment.
    staffMemberIds []string
    // The start property
    start DateTimeTimeZoneable
}
// NewBookingAppointment instantiates a new BookingAppointment and sets the default values.
func NewBookingAppointment()(*BookingAppointment) {
    m := &BookingAppointment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingAppointmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingAppointmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingAppointment(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingAppointment) GetAdditionalInformation()(*string) {
    return m.additionalInformation
}
// GetAnonymousJoinWebUrl gets the anonymousJoinWebUrl property value. The URL of the meeting to join anonymously.
func (m *BookingAppointment) GetAnonymousJoinWebUrl()(*string) {
    return m.anonymousJoinWebUrl
}
// GetCustomerEmailAddress gets the customerEmailAddress property value. The SMTP address of the bookingCustomer who is booking the appointment.
func (m *BookingAppointment) GetCustomerEmailAddress()(*string) {
    return m.customerEmailAddress
}
// GetCustomerId gets the customerId property value. The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
func (m *BookingAppointment) GetCustomerId()(*string) {
    return m.customerId
}
// GetCustomerLocation gets the customerLocation property value. Represents location information for the bookingCustomer who is booking the appointment.
func (m *BookingAppointment) GetCustomerLocation()(Locationable) {
    return m.customerLocation
}
// GetCustomerName gets the customerName property value. The customer's name.
func (m *BookingAppointment) GetCustomerName()(*string) {
    return m.customerName
}
// GetCustomerNotes gets the customerNotes property value. Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID.  You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by customerId.
func (m *BookingAppointment) GetCustomerNotes()(*string) {
    return m.customerNotes
}
// GetCustomerPhone gets the customerPhone property value. The customer's phone number.
func (m *BookingAppointment) GetCustomerPhone()(*string) {
    return m.customerPhone
}
// GetCustomers gets the customers property value. A collection of the customer properties for an appointment. An appointment will contain a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
func (m *BookingAppointment) GetCustomers()([]BookingCustomerInformationBaseable) {
    return m.customers
}
// GetCustomerTimeZone gets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingAppointment) GetCustomerTimeZone()(*string) {
    return m.customerTimeZone
}
// GetDuration gets the duration property value. The length of the appointment, denoted in ISO8601 format.
func (m *BookingAppointment) GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.duration
}
// GetEnd gets the end property value. The end property
func (m *BookingAppointment) GetEnd()(DateTimeTimeZoneable) {
    return m.end
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingAppointment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["anonymousJoinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnonymousJoinWebUrl(val)
        }
        return nil
    }
    res["customerEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerEmailAddress(val)
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["customerLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerLocation(val.(Locationable))
        }
        return nil
    }
    res["customerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerName(val)
        }
        return nil
    }
    res["customerNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNotes(val)
        }
        return nil
    }
    res["customerPhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerPhone(val)
        }
        return nil
    }
    res["customers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCustomerInformationBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomerInformationBaseable, len(val))
            for i, v := range val {
                res[i] = v.(BookingCustomerInformationBaseable)
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["customerTimeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerTimeZone(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["filledAttendeesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilledAttendeesCount(val)
        }
        return nil
    }
    res["invoiceAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceAmount(val)
        }
        return nil
    }
    res["invoiceDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDate(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["invoiceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceId(val)
        }
        return nil
    }
    res["invoiceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingInvoiceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceStatus(val.(*BookingInvoiceStatus))
        }
        return nil
    }
    res["invoiceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceUrl(val)
        }
        return nil
    }
    res["isLocationOnline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocationOnline(val)
        }
        return nil
    }
    res["joinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
        }
        return nil
    }
    res["maximumAttendeesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAttendeesCount(val)
        }
        return nil
    }
    res["onlineMeetingUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeetingUrl(val)
        }
        return nil
    }
    res["optOutOfCustomerEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptOutOfCustomerEmail(val)
        }
        return nil
    }
    res["postBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBuffer(val)
        }
        return nil
    }
    res["preBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreBuffer(val)
        }
        return nil
    }
    res["price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrice(val)
        }
        return nil
    }
    res["priceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceType(val.(*BookingPriceType))
        }
        return nil
    }
    res["reminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingReminderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingReminderable, len(val))
            for i, v := range val {
                res[i] = v.(BookingReminderable)
            }
            m.SetReminders(res)
        }
        return nil
    }
    res["selfServiceAppointmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceAppointmentId(val)
        }
        return nil
    }
    res["serviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    res["serviceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceLocation(val.(Locationable))
        }
        return nil
    }
    res["serviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    res["serviceNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceNotes(val)
        }
        return nil
    }
    res["smsNotificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsNotificationsEnabled(val)
        }
        return nil
    }
    res["staffMemberIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStaffMemberIds(res)
        }
        return nil
    }
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetFilledAttendeesCount gets the filledAttendeesCount property value. The current number of customers in the appointment.
func (m *BookingAppointment) GetFilledAttendeesCount()(*int32) {
    return m.filledAttendeesCount
}
// GetInvoiceAmount gets the invoiceAmount property value. The billed amount on the invoice.
func (m *BookingAppointment) GetInvoiceAmount()(*float64) {
    return m.invoiceAmount
}
// GetInvoiceDate gets the invoiceDate property value. The date, time, and time zone of the invoice for this appointment.
func (m *BookingAppointment) GetInvoiceDate()(DateTimeTimeZoneable) {
    return m.invoiceDate
}
// GetInvoiceId gets the invoiceId property value. The ID of the invoice.
func (m *BookingAppointment) GetInvoiceId()(*string) {
    return m.invoiceId
}
// GetInvoiceStatus gets the invoiceStatus property value. The invoiceStatus property
func (m *BookingAppointment) GetInvoiceStatus()(*BookingInvoiceStatus) {
    return m.invoiceStatus
}
// GetInvoiceUrl gets the invoiceUrl property value. The URL of the invoice in Microsoft Bookings.
func (m *BookingAppointment) GetInvoiceUrl()(*string) {
    return m.invoiceUrl
}
// GetIsLocationOnline gets the isLocationOnline property value. True indicates that the appointment will be held online. Default value is false.
func (m *BookingAppointment) GetIsLocationOnline()(*bool) {
    return m.isLocationOnline
}
// GetJoinWebUrl gets the joinWebUrl property value. The URL of the online meeting for the appointment.
func (m *BookingAppointment) GetJoinWebUrl()(*string) {
    return m.joinWebUrl
}
// GetMaximumAttendeesCount gets the maximumAttendeesCount property value. The maximum number of customers allowed in an appointment. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
func (m *BookingAppointment) GetMaximumAttendeesCount()(*int32) {
    return m.maximumAttendeesCount
}
// GetOnlineMeetingUrl gets the onlineMeetingUrl property value. The onlineMeetingUrl property
func (m *BookingAppointment) GetOnlineMeetingUrl()(*string) {
    return m.onlineMeetingUrl
}
// GetOptOutOfCustomerEmail gets the optOutOfCustomerEmail property value. True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
func (m *BookingAppointment) GetOptOutOfCustomerEmail()(*bool) {
    return m.optOutOfCustomerEmail
}
// GetPostBuffer gets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.postBuffer
}
// GetPreBuffer gets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.preBuffer
}
// GetPrice gets the price property value. The regular price for an appointment for the specified bookingService.
func (m *BookingAppointment) GetPrice()(*float64) {
    return m.price
}
// GetPriceType gets the priceType property value. Represents the type of pricing of a booking service.
func (m *BookingAppointment) GetPriceType()(*BookingPriceType) {
    return m.priceType
}
// GetReminders gets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) GetReminders()([]BookingReminderable) {
    return m.reminders
}
// GetSelfServiceAppointmentId gets the selfServiceAppointmentId property value. An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
func (m *BookingAppointment) GetSelfServiceAppointmentId()(*string) {
    return m.selfServiceAppointmentId
}
// GetServiceId gets the serviceId property value. The ID of the bookingService associated with this appointment.
func (m *BookingAppointment) GetServiceId()(*string) {
    return m.serviceId
}
// GetServiceLocation gets the serviceLocation property value. The location where the service is delivered.
func (m *BookingAppointment) GetServiceLocation()(Locationable) {
    return m.serviceLocation
}
// GetServiceName gets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
func (m *BookingAppointment) GetServiceName()(*string) {
    return m.serviceName
}
// GetServiceNotes gets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) GetServiceNotes()(*string) {
    return m.serviceNotes
}
// GetSmsNotificationsEnabled gets the smsNotificationsEnabled property value. True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
func (m *BookingAppointment) GetSmsNotificationsEnabled()(*bool) {
    return m.smsNotificationsEnabled
}
// GetStaffMemberIds gets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
func (m *BookingAppointment) GetStaffMemberIds()([]string) {
    return m.staffMemberIds
}
// GetStart gets the start property value. The start property
func (m *BookingAppointment) GetStart()(DateTimeTimeZoneable) {
    return m.start
}
// Serialize serializes information the current object
func (m *BookingAppointment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("anonymousJoinWebUrl", m.GetAnonymousJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerEmailAddress", m.GetCustomerEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customerLocation", m.GetCustomerLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerName", m.GetCustomerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerNotes", m.GetCustomerNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerPhone", m.GetCustomerPhone())
        if err != nil {
            return err
        }
    }
    if m.GetCustomers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerTimeZone", m.GetCustomerTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("invoiceAmount", m.GetInvoiceAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("invoiceDate", m.GetInvoiceDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invoiceId", m.GetInvoiceId())
        if err != nil {
            return err
        }
    }
    if m.GetInvoiceStatus() != nil {
        cast := (*m.GetInvoiceStatus()).String()
        err = writer.WriteStringValue("invoiceStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invoiceUrl", m.GetInvoiceUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocationOnline", m.GetIsLocationOnline())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAttendeesCount", m.GetMaximumAttendeesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineMeetingUrl", m.GetOnlineMeetingUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("optOutOfCustomerEmail", m.GetOptOutOfCustomerEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("postBuffer", m.GetPostBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("preBuffer", m.GetPreBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("price", m.GetPrice())
        if err != nil {
            return err
        }
    }
    if m.GetPriceType() != nil {
        cast := (*m.GetPriceType()).String()
        err = writer.WriteStringValue("priceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReminders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReminders()))
        for i, v := range m.GetReminders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("selfServiceAppointmentId", m.GetSelfServiceAppointmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("serviceLocation", m.GetServiceLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceName", m.GetServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceNotes", m.GetServiceNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smsNotificationsEnabled", m.GetSmsNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetStaffMemberIds() != nil {
        err = writer.WriteCollectionOfStringValues("staffMemberIds", m.GetStaffMemberIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalInformation sets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingAppointment) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// SetAnonymousJoinWebUrl sets the anonymousJoinWebUrl property value. The URL of the meeting to join anonymously.
func (m *BookingAppointment) SetAnonymousJoinWebUrl(value *string)() {
    m.anonymousJoinWebUrl = value
}
// SetCustomerEmailAddress sets the customerEmailAddress property value. The SMTP address of the bookingCustomer who is booking the appointment.
func (m *BookingAppointment) SetCustomerEmailAddress(value *string)() {
    m.customerEmailAddress = value
}
// SetCustomerId sets the customerId property value. The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
func (m *BookingAppointment) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetCustomerLocation sets the customerLocation property value. Represents location information for the bookingCustomer who is booking the appointment.
func (m *BookingAppointment) SetCustomerLocation(value Locationable)() {
    m.customerLocation = value
}
// SetCustomerName sets the customerName property value. The customer's name.
func (m *BookingAppointment) SetCustomerName(value *string)() {
    m.customerName = value
}
// SetCustomerNotes sets the customerNotes property value. Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID.  You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by customerId.
func (m *BookingAppointment) SetCustomerNotes(value *string)() {
    m.customerNotes = value
}
// SetCustomerPhone sets the customerPhone property value. The customer's phone number.
func (m *BookingAppointment) SetCustomerPhone(value *string)() {
    m.customerPhone = value
}
// SetCustomers sets the customers property value. A collection of the customer properties for an appointment. An appointment will contain a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
func (m *BookingAppointment) SetCustomers(value []BookingCustomerInformationBaseable)() {
    m.customers = value
}
// SetCustomerTimeZone sets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingAppointment) SetCustomerTimeZone(value *string)() {
    m.customerTimeZone = value
}
// SetDuration sets the duration property value. The length of the appointment, denoted in ISO8601 format.
func (m *BookingAppointment) SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.duration = value
}
// SetEnd sets the end property value. The end property
func (m *BookingAppointment) SetEnd(value DateTimeTimeZoneable)() {
    m.end = value
}
// SetFilledAttendeesCount sets the filledAttendeesCount property value. The current number of customers in the appointment.
func (m *BookingAppointment) SetFilledAttendeesCount(value *int32)() {
    m.filledAttendeesCount = value
}
// SetInvoiceAmount sets the invoiceAmount property value. The billed amount on the invoice.
func (m *BookingAppointment) SetInvoiceAmount(value *float64)() {
    m.invoiceAmount = value
}
// SetInvoiceDate sets the invoiceDate property value. The date, time, and time zone of the invoice for this appointment.
func (m *BookingAppointment) SetInvoiceDate(value DateTimeTimeZoneable)() {
    m.invoiceDate = value
}
// SetInvoiceId sets the invoiceId property value. The ID of the invoice.
func (m *BookingAppointment) SetInvoiceId(value *string)() {
    m.invoiceId = value
}
// SetInvoiceStatus sets the invoiceStatus property value. The invoiceStatus property
func (m *BookingAppointment) SetInvoiceStatus(value *BookingInvoiceStatus)() {
    m.invoiceStatus = value
}
// SetInvoiceUrl sets the invoiceUrl property value. The URL of the invoice in Microsoft Bookings.
func (m *BookingAppointment) SetInvoiceUrl(value *string)() {
    m.invoiceUrl = value
}
// SetIsLocationOnline sets the isLocationOnline property value. True indicates that the appointment will be held online. Default value is false.
func (m *BookingAppointment) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
// SetJoinWebUrl sets the joinWebUrl property value. The URL of the online meeting for the appointment.
func (m *BookingAppointment) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// SetMaximumAttendeesCount sets the maximumAttendeesCount property value. The maximum number of customers allowed in an appointment. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
func (m *BookingAppointment) SetMaximumAttendeesCount(value *int32)() {
    m.maximumAttendeesCount = value
}
// SetOnlineMeetingUrl sets the onlineMeetingUrl property value. The onlineMeetingUrl property
func (m *BookingAppointment) SetOnlineMeetingUrl(value *string)() {
    m.onlineMeetingUrl = value
}
// SetOptOutOfCustomerEmail sets the optOutOfCustomerEmail property value. True indicates that the bookingCustomer for this appointment does not wish to receive a confirmation for this appointment.
func (m *BookingAppointment) SetOptOutOfCustomerEmail(value *bool)() {
    m.optOutOfCustomerEmail = value
}
// SetPostBuffer sets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.postBuffer = value
}
// SetPreBuffer sets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.preBuffer = value
}
// SetPrice sets the price property value. The regular price for an appointment for the specified bookingService.
func (m *BookingAppointment) SetPrice(value *float64)() {
    m.price = value
}
// SetPriceType sets the priceType property value. Represents the type of pricing of a booking service.
func (m *BookingAppointment) SetPriceType(value *BookingPriceType)() {
    m.priceType = value
}
// SetReminders sets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) SetReminders(value []BookingReminderable)() {
    m.reminders = value
}
// SetSelfServiceAppointmentId sets the selfServiceAppointmentId property value. An additional tracking ID for the appointment, if the appointment has been created directly by the customer on the scheduling page, as opposed to by a staff member on the behalf of the customer.
func (m *BookingAppointment) SetSelfServiceAppointmentId(value *string)() {
    m.selfServiceAppointmentId = value
}
// SetServiceId sets the serviceId property value. The ID of the bookingService associated with this appointment.
func (m *BookingAppointment) SetServiceId(value *string)() {
    m.serviceId = value
}
// SetServiceLocation sets the serviceLocation property value. The location where the service is delivered.
func (m *BookingAppointment) SetServiceLocation(value Locationable)() {
    m.serviceLocation = value
}
// SetServiceName sets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it is computed from the service associated with the appointment by the serviceId property.
func (m *BookingAppointment) SetServiceName(value *string)() {
    m.serviceName = value
}
// SetServiceNotes sets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) SetServiceNotes(value *string)() {
    m.serviceNotes = value
}
// SetSmsNotificationsEnabled sets the smsNotificationsEnabled property value. True indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
func (m *BookingAppointment) SetSmsNotificationsEnabled(value *bool)() {
    m.smsNotificationsEnabled = value
}
// SetStaffMemberIds sets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
func (m *BookingAppointment) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
// SetStart sets the start property value. The start property
func (m *BookingAppointment) SetStart(value DateTimeTimeZoneable)() {
    m.start = value
}
