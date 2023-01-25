package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthRuntimeDetails 
type UserExperienceAnalyticsBatteryHealthRuntimeDetails struct {
    Entity
    // Number of active devices within the tenant. Valid values -2147483648 to 2147483647
    activeDevices *int32
    // Number of devices whose active runtime is greater than 3 hours but lesser than 5 hours. Valid values -2147483648 to 2147483647
    batteryRuntimeFair *int32
    // Number of devices  whose active runtime is greater than 5 hours. Valid values -2147483648 to 2147483647
    batteryRuntimeGood *int32
    // Number of devices whose active runtime is lesser than 3 hours. Valid values -2147483648 to 2147483647
    batteryRuntimePoor *int32
    // Recorded date time of this runtime details instance.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUserExperienceAnalyticsBatteryHealthRuntimeDetails instantiates a new userExperienceAnalyticsBatteryHealthRuntimeDetails and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthRuntimeDetails()(*UserExperienceAnalyticsBatteryHealthRuntimeDetails) {
    m := &UserExperienceAnalyticsBatteryHealthRuntimeDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthRuntimeDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBatteryHealthRuntimeDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthRuntimeDetails(), nil
}
// GetActiveDevices gets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetActiveDevices()(*int32) {
    return m.activeDevices
}
// GetBatteryRuntimeFair gets the batteryRuntimeFair property value. Number of devices whose active runtime is greater than 3 hours but lesser than 5 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetBatteryRuntimeFair()(*int32) {
    return m.batteryRuntimeFair
}
// GetBatteryRuntimeGood gets the batteryRuntimeGood property value. Number of devices  whose active runtime is greater than 5 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetBatteryRuntimeGood()(*int32) {
    return m.batteryRuntimeGood
}
// GetBatteryRuntimePoor gets the batteryRuntimePoor property value. Number of devices whose active runtime is lesser than 3 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetBatteryRuntimePoor()(*int32) {
    return m.batteryRuntimePoor
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDevices(val)
        }
        return nil
    }
    res["batteryRuntimeFair"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuntimeFair(val)
        }
        return nil
    }
    res["batteryRuntimeGood"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuntimeGood(val)
        }
        return nil
    }
    res["batteryRuntimePoor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuntimePoor(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Recorded date time of this runtime details instance.
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDevices", m.GetActiveDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryRuntimeFair", m.GetBatteryRuntimeFair())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryRuntimeGood", m.GetBatteryRuntimeGood())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryRuntimePoor", m.GetBatteryRuntimePoor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDevices sets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetActiveDevices(value *int32)() {
    m.activeDevices = value
}
// SetBatteryRuntimeFair sets the batteryRuntimeFair property value. Number of devices whose active runtime is greater than 3 hours but lesser than 5 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetBatteryRuntimeFair(value *int32)() {
    m.batteryRuntimeFair = value
}
// SetBatteryRuntimeGood sets the batteryRuntimeGood property value. Number of devices  whose active runtime is greater than 5 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetBatteryRuntimeGood(value *int32)() {
    m.batteryRuntimeGood = value
}
// SetBatteryRuntimePoor sets the batteryRuntimePoor property value. Number of devices whose active runtime is lesser than 3 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetBatteryRuntimePoor(value *int32)() {
    m.batteryRuntimePoor = value
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Recorded date time of this runtime details instance.
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
