package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequest provides operations to manage the collection of site entities.
type AccessPackageAssignmentRequest struct {
    Entity
    // The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable. Supports $expand.
    accessPackage AccessPackageable
    // For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
    accessPackageAssignment AccessPackageAssignmentable
    // Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
    answers []AccessPackageAnswerable
    // The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    completedDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of custom workflow extension instances being run on an assignment request. Read-only.
    customExtensionHandlerInstances []CustomExtensionHandlerInstanceable
    // The expirationDateTime property
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // True if the request is not to be processed for assignment.
    isValidationOnly *bool
    // The requestor's supplied justification.
    justification *string
    // The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
    requestor AccessPackageSubjectable
    // One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
    requestState *string
    // More information on the request processing status. Read-only.
    requestStatus *string
    // One of UserAdd, UserExtend, UserUpdate, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd, UserUpdate or UserRemove. Read-only.
    requestType *string
    // The range of dates that access is to be assigned to the requestor. Read-only.
    schedule RequestScheduleable
}
// NewAccessPackageAssignmentRequest instantiates a new accessPackageAssignmentRequest and sets the default values.
func NewAccessPackageAssignmentRequest()(*AccessPackageAssignmentRequest) {
    m := &AccessPackageAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequest(), nil
}
// GetAccessPackage gets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackage()(AccessPackageable) {
    return m.accessPackage
}
// GetAccessPackageAssignment gets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackageAssignment()(AccessPackageAssignmentable) {
    return m.accessPackageAssignment
}
// GetAnswers gets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
func (m *AccessPackageAssignmentRequest) GetAnswers()([]AccessPackageAnswerable) {
    return m.answers
}
// GetCompletedDate gets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCompletedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDate
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomExtensionHandlerInstances gets the customExtensionHandlerInstances property value. A collection of custom workflow extension instances being run on an assignment request. Read-only.
func (m *AccessPackageAssignmentRequest) GetCustomExtensionHandlerInstances()([]CustomExtensionHandlerInstanceable) {
    return m.customExtensionHandlerInstances
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *AccessPackageAssignmentRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(AccessPackageable))
        }
        return nil
    }
    res["accessPackageAssignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageAssignment(val.(AccessPackageAssignmentable))
        }
        return nil
    }
    res["answers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAnswerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAnswerable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAnswerable)
            }
            m.SetAnswers(res)
        }
        return nil
    }
    res["completedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDate(val)
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
    res["customExtensionHandlerInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomExtensionHandlerInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionHandlerInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(CustomExtensionHandlerInstanceable)
            }
            m.SetCustomExtensionHandlerInstances(res)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["isValidationOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["requestor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(AccessPackageSubjectable))
        }
        return nil
    }
    res["requestState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestState(val)
        }
        return nil
    }
    res["requestStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestStatus(val)
        }
        return nil
    }
    res["requestType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(RequestScheduleable))
        }
        return nil
    }
    return res
}
// GetIsValidationOnly gets the isValidationOnly property value. True if the request is not to be processed for assignment.
func (m *AccessPackageAssignmentRequest) GetIsValidationOnly()(*bool) {
    return m.isValidationOnly
}
// GetJustification gets the justification property value. The requestor's supplied justification.
func (m *AccessPackageAssignmentRequest) GetJustification()(*string) {
    return m.justification
}
// GetRequestor gets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetRequestor()(AccessPackageSubjectable) {
    return m.requestor
}
// GetRequestState gets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestState()(*string) {
    return m.requestState
}
// GetRequestStatus gets the requestStatus property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestStatus()(*string) {
    return m.requestStatus
}
// GetRequestType gets the requestType property value. One of UserAdd, UserExtend, UserUpdate, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd, UserUpdate or UserRemove. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestType()(*string) {
    return m.requestType
}
// GetSchedule gets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
func (m *AccessPackageAssignmentRequest) GetSchedule()(RequestScheduleable) {
    return m.schedule
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackage", m.GetAccessPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageAssignment", m.GetAccessPackageAssignment())
        if err != nil {
            return err
        }
    }
    if m.GetAnswers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAnswers()))
        for i, v := range m.GetAnswers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("answers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDate", m.GetCompletedDate())
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
    if m.GetCustomExtensionHandlerInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomExtensionHandlerInstances()))
        for i, v := range m.GetCustomExtensionHandlerInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customExtensionHandlerInstances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValidationOnly", m.GetIsValidationOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestor", m.GetRequestor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestState", m.GetRequestState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestStatus", m.GetRequestStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestType", m.GetRequestType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAccessPackage(value AccessPackageable)() {
    m.accessPackage = value
}
// SetAccessPackageAssignment sets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAccessPackageAssignment(value AccessPackageAssignmentable)() {
    m.accessPackageAssignment = value
}
// SetAnswers sets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
func (m *AccessPackageAssignmentRequest) SetAnswers(value []AccessPackageAnswerable)() {
    m.answers = value
}
// SetCompletedDate sets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) SetCompletedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDate = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomExtensionHandlerInstances sets the customExtensionHandlerInstances property value. A collection of custom workflow extension instances being run on an assignment request. Read-only.
func (m *AccessPackageAssignmentRequest) SetCustomExtensionHandlerInstances(value []CustomExtensionHandlerInstanceable)() {
    m.customExtensionHandlerInstances = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *AccessPackageAssignmentRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIsValidationOnly sets the isValidationOnly property value. True if the request is not to be processed for assignment.
func (m *AccessPackageAssignmentRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// SetJustification sets the justification property value. The requestor's supplied justification.
func (m *AccessPackageAssignmentRequest) SetJustification(value *string)() {
    m.justification = value
}
// SetRequestor sets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) SetRequestor(value AccessPackageSubjectable)() {
    m.requestor = value
}
// SetRequestState sets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestState(value *string)() {
    m.requestState = value
}
// SetRequestStatus sets the requestStatus property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestStatus(value *string)() {
    m.requestStatus = value
}
// SetRequestType sets the requestType property value. One of UserAdd, UserExtend, UserUpdate, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd, UserUpdate or UserRemove. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestType(value *string)() {
    m.requestType = value
}
// SetSchedule sets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
func (m *AccessPackageAssignmentRequest) SetSchedule(value RequestScheduleable)() {
    m.schedule = value
}
