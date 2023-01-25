package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingPerson represents a booking customer or staff member.
type BookingPerson struct {
    BookingNamedEntity
    // The email address of the person.
    emailAddress *string
}
// NewBookingPerson instantiates a new bookingPerson and sets the default values.
func NewBookingPerson()(*BookingPerson) {
    m := &BookingPerson{
        BookingNamedEntity: *NewBookingNamedEntity(),
    }
    odataTypeValue := "#microsoft.graph.bookingPerson";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBookingPersonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingPersonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.bookingCustomer":
                        return NewBookingCustomer(), nil
                    case "#microsoft.graph.bookingStaffMember":
                        return NewBookingStaffMember(), nil
                }
            }
        }
    }
    return NewBookingPerson(), nil
}
// GetEmailAddress gets the emailAddress property value. The email address of the person.
func (m *BookingPerson) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingPerson) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingNamedEntity.GetFieldDeserializers()
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BookingPerson) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingNamedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailAddress sets the emailAddress property value. The email address of the person.
func (m *BookingPerson) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
