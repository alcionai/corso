package dataTransport

import (
	"bytes"
	"io"
)

type ExchangeData struct {
	GraphData
	uuid    string
	message []byte
}

func NewExchangeData(name string) ExchangeData {
	exchange := ExchangeData{
		uuid:    name,
		message: make([]byte, 0),
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
		uuid:    name,
		message: bArray,
	}
	return exchange
}

func (exchange *ExchangeData) UUID() string {
	return exchange.uuid
}

func (exchange *ExchangeData) ToReader() *io.Reader {
	return exchange
}

func (exchange *ExchangeData) Read(bytes []byte) (int, error) {
	return 0, nil
}
