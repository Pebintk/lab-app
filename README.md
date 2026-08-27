# lab-app

Minimal Go HTTP service used as the CI pipeline guinea pig for the lab. Its only job is to
prove which image is running.

## Endpoints

| Path | Returns |
|---|---|
| `/healthz` | `ok`, status 200 — liveness |
| `/version` | the build version and the hostname |

## Run locally

```bash
go test ./...
go run .
curl localhost:8080/version   # version=dev hostname=...
```

## Build the image by hand

```bash
docker build --build-arg VERSION=test -t hello:test .
docker run --rm -p 8080:8080 hello:test
curl localhost:8080/version   # version=test hostname=...
```

The runtime image is `gcr.io/distroless/static-debian12` — no shell, no package manager.
If you ever genuinely need to poke around inside the container, use the corresponding
`:debug` variant of the distroless base image, which includes a busybox shell.

## Image path

```
us-central1-docker.pkg.dev/ops-lab-506804/lab-images/hello
```

## Versioning

`version` is a package-level variable compiled into the binary via `-ldflags
"-X main.version=<tag>"` at build time — not read from an environment variable. This keeps
the version and the image that serves it inseparable: `curl /version` on a running instance
tells you exactly which commit's image is deployed, which only works if the tag can't drift
independently of the binary inside it.
