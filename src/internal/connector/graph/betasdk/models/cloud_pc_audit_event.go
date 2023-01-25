package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcAuditEvent 
type CloudPcAuditEvent struct {
    Entity
    // Friendly name of the activity. Optional.
    activity *string
    // The date time in UTC when the activity was performed. Read-only.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The activityOperationType property
    activityOperationType *CloudPcAuditActivityOperationType
    // The activityResult property
    activityResult *CloudPcAuditActivityResult
    // The type of activity that was performed. Read-only.
    activityType *string
    // The actor property
    actor CloudPcAuditActorable
    // The category property
    category *CloudPcAuditCategory
    // Component name. Read-only.
    componentName *string
    // The client request identifier, used to correlate activity within the system. Read-only.
    correlationId *string
    // Event display name. Read-only.
    displayName *string
    // List of cloudPcAuditResource objects. Read-only.
    resources []CloudPcAuditResourceable
}
// NewCloudPcAuditEvent instantiates a new CloudPcAuditEvent and sets the default values.
func NewCloudPcAuditEvent()(*CloudPcAuditEvent) {
    m := &CloudPcAuditEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcAuditEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcAuditEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcAuditEvent(), nil
}
// GetActivity gets the activity property value. Friendly name of the activity. Optional.
func (m *CloudPcAuditEvent) GetActivity()(*string) {
    return m.activity
}
// GetActivityDateTime gets the activityDateTime property value. The date time in UTC when the activity was performed. Read-only.
func (m *CloudPcAuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activityDateTime
}
// GetActivityOperationType gets the activityOperationType property value. The activityOperationType property
func (m *CloudPcAuditEvent) GetActivityOperationType()(*CloudPcAuditActivityOperationType) {
    return m.activityOperationType
}
// GetActivityResult gets the activityResult property value. The activityResult property
func (m *CloudPcAuditEvent) GetActivityResult()(*CloudPcAuditActivityResult) {
    return m.activityResult
}
// GetActivityType gets the activityType property value. The type of activity that was performed. Read-only.
func (m *CloudPcAuditEvent) GetActivityType()(*string) {
    return m.activityType
}
// GetActor gets the actor property value. The actor property
func (m *CloudPcAuditEvent) GetActor()(CloudPcAuditActorable) {
    return m.actor
}
// GetCategory gets the category property value. The category property
func (m *CloudPcAuditEvent) GetCategory()(*CloudPcAuditCategory) {
    return m.category
}
// GetComponentName gets the componentName property value. Component name. Read-only.
func (m *CloudPcAuditEvent) GetComponentName()(*string) {
    return m.componentName
}
// GetCorrelationId gets the correlationId property value. The client request identifier, used to correlate activity within the system. Read-only.
func (m *CloudPcAuditEvent) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetDisplayName gets the displayName property value. Event display name. Read-only.
func (m *CloudPcAuditEvent) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcAuditEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["activityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["activityOperationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActivityOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityOperationType(val.(*CloudPcAuditActivityOperationType))
        }
        return nil
    }
    res["activityResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActivityResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityResult(val.(*CloudPcAuditActivityResult))
        }
        return nil
    }
    res["activityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcAuditActorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(CloudPcAuditActorable))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*CloudPcAuditCategory))
        }
        return nil
    }
    res["componentName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComponentName(val)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcAuditResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcAuditResourceable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcAuditResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetResources gets the resources property value. List of cloudPcAuditResource objects. Read-only.
func (m *CloudPcAuditEvent) GetResources()([]CloudPcAuditResourceable) {
    return m.resources
}
// Serialize serializes information the current object
func (m *CloudPcAuditEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetActivityOperationType() != nil {
        cast := (*m.GetActivityOperationType()).String()
        err = writer.WriteStringValue("activityOperationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActivityResult() != nil {
        cast := (*m.GetActivityResult()).String()
        err = writer.WriteStringValue("activityResult", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("componentName", m.GetComponentName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. Friendly name of the activity. Optional.
func (m *CloudPcAuditEvent) SetActivity(value *string)() {
    m.activity = value
}
// SetActivityDateTime sets the activityDateTime property value. The date time in UTC when the activity was performed. Read-only.
func (m *CloudPcAuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// SetActivityOperationType sets the activityOperationType property value. The activityOperationType property
func (m *CloudPcAuditEvent) SetActivityOperationType(value *CloudPcAuditActivityOperationType)() {
    m.activityOperationType = value
}
// SetActivityResult sets the activityResult property value. The activityResult property
func (m *CloudPcAuditEvent) SetActivityResult(value *CloudPcAuditActivityResult)() {
    m.activityResult = value
}
// SetActivityType sets the activityType property value. The type of activity that was performed. Read-only.
func (m *CloudPcAuditEvent) SetActivityType(value *string)() {
    m.activityType = value
}
// SetActor sets the actor property value. The actor property
func (m *CloudPcAuditEvent) SetActor(value CloudPcAuditActorable)() {
    m.actor = value
}
// SetCategory sets the category property value. The category property
func (m *CloudPcAuditEvent) SetCategory(value *CloudPcAuditCategory)() {
    m.category = value
}
// SetComponentName sets the componentName property value. Component name. Read-only.
func (m *CloudPcAuditEvent) SetComponentName(value *string)() {
    m.componentName = value
}
// SetCorrelationId sets the correlationId property value. The client request identifier, used to correlate activity within the system. Read-only.
func (m *CloudPcAuditEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetDisplayName sets the displayName property value. Event display name. Read-only.
func (m *CloudPcAuditEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetResources sets the resources property value. List of cloudPcAuditResource objects. Read-only.
func (m *CloudPcAuditEvent) SetResources(value []CloudPcAuditResourceable)() {
    m.resources = value
}
