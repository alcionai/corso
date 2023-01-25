package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationError provides operations to call the add method.
type EducationSynchronizationError struct {
    Entity
    // Represents the sync entity (school, section, student, teacher).
    entryType *string
    // Represents the error code for this error.
    errorCode *string
    // Contains a description of the error.
    errorMessage *string
    // The unique identifier for the entry.
    joiningValue *string
    // The time of occurrence of this error.
    recordedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identifier of this error entry.
    reportableIdentifier *string
}
// NewEducationSynchronizationError instantiates a new educationSynchronizationError and sets the default values.
func NewEducationSynchronizationError()(*EducationSynchronizationError) {
    m := &EducationSynchronizationError{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSynchronizationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationError(), nil
}
// GetEntryType gets the entryType property value. Represents the sync entity (school, section, student, teacher).
func (m *EducationSynchronizationError) GetEntryType()(*string) {
    return m.entryType
}
// GetErrorCode gets the errorCode property value. Represents the error code for this error.
func (m *EducationSynchronizationError) GetErrorCode()(*string) {
    return m.errorCode
}
// GetErrorMessage gets the errorMessage property value. Contains a description of the error.
func (m *EducationSynchronizationError) GetErrorMessage()(*string) {
    return m.errorMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntryType(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
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
    res["joiningValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoiningValue(val)
        }
        return nil
    }
    res["recordedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordedDateTime(val)
        }
        return nil
    }
    res["reportableIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportableIdentifier(val)
        }
        return nil
    }
    return res
}
// GetJoiningValue gets the joiningValue property value. The unique identifier for the entry.
func (m *EducationSynchronizationError) GetJoiningValue()(*string) {
    return m.joiningValue
}
// GetRecordedDateTime gets the recordedDateTime property value. The time of occurrence of this error.
func (m *EducationSynchronizationError) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.recordedDateTime
}
// GetReportableIdentifier gets the reportableIdentifier property value. The identifier of this error entry.
func (m *EducationSynchronizationError) GetReportableIdentifier()(*string) {
    return m.reportableIdentifier
}
// Serialize serializes information the current object
func (m *EducationSynchronizationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("entryType", m.GetEntryType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
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
    {
        err = writer.WriteStringValue("joiningValue", m.GetJoiningValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("recordedDateTime", m.GetRecordedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportableIdentifier", m.GetReportableIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEntryType sets the entryType property value. Represents the sync entity (school, section, student, teacher).
func (m *EducationSynchronizationError) SetEntryType(value *string)() {
    m.entryType = value
}
// SetErrorCode sets the errorCode property value. Represents the error code for this error.
func (m *EducationSynchronizationError) SetErrorCode(value *string)() {
    m.errorCode = value
}
// SetErrorMessage sets the errorMessage property value. Contains a description of the error.
func (m *EducationSynchronizationError) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// SetJoiningValue sets the joiningValue property value. The unique identifier for the entry.
func (m *EducationSynchronizationError) SetJoiningValue(value *string)() {
    m.joiningValue = value
}
// SetRecordedDateTime sets the recordedDateTime property value. The time of occurrence of this error.
func (m *EducationSynchronizationError) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recordedDateTime = value
}
// SetReportableIdentifier sets the reportableIdentifier property value. The identifier of this error entry.
func (m *EducationSynchronizationError) SetReportableIdentifier(value *string)() {
    m.reportableIdentifier = value
}
