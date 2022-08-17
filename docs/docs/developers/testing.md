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

## Running Tests
Standard `go test ./...` will run unit tests

Integration style tests are configured to run only if enabled by setting the
appropriate ENV variable.

e.g. `CORSO_CI_TESTS=true go test ./...`

The complete list of enviroment constants is [here](src/internal/testing/integration_runners.go)
