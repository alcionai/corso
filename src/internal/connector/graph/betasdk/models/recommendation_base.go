package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendationBase provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RecommendationBase struct {
    Entity
    // The actionSteps property
    actionSteps []ActionStepable
    // The benefits property
    benefits *string
    // The category property
    category *RecommendationCategory
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The currentScore property
    currentScore *float64
    // The displayName property
    displayName *string
    // The featureAreas property
    featureAreas []RecommendationFeatureAreas
    // The impactedResources property
    impactedResources []ImpactedResourceable
    // The impactStartDateTime property
    impactStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The impactType property
    impactType *string
    // The insights property
    insights *string
    // The lastCheckedDateTime property
    lastCheckedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastModifiedBy property
    lastModifiedBy *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The maxScore property
    maxScore *float64
    // The postponeUntilDateTime property
    postponeUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The priority property
    priority *RecommendationPriority
    // The recommendationType property
    recommendationType *RecommendationType
    // The remediationImpact property
    remediationImpact *string
    // The status property
    status *RecommendationStatus
}
// NewRecommendationBase instantiates a new recommendationBase and sets the default values.
func NewRecommendationBase()(*RecommendationBase) {
    m := &RecommendationBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRecommendationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.recommendation":
                        return NewRecommendation(), nil
                }
            }
        }
    }
    return NewRecommendationBase(), nil
}
// GetActionSteps gets the actionSteps property value. The actionSteps property
func (m *RecommendationBase) GetActionSteps()([]ActionStepable) {
    return m.actionSteps
}
// GetBenefits gets the benefits property value. The benefits property
func (m *RecommendationBase) GetBenefits()(*string) {
    return m.benefits
}
// GetCategory gets the category property value. The category property
func (m *RecommendationBase) GetCategory()(*RecommendationCategory) {
    return m.category
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *RecommendationBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCurrentScore gets the currentScore property value. The currentScore property
func (m *RecommendationBase) GetCurrentScore()(*float64) {
    return m.currentScore
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *RecommendationBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFeatureAreas gets the featureAreas property value. The featureAreas property
func (m *RecommendationBase) GetFeatureAreas()([]RecommendationFeatureAreas) {
    return m.featureAreas
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecommendationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActionStepable, len(val))
            for i, v := range val {
                res[i] = v.(ActionStepable)
            }
            m.SetActionSteps(res)
        }
        return nil
    }
    res["benefits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBenefits(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*RecommendationCategory))
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
    res["currentScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentScore(val)
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
    res["featureAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRecommendationFeatureAreas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendationFeatureAreas, len(val))
            for i, v := range val {
                res[i] = *(v.(*RecommendationFeatureAreas))
            }
            m.SetFeatureAreas(res)
        }
        return nil
    }
    res["impactedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImpactedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImpactedResourceable, len(val))
            for i, v := range val {
                res[i] = v.(ImpactedResourceable)
            }
            m.SetImpactedResources(res)
        }
        return nil
    }
    res["impactStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactStartDateTime(val)
        }
        return nil
    }
    res["impactType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactType(val)
        }
        return nil
    }
    res["insights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val)
        }
        return nil
    }
    res["lastCheckedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckedDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
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
    res["maxScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val)
        }
        return nil
    }
    res["postponeUntilDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostponeUntilDateTime(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*RecommendationPriority))
        }
        return nil
    }
    res["recommendationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationType(val.(*RecommendationType))
        }
        return nil
    }
    res["remediationImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationImpact(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*RecommendationStatus))
        }
        return nil
    }
    return res
}
// GetImpactedResources gets the impactedResources property value. The impactedResources property
func (m *RecommendationBase) GetImpactedResources()([]ImpactedResourceable) {
    return m.impactedResources
}
// GetImpactStartDateTime gets the impactStartDateTime property value. The impactStartDateTime property
func (m *RecommendationBase) GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.impactStartDateTime
}
// GetImpactType gets the impactType property value. The impactType property
func (m *RecommendationBase) GetImpactType()(*string) {
    return m.impactType
}
// GetInsights gets the insights property value. The insights property
func (m *RecommendationBase) GetInsights()(*string) {
    return m.insights
}
// GetLastCheckedDateTime gets the lastCheckedDateTime property value. The lastCheckedDateTime property
func (m *RecommendationBase) GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCheckedDateTime
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *RecommendationBase) GetLastModifiedBy()(*string) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *RecommendationBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMaxScore gets the maxScore property value. The maxScore property
func (m *RecommendationBase) GetMaxScore()(*float64) {
    return m.maxScore
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *RecommendationBase) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.postponeUntilDateTime
}
// GetPriority gets the priority property value. The priority property
func (m *RecommendationBase) GetPriority()(*RecommendationPriority) {
    return m.priority
}
// GetRecommendationType gets the recommendationType property value. The recommendationType property
func (m *RecommendationBase) GetRecommendationType()(*RecommendationType) {
    return m.recommendationType
}
// GetRemediationImpact gets the remediationImpact property value. The remediationImpact property
func (m *RecommendationBase) GetRemediationImpact()(*string) {
    return m.remediationImpact
}
// GetStatus gets the status property value. The status property
func (m *RecommendationBase) GetStatus()(*RecommendationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *RecommendationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActionSteps()))
        for i, v := range m.GetActionSteps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("actionSteps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("benefits", m.GetBenefits())
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
        err = writer.WriteFloat64Value("currentScore", m.GetCurrentScore())
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
    if m.GetFeatureAreas() != nil {
        err = writer.WriteCollectionOfStringValues("featureAreas", SerializeRecommendationFeatureAreas(m.GetFeatureAreas()))
        if err != nil {
            return err
        }
    }
    if m.GetImpactedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImpactedResources()))
        for i, v := range m.GetImpactedResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("impactedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("impactStartDateTime", m.GetImpactStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("impactType", m.GetImpactType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckedDateTime", m.GetLastCheckedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err = writer.WriteFloat64Value("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("postponeUntilDateTime", m.GetPostponeUntilDateTime())
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
    if m.GetRecommendationType() != nil {
        cast := (*m.GetRecommendationType()).String()
        err = writer.WriteStringValue("recommendationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediationImpact", m.GetRemediationImpact())
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
// SetActionSteps sets the actionSteps property value. The actionSteps property
func (m *RecommendationBase) SetActionSteps(value []ActionStepable)() {
    m.actionSteps = value
}
// SetBenefits sets the benefits property value. The benefits property
func (m *RecommendationBase) SetBenefits(value *string)() {
    m.benefits = value
}
// SetCategory sets the category property value. The category property
func (m *RecommendationBase) SetCategory(value *RecommendationCategory)() {
    m.category = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *RecommendationBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCurrentScore sets the currentScore property value. The currentScore property
func (m *RecommendationBase) SetCurrentScore(value *float64)() {
    m.currentScore = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *RecommendationBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFeatureAreas sets the featureAreas property value. The featureAreas property
func (m *RecommendationBase) SetFeatureAreas(value []RecommendationFeatureAreas)() {
    m.featureAreas = value
}
// SetImpactedResources sets the impactedResources property value. The impactedResources property
func (m *RecommendationBase) SetImpactedResources(value []ImpactedResourceable)() {
    m.impactedResources = value
}
// SetImpactStartDateTime sets the impactStartDateTime property value. The impactStartDateTime property
func (m *RecommendationBase) SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.impactStartDateTime = value
}
// SetImpactType sets the impactType property value. The impactType property
func (m *RecommendationBase) SetImpactType(value *string)() {
    m.impactType = value
}
// SetInsights sets the insights property value. The insights property
func (m *RecommendationBase) SetInsights(value *string)() {
    m.insights = value
}
// SetLastCheckedDateTime sets the lastCheckedDateTime property value. The lastCheckedDateTime property
func (m *RecommendationBase) SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckedDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *RecommendationBase) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *RecommendationBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMaxScore sets the maxScore property value. The maxScore property
func (m *RecommendationBase) SetMaxScore(value *float64)() {
    m.maxScore = value
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *RecommendationBase) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.postponeUntilDateTime = value
}
// SetPriority sets the priority property value. The priority property
func (m *RecommendationBase) SetPriority(value *RecommendationPriority)() {
    m.priority = value
}
// SetRecommendationType sets the recommendationType property value. The recommendationType property
func (m *RecommendationBase) SetRecommendationType(value *RecommendationType)() {
    m.recommendationType = value
}
// SetRemediationImpact sets the remediationImpact property value. The remediationImpact property
func (m *RecommendationBase) SetRemediationImpact(value *string)() {
    m.remediationImpact = value
}
// SetStatus sets the status property value. The status property
func (m *RecommendationBase) SetStatus(value *RecommendationStatus)() {
    m.status = value
}
