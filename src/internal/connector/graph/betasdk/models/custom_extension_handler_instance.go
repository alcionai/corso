package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionHandlerInstance 
type CustomExtensionHandlerInstance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identifier of the customAccessPackageWorkflowExtension triggered at this instance.
    customExtensionId *string
    // The unique run ID for the logic app.
    externalCorrelationId *string
    // The OdataType property
    odataType *string
    // Indicates the stage of the request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
    stage *AccessPackageCustomExtensionStage
    // Status of the request to run the access package custom extension workflow that is associated with the logic app. The possible values are: requestSent, requestReceived, unknownFutureValue.
    status *AccessPackageCustomExtensionHandlerStatus
}
// NewCustomExtensionHandlerInstance instantiates a new customExtensionHandlerInstance and sets the default values.
func NewCustomExtensionHandlerInstance()(*CustomExtensionHandlerInstance) {
    m := &CustomExtensionHandlerInstance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCustomExtensionHandlerInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionHandlerInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionHandlerInstance(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionHandlerInstance) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCustomExtensionId gets the customExtensionId property value. Identifier of the customAccessPackageWorkflowExtension triggered at this instance.
func (m *CustomExtensionHandlerInstance) GetCustomExtensionId()(*string) {
    return m.customExtensionId
}
// GetExternalCorrelationId gets the externalCorrelationId property value. The unique run ID for the logic app.
func (m *CustomExtensionHandlerInstance) GetExternalCorrelationId()(*string) {
    return m.externalCorrelationId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionHandlerInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["customExtensionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtensionId(val)
        }
        return nil
    }
    res["externalCorrelationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalCorrelationId(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionHandlerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AccessPackageCustomExtensionHandlerStatus))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomExtensionHandlerInstance) GetOdataType()(*string) {
    return m.odataType
}
// GetStage gets the stage property value. Indicates the stage of the request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) GetStage()(*AccessPackageCustomExtensionStage) {
    return m.stage
}
// GetStatus gets the status property value. Status of the request to run the access package custom extension workflow that is associated with the logic app. The possible values are: requestSent, requestReceived, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) GetStatus()(*AccessPackageCustomExtensionHandlerStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CustomExtensionHandlerInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customExtensionId", m.GetCustomExtensionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalCorrelationId", m.GetExternalCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err := writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionHandlerInstance) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCustomExtensionId sets the customExtensionId property value. Identifier of the customAccessPackageWorkflowExtension triggered at this instance.
func (m *CustomExtensionHandlerInstance) SetCustomExtensionId(value *string)() {
    m.customExtensionId = value
}
// SetExternalCorrelationId sets the externalCorrelationId property value. The unique run ID for the logic app.
func (m *CustomExtensionHandlerInstance) SetExternalCorrelationId(value *string)() {
    m.externalCorrelationId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionHandlerInstance) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStage sets the stage property value. Indicates the stage of the request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) SetStage(value *AccessPackageCustomExtensionStage)() {
    m.stage = value
}
// SetStatus sets the status property value. Status of the request to run the access package custom extension workflow that is associated with the logic app. The possible values are: requestSent, requestReceived, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) SetStatus(value *AccessPackageCustomExtensionHandlerStatus)() {
    m.status = value
}
