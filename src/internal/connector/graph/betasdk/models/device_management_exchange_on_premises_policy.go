package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExchangeOnPremisesPolicy singleton entity which represents the Exchange OnPremises policy configured for a tenant.
type DeviceManagementExchangeOnPremisesPolicy struct {
    Entity
    // The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
    accessRules []DeviceManagementExchangeAccessRuleable
    // The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
    conditionalAccessSettings OnPremisesConditionalAccessSettingsable
    // Access Level in Exchange.
    defaultAccessLevel *DeviceManagementExchangeAccessLevel
    // The list of device classes known to Exchange
    knownDeviceClasses []DeviceManagementExchangeDeviceClassable
    // Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
    notificationContent []byte
}
// NewDeviceManagementExchangeOnPremisesPolicy instantiates a new deviceManagementExchangeOnPremisesPolicy and sets the default values.
func NewDeviceManagementExchangeOnPremisesPolicy()(*DeviceManagementExchangeOnPremisesPolicy) {
    m := &DeviceManagementExchangeOnPremisesPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementExchangeOnPremisesPolicy(), nil
}
// GetAccessRules gets the accessRules property value. The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
func (m *DeviceManagementExchangeOnPremisesPolicy) GetAccessRules()([]DeviceManagementExchangeAccessRuleable) {
    return m.accessRules
}
// GetConditionalAccessSettings gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagementExchangeOnPremisesPolicy) GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable) {
    return m.conditionalAccessSettings
}
// GetDefaultAccessLevel gets the defaultAccessLevel property value. Access Level in Exchange.
func (m *DeviceManagementExchangeOnPremisesPolicy) GetDefaultAccessLevel()(*DeviceManagementExchangeAccessLevel) {
    return m.defaultAccessLevel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeOnPremisesPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExchangeAccessRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeAccessRuleable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementExchangeAccessRuleable)
            }
            m.SetAccessRules(res)
        }
        return nil
    }
    res["conditionalAccessSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessSettings(val.(OnPremisesConditionalAccessSettingsable))
        }
        return nil
    }
    res["defaultAccessLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultAccessLevel(val.(*DeviceManagementExchangeAccessLevel))
        }
        return nil
    }
    res["knownDeviceClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExchangeDeviceClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeDeviceClassable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementExchangeDeviceClassable)
            }
            m.SetKnownDeviceClasses(res)
        }
        return nil
    }
    res["notificationContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationContent(val)
        }
        return nil
    }
    return res
}
// GetKnownDeviceClasses gets the knownDeviceClasses property value. The list of device classes known to Exchange
func (m *DeviceManagementExchangeOnPremisesPolicy) GetKnownDeviceClasses()([]DeviceManagementExchangeDeviceClassable) {
    return m.knownDeviceClasses
}
// GetNotificationContent gets the notificationContent property value. Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
func (m *DeviceManagementExchangeOnPremisesPolicy) GetNotificationContent()([]byte) {
    return m.notificationContent
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeOnPremisesPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessRules()))
        for i, v := range m.GetAccessRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditionalAccessSettings", m.GetConditionalAccessSettings())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultAccessLevel() != nil {
        cast := (*m.GetDefaultAccessLevel()).String()
        err = writer.WriteStringValue("defaultAccessLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetKnownDeviceClasses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKnownDeviceClasses()))
        for i, v := range m.GetKnownDeviceClasses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("knownDeviceClasses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("notificationContent", m.GetNotificationContent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessRules sets the accessRules property value. The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
func (m *DeviceManagementExchangeOnPremisesPolicy) SetAccessRules(value []DeviceManagementExchangeAccessRuleable)() {
    m.accessRules = value
}
// SetConditionalAccessSettings sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagementExchangeOnPremisesPolicy) SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)() {
    m.conditionalAccessSettings = value
}
// SetDefaultAccessLevel sets the defaultAccessLevel property value. Access Level in Exchange.
func (m *DeviceManagementExchangeOnPremisesPolicy) SetDefaultAccessLevel(value *DeviceManagementExchangeAccessLevel)() {
    m.defaultAccessLevel = value
}
// SetKnownDeviceClasses sets the knownDeviceClasses property value. The list of device classes known to Exchange
func (m *DeviceManagementExchangeOnPremisesPolicy) SetKnownDeviceClasses(value []DeviceManagementExchangeDeviceClassable)() {
    m.knownDeviceClasses = value
}
// SetNotificationContent sets the notificationContent property value. Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
func (m *DeviceManagementExchangeOnPremisesPolicy) SetNotificationContent(value []byte)() {
    m.notificationContent = value
}
