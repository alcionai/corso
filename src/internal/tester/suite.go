package tester

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func NewUnitSuite(t *testing.T, envSets [][]string) *UnitSuite {
	MustGetEnvSets(t, envSets...)
	return new(UnitSuite)
}

type UnitSuite struct {
	suite.Suite
}

func NewIntegrationSuite(
	t *testing.T,
	envSets [][]string,
	includeGroups ...string,
) *IntegrationSuite {
	RunOnAny(
		t,
		append(
			[]string{CorsoCITests, CorsoIntegrationTests},
			includeGroups...,
		)...,
	)

	MustGetEnvSets(t, envSets...)

	return new(IntegrationSuite)
}

type IntegrationSuite struct {
	suite.Suite
}
