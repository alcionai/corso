package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionCalloutResult 
type CustomExtensionCalloutResult struct {
    AuthenticationEventHandlerResult
    // The calloutDateTime property
    calloutDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customExtensionId property
    customExtensionId *string
    // The errorCode property
    errorCode *int32
    // The httpStatus property
    httpStatus *int32
    // The numberOfAttempts property
    numberOfAttempts *int32
}
// NewCustomExtensionCalloutResult instantiates a new CustomExtensionCalloutResult and sets the default values.
func NewCustomExtensionCalloutResult()(*CustomExtensionCalloutResult) {
    m := &CustomExtensionCalloutResult{
        AuthenticationEventHandlerResult: *NewAuthenticationEventHandlerResult(),
    }
    odataTypeValue := "#microsoft.graph.customExtensionCalloutResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomExtensionCalloutResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionCalloutResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionCalloutResult(), nil
}
// GetCalloutDateTime gets the calloutDateTime property value. The calloutDateTime property
func (m *CustomExtensionCalloutResult) GetCalloutDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.calloutDateTime
}
// GetCustomExtensionId gets the customExtensionId property value. The customExtensionId property
func (m *CustomExtensionCalloutResult) GetCustomExtensionId()(*string) {
    return m.customExtensionId
}
// GetErrorCode gets the errorCode property value. The errorCode property
func (m *CustomExtensionCalloutResult) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionCalloutResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationEventHandlerResult.GetFieldDeserializers()
    res["calloutDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalloutDateTime(val)
        }
        return nil
    }
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
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["httpStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpStatus(val)
        }
        return nil
    }
    res["numberOfAttempts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfAttempts(val)
        }
        return nil
    }
    return res
}
// GetHttpStatus gets the httpStatus property value. The httpStatus property
func (m *CustomExtensionCalloutResult) GetHttpStatus()(*int32) {
    return m.httpStatus
}
// GetNumberOfAttempts gets the numberOfAttempts property value. The numberOfAttempts property
func (m *CustomExtensionCalloutResult) GetNumberOfAttempts()(*int32) {
    return m.numberOfAttempts
}
// Serialize serializes information the current object
func (m *CustomExtensionCalloutResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationEventHandlerResult.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("calloutDateTime", m.GetCalloutDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customExtensionId", m.GetCustomExtensionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("httpStatus", m.GetHttpStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfAttempts", m.GetNumberOfAttempts())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCalloutDateTime sets the calloutDateTime property value. The calloutDateTime property
func (m *CustomExtensionCalloutResult) SetCalloutDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.calloutDateTime = value
}
// SetCustomExtensionId sets the customExtensionId property value. The customExtensionId property
func (m *CustomExtensionCalloutResult) SetCustomExtensionId(value *string)() {
    m.customExtensionId = value
}
// SetErrorCode sets the errorCode property value. The errorCode property
func (m *CustomExtensionCalloutResult) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetHttpStatus sets the httpStatus property value. The httpStatus property
func (m *CustomExtensionCalloutResult) SetHttpStatus(value *int32)() {
    m.httpStatus = value
}
// SetNumberOfAttempts sets the numberOfAttempts property value. The numberOfAttempts property
func (m *CustomExtensionCalloutResult) SetNumberOfAttempts(value *int32)() {
    m.numberOfAttempts = value
}
