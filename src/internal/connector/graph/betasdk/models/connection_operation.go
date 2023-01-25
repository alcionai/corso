package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectionOperation 
type ConnectionOperation struct {
    Entity
    // The error property
    error PublicErrorable
    // The status property
    status *ConnectionOperationStatus
}
// NewConnectionOperation instantiates a new ConnectionOperation and sets the default values.
func NewConnectionOperation()(*ConnectionOperation) {
    m := &ConnectionOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConnectionOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectionOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectionOperation(), nil
}
// GetError gets the error property value. The error property
func (m *ConnectionOperation) GetError()(PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectionOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ConnectionOperationStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
func (m *ConnectionOperation) GetStatus()(*ConnectionOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ConnectionOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. The error property
func (m *ConnectionOperation) SetError(value PublicErrorable)() {
    m.error = value
}
// SetStatus sets the status property value. The status property
func (m *ConnectionOperation) SetStatus(value *ConnectionOperationStatus)() {
    m.status = value
}
