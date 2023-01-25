package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivityHistoryItem provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ActivityHistoryItem struct {
    Entity
    // The activeDurationSeconds property
    activeDurationSeconds *int32
    // The activity property
    activity UserActivityable
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The expirationDateTime property
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastActiveDateTime property
    lastActiveDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The startedDateTime property
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *Status
    // The userTimezone property
    userTimezone *string
}
// NewActivityHistoryItem instantiates a new activityHistoryItem and sets the default values.
func NewActivityHistoryItem()(*ActivityHistoryItem) {
    m := &ActivityHistoryItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateActivityHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivityHistoryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivityHistoryItem(), nil
}
// GetActiveDurationSeconds gets the activeDurationSeconds property value. The activeDurationSeconds property
func (m *ActivityHistoryItem) GetActiveDurationSeconds()(*int32) {
    return m.activeDurationSeconds
}
// GetActivity gets the activity property value. The activity property
func (m *ActivityHistoryItem) GetActivity()(UserActivityable) {
    return m.activity
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ActivityHistoryItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *ActivityHistoryItem) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivityHistoryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDurationSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDurationSeconds(val)
        }
        return nil
    }
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(UserActivityable))
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
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["lastActiveDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActiveDateTime(val)
        }
        return nil
    }
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
        val, err := n.GetEnumValue(ParseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Status))
        }
        return nil
    }
    res["userTimezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserTimezone(val)
        }
        return nil
    }
    return res
}
// GetLastActiveDateTime gets the lastActiveDateTime property value. The lastActiveDateTime property
func (m *ActivityHistoryItem) GetLastActiveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActiveDateTime
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ActivityHistoryItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetStartedDateTime gets the startedDateTime property value. The startedDateTime property
func (m *ActivityHistoryItem) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetStatus gets the status property value. The status property
func (m *ActivityHistoryItem) GetStatus()(*Status) {
    return m.status
}
// GetUserTimezone gets the userTimezone property value. The userTimezone property
func (m *ActivityHistoryItem) GetUserTimezone()(*string) {
    return m.userTimezone
}
// Serialize serializes information the current object
func (m *ActivityHistoryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDurationSeconds", m.GetActiveDurationSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
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
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActiveDateTime", m.GetLastActiveDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userTimezone", m.GetUserTimezone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDurationSeconds sets the activeDurationSeconds property value. The activeDurationSeconds property
func (m *ActivityHistoryItem) SetActiveDurationSeconds(value *int32)() {
    m.activeDurationSeconds = value
}
// SetActivity sets the activity property value. The activity property
func (m *ActivityHistoryItem) SetActivity(value UserActivityable)() {
    m.activity = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ActivityHistoryItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *ActivityHistoryItem) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetLastActiveDateTime sets the lastActiveDateTime property value. The lastActiveDateTime property
func (m *ActivityHistoryItem) SetLastActiveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActiveDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ActivityHistoryItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetStartedDateTime sets the startedDateTime property value. The startedDateTime property
func (m *ActivityHistoryItem) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *ActivityHistoryItem) SetStatus(value *Status)() {
    m.status = value
}
// SetUserTimezone sets the userTimezone property value. The userTimezone property
func (m *ActivityHistoryItem) SetUserTimezone(value *string)() {
    m.userTimezone = value
}
