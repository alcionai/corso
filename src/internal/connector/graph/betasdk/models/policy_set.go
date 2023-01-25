package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicySet a class containing the properties used for PolicySet.
type PolicySet struct {
    Entity
    // Assignments of the PolicySet.
    assignments []PolicySetAssignmentable
    // Creation time of the PolicySet.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the PolicySet.
    description *string
    // DisplayName of the PolicySet.
    displayName *string
    // The errorCode property
    errorCode *ErrorCode
    // Tags of the guided deployment
    guidedDeploymentTags []string
    // Items of the PolicySet with maximum count 100.
    items []PolicySetItemable
    // Last modified time of the PolicySet.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // RoleScopeTags of the PolicySet
    roleScopeTags []string
    // The enum to specify the status of PolicySet.
    status *PolicySetStatus
}
// NewPolicySet instantiates a new policySet and sets the default values.
func NewPolicySet()(*PolicySet) {
    m := &PolicySet{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePolicySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicySet(), nil
}
// GetAssignments gets the assignments property value. Assignments of the PolicySet.
func (m *PolicySet) GetAssignments()([]PolicySetAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. Creation time of the PolicySet.
func (m *PolicySet) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the PolicySet.
func (m *PolicySet) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. DisplayName of the PolicySet.
func (m *PolicySet) GetDisplayName()(*string) {
    return m.displayName
}
// GetErrorCode gets the errorCode property value. The errorCode property
func (m *PolicySet) GetErrorCode()(*ErrorCode) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePolicySetAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySetAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(PolicySetAssignmentable)
            }
            m.SetAssignments(res)
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
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseErrorCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val.(*ErrorCode))
        }
        return nil
    }
    res["guidedDeploymentTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGuidedDeploymentTags(res)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(PolicySetItemable)
            }
            m.SetItems(res)
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
    res["roleScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTags(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicySetStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*PolicySetStatus))
        }
        return nil
    }
    return res
}
// GetGuidedDeploymentTags gets the guidedDeploymentTags property value. Tags of the guided deployment
func (m *PolicySet) GetGuidedDeploymentTags()([]string) {
    return m.guidedDeploymentTags
}
// GetItems gets the items property value. Items of the PolicySet with maximum count 100.
func (m *PolicySet) GetItems()([]PolicySetItemable) {
    return m.items
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified time of the PolicySet.
func (m *PolicySet) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTags gets the roleScopeTags property value. RoleScopeTags of the PolicySet
func (m *PolicySet) GetRoleScopeTags()([]string) {
    return m.roleScopeTags
}
// GetStatus gets the status property value. The enum to specify the status of PolicySet.
func (m *PolicySet) GetStatus()(*PolicySetStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *PolicySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetErrorCode() != nil {
        cast := (*m.GetErrorCode()).String()
        err = writer.WriteStringValue("errorCode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetGuidedDeploymentTags() != nil {
        err = writer.WriteCollectionOfStringValues("guidedDeploymentTags", m.GetGuidedDeploymentTags())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
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
    if m.GetRoleScopeTags() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTags", m.GetRoleScopeTags())
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
// SetAssignments sets the assignments property value. Assignments of the PolicySet.
func (m *PolicySet) SetAssignments(value []PolicySetAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. Creation time of the PolicySet.
func (m *PolicySet) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the PolicySet.
func (m *PolicySet) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. DisplayName of the PolicySet.
func (m *PolicySet) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetErrorCode sets the errorCode property value. The errorCode property
func (m *PolicySet) SetErrorCode(value *ErrorCode)() {
    m.errorCode = value
}
// SetGuidedDeploymentTags sets the guidedDeploymentTags property value. Tags of the guided deployment
func (m *PolicySet) SetGuidedDeploymentTags(value []string)() {
    m.guidedDeploymentTags = value
}
// SetItems sets the items property value. Items of the PolicySet with maximum count 100.
func (m *PolicySet) SetItems(value []PolicySetItemable)() {
    m.items = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified time of the PolicySet.
func (m *PolicySet) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTags sets the roleScopeTags property value. RoleScopeTags of the PolicySet
func (m *PolicySet) SetRoleScopeTags(value []string)() {
    m.roleScopeTags = value
}
// SetStatus sets the status property value. The enum to specify the status of PolicySet.
func (m *PolicySet) SetStatus(value *PolicySetStatus)() {
    m.status = value
}
