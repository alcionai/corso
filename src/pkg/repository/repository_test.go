package repository_test

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
)

func TestInitialize(t *testing.T) {
	table := []struct {
		storage     storage.Storage
		account     repository.Account
		expectedErr string
	}{
		{
			storage.NewStorage(storage.ProviderUnknown),
			repository.Account{},
			"provider details are required",
		},
	}
	for _, test := range table {
		t.Run(test.expectedErr, func(t *testing.T) {
			_, err := repository.Initialize(context.Background(), test.account, test.storage)
			if err == nil || !strings.Contains(err.Error(), test.expectedErr) {
				t.Fatalf("expected error with [%s], got [%v]", test.expectedErr, err)
			}
		})
	}
}

// repository.Connect involves end-to-end communication with kopia, therefore this only
// tests expected error cases from
func TestConnect(t *testing.T) {
	table := []struct {
		storage     storage.Storage
		account     repository.Account
		expectedErr string
	}{
		{
			storage.NewStorage(storage.ProviderUnknown),
			repository.Account{},
			"provider details are required",
		},
	}
	for _, test := range table {
		t.Run(test.expectedErr, func(t *testing.T) {
			_, err := repository.Connect(context.Background(), test.account, test.storage)
			if err == nil || !strings.Contains(err.Error(), test.expectedErr) {
				t.Fatalf("expected error with [%s], got [%v]", test.expectedErr, err)
			}
		})
	}
}
