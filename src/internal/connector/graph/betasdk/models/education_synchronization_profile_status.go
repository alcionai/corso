package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationProfileStatus 
type EducationSynchronizationProfileStatus struct {
    Entity
    // Number of errors during synchronization.
    errorCount *int64
    // Date and time when most recent changes were observed in the profile.
    lastActivityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time of the most recent successful synchronization.
    lastSynchronizationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status of a sync. The possible values are: paused, inProgress, success, error, validationError, quarantined, unknownFutureValue, extracting, validating. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: extracting, validating.
    status *EducationSynchronizationStatus
    // Status message for the synchronization stage of the current profile.
    statusMessage *string
}
// NewEducationSynchronizationProfileStatus instantiates a new educationSynchronizationProfileStatus and sets the default values.
func NewEducationSynchronizationProfileStatus()(*EducationSynchronizationProfileStatus) {
    m := &EducationSynchronizationProfileStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSynchronizationProfileStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationProfileStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationProfileStatus(), nil
}
// GetErrorCount gets the errorCount property value. Number of errors during synchronization.
func (m *EducationSynchronizationProfileStatus) GetErrorCount()(*int64) {
    return m.errorCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationProfileStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["lastActivityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDateTime(val)
        }
        return nil
    }
    res["lastSynchronizationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSynchronizationDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSynchronizationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*EducationSynchronizationStatus))
        }
        return nil
    }
    res["statusMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusMessage(val)
        }
        return nil
    }
    return res
}
// GetLastActivityDateTime gets the lastActivityDateTime property value. Date and time when most recent changes were observed in the profile.
func (m *EducationSynchronizationProfileStatus) GetLastActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActivityDateTime
}
// GetLastSynchronizationDateTime gets the lastSynchronizationDateTime property value. Date and time of the most recent successful synchronization.
func (m *EducationSynchronizationProfileStatus) GetLastSynchronizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSynchronizationDateTime
}
// GetStatus gets the status property value. The status of a sync. The possible values are: paused, inProgress, success, error, validationError, quarantined, unknownFutureValue, extracting, validating. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: extracting, validating.
func (m *EducationSynchronizationProfileStatus) GetStatus()(*EducationSynchronizationStatus) {
    return m.status
}
// GetStatusMessage gets the statusMessage property value. Status message for the synchronization stage of the current profile.
func (m *EducationSynchronizationProfileStatus) GetStatusMessage()(*string) {
    return m.statusMessage
}
// Serialize serializes information the current object
func (m *EducationSynchronizationProfileStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActivityDateTime", m.GetLastActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSynchronizationDateTime", m.GetLastSynchronizationDateTime())
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
    {
        err = writer.WriteStringValue("statusMessage", m.GetStatusMessage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorCount sets the errorCount property value. Number of errors during synchronization.
func (m *EducationSynchronizationProfileStatus) SetErrorCount(value *int64)() {
    m.errorCount = value
}
// SetLastActivityDateTime sets the lastActivityDateTime property value. Date and time when most recent changes were observed in the profile.
func (m *EducationSynchronizationProfileStatus) SetLastActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActivityDateTime = value
}
// SetLastSynchronizationDateTime sets the lastSynchronizationDateTime property value. Date and time of the most recent successful synchronization.
func (m *EducationSynchronizationProfileStatus) SetLastSynchronizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSynchronizationDateTime = value
}
// SetStatus sets the status property value. The status of a sync. The possible values are: paused, inProgress, success, error, validationError, quarantined, unknownFutureValue, extracting, validating. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: extracting, validating.
func (m *EducationSynchronizationProfileStatus) SetStatus(value *EducationSynchronizationStatus)() {
    m.status = value
}
// SetStatusMessage sets the statusMessage property value. Status message for the synchronization stage of the current profile.
func (m *EducationSynchronizationProfileStatus) SetStatusMessage(value *string)() {
    m.statusMessage = value
}
