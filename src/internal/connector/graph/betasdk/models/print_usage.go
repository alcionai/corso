package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintUsage provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PrintUsage struct {
    Entity
    // The blackAndWhitePageCount property
    blackAndWhitePageCount *int64
    // The colorPageCount property
    colorPageCount *int64
    // The completedBlackAndWhiteJobCount property
    completedBlackAndWhiteJobCount *int64
    // The completedColorJobCount property
    completedColorJobCount *int64
    // The completedJobCount property
    completedJobCount *int64
    // The doubleSidedSheetCount property
    doubleSidedSheetCount *int64
    // The incompleteJobCount property
    incompleteJobCount *int64
    // The mediaSheetCount property
    mediaSheetCount *int64
    // The pageCount property
    pageCount *int64
    // The singleSidedSheetCount property
    singleSidedSheetCount *int64
    // The usageDate property
    usageDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewPrintUsage instantiates a new printUsage and sets the default values.
func NewPrintUsage()(*PrintUsage) {
    m := &PrintUsage{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.printUsageByPrinter":
                        return NewPrintUsageByPrinter(), nil
                    case "#microsoft.graph.printUsageByUser":
                        return NewPrintUsageByUser(), nil
                }
            }
        }
    }
    return NewPrintUsage(), nil
}
// GetBlackAndWhitePageCount gets the blackAndWhitePageCount property value. The blackAndWhitePageCount property
func (m *PrintUsage) GetBlackAndWhitePageCount()(*int64) {
    return m.blackAndWhitePageCount
}
// GetColorPageCount gets the colorPageCount property value. The colorPageCount property
func (m *PrintUsage) GetColorPageCount()(*int64) {
    return m.colorPageCount
}
// GetCompletedBlackAndWhiteJobCount gets the completedBlackAndWhiteJobCount property value. The completedBlackAndWhiteJobCount property
func (m *PrintUsage) GetCompletedBlackAndWhiteJobCount()(*int64) {
    return m.completedBlackAndWhiteJobCount
}
// GetCompletedColorJobCount gets the completedColorJobCount property value. The completedColorJobCount property
func (m *PrintUsage) GetCompletedColorJobCount()(*int64) {
    return m.completedColorJobCount
}
// GetCompletedJobCount gets the completedJobCount property value. The completedJobCount property
func (m *PrintUsage) GetCompletedJobCount()(*int64) {
    return m.completedJobCount
}
// GetDoubleSidedSheetCount gets the doubleSidedSheetCount property value. The doubleSidedSheetCount property
func (m *PrintUsage) GetDoubleSidedSheetCount()(*int64) {
    return m.doubleSidedSheetCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["blackAndWhitePageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlackAndWhitePageCount(val)
        }
        return nil
    }
    res["colorPageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorPageCount(val)
        }
        return nil
    }
    res["completedBlackAndWhiteJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedBlackAndWhiteJobCount(val)
        }
        return nil
    }
    res["completedColorJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedColorJobCount(val)
        }
        return nil
    }
    res["completedJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedJobCount(val)
        }
        return nil
    }
    res["doubleSidedSheetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDoubleSidedSheetCount(val)
        }
        return nil
    }
    res["incompleteJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteJobCount(val)
        }
        return nil
    }
    res["mediaSheetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaSheetCount(val)
        }
        return nil
    }
    res["pageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageCount(val)
        }
        return nil
    }
    res["singleSidedSheetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSidedSheetCount(val)
        }
        return nil
    }
    res["usageDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageDate(val)
        }
        return nil
    }
    return res
}
// GetIncompleteJobCount gets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrintUsage) GetIncompleteJobCount()(*int64) {
    return m.incompleteJobCount
}
// GetMediaSheetCount gets the mediaSheetCount property value. The mediaSheetCount property
func (m *PrintUsage) GetMediaSheetCount()(*int64) {
    return m.mediaSheetCount
}
// GetPageCount gets the pageCount property value. The pageCount property
func (m *PrintUsage) GetPageCount()(*int64) {
    return m.pageCount
}
// GetSingleSidedSheetCount gets the singleSidedSheetCount property value. The singleSidedSheetCount property
func (m *PrintUsage) GetSingleSidedSheetCount()(*int64) {
    return m.singleSidedSheetCount
}
// GetUsageDate gets the usageDate property value. The usageDate property
func (m *PrintUsage) GetUsageDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.usageDate
}
// Serialize serializes information the current object
func (m *PrintUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("blackAndWhitePageCount", m.GetBlackAndWhitePageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("colorPageCount", m.GetColorPageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("completedBlackAndWhiteJobCount", m.GetCompletedBlackAndWhiteJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("completedColorJobCount", m.GetCompletedColorJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("completedJobCount", m.GetCompletedJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("doubleSidedSheetCount", m.GetDoubleSidedSheetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("incompleteJobCount", m.GetIncompleteJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("mediaSheetCount", m.GetMediaSheetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("pageCount", m.GetPageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("singleSidedSheetCount", m.GetSingleSidedSheetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("usageDate", m.GetUsageDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBlackAndWhitePageCount sets the blackAndWhitePageCount property value. The blackAndWhitePageCount property
func (m *PrintUsage) SetBlackAndWhitePageCount(value *int64)() {
    m.blackAndWhitePageCount = value
}
// SetColorPageCount sets the colorPageCount property value. The colorPageCount property
func (m *PrintUsage) SetColorPageCount(value *int64)() {
    m.colorPageCount = value
}
// SetCompletedBlackAndWhiteJobCount sets the completedBlackAndWhiteJobCount property value. The completedBlackAndWhiteJobCount property
func (m *PrintUsage) SetCompletedBlackAndWhiteJobCount(value *int64)() {
    m.completedBlackAndWhiteJobCount = value
}
// SetCompletedColorJobCount sets the completedColorJobCount property value. The completedColorJobCount property
func (m *PrintUsage) SetCompletedColorJobCount(value *int64)() {
    m.completedColorJobCount = value
}
// SetCompletedJobCount sets the completedJobCount property value. The completedJobCount property
func (m *PrintUsage) SetCompletedJobCount(value *int64)() {
    m.completedJobCount = value
}
// SetDoubleSidedSheetCount sets the doubleSidedSheetCount property value. The doubleSidedSheetCount property
func (m *PrintUsage) SetDoubleSidedSheetCount(value *int64)() {
    m.doubleSidedSheetCount = value
}
// SetIncompleteJobCount sets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrintUsage) SetIncompleteJobCount(value *int64)() {
    m.incompleteJobCount = value
}
// SetMediaSheetCount sets the mediaSheetCount property value. The mediaSheetCount property
func (m *PrintUsage) SetMediaSheetCount(value *int64)() {
    m.mediaSheetCount = value
}
// SetPageCount sets the pageCount property value. The pageCount property
func (m *PrintUsage) SetPageCount(value *int64)() {
    m.pageCount = value
}
// SetSingleSidedSheetCount sets the singleSidedSheetCount property value. The singleSidedSheetCount property
func (m *PrintUsage) SetSingleSidedSheetCount(value *int64)() {
    m.singleSidedSheetCount = value
}
// SetUsageDate sets the usageDate property value. The usageDate property
func (m *PrintUsage) SetUsageDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.usageDate = value
}
