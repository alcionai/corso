package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
)

type ModelUnitSuite struct {
	tester.Suite
}

func TestModelUnitSuite(t *testing.T) {
	suite.Run(t, &ModelUnitSuite{Suite: tester.NewUnitSuite(t)})
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
		{model.RepositorySchema, assert.True},
		{model.RepositorySchema + 1, assert.False},
		{model.Schema(-1), assert.False},
		{model.Schema(100), assert.False},
	}
	for _, test := range table {
		suite.Run(test.mt.String(), func() {
			test.expect(suite.T(), test.mt.Valid())
		})
	}
}

func (suite *ModelUnitSuite) TestGetID() {
	bm := model.BaseModel{
		ID: model.StableID(uuid.NewString()),
	}

	assert.Equal(suite.T(), string(bm.ID), bm.GetID())
}
