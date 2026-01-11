# mux.dev

Portfolio TUI served over SSH using Bubble Tea and Wish.

## Run locally

```bash
go run ./cmd/cli
```

## Run SSH server

Generate a host key if you don't have one:

```bash
ssh-keygen -t ed25519 -f ssh_host_ed25519_key -N ""
```

Start the server:

```bash
go run ./cmd/ssh
```

Connect:

```bash
ssh localhost
```

## Environment variables

- `MUX_SSH_ADDR` (default `:22`)
- `MUX_SSH_HOST_KEY_PATH` (default `ssh_host_ed25519_key`)
- `MUX_HTTP_PORT` (default `80`)
- `MUX_DOMAIN` (default `mux0.dev`)

## Makefile targets

- `make test`
- `make build`
- `make tidy`
- `make docker-build`
- `make docker-run`

## Docker

Build and run:

```bash
docker build -t mux-ssh .
docker run --rm -p 22:22 -p 80:80 mux-ssh
```
