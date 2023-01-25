package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionHandler provides operations to manage the collection of site entities.
type CustomExtensionHandler struct {
    Entity
    // Indicates which custom workflow extension will be executed at this stage. Nullable. Supports $expand.
    customExtension CustomAccessPackageWorkflowExtensionable
    // Indicates the stage of the access package assignment request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
    stage *AccessPackageCustomExtensionStage
}
// NewCustomExtensionHandler instantiates a new customExtensionHandler and sets the default values.
func NewCustomExtensionHandler()(*CustomExtensionHandler) {
    m := &CustomExtensionHandler{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomExtensionHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionHandler(), nil
}
// GetCustomExtension gets the customExtension property value. Indicates which custom workflow extension will be executed at this stage. Nullable. Supports $expand.
func (m *CustomExtensionHandler) GetCustomExtension()(CustomAccessPackageWorkflowExtensionable) {
    return m.customExtension
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomAccessPackageWorkflowExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(CustomAccessPackageWorkflowExtensionable))
        }
        return nil
    }
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*AccessPackageCustomExtensionStage))
        }
        return nil
    }
    return res
}
// GetStage gets the stage property value. Indicates the stage of the access package assignment request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandler) GetStage()(*AccessPackageCustomExtensionStage) {
    return m.stage
}
// Serialize serializes information the current object
func (m *CustomExtensionHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err = writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtension sets the customExtension property value. Indicates which custom workflow extension will be executed at this stage. Nullable. Supports $expand.
func (m *CustomExtensionHandler) SetCustomExtension(value CustomAccessPackageWorkflowExtensionable)() {
    m.customExtension = value
}
// SetStage sets the stage property value. Indicates the stage of the access package assignment request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandler) SetStage(value *AccessPackageCustomExtensionStage)() {
    m.stage = value
}
