package streamstore

type Collectable struct {
	mr       Marshaller
	Unmr     Unmarshaller
	itemName string
	purpose  string
	Type     string
}

const (
	FaultErrorsType     = "fault_error"
	faultErrorsItemName = "fault_error"
	faultErrorsPurpose  = "fault_error"

	DetailsType     = "details"
	detailsItemName = "details"
	detailsPurpose  = "details"
)

// FaultErrorsCollector generates a collection of fault.Errors
// containing the marshalled bytes from the provieded marshaller.
func FaultErrorsCollector(mr Marshaller) Collectable {
	return Collectable{
		mr:       mr,
		itemName: faultErrorsItemName,
		purpose:  faultErrorsPurpose,
		Type:     FaultErrorsType,
	}
}

// FaultErrorsCollector generates a collection of details.DetailsModel
// entries containing the marshalled bytes from the provieded marshaller.
func DetailsCollector(mr Marshaller) Collectable {
	return Collectable{
		mr:       mr,
		itemName: detailsItemName,
		purpose:  detailsPurpose,
		Type:     DetailsType,
	}
}

// FaultErrorsCollector reads a collection of fault.Errors
// entries using the provided unmarshaller.
func FaultErrorsReader(unmr Unmarshaller) Collectable {
	return Collectable{
		Unmr:     unmr,
		itemName: faultErrorsItemName,
		purpose:  faultErrorsPurpose,
		Type:     FaultErrorsType,
	}
}

// FaultErrorsCollector reads a collection of details.DetailsModel
// entries using the provided unmarshaller.
func DetailsReader(unmr Unmarshaller) Collectable {
	return Collectable{
		Unmr:     unmr,
		itemName: detailsItemName,
		purpose:  detailsPurpose,
		Type:     DetailsType,
	}
}
