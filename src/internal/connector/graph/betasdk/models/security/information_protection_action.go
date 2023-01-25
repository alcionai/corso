package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionAction 
type InformationProtectionAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
}
// NewInformationProtectionAction instantiates a new informationProtectionAction and sets the default values.
func NewInformationProtectionAction()(*InformationProtectionAction) {
    m := &InformationProtectionAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInformationProtectionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.security.addContentFooterAction":
                        return NewAddContentFooterAction(), nil
                    case "#microsoft.graph.security.addContentHeaderAction":
                        return NewAddContentHeaderAction(), nil
                    case "#microsoft.graph.security.addWatermarkAction":
                        return NewAddWatermarkAction(), nil
                    case "#microsoft.graph.security.applyLabelAction":
                        return NewApplyLabelAction(), nil
                    case "#microsoft.graph.security.customAction":
                        return NewCustomAction(), nil
                    case "#microsoft.graph.security.justifyAction":
                        return NewJustifyAction(), nil
                    case "#microsoft.graph.security.metadataAction":
                        return NewMetadataAction(), nil
                    case "#microsoft.graph.security.protectAdhocAction":
                        return NewProtectAdhocAction(), nil
                    case "#microsoft.graph.security.protectByTemplateAction":
                        return NewProtectByTemplateAction(), nil
                    case "#microsoft.graph.security.protectDoNotForwardAction":
                        return NewProtectDoNotForwardAction(), nil
                    case "#microsoft.graph.security.recommendLabelAction":
                        return NewRecommendLabelAction(), nil
                    case "#microsoft.graph.security.removeContentFooterAction":
                        return NewRemoveContentFooterAction(), nil
                    case "#microsoft.graph.security.removeContentHeaderAction":
                        return NewRemoveContentHeaderAction(), nil
                    case "#microsoft.graph.security.removeProtectionAction":
                        return NewRemoveProtectionAction(), nil
                    case "#microsoft.graph.security.removeWatermarkAction":
                        return NewRemoveWatermarkAction(), nil
                }
            }
        }
    }
    return NewInformationProtectionAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *InformationProtectionAction) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *InformationProtectionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *InformationProtectionAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InformationProtectionAction) SetOdataType(value *string)() {
    m.odataType = value
}
