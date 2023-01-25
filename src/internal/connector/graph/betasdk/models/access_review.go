package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReview 
type AccessReview struct {
    Entity
    // The business flow template identifier. Required on create.  This value is case sensitive.
    businessFlowTemplateId *string
    // The user who created this review.
    createdBy UserIdentityable
    // The collection of decisions for this access review.
    decisions []AccessReviewDecisionable
    // The description provided by the access review creator, to show to the reviewers.
    description *string
    // The access review name. Required on create.
    displayName *string
    // The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of access reviews instances past, present and future, if this object is a recurring access review.
    instances []AccessReviewable
    // The collection of decisions for the caller, if the caller is a reviewer.
    myDecisions []AccessReviewDecisionable
    // The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
    reviewedEntity Identityable
    // The collection of reviewers for an access review, if access review reviewerType is of type delegated.
    reviewers []AccessReviewReviewerable
    // The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
    reviewerType *string
    // The settings of an accessReview, see type definition below.
    settings AccessReviewSettingsable
    // The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
    status *string
}
// NewAccessReview instantiates a new AccessReview and sets the default values.
func NewAccessReview()(*AccessReview) {
    m := &AccessReview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReview(), nil
}
// GetBusinessFlowTemplateId gets the businessFlowTemplateId property value. The business flow template identifier. Required on create.  This value is case sensitive.
func (m *AccessReview) GetBusinessFlowTemplateId()(*string) {
    return m.businessFlowTemplateId
}
// GetCreatedBy gets the createdBy property value. The user who created this review.
func (m *AccessReview) GetCreatedBy()(UserIdentityable) {
    return m.createdBy
}
// GetDecisions gets the decisions property value. The collection of decisions for this access review.
func (m *AccessReview) GetDecisions()([]AccessReviewDecisionable) {
    return m.decisions
}
// GetDescription gets the description property value. The description provided by the access review creator, to show to the reviewers.
func (m *AccessReview) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The access review name. Required on create.
func (m *AccessReview) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDateTime gets the endDateTime property value. The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
func (m *AccessReview) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlowTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessFlowTemplateId(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["decisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewDecisionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewDecisionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewDecisionable)
            }
            m.SetDecisions(res)
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["instances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewable)
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["myDecisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewDecisionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewDecisionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewDecisionable)
            }
            m.SetMyDecisions(res)
        }
        return nil
    }
    res["reviewedEntity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedEntity(val.(Identityable))
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewReviewerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewReviewerable)
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["reviewerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewerType(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(AccessReviewSettingsable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetInstances gets the instances property value. The collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *AccessReview) GetInstances()([]AccessReviewable) {
    return m.instances
}
// GetMyDecisions gets the myDecisions property value. The collection of decisions for the caller, if the caller is a reviewer.
func (m *AccessReview) GetMyDecisions()([]AccessReviewDecisionable) {
    return m.myDecisions
}
// GetReviewedEntity gets the reviewedEntity property value. The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
func (m *AccessReview) GetReviewedEntity()(Identityable) {
    return m.reviewedEntity
}
// GetReviewers gets the reviewers property value. The collection of reviewers for an access review, if access review reviewerType is of type delegated.
func (m *AccessReview) GetReviewers()([]AccessReviewReviewerable) {
    return m.reviewers
}
// GetReviewerType gets the reviewerType property value. The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
func (m *AccessReview) GetReviewerType()(*string) {
    return m.reviewerType
}
// GetSettings gets the settings property value. The settings of an accessReview, see type definition below.
func (m *AccessReview) GetSettings()(AccessReviewSettingsable) {
    return m.settings
}
// GetStartDateTime gets the startDateTime property value. The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
func (m *AccessReview) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
func (m *AccessReview) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *AccessReview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("businessFlowTemplateId", m.GetBusinessFlowTemplateId())
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
    if m.GetDecisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDecisions()))
        for i, v := range m.GetDecisions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("decisions", cast)
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMyDecisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMyDecisions()))
        for i, v := range m.GetMyDecisions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("myDecisions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedEntity", m.GetReviewedEntity())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewerType", m.GetReviewerType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBusinessFlowTemplateId sets the businessFlowTemplateId property value. The business flow template identifier. Required on create.  This value is case sensitive.
func (m *AccessReview) SetBusinessFlowTemplateId(value *string)() {
    m.businessFlowTemplateId = value
}
// SetCreatedBy sets the createdBy property value. The user who created this review.
func (m *AccessReview) SetCreatedBy(value UserIdentityable)() {
    m.createdBy = value
}
// SetDecisions sets the decisions property value. The collection of decisions for this access review.
func (m *AccessReview) SetDecisions(value []AccessReviewDecisionable)() {
    m.decisions = value
}
// SetDescription sets the description property value. The description provided by the access review creator, to show to the reviewers.
func (m *AccessReview) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The access review name. Required on create.
func (m *AccessReview) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDateTime sets the endDateTime property value. The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
func (m *AccessReview) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetInstances sets the instances property value. The collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *AccessReview) SetInstances(value []AccessReviewable)() {
    m.instances = value
}
// SetMyDecisions sets the myDecisions property value. The collection of decisions for the caller, if the caller is a reviewer.
func (m *AccessReview) SetMyDecisions(value []AccessReviewDecisionable)() {
    m.myDecisions = value
}
// SetReviewedEntity sets the reviewedEntity property value. The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
func (m *AccessReview) SetReviewedEntity(value Identityable)() {
    m.reviewedEntity = value
}
// SetReviewers sets the reviewers property value. The collection of reviewers for an access review, if access review reviewerType is of type delegated.
func (m *AccessReview) SetReviewers(value []AccessReviewReviewerable)() {
    m.reviewers = value
}
// SetReviewerType sets the reviewerType property value. The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
func (m *AccessReview) SetReviewerType(value *string)() {
    m.reviewerType = value
}
// SetSettings sets the settings property value. The settings of an accessReview, see type definition below.
func (m *AccessReview) SetSettings(value AccessReviewSettingsable)() {
    m.settings = value
}
// SetStartDateTime sets the startDateTime property value. The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
func (m *AccessReview) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
func (m *AccessReview) SetStatus(value *string)() {
    m.status = value
}
