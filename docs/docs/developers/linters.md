# Corso linters

Corso uses the golangci-lint GitHub action to run linters on every PR to `main`.
This helps reduce the cognitive load on reviewers and can lead to better
reviewer comments as they don’t get caught up with formatting issues.

## Installing golangci-lint locally

You can install a local version of the linter Corso uses on any platform and
there's also a docker container available. Instructions for installation are
available on the golangci-lint
[website](https://golangci-lint.run/usage/install/#local-installation). The
version that you install should match the version the GitHub workflow uses to
avoid failures even after running locally. The current version in use is
[denoted](https://github.com/alcionai/corso/blob/main/.github/workflows/lint.yml#L55)
in `.github/worflows/lint.yaml`.

## Running the linter

You can run the linter manually or with the Makefile in the repo. Running with
the Makefile will also ensure you have the proper version of golangci-lint
installed.

### Running with the Makefile

There’s a Makefile in the repo that will automatically check if the proper
golangci-lint version is installed and run it. This make rule can be run
with `make lint`. If golangci-lint isn't installed locally or the wrong version
is present it will tell you what version it expects and give a link to the
installation page.

### Running manually

You can run golangci-lint manually by executing `golangci-lint run` in the Corso
code directory. It will automatically use the `.golangci.yml` config file so it
executes with the same settings as the GitHub action.

## Adding exceptions for lint errors

Sometimes the linter will report an issue but it's not something that can or
should be fixed. In those cases there are two ways to add a linter exception.

### Single exception via comment

Adding a comment on the line before (or sometimes the offending line) with the
form `//nolint:<linter-name>` will ignore a single error. `<linter-name>` must
be replaced with the name of the linter that produced the report. Note there’s
no space between the `//` and `nolint`. Having a space between the two may
result in the linter still reporting that line.

### Global exception

The `golangci.yml` file has a list of issues that are ignored in the whole
project. These should be as targeted as possible to avoid silencing other lint
errors that are't related to the one in question. The golangci-lint
[issues configuration page](https://golangci-lint.run/usage/configuration/#issues-configuration)
has some information on this, but it's also useful to look at
[existing exceptions](https://github.com/alcionai/corso/blob/main/src/.golangci.yml)
in the repo under the `issues` section.

The config file allows for regex in the text property, so it’s useful to include
the linter/rule that triggered the message. This ensures the lint error is only
ignored for that linter. Combining the linter/rule with the error message text
specific to that error also helps avoid ignoring other lint errors.

## Interpreting linter output

Some linters have output messages that don't make clear what the issue is. The
following subsections give the version of golangci-lint that they apply to, the
linter in question, and give guidance on interpreting lint messages.

### gci `Expected 's', Found 'a' at file.go`

This applies to golangci-lint v1.45.2 for the gci linter and is due to an import
ordering issue. It occurs because imports in the file are't grouped according
to the import rules for Corso. Corso code should have three distinct import
groups, system imports, third party imports, and imports of other Corso code
like below. The most likely cause of a gci lint error is a Corso import in the
block for third party libraries.

```go
import (
    "time"

    "github.com/kopia/kopia"

    "github.com/alcionai/corso/pkg/selector"
)
