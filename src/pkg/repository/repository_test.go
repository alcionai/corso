package repository

import (
	"context"
	"errors"
	"testing"
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

func TestInitializeWith(t *testing.T) {
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
			err := initializeWith(context.Background(), test.ic)
			if (err != nil) != test.expectErr {
				t.Fatalf("unexpected error - wanted err [%v] - got [%v]", test.expectErr, err)
			}
		})
	}
}

func TestConnectWith(t *testing.T) {
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
			err := connectWith(context.Background(), test.ic)
			if (err != nil) != test.expectErr {
				t.Fatalf("unexpected error - wanted err [%v] - got [%v]", test.expectErr, err)
			}
		})
	}
}
