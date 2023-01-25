package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InstitutionData 
type InstitutionData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Short description of the institution the user studied at.
    description *string
    // Name of the institution the user studied at.
    displayName *string
    // Address or location of the institute.
    location PhysicalAddressable
    // The OdataType property
    odataType *string
    // Link to the institution or department homepage.
    webUrl *string
}
// NewInstitutionData instantiates a new institutionData and sets the default values.
func NewInstitutionData()(*InstitutionData) {
    m := &InstitutionData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInstitutionDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInstitutionDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstitutionData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstitutionData) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. Short description of the institution the user studied at.
func (m *InstitutionData) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the institution the user studied at.
func (m *InstitutionData) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InstitutionData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(PhysicalAddressable))
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
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. Address or location of the institute.
func (m *InstitutionData) GetLocation()(PhysicalAddressable) {
    return m.location
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InstitutionData) GetOdataType()(*string) {
    return m.odataType
}
// GetWebUrl gets the webUrl property value. Link to the institution or department homepage.
func (m *InstitutionData) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *InstitutionData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("location", m.GetLocation())
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
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *InstitutionData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. Short description of the institution the user studied at.
func (m *InstitutionData) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the institution the user studied at.
func (m *InstitutionData) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLocation sets the location property value. Address or location of the institute.
func (m *InstitutionData) SetLocation(value PhysicalAddressable)() {
    m.location = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InstitutionData) SetOdataType(value *string)() {
    m.odataType = value
}
// SetWebUrl sets the webUrl property value. Link to the institution or department homepage.
func (m *InstitutionData) SetWebUrl(value *string)() {
    m.webUrl = value
}
