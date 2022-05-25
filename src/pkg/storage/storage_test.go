package storage

import (
	"testing"
)

type testConfig struct {
	expect string
}

func (c testConfig) Config() config {
	return config{"expect": c.expect}
}

func TestNewStorage(t *testing.T) {
	table := []struct {
		p storageProvider
		c testConfig
	}{
		{ProviderUnknown, testConfig{"unknown"}},
		{ProviderS3, testConfig{"s3"}},
	}
	for _, test := range table {
		s := NewStorage(test.p, test.c)
		if s.Provider != test.p {
			t.Errorf("expected storage provider [%s], got [%s]", test.p, s.Provider)
		}
		if s.Config["expect"] != test.c.expect {
			t.Errorf("expected storage config [%s], got [%s]", test.c.expect, s.Config["expect"])
		}
	}
}

type fooConfig struct {
	foo string
}

func (c fooConfig) Config() config {
	return config{"foo": c.foo}
}

func TestUnionConfigs(t *testing.T) {
	te := testConfig{"test"}
	f := fooConfig{"foo"}
	cs := unionConfigs(te, f)
	if cs["expect"] != te.expect {
		t.Errorf("expected unioned config to have value [%s] at key [expect], got [%s]", te.expect, cs["expect"])
	}
	if cs["foo"] != f.foo {
		t.Errorf("expected unioned config to have value [%s] at key [foo], got [%s]", f.foo, cs["foo"])
	}
}
