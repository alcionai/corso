package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingStaffMember 
type BookingStaffMember struct {
    BookingPerson
    // True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
    availabilityIsAffectedByPersonalCalendar *bool
    // Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
    colorIndex *int32
    // True indicates that a staff member will be notified via email when a booking assigned to them is created or changed.
    isEmailNotificationEnabled *bool
    // The membershipStatus property
    membershipStatus *BookingStaffMembershipStatus
    // The role property
    role *BookingStaffRole
    // The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
    timeZone *string
    // True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
    useBusinessHours *bool
    // The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
    workingHours []BookingWorkHoursable
}
// NewBookingStaffMember instantiates a new BookingStaffMember and sets the default values.
func NewBookingStaffMember()(*BookingStaffMember) {
    m := &BookingStaffMember{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
// CreateBookingStaffMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingStaffMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingStaffMember(), nil
}
// GetAvailabilityIsAffectedByPersonalCalendar gets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) GetAvailabilityIsAffectedByPersonalCalendar()(*bool) {
    return m.availabilityIsAffectedByPersonalCalendar
}
// GetColorIndex gets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
func (m *BookingStaffMember) GetColorIndex()(*int32) {
    return m.colorIndex
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingStaffMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["availabilityIsAffectedByPersonalCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityIsAffectedByPersonalCalendar(val)
        }
        return nil
    }
    res["colorIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorIndex(val)
        }
        return nil
    }
    res["isEmailNotificationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEmailNotificationEnabled(val)
        }
        return nil
    }
    res["membershipStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffMembershipStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipStatus(val.(*BookingStaffMembershipStatus))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingStaffRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*BookingStaffRole))
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["useBusinessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseBusinessHours(val)
        }
        return nil
    }
    res["workingHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHoursable, len(val))
            for i, v := range val {
                res[i] = v.(BookingWorkHoursable)
            }
            m.SetWorkingHours(res)
        }
        return nil
    }
    return res
}
// GetIsEmailNotificationEnabled gets the isEmailNotificationEnabled property value. True indicates that a staff member will be notified via email when a booking assigned to them is created or changed.
func (m *BookingStaffMember) GetIsEmailNotificationEnabled()(*bool) {
    return m.isEmailNotificationEnabled
}
// GetMembershipStatus gets the membershipStatus property value. The membershipStatus property
func (m *BookingStaffMember) GetMembershipStatus()(*BookingStaffMembershipStatus) {
    return m.membershipStatus
}
// GetRole gets the role property value. The role property
func (m *BookingStaffMember) GetRole()(*BookingStaffRole) {
    return m.role
}
// GetTimeZone gets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) GetTimeZone()(*string) {
    return m.timeZone
}
// GetUseBusinessHours gets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) GetUseBusinessHours()(*bool) {
    return m.useBusinessHours
}
// GetWorkingHours gets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) GetWorkingHours()([]BookingWorkHoursable) {
    return m.workingHours
}
// Serialize serializes information the current object
func (m *BookingStaffMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingPerson.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("availabilityIsAffectedByPersonalCalendar", m.GetAvailabilityIsAffectedByPersonalCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("colorIndex", m.GetColorIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEmailNotificationEnabled", m.GetIsEmailNotificationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetMembershipStatus() != nil {
        cast := (*m.GetMembershipStatus()).String()
        err = writer.WriteStringValue("membershipStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useBusinessHours", m.GetUseBusinessHours())
        if err != nil {
            return err
        }
    }
    if m.GetWorkingHours() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkingHours()))
        for i, v := range m.GetWorkingHours() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("workingHours", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityIsAffectedByPersonalCalendar sets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) SetAvailabilityIsAffectedByPersonalCalendar(value *bool)() {
    m.availabilityIsAffectedByPersonalCalendar = value
}
// SetColorIndex sets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
func (m *BookingStaffMember) SetColorIndex(value *int32)() {
    m.colorIndex = value
}
// SetIsEmailNotificationEnabled sets the isEmailNotificationEnabled property value. True indicates that a staff member will be notified via email when a booking assigned to them is created or changed.
func (m *BookingStaffMember) SetIsEmailNotificationEnabled(value *bool)() {
    m.isEmailNotificationEnabled = value
}
// SetMembershipStatus sets the membershipStatus property value. The membershipStatus property
func (m *BookingStaffMember) SetMembershipStatus(value *BookingStaffMembershipStatus)() {
    m.membershipStatus = value
}
// SetRole sets the role property value. The role property
func (m *BookingStaffMember) SetRole(value *BookingStaffRole)() {
    m.role = value
}
// SetTimeZone sets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) SetTimeZone(value *string)() {
    m.timeZone = value
}
// SetUseBusinessHours sets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) SetUseBusinessHours(value *bool)() {
    m.useBusinessHours = value
}
// SetWorkingHours sets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) SetWorkingHours(value []BookingWorkHoursable)() {
    m.workingHours = value
}
