# Box

Box is to Go what `crate` is to Rust.

It provides quality of life features missing currenly in `go` binary.

## Features

- Create multiple projects.
- Automatic Makefile generation
- Automatic Extract documentation
- Setup tests.

## Prequisites
- Git 
- Bash
- Go

`box` generates a Makefile. To use it install `make`.

## Usage

```txt
Usage of box:
box [OPTIONS] [NAME]
  -profile string
    	Profile to use while creating a project (default "default")
```

Using `box` is simple.
```sh
$ box <project-name>
```

Or 
```sh
$ box -profile <profile> <project-name>
```

## Profiles

- Simple
- CLI Application

## Roadmap

- [ ] Git Init support
- [ ] Add more types of applications
    - Web API
    - Microservices
- [ ] Integrate subcommands like `cargo` does. Example `box-deb` should have support for making deb files.
- [ ] Provide `release` subcommand. Thus `box release` should compile for multiple architectures.
- [ ] Handle global package management. Provide `box install/remove` commands.
- [ ] Generate `documentation`.

