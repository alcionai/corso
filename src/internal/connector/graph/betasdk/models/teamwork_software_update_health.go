package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkSoftwareUpdateHealth 
type TeamworkSoftwareUpdateHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The software update available for the admin agent.
    adminAgentSoftwareUpdateStatus TeamworkSoftwareUpdateStatusable
    // The software update available for the company portal.
    companyPortalSoftwareUpdateStatus TeamworkSoftwareUpdateStatusable
    // The software update available for the firmware.
    firmwareSoftwareUpdateStatus TeamworkSoftwareUpdateStatusable
    // The OdataType property
    odataType *string
    // The software update available for the operating system.
    operatingSystemSoftwareUpdateStatus TeamworkSoftwareUpdateStatusable
    // The software update available for the partner agent.
    partnerAgentSoftwareUpdateStatus TeamworkSoftwareUpdateStatusable
    // The software update available for the Teams client.
    teamsClientSoftwareUpdateStatus TeamworkSoftwareUpdateStatusable
}
// NewTeamworkSoftwareUpdateHealth instantiates a new teamworkSoftwareUpdateHealth and sets the default values.
func NewTeamworkSoftwareUpdateHealth()(*TeamworkSoftwareUpdateHealth) {
    m := &TeamworkSoftwareUpdateHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkSoftwareUpdateHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkSoftwareUpdateHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSoftwareUpdateHealth(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSoftwareUpdateHealth) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdminAgentSoftwareUpdateStatus gets the adminAgentSoftwareUpdateStatus property value. The software update available for the admin agent.
func (m *TeamworkSoftwareUpdateHealth) GetAdminAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    return m.adminAgentSoftwareUpdateStatus
}
// GetCompanyPortalSoftwareUpdateStatus gets the companyPortalSoftwareUpdateStatus property value. The software update available for the company portal.
func (m *TeamworkSoftwareUpdateHealth) GetCompanyPortalSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    return m.companyPortalSoftwareUpdateStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSoftwareUpdateHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adminAgentSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAgentSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["companyPortalSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyPortalSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["firmwareSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
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
    res["operatingSystemSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["partnerAgentSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["teamsClientSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    return res
}
// GetFirmwareSoftwareUpdateStatus gets the firmwareSoftwareUpdateStatus property value. The software update available for the firmware.
func (m *TeamworkSoftwareUpdateHealth) GetFirmwareSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    return m.firmwareSoftwareUpdateStatus
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkSoftwareUpdateHealth) GetOdataType()(*string) {
    return m.odataType
}
// GetOperatingSystemSoftwareUpdateStatus gets the operatingSystemSoftwareUpdateStatus property value. The software update available for the operating system.
func (m *TeamworkSoftwareUpdateHealth) GetOperatingSystemSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    return m.operatingSystemSoftwareUpdateStatus
}
// GetPartnerAgentSoftwareUpdateStatus gets the partnerAgentSoftwareUpdateStatus property value. The software update available for the partner agent.
func (m *TeamworkSoftwareUpdateHealth) GetPartnerAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    return m.partnerAgentSoftwareUpdateStatus
}
// GetTeamsClientSoftwareUpdateStatus gets the teamsClientSoftwareUpdateStatus property value. The software update available for the Teams client.
func (m *TeamworkSoftwareUpdateHealth) GetTeamsClientSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    return m.teamsClientSoftwareUpdateStatus
}
// Serialize serializes information the current object
func (m *TeamworkSoftwareUpdateHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("adminAgentSoftwareUpdateStatus", m.GetAdminAgentSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("companyPortalSoftwareUpdateStatus", m.GetCompanyPortalSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("firmwareSoftwareUpdateStatus", m.GetFirmwareSoftwareUpdateStatus())
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
    {
        err := writer.WriteObjectValue("operatingSystemSoftwareUpdateStatus", m.GetOperatingSystemSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("partnerAgentSoftwareUpdateStatus", m.GetPartnerAgentSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("teamsClientSoftwareUpdateStatus", m.GetTeamsClientSoftwareUpdateStatus())
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
func (m *TeamworkSoftwareUpdateHealth) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdminAgentSoftwareUpdateStatus sets the adminAgentSoftwareUpdateStatus property value. The software update available for the admin agent.
func (m *TeamworkSoftwareUpdateHealth) SetAdminAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    m.adminAgentSoftwareUpdateStatus = value
}
// SetCompanyPortalSoftwareUpdateStatus sets the companyPortalSoftwareUpdateStatus property value. The software update available for the company portal.
func (m *TeamworkSoftwareUpdateHealth) SetCompanyPortalSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    m.companyPortalSoftwareUpdateStatus = value
}
// SetFirmwareSoftwareUpdateStatus sets the firmwareSoftwareUpdateStatus property value. The software update available for the firmware.
func (m *TeamworkSoftwareUpdateHealth) SetFirmwareSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    m.firmwareSoftwareUpdateStatus = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkSoftwareUpdateHealth) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperatingSystemSoftwareUpdateStatus sets the operatingSystemSoftwareUpdateStatus property value. The software update available for the operating system.
func (m *TeamworkSoftwareUpdateHealth) SetOperatingSystemSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    m.operatingSystemSoftwareUpdateStatus = value
}
// SetPartnerAgentSoftwareUpdateStatus sets the partnerAgentSoftwareUpdateStatus property value. The software update available for the partner agent.
func (m *TeamworkSoftwareUpdateHealth) SetPartnerAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    m.partnerAgentSoftwareUpdateStatus = value
}
// SetTeamsClientSoftwareUpdateStatus sets the teamsClientSoftwareUpdateStatus property value. The software update available for the Teams client.
func (m *TeamworkSoftwareUpdateHealth) SetTeamsClientSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    m.teamsClientSoftwareUpdateStatus = value
}
