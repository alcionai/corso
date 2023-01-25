package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComplianceActionItem scheduled Action for Rule
type DeviceManagementComplianceActionItem struct {
    Entity
    // Scheduled Action Type Enum
    actionType *DeviceManagementComplianceActionType
    // Number of hours to wait till the action will be enforced. Valid values 0 to 8760
    gracePeriodHours *int32
    // A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
    notificationMessageCCList []string
    // What notification Message template to use
    notificationTemplateId *string
}
// NewDeviceManagementComplianceActionItem instantiates a new deviceManagementComplianceActionItem and sets the default values.
func NewDeviceManagementComplianceActionItem()(*DeviceManagementComplianceActionItem) {
    m := &DeviceManagementComplianceActionItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementComplianceActionItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComplianceActionItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComplianceActionItem(), nil
}
// GetActionType gets the actionType property value. Scheduled Action Type Enum
func (m *DeviceManagementComplianceActionItem) GetActionType()(*DeviceManagementComplianceActionType) {
    return m.actionType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplianceActionItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementComplianceActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val.(*DeviceManagementComplianceActionType))
        }
        return nil
    }
    res["gracePeriodHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodHours(val)
        }
        return nil
    }
    res["notificationMessageCCList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotificationMessageCCList(res)
        }
        return nil
    }
    res["notificationTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTemplateId(val)
        }
        return nil
    }
    return res
}
// GetGracePeriodHours gets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
func (m *DeviceManagementComplianceActionItem) GetGracePeriodHours()(*int32) {
    return m.gracePeriodHours
}
// GetNotificationMessageCCList gets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceActionItem) GetNotificationMessageCCList()([]string) {
    return m.notificationMessageCCList
}
// GetNotificationTemplateId gets the notificationTemplateId property value. What notification Message template to use
func (m *DeviceManagementComplianceActionItem) GetNotificationTemplateId()(*string) {
    return m.notificationTemplateId
}
// Serialize serializes information the current object
func (m *DeviceManagementComplianceActionItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionType() != nil {
        cast := (*m.GetActionType()).String()
        err = writer.WriteStringValue("actionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("gracePeriodHours", m.GetGracePeriodHours())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationMessageCCList() != nil {
        err = writer.WriteCollectionOfStringValues("notificationMessageCCList", m.GetNotificationMessageCCList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationTemplateId", m.GetNotificationTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionType sets the actionType property value. Scheduled Action Type Enum
func (m *DeviceManagementComplianceActionItem) SetActionType(value *DeviceManagementComplianceActionType)() {
    m.actionType = value
}
// SetGracePeriodHours sets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
func (m *DeviceManagementComplianceActionItem) SetGracePeriodHours(value *int32)() {
    m.gracePeriodHours = value
}
// SetNotificationMessageCCList sets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceActionItem) SetNotificationMessageCCList(value []string)() {
    m.notificationMessageCCList = value
}
// SetNotificationTemplateId sets the notificationTemplateId property value. What notification Message template to use
func (m *DeviceManagementComplianceActionItem) SetNotificationTemplateId(value *string)() {
    m.notificationTemplateId = value
}
