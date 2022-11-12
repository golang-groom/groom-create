# Groom Create

This is a subproject of the main project [Groom]().

Handles project creation in Golang with ease.

## Usage

```sh
groom create [PROJECT_URL]
```

For example, creating this project.

```sh
groom create github.com/golang-groom/groom-create
```
This would create a directory `groom-create` and populate it with the default profile.

## Profiles

You can specify the profile type using `-profile`.

Currently the following profiles are available

- basic (`-profile basic`). Very basic golang project template.
- cmd (`-profile cmd`). A cli-based application project template.
- lib (`-profile lib`). A library for golang projects.
