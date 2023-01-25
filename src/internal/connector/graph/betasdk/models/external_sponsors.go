package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalSponsors 
type ExternalSponsors struct {
    UserSet
}
// NewExternalSponsors instantiates a new ExternalSponsors and sets the default values.
func NewExternalSponsors()(*ExternalSponsors) {
    m := &ExternalSponsors{
        UserSet: *NewUserSet(),
    }
    odataTypeValue := "#microsoft.graph.externalSponsors";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExternalSponsorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalSponsorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalSponsors(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalSponsors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserSet.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExternalSponsors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserSet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
