package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSystemExtensionTypeMapping represents a mapping between team identifiers for macOS system extensions and system extension types.
type MacOSSystemExtensionTypeMapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Flag enum representing the allowed macOS system extension types.
    allowedTypes *MacOSSystemExtensionType
    // The OdataType property
    odataType *string
    // Gets or sets the team identifier used to sign the system extension.
    teamIdentifier *string
}
// NewMacOSSystemExtensionTypeMapping instantiates a new macOSSystemExtensionTypeMapping and sets the default values.
func NewMacOSSystemExtensionTypeMapping()(*MacOSSystemExtensionTypeMapping) {
    m := &MacOSSystemExtensionTypeMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSSystemExtensionTypeMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSystemExtensionTypeMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSystemExtensionTypeMapping(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSSystemExtensionTypeMapping) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedTypes gets the allowedTypes property value. Flag enum representing the allowed macOS system extension types.
func (m *MacOSSystemExtensionTypeMapping) GetAllowedTypes()(*MacOSSystemExtensionType) {
    return m.allowedTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSystemExtensionTypeMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSystemExtensionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedTypes(val.(*MacOSSystemExtensionType))
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
    res["teamIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamIdentifier(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSSystemExtensionTypeMapping) GetOdataType()(*string) {
    return m.odataType
}
// GetTeamIdentifier gets the teamIdentifier property value. Gets or sets the team identifier used to sign the system extension.
func (m *MacOSSystemExtensionTypeMapping) GetTeamIdentifier()(*string) {
    return m.teamIdentifier
}
// Serialize serializes information the current object
func (m *MacOSSystemExtensionTypeMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedTypes() != nil {
        cast := (*m.GetAllowedTypes()).String()
        err := writer.WriteStringValue("allowedTypes", &cast)
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
        err := writer.WriteStringValue("teamIdentifier", m.GetTeamIdentifier())
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
func (m *MacOSSystemExtensionTypeMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedTypes sets the allowedTypes property value. Flag enum representing the allowed macOS system extension types.
func (m *MacOSSystemExtensionTypeMapping) SetAllowedTypes(value *MacOSSystemExtensionType)() {
    m.allowedTypes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSSystemExtensionTypeMapping) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTeamIdentifier sets the teamIdentifier property value. Gets or sets the team identifier used to sign the system extension.
func (m *MacOSSystemExtensionTypeMapping) SetTeamIdentifier(value *string)() {
    m.teamIdentifier = value
}
