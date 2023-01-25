package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassificationInnerError 
type ClassificationInnerError struct {
    // The activityId property
    activityId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The clientRequestId property
    clientRequestId *string
    // The code property
    code *string
    // The errorDateTime property
    errorDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
}
// NewClassificationInnerError instantiates a new classificationInnerError and sets the default values.
func NewClassificationInnerError()(*ClassificationInnerError) {
    m := &ClassificationInnerError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassificationInnerErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationInnerErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassificationInnerError(), nil
}
// GetActivityId gets the activityId property value. The activityId property
func (m *ClassificationInnerError) GetActivityId()(*string) {
    return m.activityId
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationInnerError) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClientRequestId gets the clientRequestId property value. The clientRequestId property
func (m *ClassificationInnerError) GetClientRequestId()(*string) {
    return m.clientRequestId
}
// GetCode gets the code property value. The code property
func (m *ClassificationInnerError) GetCode()(*string) {
    return m.code
}
// GetErrorDateTime gets the errorDateTime property value. The errorDateTime property
func (m *ClassificationInnerError) GetErrorDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.errorDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationInnerError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityId(val)
        }
        return nil
    }
    res["clientRequestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientRequestId(val)
        }
        return nil
    }
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["errorDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDateTime(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ClassificationInnerError) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ClassificationInnerError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityId", m.GetActivityId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientRequestId", m.GetClientRequestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("errorDateTime", m.GetErrorDateTime())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityId sets the activityId property value. The activityId property
func (m *ClassificationInnerError) SetActivityId(value *string)() {
    m.activityId = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationInnerError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClientRequestId sets the clientRequestId property value. The clientRequestId property
func (m *ClassificationInnerError) SetClientRequestId(value *string)() {
    m.clientRequestId = value
}
// SetCode sets the code property value. The code property
func (m *ClassificationInnerError) SetCode(value *string)() {
    m.code = value
}
// SetErrorDateTime sets the errorDateTime property value. The errorDateTime property
func (m *ClassificationInnerError) SetErrorDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.errorDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ClassificationInnerError) SetOdataType(value *string)() {
    m.odataType = value
}
