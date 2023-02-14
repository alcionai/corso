package tester

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite interface {
	suite.TestingSuite
	Run(name string, subtest func()) bool
}

func NewUnitSuite(t *testing.T) *unitSuite {
	return new(unitSuite)
}

type unitSuite struct {
	suite.Suite
}

func NewIntegrationSuite(
	t *testing.T,
	envSets [][]string,
	includeGroups ...string,
) *integrationSuite {
	RunOnAny(
		t,
		append(
			[]string{CorsoCITests},
			includeGroups...,
		)...,
	)

	MustGetEnvSets(t, envSets...)

	return new(integrationSuite)
}

type integrationSuite struct {
	suite.Suite
}
