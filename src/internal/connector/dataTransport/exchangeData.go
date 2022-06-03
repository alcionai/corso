package dataTransport

type ExchangeData struct {
	// Provides file data to kopia.
	uuid string
	Data []byte
}

func NewExchangeData(name string) ExchangeData {
	exchange := ExchangeData{
		uuid: name,
		Data: make([]byte, 0),
	}
	return exchange
}

func NewExchangeDataFilled(name string, bytes []byte) ExchangeData {
	var bArray []byte
	if len(bytes) > 0 {
		bArray = bytes
	} else {
		bArray = make([]byte, 0)
	}
	exchange := ExchangeData{
		uuid: name,
		Data: bArray,
	}
	return exchange
}

func (exchange *ExchangeData) UUID() string {
	return exchange.uuid
}
