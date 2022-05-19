package repository_test

import (
	"context"
	"testing"

	"github.com/kopia/kopia/repo/blob"

	"github.com/alcionai/corso/pkg/repository"
)

type testConfig struct{}

func (tc testConfig) KopiaStorage(ctx context.Context, create bool) (blob.Storage, error) {
	return nil, nil
}

func TestInitialize(t *testing.T) {
	_, err := repository.Initialize(
		context.Background(),
		repository.ProviderUnknown,
		repository.Account{},
		testConfig{})
	if err != nil {
		t.Fatalf("didn't expect initialize to error, got [%v]", err)
	}
}

func TestConnect(t *testing.T) {
	_, err := repository.Connect(
		context.Background(),
		repository.ProviderUnknown,
		repository.Account{},
		testConfig{})
	if err != nil {
		t.Fatalf("didn't expect connect to error, got [%v]", err)
	}
}
