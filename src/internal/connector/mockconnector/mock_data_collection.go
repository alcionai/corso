package mockconnector

import (
	"bytes"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
)

// MockExchangeDataCollection represents a mock exchange mailbox
type MockExchangeDataCollection struct {
	fullPath     []string
	messageCount int
	Data         [][]byte
	Names        []string
}

var (
	_ data.Collection = &MockExchangeDataCollection{}
	_ data.Stream     = &MockExchangeData{}
	_ data.StreamInfo = &MockExchangeData{}
)

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated
func NewMockExchangeDataCollection(pathRepresentation []string, numMessagesToReturn int) *MockExchangeDataCollection {
	c := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
	}

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, []byte("test message"))
		c.Names = append(c.Names, uuid.NewString())
	}
	return c
}

func NewMockLargeExchangeCollectionMail(pathRepresentation []string, numMessagesToReturn int) *MockExchangeDataCollection {
	c := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
	}

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, getMockMessageBytes())
		c.Names = append(c.Names, uuid.NewString())
	}
	return c
}

func (medc *MockExchangeDataCollection) FullPath() []string {
	return append([]string{}, medc.fullPath...)
}

// Items returns a channel that has the next items in the collection. The
// channel is closed when there are no more items available.
func (medc *MockExchangeDataCollection) Items() <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)
		for i := 0; i < medc.messageCount; i++ {
			res <- &MockExchangeData{
				medc.Names[i],
				io.NopCloser(bytes.NewReader(medc.Data[i])),
			}
		}
	}()

	return res
}

// ExchangeData represents a single item retrieved from exchange
type MockExchangeData struct {
	ID     string
	Reader io.ReadCloser
}

func (med *MockExchangeData) UUID() string {
	return med.ID
}

func (med *MockExchangeData) ToReader() io.ReadCloser {
	return med.Reader
}

func (med *MockExchangeData) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: &details.ExchangeInfo{Sender: "foo@bar.com", Subject: "Hello world!", Received: time.Now()}}
}

