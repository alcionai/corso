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
avoid failures even after running locally. The current version in use can be
[found](https://github.com/alcionai/corso/blob/main/.github/workflows/lint.yml#L55)
in `.github/worflows/lint.yaml`.

## Running the linter

You can run the linter manually or with the `Makefile` in the repository. Running with
the `Makefile` will also ensure you have the proper version of golangci-lint
installed.

### Running with the `Makefile`

There’s a `Makefile` in the corso/src that will automatically check if the proper
golangci-lint version is installed and run it. This make rule can be run
with `make lint`. If golangci-lint isn't installed locally or the wrong version
is present it will tell you what version it expects along with a link to the
installation page.

### Running manually

You can run golangci-lint manually by executing `golangci-lint run` in the corso/src
directory. This will automatically use corso's `.golangci.yml` configuration.

## Adding exceptions for lint errors

Sometimes the linter will report an issue but it's not something that you can or
should fix. In those cases there are two ways to add a linter exception:

### Single exception via comment

Adding a comment on the line before (or sometimes the offending line) with the
form `//nolint:<linter-name>` will ignore a single error. `<linter-name>` must
be replaced with the name of the linter that produced the report. Note there’s
no space between the `//` and `nolint`. Having a space between the two may
result in the linter still reporting that line.

### Global exception

The `golangci.yml` file has a list of issues that are ignored in the whole
project. These should be as targeted as possible to avoid silencing other lint
errors that aren't related to the one in question. The golangci-lint
[issues configuration page](https://golangci-lint.run/usage/configuration/#issues-configuration)
has some information on this, but it's also useful to look at
[existing exceptions](https://github.com/alcionai/corso/blob/main/src/.golangci.yml)
in the repository under the `issues` section.

The configuration file allows for regex in the text property, so it’s useful to include
the linter/rule that triggered the message. This ensures the lint error is only
ignored for that linter. Combining the linter/rule with the error message text
specific to that error also helps minimize collisions with other lint errors.

## Working with the linters

Some of the enabled linters, like `wsl`, are picky about how code is arranged.
This section provides some tips on how to organize code to reduce lint errors.

### `wsl`

`wsl` is a linter that requires blank lines in parts of the code. It helps make
the codebase uniform and ensures the code doesn't feel too compact.

#### Short-assignments versus var declarations

Go allows declaring and assigning to a variable with either short-assignments
(`x := 42`) or var assignments (`var x = 42`). `wsl` doesn't allow
grouping these two types of variable declarations together. To work around this,
you can convert all your declarations to one type or the other. Converting to
short-assignments only works if the types in question have accessible and
suitable default values.

For example, the mixed set of declarations and assignments:

```go
var err error
x := 42
```

should be changed to the following because using a short-assignment for type
`error` is cumbersome.

```go
var (
    err error
    x = 42
)
```

#### Post-increment and assignments

`wsl` doesn't allow statements before an assignment without a blank line
separating the two. Post-increment operators (e.x. `x++`) count as statements
instead of assignments and may cause `wsl` to report an error. You can avoid
this by moving the post-increment operator to be after the assignment instead of
before it if the assignment doesn't depend on the increment operation.

For example, the snippet:

```go
x++
found = true
```

should be converted to:

```go
found = true
x++
```

#### Functions using recently assigned values

`wsl` allows functions immediately after assignments, but only if the function
uses the assigned value. This requires an ordering for assignments and
function calls.

For example, the following code

```go
a := 7
b := 42
foo(a)
```

should be changed to

```go
b := 42
a := 7
foo(a)
```

If both the second assignment and function call depend on the value of the first
assignment then the assignments and function call must be separated by a blank
line.

```go
a := 7
b := a + 35

foo(a, b)
```

#### Function calls and checking returned error values

One of the other linters expects error checks to follow assignments to the error
variable without blank lines separating the two. One the other hand, `wsl` has
requirements about what statements can be mixed with assignments. To work
around this, you should separate assignments that involve an error from other
assignments. For example

```go
a := 7
b := 42
c, err := foo()
if err != nil {
  ...
}
```

should be changed to

```go
a := 7
b := 42

c, err := foo()
if err != nil {
  ...
}
```

## Common problem linters

Some linter messages aren't clear about what the issue is. Here's common
cryptic messages how you can fix the problems the linters flag.
Each subsection also includes the version of golangci-lint it applies to and the
linter in question.

### `gci` `Expected 's', Found 'a' at file.go`

This applies to golangci-lint v1.45.2 for the `gci` linter and is due to an import
ordering issue. It occurs because imports in the file aren't grouped according
to the import rules for Corso. Corso code should have three distinct import
groups: system imports, third party imports, and imports of other Corso packaged
(see the example below). Typically the cause of a `gci` lint error is a Corso import in the
block for third party libraries.

```go
import (
    "time"

    "github.com/kopia/kopia"

    "github.com/alcionai/corso/pkg/selector"
)
```
