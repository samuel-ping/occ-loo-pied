# Build server binary
FROM golang:1.23.4 AS build-stage

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/occ-loo-pied cmd/main.go

# Deploy server binary
FROM alpine:latest AS build-release-stage

COPY --from=build-stage /tmp/occ-loo-pied /usr/bin/occ-loo-pied
EXPOSE 3333
ENTRYPOINT ["occ-loo-pied"]