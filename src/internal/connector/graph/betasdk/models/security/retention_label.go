package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// RetentionLabel provides operations to manage the collection of site entities.
type RetentionLabel struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Specifies the action to take on a document with this label applied during the retention period. The possible values are: none, delete, startDispositionReview, unknownFutureValue.
    actionAfterRetentionPeriod *ActionAfterRetentionPeriod
    // Specifies how the behavior of a document with this label should be during the retention period. The possible values are: doNotRetain, retain, retainAsRecord, retainAsRegulatoryRecord, unknownFutureValue.
    behaviorDuringRetentionPeriod *BehaviorDuringRetentionPeriod
    // Represents the user who created the retentionLabel.
    createdBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable
    // Represents the date and time in which the retentionLabel is created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the locked or unlocked state of a record label when it is created.The possible values are: startLocked, startUnlocked, unknownFutureValue.
    defaultRecordBehavior *DefaultRecordBehavior
    // Provides label information for the admin. Optional.
    descriptionForAdmins *string
    // Provides the label information for the user. Optional.
    descriptionForUsers *string
    // Unique string that defines a label name.
    displayName *string
    // Review stages during which reviewers are notified to determine whether a document must be deleted or retained.
    dispositionReviewStages []DispositionReviewStageable
    // Specifies whether the label is currently being used.
    isInUse *bool
    // Specifies the replacement label to be applied automatically after the retention period of the current label ends.
    labelToBeApplied *string
    // The user who last modified the retentionLabel.
    lastModifiedBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable
    // The latest date time when the retentionLabel was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the number of days to retain the content.
    retentionDuration RetentionDurationable
    // The retentionEventType property
    retentionEventType RetentionEventTypeable
    // Specifies whether the retention duration is calculated from the content creation date, labeled date, or last modification date. The possible values are: dateLabeled, dateCreated, dateModified, dateOfEvent, unknownFutureValue.
    retentionTrigger *RetentionTrigger
}
// NewRetentionLabel instantiates a new retentionLabel and sets the default values.
func NewRetentionLabel()(*RetentionLabel) {
    m := &RetentionLabel{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateRetentionLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionLabel(), nil
}
// GetActionAfterRetentionPeriod gets the actionAfterRetentionPeriod property value. Specifies the action to take on a document with this label applied during the retention period. The possible values are: none, delete, startDispositionReview, unknownFutureValue.
func (m *RetentionLabel) GetActionAfterRetentionPeriod()(*ActionAfterRetentionPeriod) {
    return m.actionAfterRetentionPeriod
}
// GetBehaviorDuringRetentionPeriod gets the behaviorDuringRetentionPeriod property value. Specifies how the behavior of a document with this label should be during the retention period. The possible values are: doNotRetain, retain, retainAsRecord, retainAsRegulatoryRecord, unknownFutureValue.
func (m *RetentionLabel) GetBehaviorDuringRetentionPeriod()(*BehaviorDuringRetentionPeriod) {
    return m.behaviorDuringRetentionPeriod
}
// GetCreatedBy gets the createdBy property value. Represents the user who created the retentionLabel.
func (m *RetentionLabel) GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Represents the date and time in which the retentionLabel is created.
func (m *RetentionLabel) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefaultRecordBehavior gets the defaultRecordBehavior property value. Specifies the locked or unlocked state of a record label when it is created.The possible values are: startLocked, startUnlocked, unknownFutureValue.
func (m *RetentionLabel) GetDefaultRecordBehavior()(*DefaultRecordBehavior) {
    return m.defaultRecordBehavior
}
// GetDescriptionForAdmins gets the descriptionForAdmins property value. Provides label information for the admin. Optional.
func (m *RetentionLabel) GetDescriptionForAdmins()(*string) {
    return m.descriptionForAdmins
}
// GetDescriptionForUsers gets the descriptionForUsers property value. Provides the label information for the user. Optional.
func (m *RetentionLabel) GetDescriptionForUsers()(*string) {
    return m.descriptionForUsers
}
// GetDisplayName gets the displayName property value. Unique string that defines a label name.
func (m *RetentionLabel) GetDisplayName()(*string) {
    return m.displayName
}
// GetDispositionReviewStages gets the dispositionReviewStages property value. Review stages during which reviewers are notified to determine whether a document must be deleted or retained.
func (m *RetentionLabel) GetDispositionReviewStages()([]DispositionReviewStageable) {
    return m.dispositionReviewStages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionAfterRetentionPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionAfterRetentionPeriod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionAfterRetentionPeriod(val.(*ActionAfterRetentionPeriod))
        }
        return nil
    }
    res["behaviorDuringRetentionPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBehaviorDuringRetentionPeriod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBehaviorDuringRetentionPeriod(val.(*BehaviorDuringRetentionPeriod))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable))
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
    res["defaultRecordBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefaultRecordBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRecordBehavior(val.(*DefaultRecordBehavior))
        }
        return nil
    }
    res["descriptionForAdmins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescriptionForAdmins(val)
        }
        return nil
    }
    res["descriptionForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescriptionForUsers(val)
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
    res["dispositionReviewStages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDispositionReviewStageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DispositionReviewStageable, len(val))
            for i, v := range val {
                res[i] = v.(DispositionReviewStageable)
            }
            m.SetDispositionReviewStages(res)
        }
        return nil
    }
    res["isInUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInUse(val)
        }
        return nil
    }
    res["labelToBeApplied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelToBeApplied(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["retentionDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRetentionDurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionDuration(val.(RetentionDurationable))
        }
        return nil
    }
    res["retentionEventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRetentionEventTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionEventType(val.(RetentionEventTypeable))
        }
        return nil
    }
    res["retentionTrigger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRetentionTrigger)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionTrigger(val.(*RetentionTrigger))
        }
        return nil
    }
    return res
}
// GetIsInUse gets the isInUse property value. Specifies whether the label is currently being used.
func (m *RetentionLabel) GetIsInUse()(*bool) {
    return m.isInUse
}
// GetLabelToBeApplied gets the labelToBeApplied property value. Specifies the replacement label to be applied automatically after the retention period of the current label ends.
func (m *RetentionLabel) GetLabelToBeApplied()(*string) {
    return m.labelToBeApplied
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the retentionLabel.
func (m *RetentionLabel) GetLastModifiedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date time when the retentionLabel was modified.
func (m *RetentionLabel) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRetentionDuration gets the retentionDuration property value. Specifies the number of days to retain the content.
func (m *RetentionLabel) GetRetentionDuration()(RetentionDurationable) {
    return m.retentionDuration
}
// GetRetentionEventType gets the retentionEventType property value. The retentionEventType property
func (m *RetentionLabel) GetRetentionEventType()(RetentionEventTypeable) {
    return m.retentionEventType
}
// GetRetentionTrigger gets the retentionTrigger property value. Specifies whether the retention duration is calculated from the content creation date, labeled date, or last modification date. The possible values are: dateLabeled, dateCreated, dateModified, dateOfEvent, unknownFutureValue.
func (m *RetentionLabel) GetRetentionTrigger()(*RetentionTrigger) {
    return m.retentionTrigger
}
// Serialize serializes information the current object
func (m *RetentionLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionAfterRetentionPeriod() != nil {
        cast := (*m.GetActionAfterRetentionPeriod()).String()
        err = writer.WriteStringValue("actionAfterRetentionPeriod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBehaviorDuringRetentionPeriod() != nil {
        cast := (*m.GetBehaviorDuringRetentionPeriod()).String()
        err = writer.WriteStringValue("behaviorDuringRetentionPeriod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
    if m.GetDefaultRecordBehavior() != nil {
        cast := (*m.GetDefaultRecordBehavior()).String()
        err = writer.WriteStringValue("defaultRecordBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForAdmins", m.GetDescriptionForAdmins())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForUsers", m.GetDescriptionForUsers())
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
    if m.GetDispositionReviewStages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDispositionReviewStages()))
        for i, v := range m.GetDispositionReviewStages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dispositionReviewStages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInUse", m.GetIsInUse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("labelToBeApplied", m.GetLabelToBeApplied())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("retentionDuration", m.GetRetentionDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("retentionEventType", m.GetRetentionEventType())
        if err != nil {
            return err
        }
    }
    if m.GetRetentionTrigger() != nil {
        cast := (*m.GetRetentionTrigger()).String()
        err = writer.WriteStringValue("retentionTrigger", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionAfterRetentionPeriod sets the actionAfterRetentionPeriod property value. Specifies the action to take on a document with this label applied during the retention period. The possible values are: none, delete, startDispositionReview, unknownFutureValue.
func (m *RetentionLabel) SetActionAfterRetentionPeriod(value *ActionAfterRetentionPeriod)() {
    m.actionAfterRetentionPeriod = value
}
// SetBehaviorDuringRetentionPeriod sets the behaviorDuringRetentionPeriod property value. Specifies how the behavior of a document with this label should be during the retention period. The possible values are: doNotRetain, retain, retainAsRecord, retainAsRegulatoryRecord, unknownFutureValue.
func (m *RetentionLabel) SetBehaviorDuringRetentionPeriod(value *BehaviorDuringRetentionPeriod)() {
    m.behaviorDuringRetentionPeriod = value
}
// SetCreatedBy sets the createdBy property value. Represents the user who created the retentionLabel.
func (m *RetentionLabel) SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Represents the date and time in which the retentionLabel is created.
func (m *RetentionLabel) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefaultRecordBehavior sets the defaultRecordBehavior property value. Specifies the locked or unlocked state of a record label when it is created.The possible values are: startLocked, startUnlocked, unknownFutureValue.
func (m *RetentionLabel) SetDefaultRecordBehavior(value *DefaultRecordBehavior)() {
    m.defaultRecordBehavior = value
}
// SetDescriptionForAdmins sets the descriptionForAdmins property value. Provides label information for the admin. Optional.
func (m *RetentionLabel) SetDescriptionForAdmins(value *string)() {
    m.descriptionForAdmins = value
}
// SetDescriptionForUsers sets the descriptionForUsers property value. Provides the label information for the user. Optional.
func (m *RetentionLabel) SetDescriptionForUsers(value *string)() {
    m.descriptionForUsers = value
}
// SetDisplayName sets the displayName property value. Unique string that defines a label name.
func (m *RetentionLabel) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDispositionReviewStages sets the dispositionReviewStages property value. Review stages during which reviewers are notified to determine whether a document must be deleted or retained.
func (m *RetentionLabel) SetDispositionReviewStages(value []DispositionReviewStageable)() {
    m.dispositionReviewStages = value
}
// SetIsInUse sets the isInUse property value. Specifies whether the label is currently being used.
func (m *RetentionLabel) SetIsInUse(value *bool)() {
    m.isInUse = value
}
// SetLabelToBeApplied sets the labelToBeApplied property value. Specifies the replacement label to be applied automatically after the retention period of the current label ends.
func (m *RetentionLabel) SetLabelToBeApplied(value *string)() {
    m.labelToBeApplied = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the retentionLabel.
func (m *RetentionLabel) SetLastModifiedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date time when the retentionLabel was modified.
func (m *RetentionLabel) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRetentionDuration sets the retentionDuration property value. Specifies the number of days to retain the content.
func (m *RetentionLabel) SetRetentionDuration(value RetentionDurationable)() {
    m.retentionDuration = value
}
// SetRetentionEventType sets the retentionEventType property value. The retentionEventType property
func (m *RetentionLabel) SetRetentionEventType(value RetentionEventTypeable)() {
    m.retentionEventType = value
}
// SetRetentionTrigger sets the retentionTrigger property value. Specifies whether the retention duration is calculated from the content creation date, labeled date, or last modification date. The possible values are: dateLabeled, dateCreated, dateModified, dateOfEvent, unknownFutureValue.
func (m *RetentionLabel) SetRetentionTrigger(value *RetentionTrigger)() {
    m.retentionTrigger = value
}
