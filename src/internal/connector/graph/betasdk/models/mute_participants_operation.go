package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MuteParticipantsOperation 
type MuteParticipantsOperation struct {
    CommsOperation
    // The participants property
    participants []string
}
// NewMuteParticipantsOperation instantiates a new MuteParticipantsOperation and sets the default values.
func NewMuteParticipantsOperation()(*MuteParticipantsOperation) {
    m := &MuteParticipantsOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateMuteParticipantsOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMuteParticipantsOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMuteParticipantsOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MuteParticipantsOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetParticipants(res)
        }
        return nil
    }
    return res
}
// GetParticipants gets the participants property value. The participants property
func (m *MuteParticipantsOperation) GetParticipants()([]string) {
    return m.participants
}
// Serialize serializes information the current object
func (m *MuteParticipantsOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetParticipants() != nil {
        err = writer.WriteCollectionOfStringValues("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParticipants sets the participants property value. The participants property
func (m *MuteParticipantsOperation) SetParticipants(value []string)() {
    m.participants = value
}
