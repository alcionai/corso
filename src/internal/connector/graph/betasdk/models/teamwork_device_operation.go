package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceOperation provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TeamworkDeviceOperation struct {
    Entity
    // Time at which the operation reached a final state (for example, Successful, Failed, and Cancelled).
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identity of the user who created the device operation.
    createdBy IdentitySetable
    // The UTC date and time when the device operation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Error details are available only in case of a failed status.
    error OperationErrorable
    // Identity of the user who last modified the device operation.
    lastActionBy IdentitySetable
    // The UTC date and time when the device operation was last modified.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The operationType property
    operationType *TeamworkDeviceOperationType
    // Time at which the operation was started.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The current status of the async operation, for example, Queued, Scheduled, InProgress,  Successful, Cancelled, and Failed.
    status *string
}
// NewTeamworkDeviceOperation instantiates a new teamworkDeviceOperation and sets the default values.
func NewTeamworkDeviceOperation()(*TeamworkDeviceOperation) {
    m := &TeamworkDeviceOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkDeviceOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDeviceOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDeviceOperation(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. Time at which the operation reached a final state (for example, Successful, Failed, and Cancelled).
func (m *TeamworkDeviceOperation) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device operation.
func (m *TeamworkDeviceOperation) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device operation was created.
func (m *TeamworkDeviceOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetError gets the error property value. Error details are available only in case of a failed status.
func (m *TeamworkDeviceOperation) GetError()(OperationErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(OperationErrorable))
        }
        return nil
    }
    res["lastActionBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*TeamworkDeviceOperationType))
        }
        return nil
    }
    res["startedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetLastActionBy gets the lastActionBy property value. Identity of the user who last modified the device operation.
func (m *TeamworkDeviceOperation) GetLastActionBy()(IdentitySetable) {
    return m.lastActionBy
}
// GetLastActionDateTime gets the lastActionDateTime property value. The UTC date and time when the device operation was last modified.
func (m *TeamworkDeviceOperation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetOperationType gets the operationType property value. The operationType property
func (m *TeamworkDeviceOperation) GetOperationType()(*TeamworkDeviceOperationType) {
    return m.operationType
}
// GetStartedDateTime gets the startedDateTime property value. Time at which the operation was started.
func (m *TeamworkDeviceOperation) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetStatus gets the status property value. The current status of the async operation, for example, Queued, Scheduled, InProgress,  Successful, Cancelled, and Failed.
func (m *TeamworkDeviceOperation) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *TeamworkDeviceOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastActionBy", m.GetLastActionBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startedDateTime", m.GetStartedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. Time at which the operation reached a final state (for example, Successful, Failed, and Cancelled).
func (m *TeamworkDeviceOperation) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device operation.
func (m *TeamworkDeviceOperation) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device operation was created.
func (m *TeamworkDeviceOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetError sets the error property value. Error details are available only in case of a failed status.
func (m *TeamworkDeviceOperation) SetError(value OperationErrorable)() {
    m.error = value
}
// SetLastActionBy sets the lastActionBy property value. Identity of the user who last modified the device operation.
func (m *TeamworkDeviceOperation) SetLastActionBy(value IdentitySetable)() {
    m.lastActionBy = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The UTC date and time when the device operation was last modified.
func (m *TeamworkDeviceOperation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetOperationType sets the operationType property value. The operationType property
func (m *TeamworkDeviceOperation) SetOperationType(value *TeamworkDeviceOperationType)() {
    m.operationType = value
}
// SetStartedDateTime sets the startedDateTime property value. Time at which the operation was started.
func (m *TeamworkDeviceOperation) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetStatus sets the status property value. The current status of the async operation, for example, Queued, Scheduled, InProgress,  Successful, Cancelled, and Failed.
func (m *TeamworkDeviceOperation) SetStatus(value *string)() {
    m.status = value
}
