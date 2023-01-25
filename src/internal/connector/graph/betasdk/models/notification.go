package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Notification provides operations to manage the collection of site entities.
type Notification struct {
    Entity
    // Sets how long (in seconds) this notification content will stay in each platform's notification viewer. For example, when the notification is delivered to a Windows device, the value of this property is passed on to ToastNotification.ExpirationTime, which determines how long the toast notification will stay in the user's Windows Action Center.
    displayTimeToLive *int32
    // Sets a UTC expiration date and time on a user notification using ISO 8601 format (for example, midnight UTC on Jan 1, 2019 would look like this: '2019-01-01T00:00:00Z'). When time is up, the notification is removed from the Microsoft Graph notification feed store completely and is no longer part of notification history. Max value is 30 days.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the group that this notification belongs to. It is set by the developer for the purpose of grouping notifications together.
    groupName *string
    // The payload property
    payload PayloadTypesable
    // Indicates the priority of a raw user notification. Visual notifications are sent with high priority by default. Valid values are None, High and Low.
    priority *Priority
    // Represents the host name of the app to which the calling service wants to post the notification, for the given user. If targeting web endpoints (see targetPolicy.platformTypes), ensure that targetHostName is the same as the name used when creating a subscription on the client side within the application JSON property.
    targetHostName *string
    // Target policy object handles notification delivery policy for endpoint types that should be targeted (Windows, iOS, Android and WebPush) for the given user.
    targetPolicy TargetPolicyEndpointsable
}
// NewNotification instantiates a new notification and sets the default values.
func NewNotification()(*Notification) {
    m := &Notification{
        Entity: *NewEntity(),
    }
    return m
}
// CreateNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotification(), nil
}
// GetDisplayTimeToLive gets the displayTimeToLive property value. Sets how long (in seconds) this notification content will stay in each platform's notification viewer. For example, when the notification is delivered to a Windows device, the value of this property is passed on to ToastNotification.ExpirationTime, which determines how long the toast notification will stay in the user's Windows Action Center.
func (m *Notification) GetDisplayTimeToLive()(*int32) {
    return m.displayTimeToLive
}
// GetExpirationDateTime gets the expirationDateTime property value. Sets a UTC expiration date and time on a user notification using ISO 8601 format (for example, midnight UTC on Jan 1, 2019 would look like this: '2019-01-01T00:00:00Z'). When time is up, the notification is removed from the Microsoft Graph notification feed store completely and is no longer part of notification history. Max value is 30 days.
func (m *Notification) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Notification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayTimeToLive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayTimeToLive(val)
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
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadTypesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(PayloadTypesable))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*Priority))
        }
        return nil
    }
    res["targetHostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetHostName(val)
        }
        return nil
    }
    res["targetPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTargetPolicyEndpointsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPolicy(val.(TargetPolicyEndpointsable))
        }
        return nil
    }
    return res
}
// GetGroupName gets the groupName property value. The name of the group that this notification belongs to. It is set by the developer for the purpose of grouping notifications together.
func (m *Notification) GetGroupName()(*string) {
    return m.groupName
}
// GetPayload gets the payload property value. The payload property
func (m *Notification) GetPayload()(PayloadTypesable) {
    return m.payload
}
// GetPriority gets the priority property value. Indicates the priority of a raw user notification. Visual notifications are sent with high priority by default. Valid values are None, High and Low.
func (m *Notification) GetPriority()(*Priority) {
    return m.priority
}
// GetTargetHostName gets the targetHostName property value. Represents the host name of the app to which the calling service wants to post the notification, for the given user. If targeting web endpoints (see targetPolicy.platformTypes), ensure that targetHostName is the same as the name used when creating a subscription on the client side within the application JSON property.
func (m *Notification) GetTargetHostName()(*string) {
    return m.targetHostName
}
// GetTargetPolicy gets the targetPolicy property value. Target policy object handles notification delivery policy for endpoint types that should be targeted (Windows, iOS, Android and WebPush) for the given user.
func (m *Notification) GetTargetPolicy()(TargetPolicyEndpointsable) {
    return m.targetPolicy
}
// Serialize serializes information the current object
func (m *Notification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("displayTimeToLive", m.GetDisplayTimeToLive())
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
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        cast := (*m.GetPriority()).String()
        err = writer.WriteStringValue("priority", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetHostName", m.GetTargetHostName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetPolicy", m.GetTargetPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayTimeToLive sets the displayTimeToLive property value. Sets how long (in seconds) this notification content will stay in each platform's notification viewer. For example, when the notification is delivered to a Windows device, the value of this property is passed on to ToastNotification.ExpirationTime, which determines how long the toast notification will stay in the user's Windows Action Center.
func (m *Notification) SetDisplayTimeToLive(value *int32)() {
    m.displayTimeToLive = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Sets a UTC expiration date and time on a user notification using ISO 8601 format (for example, midnight UTC on Jan 1, 2019 would look like this: '2019-01-01T00:00:00Z'). When time is up, the notification is removed from the Microsoft Graph notification feed store completely and is no longer part of notification history. Max value is 30 days.
func (m *Notification) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetGroupName sets the groupName property value. The name of the group that this notification belongs to. It is set by the developer for the purpose of grouping notifications together.
func (m *Notification) SetGroupName(value *string)() {
    m.groupName = value
}
// SetPayload sets the payload property value. The payload property
func (m *Notification) SetPayload(value PayloadTypesable)() {
    m.payload = value
}
// SetPriority sets the priority property value. Indicates the priority of a raw user notification. Visual notifications are sent with high priority by default. Valid values are None, High and Low.
func (m *Notification) SetPriority(value *Priority)() {
    m.priority = value
}
// SetTargetHostName sets the targetHostName property value. Represents the host name of the app to which the calling service wants to post the notification, for the given user. If targeting web endpoints (see targetPolicy.platformTypes), ensure that targetHostName is the same as the name used when creating a subscription on the client side within the application JSON property.
func (m *Notification) SetTargetHostName(value *string)() {
    m.targetHostName = value
}
// SetTargetPolicy sets the targetPolicy property value. Target policy object handles notification delivery policy for endpoint types that should be targeted (Windows, iOS, Android and WebPush) for the given user.
func (m *Notification) SetTargetPolicy(value TargetPolicyEndpointsable)() {
    m.targetPolicy = value
}
