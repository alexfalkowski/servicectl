FROM golang:1.24.0-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/alexfalkowski/servicectl/internal/cmd.Version=${version}' -extldflags '-static'" -tags netgo -a -o servicectl main.go

FROM gcr.io/distroless/static

WORKDIR /

COPY --from=build /app/servicectl /servicectl
ENTRYPOINT ["/servicectl"]
