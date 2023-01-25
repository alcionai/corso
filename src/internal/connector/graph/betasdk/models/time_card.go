package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCard 
type TimeCard struct {
    ChangeTrackedEntity
    // The list of breaks associated with the timeCard.
    breaks []TimeCardBreakable
    // The clock-in event of the timeCard.
    clockInEvent TimeCardEventable
    // The clock-out event of the timeCard.
    clockOutEvent TimeCardEventable
    // Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
    confirmedBy *ConfirmedBy
    // Notes about the timeCard.
    notes ItemBodyable
    // The original timeCardEntry of the timeCard, before user edits.
    originalEntry TimeCardEntryable
    // The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
    state *TimeCardState
    // User ID to which  the timeCard belongs.
    userId *string
}
// NewTimeCard instantiates a new TimeCard and sets the default values.
func NewTimeCard()(*TimeCard) {
    m := &TimeCard{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.timeCard";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTimeCardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeCard(), nil
}
// GetBreaks gets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCard) GetBreaks()([]TimeCardBreakable) {
    return m.breaks
}
// GetClockInEvent gets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCard) GetClockInEvent()(TimeCardEventable) {
    return m.clockInEvent
}
// GetClockOutEvent gets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCard) GetClockOutEvent()(TimeCardEventable) {
    return m.clockOutEvent
}
// GetConfirmedBy gets the confirmedBy property value. Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
func (m *TimeCard) GetConfirmedBy()(*ConfirmedBy) {
    return m.confirmedBy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCard) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["breaks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeCardBreakFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeCardBreakable, len(val))
            for i, v := range val {
                res[i] = v.(TimeCardBreakable)
            }
            m.SetBreaks(res)
        }
        return nil
    }
    res["clockInEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockInEvent(val.(TimeCardEventable))
        }
        return nil
    }
    res["clockOutEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClockOutEvent(val.(TimeCardEventable))
        }
        return nil
    }
    res["confirmedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfirmedBy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfirmedBy(val.(*ConfirmedBy))
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    res["originalEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalEntry(val.(TimeCardEntryable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTimeCardState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*TimeCardState))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. Notes about the timeCard.
func (m *TimeCard) GetNotes()(ItemBodyable) {
    return m.notes
}
// GetOriginalEntry gets the originalEntry property value. The original timeCardEntry of the timeCard, before user edits.
func (m *TimeCard) GetOriginalEntry()(TimeCardEntryable) {
    return m.originalEntry
}
// GetState gets the state property value. The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
func (m *TimeCard) GetState()(*TimeCardState) {
    return m.state
}
// GetUserId gets the userId property value. User ID to which  the timeCard belongs.
func (m *TimeCard) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *TimeCard) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBreaks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBreaks()))
        for i, v := range m.GetBreaks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("breaks", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clockInEvent", m.GetClockInEvent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clockOutEvent", m.GetClockOutEvent())
        if err != nil {
            return err
        }
    }
    if m.GetConfirmedBy() != nil {
        cast := (*m.GetConfirmedBy()).String()
        err = writer.WriteStringValue("confirmedBy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("originalEntry", m.GetOriginalEntry())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBreaks sets the breaks property value. The list of breaks associated with the timeCard.
func (m *TimeCard) SetBreaks(value []TimeCardBreakable)() {
    m.breaks = value
}
// SetClockInEvent sets the clockInEvent property value. The clock-in event of the timeCard.
func (m *TimeCard) SetClockInEvent(value TimeCardEventable)() {
    m.clockInEvent = value
}
// SetClockOutEvent sets the clockOutEvent property value. The clock-out event of the timeCard.
func (m *TimeCard) SetClockOutEvent(value TimeCardEventable)() {
    m.clockOutEvent = value
}
// SetConfirmedBy sets the confirmedBy property value. Indicate if this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
func (m *TimeCard) SetConfirmedBy(value *ConfirmedBy)() {
    m.confirmedBy = value
}
// SetNotes sets the notes property value. Notes about the timeCard.
func (m *TimeCard) SetNotes(value ItemBodyable)() {
    m.notes = value
}
// SetOriginalEntry sets the originalEntry property value. The original timeCardEntry of the timeCard, before user edits.
func (m *TimeCard) SetOriginalEntry(value TimeCardEntryable)() {
    m.originalEntry = value
}
// SetState sets the state property value. The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut, unknownFutureValue.
func (m *TimeCard) SetState(value *TimeCardState)() {
    m.state = value
}
// SetUserId sets the userId property value. User ID to which  the timeCard belongs.
func (m *TimeCard) SetUserId(value *string)() {
    m.userId = value
}
