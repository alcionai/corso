package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSiteHistory the history for the site modifications
type BrowserSiteHistory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
    allowRedirect *bool
    // The comment for the site.
    comment *string
    // Controls what compatibility setting is used for specific sites or domains. The possible values are: default, internetExplorer8Enterprise, internetExplorer7Enterprise, internetExplorer11, internetExplorer10, internetExplorer9, internetExplorer8, internetExplorer7, internetExplorer5, unknownFutureValue.
    compatibilityMode *BrowserSiteCompatibilityMode
    // The user who last modified the site.
    lastModifiedBy IdentitySetable
    // The merge type of the site. The possible values are: noMerge, default, unknownFutureValue.
    mergeType *BrowserSiteMergeType
    // The OdataType property
    odataType *string
    // The date and time when the site was last published.
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The target environment that the site should open in. The possible values are: internetExplorerMode, internetExplorer11, microsoftEdge, configurable, none, unknownFutureValue.Prior to June 15, 2022, the internetExplorer11 option would allow opening a site in the Internet Explorer 11 (IE11) desktop application. Following the retirement of IE11 on June 15, 2022, the internetExplorer11 option will no longer open an IE11 window and will instead behave the same as the internetExplorerMode option.
    targetEnvironment *BrowserSiteTargetEnvironment
}
// NewBrowserSiteHistory instantiates a new browserSiteHistory and sets the default values.
func NewBrowserSiteHistory()(*BrowserSiteHistory) {
    m := &BrowserSiteHistory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBrowserSiteHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSiteHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSiteHistory(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BrowserSiteHistory) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowRedirect gets the allowRedirect property value. Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
func (m *BrowserSiteHistory) GetAllowRedirect()(*bool) {
    return m.allowRedirect
}
// GetComment gets the comment property value. The comment for the site.
func (m *BrowserSiteHistory) GetComment()(*string) {
    return m.comment
}
// GetCompatibilityMode gets the compatibilityMode property value. Controls what compatibility setting is used for specific sites or domains. The possible values are: default, internetExplorer8Enterprise, internetExplorer7Enterprise, internetExplorer11, internetExplorer10, internetExplorer9, internetExplorer8, internetExplorer7, internetExplorer5, unknownFutureValue.
func (m *BrowserSiteHistory) GetCompatibilityMode()(*BrowserSiteCompatibilityMode) {
    return m.compatibilityMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSiteHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowRedirect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowRedirect(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["compatibilityMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteCompatibilityMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompatibilityMode(val.(*BrowserSiteCompatibilityMode))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["mergeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteMergeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergeType(val.(*BrowserSiteMergeType))
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
    res["publishedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    res["targetEnvironment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBrowserSiteTargetEnvironment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetEnvironment(val.(*BrowserSiteTargetEnvironment))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the site.
func (m *BrowserSiteHistory) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetMergeType gets the mergeType property value. The merge type of the site. The possible values are: noMerge, default, unknownFutureValue.
func (m *BrowserSiteHistory) GetMergeType()(*BrowserSiteMergeType) {
    return m.mergeType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BrowserSiteHistory) GetOdataType()(*string) {
    return m.odataType
}
// GetPublishedDateTime gets the publishedDateTime property value. The date and time when the site was last published.
func (m *BrowserSiteHistory) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.publishedDateTime
}
// GetTargetEnvironment gets the targetEnvironment property value. The target environment that the site should open in. The possible values are: internetExplorerMode, internetExplorer11, microsoftEdge, configurable, none, unknownFutureValue.Prior to June 15, 2022, the internetExplorer11 option would allow opening a site in the Internet Explorer 11 (IE11) desktop application. Following the retirement of IE11 on June 15, 2022, the internetExplorer11 option will no longer open an IE11 window and will instead behave the same as the internetExplorerMode option.
func (m *BrowserSiteHistory) GetTargetEnvironment()(*BrowserSiteTargetEnvironment) {
    return m.targetEnvironment
}
// Serialize serializes information the current object
func (m *BrowserSiteHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowRedirect", m.GetAllowRedirect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    if m.GetCompatibilityMode() != nil {
        cast := (*m.GetCompatibilityMode()).String()
        err := writer.WriteStringValue("compatibilityMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    if m.GetMergeType() != nil {
        cast := (*m.GetMergeType()).String()
        err := writer.WriteStringValue("mergeType", &cast)
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
        err := writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetTargetEnvironment() != nil {
        cast := (*m.GetTargetEnvironment()).String()
        err := writer.WriteStringValue("targetEnvironment", &cast)
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
func (m *BrowserSiteHistory) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowRedirect sets the allowRedirect property value. Controls the behavior of redirected sites. If true, indicates that the site will open in Internet Explorer 11 or Microsoft Edge even if the site is navigated to as part of a HTTP or meta refresh redirection chain.
func (m *BrowserSiteHistory) SetAllowRedirect(value *bool)() {
    m.allowRedirect = value
}
// SetComment sets the comment property value. The comment for the site.
func (m *BrowserSiteHistory) SetComment(value *string)() {
    m.comment = value
}
// SetCompatibilityMode sets the compatibilityMode property value. Controls what compatibility setting is used for specific sites or domains. The possible values are: default, internetExplorer8Enterprise, internetExplorer7Enterprise, internetExplorer11, internetExplorer10, internetExplorer9, internetExplorer8, internetExplorer7, internetExplorer5, unknownFutureValue.
func (m *BrowserSiteHistory) SetCompatibilityMode(value *BrowserSiteCompatibilityMode)() {
    m.compatibilityMode = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the site.
func (m *BrowserSiteHistory) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetMergeType sets the mergeType property value. The merge type of the site. The possible values are: noMerge, default, unknownFutureValue.
func (m *BrowserSiteHistory) SetMergeType(value *BrowserSiteMergeType)() {
    m.mergeType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BrowserSiteHistory) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPublishedDateTime sets the publishedDateTime property value. The date and time when the site was last published.
func (m *BrowserSiteHistory) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
// SetTargetEnvironment sets the targetEnvironment property value. The target environment that the site should open in. The possible values are: internetExplorerMode, internetExplorer11, microsoftEdge, configurable, none, unknownFutureValue.Prior to June 15, 2022, the internetExplorer11 option would allow opening a site in the Internet Explorer 11 (IE11) desktop application. Following the retirement of IE11 on June 15, 2022, the internetExplorer11 option will no longer open an IE11 window and will instead behave the same as the internetExplorerMode option.
func (m *BrowserSiteHistory) SetTargetEnvironment(value *BrowserSiteTargetEnvironment)() {
    m.targetEnvironment = value
}
