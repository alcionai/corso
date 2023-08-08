# Running tests locally

## Prerequisites

- Set `CORSO_PASSPHRASE` environment variable

    ```bash
    export CORSO_PASSPHRASE=<some password>
    ```

- Set AWS credential (needed for tests that use S3) environment variables

    ```bash
    export AWS_ACCESS_KEY_ID="...."
    export AWS_SECRET_ACCESS_KEY="..."
    export AWS_SESSION_TOKEN="..."
    ```

- Create a config file with the S3 bucket used for testing

    ```toml
    bucket = '<bucket name>'
    ```

- Set `CORSO_TEST_CONFIG_FILE` to use the test config file

    ```bash
    export CORSO_TEST_CONFIG_FILE=~/.corso_test.toml
    ```

- Set M365 Credentials environment variables

    > You can find more information on how to get these values in our [M365 docs](../../setup/m365-access/).

    ```bash
    export AZURE_CLIENT_ID=<id>
    export AZURE_CLIENT_SECRET=<secret>
    export AZURE_TENANT_ID=<tenant>
    ```

## Running tests

Standard `go test ./...` will run unit tests

Integration style tests run when enabled by setting the appropriate environment variable.

For example, `CORSO_CI_TESTS=true go test ./...`

The complete list of environment constants is available at
`.../src/internal/tester/integration_runners.go`.

## Advanced options

- To override the M365 user for tests, use `CORSO_M365_TEST_USER_ID`

    ```bash
    export CORSO_M365_TEST_USER_ID="..."
    ```
