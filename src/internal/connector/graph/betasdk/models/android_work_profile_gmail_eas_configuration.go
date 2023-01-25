package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileGmailEasConfiguration 
type AndroidWorkProfileGmailEasConfiguration struct {
    AndroidWorkProfileEasEmailProfileBase
}
// NewAndroidWorkProfileGmailEasConfiguration instantiates a new AndroidWorkProfileGmailEasConfiguration and sets the default values.
func NewAndroidWorkProfileGmailEasConfiguration()(*AndroidWorkProfileGmailEasConfiguration) {
    m := &AndroidWorkProfileGmailEasConfiguration{
        AndroidWorkProfileEasEmailProfileBase: *NewAndroidWorkProfileEasEmailProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileGmailEasConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidWorkProfileGmailEasConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileGmailEasConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileGmailEasConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileGmailEasConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidWorkProfileEasEmailProfileBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileGmailEasConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidWorkProfileEasEmailProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
