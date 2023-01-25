package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewSettings 
type AccessReviewSettings struct {
    // Indicates whether showing recommendations to reviewers is enabled.
    accessRecommendationsEnabled *bool
    // The number of days of user activities to show to reviewers.
    activityDurationInDays *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
    autoApplyReviewResultsEnabled *bool
    // Indicates whether a decision should be set if the reviewer did not supply one. For use when auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
    autoReviewEnabled *bool
    // Detailed settings for how the feature should set the review decision. For use when auto-apply is enabled.
    autoReviewSettings AutoReviewSettingsable
    // Indicates whether reviewers are required to provide a justification when reviewing access.
    justificationRequiredOnApproval *bool
    // Indicates whether sending mails to reviewers and the review creator is enabled.
    mailNotificationsEnabled *bool
    // The OdataType property
    odataType *string
    // Detailed settings for recurrence.
    recurrenceSettings AccessReviewRecurrenceSettingsable
    // Indicates whether sending reminder emails to reviewers is enabled.
    remindersEnabled *bool
}
// NewAccessReviewSettings instantiates a new accessReviewSettings and sets the default values.
func NewAccessReviewSettings()(*AccessReviewSettings) {
    m := &AccessReviewSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessReviewSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.businessFlowSettings":
                        return NewBusinessFlowSettings(), nil
                }
            }
        }
    }
    return NewAccessReviewSettings(), nil
}
// GetAccessRecommendationsEnabled gets the accessRecommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled.
func (m *AccessReviewSettings) GetAccessRecommendationsEnabled()(*bool) {
    return m.accessRecommendationsEnabled
}
// GetActivityDurationInDays gets the activityDurationInDays property value. The number of days of user activities to show to reviewers.
func (m *AccessReviewSettings) GetActivityDurationInDays()(*int32) {
    return m.activityDurationInDays
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAutoApplyReviewResultsEnabled gets the autoApplyReviewResultsEnabled property value. Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
func (m *AccessReviewSettings) GetAutoApplyReviewResultsEnabled()(*bool) {
    return m.autoApplyReviewResultsEnabled
}
// GetAutoReviewEnabled gets the autoReviewEnabled property value. Indicates whether a decision should be set if the reviewer did not supply one. For use when auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
func (m *AccessReviewSettings) GetAutoReviewEnabled()(*bool) {
    return m.autoReviewEnabled
}
// GetAutoReviewSettings gets the autoReviewSettings property value. Detailed settings for how the feature should set the review decision. For use when auto-apply is enabled.
func (m *AccessReviewSettings) GetAutoReviewSettings()(AutoReviewSettingsable) {
    return m.autoReviewSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessRecommendationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessRecommendationsEnabled(val)
        }
        return nil
    }
    res["activityDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDurationInDays(val)
        }
        return nil
    }
    res["autoApplyReviewResultsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoApplyReviewResultsEnabled(val)
        }
        return nil
    }
    res["autoReviewEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoReviewEnabled(val)
        }
        return nil
    }
    res["autoReviewSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutoReviewSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoReviewSettings(val.(AutoReviewSettingsable))
        }
        return nil
    }
    res["justificationRequiredOnApproval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationRequiredOnApproval(val)
        }
        return nil
    }
    res["mailNotificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNotificationsEnabled(val)
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
    res["recurrenceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewRecurrenceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceSettings(val.(AccessReviewRecurrenceSettingsable))
        }
        return nil
    }
    res["remindersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemindersEnabled(val)
        }
        return nil
    }
    return res
}
// GetJustificationRequiredOnApproval gets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide a justification when reviewing access.
func (m *AccessReviewSettings) GetJustificationRequiredOnApproval()(*bool) {
    return m.justificationRequiredOnApproval
}
// GetMailNotificationsEnabled gets the mailNotificationsEnabled property value. Indicates whether sending mails to reviewers and the review creator is enabled.
func (m *AccessReviewSettings) GetMailNotificationsEnabled()(*bool) {
    return m.mailNotificationsEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRecurrenceSettings gets the recurrenceSettings property value. Detailed settings for recurrence.
func (m *AccessReviewSettings) GetRecurrenceSettings()(AccessReviewRecurrenceSettingsable) {
    return m.recurrenceSettings
}
// GetRemindersEnabled gets the remindersEnabled property value. Indicates whether sending reminder emails to reviewers is enabled.
func (m *AccessReviewSettings) GetRemindersEnabled()(*bool) {
    return m.remindersEnabled
}
// Serialize serializes information the current object
func (m *AccessReviewSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("accessRecommendationsEnabled", m.GetAccessRecommendationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("activityDurationInDays", m.GetActivityDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("autoApplyReviewResultsEnabled", m.GetAutoApplyReviewResultsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("autoReviewEnabled", m.GetAutoReviewEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("autoReviewSettings", m.GetAutoReviewSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("justificationRequiredOnApproval", m.GetJustificationRequiredOnApproval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mailNotificationsEnabled", m.GetMailNotificationsEnabled())
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
    {
        err := writer.WriteObjectValue("recurrenceSettings", m.GetRecurrenceSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("remindersEnabled", m.GetRemindersEnabled())
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
// SetAccessRecommendationsEnabled sets the accessRecommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled.
func (m *AccessReviewSettings) SetAccessRecommendationsEnabled(value *bool)() {
    m.accessRecommendationsEnabled = value
}
// SetActivityDurationInDays sets the activityDurationInDays property value. The number of days of user activities to show to reviewers.
func (m *AccessReviewSettings) SetActivityDurationInDays(value *int32)() {
    m.activityDurationInDays = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAutoApplyReviewResultsEnabled sets the autoApplyReviewResultsEnabled property value. Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
func (m *AccessReviewSettings) SetAutoApplyReviewResultsEnabled(value *bool)() {
    m.autoApplyReviewResultsEnabled = value
}
// SetAutoReviewEnabled sets the autoReviewEnabled property value. Indicates whether a decision should be set if the reviewer did not supply one. For use when auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
func (m *AccessReviewSettings) SetAutoReviewEnabled(value *bool)() {
    m.autoReviewEnabled = value
}
// SetAutoReviewSettings sets the autoReviewSettings property value. Detailed settings for how the feature should set the review decision. For use when auto-apply is enabled.
func (m *AccessReviewSettings) SetAutoReviewSettings(value AutoReviewSettingsable)() {
    m.autoReviewSettings = value
}
// SetJustificationRequiredOnApproval sets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide a justification when reviewing access.
func (m *AccessReviewSettings) SetJustificationRequiredOnApproval(value *bool)() {
    m.justificationRequiredOnApproval = value
}
// SetMailNotificationsEnabled sets the mailNotificationsEnabled property value. Indicates whether sending mails to reviewers and the review creator is enabled.
func (m *AccessReviewSettings) SetMailNotificationsEnabled(value *bool)() {
    m.mailNotificationsEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecurrenceSettings sets the recurrenceSettings property value. Detailed settings for recurrence.
func (m *AccessReviewSettings) SetRecurrenceSettings(value AccessReviewRecurrenceSettingsable)() {
    m.recurrenceSettings = value
}
// SetRemindersEnabled sets the remindersEnabled property value. Indicates whether sending reminder emails to reviewers is enabled.
func (m *AccessReviewSettings) SetRemindersEnabled(value *bool)() {
    m.remindersEnabled = value
}
