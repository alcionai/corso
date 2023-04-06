# Graph SDK Powershell Troubleshooter

In certain cases, troubleshooting would be significantly simplified if a Corso
user had a simple mechanism to execute targeted MS Graph API commands against
their environment.

One convenient mechanism to accomplish this without going down to the level of
wrapping individual Graph API calls is to use the
[Microsoft Graph PowerShell](https://learn.microsoft.com/en-us/powershell/microsoftgraph/overview?view=graph-powershell-1.0).
It provides a convenient wrapper and great coverage of the API surface.

## Build container

Before using the tool you want to build the container that packages it.

```sh
docker build -t corso/graph_pwsh:latest .
```

## Prerequisites

### Docker

You need to have Docker installed on your system.

### Azure AD app credentials

The tool uses your existing Corso app to make Graph calls and for authentication
you want `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` to be
set as environment variables. You can read more about this [here](https://corsobackup.io/docs/setup/m365-access/).
You will then pass these into the container run so that authentication can be completed.

## Using the tool

### Interactive use

This is suitable if you would like to issue a number of MS Graph API commands from an
interactive shell in the container.

```sh
docker run --rm -it -v $(pwd):/usr/pwsh -e AZURE_TENANT_ID -e AZURE_CLIENT_ID -e AZURE_CLIENT_SECRET corso/graph_pwsh pwsh
```

Alternatively you can use an environment variable file `env_names` that has the names of the required environment variables

```sh
docker run --rm -it -v $(pwd):/usr/pwsh --env-file env_names corso/graph_pwsh pwsh
```

Before you run any command you want to authenticate with Graph using a convenient script
that will create a connection using the default permissions granted to the app.

```powershell
PS> ./Auth-Graph.ps1
```

If you know what you are doing feel free to use `Connect-MgGraph` directly.

### Specific command use

Suitable when you want to run just a single command. Essentially running the `Auth-Graph.ps1`
before the actual command you want to run.

```sh
docker run --rm -it -v $(pwd):/usr/pwsh --env-file env_names corso/graph_pwsh \
       pwsh -c "<your Graph command>"
```

Here is a complete example to get all users

```sh
# This is the equivalent of GET https://graph.microsoft.com/v1.0/users
docker run --rm -it -v $(pwd):/usr/pwsh --env-file env_names corso/graph_pwsh \
       pwsh -c "Get-MgUser -All"
```

Another example to retrieve an email message for a given user by ID.

```sh
# This is the equivalent of GET https://graph.microsoft.com/v1.0/<userID>/messages/<messageId>
docker run --rm -it -v $(pwd):/usr/pwsh --env-file env_names corso/graph_pwsh \
       pwsh -c "Get-MgUserMessage -UserId <userID or UPN> -MessageID <messageID>"
```

## Debug output

To see the requests and responses made by the specific Graph PowerShell commands, add `-Debug` to you command,
similar to the example below.

```sh
# This is the equivalent of GET https://graph.microsoft.com/v1.0/users
docker run --rm -it -v $(pwd):/usr/pwsh --env-file env_names corso/graph_pwsh \
       pwsh -c "Get-MgUser -All -Debug"
```

## Using Beta API calls

In order to use the Beta Graph API, make sure you have done `export MSGRAPH_USE_BETA=1`
before running the container and pass the environment variable in.

Alternatively you can do the following:

```sh
# This is the equivalent of GET https://graph.microsoft.com/v1.0/users
docker run --rm -it -v $(pwd):/usr/pwsh --env-file env_names corso/graph_pwsh \
       pwsh -c "Select-MgProfile -Name "beta" && Get-MgUser -All"
```

## Graph PowerShell reference

To learn about specific commands, see the
[Graph PowerShell Reference](https://learn.microsoft.com/en-us/powershell/microsoftgraph/get-started?view=graph-powershell-1.0)
