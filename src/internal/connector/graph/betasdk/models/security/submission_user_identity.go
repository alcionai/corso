package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// SubmissionUserIdentity 
type SubmissionUserIdentity struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Identity
    // The email of user who is making the submission when logged in (delegated token case).
    email *string
}
// NewSubmissionUserIdentity instantiates a new SubmissionUserIdentity and sets the default values.
func NewSubmissionUserIdentity()(*SubmissionUserIdentity) {
    m := &SubmissionUserIdentity{
        Identity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.security.submissionUserIdentity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSubmissionUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubmissionUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubmissionUserIdentity(), nil
}
// GetEmail gets the email property value. The email of user who is making the submission when logged in (delegated token case).
func (m *SubmissionUserIdentity) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubmissionUserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SubmissionUserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmail sets the email property value. The email of user who is making the submission when logged in (delegated token case).
func (m *SubmissionUserIdentity) SetEmail(value *string)() {
    m.email = value
}
