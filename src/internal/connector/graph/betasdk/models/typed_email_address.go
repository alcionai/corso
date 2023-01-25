package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TypedEmailAddress 
type TypedEmailAddress struct {
    EmailAddress
    // To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
    otherLabel *string
    // The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
    type_escaped *EmailType
}
// NewTypedEmailAddress instantiates a new TypedEmailAddress and sets the default values.
func NewTypedEmailAddress()(*TypedEmailAddress) {
    m := &TypedEmailAddress{
        EmailAddress: *NewEmailAddress(),
    }
    odataTypeValue := "#microsoft.graph.typedEmailAddress";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTypedEmailAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTypedEmailAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTypedEmailAddress(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TypedEmailAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EmailAddress.GetFieldDeserializers()
    res["otherLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherLabel(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*EmailType))
        }
        return nil
    }
    return res
}
// GetOtherLabel gets the otherLabel property value. To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
func (m *TypedEmailAddress) GetOtherLabel()(*string) {
    return m.otherLabel
}
// GetType gets the type property value. The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
func (m *TypedEmailAddress) GetType()(*EmailType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *TypedEmailAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EmailAddress.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("otherLabel", m.GetOtherLabel())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOtherLabel sets the otherLabel property value. To specify a custom type of email address, set type to other, and assign otherLabel to a custom string. For example, you may use a specific email address for your volunteer activities. Set type to other, and set otherLabel to a custom string such as Volunteer work.
func (m *TypedEmailAddress) SetOtherLabel(value *string)() {
    m.otherLabel = value
}
// SetType sets the type property value. The type of email address. Possible values are: unknown, work, personal, main, other. The default value is unknown, which means address has not been set as a specific type.
func (m *TypedEmailAddress) SetType(value *EmailType)() {
    m.type_escaped = value
}
