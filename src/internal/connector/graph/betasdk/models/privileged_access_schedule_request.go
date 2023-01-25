package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessScheduleRequest 
type PrivilegedAccessScheduleRequest struct {
    Request
    // The action property
    action *ScheduleRequestActions
    // The isValidationOnly property
    isValidationOnly *bool
    // The justification property
    justification *string
    // The scheduleInfo property
    scheduleInfo RequestScheduleable
    // The ticketInfo property
    ticketInfo TicketInfoable
}
// NewPrivilegedAccessScheduleRequest instantiates a new PrivilegedAccessScheduleRequest and sets the default values.
func NewPrivilegedAccessScheduleRequest()(*PrivilegedAccessScheduleRequest) {
    m := &PrivilegedAccessScheduleRequest{
        Request: *NewRequest(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessScheduleRequest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrivilegedAccessScheduleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessScheduleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest":
                        return NewPrivilegedAccessGroupAssignmentScheduleRequest(), nil
                    case "#microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest":
                        return NewPrivilegedAccessGroupEligibilityScheduleRequest(), nil
                }
            }
        }
    }
    return NewPrivilegedAccessScheduleRequest(), nil
}
// GetAction gets the action property value. The action property
func (m *PrivilegedAccessScheduleRequest) GetAction()(*ScheduleRequestActions) {
    return m.action
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessScheduleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScheduleRequestActions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*ScheduleRequestActions))
        }
        return nil
    }
    res["isValidationOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
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
    res["ticketInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTicketInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketInfo(val.(TicketInfoable))
        }
        return nil
    }
    return res
}
// GetIsValidationOnly gets the isValidationOnly property value. The isValidationOnly property
func (m *PrivilegedAccessScheduleRequest) GetIsValidationOnly()(*bool) {
    return m.isValidationOnly
}
// GetJustification gets the justification property value. The justification property
func (m *PrivilegedAccessScheduleRequest) GetJustification()(*string) {
    return m.justification
}
// GetScheduleInfo gets the scheduleInfo property value. The scheduleInfo property
func (m *PrivilegedAccessScheduleRequest) GetScheduleInfo()(RequestScheduleable) {
    return m.scheduleInfo
}
// GetTicketInfo gets the ticketInfo property value. The ticketInfo property
func (m *PrivilegedAccessScheduleRequest) GetTicketInfo()(TicketInfoable) {
    return m.ticketInfo
}
// Serialize serializes information the current object
func (m *PrivilegedAccessScheduleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Request.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValidationOnly", m.GetIsValidationOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
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
        err = writer.WriteObjectValue("ticketInfo", m.GetTicketInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *PrivilegedAccessScheduleRequest) SetAction(value *ScheduleRequestActions)() {
    m.action = value
}
// SetIsValidationOnly sets the isValidationOnly property value. The isValidationOnly property
func (m *PrivilegedAccessScheduleRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// SetJustification sets the justification property value. The justification property
func (m *PrivilegedAccessScheduleRequest) SetJustification(value *string)() {
    m.justification = value
}
// SetScheduleInfo sets the scheduleInfo property value. The scheduleInfo property
func (m *PrivilegedAccessScheduleRequest) SetScheduleInfo(value RequestScheduleable)() {
    m.scheduleInfo = value
}
// SetTicketInfo sets the ticketInfo property value. The ticketInfo property
func (m *PrivilegedAccessScheduleRequest) SetTicketInfo(value TicketInfoable)() {
    m.ticketInfo = value
}
