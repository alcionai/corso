package tester

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func NewUnitSuite(t *testing.T) *UnitSuite {
	return new(UnitSuite)
}

type UnitSuite struct {
	suite.Suite
}

func NewIntegrationSuite(t *testing.T, includeGroups ...string) *IntegrationSuite {
	RunOnAny(
		t,
		append(
			[]string{CorsoCITests, CorsoIntegrationTests},
			includeGroups...,
		)...,
	)

	return new(IntegrationSuite)
}

type IntegrationSuite struct {
	suite.Suite
}
