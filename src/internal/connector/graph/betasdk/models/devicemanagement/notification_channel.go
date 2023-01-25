package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationChannel 
type NotificationChannel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type of the notification channel. The possible values are: portal, email, phoneCall, sms, unknownFutureValue.
    notificationChannelType *NotificationChannelType
    // Information about the notification receivers, such as locale and contact information. For example, en-us for locale and serena.davis@contoso.com for contact information.
    notificationReceivers []NotificationReceiverable
    // The OdataType property
    odataType *string
    // The contact information about the notification receivers, such as email addresses. For portal notifications, receivers can be left blank. For email notifications, receivers consists of email addresses such as serena.davis@contoso.com.
    receivers []string
}
// NewNotificationChannel instantiates a new notificationChannel and sets the default values.
func NewNotificationChannel()(*NotificationChannel) {
    m := &NotificationChannel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateNotificationChannelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationChannelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationChannel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NotificationChannel) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationChannel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notificationChannelType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationChannelType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationChannelType(val.(*NotificationChannelType))
        }
        return nil
    }
    res["notificationReceivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationReceiverFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationReceiverable, len(val))
            for i, v := range val {
                res[i] = v.(NotificationReceiverable)
            }
            m.SetNotificationReceivers(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["receivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetReceivers(res)
        }
        return nil
    }
    return res
}
// GetNotificationChannelType gets the notificationChannelType property value. The type of the notification channel. The possible values are: portal, email, phoneCall, sms, unknownFutureValue.
func (m *NotificationChannel) GetNotificationChannelType()(*NotificationChannelType) {
    return m.notificationChannelType
}
// GetNotificationReceivers gets the notificationReceivers property value. Information about the notification receivers, such as locale and contact information. For example, en-us for locale and serena.davis@contoso.com for contact information.
func (m *NotificationChannel) GetNotificationReceivers()([]NotificationReceiverable) {
    return m.notificationReceivers
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NotificationChannel) GetOdataType()(*string) {
    return m.odataType
}
// GetReceivers gets the receivers property value. The contact information about the notification receivers, such as email addresses. For portal notifications, receivers can be left blank. For email notifications, receivers consists of email addresses such as serena.davis@contoso.com.
func (m *NotificationChannel) GetReceivers()([]string) {
    return m.receivers
}
// Serialize serializes information the current object
func (m *NotificationChannel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNotificationChannelType() != nil {
        cast := (*m.GetNotificationChannelType()).String()
        err := writer.WriteStringValue("notificationChannelType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationReceivers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationReceivers()))
        for i, v := range m.GetNotificationReceivers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("notificationReceivers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetReceivers() != nil {
        err := writer.WriteCollectionOfStringValues("receivers", m.GetReceivers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NotificationChannel) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNotificationChannelType sets the notificationChannelType property value. The type of the notification channel. The possible values are: portal, email, phoneCall, sms, unknownFutureValue.
func (m *NotificationChannel) SetNotificationChannelType(value *NotificationChannelType)() {
    m.notificationChannelType = value
}
// SetNotificationReceivers sets the notificationReceivers property value. Information about the notification receivers, such as locale and contact information. For example, en-us for locale and serena.davis@contoso.com for contact information.
func (m *NotificationChannel) SetNotificationReceivers(value []NotificationReceiverable)() {
    m.notificationReceivers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NotificationChannel) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReceivers sets the receivers property value. The contact information about the notification receivers, such as email addresses. For portal notifications, receivers can be left blank. For email notifications, receivers consists of email addresses such as serena.davis@contoso.com.
func (m *NotificationChannel) SetReceivers(value []string)() {
    m.receivers = value
}
