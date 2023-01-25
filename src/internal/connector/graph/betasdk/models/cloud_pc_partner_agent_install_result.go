package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcPartnerAgentInstallResult 
type CloudPcPartnerAgentInstallResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The status of a partner agent installation. Possible values are: installed, installFailed, installing, uninstalling, uninstallFailed and licensed. Read-Only.
    installStatus *CloudPcPartnerAgentInstallStatus
    // Indicates if the partner agent is a third party. When 'TRUE', the agent is a third-party (non-Microsoft) agent.  When 'FALSE', the agent is a Microsoft agent or is not known.  The default value is 'FALSE'.
    isThirdPartyPartner *bool
    // The OdataType property
    odataType *string
    // Indicates the name of a partner agent and includes first-party and third-party. Currently, Citrix is the only third-party value. Read-Only.
    partnerAgentName *CloudPcPartnerAgentName
    // Indicates if the partner agent is a third party. When 'TRUE', the agent is a third-party (non-Microsoft) agent. When 'FALSE', the agent is a Microsoft agent or is not known. The default value is 'FALSE'.
    retriable *bool
}
// NewCloudPcPartnerAgentInstallResult instantiates a new cloudPcPartnerAgentInstallResult and sets the default values.
func NewCloudPcPartnerAgentInstallResult()(*CloudPcPartnerAgentInstallResult) {
    m := &CloudPcPartnerAgentInstallResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcPartnerAgentInstallResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcPartnerAgentInstallResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcPartnerAgentInstallResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcPartnerAgentInstallResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcPartnerAgentInstallResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["installStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcPartnerAgentInstallStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallStatus(val.(*CloudPcPartnerAgentInstallStatus))
        }
        return nil
    }
    res["isThirdPartyPartner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsThirdPartyPartner(val)
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
    res["partnerAgentName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcPartnerAgentName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentName(val.(*CloudPcPartnerAgentName))
        }
        return nil
    }
    res["retriable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetriable(val)
        }
        return nil
    }
    return res
}
// GetInstallStatus gets the installStatus property value. The status of a partner agent installation. Possible values are: installed, installFailed, installing, uninstalling, uninstallFailed and licensed. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) GetInstallStatus()(*CloudPcPartnerAgentInstallStatus) {
    return m.installStatus
}
// GetIsThirdPartyPartner gets the isThirdPartyPartner property value. Indicates if the partner agent is a third party. When 'TRUE', the agent is a third-party (non-Microsoft) agent.  When 'FALSE', the agent is a Microsoft agent or is not known.  The default value is 'FALSE'.
func (m *CloudPcPartnerAgentInstallResult) GetIsThirdPartyPartner()(*bool) {
    return m.isThirdPartyPartner
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcPartnerAgentInstallResult) GetOdataType()(*string) {
    return m.odataType
}
// GetPartnerAgentName gets the partnerAgentName property value. Indicates the name of a partner agent and includes first-party and third-party. Currently, Citrix is the only third-party value. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) GetPartnerAgentName()(*CloudPcPartnerAgentName) {
    return m.partnerAgentName
}
// GetRetriable gets the retriable property value. Indicates if the partner agent is a third party. When 'TRUE', the agent is a third-party (non-Microsoft) agent. When 'FALSE', the agent is a Microsoft agent or is not known. The default value is 'FALSE'.
func (m *CloudPcPartnerAgentInstallResult) GetRetriable()(*bool) {
    return m.retriable
}
// Serialize serializes information the current object
func (m *CloudPcPartnerAgentInstallResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInstallStatus() != nil {
        cast := (*m.GetInstallStatus()).String()
        err := writer.WriteStringValue("installStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isThirdPartyPartner", m.GetIsThirdPartyPartner())
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
    if m.GetPartnerAgentName() != nil {
        cast := (*m.GetPartnerAgentName()).String()
        err := writer.WriteStringValue("partnerAgentName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("retriable", m.GetRetriable())
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
func (m *CloudPcPartnerAgentInstallResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetInstallStatus sets the installStatus property value. The status of a partner agent installation. Possible values are: installed, installFailed, installing, uninstalling, uninstallFailed and licensed. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) SetInstallStatus(value *CloudPcPartnerAgentInstallStatus)() {
    m.installStatus = value
}
// SetIsThirdPartyPartner sets the isThirdPartyPartner property value. Indicates if the partner agent is a third party. When 'TRUE', the agent is a third-party (non-Microsoft) agent.  When 'FALSE', the agent is a Microsoft agent or is not known.  The default value is 'FALSE'.
func (m *CloudPcPartnerAgentInstallResult) SetIsThirdPartyPartner(value *bool)() {
    m.isThirdPartyPartner = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcPartnerAgentInstallResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPartnerAgentName sets the partnerAgentName property value. Indicates the name of a partner agent and includes first-party and third-party. Currently, Citrix is the only third-party value. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) SetPartnerAgentName(value *CloudPcPartnerAgentName)() {
    m.partnerAgentName = value
}
// SetRetriable sets the retriable property value. Indicates if the partner agent is a third party. When 'TRUE', the agent is a third-party (non-Microsoft) agent. When 'FALSE', the agent is a Microsoft agent or is not known. The default value is 'FALSE'.
func (m *CloudPcPartnerAgentInstallResult) SetRetriable(value *bool)() {
    m.retriable = value
}
