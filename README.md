# Templative

## Install

```sh
go get -u github.com/cwd-k2/templative/cmd/...
```

## Prerequisites

- `git`

## Run

```sh
# tp <owner>/<repo> <target-directory-name>

# Get a template from GitHub, over https.
# Create a project from template

# Cloned to $TEMPLATIVE_DIR/<owner>/<repo>
# .git directory in template ignored

$ tp s10akir/typescript-boilerplate awesome-ts-project

# There we go
```

## Environment Variable

- `TEMPLATIVE_DIR`: Where templates to be stored. (default: `${XDG_DATA_HOME:-$HOME/.local/share}/templative`)
