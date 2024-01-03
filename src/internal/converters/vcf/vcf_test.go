package vcf

import (
	"strings"
	"testing"
	"time"

	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/converters/vcf/testdata"
	"github.com/alcionai/corso/src/internal/tester"
)

type VCFUnitSuite struct {
	tester.Suite
}

func TestVCFUnitSuite(t *testing.T) {
	suite.Run(t, &VCFUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *VCFUnitSuite) TestConvert_contactable_to_vcf() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(testdata.ContactsInput)

	bytes, err := FromJSON(ctx, body)
	require.NoError(t, err, "convert")

	out := strings.ReplaceAll(string(bytes), "\r", "") // output contains \r
	assert.Equal(t, strings.TrimSpace(testdata.ContactsOutput), strings.TrimSpace(string(out)))
}

func (suite *VCFUnitSuite) TestConvert_contactable_cases() {
	t := suite.T()

	tests := []struct {
		name           string
		transformation func(contact models.Contactable)
		check          string
	}{
		{
			name: "name",
			transformation: func(contact models.Contactable) {
				contact.SetGivenName(ptr.To("given"))
				contact.SetSurname(ptr.To("sur"))
			},
			check: "N:sur;given;;;",
		},
		{
			name: "all name related",
			transformation: func(contact models.Contactable) {
				contact.SetGivenName(ptr.To("given"))
				contact.SetSurname(ptr.To("sur"))
				contact.SetMiddleName(ptr.To("middle"))
				contact.SetTitle(ptr.To("title"))
				contact.SetGeneration(ptr.To("gen"))
			},
			check: "N:sur;given;middle;title;gen",
		},
		{
			name: "org",
			transformation: func(contact models.Contactable) {
				contact.SetCompanyName(ptr.To("org"))
			},
			check: "ORG:org",
		},
		{
			name: "org,dept,prof",
			transformation: func(contact models.Contactable) {
				contact.SetCompanyName(ptr.To("org"))
				contact.SetDepartment(ptr.To("dept"))
				contact.SetProfession(ptr.To("prof"))
			},
			check: "ORG:org;dept;prof",
		},
		{
			name: "dept,prof without org name",
			transformation: func(contact models.Contactable) {
				contact.SetDepartment(ptr.To("dept"))
				contact.SetProfession(ptr.To("prof"))
			},
			check: "ORG:dept;prof",
		},
		{
			name: "org,prof without dept",
			transformation: func(contact models.Contactable) {
				contact.SetCompanyName(ptr.To("org"))
				contact.SetProfession(ptr.To("prof"))
			},
			check: "ORG:org;prof",
		},
		{
			name: "birthday",
			transformation: func(contact models.Contactable) {
				date := time.Time(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
				contact.SetBirthday(ptr.To(date))
			},
			check: "BDAY:2000-01-01",
		},
		{
			name: "address",
			transformation: func(contact models.Contactable) {
				add := models.NewPhysicalAddress()
				add.SetStreet(ptr.To("street"))
				add.SetCity(ptr.To("city"))
				add.SetState(ptr.To("state"))
				add.SetCountryOrRegion(ptr.To("country"))
				add.SetPostalCode(ptr.To("zip"))
				contact.SetHomeAddress(add)
			},
			check: "ADR;TYPE=home:;;street;city;state;zip;country",
		},
		{
			name: "mobile",
			transformation: func(contact models.Contactable) {
				contact.SetMobilePhone(ptr.To("mobile"))
			},
			check: "TEL;TYPE=cell:mobile",
		},
		{
			name: "home",
			transformation: func(contact models.Contactable) {
				contact.SetHomePhones([]string{"home"})
			},
			check: "TEL;TYPE=home:home", // ideally check both
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, flush := tester.NewContext(t)
			defer flush()

			contact := models.NewContact()
			tt.transformation(contact)

			writer := kjson.NewJsonSerializationWriter()
			defer writer.Close()

			err := writer.WriteObjectValue("", contact)
			require.NoError(t, err, "serializing contact")

			nbody, err := writer.GetSerializedContent()
			require.NoError(t, err, "getting serialized content")

			bytes, err := FromJSON(ctx, nbody)
			require.NoError(t, err, "convert")

			assert.Contains(t, string(bytes), tt.check)
		})
	}
}
