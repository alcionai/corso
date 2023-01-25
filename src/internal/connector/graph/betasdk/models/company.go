package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Company 
type Company struct {
    Entity
    // The accounts property
    accounts []Accountable
    // The agedAccountsPayable property
    agedAccountsPayable []AgedAccountsPayableable
    // The agedAccountsReceivable property
    agedAccountsReceivable []AgedAccountsReceivableable
    // The businessProfileId property
    businessProfileId *string
    // The companyInformation property
    companyInformation []CompanyInformationable
    // The countriesRegions property
    countriesRegions []CountryRegionable
    // The currencies property
    currencies []Currencyable
    // The customerPaymentJournals property
    customerPaymentJournals []CustomerPaymentJournalable
    // The customerPayments property
    customerPayments []CustomerPaymentable
    // The customers property
    customers []Customerable
    // The dimensions property
    dimensions []Dimensionable
    // The dimensionValues property
    dimensionValues []DimensionValueable
    // The displayName property
    displayName *string
    // The employees property
    employees []Employeeable
    // The generalLedgerEntries property
    generalLedgerEntries []GeneralLedgerEntryable
    // The itemCategories property
    itemCategories []ItemCategoryable
    // The items property
    items []Itemable
    // The journalLines property
    journalLines []JournalLineable
    // The journals property
    journals []Journalable
    // The name property
    name *string
    // The paymentMethods property
    paymentMethods []PaymentMethodable
    // The paymentTerms property
    paymentTerms []PaymentTermable
    // The picture property
    picture []Pictureable
    // The purchaseInvoiceLines property
    purchaseInvoiceLines []PurchaseInvoiceLineable
    // The purchaseInvoices property
    purchaseInvoices []PurchaseInvoiceable
    // The salesCreditMemoLines property
    salesCreditMemoLines []SalesCreditMemoLineable
    // The salesCreditMemos property
    salesCreditMemos []SalesCreditMemoable
    // The salesInvoiceLines property
    salesInvoiceLines []SalesInvoiceLineable
    // The salesInvoices property
    salesInvoices []SalesInvoiceable
    // The salesOrderLines property
    salesOrderLines []SalesOrderLineable
    // The salesOrders property
    salesOrders []SalesOrderable
    // The salesQuoteLines property
    salesQuoteLines []SalesQuoteLineable
    // The salesQuotes property
    salesQuotes []SalesQuoteable
    // The shipmentMethods property
    shipmentMethods []ShipmentMethodable
    // The systemVersion property
    systemVersion *string
    // The taxAreas property
    taxAreas []TaxAreaable
    // The taxGroups property
    taxGroups []TaxGroupable
    // The unitsOfMeasure property
    unitsOfMeasure []UnitOfMeasureable
    // The vendors property
    vendors []Vendor_escapedable
}
// NewCompany instantiates a new Company and sets the default values.
func NewCompany()(*Company) {
    m := &Company{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCompanyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompany(), nil
}
// GetAccounts gets the accounts property value. The accounts property
func (m *Company) GetAccounts()([]Accountable) {
    return m.accounts
}
// GetAgedAccountsPayable gets the agedAccountsPayable property value. The agedAccountsPayable property
func (m *Company) GetAgedAccountsPayable()([]AgedAccountsPayableable) {
    return m.agedAccountsPayable
}
// GetAgedAccountsReceivable gets the agedAccountsReceivable property value. The agedAccountsReceivable property
func (m *Company) GetAgedAccountsReceivable()([]AgedAccountsReceivableable) {
    return m.agedAccountsReceivable
}
// GetBusinessProfileId gets the businessProfileId property value. The businessProfileId property
func (m *Company) GetBusinessProfileId()(*string) {
    return m.businessProfileId
}
// GetCompanyInformation gets the companyInformation property value. The companyInformation property
func (m *Company) GetCompanyInformation()([]CompanyInformationable) {
    return m.companyInformation
}
// GetCountriesRegions gets the countriesRegions property value. The countriesRegions property
func (m *Company) GetCountriesRegions()([]CountryRegionable) {
    return m.countriesRegions
}
// GetCurrencies gets the currencies property value. The currencies property
func (m *Company) GetCurrencies()([]Currencyable) {
    return m.currencies
}
// GetCustomerPaymentJournals gets the customerPaymentJournals property value. The customerPaymentJournals property
func (m *Company) GetCustomerPaymentJournals()([]CustomerPaymentJournalable) {
    return m.customerPaymentJournals
}
// GetCustomerPayments gets the customerPayments property value. The customerPayments property
func (m *Company) GetCustomerPayments()([]CustomerPaymentable) {
    return m.customerPayments
}
// GetCustomers gets the customers property value. The customers property
func (m *Company) GetCustomers()([]Customerable) {
    return m.customers
}
// GetDimensions gets the dimensions property value. The dimensions property
func (m *Company) GetDimensions()([]Dimensionable) {
    return m.dimensions
}
// GetDimensionValues gets the dimensionValues property value. The dimensionValues property
func (m *Company) GetDimensionValues()([]DimensionValueable) {
    return m.dimensionValues
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Company) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmployees gets the employees property value. The employees property
func (m *Company) GetEmployees()([]Employeeable) {
    return m.employees
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Company) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Accountable, len(val))
            for i, v := range val {
                res[i] = v.(Accountable)
            }
            m.SetAccounts(res)
        }
        return nil
    }
    res["agedAccountsPayable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgedAccountsPayableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgedAccountsPayableable, len(val))
            for i, v := range val {
                res[i] = v.(AgedAccountsPayableable)
            }
            m.SetAgedAccountsPayable(res)
        }
        return nil
    }
    res["agedAccountsReceivable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgedAccountsReceivableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgedAccountsReceivableable, len(val))
            for i, v := range val {
                res[i] = v.(AgedAccountsReceivableable)
            }
            m.SetAgedAccountsReceivable(res)
        }
        return nil
    }
    res["businessProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessProfileId(val)
        }
        return nil
    }
    res["companyInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompanyInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanyInformationable, len(val))
            for i, v := range val {
                res[i] = v.(CompanyInformationable)
            }
            m.SetCompanyInformation(res)
        }
        return nil
    }
    res["countriesRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCountryRegionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CountryRegionable, len(val))
            for i, v := range val {
                res[i] = v.(CountryRegionable)
            }
            m.SetCountriesRegions(res)
        }
        return nil
    }
    res["currencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCurrencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Currencyable, len(val))
            for i, v := range val {
                res[i] = v.(Currencyable)
            }
            m.SetCurrencies(res)
        }
        return nil
    }
    res["customerPaymentJournals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerPaymentJournalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerPaymentJournalable, len(val))
            for i, v := range val {
                res[i] = v.(CustomerPaymentJournalable)
            }
            m.SetCustomerPaymentJournals(res)
        }
        return nil
    }
    res["customerPayments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerPaymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerPaymentable, len(val))
            for i, v := range val {
                res[i] = v.(CustomerPaymentable)
            }
            m.SetCustomerPayments(res)
        }
        return nil
    }
    res["customers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Customerable, len(val))
            for i, v := range val {
                res[i] = v.(Customerable)
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["dimensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDimensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Dimensionable, len(val))
            for i, v := range val {
                res[i] = v.(Dimensionable)
            }
            m.SetDimensions(res)
        }
        return nil
    }
    res["dimensionValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDimensionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DimensionValueable, len(val))
            for i, v := range val {
                res[i] = v.(DimensionValueable)
            }
            m.SetDimensionValues(res)
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
    res["employees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmployeeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Employeeable, len(val))
            for i, v := range val {
                res[i] = v.(Employeeable)
            }
            m.SetEmployees(res)
        }
        return nil
    }
    res["generalLedgerEntries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGeneralLedgerEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GeneralLedgerEntryable, len(val))
            for i, v := range val {
                res[i] = v.(GeneralLedgerEntryable)
            }
            m.SetGeneralLedgerEntries(res)
        }
        return nil
    }
    res["itemCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(ItemCategoryable)
            }
            m.SetItemCategories(res)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Itemable, len(val))
            for i, v := range val {
                res[i] = v.(Itemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["journalLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJournalLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JournalLineable, len(val))
            for i, v := range val {
                res[i] = v.(JournalLineable)
            }
            m.SetJournalLines(res)
        }
        return nil
    }
    res["journals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJournalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Journalable, len(val))
            for i, v := range val {
                res[i] = v.(Journalable)
            }
            m.SetJournals(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["paymentMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePaymentMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PaymentMethodable, len(val))
            for i, v := range val {
                res[i] = v.(PaymentMethodable)
            }
            m.SetPaymentMethods(res)
        }
        return nil
    }
    res["paymentTerms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePaymentTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PaymentTermable, len(val))
            for i, v := range val {
                res[i] = v.(PaymentTermable)
            }
            m.SetPaymentTerms(res)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pictureable, len(val))
            for i, v := range val {
                res[i] = v.(Pictureable)
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["purchaseInvoiceLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePurchaseInvoiceLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceLineable, len(val))
            for i, v := range val {
                res[i] = v.(PurchaseInvoiceLineable)
            }
            m.SetPurchaseInvoiceLines(res)
        }
        return nil
    }
    res["purchaseInvoices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePurchaseInvoiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceable, len(val))
            for i, v := range val {
                res[i] = v.(PurchaseInvoiceable)
            }
            m.SetPurchaseInvoices(res)
        }
        return nil
    }
    res["salesCreditMemoLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesCreditMemoLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesCreditMemoLineable, len(val))
            for i, v := range val {
                res[i] = v.(SalesCreditMemoLineable)
            }
            m.SetSalesCreditMemoLines(res)
        }
        return nil
    }
    res["salesCreditMemos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesCreditMemoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesCreditMemoable, len(val))
            for i, v := range val {
                res[i] = v.(SalesCreditMemoable)
            }
            m.SetSalesCreditMemos(res)
        }
        return nil
    }
    res["salesInvoiceLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesInvoiceLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesInvoiceLineable, len(val))
            for i, v := range val {
                res[i] = v.(SalesInvoiceLineable)
            }
            m.SetSalesInvoiceLines(res)
        }
        return nil
    }
    res["salesInvoices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesInvoiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesInvoiceable, len(val))
            for i, v := range val {
                res[i] = v.(SalesInvoiceable)
            }
            m.SetSalesInvoices(res)
        }
        return nil
    }
    res["salesOrderLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesOrderLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesOrderLineable, len(val))
            for i, v := range val {
                res[i] = v.(SalesOrderLineable)
            }
            m.SetSalesOrderLines(res)
        }
        return nil
    }
    res["salesOrders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesOrderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesOrderable, len(val))
            for i, v := range val {
                res[i] = v.(SalesOrderable)
            }
            m.SetSalesOrders(res)
        }
        return nil
    }
    res["salesQuoteLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesQuoteLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuoteLineable, len(val))
            for i, v := range val {
                res[i] = v.(SalesQuoteLineable)
            }
            m.SetSalesQuoteLines(res)
        }
        return nil
    }
    res["salesQuotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSalesQuoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesQuoteable, len(val))
            for i, v := range val {
                res[i] = v.(SalesQuoteable)
            }
            m.SetSalesQuotes(res)
        }
        return nil
    }
    res["shipmentMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateShipmentMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ShipmentMethodable, len(val))
            for i, v := range val {
                res[i] = v.(ShipmentMethodable)
            }
            m.SetShipmentMethods(res)
        }
        return nil
    }
    res["systemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemVersion(val)
        }
        return nil
    }
    res["taxAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaxAreaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaxAreaable, len(val))
            for i, v := range val {
                res[i] = v.(TaxAreaable)
            }
            m.SetTaxAreas(res)
        }
        return nil
    }
    res["taxGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaxGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaxGroupable, len(val))
            for i, v := range val {
                res[i] = v.(TaxGroupable)
            }
            m.SetTaxGroups(res)
        }
        return nil
    }
    res["unitsOfMeasure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnitOfMeasureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnitOfMeasureable, len(val))
            for i, v := range val {
                res[i] = v.(UnitOfMeasureable)
            }
            m.SetUnitsOfMeasure(res)
        }
        return nil
    }
    res["vendors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVendor_escapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Vendor_escapedable, len(val))
            for i, v := range val {
                res[i] = v.(Vendor_escapedable)
            }
            m.SetVendors(res)
        }
        return nil
    }
    return res
}
// GetGeneralLedgerEntries gets the generalLedgerEntries property value. The generalLedgerEntries property
func (m *Company) GetGeneralLedgerEntries()([]GeneralLedgerEntryable) {
    return m.generalLedgerEntries
}
// GetItemCategories gets the itemCategories property value. The itemCategories property
func (m *Company) GetItemCategories()([]ItemCategoryable) {
    return m.itemCategories
}
// GetItems gets the items property value. The items property
func (m *Company) GetItems()([]Itemable) {
    return m.items
}
// GetJournalLines gets the journalLines property value. The journalLines property
func (m *Company) GetJournalLines()([]JournalLineable) {
    return m.journalLines
}
// GetJournals gets the journals property value. The journals property
func (m *Company) GetJournals()([]Journalable) {
    return m.journals
}
// GetName gets the name property value. The name property
func (m *Company) GetName()(*string) {
    return m.name
}
// GetPaymentMethods gets the paymentMethods property value. The paymentMethods property
func (m *Company) GetPaymentMethods()([]PaymentMethodable) {
    return m.paymentMethods
}
// GetPaymentTerms gets the paymentTerms property value. The paymentTerms property
func (m *Company) GetPaymentTerms()([]PaymentTermable) {
    return m.paymentTerms
}
// GetPicture gets the picture property value. The picture property
func (m *Company) GetPicture()([]Pictureable) {
    return m.picture
}
// GetPurchaseInvoiceLines gets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *Company) GetPurchaseInvoiceLines()([]PurchaseInvoiceLineable) {
    return m.purchaseInvoiceLines
}
// GetPurchaseInvoices gets the purchaseInvoices property value. The purchaseInvoices property
func (m *Company) GetPurchaseInvoices()([]PurchaseInvoiceable) {
    return m.purchaseInvoices
}
// GetSalesCreditMemoLines gets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *Company) GetSalesCreditMemoLines()([]SalesCreditMemoLineable) {
    return m.salesCreditMemoLines
}
// GetSalesCreditMemos gets the salesCreditMemos property value. The salesCreditMemos property
func (m *Company) GetSalesCreditMemos()([]SalesCreditMemoable) {
    return m.salesCreditMemos
}
// GetSalesInvoiceLines gets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *Company) GetSalesInvoiceLines()([]SalesInvoiceLineable) {
    return m.salesInvoiceLines
}
// GetSalesInvoices gets the salesInvoices property value. The salesInvoices property
func (m *Company) GetSalesInvoices()([]SalesInvoiceable) {
    return m.salesInvoices
}
// GetSalesOrderLines gets the salesOrderLines property value. The salesOrderLines property
func (m *Company) GetSalesOrderLines()([]SalesOrderLineable) {
    return m.salesOrderLines
}
// GetSalesOrders gets the salesOrders property value. The salesOrders property
func (m *Company) GetSalesOrders()([]SalesOrderable) {
    return m.salesOrders
}
// GetSalesQuoteLines gets the salesQuoteLines property value. The salesQuoteLines property
func (m *Company) GetSalesQuoteLines()([]SalesQuoteLineable) {
    return m.salesQuoteLines
}
// GetSalesQuotes gets the salesQuotes property value. The salesQuotes property
func (m *Company) GetSalesQuotes()([]SalesQuoteable) {
    return m.salesQuotes
}
// GetShipmentMethods gets the shipmentMethods property value. The shipmentMethods property
func (m *Company) GetShipmentMethods()([]ShipmentMethodable) {
    return m.shipmentMethods
}
// GetSystemVersion gets the systemVersion property value. The systemVersion property
func (m *Company) GetSystemVersion()(*string) {
    return m.systemVersion
}
// GetTaxAreas gets the taxAreas property value. The taxAreas property
func (m *Company) GetTaxAreas()([]TaxAreaable) {
    return m.taxAreas
}
// GetTaxGroups gets the taxGroups property value. The taxGroups property
func (m *Company) GetTaxGroups()([]TaxGroupable) {
    return m.taxGroups
}
// GetUnitsOfMeasure gets the unitsOfMeasure property value. The unitsOfMeasure property
func (m *Company) GetUnitsOfMeasure()([]UnitOfMeasureable) {
    return m.unitsOfMeasure
}
// GetVendors gets the vendors property value. The vendors property
func (m *Company) GetVendors()([]Vendor_escapedable) {
    return m.vendors
}
// Serialize serializes information the current object
func (m *Company) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsPayable() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgedAccountsPayable()))
        for i, v := range m.GetAgedAccountsPayable() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agedAccountsPayable", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgedAccountsReceivable() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgedAccountsReceivable()))
        for i, v := range m.GetAgedAccountsReceivable() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agedAccountsReceivable", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessProfileId", m.GetBusinessProfileId())
        if err != nil {
            return err
        }
    }
    if m.GetCompanyInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyInformation()))
        for i, v := range m.GetCompanyInformation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("companyInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCountriesRegions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCountriesRegions()))
        for i, v := range m.GetCountriesRegions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("countriesRegions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCurrencies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCurrencies()))
        for i, v := range m.GetCurrencies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("currencies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPaymentJournals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomerPaymentJournals()))
        for i, v := range m.GetCustomerPaymentJournals() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customerPaymentJournals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomerPayments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomerPayments()))
        for i, v := range m.GetCustomerPayments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customerPayments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDimensions()))
        for i, v := range m.GetDimensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dimensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDimensionValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDimensionValues()))
        for i, v := range m.GetDimensionValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dimensionValues", cast)
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
    if m.GetEmployees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmployees()))
        for i, v := range m.GetEmployees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("employees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGeneralLedgerEntries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGeneralLedgerEntries()))
        for i, v := range m.GetGeneralLedgerEntries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("generalLedgerEntries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItemCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemCategories()))
        for i, v := range m.GetItemCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("itemCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournalLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJournalLines()))
        for i, v := range m.GetJournalLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("journalLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJournals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJournals()))
        for i, v := range m.GetJournals() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("journals", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetPaymentMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPaymentMethods()))
        for i, v := range m.GetPaymentMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("paymentMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPaymentTerms() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPaymentTerms()))
        for i, v := range m.GetPaymentTerms() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("paymentTerms", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPicture() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoiceLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPurchaseInvoiceLines()))
        for i, v := range m.GetPurchaseInvoiceLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("purchaseInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPurchaseInvoices()))
        for i, v := range m.GetPurchaseInvoices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("purchaseInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemoLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesCreditMemoLines()))
        for i, v := range m.GetSalesCreditMemoLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesCreditMemoLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesCreditMemos() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesCreditMemos()))
        for i, v := range m.GetSalesCreditMemos() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesCreditMemos", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoiceLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesInvoiceLines()))
        for i, v := range m.GetSalesInvoiceLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesInvoices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesInvoices()))
        for i, v := range m.GetSalesInvoices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesInvoices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrderLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesOrderLines()))
        for i, v := range m.GetSalesOrderLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesOrderLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesOrders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesOrders()))
        for i, v := range m.GetSalesOrders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesOrders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuoteLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesQuoteLines()))
        for i, v := range m.GetSalesQuoteLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesQuoteLines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSalesQuotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSalesQuotes()))
        for i, v := range m.GetSalesQuotes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("salesQuotes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetShipmentMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShipmentMethods()))
        for i, v := range m.GetShipmentMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("shipmentMethods", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("systemVersion", m.GetSystemVersion())
        if err != nil {
            return err
        }
    }
    if m.GetTaxAreas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaxAreas()))
        for i, v := range m.GetTaxAreas() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taxAreas", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaxGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaxGroups()))
        for i, v := range m.GetTaxGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taxGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUnitsOfMeasure() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnitsOfMeasure()))
        for i, v := range m.GetUnitsOfMeasure() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("unitsOfMeasure", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVendors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVendors()))
        for i, v := range m.GetVendors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("vendors", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccounts sets the accounts property value. The accounts property
func (m *Company) SetAccounts(value []Accountable)() {
    m.accounts = value
}
// SetAgedAccountsPayable sets the agedAccountsPayable property value. The agedAccountsPayable property
func (m *Company) SetAgedAccountsPayable(value []AgedAccountsPayableable)() {
    m.agedAccountsPayable = value
}
// SetAgedAccountsReceivable sets the agedAccountsReceivable property value. The agedAccountsReceivable property
func (m *Company) SetAgedAccountsReceivable(value []AgedAccountsReceivableable)() {
    m.agedAccountsReceivable = value
}
// SetBusinessProfileId sets the businessProfileId property value. The businessProfileId property
func (m *Company) SetBusinessProfileId(value *string)() {
    m.businessProfileId = value
}
// SetCompanyInformation sets the companyInformation property value. The companyInformation property
func (m *Company) SetCompanyInformation(value []CompanyInformationable)() {
    m.companyInformation = value
}
// SetCountriesRegions sets the countriesRegions property value. The countriesRegions property
func (m *Company) SetCountriesRegions(value []CountryRegionable)() {
    m.countriesRegions = value
}
// SetCurrencies sets the currencies property value. The currencies property
func (m *Company) SetCurrencies(value []Currencyable)() {
    m.currencies = value
}
// SetCustomerPaymentJournals sets the customerPaymentJournals property value. The customerPaymentJournals property
func (m *Company) SetCustomerPaymentJournals(value []CustomerPaymentJournalable)() {
    m.customerPaymentJournals = value
}
// SetCustomerPayments sets the customerPayments property value. The customerPayments property
func (m *Company) SetCustomerPayments(value []CustomerPaymentable)() {
    m.customerPayments = value
}
// SetCustomers sets the customers property value. The customers property
func (m *Company) SetCustomers(value []Customerable)() {
    m.customers = value
}
// SetDimensions sets the dimensions property value. The dimensions property
func (m *Company) SetDimensions(value []Dimensionable)() {
    m.dimensions = value
}
// SetDimensionValues sets the dimensionValues property value. The dimensionValues property
func (m *Company) SetDimensionValues(value []DimensionValueable)() {
    m.dimensionValues = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Company) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmployees sets the employees property value. The employees property
func (m *Company) SetEmployees(value []Employeeable)() {
    m.employees = value
}
// SetGeneralLedgerEntries sets the generalLedgerEntries property value. The generalLedgerEntries property
func (m *Company) SetGeneralLedgerEntries(value []GeneralLedgerEntryable)() {
    m.generalLedgerEntries = value
}
// SetItemCategories sets the itemCategories property value. The itemCategories property
func (m *Company) SetItemCategories(value []ItemCategoryable)() {
    m.itemCategories = value
}
// SetItems sets the items property value. The items property
func (m *Company) SetItems(value []Itemable)() {
    m.items = value
}
// SetJournalLines sets the journalLines property value. The journalLines property
func (m *Company) SetJournalLines(value []JournalLineable)() {
    m.journalLines = value
}
// SetJournals sets the journals property value. The journals property
func (m *Company) SetJournals(value []Journalable)() {
    m.journals = value
}
// SetName sets the name property value. The name property
func (m *Company) SetName(value *string)() {
    m.name = value
}
// SetPaymentMethods sets the paymentMethods property value. The paymentMethods property
func (m *Company) SetPaymentMethods(value []PaymentMethodable)() {
    m.paymentMethods = value
}
// SetPaymentTerms sets the paymentTerms property value. The paymentTerms property
func (m *Company) SetPaymentTerms(value []PaymentTermable)() {
    m.paymentTerms = value
}
// SetPicture sets the picture property value. The picture property
func (m *Company) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetPurchaseInvoiceLines sets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *Company) SetPurchaseInvoiceLines(value []PurchaseInvoiceLineable)() {
    m.purchaseInvoiceLines = value
}
// SetPurchaseInvoices sets the purchaseInvoices property value. The purchaseInvoices property
func (m *Company) SetPurchaseInvoices(value []PurchaseInvoiceable)() {
    m.purchaseInvoices = value
}
// SetSalesCreditMemoLines sets the salesCreditMemoLines property value. The salesCreditMemoLines property
func (m *Company) SetSalesCreditMemoLines(value []SalesCreditMemoLineable)() {
    m.salesCreditMemoLines = value
}
// SetSalesCreditMemos sets the salesCreditMemos property value. The salesCreditMemos property
func (m *Company) SetSalesCreditMemos(value []SalesCreditMemoable)() {
    m.salesCreditMemos = value
}
// SetSalesInvoiceLines sets the salesInvoiceLines property value. The salesInvoiceLines property
func (m *Company) SetSalesInvoiceLines(value []SalesInvoiceLineable)() {
    m.salesInvoiceLines = value
}
// SetSalesInvoices sets the salesInvoices property value. The salesInvoices property
func (m *Company) SetSalesInvoices(value []SalesInvoiceable)() {
    m.salesInvoices = value
}
// SetSalesOrderLines sets the salesOrderLines property value. The salesOrderLines property
func (m *Company) SetSalesOrderLines(value []SalesOrderLineable)() {
    m.salesOrderLines = value
}
// SetSalesOrders sets the salesOrders property value. The salesOrders property
func (m *Company) SetSalesOrders(value []SalesOrderable)() {
    m.salesOrders = value
}
// SetSalesQuoteLines sets the salesQuoteLines property value. The salesQuoteLines property
func (m *Company) SetSalesQuoteLines(value []SalesQuoteLineable)() {
    m.salesQuoteLines = value
}
// SetSalesQuotes sets the salesQuotes property value. The salesQuotes property
func (m *Company) SetSalesQuotes(value []SalesQuoteable)() {
    m.salesQuotes = value
}
// SetShipmentMethods sets the shipmentMethods property value. The shipmentMethods property
func (m *Company) SetShipmentMethods(value []ShipmentMethodable)() {
    m.shipmentMethods = value
}
// SetSystemVersion sets the systemVersion property value. The systemVersion property
func (m *Company) SetSystemVersion(value *string)() {
    m.systemVersion = value
}
// SetTaxAreas sets the taxAreas property value. The taxAreas property
func (m *Company) SetTaxAreas(value []TaxAreaable)() {
    m.taxAreas = value
}
// SetTaxGroups sets the taxGroups property value. The taxGroups property
func (m *Company) SetTaxGroups(value []TaxGroupable)() {
    m.taxGroups = value
}
// SetUnitsOfMeasure sets the unitsOfMeasure property value. The unitsOfMeasure property
func (m *Company) SetUnitsOfMeasure(value []UnitOfMeasureable)() {
    m.unitsOfMeasure = value
}
// SetVendors sets the vendors property value. The vendors property
func (m *Company) SetVendors(value []Vendor_escapedable)() {
    m.vendors = value
}
