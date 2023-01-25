package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScript intune will provide customer the ability to run their Powershell Compliance scripts (detection) on the enrolled windows 10 Azure Active Directory joined devices.
type DeviceComplianceScript struct {
    Entity
    // The list of group assignments for the device compliance script
    assignments []DeviceHealthScriptAssignmentable
    // The timestamp of when the device compliance script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the device compliance script
    description *string
    // The entire content of the detection powershell script
    detectionScriptContent []byte
    // List of run states for the device compliance script across all devices
    deviceRunStates []DeviceComplianceScriptDeviceStateable
    // Name of the device compliance script
    displayName *string
    // Indicate whether the script signature needs be checked
    enforceSignatureCheck *bool
    // The timestamp of when the device compliance script was modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the device compliance script publisher
    publisher *string
    // List of Scope Tag IDs for the device compliance script
    roleScopeTagIds []string
    // Indicate whether PowerShell script(s) should run as 32-bit
    runAs32Bit *bool
    // Indicates the type of execution context the app runs in.
    runAsAccount *RunAsAccountType
    // High level run summary for device compliance script.
    runSummary DeviceComplianceScriptRunSummaryable
    // Version of the device compliance script
    version *string
}
// NewDeviceComplianceScript instantiates a new deviceComplianceScript and sets the default values.
func NewDeviceComplianceScript()(*DeviceComplianceScript) {
    m := &DeviceComplianceScript{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceComplianceScriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceScript(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the device compliance script
func (m *DeviceComplianceScript) GetAssignments()([]DeviceHealthScriptAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The timestamp of when the device compliance script was created. This property is read-only.
func (m *DeviceComplianceScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the device compliance script
func (m *DeviceComplianceScript) GetDescription()(*string) {
    return m.description
}
// GetDetectionScriptContent gets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceComplianceScript) GetDetectionScriptContent()([]byte) {
    return m.detectionScriptContent
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for the device compliance script across all devices
func (m *DeviceComplianceScript) GetDeviceRunStates()([]DeviceComplianceScriptDeviceStateable) {
    return m.deviceRunStates
}
// GetDisplayName gets the displayName property value. Name of the device compliance script
func (m *DeviceComplianceScript) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceComplianceScript) GetEnforceSignatureCheck()(*bool) {
    return m.enforceSignatureCheck
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScript) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceHealthScriptAssignmentable)
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
    res["detectionScriptContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptContent(val)
        }
        return nil
    }
    res["deviceRunStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptDeviceStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptDeviceStateable)
            }
            m.SetDeviceRunStates(res)
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
    res["enforceSignatureCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceSignatureCheck(val)
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
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
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
    res["runAs32Bit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAs32Bit(val)
        }
        return nil
    }
    res["runAsAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    res["runSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceComplianceScriptRunSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(DeviceComplianceScriptRunSummaryable))
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The timestamp of when the device compliance script was modified. This property is read-only.
func (m *DeviceComplianceScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPublisher gets the publisher property value. Name of the device compliance script publisher
func (m *DeviceComplianceScript) GetPublisher()(*string) {
    return m.publisher
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for the device compliance script
func (m *DeviceComplianceScript) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetRunAs32Bit gets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceComplianceScript) GetRunAs32Bit()(*bool) {
    return m.runAs32Bit
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceComplianceScript) GetRunAsAccount()(*RunAsAccountType) {
    return m.runAsAccount
}
// GetRunSummary gets the runSummary property value. High level run summary for device compliance script.
func (m *DeviceComplianceScript) GetRunSummary()(DeviceComplianceScriptRunSummaryable) {
    return m.runSummary
}
// GetVersion gets the version property value. Version of the device compliance script
func (m *DeviceComplianceScript) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceComplianceScript) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("detectionScriptContent", m.GetDetectionScriptContent())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRunStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
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
        err = writer.WriteBoolValue("enforceSignatureCheck", m.GetEnforceSignatureCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
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
        err = writer.WriteBoolValue("runAs32Bit", m.GetRunAs32Bit())
        if err != nil {
            return err
        }
    }
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err = writer.WriteStringValue("runAsAccount", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("runSummary", m.GetRunSummary())
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
// SetAssignments sets the assignments property value. The list of group assignments for the device compliance script
func (m *DeviceComplianceScript) SetAssignments(value []DeviceHealthScriptAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The timestamp of when the device compliance script was created. This property is read-only.
func (m *DeviceComplianceScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the device compliance script
func (m *DeviceComplianceScript) SetDescription(value *string)() {
    m.description = value
}
// SetDetectionScriptContent sets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceComplianceScript) SetDetectionScriptContent(value []byte)() {
    m.detectionScriptContent = value
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for the device compliance script across all devices
func (m *DeviceComplianceScript) SetDeviceRunStates(value []DeviceComplianceScriptDeviceStateable)() {
    m.deviceRunStates = value
}
// SetDisplayName sets the displayName property value. Name of the device compliance script
func (m *DeviceComplianceScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceComplianceScript) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The timestamp of when the device compliance script was modified. This property is read-only.
func (m *DeviceComplianceScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPublisher sets the publisher property value. Name of the device compliance script publisher
func (m *DeviceComplianceScript) SetPublisher(value *string)() {
    m.publisher = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for the device compliance script
func (m *DeviceComplianceScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetRunAs32Bit sets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceComplianceScript) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceComplianceScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// SetRunSummary sets the runSummary property value. High level run summary for device compliance script.
func (m *DeviceComplianceScript) SetRunSummary(value DeviceComplianceScriptRunSummaryable)() {
    m.runSummary = value
}
// SetVersion sets the version property value. Version of the device compliance script
func (m *DeviceComplianceScript) SetVersion(value *string)() {
    m.version = value
}
