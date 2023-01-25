package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTaskFolder provides operations to manage the collection of site entities.
type OutlookTaskFolder struct {
    Entity
    // The version of the task folder.
    changeKey *string
    // True if the folder is the default task folder.
    isDefaultFolder *bool
    // The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The name of the task folder.
    name *string
    // The unique GUID identifier for the task folder's parent group.
    parentGroupKey *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The tasks in this task folder. Read-only. Nullable.
    tasks []OutlookTaskable
}
// NewOutlookTaskFolder instantiates a new outlookTaskFolder and sets the default values.
func NewOutlookTaskFolder()(*OutlookTaskFolder) {
    m := &OutlookTaskFolder{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOutlookTaskFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookTaskFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookTaskFolder(), nil
}
// GetChangeKey gets the changeKey property value. The version of the task folder.
func (m *OutlookTaskFolder) GetChangeKey()(*string) {
    return m.changeKey
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTaskFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["changeKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["isDefaultFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultFolder(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(MultiValueLegacyExtendedPropertyable)
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parentGroupKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentGroupKey(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SingleValueLegacyExtendedPropertyable)
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskable, len(val))
            for i, v := range val {
                res[i] = v.(OutlookTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetIsDefaultFolder gets the isDefaultFolder property value. True if the folder is the default task folder.
func (m *OutlookTaskFolder) GetIsDefaultFolder()(*bool) {
    return m.isDefaultFolder
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetName gets the name property value. The name of the task folder.
func (m *OutlookTaskFolder) GetName()(*string) {
    return m.name
}
// GetParentGroupKey gets the parentGroupKey property value. The unique GUID identifier for the task folder's parent group.
func (m *OutlookTaskFolder) GetParentGroupKey()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.parentGroupKey
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetTasks gets the tasks property value. The tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetTasks()([]OutlookTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *OutlookTaskFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultFolder", m.GetIsDefaultFolder())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("parentGroupKey", m.GetParentGroupKey())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChangeKey sets the changeKey property value. The version of the task folder.
func (m *OutlookTaskFolder) SetChangeKey(value *string)() {
    m.changeKey = value
}
// SetIsDefaultFolder sets the isDefaultFolder property value. True if the folder is the default task folder.
func (m *OutlookTaskFolder) SetIsDefaultFolder(value *bool)() {
    m.isDefaultFolder = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetName sets the name property value. The name of the task folder.
func (m *OutlookTaskFolder) SetName(value *string)() {
    m.name = value
}
// SetParentGroupKey sets the parentGroupKey property value. The unique GUID identifier for the task folder's parent group.
func (m *OutlookTaskFolder) SetParentGroupKey(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.parentGroupKey = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetTasks sets the tasks property value. The tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetTasks(value []OutlookTaskable)() {
    m.tasks = value
}
