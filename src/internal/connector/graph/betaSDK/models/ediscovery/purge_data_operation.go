package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PurgeDataOperation 
type PurgeDataOperation struct {
    CaseOperation
}
// NewPurgeDataOperation instantiates a new PurgeDataOperation and sets the default values.
func NewPurgeDataOperation()(*PurgeDataOperation) {
    m := &PurgeDataOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreatePurgeDataOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePurgeDataOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPurgeDataOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PurgeDataOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PurgeDataOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
