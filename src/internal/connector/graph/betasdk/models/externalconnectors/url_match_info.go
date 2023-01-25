package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UrlMatchInfo 
type UrlMatchInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A list of the URL prefixes that must match URLs to be processed by this URL-to-item-resolver.
    baseUrls []string
    // The OdataType property
    odataType *string
    // A regular expression that will be matched towards the URL that is processed by this URL-to-item-resolver. The ECMAScript specification for regular expressions (ECMA-262) is used for the evaluation. The named groups defined by the regular expression will be used later to extract values from the URL.
    urlPattern *string
}
// NewUrlMatchInfo instantiates a new urlMatchInfo and sets the default values.
func NewUrlMatchInfo()(*UrlMatchInfo) {
    m := &UrlMatchInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUrlMatchInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUrlMatchInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUrlMatchInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UrlMatchInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBaseUrls gets the baseUrls property value. A list of the URL prefixes that must match URLs to be processed by this URL-to-item-resolver.
func (m *UrlMatchInfo) GetBaseUrls()([]string) {
    return m.baseUrls
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UrlMatchInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["baseUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBaseUrls(res)
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
    res["urlPattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrlPattern(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UrlMatchInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetUrlPattern gets the urlPattern property value. A regular expression that will be matched towards the URL that is processed by this URL-to-item-resolver. The ECMAScript specification for regular expressions (ECMA-262) is used for the evaluation. The named groups defined by the regular expression will be used later to extract values from the URL.
func (m *UrlMatchInfo) GetUrlPattern()(*string) {
    return m.urlPattern
}
// Serialize serializes information the current object
func (m *UrlMatchInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBaseUrls() != nil {
        err := writer.WriteCollectionOfStringValues("baseUrls", m.GetBaseUrls())
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
        err := writer.WriteStringValue("urlPattern", m.GetUrlPattern())
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
func (m *UrlMatchInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBaseUrls sets the baseUrls property value. A list of the URL prefixes that must match URLs to be processed by this URL-to-item-resolver.
func (m *UrlMatchInfo) SetBaseUrls(value []string)() {
    m.baseUrls = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UrlMatchInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUrlPattern sets the urlPattern property value. A regular expression that will be matched towards the URL that is processed by this URL-to-item-resolver. The ECMAScript specification for regular expressions (ECMA-262) is used for the evaluation. The named groups defined by the regular expression will be used later to extract values from the URL.
func (m *UrlMatchInfo) SetUrlPattern(value *string)() {
    m.urlPattern = value
}
