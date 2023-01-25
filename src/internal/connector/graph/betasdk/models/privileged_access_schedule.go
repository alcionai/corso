package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessSchedule 
type PrivilegedAccessSchedule struct {
    Entity
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The createdUsing property
    createdUsing *string
    // The modifiedDateTime property
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The scheduleInfo property
    scheduleInfo RequestScheduleable
    // The status property
    status *string
}
// NewPrivilegedAccessSchedule instantiates a new privilegedAccessSchedule and sets the default values.
func NewPrivilegedAccessSchedule()(*PrivilegedAccessSchedule) {
    m := &PrivilegedAccessSchedule{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedAccessScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.privilegedAccessGroupAssignmentSchedule":
                        return NewPrivilegedAccessGroupAssignmentSchedule(), nil
                    case "#microsoft.graph.privilegedAccessGroupEligibilitySchedule":
                        return NewPrivilegedAccessGroupEligibilitySchedule(), nil
                }
            }
        }
    }
    return NewPrivilegedAccessSchedule(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *PrivilegedAccessSchedule) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreatedUsing gets the createdUsing property value. The createdUsing property
func (m *PrivilegedAccessSchedule) GetCreatedUsing()(*string) {
    return m.createdUsing
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["createdUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedUsing(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["scheduleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(RequestScheduleable))
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
// GetModifiedDateTime gets the modifiedDateTime property value. The modifiedDateTime property
func (m *PrivilegedAccessSchedule) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetScheduleInfo gets the scheduleInfo property value. The scheduleInfo property
func (m *PrivilegedAccessSchedule) GetScheduleInfo()(RequestScheduleable) {
    return m.scheduleInfo
}
// GetStatus gets the status property value. The status property
func (m *PrivilegedAccessSchedule) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *PrivilegedAccessSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdUsing", m.GetCreatedUsing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
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
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *PrivilegedAccessSchedule) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedUsing sets the createdUsing property value. The createdUsing property
func (m *PrivilegedAccessSchedule) SetCreatedUsing(value *string)() {
    m.createdUsing = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The modifiedDateTime property
func (m *PrivilegedAccessSchedule) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetScheduleInfo sets the scheduleInfo property value. The scheduleInfo property
func (m *PrivilegedAccessSchedule) SetScheduleInfo(value RequestScheduleable)() {
    m.scheduleInfo = value
}
// SetStatus sets the status property value. The status property
func (m *PrivilegedAccessSchedule) SetStatus(value *string)() {
    m.status = value
}
