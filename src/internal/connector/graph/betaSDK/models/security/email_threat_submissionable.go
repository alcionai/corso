package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailThreatSubmissionable 
type EmailThreatSubmissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ThreatSubmissionable
    GetAttackSimulationInfo()(AttackSimulationInfoable)
    GetInternetMessageId()(*string)
    GetOriginalCategory()(*SubmissionCategory)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipientEmailAddress()(*string)
    GetSender()(*string)
    GetSenderIP()(*string)
    GetSubject()(*string)
    GetTenantAllowOrBlockListAction()(TenantAllowOrBlockListActionable)
    SetAttackSimulationInfo(value AttackSimulationInfoable)()
    SetInternetMessageId(value *string)()
    SetOriginalCategory(value *SubmissionCategory)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipientEmailAddress(value *string)()
    SetSender(value *string)()
    SetSenderIP(value *string)()
    SetSubject(value *string)()
    SetTenantAllowOrBlockListAction(value TenantAllowOrBlockListActionable)()
}
