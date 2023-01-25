package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceSoftwareVersions 
type TeamworkDeviceSoftwareVersions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The software version for the admin agent running on the device.
    adminAgentSoftwareVersion *string
    // The software version for the firmware running on the device.
    firmwareSoftwareVersion *string
    // The OdataType property
    odataType *string
    // The software version for the operating system on the device.
    operatingSystemSoftwareVersion *string
    // The software version for the partner agent running on the device.
    partnerAgentSoftwareVersion *string
    // The software version for the Teams client running on the device.
    teamsClientSoftwareVersion *string
}
// NewTeamworkDeviceSoftwareVersions instantiates a new teamworkDeviceSoftwareVersions and sets the default values.
func NewTeamworkDeviceSoftwareVersions()(*TeamworkDeviceSoftwareVersions) {
    m := &TeamworkDeviceSoftwareVersions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkDeviceSoftwareVersionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDeviceSoftwareVersionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDeviceSoftwareVersions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDeviceSoftwareVersions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdminAgentSoftwareVersion gets the adminAgentSoftwareVersion property value. The software version for the admin agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetAdminAgentSoftwareVersion()(*string) {
    return m.adminAgentSoftwareVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceSoftwareVersions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adminAgentSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAgentSoftwareVersion(val)
        }
        return nil
    }
    res["firmwareSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareSoftwareVersion(val)
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
    res["operatingSystemSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemSoftwareVersion(val)
        }
        return nil
    }
    res["partnerAgentSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentSoftwareVersion(val)
        }
        return nil
    }
    res["teamsClientSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientSoftwareVersion(val)
        }
        return nil
    }
    return res
}
// GetFirmwareSoftwareVersion gets the firmwareSoftwareVersion property value. The software version for the firmware running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetFirmwareSoftwareVersion()(*string) {
    return m.firmwareSoftwareVersion
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkDeviceSoftwareVersions) GetOdataType()(*string) {
    return m.odataType
}
// GetOperatingSystemSoftwareVersion gets the operatingSystemSoftwareVersion property value. The software version for the operating system on the device.
func (m *TeamworkDeviceSoftwareVersions) GetOperatingSystemSoftwareVersion()(*string) {
    return m.operatingSystemSoftwareVersion
}
// GetPartnerAgentSoftwareVersion gets the partnerAgentSoftwareVersion property value. The software version for the partner agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetPartnerAgentSoftwareVersion()(*string) {
    return m.partnerAgentSoftwareVersion
}
// GetTeamsClientSoftwareVersion gets the teamsClientSoftwareVersion property value. The software version for the Teams client running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetTeamsClientSoftwareVersion()(*string) {
    return m.teamsClientSoftwareVersion
}
// Serialize serializes information the current object
func (m *TeamworkDeviceSoftwareVersions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adminAgentSoftwareVersion", m.GetAdminAgentSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareSoftwareVersion", m.GetFirmwareSoftwareVersion())
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
        err := writer.WriteStringValue("operatingSystemSoftwareVersion", m.GetOperatingSystemSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("partnerAgentSoftwareVersion", m.GetPartnerAgentSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamsClientSoftwareVersion", m.GetTeamsClientSoftwareVersion())
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
func (m *TeamworkDeviceSoftwareVersions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdminAgentSoftwareVersion sets the adminAgentSoftwareVersion property value. The software version for the admin agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetAdminAgentSoftwareVersion(value *string)() {
    m.adminAgentSoftwareVersion = value
}
// SetFirmwareSoftwareVersion sets the firmwareSoftwareVersion property value. The software version for the firmware running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetFirmwareSoftwareVersion(value *string)() {
    m.firmwareSoftwareVersion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDeviceSoftwareVersions) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperatingSystemSoftwareVersion sets the operatingSystemSoftwareVersion property value. The software version for the operating system on the device.
func (m *TeamworkDeviceSoftwareVersions) SetOperatingSystemSoftwareVersion(value *string)() {
    m.operatingSystemSoftwareVersion = value
}
// SetPartnerAgentSoftwareVersion sets the partnerAgentSoftwareVersion property value. The software version for the partner agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetPartnerAgentSoftwareVersion(value *string)() {
    m.partnerAgentSoftwareVersion = value
}
// SetTeamsClientSoftwareVersion sets the teamsClientSoftwareVersion property value. The software version for the Teams client running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetTeamsClientSoftwareVersion(value *string)() {
    m.teamsClientSoftwareVersion = value
}
