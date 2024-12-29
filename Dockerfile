# Build static client files
FROM node:22.12 AS client-build-stage

WORKDIR /app
COPY . .
WORKDIR /app/web
RUN npm install
RUN npm run build
COPY ./web/build /tmp/build

# Build server binary
FROM golang:1.23.4 AS server-build-stage

WORKDIR /app
COPY . .
COPY --from=client-build-stage /tmp/build ./web
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/occ-loo-pied-server cmd/main.go

# Deploy server binary
FROM alpine:latest AS build-release-stage

COPY --from=server-build-stage /tmp/occ-loo-pied-server /usr/bin/occ-loo-pied
EXPOSE 3333
ENTRYPOINT ["occ-loo-pied"]