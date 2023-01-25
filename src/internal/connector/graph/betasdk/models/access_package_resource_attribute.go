package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceAttribute 
type AccessPackageResourceAttribute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore object type.
    attributeDestination AccessPackageResourceAttributeDestinationable
    // The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension_2b676109c7c74ae2b41549205f1947ed_personalTitle.
    attributeName *string
    // Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion object type.
    attributeSource AccessPackageResourceAttributeSourceable
    // Unique identifier for the attribute on the access package resource. Read-only.
    id *string
    // Specifies whether or not an existing attribute value can be edited by the requester.
    isEditable *bool
    // Specifies whether the attribute will remain in the end system after an assignment ends.
    isPersistedOnAssignmentRemoval *bool
    // The OdataType property
    odataType *string
}
// NewAccessPackageResourceAttribute instantiates a new accessPackageResourceAttribute and sets the default values.
func NewAccessPackageResourceAttribute()(*AccessPackageResourceAttribute) {
    m := &AccessPackageResourceAttribute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageResourceAttributeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceAttributeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceAttribute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageResourceAttribute) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttributeDestination gets the attributeDestination property value. Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore object type.
func (m *AccessPackageResourceAttribute) GetAttributeDestination()(AccessPackageResourceAttributeDestinationable) {
    return m.attributeDestination
}
// GetAttributeName gets the attributeName property value. The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension_2b676109c7c74ae2b41549205f1947ed_personalTitle.
func (m *AccessPackageResourceAttribute) GetAttributeName()(*string) {
    return m.attributeName
}
// GetAttributeSource gets the attributeSource property value. Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion object type.
func (m *AccessPackageResourceAttribute) GetAttributeSource()(AccessPackageResourceAttributeSourceable) {
    return m.attributeSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceAttribute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributeDestination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceAttributeDestinationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeDestination(val.(AccessPackageResourceAttributeDestinationable))
        }
        return nil
    }
    res["attributeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeName(val)
        }
        return nil
    }
    res["attributeSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceAttributeSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeSource(val.(AccessPackageResourceAttributeSourceable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isEditable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEditable(val)
        }
        return nil
    }
    res["isPersistedOnAssignmentRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPersistedOnAssignmentRemoval(val)
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
// GetId gets the id property value. Unique identifier for the attribute on the access package resource. Read-only.
func (m *AccessPackageResourceAttribute) GetId()(*string) {
    return m.id
}
// GetIsEditable gets the isEditable property value. Specifies whether or not an existing attribute value can be edited by the requester.
func (m *AccessPackageResourceAttribute) GetIsEditable()(*bool) {
    return m.isEditable
}
// GetIsPersistedOnAssignmentRemoval gets the isPersistedOnAssignmentRemoval property value. Specifies whether the attribute will remain in the end system after an assignment ends.
func (m *AccessPackageResourceAttribute) GetIsPersistedOnAssignmentRemoval()(*bool) {
    return m.isPersistedOnAssignmentRemoval
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageResourceAttribute) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AccessPackageResourceAttribute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attributeDestination", m.GetAttributeDestination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("attributeName", m.GetAttributeName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("attributeSource", m.GetAttributeSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEditable", m.GetIsEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPersistedOnAssignmentRemoval", m.GetIsPersistedOnAssignmentRemoval())
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
func (m *AccessPackageResourceAttribute) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttributeDestination sets the attributeDestination property value. Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore object type.
func (m *AccessPackageResourceAttribute) SetAttributeDestination(value AccessPackageResourceAttributeDestinationable)() {
    m.attributeDestination = value
}
// SetAttributeName sets the attributeName property value. The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension_2b676109c7c74ae2b41549205f1947ed_personalTitle.
func (m *AccessPackageResourceAttribute) SetAttributeName(value *string)() {
    m.attributeName = value
}
// SetAttributeSource sets the attributeSource property value. Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion object type.
func (m *AccessPackageResourceAttribute) SetAttributeSource(value AccessPackageResourceAttributeSourceable)() {
    m.attributeSource = value
}
// SetId sets the id property value. Unique identifier for the attribute on the access package resource. Read-only.
func (m *AccessPackageResourceAttribute) SetId(value *string)() {
    m.id = value
}
// SetIsEditable sets the isEditable property value. Specifies whether or not an existing attribute value can be edited by the requester.
func (m *AccessPackageResourceAttribute) SetIsEditable(value *bool)() {
    m.isEditable = value
}
// SetIsPersistedOnAssignmentRemoval sets the isPersistedOnAssignmentRemoval property value. Specifies whether the attribute will remain in the end system after an assignment ends.
func (m *AccessPackageResourceAttribute) SetIsPersistedOnAssignmentRemoval(value *bool)() {
    m.isPersistedOnAssignmentRemoval = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageResourceAttribute) SetOdataType(value *string)() {
    m.odataType = value
}
