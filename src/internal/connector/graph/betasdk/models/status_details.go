package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusDetails 
type StatusDetails struct {
    StatusBase
    // Additional details in case of error.
    additionalDetails *string
    // Categorizes the error code. Possible values are Failure, NonServiceFailure, Success.
    errorCategory *ProvisioningStatusErrorCategory
    // Unique error code if any occurred. Learn more
    errorCode *string
    // Summarizes the status and describes why the status happened.
    reason *string
    // Provides the resolution for the corresponding error.
    recommendedAction *string
}
// NewStatusDetails instantiates a new StatusDetails and sets the default values.
func NewStatusDetails()(*StatusDetails) {
    m := &StatusDetails{
        StatusBase: *NewStatusBase(),
    }
    odataTypeValue := "#microsoft.graph.statusDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatusDetails(), nil
}
// GetAdditionalDetails gets the additionalDetails property value. Additional details in case of error.
func (m *StatusDetails) GetAdditionalDetails()(*string) {
    return m.additionalDetails
}
// GetErrorCategory gets the errorCategory property value. Categorizes the error code. Possible values are Failure, NonServiceFailure, Success.
func (m *StatusDetails) GetErrorCategory()(*ProvisioningStatusErrorCategory) {
    return m.errorCategory
}
// GetErrorCode gets the errorCode property value. Unique error code if any occurred. Learn more
func (m *StatusDetails) GetErrorCode()(*string) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StatusBase.GetFieldDeserializers()
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetails(val)
        }
        return nil
    }
    res["errorCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningStatusErrorCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCategory(val.(*ProvisioningStatusErrorCategory))
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["recommendedAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedAction(val)
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. Summarizes the status and describes why the status happened.
func (m *StatusDetails) GetReason()(*string) {
    return m.reason
}
// GetRecommendedAction gets the recommendedAction property value. Provides the resolution for the corresponding error.
func (m *StatusDetails) GetRecommendedAction()(*string) {
    return m.recommendedAction
}
// Serialize serializes information the current object
func (m *StatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StatusBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    if m.GetErrorCategory() != nil {
        cast := (*m.GetErrorCategory()).String()
        err = writer.WriteStringValue("errorCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalDetails sets the additionalDetails property value. Additional details in case of error.
func (m *StatusDetails) SetAdditionalDetails(value *string)() {
    m.additionalDetails = value
}
// SetErrorCategory sets the errorCategory property value. Categorizes the error code. Possible values are Failure, NonServiceFailure, Success.
func (m *StatusDetails) SetErrorCategory(value *ProvisioningStatusErrorCategory)() {
    m.errorCategory = value
}
// SetErrorCode sets the errorCode property value. Unique error code if any occurred. Learn more
func (m *StatusDetails) SetErrorCode(value *string)() {
    m.errorCode = value
}
// SetReason sets the reason property value. Summarizes the status and describes why the status happened.
func (m *StatusDetails) SetReason(value *string)() {
    m.reason = value
}
// SetRecommendedAction sets the recommendedAction property value. Provides the resolution for the corresponding error.
func (m *StatusDetails) SetRecommendedAction(value *string)() {
    m.recommendedAction = value
}
