package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExcludedApps contains properties for Excluded Office365 Apps.
type ExcludedApps struct {
    // The value for if MS Office Access should be excluded or not.
    access *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value for if Microsoft Search as default should be excluded or not.
    bing *bool
    // The value for if MS Office Excel should be excluded or not.
    excel *bool
    // The value for if MS Office OneDrive for Business - Groove should be excluded or not.
    groove *bool
    // The value for if MS Office InfoPath should be excluded or not.
    infoPath *bool
    // The value for if MS Office Skype for Business - Lync should be excluded or not.
    lync *bool
    // The OdataType property
    odataType *string
    // The value for if MS Office OneDrive should be excluded or not.
    oneDrive *bool
    // The value for if MS Office OneNote should be excluded or not.
    oneNote *bool
    // The value for if MS Office Outlook should be excluded or not.
    outlook *bool
    // The value for if MS Office PowerPoint should be excluded or not.
    powerPoint *bool
    // The value for if MS Office Publisher should be excluded or not.
    publisher *bool
    // The value for if MS Office SharePointDesigner should be excluded or not.
    sharePointDesigner *bool
    // The value for if MS Office Teams should be excluded or not.
    teams *bool
    // The value for if MS Office Visio should be excluded or not.
    visio *bool
    // The value for if MS Office Word should be excluded or not.
    word *bool
}
// NewExcludedApps instantiates a new excludedApps and sets the default values.
func NewExcludedApps()(*ExcludedApps) {
    m := &ExcludedApps{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExcludedAppsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExcludedAppsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExcludedApps(), nil
}
// GetAccess gets the access property value. The value for if MS Office Access should be excluded or not.
func (m *ExcludedApps) GetAccess()(*bool) {
    return m.access
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExcludedApps) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBing gets the bing property value. The value for if Microsoft Search as default should be excluded or not.
func (m *ExcludedApps) GetBing()(*bool) {
    return m.bing
}
// GetExcel gets the excel property value. The value for if MS Office Excel should be excluded or not.
func (m *ExcludedApps) GetExcel()(*bool) {
    return m.excel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExcludedApps) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val)
        }
        return nil
    }
    res["bing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBing(val)
        }
        return nil
    }
    res["excel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcel(val)
        }
        return nil
    }
    res["groove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroove(val)
        }
        return nil
    }
    res["infoPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfoPath(val)
        }
        return nil
    }
    res["lync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLync(val)
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
    res["oneDrive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDrive(val)
        }
        return nil
    }
    res["oneNote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNote(val)
        }
        return nil
    }
    res["outlook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlook(val)
        }
        return nil
    }
    res["powerPoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerPoint(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["sharePointDesigner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointDesigner(val)
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeams(val)
        }
        return nil
    }
    res["visio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisio(val)
        }
        return nil
    }
    res["word"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWord(val)
        }
        return nil
    }
    return res
}
// GetGroove gets the groove property value. The value for if MS Office OneDrive for Business - Groove should be excluded or not.
func (m *ExcludedApps) GetGroove()(*bool) {
    return m.groove
}
// GetInfoPath gets the infoPath property value. The value for if MS Office InfoPath should be excluded or not.
func (m *ExcludedApps) GetInfoPath()(*bool) {
    return m.infoPath
}
// GetLync gets the lync property value. The value for if MS Office Skype for Business - Lync should be excluded or not.
func (m *ExcludedApps) GetLync()(*bool) {
    return m.lync
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExcludedApps) GetOdataType()(*string) {
    return m.odataType
}
// GetOneDrive gets the oneDrive property value. The value for if MS Office OneDrive should be excluded or not.
func (m *ExcludedApps) GetOneDrive()(*bool) {
    return m.oneDrive
}
// GetOneNote gets the oneNote property value. The value for if MS Office OneNote should be excluded or not.
func (m *ExcludedApps) GetOneNote()(*bool) {
    return m.oneNote
}
// GetOutlook gets the outlook property value. The value for if MS Office Outlook should be excluded or not.
func (m *ExcludedApps) GetOutlook()(*bool) {
    return m.outlook
}
// GetPowerPoint gets the powerPoint property value. The value for if MS Office PowerPoint should be excluded or not.
func (m *ExcludedApps) GetPowerPoint()(*bool) {
    return m.powerPoint
}
// GetPublisher gets the publisher property value. The value for if MS Office Publisher should be excluded or not.
func (m *ExcludedApps) GetPublisher()(*bool) {
    return m.publisher
}
// GetSharePointDesigner gets the sharePointDesigner property value. The value for if MS Office SharePointDesigner should be excluded or not.
func (m *ExcludedApps) GetSharePointDesigner()(*bool) {
    return m.sharePointDesigner
}
// GetTeams gets the teams property value. The value for if MS Office Teams should be excluded or not.
func (m *ExcludedApps) GetTeams()(*bool) {
    return m.teams
}
// GetVisio gets the visio property value. The value for if MS Office Visio should be excluded or not.
func (m *ExcludedApps) GetVisio()(*bool) {
    return m.visio
}
// GetWord gets the word property value. The value for if MS Office Word should be excluded or not.
func (m *ExcludedApps) GetWord()(*bool) {
    return m.word
}
// Serialize serializes information the current object
func (m *ExcludedApps) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("access", m.GetAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("bing", m.GetBing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("excel", m.GetExcel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("groove", m.GetGroove())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("infoPath", m.GetInfoPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("lync", m.GetLync())
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
        err := writer.WriteBoolValue("oneDrive", m.GetOneDrive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("oneNote", m.GetOneNote())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("outlook", m.GetOutlook())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("powerPoint", m.GetPowerPoint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sharePointDesigner", m.GetSharePointDesigner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("teams", m.GetTeams())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("visio", m.GetVisio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("word", m.GetWord())
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
// SetAccess sets the access property value. The value for if MS Office Access should be excluded or not.
func (m *ExcludedApps) SetAccess(value *bool)() {
    m.access = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExcludedApps) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBing sets the bing property value. The value for if Microsoft Search as default should be excluded or not.
func (m *ExcludedApps) SetBing(value *bool)() {
    m.bing = value
}
// SetExcel sets the excel property value. The value for if MS Office Excel should be excluded or not.
func (m *ExcludedApps) SetExcel(value *bool)() {
    m.excel = value
}
// SetGroove sets the groove property value. The value for if MS Office OneDrive for Business - Groove should be excluded or not.
func (m *ExcludedApps) SetGroove(value *bool)() {
    m.groove = value
}
// SetInfoPath sets the infoPath property value. The value for if MS Office InfoPath should be excluded or not.
func (m *ExcludedApps) SetInfoPath(value *bool)() {
    m.infoPath = value
}
// SetLync sets the lync property value. The value for if MS Office Skype for Business - Lync should be excluded or not.
func (m *ExcludedApps) SetLync(value *bool)() {
    m.lync = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExcludedApps) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOneDrive sets the oneDrive property value. The value for if MS Office OneDrive should be excluded or not.
func (m *ExcludedApps) SetOneDrive(value *bool)() {
    m.oneDrive = value
}
// SetOneNote sets the oneNote property value. The value for if MS Office OneNote should be excluded or not.
func (m *ExcludedApps) SetOneNote(value *bool)() {
    m.oneNote = value
}
// SetOutlook sets the outlook property value. The value for if MS Office Outlook should be excluded or not.
func (m *ExcludedApps) SetOutlook(value *bool)() {
    m.outlook = value
}
// SetPowerPoint sets the powerPoint property value. The value for if MS Office PowerPoint should be excluded or not.
func (m *ExcludedApps) SetPowerPoint(value *bool)() {
    m.powerPoint = value
}
// SetPublisher sets the publisher property value. The value for if MS Office Publisher should be excluded or not.
func (m *ExcludedApps) SetPublisher(value *bool)() {
    m.publisher = value
}
// SetSharePointDesigner sets the sharePointDesigner property value. The value for if MS Office SharePointDesigner should be excluded or not.
func (m *ExcludedApps) SetSharePointDesigner(value *bool)() {
    m.sharePointDesigner = value
}
// SetTeams sets the teams property value. The value for if MS Office Teams should be excluded or not.
func (m *ExcludedApps) SetTeams(value *bool)() {
    m.teams = value
}
// SetVisio sets the visio property value. The value for if MS Office Visio should be excluded or not.
func (m *ExcludedApps) SetVisio(value *bool)() {
    m.visio = value
}
// SetWord sets the word property value. The value for if MS Office Word should be excluded or not.
func (m *ExcludedApps) SetWord(value *bool)() {
    m.word = value
}
