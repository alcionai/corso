package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyOperation the entity represents an group policy operation.
type GroupPolicyOperation struct {
    Entity
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Type of Group Policy operation status.
    operationStatus *GroupPolicyOperationStatus
    // Type of Group Policy operation.
    operationType *GroupPolicyOperationType
    // The group policy operation status detail.
    statusDetails *string
}
// NewGroupPolicyOperation instantiates a new groupPolicyOperation and sets the default values.
func NewGroupPolicyOperation()(*GroupPolicyOperation) {
    m := &GroupPolicyOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["operationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationStatus(val.(*GroupPolicyOperationStatus))
        }
        return nil
    }
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*GroupPolicyOperationType))
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyOperation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOperationStatus gets the operationStatus property value. Type of Group Policy operation status.
func (m *GroupPolicyOperation) GetOperationStatus()(*GroupPolicyOperationStatus) {
    return m.operationStatus
}
// GetOperationType gets the operationType property value. Type of Group Policy operation.
func (m *GroupPolicyOperation) GetOperationType()(*GroupPolicyOperationType) {
    return m.operationType
}
// GetStatusDetails gets the statusDetails property value. The group policy operation status detail.
func (m *GroupPolicyOperation) GetStatusDetails()(*string) {
    return m.statusDetails
}
// Serialize serializes information the current object
func (m *GroupPolicyOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOperationStatus() != nil {
        cast := (*m.GetOperationStatus()).String()
        err = writer.WriteStringValue("operationStatus", &cast)
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
        err = writer.WriteStringValue("statusDetails", m.GetStatusDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyOperation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOperationStatus sets the operationStatus property value. Type of Group Policy operation status.
func (m *GroupPolicyOperation) SetOperationStatus(value *GroupPolicyOperationStatus)() {
    m.operationStatus = value
}
// SetOperationType sets the operationType property value. Type of Group Policy operation.
func (m *GroupPolicyOperation) SetOperationType(value *GroupPolicyOperationType)() {
    m.operationType = value
}
// SetStatusDetails sets the statusDetails property value. The group policy operation status detail.
func (m *GroupPolicyOperation) SetStatusDetails(value *string)() {
    m.statusDetails = value
}
