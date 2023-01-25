package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppIntentAndStateDetail mobile App Intent and Install State for a given device.
type MobileAppIntentAndStateDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // MobieApp identifier.
    applicationId *string
    // The admin provided or imported title of the app.
    displayName *string
    // Human readable version of the application
    displayVersion *string
    // A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
    installState *ResultantAppState
    // Indicates the status of the mobile app on the device.
    mobileAppIntent *MobileAppIntent
    // The OdataType property
    odataType *string
    // The supported platforms for the app.
    supportedDeviceTypes []MobileAppSupportedDeviceTypeable
}
// NewMobileAppIntentAndStateDetail instantiates a new mobileAppIntentAndStateDetail and sets the default values.
func NewMobileAppIntentAndStateDetail()(*MobileAppIntentAndStateDetail) {
    m := &MobileAppIntentAndStateDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppIntentAndStateDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppIntentAndStateDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppIntentAndStateDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppIntentAndStateDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicationId gets the applicationId property value. MobieApp identifier.
func (m *MobileAppIntentAndStateDetail) GetApplicationId()(*string) {
    return m.applicationId
}
// GetDisplayName gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileAppIntentAndStateDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetDisplayVersion gets the displayVersion property value. Human readable version of the application
func (m *MobileAppIntentAndStateDetail) GetDisplayVersion()(*string) {
    return m.displayVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppIntentAndStateDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
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
    res["displayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["installState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallState(val.(*ResultantAppState))
        }
        return nil
    }
    res["mobileAppIntent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppIntent(val.(*MobileAppIntent))
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
    res["supportedDeviceTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppSupportedDeviceTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppSupportedDeviceTypeable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppSupportedDeviceTypeable)
            }
            m.SetSupportedDeviceTypes(res)
        }
        return nil
    }
    return res
}
// GetInstallState gets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppIntentAndStateDetail) GetInstallState()(*ResultantAppState) {
    return m.installState
}
// GetMobileAppIntent gets the mobileAppIntent property value. Indicates the status of the mobile app on the device.
func (m *MobileAppIntentAndStateDetail) GetMobileAppIntent()(*MobileAppIntent) {
    return m.mobileAppIntent
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MobileAppIntentAndStateDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetSupportedDeviceTypes gets the supportedDeviceTypes property value. The supported platforms for the app.
func (m *MobileAppIntentAndStateDetail) GetSupportedDeviceTypes()([]MobileAppSupportedDeviceTypeable) {
    return m.supportedDeviceTypes
}
// Serialize serializes information the current object
func (m *MobileAppIntentAndStateDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
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
        err := writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := (*m.GetInstallState()).String()
        err := writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppIntent() != nil {
        cast := (*m.GetMobileAppIntent()).String()
        err := writer.WriteStringValue("mobileAppIntent", &cast)
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
    if m.GetSupportedDeviceTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSupportedDeviceTypes()))
        for i, v := range m.GetSupportedDeviceTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("supportedDeviceTypes", cast)
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
func (m *MobileAppIntentAndStateDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplicationId sets the applicationId property value. MobieApp identifier.
func (m *MobileAppIntentAndStateDetail) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetDisplayName sets the displayName property value. The admin provided or imported title of the app.
func (m *MobileAppIntentAndStateDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDisplayVersion sets the displayVersion property value. Human readable version of the application
func (m *MobileAppIntentAndStateDetail) SetDisplayVersion(value *string)() {
    m.displayVersion = value
}
// SetInstallState sets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppIntentAndStateDetail) SetInstallState(value *ResultantAppState)() {
    m.installState = value
}
// SetMobileAppIntent sets the mobileAppIntent property value. Indicates the status of the mobile app on the device.
func (m *MobileAppIntentAndStateDetail) SetMobileAppIntent(value *MobileAppIntent)() {
    m.mobileAppIntent = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MobileAppIntentAndStateDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSupportedDeviceTypes sets the supportedDeviceTypes property value. The supported platforms for the app.
func (m *MobileAppIntentAndStateDetail) SetSupportedDeviceTypes(value []MobileAppSupportedDeviceTypeable)() {
    m.supportedDeviceTypes = value
}
