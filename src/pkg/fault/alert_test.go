package fault_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type AlertUnitSuite struct {
	tester.Suite
}

func TestAlertUnitSuite(t *testing.T) {
	suite.Run(t, &AlertUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *AlertUnitSuite) TestAlert_String() {
	var (
		t = suite.T()
		a fault.Alert
	)

	assert.Contains(t, a.String(), "Alert: <missing>")

	a = fault.Alert{
		Item:    fault.Item{},
		Message: "",
	}
	assert.Contains(t, a.String(), "Alert: <missing>")

	a = fault.Alert{
		Item: fault.Item{
			ID: "item_id",
		},
		Message: "msg",
	}
	assert.NotContains(t, a.String(), "item_id")
	assert.Contains(t, a.String(), "Alert: msg")
}

func (suite *AlertUnitSuite) TestNewAlert() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	a := fault.NewAlert("message-to-show", "ns", "item_id", "item_name", addtl)

	expect := fault.Alert{
		Item: fault.Item{
			Namespace:  "ns",
			ID:         "item_id",
			Name:       "item_name",
			Additional: addtl,
		},
		Message: "message-to-show",
	}

	assert.Equal(t, expect, *a)
}

func (suite *AlertUnitSuite) TestAlert_HeadersValues() {
	addtl := map[string]any{
		fault.AddtlContainerID:   "cid",
		fault.AddtlContainerName: "cname",
	}

	table := []struct {
		name   string
		alert  *fault.Alert
		expect []string
	}{
		{
			name:   "new alert",
			alert:  fault.NewAlert("message-to-show", "ns", "id", "name", addtl),
			expect: []string{"Alert", "message-to-show", "cname", "name", "id"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, []string{"Action", "Message", "Container", "Name", "ID"}, test.alert.Headers(false))
			assert.Equal(t, test.expect, test.alert.Values(false))

			assert.Equal(t, []string{"Action", "Message", "Container", "Name", "ID"}, test.alert.Headers(true))
			assert.Equal(t, test.expect, test.alert.Values(true))
		})
	}
}
