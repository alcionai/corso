package mock

import "fmt"

const (
	// Order of fields to fill in:
	// 1. displayName
	// 2. fileAsName
	// 3. phone
	// 4. givenName
	// 5. middleName
	// 6. surname
	//nolint:lll
	contactTmpl = `{
	"id":"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEOAADSEBNbUIB9RL6ePDeF3FIYAABS7DZnAAA=",
	"@odata.context":"https://graph.microsoft.com/v1.0/$metadata#users('foobar%%408qzvrj.onmicrosoft.com')/contacts/$entity",
	"@odata.etag":"W/\"EQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAABSx4Tr\"",
	"categories":[],
	"changeKey":"EQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAABSx4Tr",
	"createdDateTime":"2019-08-04T06:55:33Z",
	"lastModifiedDateTime":"2019-08-04T06:55:33Z",
	"businessAddress":{},
	"businessPhones":[],
	"children":[],
	"displayName":"%s",
	"emailAddresses":[],
	"fileAs":"%s",
	"mobilePhone":"%s",
	"givenName":"%s",
	"homeAddress":{},
	"homePhones":[],
	"imAddresses":[],
	"otherAddress":{},
	"parentFolderId":"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9FIYAAAAAAEOAAA=",
	"personalNotes":"",
	"middleName":"%s",
	"surname":"%s"
}`

	defaultContactDisplayName = "Santiago Quail"
	defaultContactFileAsName  = "Quail, Santiago"
	defaultContactGivenName   = "Santiago"
	defaultContactSurname     = "Quail"
)

// ContactBytes returns bytes for Contactable item.
// When hydrated: contact.GetGivenName() shows differences
func ContactBytes(middleName string) []byte {
	phone := generatePhoneNumber()

	return ContactBytesWith(
		defaultContactDisplayName,
		defaultContactFileAsName,
		defaultContactGivenName,
		middleName,
		defaultContactSurname,
		phone,
	)
}

func ContactBytesWith(
	displayName, fileAsName,
	givenName, middleName, surname,
	phone string,
) []byte {
	return []byte(fmt.Sprintf(
		contactTmpl,
		displayName,
		fileAsName,
		phone,
		givenName,
		middleName,
		surname,
	))
}
