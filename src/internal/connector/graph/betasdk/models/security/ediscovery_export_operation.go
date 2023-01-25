package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryExportOperation 
type EdiscoveryExportOperation struct {
    CaseOperation
    // The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
    azureBlobContainer *string
    // The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
    azureBlobToken *string
    // The description provided for the export.
    description *string
    // The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
    exportOptions *ExportOptions
    // The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
    exportStructure *ExportFileStructure
    // The outputFolderId property
    outputFolderId *string
    // The name provided for the export.
    outputName *string
    // Review set from where documents are exported.
    reviewSet EdiscoveryReviewSetable
    // The review set query which is used to filter the documents for export.
    reviewSetQuery EdiscoveryReviewSetQueryable
}
// NewEdiscoveryExportOperation instantiates a new EdiscoveryExportOperation and sets the default values.
func NewEdiscoveryExportOperation()(*EdiscoveryExportOperation) {
    m := &EdiscoveryExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryExportOperation(), nil
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
func (m *EdiscoveryExportOperation) GetAzureBlobContainer()(*string) {
    return m.azureBlobContainer
}
// GetAzureBlobToken gets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *EdiscoveryExportOperation) GetAzureBlobToken()(*string) {
    return m.azureBlobToken
}
// GetDescription gets the description property value. The description provided for the export.
func (m *EdiscoveryExportOperation) GetDescription()(*string) {
    return m.description
}
// GetExportOptions gets the exportOptions property value. The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
func (m *EdiscoveryExportOperation) GetExportOptions()(*ExportOptions) {
    return m.exportOptions
}
// GetExportStructure gets the exportStructure property value. The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
func (m *EdiscoveryExportOperation) GetExportStructure()(*ExportFileStructure) {
    return m.exportStructure
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["azureBlobContainer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureBlobContainer(val)
        }
        return nil
    }
    res["azureBlobToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureBlobToken(val)
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
    res["exportOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportOptions(val.(*ExportOptions))
        }
        return nil
    }
    res["exportStructure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportFileStructure)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportStructure(val.(*ExportFileStructure))
        }
        return nil
    }
    res["outputFolderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputFolderId(val)
        }
        return nil
    }
    res["outputName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputName(val)
        }
        return nil
    }
    res["reviewSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(EdiscoveryReviewSetable))
        }
        return nil
    }
    res["reviewSetQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSetQuery(val.(EdiscoveryReviewSetQueryable))
        }
        return nil
    }
    return res
}
// GetOutputFolderId gets the outputFolderId property value. The outputFolderId property
func (m *EdiscoveryExportOperation) GetOutputFolderId()(*string) {
    return m.outputFolderId
}
// GetOutputName gets the outputName property value. The name provided for the export.
func (m *EdiscoveryExportOperation) GetOutputName()(*string) {
    return m.outputName
}
// GetReviewSet gets the reviewSet property value. Review set from where documents are exported.
func (m *EdiscoveryExportOperation) GetReviewSet()(EdiscoveryReviewSetable) {
    return m.reviewSet
}
// GetReviewSetQuery gets the reviewSetQuery property value. The review set query which is used to filter the documents for export.
func (m *EdiscoveryExportOperation) GetReviewSetQuery()(EdiscoveryReviewSetQueryable) {
    return m.reviewSetQuery
}
// Serialize serializes information the current object
func (m *EdiscoveryExportOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureBlobContainer", m.GetAzureBlobContainer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureBlobToken", m.GetAzureBlobToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetExportOptions() != nil {
        cast := (*m.GetExportOptions()).String()
        err = writer.WriteStringValue("exportOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportStructure() != nil {
        cast := (*m.GetExportStructure()).String()
        err = writer.WriteStringValue("exportStructure", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outputFolderId", m.GetOutputFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outputName", m.GetOutputName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSet", m.GetReviewSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSetQuery", m.GetReviewSetQuery())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
func (m *EdiscoveryExportOperation) SetAzureBlobContainer(value *string)() {
    m.azureBlobContainer = value
}
// SetAzureBlobToken sets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *EdiscoveryExportOperation) SetAzureBlobToken(value *string)() {
    m.azureBlobToken = value
}
// SetDescription sets the description property value. The description provided for the export.
func (m *EdiscoveryExportOperation) SetDescription(value *string)() {
    m.description = value
}
// SetExportOptions sets the exportOptions property value. The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
func (m *EdiscoveryExportOperation) SetExportOptions(value *ExportOptions)() {
    m.exportOptions = value
}
// SetExportStructure sets the exportStructure property value. The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
func (m *EdiscoveryExportOperation) SetExportStructure(value *ExportFileStructure)() {
    m.exportStructure = value
}
// SetOutputFolderId sets the outputFolderId property value. The outputFolderId property
func (m *EdiscoveryExportOperation) SetOutputFolderId(value *string)() {
    m.outputFolderId = value
}
// SetOutputName sets the outputName property value. The name provided for the export.
func (m *EdiscoveryExportOperation) SetOutputName(value *string)() {
    m.outputName = value
}
// SetReviewSet sets the reviewSet property value. Review set from where documents are exported.
func (m *EdiscoveryExportOperation) SetReviewSet(value EdiscoveryReviewSetable)() {
    m.reviewSet = value
}
// SetReviewSetQuery sets the reviewSetQuery property value. The review set query which is used to filter the documents for export.
func (m *EdiscoveryExportOperation) SetReviewSetQuery(value EdiscoveryReviewSetQueryable)() {
    m.reviewSetQuery = value
}
