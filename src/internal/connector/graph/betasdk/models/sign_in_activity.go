package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SignInActivity 
type SignInActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The last non-interactive sign-in date for a specific user. You can use this field to calculate the last time a client signed in to the directory on behalf of a user. Because some users may use clients to access tenant resources rather than signing into your tenant directly, you can use the non-interactive sign-in date to along with lastSignInDateTime to identify inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. Azure AD maintains non-interactive sign-ins going back to May 2020. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
    lastNonInteractiveSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Request identifier of the last non-interactive sign-in performed by this user.
    lastNonInteractiveSignInRequestId *string
    // The last interactive sign-in date and time for a specific user. You can use this field to calculate the last time a user signed in to the directory with an interactive authentication method. This field can be used to build reports, such as inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. Azure AD maintains interactive sign-ins going back to April 2020. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
    lastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Request identifier of the last interactive sign-in performed by this user.
    lastSignInRequestId *string
    // The OdataType property
    odataType *string
}
// NewSignInActivity instantiates a new signInActivity and sets the default values.
func NewSignInActivity()(*SignInActivity) {
    m := &SignInActivity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSignInActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSignInActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignInActivity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SignInActivity) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignInActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lastNonInteractiveSignInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastNonInteractiveSignInDateTime(val)
        }
        return nil
    }
    res["lastNonInteractiveSignInRequestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastNonInteractiveSignInRequestId(val)
        }
        return nil
    }
    res["lastSignInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSignInDateTime(val)
        }
        return nil
    }
    res["lastSignInRequestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSignInRequestId(val)
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
// GetLastNonInteractiveSignInDateTime gets the lastNonInteractiveSignInDateTime property value. The last non-interactive sign-in date for a specific user. You can use this field to calculate the last time a client signed in to the directory on behalf of a user. Because some users may use clients to access tenant resources rather than signing into your tenant directly, you can use the non-interactive sign-in date to along with lastSignInDateTime to identify inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. Azure AD maintains non-interactive sign-ins going back to May 2020. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
func (m *SignInActivity) GetLastNonInteractiveSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastNonInteractiveSignInDateTime
}
// GetLastNonInteractiveSignInRequestId gets the lastNonInteractiveSignInRequestId property value. Request identifier of the last non-interactive sign-in performed by this user.
func (m *SignInActivity) GetLastNonInteractiveSignInRequestId()(*string) {
    return m.lastNonInteractiveSignInRequestId
}
// GetLastSignInDateTime gets the lastSignInDateTime property value. The last interactive sign-in date and time for a specific user. You can use this field to calculate the last time a user signed in to the directory with an interactive authentication method. This field can be used to build reports, such as inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. Azure AD maintains interactive sign-ins going back to April 2020. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
func (m *SignInActivity) GetLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSignInDateTime
}
// GetLastSignInRequestId gets the lastSignInRequestId property value. Request identifier of the last interactive sign-in performed by this user.
func (m *SignInActivity) GetLastSignInRequestId()(*string) {
    return m.lastSignInRequestId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SignInActivity) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SignInActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastNonInteractiveSignInDateTime", m.GetLastNonInteractiveSignInDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastNonInteractiveSignInRequestId", m.GetLastNonInteractiveSignInRequestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSignInDateTime", m.GetLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastSignInRequestId", m.GetLastSignInRequestId())
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
func (m *SignInActivity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLastNonInteractiveSignInDateTime sets the lastNonInteractiveSignInDateTime property value. The last non-interactive sign-in date for a specific user. You can use this field to calculate the last time a client signed in to the directory on behalf of a user. Because some users may use clients to access tenant resources rather than signing into your tenant directly, you can use the non-interactive sign-in date to along with lastSignInDateTime to identify inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. Azure AD maintains non-interactive sign-ins going back to May 2020. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
func (m *SignInActivity) SetLastNonInteractiveSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastNonInteractiveSignInDateTime = value
}
// SetLastNonInteractiveSignInRequestId sets the lastNonInteractiveSignInRequestId property value. Request identifier of the last non-interactive sign-in performed by this user.
func (m *SignInActivity) SetLastNonInteractiveSignInRequestId(value *string)() {
    m.lastNonInteractiveSignInRequestId = value
}
// SetLastSignInDateTime sets the lastSignInDateTime property value. The last interactive sign-in date and time for a specific user. You can use this field to calculate the last time a user signed in to the directory with an interactive authentication method. This field can be used to build reports, such as inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. Azure AD maintains interactive sign-ins going back to April 2020. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
func (m *SignInActivity) SetLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSignInDateTime = value
}
// SetLastSignInRequestId sets the lastSignInRequestId property value. Request identifier of the last interactive sign-in performed by this user.
func (m *SignInActivity) SetLastSignInRequestId(value *string)() {
    m.lastSignInRequestId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SignInActivity) SetOdataType(value *string)() {
    m.odataType = value
}
