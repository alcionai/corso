package ptr_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
)

type PointerSuite struct {
	tester.Suite
}

func TestPointerSuite(t *testing.T) {
	s := &PointerSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

// TestVal checks to ptr derefencing for the
// following types:
// - *string
// - *bool
// - *time.Time
func (suite *PointerSuite) TestVal() {
	var (
		t          = suite.T()
		created    *time.Time
		testString *string
		testBool   *bool
		testInt    *int
		testInt32  *int32
		testInt64  *int64
	)

	// String Checks
	subject := ptr.Val(testString)
	assert.Empty(t, subject)

	hello := "Hello World"
	testString = &hello
	subject = ptr.Val(testString)

	t.Logf("Received: %s", subject)
	assert.NotEmpty(t, subject)

	// Time Checks

	myTime := ptr.Val(created)
	assert.Empty(t, myTime)
	assert.NotNil(t, myTime)

	now := time.Now()
	created = &now
	myTime = ptr.Val(created)
	assert.NotEmpty(t, myTime)

	// Bool Checks
	truth := true
	myBool := ptr.Val(testBool)
	assert.NotNil(t, myBool)
	assert.False(t, myBool)

	testBool = &truth
	myBool = ptr.Val(testBool)
	assert.NotNil(t, myBool)
	assert.True(t, myBool)

	// Int checks
	myInt := ptr.Val(testInt)
	myInt32 := ptr.Val(testInt32)
	myInt64 := ptr.Val(testInt64)

	assert.NotNil(t, myInt)
	assert.NotNil(t, myInt32)
	assert.NotNil(t, myInt64)
	assert.Empty(t, myInt)
	assert.Empty(t, myInt32)
	assert.Empty(t, myInt64)

	num := 4071
	num32 := int32(num * 32)
	num64 := int64(num * 2048)
	testInt = &num
	testInt32 = &num32
	testInt64 = &num64

	myInt = ptr.Val(testInt)
	myInt32 = ptr.Val(testInt32)
	myInt64 = ptr.Val(testInt64)

	assert.NotNil(t, myInt)
	assert.NotNil(t, myInt32)
	assert.NotNil(t, myInt64)
	assert.NotEmpty(t, myInt)
	assert.NotEmpty(t, myInt32)
	assert.NotEmpty(t, myInt64)
}
