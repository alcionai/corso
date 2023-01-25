package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpPoliciesJobResult 
type DlpPoliciesJobResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The auditCorrelationId property
    auditCorrelationId *string
    // The evaluationDateTime property
    evaluationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The matchingRules property
    matchingRules []MatchingDlpRuleable
    // The OdataType property
    odataType *string
}
// NewDlpPoliciesJobResult instantiates a new dlpPoliciesJobResult and sets the default values.
func NewDlpPoliciesJobResult()(*DlpPoliciesJobResult) {
    m := &DlpPoliciesJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDlpPoliciesJobResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpPoliciesJobResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDlpPoliciesJobResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpPoliciesJobResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAuditCorrelationId gets the auditCorrelationId property value. The auditCorrelationId property
func (m *DlpPoliciesJobResult) GetAuditCorrelationId()(*string) {
    return m.auditCorrelationId
}
// GetEvaluationDateTime gets the evaluationDateTime property value. The evaluationDateTime property
func (m *DlpPoliciesJobResult) GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.evaluationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpPoliciesJobResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["auditCorrelationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditCorrelationId(val)
        }
        return nil
    }
    res["evaluationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationDateTime(val)
        }
        return nil
    }
    res["matchingRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMatchingDlpRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MatchingDlpRuleable, len(val))
            for i, v := range val {
                res[i] = v.(MatchingDlpRuleable)
            }
            m.SetMatchingRules(res)
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
    return res
}
// GetMatchingRules gets the matchingRules property value. The matchingRules property
func (m *DlpPoliciesJobResult) GetMatchingRules()([]MatchingDlpRuleable) {
    return m.matchingRules
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DlpPoliciesJobResult) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DlpPoliciesJobResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("auditCorrelationId", m.GetAuditCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("evaluationDateTime", m.GetEvaluationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMatchingRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatchingRules()))
        for i, v := range m.GetMatchingRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("matchingRules", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpPoliciesJobResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAuditCorrelationId sets the auditCorrelationId property value. The auditCorrelationId property
func (m *DlpPoliciesJobResult) SetAuditCorrelationId(value *string)() {
    m.auditCorrelationId = value
}
// SetEvaluationDateTime sets the evaluationDateTime property value. The evaluationDateTime property
func (m *DlpPoliciesJobResult) SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.evaluationDateTime = value
}
// SetMatchingRules sets the matchingRules property value. The matchingRules property
func (m *DlpPoliciesJobResult) SetMatchingRules(value []MatchingDlpRuleable)() {
    m.matchingRules = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DlpPoliciesJobResult) SetOdataType(value *string)() {
    m.odataType = value
}
