package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicySetItem a class containing the properties used for PolicySet Item.
type PolicySetItem struct {
    Entity
    // Creation time of the PolicySetItem.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // DisplayName of the PolicySetItem.
    displayName *string
    // The errorCode property
    errorCode *ErrorCode
    // Tags of the guided deployment
    guidedDeploymentTags []string
    // policySetType of the PolicySetItem.
    itemType *string
    // Last modified time of the PolicySetItem.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // PayloadId of the PolicySetItem.
    payloadId *string
    // The enum to specify the status of PolicySet.
    status *PolicySetStatus
}
// NewPolicySetItem instantiates a new policySetItem and sets the default values.
func NewPolicySetItem()(*PolicySetItem) {
    m := &PolicySetItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceCompliancePolicyPolicySetItem":
                        return NewDeviceCompliancePolicyPolicySetItem(), nil
                    case "#microsoft.graph.deviceConfigurationPolicySetItem":
                        return NewDeviceConfigurationPolicySetItem(), nil
                    case "#microsoft.graph.deviceManagementConfigurationPolicyPolicySetItem":
                        return NewDeviceManagementConfigurationPolicyPolicySetItem(), nil
                    case "#microsoft.graph.deviceManagementScriptPolicySetItem":
                        return NewDeviceManagementScriptPolicySetItem(), nil
                    case "#microsoft.graph.enrollmentRestrictionsConfigurationPolicySetItem":
                        return NewEnrollmentRestrictionsConfigurationPolicySetItem(), nil
                    case "#microsoft.graph.iosLobAppProvisioningConfigurationPolicySetItem":
                        return NewIosLobAppProvisioningConfigurationPolicySetItem(), nil
                    case "#microsoft.graph.managedAppProtectionPolicySetItem":
                        return NewManagedAppProtectionPolicySetItem(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationPolicySetItem":
                        return NewManagedDeviceMobileAppConfigurationPolicySetItem(), nil
                    case "#microsoft.graph.mdmWindowsInformationProtectionPolicyPolicySetItem":
                        return NewMdmWindowsInformationProtectionPolicyPolicySetItem(), nil
                    case "#microsoft.graph.mobileAppPolicySetItem":
                        return NewMobileAppPolicySetItem(), nil
                    case "#microsoft.graph.targetedManagedAppConfigurationPolicySetItem":
                        return NewTargetedManagedAppConfigurationPolicySetItem(), nil
                    case "#microsoft.graph.windows10EnrollmentCompletionPageConfigurationPolicySetItem":
                        return NewWindows10EnrollmentCompletionPageConfigurationPolicySetItem(), nil
                    case "#microsoft.graph.windowsAutopilotDeploymentProfilePolicySetItem":
                        return NewWindowsAutopilotDeploymentProfilePolicySetItem(), nil
                }
            }
        }
    }
    return NewPolicySetItem(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Creation time of the PolicySetItem.
func (m *PolicySetItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. DisplayName of the PolicySetItem.
func (m *PolicySetItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetErrorCode gets the errorCode property value. The errorCode property
func (m *PolicySetItem) GetErrorCode()(*ErrorCode) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["itemType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemType(val)
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
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
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
func (m *PolicySetItem) GetGuidedDeploymentTags()([]string) {
    return m.guidedDeploymentTags
}
// GetItemType gets the itemType property value. policySetType of the PolicySetItem.
func (m *PolicySetItem) GetItemType()(*string) {
    return m.itemType
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified time of the PolicySetItem.
func (m *PolicySetItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPayloadId gets the payloadId property value. PayloadId of the PolicySetItem.
func (m *PolicySetItem) GetPayloadId()(*string) {
    return m.payloadId
}
// GetStatus gets the status property value. The enum to specify the status of PolicySet.
func (m *PolicySetItem) GetStatus()(*PolicySetStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *PolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    {
        err = writer.WriteStringValue("itemType", m.GetItemType())
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
        err = writer.WriteStringValue("payloadId", m.GetPayloadId())
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
// SetCreatedDateTime sets the createdDateTime property value. Creation time of the PolicySetItem.
func (m *PolicySetItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. DisplayName of the PolicySetItem.
func (m *PolicySetItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetErrorCode sets the errorCode property value. The errorCode property
func (m *PolicySetItem) SetErrorCode(value *ErrorCode)() {
    m.errorCode = value
}
// SetGuidedDeploymentTags sets the guidedDeploymentTags property value. Tags of the guided deployment
func (m *PolicySetItem) SetGuidedDeploymentTags(value []string)() {
    m.guidedDeploymentTags = value
}
// SetItemType sets the itemType property value. policySetType of the PolicySetItem.
func (m *PolicySetItem) SetItemType(value *string)() {
    m.itemType = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified time of the PolicySetItem.
func (m *PolicySetItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPayloadId sets the payloadId property value. PayloadId of the PolicySetItem.
func (m *PolicySetItem) SetPayloadId(value *string)() {
    m.payloadId = value
}
// SetStatus sets the status property value. The enum to specify the status of PolicySet.
func (m *PolicySetItem) SetStatus(value *PolicySetStatus)() {
    m.status = value
}
