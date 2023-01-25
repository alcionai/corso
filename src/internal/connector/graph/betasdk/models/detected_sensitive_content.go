package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedSensitiveContent 
type DetectedSensitiveContent struct {
    DetectedSensitiveContentBase
    // The classificationAttributes property
    classificationAttributes []ClassificationAttributeable
    // The classificationMethod property
    classificationMethod *ClassificationMethod
    // The matches property
    matches []SensitiveContentLocationable
    // The scope property
    scope *SensitiveTypeScope
    // The sensitiveTypeSource property
    sensitiveTypeSource *SensitiveTypeSource
}
// NewDetectedSensitiveContent instantiates a new DetectedSensitiveContent and sets the default values.
func NewDetectedSensitiveContent()(*DetectedSensitiveContent) {
    m := &DetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// CreateDetectedSensitiveContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDetectedSensitiveContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.machineLearningDetectedSensitiveContent":
                        return NewMachineLearningDetectedSensitiveContent(), nil
                }
            }
        }
    }
    return NewDetectedSensitiveContent(), nil
}
// GetClassificationAttributes gets the classificationAttributes property value. The classificationAttributes property
func (m *DetectedSensitiveContent) GetClassificationAttributes()([]ClassificationAttributeable) {
    return m.classificationAttributes
}
// GetClassificationMethod gets the classificationMethod property value. The classificationMethod property
func (m *DetectedSensitiveContent) GetClassificationMethod()(*ClassificationMethod) {
    return m.classificationMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedSensitiveContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["classificationAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassificationAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationAttributeable, len(val))
            for i, v := range val {
                res[i] = v.(ClassificationAttributeable)
            }
            m.SetClassificationAttributes(res)
        }
        return nil
    }
    res["classificationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClassificationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassificationMethod(val.(*ClassificationMethod))
        }
        return nil
    }
    res["matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitiveContentLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveContentLocationable, len(val))
            for i, v := range val {
                res[i] = v.(SensitiveContentLocationable)
            }
            m.SetMatches(res)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*SensitiveTypeScope))
        }
        return nil
    }
    res["sensitiveTypeSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitiveTypeSource(val.(*SensitiveTypeSource))
        }
        return nil
    }
    return res
}
// GetMatches gets the matches property value. The matches property
func (m *DetectedSensitiveContent) GetMatches()([]SensitiveContentLocationable) {
    return m.matches
}
// GetScope gets the scope property value. The scope property
func (m *DetectedSensitiveContent) GetScope()(*SensitiveTypeScope) {
    return m.scope
}
// GetSensitiveTypeSource gets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *DetectedSensitiveContent) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    return m.sensitiveTypeSource
}
// Serialize serializes information the current object
func (m *DetectedSensitiveContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassificationAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassificationAttributes()))
        for i, v := range m.GetClassificationAttributes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classificationAttributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassificationMethod() != nil {
        cast := (*m.GetClassificationMethod()).String()
        err = writer.WriteStringValue("classificationMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := (*m.GetScope()).String()
        err = writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeSource() != nil {
        cast := (*m.GetSensitiveTypeSource()).String()
        err = writer.WriteStringValue("sensitiveTypeSource", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassificationAttributes sets the classificationAttributes property value. The classificationAttributes property
func (m *DetectedSensitiveContent) SetClassificationAttributes(value []ClassificationAttributeable)() {
    m.classificationAttributes = value
}
// SetClassificationMethod sets the classificationMethod property value. The classificationMethod property
func (m *DetectedSensitiveContent) SetClassificationMethod(value *ClassificationMethod)() {
    m.classificationMethod = value
}
// SetMatches sets the matches property value. The matches property
func (m *DetectedSensitiveContent) SetMatches(value []SensitiveContentLocationable)() {
    m.matches = value
}
// SetScope sets the scope property value. The scope property
func (m *DetectedSensitiveContent) SetScope(value *SensitiveTypeScope)() {
    m.scope = value
}
// SetSensitiveTypeSource sets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *DetectedSensitiveContent) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    m.sensitiveTypeSource = value
}
