package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicy provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsDefenderApplicationControlSupplementalPolicy struct {
    Entity
    // The associated group assignments for this WindowsDefenderApplicationControl supplemental policy.
    assignments []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable
    // The WindowsDefenderApplicationControl supplemental policy content in byte array format.
    content []byte
    // The WindowsDefenderApplicationControl supplemental policy content's file name.
    contentFileName *string
    // The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // WindowsDefenderApplicationControl supplemental policy deployment summary.
    deploySummary WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable
    // The description of WindowsDefenderApplicationControl supplemental policy.
    description *string
    // The list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
    deviceStatuses []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable
    // The display name of WindowsDefenderApplicationControl supplemental policy.
    displayName *string
    // The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
    roleScopeTagIds []string
    // The WindowsDefenderApplicationControl supplemental policy's version.
    version *string
}
// NewWindowsDefenderApplicationControlSupplementalPolicy instantiates a new windowsDefenderApplicationControlSupplementalPolicy and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicy()(*WindowsDefenderApplicationControlSupplementalPolicy) {
    m := &WindowsDefenderApplicationControlSupplementalPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDefenderApplicationControlSupplementalPolicy(), nil
}
// GetAssignments gets the assignments property value. The associated group assignments for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetAssignments()([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable) {
    return m.assignments
}
// GetContent gets the content property value. The WindowsDefenderApplicationControl supplemental policy content in byte array format.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetContent()([]byte) {
    return m.content
}
// GetContentFileName gets the contentFileName property value. The WindowsDefenderApplicationControl supplemental policy content's file name.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetContentFileName()(*string) {
    return m.contentFileName
}
// GetCreationDateTime gets the creationDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetDeploySummary gets the deploySummary property value. WindowsDefenderApplicationControl supplemental policy deployment summary.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDeploySummary()(WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable) {
    return m.deploySummary
}
// GetDescription gets the description property value. The description of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDescription()(*string) {
    return m.description
}
// GetDeviceStatuses gets the deviceStatuses property value. The list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDeviceStatuses()([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable) {
    return m.deviceStatuses
}
// GetDisplayName gets the displayName property value. The display name of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDefenderApplicationControlSupplementalPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentFileName(val)
        }
        return nil
    }
    res["creationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
        }
        return nil
    }
    res["deploySummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploySummary(val.(WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable))
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
    res["deviceStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable)
            }
            m.SetDeviceStatuses(res)
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
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetVersion gets the version property value. The WindowsDefenderApplicationControl supplemental policy's version.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentFileName", m.GetContentFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploySummary", m.GetDeploySummary())
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
    if m.GetDeviceStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The associated group assignments for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetAssignments(value []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)() {
    m.assignments = value
}
// SetContent sets the content property value. The WindowsDefenderApplicationControl supplemental policy content in byte array format.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetContent(value []byte)() {
    m.content = value
}
// SetContentFileName sets the contentFileName property value. The WindowsDefenderApplicationControl supplemental policy content's file name.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetContentFileName(value *string)() {
    m.contentFileName = value
}
// SetCreationDateTime sets the creationDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetDeploySummary sets the deploySummary property value. WindowsDefenderApplicationControl supplemental policy deployment summary.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDeploySummary(value WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable)() {
    m.deploySummary = value
}
// SetDescription sets the description property value. The description of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceStatuses sets the deviceStatuses property value. The list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDeviceStatuses(value []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable)() {
    m.deviceStatuses = value
}
// SetDisplayName sets the displayName property value. The display name of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetVersion sets the version property value. The WindowsDefenderApplicationControl supplemental policy's version.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetVersion(value *string)() {
    m.version = value
}
