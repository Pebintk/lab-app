FROM golang:1.25-alpine AS build
WORKDIR /src
ARG VERSION=dev

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-X main.version=${VERSION}" -o /app .

FROM gcr.io/distroless/static-debian12
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
