package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailThreatSubmission 
type EmailThreatSubmission struct {
    ThreatSubmission
    // If the email is phishing simulation, this field will not be null.
    attackSimulationInfo AttackSimulationInfoable
    // Specifies the internet message id of the email being submitted. This information is present in the email header.
    internetMessageId *string
    // The original category of the submission. The possible values are: notJunk, spam, phishing, malware and unkownFutureValue.
    originalCategory *SubmissionCategory
    // Specifies the date and time stamp when the email was received.
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the email address (in smtp format) of the recipient who received the email.
    recipientEmailAddress *string
    // Specifies the email address of the sender.
    sender *string
    // Specifies the IP address of the sender.
    senderIP *string
    // Specifies the subject of the email .
    subject *string
    // It is used to automatically add allows for the components such as URL, file, sender; which are deemed bad by Microsoft so that similar messages in the future can be allowed.
    tenantAllowOrBlockListAction TenantAllowOrBlockListActionable
}
// NewEmailThreatSubmission instantiates a new EmailThreatSubmission and sets the default values.
func NewEmailThreatSubmission()(*EmailThreatSubmission) {
    m := &EmailThreatSubmission{
        ThreatSubmission: *NewThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.emailThreatSubmission";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEmailThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.security.emailContentThreatSubmission":
                        return NewEmailContentThreatSubmission(), nil
                    case "#microsoft.graph.security.emailUrlThreatSubmission":
                        return NewEmailUrlThreatSubmission(), nil
                }
            }
        }
    }
    return NewEmailThreatSubmission(), nil
}
// GetAttackSimulationInfo gets the attackSimulationInfo property value. If the email is phishing simulation, this field will not be null.
func (m *EmailThreatSubmission) GetAttackSimulationInfo()(AttackSimulationInfoable) {
    return m.attackSimulationInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatSubmission.GetFieldDeserializers()
    res["attackSimulationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulationInfo(val.(AttackSimulationInfoable))
        }
        return nil
    }
    res["internetMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["originalCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalCategory(val.(*SubmissionCategory))
        }
        return nil
    }
    res["receivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["recipientEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientEmailAddress(val)
        }
        return nil
    }
    res["sender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSender(val)
        }
        return nil
    }
    res["senderIP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderIP(val)
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
    res["tenantAllowOrBlockListAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantAllowOrBlockListActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAllowOrBlockListAction(val.(TenantAllowOrBlockListActionable))
        }
        return nil
    }
    return res
}
// GetInternetMessageId gets the internetMessageId property value. Specifies the internet message id of the email being submitted. This information is present in the email header.
func (m *EmailThreatSubmission) GetInternetMessageId()(*string) {
    return m.internetMessageId
}
// GetOriginalCategory gets the originalCategory property value. The original category of the submission. The possible values are: notJunk, spam, phishing, malware and unkownFutureValue.
func (m *EmailThreatSubmission) GetOriginalCategory()(*SubmissionCategory) {
    return m.originalCategory
}
// GetReceivedDateTime gets the receivedDateTime property value. Specifies the date and time stamp when the email was received.
func (m *EmailThreatSubmission) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
// GetRecipientEmailAddress gets the recipientEmailAddress property value. Specifies the email address (in smtp format) of the recipient who received the email.
func (m *EmailThreatSubmission) GetRecipientEmailAddress()(*string) {
    return m.recipientEmailAddress
}
// GetSender gets the sender property value. Specifies the email address of the sender.
func (m *EmailThreatSubmission) GetSender()(*string) {
    return m.sender
}
// GetSenderIP gets the senderIP property value. Specifies the IP address of the sender.
func (m *EmailThreatSubmission) GetSenderIP()(*string) {
    return m.senderIP
}
// GetSubject gets the subject property value. Specifies the subject of the email .
func (m *EmailThreatSubmission) GetSubject()(*string) {
    return m.subject
}
// GetTenantAllowOrBlockListAction gets the tenantAllowOrBlockListAction property value. It is used to automatically add allows for the components such as URL, file, sender; which are deemed bad by Microsoft so that similar messages in the future can be allowed.
func (m *EmailThreatSubmission) GetTenantAllowOrBlockListAction()(TenantAllowOrBlockListActionable) {
    return m.tenantAllowOrBlockListAction
}
// Serialize serializes information the current object
func (m *EmailThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("attackSimulationInfo", m.GetAttackSimulationInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    if m.GetOriginalCategory() != nil {
        cast := (*m.GetOriginalCategory()).String()
        err = writer.WriteStringValue("originalCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientEmailAddress", m.GetRecipientEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sender", m.GetSender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderIP", m.GetSenderIP())
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
    {
        err = writer.WriteObjectValue("tenantAllowOrBlockListAction", m.GetTenantAllowOrBlockListAction())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttackSimulationInfo sets the attackSimulationInfo property value. If the email is phishing simulation, this field will not be null.
func (m *EmailThreatSubmission) SetAttackSimulationInfo(value AttackSimulationInfoable)() {
    m.attackSimulationInfo = value
}
// SetInternetMessageId sets the internetMessageId property value. Specifies the internet message id of the email being submitted. This information is present in the email header.
func (m *EmailThreatSubmission) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetOriginalCategory sets the originalCategory property value. The original category of the submission. The possible values are: notJunk, spam, phishing, malware and unkownFutureValue.
func (m *EmailThreatSubmission) SetOriginalCategory(value *SubmissionCategory)() {
    m.originalCategory = value
}
// SetReceivedDateTime sets the receivedDateTime property value. Specifies the date and time stamp when the email was received.
func (m *EmailThreatSubmission) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// SetRecipientEmailAddress sets the recipientEmailAddress property value. Specifies the email address (in smtp format) of the recipient who received the email.
func (m *EmailThreatSubmission) SetRecipientEmailAddress(value *string)() {
    m.recipientEmailAddress = value
}
// SetSender sets the sender property value. Specifies the email address of the sender.
func (m *EmailThreatSubmission) SetSender(value *string)() {
    m.sender = value
}
// SetSenderIP sets the senderIP property value. Specifies the IP address of the sender.
func (m *EmailThreatSubmission) SetSenderIP(value *string)() {
    m.senderIP = value
}
// SetSubject sets the subject property value. Specifies the subject of the email .
func (m *EmailThreatSubmission) SetSubject(value *string)() {
    m.subject = value
}
// SetTenantAllowOrBlockListAction sets the tenantAllowOrBlockListAction property value. It is used to automatically add allows for the components such as URL, file, sender; which are deemed bad by Microsoft so that similar messages in the future can be allowed.
func (m *EmailThreatSubmission) SetTenantAllowOrBlockListAction(value TenantAllowOrBlockListActionable)() {
    m.tenantAllowOrBlockListAction = value
}
