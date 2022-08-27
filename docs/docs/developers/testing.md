# Running tests locally

## Prerequisites

- Set `CORSO_PASSWORD` environment variable

    ```bash
    export CORSO_PASSWORD=<some password>
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

    ```bash
    export TENANT_ID=<tenant>
    export CLIENT_ID=<id>
    export CLIENT_SECRET=<secret>
    ```

## Running tests

Standard `go test ./...` will run unit tests

Integration style tests run when enabled by setting the appropriate environment variable.

For example, `CORSO_CI_TESTS=true go test ./...`

The complete list of environment constants is available at
`.../src/internal/tester/integration_runners.go`.

## Advanced options

- To override the M365 user for tests, use `CORSO_M356_TEST_USER_ID`

    ```bash
    export CORSO_M356_TEST_USER_ID="..."
    ```
