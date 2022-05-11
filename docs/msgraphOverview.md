# Microsoft Graph SDK for Go

The msgraph-sdk-go is an instrumental technology that is used to transport data
from Microsoft 365 to users. The [Microsoft Graph
SDK](https://github.com/microsoftgraph/msgraph-sdk-go) source code is an open
source. The document includes some information of the nuances of msgraph-sdk.
If a reader desires a more in depth dive into the SDK, one can view the source
code directly or read up on the documentation found
[here](https://docs.microsoft.com/en-us/graph/sdks/sdks-overview)

## Complications

Supporting libraries for msgraph are unstable at present date. The [change
log](https://github.com/microsoftgraph/msgraph-sdk-go/blob/main/CHANGELOG.md)
will speak to the difference in the new release. They do not always contain
instructions on how to navigate those changes.

## Additional Details
For more detailed specifics on how Corso uses msgraph components as it fetches
data from Microsoft 365:
- [Authorization](msgraphAuth.md)
- [Data Transport](msgraphTransport.md)
- [Restore](msgraphRestore.md)

