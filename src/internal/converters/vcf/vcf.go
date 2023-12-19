package vcf

import (
	"bytes"
	"context"

	"github.com/alcionai/clues"
	"github.com/emersion/go-vcard"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// This package helps convert the json response backed up from graph
// API to a vCard file.
// Ref: https://learn.microsoft.com/en-us/graph/api/resources/contact?view=graph-rest-1.0
// Ref: https://datatracker.ietf.org/doc/html/rfc6350

// TODO: items that are only available via beta api and not mapped
// weddingAnniversary, gender, websites

func addAddress(iaddr models.PhysicalAddressable, addrType string, vc *vcard.Card) {
	if iaddr == nil {
		return
	}

	// return if every value is empty
	if len(ptr.Val(iaddr.GetStreet())) == 0 &&
		len(ptr.Val(iaddr.GetCity())) == 0 &&
		len(ptr.Val(iaddr.GetState())) == 0 &&
		len(ptr.Val(iaddr.GetPostalCode())) == 0 &&
		len(ptr.Val(iaddr.GetCountryOrRegion())) == 0 {
		return
	}

	addr := vcard.Address{
		StreetAddress: ptr.Val(iaddr.GetStreet()),
		Locality:      ptr.Val(iaddr.GetCity()),
		Region:        ptr.Val(iaddr.GetState()),
		PostalCode:    ptr.Val(iaddr.GetPostalCode()),
		Country:       ptr.Val(iaddr.GetCountryOrRegion()),
	}

	if len(addrType) > 0 {
		addr.Field = &vcard.Field{}
		addr.Params = vcard.Params{"TYPE": []string{addrType}}
	}

	vc.AddAddress(&addr)
}

func addPhones(phones []string, phoneType string, vc *vcard.Card) {
	for _, phone := range phones {
		vc.Add(
			vcard.FieldTelephone,
			&vcard.Field{Value: phone, Params: vcard.Params{"TYPE": []string{phoneType}}})
	}
}

func addEmails(emails []models.EmailAddressable, vc *vcard.Card) {
	for _, email := range emails {
		etype, _ := email.GetBackingStore().Get("type")
		if etype == "unknown" {
			etype = nil
		}

		if etype != nil {
			vc.Add(
				vcard.FieldEmail,
				&vcard.Field{
					Value:  ptr.Val(email.GetAddress()),
					Params: vcard.Params{"TYPE": []string{etype.(string)}},
				})
		} else {
			vc.Add(
				vcard.FieldEmail,
				&vcard.Field{Value: ptr.Val(email.GetAddress())})
		}
	}
}

func FromJSON(ctx context.Context, body []byte) (string, error) {
	vc := vcard.Card{}
	vcard.ToV4(vc)

	data, err := api.BytesToContactable(body)
	if err != nil {
		return "", clues.Wrap(err, "converting to contactable")
	}

	name := vcard.Name{
		GivenName:       ptr.Val(data.GetGivenName()),
		FamilyName:      ptr.Val(data.GetSurname()),
		AdditionalName:  ptr.Val(data.GetMiddleName()),
		HonorificPrefix: ptr.Val(data.GetTitle()),
		HonorificSuffix: ptr.Val(data.GetGeneration()),
	}
	vc.SetName(&name)

	nick := data.GetNickName()
	if nick != nil {
		vc.Set(vcard.FieldNickname, &vcard.Field{Value: *nick})
	}

	bday := data.GetBirthday()
	if bday != nil {
		vc.Set(vcard.FieldBirthday, &vcard.Field{Value: bday.Format("2006-01-02")})
	}

	addAddress(data.GetHomeAddress(), vcard.TypeHome, &vc)
	addAddress(data.GetBusinessAddress(), vcard.TypeWork, &vc)
	addAddress(data.GetOtherAddress(), "other", &vc)

	mob := data.GetMobilePhone()
	if mob != nil && len(ptr.Val(mob)) > 0 {
		addPhones([]string{*mob}, vcard.TypeCell, &vc)
	}

	addPhones(data.GetBusinessPhones(), vcard.TypeWork, &vc)
	addPhones(data.GetHomePhones(), vcard.TypeHome, &vc)

	addEmails(data.GetEmailAddresses(), &vc) // no type?

	im := data.GetImAddresses()
	for _, imaddr := range im {
		vc.Add(vcard.FieldIMPP, &vcard.Field{Value: imaddr})
	}

	orgFull := ""

	org := data.GetCompanyName()
	if org != nil {
		orgFull = *org
	}

	dept := data.GetDepartment()
	if dept != nil {
		if len(orgFull) > 0 {
			orgFull += ";"
		}

		orgFull += *dept
	}

	profession := data.GetProfession()
	if profession != nil {
		if len(orgFull) > 0 {
			orgFull += ";"
		}

		orgFull += *profession
	}

	if len(orgFull) > 0 {
		vc.Set(vcard.FieldOrganization, &vcard.Field{Value: orgFull})
	}

	job := data.GetJobTitle()
	if job != nil && len(ptr.Val(job)) > 0 {
		vc.Set(vcard.FieldTitle, &vcard.Field{Value: *job})
	}

	children := data.GetChildren()
	for _, child := range children {
		vc.Add(
			vcard.FieldRelated,
			&vcard.Field{Value: child, Params: vcard.Params{"TYPE": []string{vcard.TypeChild}}})
	}

	spouse := data.GetSpouseName()
	if spouse != nil && len(ptr.Val(spouse)) > 0 {
		vc.Add(
			vcard.FieldRelated,
			&vcard.Field{Value: *spouse, Params: vcard.Params{"TYPE": []string{vcard.TypeSpouse}}})
	}

	manager := data.GetManager()
	if manager != nil && len(ptr.Val(manager)) > 0 {
		vc.Add(
			vcard.FieldRelated,
			&vcard.Field{Value: *manager, Params: vcard.Params{"TYPE": []string{"manager"}}})
	}

	assistant := data.GetAssistantName()
	if assistant != nil && len(ptr.Val(assistant)) > 0 {
		vc.Add(
			vcard.FieldRelated,
			&vcard.Field{Value: *assistant, Params: vcard.Params{"TYPE": []string{"assistant"}}})
	}

	notes := data.GetPersonalNotes()
	if notes != nil && len(ptr.Val(notes)) > 0 {
		vc.Set(vcard.FieldNote, &vcard.Field{Value: *notes})
	}

	out := bytes.NewBuffer(nil)
	enc := vcard.NewEncoder(out)

	err = enc.Encode(vc)
	if err != nil {
		return "", clues.Wrap(err, "encoding vcard")
	}

	return out.String(), nil
}
