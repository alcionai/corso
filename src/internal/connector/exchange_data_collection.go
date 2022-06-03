package connector

// A DataStream provides an iterator to consume a
// collection of graph data of the same type (e.g. mail)
type DataStream interface{}

// ExchangeDataCollection represents exchange mailbox
// data for a single user. It implements the DataStream
// interface which allows reading data in the collection
type ExchangeDataCollection struct {
	user string
}
