package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelingOptions 
type LabelingOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentMethod property
    assignmentMethod *AssignmentMethod
    // The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
    downgradeJustification DowngradeJustificationable
    // Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
    extendedProperties []KeyValuePairable
    // The GUID of the label that should be applied to the information.
    labelId *string
    // The OdataType property
    odataType *string
}
// NewLabelingOptions instantiates a new labelingOptions and sets the default values.
func NewLabelingOptions()(*LabelingOptions) {
    m := &LabelingOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLabelingOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelingOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelingOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LabelingOptions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentMethod gets the assignmentMethod property value. The assignmentMethod property
func (m *LabelingOptions) GetAssignmentMethod()(*AssignmentMethod) {
    return m.assignmentMethod
}
// GetDowngradeJustification gets the downgradeJustification property value. The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
func (m *LabelingOptions) GetDowngradeJustification()(DowngradeJustificationable) {
    return m.downgradeJustification
}
// GetExtendedProperties gets the extendedProperties property value. Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
func (m *LabelingOptions) GetExtendedProperties()([]KeyValuePairable) {
    return m.extendedProperties
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelingOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentMethod(val.(*AssignmentMethod))
        }
        return nil
    }
    res["downgradeJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDowngradeJustificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeJustification(val.(DowngradeJustificationable))
        }
        return nil
    }
    res["extendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetExtendedProperties(res)
        }
        return nil
    }
    res["labelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelId(val)
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
// GetLabelId gets the labelId property value. The GUID of the label that should be applied to the information.
func (m *LabelingOptions) GetLabelId()(*string) {
    return m.labelId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LabelingOptions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LabelingOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("downgradeJustification", m.GetDowngradeJustification())
        if err != nil {
            return err
        }
    }
    if m.GetExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtendedProperties()))
        for i, v := range m.GetExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("extendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
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
func (m *LabelingOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentMethod sets the assignmentMethod property value. The assignmentMethod property
func (m *LabelingOptions) SetAssignmentMethod(value *AssignmentMethod)() {
    m.assignmentMethod = value
}
// SetDowngradeJustification sets the downgradeJustification property value. The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
func (m *LabelingOptions) SetDowngradeJustification(value DowngradeJustificationable)() {
    m.downgradeJustification = value
}
// SetExtendedProperties sets the extendedProperties property value. Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
func (m *LabelingOptions) SetExtendedProperties(value []KeyValuePairable)() {
    m.extendedProperties = value
}
// SetLabelId sets the labelId property value. The GUID of the label that should be applied to the information.
func (m *LabelingOptions) SetLabelId(value *string)() {
    m.labelId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LabelingOptions) SetOdataType(value *string)() {
    m.odataType = value
}
