package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcConnectivityResult 
type CloudPcConnectivityResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A list of failed health check items. If the status property is available, this property will be empty.
    failedHealthCheckItems []CloudPcHealthCheckItemable
    // The OdataType property
    odataType *string
    // The status property
    status *CloudPcConnectivityStatus
    // Datetime when the status was updated. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
    updatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCloudPcConnectivityResult instantiates a new cloudPcConnectivityResult and sets the default values.
func NewCloudPcConnectivityResult()(*CloudPcConnectivityResult) {
    m := &CloudPcConnectivityResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcConnectivityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcConnectivityResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcConnectivityResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcConnectivityResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFailedHealthCheckItems gets the failedHealthCheckItems property value. A list of failed health check items. If the status property is available, this property will be empty.
func (m *CloudPcConnectivityResult) GetFailedHealthCheckItems()([]CloudPcHealthCheckItemable) {
    return m.failedHealthCheckItems
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcConnectivityResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedHealthCheckItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcHealthCheckItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcHealthCheckItemable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcHealthCheckItemable)
            }
            m.SetFailedHealthCheckItems(res)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcConnectivityStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcConnectivityStatus))
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcConnectivityResult) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status property
func (m *CloudPcConnectivityResult) GetStatus()(*CloudPcConnectivityStatus) {
    return m.status
}
// GetUpdatedDateTime gets the updatedDateTime property value. Datetime when the status was updated. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcConnectivityResult) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedDateTime
}
// Serialize serializes information the current object
func (m *CloudPcConnectivityResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFailedHealthCheckItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFailedHealthCheckItems()))
        for i, v := range m.GetFailedHealthCheckItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("failedHealthCheckItems", cast)
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
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
func (m *CloudPcConnectivityResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFailedHealthCheckItems sets the failedHealthCheckItems property value. A list of failed health check items. If the status property is available, this property will be empty.
func (m *CloudPcConnectivityResult) SetFailedHealthCheckItems(value []CloudPcHealthCheckItemable)() {
    m.failedHealthCheckItems = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcConnectivityResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status property
func (m *CloudPcConnectivityResult) SetStatus(value *CloudPcConnectivityStatus)() {
    m.status = value
}
// SetUpdatedDateTime sets the updatedDateTime property value. Datetime when the status was updated. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcConnectivityResult) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedDateTime = value
}
