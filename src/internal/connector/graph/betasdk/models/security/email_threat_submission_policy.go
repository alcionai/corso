package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// EmailThreatSubmissionPolicy provides operations to call the add method.
type EmailThreatSubmissionPolicy struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
    customizedNotificationSenderEmailAddress *string
    // Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
    customizedReportRecipientEmailAddress *string
    // Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
    isAlwaysReportEnabledForUsers *bool
    // Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
    isAskMeEnabledForUsers *bool
    // Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
    isCustomizedMessageEnabled *bool
    // If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
    isCustomizedMessageEnabledForPhishing *bool
    // Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
    isCustomizedNotificationSenderEnabled *bool
    // Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
    isNeverReportEnabledForUsers *bool
    // Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
    isOrganizationBrandingEnabled *bool
    // Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
    isReportFromQuarantineEnabled *bool
    // Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
    isReportToCustomizedEmailAddressEnabled *bool
    // If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
    isReportToMicrosoftEnabled *bool
    // Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
    isReviewEmailNotificationEnabled *bool
}
// NewEmailThreatSubmissionPolicy instantiates a new emailThreatSubmissionPolicy and sets the default values.
func NewEmailThreatSubmissionPolicy()(*EmailThreatSubmissionPolicy) {
    m := &EmailThreatSubmissionPolicy{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateEmailThreatSubmissionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailThreatSubmissionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailThreatSubmissionPolicy(), nil
}
// GetCustomizedNotificationSenderEmailAddress gets the customizedNotificationSenderEmailAddress property value. Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetCustomizedNotificationSenderEmailAddress()(*string) {
    return m.customizedNotificationSenderEmailAddress
}
// GetCustomizedReportRecipientEmailAddress gets the customizedReportRecipientEmailAddress property value. Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetCustomizedReportRecipientEmailAddress()(*string) {
    return m.customizedReportRecipientEmailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailThreatSubmissionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customizedNotificationSenderEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizedNotificationSenderEmailAddress(val)
        }
        return nil
    }
    res["customizedReportRecipientEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizedReportRecipientEmailAddress(val)
        }
        return nil
    }
    res["isAlwaysReportEnabledForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAlwaysReportEnabledForUsers(val)
        }
        return nil
    }
    res["isAskMeEnabledForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAskMeEnabledForUsers(val)
        }
        return nil
    }
    res["isCustomizedMessageEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomizedMessageEnabled(val)
        }
        return nil
    }
    res["isCustomizedMessageEnabledForPhishing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomizedMessageEnabledForPhishing(val)
        }
        return nil
    }
    res["isCustomizedNotificationSenderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomizedNotificationSenderEnabled(val)
        }
        return nil
    }
    res["isNeverReportEnabledForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsNeverReportEnabledForUsers(val)
        }
        return nil
    }
    res["isOrganizationBrandingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizationBrandingEnabled(val)
        }
        return nil
    }
    res["isReportFromQuarantineEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReportFromQuarantineEnabled(val)
        }
        return nil
    }
    res["isReportToCustomizedEmailAddressEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReportToCustomizedEmailAddressEnabled(val)
        }
        return nil
    }
    res["isReportToMicrosoftEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReportToMicrosoftEnabled(val)
        }
        return nil
    }
    res["isReviewEmailNotificationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReviewEmailNotificationEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsAlwaysReportEnabledForUsers gets the isAlwaysReportEnabledForUsers property value. Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsAlwaysReportEnabledForUsers()(*bool) {
    return m.isAlwaysReportEnabledForUsers
}
// GetIsAskMeEnabledForUsers gets the isAskMeEnabledForUsers property value. Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsAskMeEnabledForUsers()(*bool) {
    return m.isAskMeEnabledForUsers
}
// GetIsCustomizedMessageEnabled gets the isCustomizedMessageEnabled property value. Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabled()(*bool) {
    return m.isCustomizedMessageEnabled
}
// GetIsCustomizedMessageEnabledForPhishing gets the isCustomizedMessageEnabledForPhishing property value. If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabledForPhishing()(*bool) {
    return m.isCustomizedMessageEnabledForPhishing
}
// GetIsCustomizedNotificationSenderEnabled gets the isCustomizedNotificationSenderEnabled property value. Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedNotificationSenderEnabled()(*bool) {
    return m.isCustomizedNotificationSenderEnabled
}
// GetIsNeverReportEnabledForUsers gets the isNeverReportEnabledForUsers property value. Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsNeverReportEnabledForUsers()(*bool) {
    return m.isNeverReportEnabledForUsers
}
// GetIsOrganizationBrandingEnabled gets the isOrganizationBrandingEnabled property value. Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsOrganizationBrandingEnabled()(*bool) {
    return m.isOrganizationBrandingEnabled
}
// GetIsReportFromQuarantineEnabled gets the isReportFromQuarantineEnabled property value. Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportFromQuarantineEnabled()(*bool) {
    return m.isReportFromQuarantineEnabled
}
// GetIsReportToCustomizedEmailAddressEnabled gets the isReportToCustomizedEmailAddressEnabled property value. Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportToCustomizedEmailAddressEnabled()(*bool) {
    return m.isReportToCustomizedEmailAddressEnabled
}
// GetIsReportToMicrosoftEnabled gets the isReportToMicrosoftEnabled property value. If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportToMicrosoftEnabled()(*bool) {
    return m.isReportToMicrosoftEnabled
}
// GetIsReviewEmailNotificationEnabled gets the isReviewEmailNotificationEnabled property value. Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReviewEmailNotificationEnabled()(*bool) {
    return m.isReviewEmailNotificationEnabled
}
// Serialize serializes information the current object
func (m *EmailThreatSubmissionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("customizedNotificationSenderEmailAddress", m.GetCustomizedNotificationSenderEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customizedReportRecipientEmailAddress", m.GetCustomizedReportRecipientEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAlwaysReportEnabledForUsers", m.GetIsAlwaysReportEnabledForUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAskMeEnabledForUsers", m.GetIsAskMeEnabledForUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomizedMessageEnabled", m.GetIsCustomizedMessageEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomizedMessageEnabledForPhishing", m.GetIsCustomizedMessageEnabledForPhishing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomizedNotificationSenderEnabled", m.GetIsCustomizedNotificationSenderEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isNeverReportEnabledForUsers", m.GetIsNeverReportEnabledForUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizationBrandingEnabled", m.GetIsOrganizationBrandingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReportFromQuarantineEnabled", m.GetIsReportFromQuarantineEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReportToCustomizedEmailAddressEnabled", m.GetIsReportToCustomizedEmailAddressEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReportToMicrosoftEnabled", m.GetIsReportToMicrosoftEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReviewEmailNotificationEnabled", m.GetIsReviewEmailNotificationEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomizedNotificationSenderEmailAddress sets the customizedNotificationSenderEmailAddress property value. Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetCustomizedNotificationSenderEmailAddress(value *string)() {
    m.customizedNotificationSenderEmailAddress = value
}
// SetCustomizedReportRecipientEmailAddress sets the customizedReportRecipientEmailAddress property value. Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetCustomizedReportRecipientEmailAddress(value *string)() {
    m.customizedReportRecipientEmailAddress = value
}
// SetIsAlwaysReportEnabledForUsers sets the isAlwaysReportEnabledForUsers property value. Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsAlwaysReportEnabledForUsers(value *bool)() {
    m.isAlwaysReportEnabledForUsers = value
}
// SetIsAskMeEnabledForUsers sets the isAskMeEnabledForUsers property value. Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsAskMeEnabledForUsers(value *bool)() {
    m.isAskMeEnabledForUsers = value
}
// SetIsCustomizedMessageEnabled sets the isCustomizedMessageEnabled property value. Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabled(value *bool)() {
    m.isCustomizedMessageEnabled = value
}
// SetIsCustomizedMessageEnabledForPhishing sets the isCustomizedMessageEnabledForPhishing property value. If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabledForPhishing(value *bool)() {
    m.isCustomizedMessageEnabledForPhishing = value
}
// SetIsCustomizedNotificationSenderEnabled sets the isCustomizedNotificationSenderEnabled property value. Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedNotificationSenderEnabled(value *bool)() {
    m.isCustomizedNotificationSenderEnabled = value
}
// SetIsNeverReportEnabledForUsers sets the isNeverReportEnabledForUsers property value. Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsNeverReportEnabledForUsers(value *bool)() {
    m.isNeverReportEnabledForUsers = value
}
// SetIsOrganizationBrandingEnabled sets the isOrganizationBrandingEnabled property value. Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsOrganizationBrandingEnabled(value *bool)() {
    m.isOrganizationBrandingEnabled = value
}
// SetIsReportFromQuarantineEnabled sets the isReportFromQuarantineEnabled property value. Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportFromQuarantineEnabled(value *bool)() {
    m.isReportFromQuarantineEnabled = value
}
// SetIsReportToCustomizedEmailAddressEnabled sets the isReportToCustomizedEmailAddressEnabled property value. Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportToCustomizedEmailAddressEnabled(value *bool)() {
    m.isReportToCustomizedEmailAddressEnabled = value
}
// SetIsReportToMicrosoftEnabled sets the isReportToMicrosoftEnabled property value. If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportToMicrosoftEnabled(value *bool)() {
    m.isReportToMicrosoftEnabled = value
}
// SetIsReviewEmailNotificationEnabled sets the isReviewEmailNotificationEnabled property value. Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReviewEmailNotificationEnabled(value *bool)() {
    m.isReviewEmailNotificationEnabled = value
}
