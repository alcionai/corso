package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementTask a device app management task.
type DeviceAppManagementTask struct {
    Entity
    // The name or email of the admin this task is assigned to.
    assignedTo *string
    // Device app management task category.
    category *DeviceAppManagementTaskCategory
    // The created date.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The email address of the creator.
    creator *string
    // Notes from the creator.
    creatorNotes *string
    // The description.
    description *string
    // The name.
    displayName *string
    // The due date.
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device app management task priority.
    priority *DeviceAppManagementTaskPriority
    // Device app management task status.
    status *DeviceAppManagementTaskStatus
}
// NewDeviceAppManagementTask instantiates a new deviceAppManagementTask and sets the default values.
func NewDeviceAppManagementTask()(*DeviceAppManagementTask) {
    m := &DeviceAppManagementTask{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceAppManagementTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.appVulnerabilityTask":
                        return NewAppVulnerabilityTask(), nil
                    case "#microsoft.graph.securityConfigurationTask":
                        return NewSecurityConfigurationTask(), nil
                    case "#microsoft.graph.unmanagedDeviceDiscoveryTask":
                        return NewUnmanagedDeviceDiscoveryTask(), nil
                }
            }
        }
    }
    return NewDeviceAppManagementTask(), nil
}
// GetAssignedTo gets the assignedTo property value. The name or email of the admin this task is assigned to.
func (m *DeviceAppManagementTask) GetAssignedTo()(*string) {
    return m.assignedTo
}
// GetCategory gets the category property value. Device app management task category.
func (m *DeviceAppManagementTask) GetCategory()(*DeviceAppManagementTaskCategory) {
    return m.category
}
// GetCreatedDateTime gets the createdDateTime property value. The created date.
func (m *DeviceAppManagementTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreator gets the creator property value. The email address of the creator.
func (m *DeviceAppManagementTask) GetCreator()(*string) {
    return m.creator
}
// GetCreatorNotes gets the creatorNotes property value. Notes from the creator.
func (m *DeviceAppManagementTask) GetCreatorNotes()(*string) {
    return m.creatorNotes
}
// GetDescription gets the description property value. The description.
func (m *DeviceAppManagementTask) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name.
func (m *DeviceAppManagementTask) GetDisplayName()(*string) {
    return m.displayName
}
// GetDueDateTime gets the dueDateTime property value. The due date.
func (m *DeviceAppManagementTask) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dueDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*DeviceAppManagementTaskCategory))
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
    res["creator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val)
        }
        return nil
    }
    res["creatorNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatorNotes(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*DeviceAppManagementTaskPriority))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceAppManagementTaskStatus))
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. Device app management task priority.
func (m *DeviceAppManagementTask) GetPriority()(*DeviceAppManagementTaskPriority) {
    return m.priority
}
// GetStatus gets the status property value. Device app management task status.
func (m *DeviceAppManagementTask) GetStatus()(*DeviceAppManagementTaskStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DeviceAppManagementTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creator", m.GetCreator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creatorNotes", m.GetCreatorNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    {
        err = writer.WriteTimeValue("dueDateTime", m.GetDueDateTime())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. The name or email of the admin this task is assigned to.
func (m *DeviceAppManagementTask) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// SetCategory sets the category property value. Device app management task category.
func (m *DeviceAppManagementTask) SetCategory(value *DeviceAppManagementTaskCategory)() {
    m.category = value
}
// SetCreatedDateTime sets the createdDateTime property value. The created date.
func (m *DeviceAppManagementTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreator sets the creator property value. The email address of the creator.
func (m *DeviceAppManagementTask) SetCreator(value *string)() {
    m.creator = value
}
// SetCreatorNotes sets the creatorNotes property value. Notes from the creator.
func (m *DeviceAppManagementTask) SetCreatorNotes(value *string)() {
    m.creatorNotes = value
}
// SetDescription sets the description property value. The description.
func (m *DeviceAppManagementTask) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name.
func (m *DeviceAppManagementTask) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDueDateTime sets the dueDateTime property value. The due date.
func (m *DeviceAppManagementTask) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDateTime = value
}
// SetPriority sets the priority property value. Device app management task priority.
func (m *DeviceAppManagementTask) SetPriority(value *DeviceAppManagementTaskPriority)() {
    m.priority = value
}
// SetStatus sets the status property value. Device app management task status.
func (m *DeviceAppManagementTask) SetStatus(value *DeviceAppManagementTaskStatus)() {
    m.status = value
}
