package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MembershipRuleProcessingStatus 
type MembershipRuleProcessingStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
    errorMessage *string
    // Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
    lastMembershipUpdated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.  Required. Read-only.
    status *MembershipRuleProcessingStatusDetails
}
// NewMembershipRuleProcessingStatus instantiates a new membershipRuleProcessingStatus and sets the default values.
func NewMembershipRuleProcessingStatus()(*MembershipRuleProcessingStatus) {
    m := &MembershipRuleProcessingStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMembershipRuleProcessingStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMembershipRuleProcessingStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMembershipRuleProcessingStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MembershipRuleProcessingStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetErrorMessage gets the errorMessage property value. Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) GetErrorMessage()(*string) {
    return m.errorMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MembershipRuleProcessingStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["lastMembershipUpdated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastMembershipUpdated(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMembershipRuleProcessingStatusDetails)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*MembershipRuleProcessingStatusDetails))
        }
        return nil
    }
    return res
}
// GetLastMembershipUpdated gets the lastMembershipUpdated property value. Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) GetLastMembershipUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastMembershipUpdated
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MembershipRuleProcessingStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.  Required. Read-only.
func (m *MembershipRuleProcessingStatus) GetStatus()(*MembershipRuleProcessingStatusDetails) {
    return m.status
}
// Serialize serializes information the current object
func (m *MembershipRuleProcessingStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastMembershipUpdated", m.GetLastMembershipUpdated())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *MembershipRuleProcessingStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetErrorMessage sets the errorMessage property value. Detailed error message if dynamic group processing ran into an error.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// SetLastMembershipUpdated sets the lastMembershipUpdated property value. Most recent date and time when membership of a dynamic group was updated.  Optional. Read-only.
func (m *MembershipRuleProcessingStatus) SetLastMembershipUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastMembershipUpdated = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MembershipRuleProcessingStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. Current status of a dynamic group processing. Possible values are: NotStarted, Running, Succeeded, Failed, and UnknownFutureValue.  Required. Read-only.
func (m *MembershipRuleProcessingStatus) SetStatus(value *MembershipRuleProcessingStatusDetails)() {
    m.status = value
}
