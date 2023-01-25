package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedDeviceIdentity the importedDeviceIdentity resource represents a unique hardware identity of a device that has been pre-staged for pre-enrollment configuration.
type ImportedDeviceIdentity struct {
    Entity
    // Created Date Time of the device
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of the device
    description *string
    // The enrollmentState property
    enrollmentState *EnrollmentState
    // Imported Device Identifier
    importedDeviceIdentifier *string
    // The importedDeviceIdentityType property
    importedDeviceIdentityType *ImportedDeviceIdentityType
    // Last Contacted Date Time of the device
    lastContactedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last Modified DateTime of the description
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The platform property
    platform *Platform
}
// NewImportedDeviceIdentity instantiates a new importedDeviceIdentity and sets the default values.
func NewImportedDeviceIdentity()(*ImportedDeviceIdentity) {
    m := &ImportedDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateImportedDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.importedDeviceIdentityResult":
                        return NewImportedDeviceIdentityResult(), nil
                }
            }
        }
    }
    return NewImportedDeviceIdentity(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Created Date Time of the device
func (m *ImportedDeviceIdentity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of the device
func (m *ImportedDeviceIdentity) GetDescription()(*string) {
    return m.description
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
func (m *ImportedDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    return m.enrollmentState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedDeviceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["enrollmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["importedDeviceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportedDeviceIdentifier(val)
        }
        return nil
    }
    res["importedDeviceIdentityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportedDeviceIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportedDeviceIdentityType(val.(*ImportedDeviceIdentityType))
        }
        return nil
    }
    res["lastContactedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastContactedDateTime(val)
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*Platform))
        }
        return nil
    }
    return res
}
// GetImportedDeviceIdentifier gets the importedDeviceIdentifier property value. Imported Device Identifier
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentifier()(*string) {
    return m.importedDeviceIdentifier
}
// GetImportedDeviceIdentityType gets the importedDeviceIdentityType property value. The importedDeviceIdentityType property
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentityType()(*ImportedDeviceIdentityType) {
    return m.importedDeviceIdentityType
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastContactedDateTime
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last Modified DateTime of the description
func (m *ImportedDeviceIdentity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPlatform gets the platform property value. The platform property
func (m *ImportedDeviceIdentity) GetPlatform()(*Platform) {
    return m.platform
}
// Serialize serializes information the current object
func (m *ImportedDeviceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("importedDeviceIdentifier", m.GetImportedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetImportedDeviceIdentityType() != nil {
        cast := (*m.GetImportedDeviceIdentityType()).String()
        err = writer.WriteStringValue("importedDeviceIdentityType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastContactedDateTime", m.GetLastContactedDateTime())
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created Date Time of the device
func (m *ImportedDeviceIdentity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the device
func (m *ImportedDeviceIdentity) SetDescription(value *string)() {
    m.description = value
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *ImportedDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
// SetImportedDeviceIdentifier sets the importedDeviceIdentifier property value. Imported Device Identifier
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentifier(value *string)() {
    m.importedDeviceIdentifier = value
}
// SetImportedDeviceIdentityType sets the importedDeviceIdentityType property value. The importedDeviceIdentityType property
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentityType(value *ImportedDeviceIdentityType)() {
    m.importedDeviceIdentityType = value
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastContactedDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last Modified DateTime of the description
func (m *ImportedDeviceIdentity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPlatform sets the platform property value. The platform property
func (m *ImportedDeviceIdentity) SetPlatform(value *Platform)() {
    m.platform = value
}
