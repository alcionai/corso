package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SingleSignOnExtension represents an Apple Single Sign-On Extension.
type SingleSignOnExtension struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
}
// NewSingleSignOnExtension instantiates a new singleSignOnExtension and sets the default values.
func NewSingleSignOnExtension()(*SingleSignOnExtension) {
    m := &SingleSignOnExtension{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.credentialSingleSignOnExtension":
                        return NewCredentialSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosAzureAdSingleSignOnExtension":
                        return NewIosAzureAdSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosCredentialSingleSignOnExtension":
                        return NewIosCredentialSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosKerberosSingleSignOnExtension":
                        return NewIosKerberosSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosRedirectSingleSignOnExtension":
                        return NewIosRedirectSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosSingleSignOnExtension":
                        return NewIosSingleSignOnExtension(), nil
                    case "#microsoft.graph.kerberosSingleSignOnExtension":
                        return NewKerberosSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSAzureAdSingleSignOnExtension":
                        return NewMacOSAzureAdSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSCredentialSingleSignOnExtension":
                        return NewMacOSCredentialSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSKerberosSingleSignOnExtension":
                        return NewMacOSKerberosSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSRedirectSingleSignOnExtension":
                        return NewMacOSRedirectSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSSingleSignOnExtension":
                        return NewMacOSSingleSignOnExtension(), nil
                    case "#microsoft.graph.redirectSingleSignOnExtension":
                        return NewRedirectSingleSignOnExtension(), nil
                }
            }
        }
    }
    return NewSingleSignOnExtension(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SingleSignOnExtension) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SingleSignOnExtension) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SingleSignOnExtension) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SingleSignOnExtension) SetOdataType(value *string)() {
    m.odataType = value
}
