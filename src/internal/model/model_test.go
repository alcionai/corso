package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/model"
)

type ModelUnitSuite struct {
	suite.Suite
}

func TestModelUnitSuite(t *testing.T) {
	suite.Run(t, new(ModelUnitSuite))
}

func (suite *ModelUnitSuite) TestValid() {
	table := []struct {
		mt     model.Schema
		expect assert.BoolAssertionFunc
	}{
		{model.UnknownSchema, assert.False},
		{model.BackupOpSchema, assert.True},
		{model.RestoreOpSchema, assert.True},
		{model.BackupSchema, assert.True},
		{model.BackupDetailsSchema, assert.True},
		{model.Schema(-1), assert.False},
		{model.Schema(100), assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.mt.String(), func(t *testing.T) {
			test.expect(t, test.mt.Valid())
		})
	}
}
