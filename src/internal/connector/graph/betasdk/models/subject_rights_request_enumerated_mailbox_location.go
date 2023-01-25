package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestEnumeratedMailboxLocation 
type SubjectRightsRequestEnumeratedMailboxLocation struct {
    SubjectRightsRequestMailboxLocation
    // Collection of mailboxes that should be included in the search. Includes the UPN (user principal name) of each mailbox, for example, Monica.Thompson@contoso.com.
    upns []string
}
// NewSubjectRightsRequestEnumeratedMailboxLocation instantiates a new SubjectRightsRequestEnumeratedMailboxLocation and sets the default values.
func NewSubjectRightsRequestEnumeratedMailboxLocation()(*SubjectRightsRequestEnumeratedMailboxLocation) {
    m := &SubjectRightsRequestEnumeratedMailboxLocation{
        SubjectRightsRequestMailboxLocation: *NewSubjectRightsRequestMailboxLocation(),
    }
    odataTypeValue := "#microsoft.graph.subjectRightsRequestEnumeratedMailboxLocation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSubjectRightsRequestEnumeratedMailboxLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestEnumeratedMailboxLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestEnumeratedMailboxLocation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestEnumeratedMailboxLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectRightsRequestMailboxLocation.GetFieldDeserializers()
    res["upns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUpns(res)
        }
        return nil
    }
    return res
}
// GetUpns gets the upns property value. Collection of mailboxes that should be included in the search. Includes the UPN (user principal name) of each mailbox, for example, Monica.Thompson@contoso.com.
func (m *SubjectRightsRequestEnumeratedMailboxLocation) GetUpns()([]string) {
    return m.upns
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestEnumeratedMailboxLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectRightsRequestMailboxLocation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUpns() != nil {
        err = writer.WriteCollectionOfStringValues("upns", m.GetUpns())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpns sets the upns property value. Collection of mailboxes that should be included in the search. Includes the UPN (user principal name) of each mailbox, for example, Monica.Thompson@contoso.com.
func (m *SubjectRightsRequestEnumeratedMailboxLocation) SetUpns(value []string)() {
    m.upns = value
}
