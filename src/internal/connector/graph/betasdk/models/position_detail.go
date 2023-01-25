package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PositionDetail 
type PositionDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Detail about the company or employer.
    company CompanyDetailable
    // Description of the position in question.
    description *string
    // When the position ended.
    endMonthYear *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The title held when in that position.
    jobTitle *string
    // The OdataType property
    odataType *string
    // The role the position entailed.
    role *string
    // The start month and year of the position.
    startMonthYear *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Short summary of the position.
    summary *string
}
// NewPositionDetail instantiates a new positionDetail and sets the default values.
func NewPositionDetail()(*PositionDetail) {
    m := &PositionDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePositionDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePositionDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPositionDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PositionDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompany gets the company property value. Detail about the company or employer.
func (m *PositionDetail) GetCompany()(CompanyDetailable) {
    return m.company
}
// GetDescription gets the description property value. Description of the position in question.
func (m *PositionDetail) GetDescription()(*string) {
    return m.description
}
// GetEndMonthYear gets the endMonthYear property value. When the position ended.
func (m *PositionDetail) GetEndMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endMonthYear
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PositionDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["company"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCompanyDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val.(CompanyDetailable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["endMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonthYear(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
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
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["startMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonthYear(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    return res
}
// GetJobTitle gets the jobTitle property value. The title held when in that position.
func (m *PositionDetail) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PositionDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetRole gets the role property value. The role the position entailed.
func (m *PositionDetail) GetRole()(*string) {
    return m.role
}
// GetStartMonthYear gets the startMonthYear property value. The start month and year of the position.
func (m *PositionDetail) GetStartMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.startMonthYear
}
// GetSummary gets the summary property value. Short summary of the position.
func (m *PositionDetail) GetSummary()(*string) {
    return m.summary
}
// Serialize serializes information the current object
func (m *PositionDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("endMonthYear", m.GetEndMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jobTitle", m.GetJobTitle())
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
        err := writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("startMonthYear", m.GetStartMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
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
func (m *PositionDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompany sets the company property value. Detail about the company or employer.
func (m *PositionDetail) SetCompany(value CompanyDetailable)() {
    m.company = value
}
// SetDescription sets the description property value. Description of the position in question.
func (m *PositionDetail) SetDescription(value *string)() {
    m.description = value
}
// SetEndMonthYear sets the endMonthYear property value. When the position ended.
func (m *PositionDetail) SetEndMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endMonthYear = value
}
// SetJobTitle sets the jobTitle property value. The title held when in that position.
func (m *PositionDetail) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PositionDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRole sets the role property value. The role the position entailed.
func (m *PositionDetail) SetRole(value *string)() {
    m.role = value
}
// SetStartMonthYear sets the startMonthYear property value. The start month and year of the position.
func (m *PositionDetail) SetStartMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.startMonthYear = value
}
// SetSummary sets the summary property value. Short summary of the position.
func (m *PositionDetail) SetSummary(value *string)() {
    m.summary = value
}
