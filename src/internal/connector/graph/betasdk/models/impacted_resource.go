package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImpactedResource provides operations to call the add method.
type ImpactedResource struct {
    Entity
    // The addedDateTime property
    addedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The additionalDetails property
    additionalDetails []KeyValueable
    // The apiUrl property
    apiUrl *string
    // The displayName property
    displayName *string
    // The lastModifiedBy property
    lastModifiedBy *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *string
    // The owner property
    owner *string
    // The portalUrl property
    portalUrl *string
    // The postponeUntilDateTime property
    postponeUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The rank property
    rank *int32
    // The recommendationId property
    recommendationId *string
    // The resourceType property
    resourceType *string
    // The status property
    status *RecommendationStatus
    // The subjectId property
    subjectId *string
}
// NewImpactedResource instantiates a new impactedResource and sets the default values.
func NewImpactedResource()(*ImpactedResource) {
    m := &ImpactedResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateImpactedResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImpactedResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImpactedResource(), nil
}
// GetAddedDateTime gets the addedDateTime property value. The addedDateTime property
func (m *ImpactedResource) GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.addedDateTime
}
// GetAdditionalDetails gets the additionalDetails property value. The additionalDetails property
func (m *ImpactedResource) GetAdditionalDetails()([]KeyValueable) {
    return m.additionalDetails
}
// GetApiUrl gets the apiUrl property value. The apiUrl property
func (m *ImpactedResource) GetApiUrl()(*string) {
    return m.apiUrl
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ImpactedResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImpactedResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedDateTime(val)
        }
        return nil
    }
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValueable)
            }
            m.SetAdditionalDetails(res)
        }
        return nil
    }
    res["apiUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiUrl(val)
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["portalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPortalUrl(val)
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
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["recommendationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationId(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
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
    res["subjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectId(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *ImpactedResource) GetLastModifiedBy()(*string) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ImpactedResource) GetLastModifiedDateTime()(*string) {
    return m.lastModifiedDateTime
}
// GetOwner gets the owner property value. The owner property
func (m *ImpactedResource) GetOwner()(*string) {
    return m.owner
}
// GetPortalUrl gets the portalUrl property value. The portalUrl property
func (m *ImpactedResource) GetPortalUrl()(*string) {
    return m.portalUrl
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *ImpactedResource) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.postponeUntilDateTime
}
// GetRank gets the rank property value. The rank property
func (m *ImpactedResource) GetRank()(*int32) {
    return m.rank
}
// GetRecommendationId gets the recommendationId property value. The recommendationId property
func (m *ImpactedResource) GetRecommendationId()(*string) {
    return m.recommendationId
}
// GetResourceType gets the resourceType property value. The resourceType property
func (m *ImpactedResource) GetResourceType()(*string) {
    return m.resourceType
}
// GetStatus gets the status property value. The status property
func (m *ImpactedResource) GetStatus()(*RecommendationStatus) {
    return m.status
}
// GetSubjectId gets the subjectId property value. The subjectId property
func (m *ImpactedResource) GetSubjectId()(*string) {
    return m.subjectId
}
// Serialize serializes information the current object
func (m *ImpactedResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("addedDateTime", m.GetAddedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetAdditionalDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalDetails()))
        for i, v := range m.GetAdditionalDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("additionalDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("apiUrl", m.GetApiUrl())
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
    {
        err = writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
        err = writer.WriteStringValue("portalUrl", m.GetPortalUrl())
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
    {
        err = writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendationId", m.GetRecommendationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceType", m.GetResourceType())
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
        err = writer.WriteStringValue("subjectId", m.GetSubjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedDateTime sets the addedDateTime property value. The addedDateTime property
func (m *ImpactedResource) SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.addedDateTime = value
}
// SetAdditionalDetails sets the additionalDetails property value. The additionalDetails property
func (m *ImpactedResource) SetAdditionalDetails(value []KeyValueable)() {
    m.additionalDetails = value
}
// SetApiUrl sets the apiUrl property value. The apiUrl property
func (m *ImpactedResource) SetApiUrl(value *string)() {
    m.apiUrl = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ImpactedResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *ImpactedResource) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ImpactedResource) SetLastModifiedDateTime(value *string)() {
    m.lastModifiedDateTime = value
}
// SetOwner sets the owner property value. The owner property
func (m *ImpactedResource) SetOwner(value *string)() {
    m.owner = value
}
// SetPortalUrl sets the portalUrl property value. The portalUrl property
func (m *ImpactedResource) SetPortalUrl(value *string)() {
    m.portalUrl = value
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *ImpactedResource) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.postponeUntilDateTime = value
}
// SetRank sets the rank property value. The rank property
func (m *ImpactedResource) SetRank(value *int32)() {
    m.rank = value
}
// SetRecommendationId sets the recommendationId property value. The recommendationId property
func (m *ImpactedResource) SetRecommendationId(value *string)() {
    m.recommendationId = value
}
// SetResourceType sets the resourceType property value. The resourceType property
func (m *ImpactedResource) SetResourceType(value *string)() {
    m.resourceType = value
}
// SetStatus sets the status property value. The status property
func (m *ImpactedResource) SetStatus(value *RecommendationStatus)() {
    m.status = value
}
// SetSubjectId sets the subjectId property value. The subjectId property
func (m *ImpactedResource) SetSubjectId(value *string)() {
    m.subjectId = value
}
