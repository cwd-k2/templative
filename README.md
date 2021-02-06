# Templative

## Install

```sh
go get -u github.com/cwd-k2/templative/cmd/...
```

## Run

```sh
# tp <owner>/<repo> <target-directory-name>

# Get a template from GitHub, over https.
# Create a project from template
$ tp s10akir/typescript-boilerplate awesome-ts-project

# tp yaml <path-to-yaml> <target-directory-name>

# Create a directory structure from yaml
$ tp yaml examples/project-layout.yml my-project
$ exa --tree my-project
 my-project
├──  LICENSE
├──  Makefile
├──  README.md
├──  assets
├──  build
├──  cmd
│  └──  somecommand
│     └──  main.go
├──  configs
├──  docs
├──  examples
├──  internal
│  ├──  app
│  └──  pkg
├──  pkg
├──  scripts
├──  test
└──  tools
```
