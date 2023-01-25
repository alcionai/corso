package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsMinimumOperatingSystem the minimum operating system required for a Windows mobile app.
type WindowsMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Windows version 10.0 or later.
    v10_0 *bool
    // Windows 10 1607 or later.
    v10_1607 *bool
    // Windows 10 1703 or later.
    v10_1703 *bool
    // Windows 10 1709 or later.
    v10_1709 *bool
    // Windows 10 1803 or later.
    v10_1803 *bool
    // Windows 10 1809 or later.
    v10_1809 *bool
    // Windows 10 1903 or later.
    v10_1903 *bool
    // Windows 10 1909 or later.
    v10_1909 *bool
    // Windows 10 2004 or later.
    v10_2004 *bool
    // Windows 10 21H1 or later.
    v10_21H1 *bool
    // Windows 10 2H20 or later.
    v10_2H20 *bool
    // Windows version 8.0 or later.
    v8_0 *bool
    // Windows version 8.1 or later.
    v8_1 *bool
}
// NewWindowsMinimumOperatingSystem instantiates a new windowsMinimumOperatingSystem and sets the default values.
func NewWindowsMinimumOperatingSystem()(*WindowsMinimumOperatingSystem) {
    m := &WindowsMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsMinimumOperatingSystem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["v10_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_0(val)
        }
        return nil
    }
    res["v10_1607"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_1607(val)
        }
        return nil
    }
    res["v10_1703"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_1703(val)
        }
        return nil
    }
    res["v10_1709"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_1709(val)
        }
        return nil
    }
    res["v10_1803"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_1803(val)
        }
        return nil
    }
    res["v10_1809"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_1809(val)
        }
        return nil
    }
    res["v10_1903"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_1903(val)
        }
        return nil
    }
    res["v10_1909"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_1909(val)
        }
        return nil
    }
    res["v10_2004"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_2004(val)
        }
        return nil
    }
    res["v10_21H1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_21H1(val)
        }
        return nil
    }
    res["v10_2H20"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_2H20(val)
        }
        return nil
    }
    res["v8_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV8_0(val)
        }
        return nil
    }
    res["v8_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV8_1(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV10_0 gets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_0()(*bool) {
    return m.v10_0
}
// GetV10_1607 gets the v10_1607 property value. Windows 10 1607 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_1607()(*bool) {
    return m.v10_1607
}
// GetV10_1703 gets the v10_1703 property value. Windows 10 1703 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_1703()(*bool) {
    return m.v10_1703
}
// GetV10_1709 gets the v10_1709 property value. Windows 10 1709 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_1709()(*bool) {
    return m.v10_1709
}
// GetV10_1803 gets the v10_1803 property value. Windows 10 1803 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_1803()(*bool) {
    return m.v10_1803
}
// GetV10_1809 gets the v10_1809 property value. Windows 10 1809 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_1809()(*bool) {
    return m.v10_1809
}
// GetV10_1903 gets the v10_1903 property value. Windows 10 1903 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_1903()(*bool) {
    return m.v10_1903
}
// GetV10_1909 gets the v10_1909 property value. Windows 10 1909 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_1909()(*bool) {
    return m.v10_1909
}
// GetV10_2004 gets the v10_2004 property value. Windows 10 2004 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_2004()(*bool) {
    return m.v10_2004
}
// GetV10_21H1 gets the v10_21H1 property value. Windows 10 21H1 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_21H1()(*bool) {
    return m.v10_21H1
}
// GetV10_2H20 gets the v10_2H20 property value. Windows 10 2H20 or later.
func (m *WindowsMinimumOperatingSystem) GetV10_2H20()(*bool) {
    return m.v10_2H20
}
// GetV8_0 gets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV8_0()(*bool) {
    return m.v8_0
}
// GetV8_1 gets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) GetV8_1()(*bool) {
    return m.v8_1
}
// Serialize serializes information the current object
func (m *WindowsMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_0", m.GetV10_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1607", m.GetV10_1607())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1703", m.GetV10_1703())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1709", m.GetV10_1709())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1803", m.GetV10_1803())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1809", m.GetV10_1809())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1903", m.GetV10_1903())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1909", m.GetV10_1909())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_2004", m.GetV10_2004())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_21H1", m.GetV10_21H1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_2H20", m.GetV10_2H20())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_0", m.GetV8_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_1", m.GetV8_1())
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
func (m *WindowsMinimumOperatingSystem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV10_0 sets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_0(value *bool)() {
    m.v10_0 = value
}
// SetV10_1607 sets the v10_1607 property value. Windows 10 1607 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_1607(value *bool)() {
    m.v10_1607 = value
}
// SetV10_1703 sets the v10_1703 property value. Windows 10 1703 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_1703(value *bool)() {
    m.v10_1703 = value
}
// SetV10_1709 sets the v10_1709 property value. Windows 10 1709 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_1709(value *bool)() {
    m.v10_1709 = value
}
// SetV10_1803 sets the v10_1803 property value. Windows 10 1803 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_1803(value *bool)() {
    m.v10_1803 = value
}
// SetV10_1809 sets the v10_1809 property value. Windows 10 1809 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_1809(value *bool)() {
    m.v10_1809 = value
}
// SetV10_1903 sets the v10_1903 property value. Windows 10 1903 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_1903(value *bool)() {
    m.v10_1903 = value
}
// SetV10_1909 sets the v10_1909 property value. Windows 10 1909 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_1909(value *bool)() {
    m.v10_1909 = value
}
// SetV10_2004 sets the v10_2004 property value. Windows 10 2004 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_2004(value *bool)() {
    m.v10_2004 = value
}
// SetV10_21H1 sets the v10_21H1 property value. Windows 10 21H1 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_21H1(value *bool)() {
    m.v10_21H1 = value
}
// SetV10_2H20 sets the v10_2H20 property value. Windows 10 2H20 or later.
func (m *WindowsMinimumOperatingSystem) SetV10_2H20(value *bool)() {
    m.v10_2H20 = value
}
// SetV8_0 sets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV8_0(value *bool)() {
    m.v8_0 = value
}
// SetV8_1 sets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) SetV8_1(value *bool)() {
    m.v8_1 = value
}
