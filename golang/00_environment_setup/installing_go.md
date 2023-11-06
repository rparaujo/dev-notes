# Installing Go (Golang)

Below you'll find the steps to install Go on different operating systems. For the most up-to-date installation methods and best practices, refer to the official Go downloads page at [https://go.dev/dl/](https://go.dev/dl/) .

## Mac
If you use a Mac for development, you can install Go using [Homebrew](https://formulae.brew.sh/) using the command `brew install go`

## Windows
If you use Windows for development, you can install Go using [Chocolatey](https://chocolatey.org/) using the command `choco install golang`

## Linux/FreeBSD
For Linux you can you can download the _gzipped_ tar files and expand the the directory _/usr/local/go_.
Add _/usr/local/go/bin_ to your `$PATH` to that the `go` command is available

```
$ tar -C /usr/local -xzf go<version>.linux-amd64.tar.gz
$ echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile
$ source $HOME/.profile
```

To test that the installation was successful you can run:
```
go version
```
You should see something like the following output:
```
go version go<version> darwin/amd64
```

Next Steps
Explore Go packages: Take a look at the plethora of community-contributed packages at https://pkg.go.dev/.
Learn the language: Dive into Go's official documentation and interactive tutorials to start learning the language: https://go.dev/doc/.

<div style="text-align: center;">

[Previous](./pre_requisits.md) | [Home](../README.md#environment-setup) | [Next](./ide.md)

</div>