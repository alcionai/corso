package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Command provides operations to manage the collection of site entities.
type Command struct {
    Entity
    // The appServiceName property
    appServiceName *string
    // The error property
    error *string
    // The packageFamilyName property
    packageFamilyName *string
    // The payload property
    payload PayloadRequestable
    // The permissionTicket property
    permissionTicket *string
    // The postBackUri property
    postBackUri *string
    // The responsepayload property
    responsepayload PayloadResponseable
    // The status property
    status *string
    // The type property
    type_escaped *string
}
// NewCommand instantiates a new command and sets the default values.
func NewCommand()(*Command) {
    m := &Command{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCommandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommand(), nil
}
// GetAppServiceName gets the appServiceName property value. The appServiceName property
func (m *Command) GetAppServiceName()(*string) {
    return m.appServiceName
}
// GetError gets the error property value. The error property
func (m *Command) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Command) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appServiceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppServiceName(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["packageFamilyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageFamilyName(val)
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(PayloadRequestable))
        }
        return nil
    }
    res["permissionTicket"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionTicket(val)
        }
        return nil
    }
    res["postBackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBackUri(val)
        }
        return nil
    }
    res["responsepayload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsepayload(val.(PayloadResponseable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetPackageFamilyName gets the packageFamilyName property value. The packageFamilyName property
func (m *Command) GetPackageFamilyName()(*string) {
    return m.packageFamilyName
}
// GetPayload gets the payload property value. The payload property
func (m *Command) GetPayload()(PayloadRequestable) {
    return m.payload
}
// GetPermissionTicket gets the permissionTicket property value. The permissionTicket property
func (m *Command) GetPermissionTicket()(*string) {
    return m.permissionTicket
}
// GetPostBackUri gets the postBackUri property value. The postBackUri property
func (m *Command) GetPostBackUri()(*string) {
    return m.postBackUri
}
// GetResponsepayload gets the responsepayload property value. The responsepayload property
func (m *Command) GetResponsepayload()(PayloadResponseable) {
    return m.responsepayload
}
// GetStatus gets the status property value. The status property
func (m *Command) GetStatus()(*string) {
    return m.status
}
// GetType gets the type property value. The type property
func (m *Command) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *Command) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appServiceName", m.GetAppServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageFamilyName", m.GetPackageFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionTicket", m.GetPermissionTicket())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postBackUri", m.GetPostBackUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("responsepayload", m.GetResponsepayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppServiceName sets the appServiceName property value. The appServiceName property
func (m *Command) SetAppServiceName(value *string)() {
    m.appServiceName = value
}
// SetError sets the error property value. The error property
func (m *Command) SetError(value *string)() {
    m.error = value
}
// SetPackageFamilyName sets the packageFamilyName property value. The packageFamilyName property
func (m *Command) SetPackageFamilyName(value *string)() {
    m.packageFamilyName = value
}
// SetPayload sets the payload property value. The payload property
func (m *Command) SetPayload(value PayloadRequestable)() {
    m.payload = value
}
// SetPermissionTicket sets the permissionTicket property value. The permissionTicket property
func (m *Command) SetPermissionTicket(value *string)() {
    m.permissionTicket = value
}
// SetPostBackUri sets the postBackUri property value. The postBackUri property
func (m *Command) SetPostBackUri(value *string)() {
    m.postBackUri = value
}
// SetResponsepayload sets the responsepayload property value. The responsepayload property
func (m *Command) SetResponsepayload(value PayloadResponseable)() {
    m.responsepayload = value
}
// SetStatus sets the status property value. The status property
func (m *Command) SetStatus(value *string)() {
    m.status = value
}
// SetType sets the type property value. The type property
func (m *Command) SetType(value *string)() {
    m.type_escaped = value
}
