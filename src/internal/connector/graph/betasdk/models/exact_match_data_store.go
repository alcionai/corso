package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchDataStore 
type ExactMatchDataStore struct {
    ExactMatchDataStoreBase
    // The sessions property
    sessions []ExactMatchSessionable
}
// NewExactMatchDataStore instantiates a new ExactMatchDataStore and sets the default values.
func NewExactMatchDataStore()(*ExactMatchDataStore) {
    m := &ExactMatchDataStore{
        ExactMatchDataStoreBase: *NewExactMatchDataStoreBase(),
    }
    return m
}
// CreateExactMatchDataStoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDataStoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchDataStore(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDataStore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExactMatchDataStoreBase.GetFieldDeserializers()
    res["sessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactMatchSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchSessionable, len(val))
            for i, v := range val {
                res[i] = v.(ExactMatchSessionable)
            }
            m.SetSessions(res)
        }
        return nil
    }
    return res
}
// GetSessions gets the sessions property value. The sessions property
func (m *ExactMatchDataStore) GetSessions()([]ExactMatchSessionable) {
    return m.sessions
}
// Serialize serializes information the current object
func (m *ExactMatchDataStore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExactMatchDataStoreBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSessions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSessions sets the sessions property value. The sessions property
func (m *ExactMatchDataStore) SetSessions(value []ExactMatchSessionable)() {
    m.sessions = value
}
