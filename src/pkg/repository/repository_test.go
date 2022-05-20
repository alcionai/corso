package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alcionai/corso/pkg/repository"
)

type testInitConn struct {
	err error
}

func (tic testInitConn) Initialize(ctx context.Context) error {
	return tic.err
}

func (tic testInitConn) Connect(ctx context.Context) error {
	return tic.err
}

func TestInitialize(t *testing.T) {
	table := []struct {
		name      string
		ic        testInitConn
		expectErr bool
	}{
		{"no errors", testInitConn{}, false},
		{"errors", testInitConn{errors.New("an error")}, true},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			_, err := repository.Initialize(
				context.Background(),
				repository.ProviderUnknown,
				repository.Account{},
				test.ic)
			if (err != nil) != test.expectErr {
				t.Fatalf("unexpected error - wanted err [%v] - got [%v]", test.expectErr, err)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	table := []struct {
		name      string
		ic        testInitConn
		expectErr bool
	}{
		{"no errors", testInitConn{}, false},
		{"errors", testInitConn{errors.New("an error")}, true},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			_, err := repository.Connect(
				context.Background(),
				repository.ProviderUnknown,
				repository.Account{},
				test.ic)
			if (err != nil) != test.expectErr {
				t.Fatalf("unexpected error - wanted err [%v] - got [%v]", test.expectErr, err)
			}
		})
	}
}
