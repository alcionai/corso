package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Currency provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Currency struct {
    Entity
    // The amountDecimalPlaces property
    amountDecimalPlaces *string
    // The amountRoundingPrecision property
    amountRoundingPrecision *float64
    // The code property
    code *string
    // The displayName property
    displayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The symbol property
    symbol *string
}
// NewCurrency instantiates a new currency and sets the default values.
func NewCurrency()(*Currency) {
    m := &Currency{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCurrencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCurrencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCurrency(), nil
}
// GetAmountDecimalPlaces gets the amountDecimalPlaces property value. The amountDecimalPlaces property
func (m *Currency) GetAmountDecimalPlaces()(*string) {
    return m.amountDecimalPlaces
}
// GetAmountRoundingPrecision gets the amountRoundingPrecision property value. The amountRoundingPrecision property
func (m *Currency) GetAmountRoundingPrecision()(*float64) {
    return m.amountRoundingPrecision
}
// GetCode gets the code property value. The code property
func (m *Currency) GetCode()(*string) {
    return m.code
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Currency) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Currency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["amountDecimalPlaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountDecimalPlaces(val)
        }
        return nil
    }
    res["amountRoundingPrecision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountRoundingPrecision(val)
        }
        return nil
    }
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["symbol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSymbol(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Currency) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetSymbol gets the symbol property value. The symbol property
func (m *Currency) GetSymbol()(*string) {
    return m.symbol
}
// Serialize serializes information the current object
func (m *Currency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("amountDecimalPlaces", m.GetAmountDecimalPlaces())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("amountRoundingPrecision", m.GetAmountRoundingPrecision())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("symbol", m.GetSymbol())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAmountDecimalPlaces sets the amountDecimalPlaces property value. The amountDecimalPlaces property
func (m *Currency) SetAmountDecimalPlaces(value *string)() {
    m.amountDecimalPlaces = value
}
// SetAmountRoundingPrecision sets the amountRoundingPrecision property value. The amountRoundingPrecision property
func (m *Currency) SetAmountRoundingPrecision(value *float64)() {
    m.amountRoundingPrecision = value
}
// SetCode sets the code property value. The code property
func (m *Currency) SetCode(value *string)() {
    m.code = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Currency) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Currency) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetSymbol sets the symbol property value. The symbol property
func (m *Currency) SetSymbol(value *string)() {
    m.symbol = value
}
