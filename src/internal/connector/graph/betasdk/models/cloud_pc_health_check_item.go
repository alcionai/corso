package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcHealthCheckItem 
type CloudPcHealthCheckItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Additional message for this health check.
    additionalDetails *string
    // The connectivity health check item name.
    displayName *string
    // Timestamp when the last check occurs. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
    lastHealthCheckDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The result property
    result *CloudPcConnectivityEventResult
}
// NewCloudPcHealthCheckItem instantiates a new cloudPcHealthCheckItem and sets the default values.
func NewCloudPcHealthCheckItem()(*CloudPcHealthCheckItem) {
    m := &CloudPcHealthCheckItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcHealthCheckItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcHealthCheckItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcHealthCheckItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcHealthCheckItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdditionalDetails gets the additionalDetails property value. Additional message for this health check.
func (m *CloudPcHealthCheckItem) GetAdditionalDetails()(*string) {
    return m.additionalDetails
}
// GetDisplayName gets the displayName property value. The connectivity health check item name.
func (m *CloudPcHealthCheckItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcHealthCheckItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetails(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastHealthCheckDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHealthCheckDateTime(val)
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
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcConnectivityEventResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*CloudPcConnectivityEventResult))
        }
        return nil
    }
    return res
}
// GetLastHealthCheckDateTime gets the lastHealthCheckDateTime property value. Timestamp when the last check occurs. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcHealthCheckItem) GetLastHealthCheckDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastHealthCheckDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcHealthCheckItem) GetOdataType()(*string) {
    return m.odataType
}
// GetResult gets the result property value. The result property
func (m *CloudPcHealthCheckItem) GetResult()(*CloudPcConnectivityEventResult) {
    return m.result
}
// Serialize serializes information the current object
func (m *CloudPcHealthCheckItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastHealthCheckDateTime", m.GetLastHealthCheckDateTime())
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
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err := writer.WriteStringValue("result", &cast)
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
func (m *CloudPcHealthCheckItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdditionalDetails sets the additionalDetails property value. Additional message for this health check.
func (m *CloudPcHealthCheckItem) SetAdditionalDetails(value *string)() {
    m.additionalDetails = value
}
// SetDisplayName sets the displayName property value. The connectivity health check item name.
func (m *CloudPcHealthCheckItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastHealthCheckDateTime sets the lastHealthCheckDateTime property value. Timestamp when the last check occurs. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcHealthCheckItem) SetLastHealthCheckDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHealthCheckDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcHealthCheckItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResult sets the result property value. The result property
func (m *CloudPcHealthCheckItem) SetResult(value *CloudPcConnectivityEventResult)() {
    m.result = value
}
