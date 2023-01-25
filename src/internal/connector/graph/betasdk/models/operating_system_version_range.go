package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperatingSystemVersionRange operating System version range.
type OperatingSystemVersionRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The description of this range (e.g. Valid 1702 builds)
    description *string
    // The highest inclusive version that this range contains.
    highestVersion *string
    // The lowest inclusive version that this range contains.
    lowestVersion *string
    // The OdataType property
    odataType *string
}
// NewOperatingSystemVersionRange instantiates a new operatingSystemVersionRange and sets the default values.
func NewOperatingSystemVersionRange()(*OperatingSystemVersionRange) {
    m := &OperatingSystemVersionRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOperatingSystemVersionRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperatingSystemVersionRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperatingSystemVersionRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OperatingSystemVersionRange) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of this range (e.g. Valid 1702 builds)
func (m *OperatingSystemVersionRange) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OperatingSystemVersionRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["highestVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighestVersion(val)
        }
        return nil
    }
    res["lowestVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowestVersion(val)
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
    return res
}
// GetHighestVersion gets the highestVersion property value. The highest inclusive version that this range contains.
func (m *OperatingSystemVersionRange) GetHighestVersion()(*string) {
    return m.highestVersion
}
// GetLowestVersion gets the lowestVersion property value. The lowest inclusive version that this range contains.
func (m *OperatingSystemVersionRange) GetLowestVersion()(*string) {
    return m.lowestVersion
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OperatingSystemVersionRange) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OperatingSystemVersionRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("highestVersion", m.GetHighestVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lowestVersion", m.GetLowestVersion())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OperatingSystemVersionRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of this range (e.g. Valid 1702 builds)
func (m *OperatingSystemVersionRange) SetDescription(value *string)() {
    m.description = value
}
// SetHighestVersion sets the highestVersion property value. The highest inclusive version that this range contains.
func (m *OperatingSystemVersionRange) SetHighestVersion(value *string)() {
    m.highestVersion = value
}
// SetLowestVersion sets the lowestVersion property value. The lowest inclusive version that this range contains.
func (m *OperatingSystemVersionRange) SetLowestVersion(value *string)() {
    m.lowestVersion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OperatingSystemVersionRange) SetOdataType(value *string)() {
    m.odataType = value
}
