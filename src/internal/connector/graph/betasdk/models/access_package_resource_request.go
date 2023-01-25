package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceRequest provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AccessPackageResourceRequest struct {
    Entity
    // The accessPackageResource property
    accessPackageResource AccessPackageResourceable
    // The unique ID of the access package catalog.
    catalogId *string
    // The executeImmediately property
    executeImmediately *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // If set, does not add the resource.
    isValidationOnly *bool
    // The requestor's justification for adding or removing the resource.
    justification *string
    // Read-only. Nullable. Supports $expand.
    requestor AccessPackageSubjectable
    // The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
    requestState *string
    // The requestStatus property
    requestStatus *string
    // Use AdminAdd to add a resource, if the caller is an administrator or resource owner, AdminUpdate to update a resource, or AdminRemove to remove a resource.
    requestType *string
}
// NewAccessPackageResourceRequest instantiates a new accessPackageResourceRequest and sets the default values.
func NewAccessPackageResourceRequest()(*AccessPackageResourceRequest) {
    m := &AccessPackageResourceRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceRequest(), nil
}
// GetAccessPackageResource gets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceRequest) GetAccessPackageResource()(AccessPackageResourceable) {
    return m.accessPackageResource
}
// GetCatalogId gets the catalogId property value. The unique ID of the access package catalog.
func (m *AccessPackageResourceRequest) GetCatalogId()(*string) {
    return m.catalogId
}
// GetExecuteImmediately gets the executeImmediately property value. The executeImmediately property
func (m *AccessPackageResourceRequest) GetExecuteImmediately()(*bool) {
    return m.executeImmediately
}
// GetExpirationDateTime gets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResourceRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResource(val.(AccessPackageResourceable))
        }
        return nil
    }
    res["catalogId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    res["executeImmediately"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecuteImmediately(val)
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
    return res
}
// GetIsValidationOnly gets the isValidationOnly property value. If set, does not add the resource.
func (m *AccessPackageResourceRequest) GetIsValidationOnly()(*bool) {
    return m.isValidationOnly
}
// GetJustification gets the justification property value. The requestor's justification for adding or removing the resource.
func (m *AccessPackageResourceRequest) GetJustification()(*string) {
    return m.justification
}
// GetRequestor gets the requestor property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResourceRequest) GetRequestor()(AccessPackageSubjectable) {
    return m.requestor
}
// GetRequestState gets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
func (m *AccessPackageResourceRequest) GetRequestState()(*string) {
    return m.requestState
}
// GetRequestStatus gets the requestStatus property value. The requestStatus property
func (m *AccessPackageResourceRequest) GetRequestStatus()(*string) {
    return m.requestStatus
}
// GetRequestType gets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, AdminUpdate to update a resource, or AdminRemove to remove a resource.
func (m *AccessPackageResourceRequest) GetRequestType()(*string) {
    return m.requestType
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackageResource", m.GetAccessPackageResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("catalogId", m.GetCatalogId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("executeImmediately", m.GetExecuteImmediately())
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
    return nil
}
// SetAccessPackageResource sets the accessPackageResource property value. The accessPackageResource property
func (m *AccessPackageResourceRequest) SetAccessPackageResource(value AccessPackageResourceable)() {
    m.accessPackageResource = value
}
// SetCatalogId sets the catalogId property value. The unique ID of the access package catalog.
func (m *AccessPackageResourceRequest) SetCatalogId(value *string)() {
    m.catalogId = value
}
// SetExecuteImmediately sets the executeImmediately property value. The executeImmediately property
func (m *AccessPackageResourceRequest) SetExecuteImmediately(value *bool)() {
    m.executeImmediately = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageResourceRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIsValidationOnly sets the isValidationOnly property value. If set, does not add the resource.
func (m *AccessPackageResourceRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// SetJustification sets the justification property value. The requestor's justification for adding or removing the resource.
func (m *AccessPackageResourceRequest) SetJustification(value *string)() {
    m.justification = value
}
// SetRequestor sets the requestor property value. Read-only. Nullable. Supports $expand.
func (m *AccessPackageResourceRequest) SetRequestor(value AccessPackageSubjectable)() {
    m.requestor = value
}
// SetRequestState sets the requestState property value. The outcome of whether the service was able to add the resource to the catalog.  The value is Delivered if the resource was added or removed. Read-Only.
func (m *AccessPackageResourceRequest) SetRequestState(value *string)() {
    m.requestState = value
}
// SetRequestStatus sets the requestStatus property value. The requestStatus property
func (m *AccessPackageResourceRequest) SetRequestStatus(value *string)() {
    m.requestStatus = value
}
// SetRequestType sets the requestType property value. Use AdminAdd to add a resource, if the caller is an administrator or resource owner, AdminUpdate to update a resource, or AdminRemove to remove a resource.
func (m *AccessPackageResourceRequest) SetRequestType(value *string)() {
    m.requestType = value
}
