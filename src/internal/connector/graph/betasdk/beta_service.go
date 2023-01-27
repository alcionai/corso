package betasdk

import (
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

// Service wraps BetaClient's functionality.
// Abstraction created to comply loosely with graph.Servicer
// methods for ease of switching between v1.0 and beta connnectors
type Service struct {
	client *BetaClient
}

func (s Service) Adapter() *msgraphsdk.GraphRequestAdapter {
	return s.client.requestAdapter
}

func (s Service) Client() *BetaClient {
	return s.client
}
