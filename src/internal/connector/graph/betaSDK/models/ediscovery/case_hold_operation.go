package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseHoldOperation 
type CaseHoldOperation struct {
    CaseOperation
}
// NewCaseHoldOperation instantiates a new CaseHoldOperation and sets the default values.
func NewCaseHoldOperation()(*CaseHoldOperation) {
    m := &CaseHoldOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateCaseHoldOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCaseHoldOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCaseHoldOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CaseHoldOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CaseHoldOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
