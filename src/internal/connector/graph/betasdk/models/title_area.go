package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TitleArea 
type TitleArea struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Alternative text on the title area.
    alternativeText *string
    // Indicates whether the title area has a gradient effect enabled.
    enableGradientEffect *bool
    // URL of the image in the title area.
    imageWebUrl *string
    // Enumeration value that indicates the layout of the title area. The possible values are: imageAndTitle, plain, colorBlock, overlap, unknownFutureValue.
    layout *TitleAreaLayoutType
    // The OdataType property
    odataType *string
    // Contains collections of data that can be processed by server side services like search index and link fixup.
    serverProcessedContent ServerProcessedContentable
    // Indicates whether the author should be shown in title area.
    showAuthor *bool
    // Indicates whether the published date should be shown in title area.
    showPublishedDate *bool
    // Indicates whether the text block above title should be shown in title area.
    showTextBlockAboveTitle *bool
    // The text above title line.
    textAboveTitle *string
    // Enumeration value that indicates the text alignment of the title area. The possible values are: left, center, unknownFutureValue.
    textAlignment *TitleAreaTextAlignmentType
}
// NewTitleArea instantiates a new titleArea and sets the default values.
func NewTitleArea()(*TitleArea) {
    m := &TitleArea{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTitleAreaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTitleAreaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTitleArea(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TitleArea) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAlternativeText gets the alternativeText property value. Alternative text on the title area.
func (m *TitleArea) GetAlternativeText()(*string) {
    return m.alternativeText
}
// GetEnableGradientEffect gets the enableGradientEffect property value. Indicates whether the title area has a gradient effect enabled.
func (m *TitleArea) GetEnableGradientEffect()(*bool) {
    return m.enableGradientEffect
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TitleArea) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alternativeText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternativeText(val)
        }
        return nil
    }
    res["enableGradientEffect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableGradientEffect(val)
        }
        return nil
    }
    res["imageWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageWebUrl(val)
        }
        return nil
    }
    res["layout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTitleAreaLayoutType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLayout(val.(*TitleAreaLayoutType))
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
    res["serverProcessedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerProcessedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerProcessedContent(val.(ServerProcessedContentable))
        }
        return nil
    }
    res["showAuthor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowAuthor(val)
        }
        return nil
    }
    res["showPublishedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowPublishedDate(val)
        }
        return nil
    }
    res["showTextBlockAboveTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowTextBlockAboveTitle(val)
        }
        return nil
    }
    res["textAboveTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTextAboveTitle(val)
        }
        return nil
    }
    res["textAlignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTitleAreaTextAlignmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTextAlignment(val.(*TitleAreaTextAlignmentType))
        }
        return nil
    }
    return res
}
// GetImageWebUrl gets the imageWebUrl property value. URL of the image in the title area.
func (m *TitleArea) GetImageWebUrl()(*string) {
    return m.imageWebUrl
}
// GetLayout gets the layout property value. Enumeration value that indicates the layout of the title area. The possible values are: imageAndTitle, plain, colorBlock, overlap, unknownFutureValue.
func (m *TitleArea) GetLayout()(*TitleAreaLayoutType) {
    return m.layout
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TitleArea) GetOdataType()(*string) {
    return m.odataType
}
// GetServerProcessedContent gets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
func (m *TitleArea) GetServerProcessedContent()(ServerProcessedContentable) {
    return m.serverProcessedContent
}
// GetShowAuthor gets the showAuthor property value. Indicates whether the author should be shown in title area.
func (m *TitleArea) GetShowAuthor()(*bool) {
    return m.showAuthor
}
// GetShowPublishedDate gets the showPublishedDate property value. Indicates whether the published date should be shown in title area.
func (m *TitleArea) GetShowPublishedDate()(*bool) {
    return m.showPublishedDate
}
// GetShowTextBlockAboveTitle gets the showTextBlockAboveTitle property value. Indicates whether the text block above title should be shown in title area.
func (m *TitleArea) GetShowTextBlockAboveTitle()(*bool) {
    return m.showTextBlockAboveTitle
}
// GetTextAboveTitle gets the textAboveTitle property value. The text above title line.
func (m *TitleArea) GetTextAboveTitle()(*string) {
    return m.textAboveTitle
}
// GetTextAlignment gets the textAlignment property value. Enumeration value that indicates the text alignment of the title area. The possible values are: left, center, unknownFutureValue.
func (m *TitleArea) GetTextAlignment()(*TitleAreaTextAlignmentType) {
    return m.textAlignment
}
// Serialize serializes information the current object
func (m *TitleArea) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alternativeText", m.GetAlternativeText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableGradientEffect", m.GetEnableGradientEffect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("imageWebUrl", m.GetImageWebUrl())
        if err != nil {
            return err
        }
    }
    if m.GetLayout() != nil {
        cast := (*m.GetLayout()).String()
        err := writer.WriteStringValue("layout", &cast)
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
        err := writer.WriteObjectValue("serverProcessedContent", m.GetServerProcessedContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showAuthor", m.GetShowAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showPublishedDate", m.GetShowPublishedDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showTextBlockAboveTitle", m.GetShowTextBlockAboveTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("textAboveTitle", m.GetTextAboveTitle())
        if err != nil {
            return err
        }
    }
    if m.GetTextAlignment() != nil {
        cast := (*m.GetTextAlignment()).String()
        err := writer.WriteStringValue("textAlignment", &cast)
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
func (m *TitleArea) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAlternativeText sets the alternativeText property value. Alternative text on the title area.
func (m *TitleArea) SetAlternativeText(value *string)() {
    m.alternativeText = value
}
// SetEnableGradientEffect sets the enableGradientEffect property value. Indicates whether the title area has a gradient effect enabled.
func (m *TitleArea) SetEnableGradientEffect(value *bool)() {
    m.enableGradientEffect = value
}
// SetImageWebUrl sets the imageWebUrl property value. URL of the image in the title area.
func (m *TitleArea) SetImageWebUrl(value *string)() {
    m.imageWebUrl = value
}
// SetLayout sets the layout property value. Enumeration value that indicates the layout of the title area. The possible values are: imageAndTitle, plain, colorBlock, overlap, unknownFutureValue.
func (m *TitleArea) SetLayout(value *TitleAreaLayoutType)() {
    m.layout = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TitleArea) SetOdataType(value *string)() {
    m.odataType = value
}
// SetServerProcessedContent sets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
func (m *TitleArea) SetServerProcessedContent(value ServerProcessedContentable)() {
    m.serverProcessedContent = value
}
// SetShowAuthor sets the showAuthor property value. Indicates whether the author should be shown in title area.
func (m *TitleArea) SetShowAuthor(value *bool)() {
    m.showAuthor = value
}
// SetShowPublishedDate sets the showPublishedDate property value. Indicates whether the published date should be shown in title area.
func (m *TitleArea) SetShowPublishedDate(value *bool)() {
    m.showPublishedDate = value
}
// SetShowTextBlockAboveTitle sets the showTextBlockAboveTitle property value. Indicates whether the text block above title should be shown in title area.
func (m *TitleArea) SetShowTextBlockAboveTitle(value *bool)() {
    m.showTextBlockAboveTitle = value
}
// SetTextAboveTitle sets the textAboveTitle property value. The text above title line.
func (m *TitleArea) SetTextAboveTitle(value *string)() {
    m.textAboveTitle = value
}
// SetTextAlignment sets the textAlignment property value. Enumeration value that indicates the text alignment of the title area. The possible values are: left, center, unknownFutureValue.
func (m *TitleArea) SetTextAlignment(value *TitleAreaTextAlignmentType)() {
    m.textAlignment = value
}
