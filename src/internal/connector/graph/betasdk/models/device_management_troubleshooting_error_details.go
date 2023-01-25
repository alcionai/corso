package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTroubleshootingErrorDetails object containing detailed information about the error and its remediation.
type DeviceManagementTroubleshootingErrorDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Not yet documented
    context *string
    // Not yet documented
    failure *string
    // The detailed description of what went wrong.
    failureDetails *string
    // The OdataType property
    odataType *string
    // The detailed description of how to remediate this issue.
    remediation *string
    // Links to helpful documentation about this failure.
    resources []DeviceManagementTroubleshootingErrorResourceable
}
// NewDeviceManagementTroubleshootingErrorDetails instantiates a new deviceManagementTroubleshootingErrorDetails and sets the default values.
func NewDeviceManagementTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    m := &DeviceManagementTroubleshootingErrorDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementTroubleshootingErrorDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContext gets the context property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetContext()(*string) {
    return m.context
}
// GetFailure gets the failure property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailure()(*string) {
    return m.failure
}
// GetFailureDetails gets the failureDetails property value. The detailed description of what went wrong.
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailureDetails()(*string) {
    return m.failureDetails
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["failure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailure(val)
        }
        return nil
    }
    res["failureDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureDetails(val)
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
    res["remediation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediation(val)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingErrorResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingErrorResourceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementTroubleshootingErrorResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementTroubleshootingErrorDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetRemediation gets the remediation property value. The detailed description of how to remediate this issue.
func (m *DeviceManagementTroubleshootingErrorDetails) GetRemediation()(*string) {
    return m.remediation
}
// GetResources gets the resources property value. Links to helpful documentation about this failure.
func (m *DeviceManagementTroubleshootingErrorDetails) GetResources()([]DeviceManagementTroubleshootingErrorResourceable) {
    return m.resources
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failure", m.GetFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failureDetails", m.GetFailureDetails())
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
        err := writer.WriteStringValue("remediation", m.GetRemediation())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("resources", cast)
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
func (m *DeviceManagementTroubleshootingErrorDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContext sets the context property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) SetContext(value *string)() {
    m.context = value
}
// SetFailure sets the failure property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailure(value *string)() {
    m.failure = value
}
// SetFailureDetails sets the failureDetails property value. The detailed description of what went wrong.
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailureDetails(value *string)() {
    m.failureDetails = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementTroubleshootingErrorDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemediation sets the remediation property value. The detailed description of how to remediate this issue.
func (m *DeviceManagementTroubleshootingErrorDetails) SetRemediation(value *string)() {
    m.remediation = value
}
// SetResources sets the resources property value. Links to helpful documentation about this failure.
func (m *DeviceManagementTroubleshootingErrorDetails) SetResources(value []DeviceManagementTroubleshootingErrorResourceable)() {
    m.resources = value
}
