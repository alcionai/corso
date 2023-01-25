package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppLogCollectionRequest appLogCollectionRequest Entity.
type AppLogCollectionRequest struct {
    Entity
    // Time at which the upload log request reached a terminal state
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of log folders.
    customLogFolders []string
    // Error message if any during the upload process
    errorMessage *string
    // AppLogUploadStatus
    status *AppLogUploadState
}
// NewAppLogCollectionRequest instantiates a new appLogCollectionRequest and sets the default values.
func NewAppLogCollectionRequest()(*AppLogCollectionRequest) {
    m := &AppLogCollectionRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppLogCollectionRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppLogCollectionRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppLogCollectionRequest(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. Time at which the upload log request reached a terminal state
func (m *AppLogCollectionRequest) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCustomLogFolders gets the customLogFolders property value. List of log folders.
func (m *AppLogCollectionRequest) GetCustomLogFolders()([]string) {
    return m.customLogFolders
}
// GetErrorMessage gets the errorMessage property value. Error message if any during the upload process
func (m *AppLogCollectionRequest) GetErrorMessage()(*string) {
    return m.errorMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppLogCollectionRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["customLogFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCustomLogFolders(res)
        }
        return nil
    }
    res["errorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLogUploadState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AppLogUploadState))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. AppLogUploadStatus
func (m *AppLogCollectionRequest) GetStatus()(*AppLogUploadState) {
    return m.status
}
// Serialize serializes information the current object
func (m *AppLogCollectionRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCustomLogFolders() != nil {
        err = writer.WriteCollectionOfStringValues("customLogFolders", m.GetCustomLogFolders())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorMessage", m.GetErrorMessage())
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
// SetCompletedDateTime sets the completedDateTime property value. Time at which the upload log request reached a terminal state
func (m *AppLogCollectionRequest) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCustomLogFolders sets the customLogFolders property value. List of log folders.
func (m *AppLogCollectionRequest) SetCustomLogFolders(value []string)() {
    m.customLogFolders = value
}
// SetErrorMessage sets the errorMessage property value. Error message if any during the upload process
func (m *AppLogCollectionRequest) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// SetStatus sets the status property value. AppLogUploadStatus
func (m *AppLogCollectionRequest) SetStatus(value *AppLogUploadState)() {
    m.status = value
}
