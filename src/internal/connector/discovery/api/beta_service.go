package api

import (
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk"
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/pkg/errors"
)

// Service wraps BetaClient's functionality.
// Abstraction created to comply loosely with graph.Servicer
// methods for ease of switching between v1.0 and beta connnectors
type Service struct {
	client *betasdk.BetaClient
}

func (s Service) Client() *betasdk.BetaClient {
	return s.client
}

func NewBetaService(adpt *msgraphsdk.GraphRequestAdapter) *Service {
	return &Service{
		client: betasdk.NewBetaClient(adpt),
	}
}

// Seraialize writes an M365 parsable object into a byte array using the built-in
// application/json writer within the adapter.
func (s Service) Serialize(object absser.Parsable) ([]byte, error) {
	writer, err := s.client.Adapter().
		GetSerializationWriterFactory().
		GetSerializationWriter("application/json")
	if err != nil || writer == nil {
		return nil, errors.Wrap(err, "creating json serialization writer")
	}

	err = writer.WriteObjectValue("", object)
	if err != nil {
		return nil, errors.Wrap(err, "writeObjecValue serialization")
	}

	return writer.GetSerializedContent()
}
