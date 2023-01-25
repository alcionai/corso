package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KerberosSignOnSettings 
type KerberosSignOnSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
    kerberosServicePrincipalName *string
    // The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
    kerberosSignOnMappingAttributeType *KerberosSignOnMappingAttributeType
    // The OdataType property
    odataType *string
}
// NewKerberosSignOnSettings instantiates a new kerberosSignOnSettings and sets the default values.
func NewKerberosSignOnSettings()(*KerberosSignOnSettings) {
    m := &KerberosSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateKerberosSignOnSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKerberosSignOnSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKerberosSignOnSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KerberosSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KerberosSignOnSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["kerberosServicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosServicePrincipalName(val)
        }
        return nil
    }
    res["kerberosSignOnMappingAttributeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKerberosSignOnMappingAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnMappingAttributeType(val.(*KerberosSignOnMappingAttributeType))
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
// GetKerberosServicePrincipalName gets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) GetKerberosServicePrincipalName()(*string) {
    return m.kerberosServicePrincipalName
}
// GetKerberosSignOnMappingAttributeType gets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType) {
    return m.kerberosSignOnMappingAttributeType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *KerberosSignOnSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *KerberosSignOnSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("kerberosServicePrincipalName", m.GetKerberosServicePrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetKerberosSignOnMappingAttributeType() != nil {
        cast := (*m.GetKerberosSignOnMappingAttributeType()).String()
        err := writer.WriteStringValue("kerberosSignOnMappingAttributeType", &cast)
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
func (m *KerberosSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKerberosServicePrincipalName sets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) SetKerberosServicePrincipalName(value *string)() {
    m.kerberosServicePrincipalName = value
}
// SetKerberosSignOnMappingAttributeType sets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)() {
    m.kerberosSignOnMappingAttributeType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *KerberosSignOnSettings) SetOdataType(value *string)() {
    m.odataType = value
}
