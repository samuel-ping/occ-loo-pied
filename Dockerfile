# Build static client files
FROM node:22.12 AS client-build-stage

WORKDIR /app
COPY web/package.json .
RUN npm install
COPY ./web .
RUN npm run build

# Build server binary
FROM golang:1.23.4 AS server-build-stage

WORKDIR /app
COPY . .
COPY --from=client-build-stage /app/build ./web/build
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/occ-loo-pied cmd/main.go

# Deploy server binary
FROM alpine:latest AS build-release-stage

COPY --from=server-build-stage /tmp/occ-loo-pied /usr/bin/occ-loo-pied
EXPOSE 3333
ENTRYPOINT ["occ-loo-pied"]