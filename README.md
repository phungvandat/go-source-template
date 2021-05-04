# BACKEND SOURCE TEMPLATE

## Development setup

```shell
# Init DB, ENV
make init-dev
```

## How to run

```shell
make dev
```

# How to gen docs

- Install go swagger: `go get -u github.com/go-swagger/go-swagger/cmd/swagger`
- Gen docs: `make gen-docs`
- View UI: `make docs-ui`
