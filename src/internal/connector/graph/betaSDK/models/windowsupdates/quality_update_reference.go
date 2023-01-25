package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// QualityUpdateReference 
type QualityUpdateReference struct {
    WindowsUpdateReference
    // Specifies the classification of the referenced content. Supports a subset of the values for qualityUpdateClassification. Possible values are: security, unknownFutureValue.
    classification *QualityUpdateClassification
    // Specifies a quality update in the given servicingChannel with the given classification by date (i.e. the last update published on the specified date). Default value is security.
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewQualityUpdateReference instantiates a new QualityUpdateReference and sets the default values.
func NewQualityUpdateReference()(*QualityUpdateReference) {
    m := &QualityUpdateReference{
        WindowsUpdateReference: *NewWindowsUpdateReference(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.qualityUpdateReference";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateQualityUpdateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQualityUpdateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsUpdates.expeditedQualityUpdateReference":
                        return NewExpeditedQualityUpdateReference(), nil
                }
            }
        }
    }
    return NewQualityUpdateReference(), nil
}
// GetClassification gets the classification property value. Specifies the classification of the referenced content. Supports a subset of the values for qualityUpdateClassification. Possible values are: security, unknownFutureValue.
func (m *QualityUpdateReference) GetClassification()(*QualityUpdateClassification) {
    return m.classification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *QualityUpdateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateReference.GetFieldDeserializers()
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQualityUpdateClassification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*QualityUpdateClassification))
        }
        return nil
    }
    res["releaseDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    return res
}
// GetReleaseDateTime gets the releaseDateTime property value. Specifies a quality update in the given servicingChannel with the given classification by date (i.e. the last update published on the specified date). Default value is security.
func (m *QualityUpdateReference) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.releaseDateTime
}
// Serialize serializes information the current object
func (m *QualityUpdateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateReference.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassification sets the classification property value. Specifies the classification of the referenced content. Supports a subset of the values for qualityUpdateClassification. Possible values are: security, unknownFutureValue.
func (m *QualityUpdateReference) SetClassification(value *QualityUpdateClassification)() {
    m.classification = value
}
// SetReleaseDateTime sets the releaseDateTime property value. Specifies a quality update in the given servicingChannel with the given classification by date (i.e. the last update published on the specified date). Default value is security.
func (m *QualityUpdateReference) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releaseDateTime = value
}
