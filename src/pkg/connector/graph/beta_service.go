package graphapi

import (
	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/internal/connector/graph/betasdk"
)

// Service wraps BetaClient's functionality.
// Abstraction created to comply loosely with graph.Servicer
// methods for ease of switching between v1.0 and beta connnectors
type BetaService struct {
	client *betasdk.BetaClient
}

func (s BetaService) Client() *betasdk.BetaClient {
	return s.client
}

func NewBetaService(adpt *msgraphsdkgo.GraphRequestAdapter) *BetaService {
	return &BetaService{
		client: betasdk.NewBetaClient(adpt),
	}
}

// Seraialize writes an M365 parsable object into a byte array using the built-in
// application/json writer within the adapter.
func (s BetaService) Serialize(object serialization.Parsable) ([]byte, error) {
	writer, err := s.client.Adapter().
		GetSerializationWriterFactory().
		GetSerializationWriter("application/json")
	if err != nil || writer == nil {
		return nil, clues.Wrap(err, "creating json serialization writer")
	}

	err = writer.WriteObjectValue("", object)
	if err != nil {
		return nil, clues.Wrap(err, "writeObjectValue serialization")
	}

	return writer.GetSerializedContent()
}
