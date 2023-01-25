package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Payload provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Payload struct {
    Entity
    // The branch of a payload. Possible values are: unknown, other, americanExpress, capitalOne, dhl, docuSign, dropbox, facebook, firstAmerican, microsoft, netflix, scotiabank, stewartTitle, tesco, wellsFargo, syrinxCloud, adobe, teams, zoom, unknownFutureValue.
    brand *PayloadBrand
    // The complexity of a payload.Possible values are: unknown, low, medium, high, unknownFutureValue
    complexity *PayloadComplexity
    // Identity of the user who created the attack simulation and training campaign payload.
    createdBy EmailIdentityable
    // Date and time when the attack simulation and training campaign payload.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the attack simulation and training campaign payload.
    description *string
    // Additional details about the payload.
    detail PayloadDetailable
    // Display name of the attack simulation and training campaign payload. Supports $filter and $orderby.
    displayName *string
    // Industry of a payload. Possible values are: unknown, other, banking, businessServices, consumerServices, education, energy, construction, consulting, financialServices, government, hospitality, insurance, legal, courierServices, IT, healthcare, manufacturing, retail, telecom, realEstate, unknownFutureValue.
    industry *PayloadIndustry
    // Indicates whether the attack simulation and training campaign payload was created from an automation flow. Supports $filter and $orderby.
    isAutomated *bool
    // Indicates whether the payload is controversial.
    isControversial *bool
    // Indicates whether the payload is from any recent event.
    isCurrentEvent *bool
    // Payload language.
    language *string
    // Identity of the user who most recently modified the attack simulation and training campaign payload.
    lastModifiedBy EmailIdentityable
    // Date and time when the attack simulation and training campaign payload was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Free text tags for a payload.
    payloadTags []string
    // The payload delivery platform for a simulation. Possible values are: unknown, sms, email, teams, unknownFutureValue.
    platform *PayloadDeliveryPlatform
    // Predicted probability for a payload to phish a targeted user.
    predictedCompromiseRate *float64
    // Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
    simulationAttackType *SimulationAttackType
    // The source property
    source *SimulationContentSource
    // Simulation content status. Supports $filter and $orderby. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue. Inherited from simulation.
    status *SimulationContentStatus
    // The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue. For more information on the types of social engineering attack techniques, see simulations.
    technique *SimulationAttackTechnique
    // The theme of a payload. Possible values are: unknown, other, accountActivation, accountVerification, billing, cleanUpMail, controversial, documentReceived, expense, incomingMessages, invoice, itemReceived, loginAlert, mailReceived, password, payment, payroll, personalizedOffer, quarantine, remoteWork, reviewMessage, securityUpdate, serviceSuspended, signatureRequired, upgradeMailboxStorage, verifyMailbox, voicemail, advertisement, employeeEngagement, unknownFutureValue.
    theme *PayloadTheme
}
// NewPayload instantiates a new payload and sets the default values.
func NewPayload()(*Payload) {
    m := &Payload{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePayloadFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayload(), nil
}
// GetBrand gets the brand property value. The branch of a payload. Possible values are: unknown, other, americanExpress, capitalOne, dhl, docuSign, dropbox, facebook, firstAmerican, microsoft, netflix, scotiabank, stewartTitle, tesco, wellsFargo, syrinxCloud, adobe, teams, zoom, unknownFutureValue.
func (m *Payload) GetBrand()(*PayloadBrand) {
    return m.brand
}
// GetComplexity gets the complexity property value. The complexity of a payload.Possible values are: unknown, low, medium, high, unknownFutureValue
func (m *Payload) GetComplexity()(*PayloadComplexity) {
    return m.complexity
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the attack simulation and training campaign payload.
func (m *Payload) GetCreatedBy()(EmailIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the attack simulation and training campaign payload.
func (m *Payload) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the attack simulation and training campaign payload.
func (m *Payload) GetDescription()(*string) {
    return m.description
}
// GetDetail gets the detail property value. Additional details about the payload.
func (m *Payload) GetDetail()(PayloadDetailable) {
    return m.detail
}
// GetDisplayName gets the displayName property value. Display name of the attack simulation and training campaign payload. Supports $filter and $orderby.
func (m *Payload) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Payload) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["brand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadBrand)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrand(val.(*PayloadBrand))
        }
        return nil
    }
    res["complexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplexity(val.(*PayloadComplexity))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
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
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PayloadDetailable))
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
    res["industry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadIndustry)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustry(val.(*PayloadIndustry))
        }
        return nil
    }
    res["isAutomated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutomated(val)
        }
        return nil
    }
    res["isControversial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsControversial(val)
        }
        return nil
    }
    res["isCurrentEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCurrentEvent(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
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
    res["payloadTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPayloadTags(res)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadDeliveryPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*PayloadDeliveryPlatform))
        }
        return nil
    }
    res["predictedCompromiseRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPredictedCompromiseRate(val)
        }
        return nil
    }
    res["simulationAttackType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationAttackType(val.(*SimulationAttackType))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*SimulationContentSource))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SimulationContentStatus))
        }
        return nil
    }
    res["technique"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackTechnique)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnique(val.(*SimulationAttackTechnique))
        }
        return nil
    }
    res["theme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadTheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val.(*PayloadTheme))
        }
        return nil
    }
    return res
}
// GetIndustry gets the industry property value. Industry of a payload. Possible values are: unknown, other, banking, businessServices, consumerServices, education, energy, construction, consulting, financialServices, government, hospitality, insurance, legal, courierServices, IT, healthcare, manufacturing, retail, telecom, realEstate, unknownFutureValue.
func (m *Payload) GetIndustry()(*PayloadIndustry) {
    return m.industry
}
// GetIsAutomated gets the isAutomated property value. Indicates whether the attack simulation and training campaign payload was created from an automation flow. Supports $filter and $orderby.
func (m *Payload) GetIsAutomated()(*bool) {
    return m.isAutomated
}
// GetIsControversial gets the isControversial property value. Indicates whether the payload is controversial.
func (m *Payload) GetIsControversial()(*bool) {
    return m.isControversial
}
// GetIsCurrentEvent gets the isCurrentEvent property value. Indicates whether the payload is from any recent event.
func (m *Payload) GetIsCurrentEvent()(*bool) {
    return m.isCurrentEvent
}
// GetLanguage gets the language property value. Payload language.
func (m *Payload) GetLanguage()(*string) {
    return m.language
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign payload.
func (m *Payload) GetLastModifiedBy()(EmailIdentityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time when the attack simulation and training campaign payload was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Payload) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPayloadTags gets the payloadTags property value. Free text tags for a payload.
func (m *Payload) GetPayloadTags()([]string) {
    return m.payloadTags
}
// GetPlatform gets the platform property value. The payload delivery platform for a simulation. Possible values are: unknown, sms, email, teams, unknownFutureValue.
func (m *Payload) GetPlatform()(*PayloadDeliveryPlatform) {
    return m.platform
}
// GetPredictedCompromiseRate gets the predictedCompromiseRate property value. Predicted probability for a payload to phish a targeted user.
func (m *Payload) GetPredictedCompromiseRate()(*float64) {
    return m.predictedCompromiseRate
}
// GetSimulationAttackType gets the simulationAttackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
func (m *Payload) GetSimulationAttackType()(*SimulationAttackType) {
    return m.simulationAttackType
}
// GetSource gets the source property value. The source property
func (m *Payload) GetSource()(*SimulationContentSource) {
    return m.source
}
// GetStatus gets the status property value. Simulation content status. Supports $filter and $orderby. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue. Inherited from simulation.
func (m *Payload) GetStatus()(*SimulationContentStatus) {
    return m.status
}
// GetTechnique gets the technique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue. For more information on the types of social engineering attack techniques, see simulations.
func (m *Payload) GetTechnique()(*SimulationAttackTechnique) {
    return m.technique
}
// GetTheme gets the theme property value. The theme of a payload. Possible values are: unknown, other, accountActivation, accountVerification, billing, cleanUpMail, controversial, documentReceived, expense, incomingMessages, invoice, itemReceived, loginAlert, mailReceived, password, payment, payroll, personalizedOffer, quarantine, remoteWork, reviewMessage, securityUpdate, serviceSuspended, signatureRequired, upgradeMailboxStorage, verifyMailbox, voicemail, advertisement, employeeEngagement, unknownFutureValue.
func (m *Payload) GetTheme()(*PayloadTheme) {
    return m.theme
}
// Serialize serializes information the current object
func (m *Payload) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBrand() != nil {
        cast := (*m.GetBrand()).String()
        err = writer.WriteStringValue("brand", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplexity() != nil {
        cast := (*m.GetComplexity()).String()
        err = writer.WriteStringValue("complexity", &cast)
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
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
    if m.GetIndustry() != nil {
        cast := (*m.GetIndustry()).String()
        err = writer.WriteStringValue("industry", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAutomated", m.GetIsAutomated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isControversial", m.GetIsControversial())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCurrentEvent", m.GetIsCurrentEvent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
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
    if m.GetPayloadTags() != nil {
        err = writer.WriteCollectionOfStringValues("payloadTags", m.GetPayloadTags())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("predictedCompromiseRate", m.GetPredictedCompromiseRate())
        if err != nil {
            return err
        }
    }
    if m.GetSimulationAttackType() != nil {
        cast := (*m.GetSimulationAttackType()).String()
        err = writer.WriteStringValue("simulationAttackType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
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
    if m.GetTechnique() != nil {
        cast := (*m.GetTechnique()).String()
        err = writer.WriteStringValue("technique", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTheme() != nil {
        cast := (*m.GetTheme()).String()
        err = writer.WriteStringValue("theme", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBrand sets the brand property value. The branch of a payload. Possible values are: unknown, other, americanExpress, capitalOne, dhl, docuSign, dropbox, facebook, firstAmerican, microsoft, netflix, scotiabank, stewartTitle, tesco, wellsFargo, syrinxCloud, adobe, teams, zoom, unknownFutureValue.
func (m *Payload) SetBrand(value *PayloadBrand)() {
    m.brand = value
}
// SetComplexity sets the complexity property value. The complexity of a payload.Possible values are: unknown, low, medium, high, unknownFutureValue
func (m *Payload) SetComplexity(value *PayloadComplexity)() {
    m.complexity = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the attack simulation and training campaign payload.
func (m *Payload) SetCreatedBy(value EmailIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the attack simulation and training campaign payload.
func (m *Payload) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the attack simulation and training campaign payload.
func (m *Payload) SetDescription(value *string)() {
    m.description = value
}
// SetDetail sets the detail property value. Additional details about the payload.
func (m *Payload) SetDetail(value PayloadDetailable)() {
    m.detail = value
}
// SetDisplayName sets the displayName property value. Display name of the attack simulation and training campaign payload. Supports $filter and $orderby.
func (m *Payload) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIndustry sets the industry property value. Industry of a payload. Possible values are: unknown, other, banking, businessServices, consumerServices, education, energy, construction, consulting, financialServices, government, hospitality, insurance, legal, courierServices, IT, healthcare, manufacturing, retail, telecom, realEstate, unknownFutureValue.
func (m *Payload) SetIndustry(value *PayloadIndustry)() {
    m.industry = value
}
// SetIsAutomated sets the isAutomated property value. Indicates whether the attack simulation and training campaign payload was created from an automation flow. Supports $filter and $orderby.
func (m *Payload) SetIsAutomated(value *bool)() {
    m.isAutomated = value
}
// SetIsControversial sets the isControversial property value. Indicates whether the payload is controversial.
func (m *Payload) SetIsControversial(value *bool)() {
    m.isControversial = value
}
// SetIsCurrentEvent sets the isCurrentEvent property value. Indicates whether the payload is from any recent event.
func (m *Payload) SetIsCurrentEvent(value *bool)() {
    m.isCurrentEvent = value
}
// SetLanguage sets the language property value. Payload language.
func (m *Payload) SetLanguage(value *string)() {
    m.language = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign payload.
func (m *Payload) SetLastModifiedBy(value EmailIdentityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time when the attack simulation and training campaign payload was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Payload) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPayloadTags sets the payloadTags property value. Free text tags for a payload.
func (m *Payload) SetPayloadTags(value []string)() {
    m.payloadTags = value
}
// SetPlatform sets the platform property value. The payload delivery platform for a simulation. Possible values are: unknown, sms, email, teams, unknownFutureValue.
func (m *Payload) SetPlatform(value *PayloadDeliveryPlatform)() {
    m.platform = value
}
// SetPredictedCompromiseRate sets the predictedCompromiseRate property value. Predicted probability for a payload to phish a targeted user.
func (m *Payload) SetPredictedCompromiseRate(value *float64)() {
    m.predictedCompromiseRate = value
}
// SetSimulationAttackType sets the simulationAttackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
func (m *Payload) SetSimulationAttackType(value *SimulationAttackType)() {
    m.simulationAttackType = value
}
// SetSource sets the source property value. The source property
func (m *Payload) SetSource(value *SimulationContentSource)() {
    m.source = value
}
// SetStatus sets the status property value. Simulation content status. Supports $filter and $orderby. Possible values are: unknown, draft, ready, archive, delete, unknownFutureValue. Inherited from simulation.
func (m *Payload) SetStatus(value *SimulationContentStatus)() {
    m.status = value
}
// SetTechnique sets the technique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue. For more information on the types of social engineering attack techniques, see simulations.
func (m *Payload) SetTechnique(value *SimulationAttackTechnique)() {
    m.technique = value
}
// SetTheme sets the theme property value. The theme of a payload. Possible values are: unknown, other, accountActivation, accountVerification, billing, cleanUpMail, controversial, documentReceived, expense, incomingMessages, invoice, itemReceived, loginAlert, mailReceived, password, payment, payroll, personalizedOffer, quarantine, remoteWork, reviewMessage, securityUpdate, serviceSuspended, signatureRequired, upgradeMailboxStorage, verifyMailbox, voicemail, advertisement, employeeEngagement, unknownFutureValue.
func (m *Payload) SetTheme(value *PayloadTheme)() {
    m.theme = value
}
