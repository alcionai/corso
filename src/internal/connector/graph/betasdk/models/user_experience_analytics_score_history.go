package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsScoreHistory the user experience analytics device startup score history.
type UserExperienceAnalyticsScoreHistory struct {
    Entity
    // The user experience analytics device startup date time.
    startupDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUserExperienceAnalyticsScoreHistory instantiates a new userExperienceAnalyticsScoreHistory and sets the default values.
func NewUserExperienceAnalyticsScoreHistory()(*UserExperienceAnalyticsScoreHistory) {
    m := &UserExperienceAnalyticsScoreHistory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsScoreHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsScoreHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsScoreHistory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsScoreHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["startupDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupDateTime(val)
        }
        return nil
    }
    return res
}
// GetStartupDateTime gets the startupDateTime property value. The user experience analytics device startup date time.
func (m *UserExperienceAnalyticsScoreHistory) GetStartupDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startupDateTime
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsScoreHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("startupDateTime", m.GetStartupDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStartupDateTime sets the startupDateTime property value. The user experience analytics device startup date time.
func (m *UserExperienceAnalyticsScoreHistory) SetStartupDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startupDateTime = value
}
