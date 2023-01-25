package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterUsageSummary 
type PrinterUsageSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The completedJobCount property
    completedJobCount *int32
    // The incompleteJobCount property
    incompleteJobCount *int32
    // The OdataType property
    odataType *string
    // The printer property
    printer DirectoryObjectable
    // The printerDisplayName property
    printerDisplayName *string
    // The printerId property
    printerId *string
    // The printerManufacturer property
    printerManufacturer *string
    // The printerModel property
    printerModel *string
}
// NewPrinterUsageSummary instantiates a new printerUsageSummary and sets the default values.
func NewPrinterUsageSummary()(*PrinterUsageSummary) {
    m := &PrinterUsageSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrinterUsageSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterUsageSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterUsageSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterUsageSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompletedJobCount gets the completedJobCount property value. The completedJobCount property
func (m *PrinterUsageSummary) GetCompletedJobCount()(*int32) {
    return m.completedJobCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterUsageSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedJobCount(val)
        }
        return nil
    }
    res["incompleteJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteJobCount(val)
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
    res["printer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinter(val.(DirectoryObjectable))
        }
        return nil
    }
    res["printerDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterDisplayName(val)
        }
        return nil
    }
    res["printerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterId(val)
        }
        return nil
    }
    res["printerManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterManufacturer(val)
        }
        return nil
    }
    res["printerModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterModel(val)
        }
        return nil
    }
    return res
}
// GetIncompleteJobCount gets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrinterUsageSummary) GetIncompleteJobCount()(*int32) {
    return m.incompleteJobCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterUsageSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetPrinter gets the printer property value. The printer property
func (m *PrinterUsageSummary) GetPrinter()(DirectoryObjectable) {
    return m.printer
}
// GetPrinterDisplayName gets the printerDisplayName property value. The printerDisplayName property
func (m *PrinterUsageSummary) GetPrinterDisplayName()(*string) {
    return m.printerDisplayName
}
// GetPrinterId gets the printerId property value. The printerId property
func (m *PrinterUsageSummary) GetPrinterId()(*string) {
    return m.printerId
}
// GetPrinterManufacturer gets the printerManufacturer property value. The printerManufacturer property
func (m *PrinterUsageSummary) GetPrinterManufacturer()(*string) {
    return m.printerManufacturer
}
// GetPrinterModel gets the printerModel property value. The printerModel property
func (m *PrinterUsageSummary) GetPrinterModel()(*string) {
    return m.printerModel
}
// Serialize serializes information the current object
func (m *PrinterUsageSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("completedJobCount", m.GetCompletedJobCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("incompleteJobCount", m.GetIncompleteJobCount())
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
        err := writer.WriteObjectValue("printer", m.GetPrinter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerDisplayName", m.GetPrinterDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerId", m.GetPrinterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerManufacturer", m.GetPrinterManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerModel", m.GetPrinterModel())
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
func (m *PrinterUsageSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompletedJobCount sets the completedJobCount property value. The completedJobCount property
func (m *PrinterUsageSummary) SetCompletedJobCount(value *int32)() {
    m.completedJobCount = value
}
// SetIncompleteJobCount sets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrinterUsageSummary) SetIncompleteJobCount(value *int32)() {
    m.incompleteJobCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterUsageSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrinter sets the printer property value. The printer property
func (m *PrinterUsageSummary) SetPrinter(value DirectoryObjectable)() {
    m.printer = value
}
// SetPrinterDisplayName sets the printerDisplayName property value. The printerDisplayName property
func (m *PrinterUsageSummary) SetPrinterDisplayName(value *string)() {
    m.printerDisplayName = value
}
// SetPrinterId sets the printerId property value. The printerId property
func (m *PrinterUsageSummary) SetPrinterId(value *string)() {
    m.printerId = value
}
// SetPrinterManufacturer sets the printerManufacturer property value. The printerManufacturer property
func (m *PrinterUsageSummary) SetPrinterManufacturer(value *string)() {
    m.printerManufacturer = value
}
// SetPrinterModel sets the printerModel property value. The printerModel property
func (m *PrinterUsageSummary) SetPrinterModel(value *string)() {
    m.printerModel = value
}
