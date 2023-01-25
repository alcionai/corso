package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarkContent 
type MarkContent struct {
    LabelActionBase
    // The fontColor property
    fontColor *string
    // The fontSize property
    fontSize *int64
    // The text property
    text *string
}
// NewMarkContent instantiates a new MarkContent and sets the default values.
func NewMarkContent()(*MarkContent) {
    m := &MarkContent{
        LabelActionBase: *NewLabelActionBase(),
    }
    odataTypeValue := "#microsoft.graph.markContent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMarkContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMarkContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.addFooter":
                        return NewAddFooter(), nil
                    case "#microsoft.graph.addHeader":
                        return NewAddHeader(), nil
                    case "#microsoft.graph.addWatermark":
                        return NewAddWatermark(), nil
                }
            }
        }
    }
    return NewMarkContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MarkContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["fontColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontColor(val)
        }
        return nil
    }
    res["fontSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontSize(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetFontColor gets the fontColor property value. The fontColor property
func (m *MarkContent) GetFontColor()(*string) {
    return m.fontColor
}
// GetFontSize gets the fontSize property value. The fontSize property
func (m *MarkContent) GetFontSize()(*int64) {
    return m.fontSize
}
// GetText gets the text property value. The text property
func (m *MarkContent) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *MarkContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fontColor", m.GetFontColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("fontSize", m.GetFontSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFontColor sets the fontColor property value. The fontColor property
func (m *MarkContent) SetFontColor(value *string)() {
    m.fontColor = value
}
// SetFontSize sets the fontSize property value. The fontSize property
func (m *MarkContent) SetFontSize(value *int64)() {
    m.fontSize = value
}
// SetText sets the text property value. The text property
func (m *MarkContent) SetText(value *string)() {
    m.text = value
}
