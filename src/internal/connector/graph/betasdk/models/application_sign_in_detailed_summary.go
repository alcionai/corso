package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationSignInDetailedSummary 
type ApplicationSignInDetailedSummary struct {
    Entity
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    aggregatedEventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the application that the user signed in to.
    appDisplayName *string
    // ID of the application that the user signed in to.
    appId *string
    // Count of sign-ins made by the application.
    signInCount *int64
    // Details of the sign-in status.
    status SignInStatusable
}
// NewApplicationSignInDetailedSummary instantiates a new ApplicationSignInDetailedSummary and sets the default values.
func NewApplicationSignInDetailedSummary()(*ApplicationSignInDetailedSummary) {
    m := &ApplicationSignInDetailedSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApplicationSignInDetailedSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationSignInDetailedSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationSignInDetailedSummary(), nil
}
// GetAggregatedEventDateTime gets the aggregatedEventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ApplicationSignInDetailedSummary) GetAggregatedEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.aggregatedEventDateTime
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppId gets the appId property value. ID of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationSignInDetailedSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedEventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAggregatedEventDateTime(val)
        }
        return nil
    }
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["signInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInCount(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(SignInStatusable))
        }
        return nil
    }
    return res
}
// GetSignInCount gets the signInCount property value. Count of sign-ins made by the application.
func (m *ApplicationSignInDetailedSummary) GetSignInCount()(*int64) {
    return m.signInCount
}
// GetStatus gets the status property value. Details of the sign-in status.
func (m *ApplicationSignInDetailedSummary) GetStatus()(SignInStatusable) {
    return m.status
}
// Serialize serializes information the current object
func (m *ApplicationSignInDetailedSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("aggregatedEventDateTime", m.GetAggregatedEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("signInCount", m.GetSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAggregatedEventDateTime sets the aggregatedEventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ApplicationSignInDetailedSummary) SetAggregatedEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.aggregatedEventDateTime = value
}
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppId sets the appId property value. ID of the application that the user signed in to.
func (m *ApplicationSignInDetailedSummary) SetAppId(value *string)() {
    m.appId = value
}
// SetSignInCount sets the signInCount property value. Count of sign-ins made by the application.
func (m *ApplicationSignInDetailedSummary) SetSignInCount(value *int64)() {
    m.signInCount = value
}
// SetStatus sets the status property value. Details of the sign-in status.
func (m *ApplicationSignInDetailedSummary) SetStatus(value SignInStatusable)() {
    m.status = value
}