// getMockMessageBytes returns bytes for Messageable item. Details based on
// kiota-json-serialization package
func getMockMessageBytes() []byte {

	message := "{" +
		"\"@odata.context\": \"https://graph.microsoft.com/v1.0/$metadata#users('vincent%40biret365.onmicrosoft.com')/messages\"," +
		"\"@odata.nextLink\": \"https://graph.microsoft.com/v1.0/users/vincent@biret365.onmicrosoft.com/messages?$skip=10\"," +
		"\"value\": [" +
		"{" +
		"\"@odata.etag\": \"W/\\\"CQAAABYAAAAs+XSiyjZdS4Rhtwk0v1pGAAA4Xv0v\\\"\"," +
		"\"id\": \"AAMkAGNmMGZiNjM5LTZmMDgtNGU2OS1iYmUwLWYwZDc4M2ZkOGY1ZQBGAAAAAAAK20ulGawAT7z-yx90ohp-BwAs_XSiyjZdS4Rhtwk0v1pGAAAAAAEMAAAs_XSiyjZdS4Rhtwk0v1pGAAA4dw6TAAA=\"," +
		"\"createdDateTime\": \"2021-10-14T09:19:01Z\"," +
		"\"lastModifiedDateTime\": \"2021-10-14T09:19:03Z\"," +
		"\"changeKey\": \"CQAAABYAAAAs+XSiyjZdS4Rhtwk0v1pGAAA4Xv0v\"," +
		"\"categories\": []," +
		"\"receivedDateTime\": \"2021-10-14T09:19:02Z\"," +
		"\"sentDateTime\": \"2021-10-14T09:18:59Z\"," +
		"\"hasAttachments\": false," +
		"\"internetMessageId\": \"<608fed24166f421aa1e27a6c822074ba-JFBVALKQOJXWILKNK4YVA7CPGM3DKTLFONZWCZ3FINSW45DFOJ6E2ZLTONQWOZKDMVXHIZLSL5GUGMRZGEYDQOD4KNWXI4A=@microsoft.com>\"," +
		"\"subject\": \"Major update from Message center\"," +
		"\"bodyPreview\": \"(Updated) Microsoft 365 Compliance Center Core eDiscovery - Search by ID list retirementMC291088 · BIRET365Updated October 13, 2021: We have updated this message with additional details for clarity.We will be retiring the option to Search by ID,\"," +
		"\"importance\": \"normal\"," +
		"\"parentFolderId\": \"AQMkAGNmMGZiNjM5LTZmMDgtNGU2OS1iYgBlMC1mMGQ3ODNmZDhmNWUALgAAAwrbS6UZrABPvP-LH3SiGn8BACz5dKLKNl1LhGG3CTS-WkYAAAIBDAAAAA==\"," +
		"\"conversationId\": \"AAQkAGNmMGZiNjM5LTZmMDgtNGU2OS1iYmUwLWYwZDc4M2ZkOGY1ZQAQANari86tqeZDsqpmA19AXLQ=\"," +
		"\"conversationIndex\": \"AQHXwNyG1quLzq2p5kOyqmYDX0BctA==\"," +
		"\"isDeliveryReceiptRequested\": null," +
		"\"isReadReceiptRequested\": false," +
		"\"isRead\": false," +
		"\"isDraft\": false," +
		"\"webLink\": \"https://outlook.office365.com/owa/?ItemID=AAMkAGNmMGZiNjM5LTZmMDgtNGU2OS1iYmUwLWYwZDc4M2ZkOGY1ZQBGAAAAAAAK20ulGawAT7z%2Fyx90ohp%2FBwAs%2BXSiyjZdS4Rhtwk0v1pGAAAAAAEMAAAs%2BXSiyjZdS4Rhtwk0v1pGAAA4dw6TAAA%3D&exvsurl=1&viewmodel=ReadMessageItem\"," +
		"\"inferenceClassification\": \"other\"," +
		"\"body\": {" +
		"\"contentType\": \"html\"," +
		"\"content\": \"<html><head><meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><meta name=\\\"viewport\\\" content=\\\"width=device-width, initial-scale=1\\\"><meta content=\\\"IE=edge\\\"><style><!--body, table, td{font-family:Segoe UI,Helvetica,Arial,sans-serif!important}a{color:#006CBE;text-decoration:none}--></style></head><body><div style=\\\"background:white; min-height:100vh; color:#323130; font-size:14px\\\"><table border=\\\"0\\\" cellpadding=\\\"0\\\" cellspacing=\\\"0\\\" width=\\\"100%\\\" height=\\\"100%\\\"><tbody><tr><td></td><td width=\\\"640\\\"><table border=\\\"0\\\" cellpadding=\\\"0\\\" cellspacing=\\\"0\\\" style=\\\"min-width:100%; background:white\\\"><tbody><tr><td style=\\\"padding:24px 24px 45px\\\"><img src=\\\"https://eus-contentstorage.osi.office.net/images/retailer.images/centralizeddeployment/logos/112fec798b78aa02.png\\\" width=\\\"100\\\" height=\\\"21\\\" alt=\\\"Microsoft\\\"> </td></tr><tr><td style=\\\"font-size:28px; padding:0 24px; font-weight:bold; color:#000000\\\">(Updated) Microsoft 365 Compliance Center Core eDiscovery - Search by ID list retirement</td></tr><tr><td style=\\\"color:#323130; padding:20px 24px 40px 24px\\\"><span style=\\\"font-weight:600\\\">MC291088 · BIRET365</span></td></tr><tr><td style=\\\"padding:0 24px 44px\\\"><div><p style=\\\"margin-top:0\\\">Updated October 13, 2021: We have updated this message with additional details for clarity.</p><p>We will be retiring the option to Search by ID list, as it is not functioning to an adequate level and creates significant challenges for organizations who depend on consistent and repeatable results for eDiscovery workflows.<br></p><p><b style=\\\"font-weight:600\\\">When will this happen:</b></p><p>We will begin making this change in mid-November and expect to complete by the end of November.</p><p><b style=\\\"font-weight:600\\\">How this will affect your organization:</b><br></p><p>You are receiving this message because our reporting indicates your organization may be using Search by ID list.</p><p>Once this change is made, the option to Search by ID list will be removed. We suggest focusing on search by query, condition and/or locations rather that ID.</p><p><b style=\\\"font-weight:600\\\">What you need to do to prepare:</b><br></p><p>To fix this problem you need to review your eDiscovery search process, and update the workflow to focus on search by Subjects and dates rather than Search by ID list. Upon export from Core eDiscovery you can explore options to refine to only the messages of interest. </p><p></p><p>Click Additional Information to find out more.<br></p><a href=\\\"https://docs.microsoft.com/microsoft-365/compliance/search-for-content-in-core-ediscovery?view=o365-worldwide\\\" title=\\\"Additional Information\\\">Additional Information</a> </div><div style=\\\"padding-top:3px\\\"><a href=\\\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC291088?MCLinkSource=MajorUpdate\\\" title=\\\"view message\\\" target=\\\"_blank\\\">View this message in the Microsoft 365 admin center</a> </div></td></tr><tr><td><table border=\\\"0\\\" cellpadding=\\\"0\\\" cellspacing=\\\"0\\\" width=\\\"100%\\\" style=\\\"min-width:100%; background-color:#F3F2F1\\\"><tbody><tr><td style=\\\"padding:44px 24px 3px; font-size:10px; color:#484644\\\">You're subscribed to this email using vincent@biret365.onmicrosoft.com. If you're an IT admin, you're subscribed by default, but you can <a href=\\\"https://admin.microsoft.com/adminportal/home#/MessageCenter/:/mcpreferences\\\" target=\\\"_blank\\\">unsubscribe at any time</a>. If you're not an IT admin, ask your admin to remove your email address from Microsoft 365 message center preferences.<br><br><a href=\\\"https://docs.microsoft.com/en-us/microsoft-365/admin/manage/language-translation-for-message-center-posts?view=o365-worldwide\\\" target=\\\"_blank\\\">How to view translated messages</a><br></td></tr><tr><td style=\\\"padding:25px 24px 24px; font-size:12px\\\"><div style=\\\"color:#696969\\\">This is a mandatory service communication. To set your contact preferences or to unsubcribe from other communications, visit the <a href=\\\"https://go.microsoft.com/fwlink/?LinkId=243189\\\" target=\\\"_blank\\\" style=\\\"color:#696969; text-decoration:underline; text-decoration-color:#696969\\\">Promotional Communications Manager</a>. <a href=\\\"https://go.microsoft.com/fwlink/?LinkId=521839\\\" target=\\\"_blank\\\" style=\\\"color:#696969; text-decoration:underline; text-decoration-color:#696969\\\">Privacy statement</a>. <br><br>Il s’agit de communications obligatoires. Pour configurer vos préférences de contact pour d’autres communications, accédez au <a href=\\\"https://go.microsoft.com/fwlink/?LinkId=243189\\\" target=\\\"_blank\\\" style=\\\"color:#696969; text-decoration:underline; text-decoration-color:#696969\\\">gestionnaire de communications promotionnelles</a>. <a href=\\\"https://go.microsoft.com/fwlink/?LinkId=521839\\\" target=\\\"_blank\\\" style=\\\"color:#696969; text-decoration:underline; text-decoration-color:#696969\\\">Déclaration de confidentialité</a>. </div><div style=\\\"color:#696969; margin-top:10px; margin-bottom:13px\\\">Microsoft Corporation, One Microsoft Way, Redmond WA 98052 USA</div><img src=\\\"https://eus-contentstorage.osi.office.net/images/retailer.images/centralizeddeployment/logos/112fec798b78aa02.png\\\" width=\\\"94\\\" height=\\\"20\\\" alt=\\\"Microsoft\\\"> </td></tr></tbody></table></td></tr></tbody></table></td><td></td></tr></tbody></table></div><img src=\\\"https://mucp.api.account.microsoft.com/m/v2/v?d=AIAAD2ON6I4P6T45JIHQXRZ6AI7WMQVRDMGBPOFLIPLXZDLYEKNQK44CEBYSPPTPDHET337ASHWG3BMEXD6NQZGTF442DPYPANRAMYRCB5XW3VUZYYL7MXCMJU7NIFJFML3F22PJFGPVVKXDWKRH374HXHZFHRY&amp;i=AIAADOZFMOPSOOEFOUHZD4HWEDARG3W3DMLBKJLS4RUJB6O5L7UJYE5NWIJQFRZTMSB74FMTRBBXRGSZEHD6UYCOLJNM7JTG27THR2WYKQWVGJXJGXJDIRHKWQDFKHWPZPZGXDKOGME5EPT3MJK3LLV7VUODVXG2VLJW5SS6POXQKSQXJWFFBHDP6VMQQEX6MHHWYLSUJG4EPHC4U23LQ7P2IKBLOLB5TTYXB5WQPHDYUDO6WN7BVWK4JGZFE7JOGWQTWAGYP7NKV7L3W3XV2W2E7NOXLUQ\\\" width=\\\"1\\\" height=\\\"1\\\" tabindex=\\\"-1\\\" aria-hidden=\\\"true\\\" alt=\\\"\\\"> </body></html>\"" +
		"}," +
		"\"sender\": {" +
		"\"emailAddress\": {" +
		"\"name\": \"Microsoft 365 Message center\"," +
		"\"address\": \"o365mc@microsoft.com\"" +
		"}" +
		"}," +
		"\"from\": {" +
		"\"emailAddress\": {" +
		"\"name\": \"Microsoft 365 Message center\"," +
		"\"address\": \"o365mc@microsoft.com\"" +
		"}" +
		"}," +
		"\"toRecipients\": [" +
		"{" +
		"\"emailAddress\": {" +
		"\"name\": \"Vincent BIRET\"," +
		"\"address\": \"lidiah@8qzvrj.onmicrosoft.com\"" +
		"}" +
		"}" +
		"]," +
		"\"ccRecipients\": []," +
		"\"bccRecipients\": []," +
		"\"replyTo\": []," +
		"\"flag\": {" +
		"\"flagStatus\": \"notFlagged\"" +
		"}" +
		"}" +
		"]" +
		"}"

	return []byte(message)
}
