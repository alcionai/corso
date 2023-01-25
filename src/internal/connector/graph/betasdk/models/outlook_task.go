package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTask 
type OutlookTask struct {
    OutlookItem
    // The name of the person who has been assigned the task in Outlook. Read-only.
    assignedTo *string
    // The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
    attachments []Attachmentable
    // The task body that typically contains information about the task. Note that only HTML type is supported.
    body ItemBodyable
    // The date in the specified time zone that the task was finished.
    completedDateTime DateTimeTimeZoneable
    // The date in the specified time zone that the task is to be finished.
    dueDateTime DateTimeTimeZoneable
    // Set to true if the task has attachments.
    hasAttachments *bool
    // The importance property
    importance *Importance
    // The isReminderOn property
    isReminderOn *bool
    // The collection of multi-value extended properties defined for the task. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The owner property
    owner *string
    // The parentFolderId property
    parentFolderId *string
    // The recurrence property
    recurrence PatternedRecurrenceable
    // The reminderDateTime property
    reminderDateTime DateTimeTimeZoneable
    // The sensitivity property
    sensitivity *Sensitivity
    // The collection of single-value extended properties defined for the task. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The startDateTime property
    startDateTime DateTimeTimeZoneable
    // The status property
    status *TaskStatus
    // The subject property
    subject *string
}
// NewOutlookTask instantiates a new OutlookTask and sets the default values.
func NewOutlookTask()(*OutlookTask) {
    m := &OutlookTask{
        OutlookItem: *NewOutlookItem(),
    }
    odataTypeValue := "#microsoft.graph.outlookTask";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOutlookTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookTask(), nil
}
// GetAssignedTo gets the assignedTo property value. The name of the person who has been assigned the task in Outlook. Read-only.
func (m *OutlookTask) GetAssignedTo()(*string) {
    return m.assignedTo
}
// GetAttachments gets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *OutlookTask) GetAttachments()([]Attachmentable) {
    return m.attachments
}
// GetBody gets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
func (m *OutlookTask) GetBody()(ItemBodyable) {
    return m.body
}
// GetCompletedDateTime gets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *OutlookTask) GetCompletedDateTime()(DateTimeTimeZoneable) {
    return m.completedDateTime
}
// GetDueDateTime gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *OutlookTask) GetDueDateTime()(DateTimeTimeZoneable) {
    return m.dueDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
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
    res["attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Attachmentable, len(val))
            for i, v := range val {
                res[i] = v.(Attachmentable)
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ItemBodyable))
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["hasAttachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["isReminderOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReminderOn(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["parentFolderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["reminderDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["sensitivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val.(*Sensitivity))
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TaskStatus))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetHasAttachments gets the hasAttachments property value. Set to true if the task has attachments.
func (m *OutlookTask) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetImportance gets the importance property value. The importance property
func (m *OutlookTask) GetImportance()(*Importance) {
    return m.importance
}
// GetIsReminderOn gets the isReminderOn property value. The isReminderOn property
func (m *OutlookTask) GetIsReminderOn()(*bool) {
    return m.isReminderOn
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetOwner gets the owner property value. The owner property
func (m *OutlookTask) GetOwner()(*string) {
    return m.owner
}
// GetParentFolderId gets the parentFolderId property value. The parentFolderId property
func (m *OutlookTask) GetParentFolderId()(*string) {
    return m.parentFolderId
}
// GetRecurrence gets the recurrence property value. The recurrence property
func (m *OutlookTask) GetRecurrence()(PatternedRecurrenceable) {
    return m.recurrence
}
// GetReminderDateTime gets the reminderDateTime property value. The reminderDateTime property
func (m *OutlookTask) GetReminderDateTime()(DateTimeTimeZoneable) {
    return m.reminderDateTime
}
// GetSensitivity gets the sensitivity property value. The sensitivity property
func (m *OutlookTask) GetSensitivity()(*Sensitivity) {
    return m.sensitivity
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *OutlookTask) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status property
func (m *OutlookTask) GetStatus()(*TaskStatus) {
    return m.status
}
// GetSubject gets the subject property value. The subject property
func (m *OutlookTask) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *OutlookTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetAttachments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReminderOn", m.GetIsReminderOn())
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
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reminderDateTime", m.GetReminderDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := (*m.GetSensitivity()).String()
        err = writer.WriteStringValue("sensitivity", &cast)
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
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
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
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. The name of the person who has been assigned the task in Outlook. Read-only.
func (m *OutlookTask) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// SetAttachments sets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *OutlookTask) SetAttachments(value []Attachmentable)() {
    m.attachments = value
}
// SetBody sets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
func (m *OutlookTask) SetBody(value ItemBodyable)() {
    m.body = value
}
// SetCompletedDateTime sets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *OutlookTask) SetCompletedDateTime(value DateTimeTimeZoneable)() {
    m.completedDateTime = value
}
// SetDueDateTime sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *OutlookTask) SetDueDateTime(value DateTimeTimeZoneable)() {
    m.dueDateTime = value
}
// SetHasAttachments sets the hasAttachments property value. Set to true if the task has attachments.
func (m *OutlookTask) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetImportance sets the importance property value. The importance property
func (m *OutlookTask) SetImportance(value *Importance)() {
    m.importance = value
}
// SetIsReminderOn sets the isReminderOn property value. The isReminderOn property
func (m *OutlookTask) SetIsReminderOn(value *bool)() {
    m.isReminderOn = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetOwner sets the owner property value. The owner property
func (m *OutlookTask) SetOwner(value *string)() {
    m.owner = value
}
// SetParentFolderId sets the parentFolderId property value. The parentFolderId property
func (m *OutlookTask) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetRecurrence sets the recurrence property value. The recurrence property
func (m *OutlookTask) SetRecurrence(value PatternedRecurrenceable)() {
    m.recurrence = value
}
// SetReminderDateTime sets the reminderDateTime property value. The reminderDateTime property
func (m *OutlookTask) SetReminderDateTime(value DateTimeTimeZoneable)() {
    m.reminderDateTime = value
}
// SetSensitivity sets the sensitivity property value. The sensitivity property
func (m *OutlookTask) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *OutlookTask) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *OutlookTask) SetStatus(value *TaskStatus)() {
    m.status = value
}
// SetSubject sets the subject property value. The subject property
func (m *OutlookTask) SetSubject(value *string)() {
    m.subject = value
}
