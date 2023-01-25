package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcess the user experience analytics device startup process details.
type UserExperienceAnalyticsDeviceStartupProcess struct {
    Entity
    // The user experience analytics device id.
    managedDeviceId *string
    // User experience analytics device startup process name.
    processName *string
    // The user experience analytics device startup process product name.
    productName *string
    // The User experience analytics device startup process publisher.
    publisher *string
    // User experience analytics device startup process impact in milliseconds.
    startupImpactInMs *int32
}
// NewUserExperienceAnalyticsDeviceStartupProcess instantiates a new userExperienceAnalyticsDeviceStartupProcess and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcess()(*UserExperienceAnalyticsDeviceStartupProcess) {
    m := &UserExperienceAnalyticsDeviceStartupProcess{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceStartupProcess(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["processName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessName(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["startupImpactInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupImpactInMs(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetProcessName gets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetProcessName()(*string) {
    return m.processName
}
// GetProductName gets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetProductName()(*string) {
    return m.productName
}
// GetPublisher gets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetPublisher()(*string) {
    return m.publisher
}
// GetStartupImpactInMs gets the startupImpactInMs property value. User experience analytics device startup process impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetStartupImpactInMs()(*int32) {
    return m.startupImpactInMs
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceStartupProcess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("processName", m.GetProcessName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("startupImpactInMs", m.GetStartupImpactInMs())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManagedDeviceId sets the managedDeviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetProcessName sets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetProcessName(value *string)() {
    m.processName = value
}
// SetProductName sets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetProductName(value *string)() {
    m.productName = value
}
// SetPublisher sets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetPublisher(value *string)() {
    m.publisher = value
}
// SetStartupImpactInMs sets the startupImpactInMs property value. User experience analytics device startup process impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetStartupImpactInMs(value *int32)() {
    m.startupImpactInMs = value
}
