package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeClientCheckinStatus 
type OfficeClientCheckinStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of policies delivered to the device as last checkin.
    appliedPolicies []string
    // Last device check-in time in UTC.
    checkinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device name trying to check-in.
    deviceName *string
    // Device platform trying to check-in.
    devicePlatform *string
    // Device platform version trying to check-in.
    devicePlatformVersion *string
    // Error message if any associated for the last checkin.
    errorMessage *string
    // The OdataType property
    odataType *string
    // User identifier using the device.
    userId *string
    // User principal name using the device.
    userPrincipalName *string
    // If the last checkin was successful.
    wasSuccessful *bool
}
// NewOfficeClientCheckinStatus instantiates a new officeClientCheckinStatus and sets the default values.
func NewOfficeClientCheckinStatus()(*OfficeClientCheckinStatus) {
    m := &OfficeClientCheckinStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOfficeClientCheckinStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeClientCheckinStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeClientCheckinStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeClientCheckinStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppliedPolicies gets the appliedPolicies property value. List of policies delivered to the device as last checkin.
func (m *OfficeClientCheckinStatus) GetAppliedPolicies()([]string) {
    return m.appliedPolicies
}
// GetCheckinDateTime gets the checkinDateTime property value. Last device check-in time in UTC.
func (m *OfficeClientCheckinStatus) GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.checkinDateTime
}
// GetDeviceName gets the deviceName property value. Device name trying to check-in.
func (m *OfficeClientCheckinStatus) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDevicePlatform gets the devicePlatform property value. Device platform trying to check-in.
func (m *OfficeClientCheckinStatus) GetDevicePlatform()(*string) {
    return m.devicePlatform
}
// GetDevicePlatformVersion gets the devicePlatformVersion property value. Device platform version trying to check-in.
func (m *OfficeClientCheckinStatus) GetDevicePlatformVersion()(*string) {
    return m.devicePlatformVersion
}
// GetErrorMessage gets the errorMessage property value. Error message if any associated for the last checkin.
func (m *OfficeClientCheckinStatus) GetErrorMessage()(*string) {
    return m.errorMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeClientCheckinStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliedPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAppliedPolicies(res)
        }
        return nil
    }
    res["checkinDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckinDateTime(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["devicePlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatform(val)
        }
        return nil
    }
    res["devicePlatformVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatformVersion(val)
        }
        return nil
    }
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
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["wasSuccessful"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWasSuccessful(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OfficeClientCheckinStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetUserId gets the userId property value. User identifier using the device.
func (m *OfficeClientCheckinStatus) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name using the device.
func (m *OfficeClientCheckinStatus) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetWasSuccessful gets the wasSuccessful property value. If the last checkin was successful.
func (m *OfficeClientCheckinStatus) GetWasSuccessful()(*bool) {
    return m.wasSuccessful
}
// Serialize serializes information the current object
func (m *OfficeClientCheckinStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppliedPolicies() != nil {
        err := writer.WriteCollectionOfStringValues("appliedPolicies", m.GetAppliedPolicies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("checkinDateTime", m.GetCheckinDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("devicePlatform", m.GetDevicePlatform())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("devicePlatformVersion", m.GetDevicePlatformVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorMessage", m.GetErrorMessage())
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
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wasSuccessful", m.GetWasSuccessful())
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
func (m *OfficeClientCheckinStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppliedPolicies sets the appliedPolicies property value. List of policies delivered to the device as last checkin.
func (m *OfficeClientCheckinStatus) SetAppliedPolicies(value []string)() {
    m.appliedPolicies = value
}
// SetCheckinDateTime sets the checkinDateTime property value. Last device check-in time in UTC.
func (m *OfficeClientCheckinStatus) SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.checkinDateTime = value
}
// SetDeviceName sets the deviceName property value. Device name trying to check-in.
func (m *OfficeClientCheckinStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDevicePlatform sets the devicePlatform property value. Device platform trying to check-in.
func (m *OfficeClientCheckinStatus) SetDevicePlatform(value *string)() {
    m.devicePlatform = value
}
// SetDevicePlatformVersion sets the devicePlatformVersion property value. Device platform version trying to check-in.
func (m *OfficeClientCheckinStatus) SetDevicePlatformVersion(value *string)() {
    m.devicePlatformVersion = value
}
// SetErrorMessage sets the errorMessage property value. Error message if any associated for the last checkin.
func (m *OfficeClientCheckinStatus) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OfficeClientCheckinStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserId sets the userId property value. User identifier using the device.
func (m *OfficeClientCheckinStatus) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name using the device.
func (m *OfficeClientCheckinStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetWasSuccessful sets the wasSuccessful property value. If the last checkin was successful.
func (m *OfficeClientCheckinStatus) SetWasSuccessful(value *bool)() {
    m.wasSuccessful = value
}
