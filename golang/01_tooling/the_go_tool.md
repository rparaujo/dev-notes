# Go tools

Go comes with batteries included, this means that Go comes with all the development tools you need and a rich standard library.
Out of the box, Go comes with a compiler, linter, formatter, dependency manager, tests runners, among others.

## Create a new Go project
```
go mod init <project_name>
```

## Compile and execute the current project
```
go run .
```
The `go run` command builds your source code and produces a binary, stores the binary in a temporary directory and executes the binary from there. After your code finishes executing the binary is deleted.

## Build the current project
```
go build .
```
This command builds your project and creates an executable file in the current directory.
If you want to change the name of the executable file, you can use the `-o` flag
```
go build -o hello .
```
This will create a `hello` binary if you are in Mac or Linux, or a `hello.exe` if you are in Windows.

## Install go tools
```
go install <path-to-tool>
```
By default the third party Go tools are installed in the default Go workspace in _$HOME/go_ with their source code stored in _$HOME/go/src_ and the compiled binaries in _$HOME/go/bin_.
If you whish to change the default Go workspace, you can do so by setting the `GOPATH` environment variable.

[Previous](../00_environment_setup/ide.md) | [Home](../README.md#environment-setup) | [Next](./formatting.md)

