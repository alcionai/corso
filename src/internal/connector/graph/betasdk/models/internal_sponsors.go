package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InternalSponsors 
type InternalSponsors struct {
    UserSet
}
// NewInternalSponsors instantiates a new InternalSponsors and sets the default values.
func NewInternalSponsors()(*InternalSponsors) {
    m := &InternalSponsors{
        UserSet: *NewUserSet(),
    }
    odataTypeValue := "#microsoft.graph.internalSponsors";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateInternalSponsorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInternalSponsorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInternalSponsors(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternalSponsors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserSet.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InternalSponsors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserSet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
