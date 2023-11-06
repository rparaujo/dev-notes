# Linting

```
go install golang.org/x/lint/golint@latest
```

Run linting in the whole project with:
```
golint ./...
```

```
go vet
```

```
golangci-lint run
```
You can configure the linters that `golangci` runs by including a `.golangci.yaml` file in your project.
Check the [documentation](https://golangci-lint.run/usage/linters/) for available linters and their configurations.