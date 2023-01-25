package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InboundSharedUserProfile 
type InboundSharedUserProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The homeTenantId property
    homeTenantId *string
    // The OdataType property
    odataType *string
    // The userId property
    userId *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewInboundSharedUserProfile instantiates a new inboundSharedUserProfile and sets the default values.
func NewInboundSharedUserProfile()(*InboundSharedUserProfile) {
    m := &InboundSharedUserProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInboundSharedUserProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInboundSharedUserProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInboundSharedUserProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InboundSharedUserProfile) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *InboundSharedUserProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InboundSharedUserProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["homeTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
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
    return res
}
// GetHomeTenantId gets the homeTenantId property value. The homeTenantId property
func (m *InboundSharedUserProfile) GetHomeTenantId()(*string) {
    return m.homeTenantId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InboundSharedUserProfile) GetOdataType()(*string) {
    return m.odataType
}
// GetUserId gets the userId property value. The userId property
func (m *InboundSharedUserProfile) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *InboundSharedUserProfile) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *InboundSharedUserProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("homeTenantId", m.GetHomeTenantId())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InboundSharedUserProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *InboundSharedUserProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHomeTenantId sets the homeTenantId property value. The homeTenantId property
func (m *InboundSharedUserProfile) SetHomeTenantId(value *string)() {
    m.homeTenantId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InboundSharedUserProfile) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserId sets the userId property value. The userId property
func (m *InboundSharedUserProfile) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *InboundSharedUserProfile) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
