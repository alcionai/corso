package connector_test

import (
	"os"
	"testing"

	graph "github.com/alcionai/corso/internal/connector"
	"github.com/stretchr/testify/assert"
)

func TestBadConnection(t *testing.T) {
	gc := graph.NewGraphConnector("Test", "without", "data")
	assert.True(t, gc.HasConnectorErrors())
	assert.Equal(t, len(gc.GetUsers()), 0)

}

func TestConnectWithEnvVariables(t *testing.T) {
	tenant := os.Getenv("TENANT_ID")
	client := os.Getenv("CLIENT_ID")
	secret := os.Getenv("CLIENT_SECRET")

	if tenant == "" || client == "" || secret == "" {
		t.Logf("Connection Test Skipped\n")
	} else {
		gc := graph.NewGraphConnector(tenant, client, secret)
		assert.False(t, gc.HasConnectorErrors())
		t.Logf("Users: %v\n", gc.GetUsers())
		assert.True(t, len(gc.GetUsers()) > 0)
	}
}
